package pipeline

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/thingful/zenroom-go"

	kitlog "github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/thingful/iotencoder/pkg/postgres"

	"github.com/thingful/iotencoder/pkg/lua"

	datastore "github.com/thingful/twirp-datastore-go"
)

var (
	// datastoreErrorCounter is a prometheus counter recording a count of any
	// errors that occur when writing to the datastore
	datastoreErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "decode_datastore_errors",
			Help: "Count of errors writing to datastore",
		},
	)

	// zenroomErrorCounter is a prometheus counter recording a count of any errors
	// that occur when invoking zenroom.
	zenroomErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "decode_zenroom_errors",
			Help: "Count of errors invoking zenroom",
		},
	)

	// datastoreWriteHistogram is a prometheus histogram recording successful
	// writes to the datastore. We use the default bucket distributions.
	datastoreWriteHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "decode_datastore_writes",
			Help: "Datastore writes duration distribution",
		},
	)

	// zenroomHistogram is a prometheus histogram recording execution times of
	// calls to zenroom to exec some script.
	zenroomHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "decode_zenroom_exec",
			Help: "Execution time of zenroom scripts",
		},
	)
)

func init() {
	prometheus.MustRegister(datastoreErrorCounter)
	prometheus.MustRegister(datastoreWriteHistogram)
	prometheus.MustRegister(zenroomErrorCounter)
	prometheus.MustRegister(zenroomHistogram)
}

// Processor is an interface we define to handle processing all the streams for
// a device, where processing means reading all streams for the device, applying
// whatever operations that stream specifies in terms of filtering / aggregation
// / bucketing, encrypting the result and then writing the encrypted body to the
// datastore.
type Processor interface {
	// Process takes an input a device which will have one or more attached
	// streams, as well as the received payload from the device. Internally it is
	// responsible for processing the data for each stream and then writing the
	// encrypted result to the remote datastore.
	Process(device *postgres.Device, payload []byte) error
}

// processor is our internal type that implements the above interface
type processor struct {
	datastore datastore.Datastore
	logger    kitlog.Logger
	verbose   bool
}

// NewProcessor is a constructor function that takes as input an instantiated
// datastore client, and a logger. It returns the instantiated processor which
// is ready for use. Note we pass in the datastore instance so that we can
// supply a mock for testing.
func NewProcessor(ds datastore.Datastore, verbose bool, logger kitlog.Logger) Processor {
	logger = kitlog.With(logger, "module", "pipeline")

	return &processor{
		datastore: ds,
		logger:    logger,
		verbose:   verbose,
	}
}

// Process is the function that actually does the work of dispatching the
// received data to all destination streams after applying whatever processing
// the stream specifies. Currently we do the simplest thing of just writing the
// data directly to the datastore.
func (p *processor) Process(device *postgres.Device, payload []byte) error {

	// check payload
	if payload == nil {
		return errors.New("empty payload received")
	}

	// read script from go-bindata asset
	script, err := lua.Asset("encrypt.lua")
	if err != nil {
		return errors.Wrap(err, "failed to read zenroom script")
	}

	for _, stream := range device.Streams {
		if p.verbose {
			p.logger.Log("public_key", stream.PublicKey, "user_uid", device.UserUID, "msg", "writing data")
		}

		start := time.Now()

		encodedPayload, err := zenroom.Exec(
			script,
			[]byte(fmt.Sprintf(`{"public": "%s", "private": "%s" }`, stream.PublicKey, device.PrivateKey)),
			payload,
		)

		duration := time.Since(start)

		if err != nil {
			zenroomErrorCounter.Inc()
			return err
		}

		zenroomHistogram.Observe(duration.Seconds())

		start = time.Now()

		_, err = p.datastore.WriteData(context.Background(), &datastore.WriteRequest{
			PublicKey: stream.PublicKey,
			UserUid:   device.UserUID,
			Data:      encodedPayload,
		})

		duration = time.Since(start)

		if err != nil {
			datastoreErrorCounter.Inc()

			return err
		}

		datastoreWriteHistogram.Observe(duration.Seconds())
	}

	return nil
}
