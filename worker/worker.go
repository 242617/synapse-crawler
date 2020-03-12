package worker

import (
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/242617/synapse-crawler/jobs/ping"
	"github.com/242617/synapse-crawler/log"
)

var logger log.Logger

func Init(base log.Logger) error {
	logger = base

	go func() {
		for {

			job := ping.NewJob()
			_, err := job.Do()
			if err != nil {
				logger.Warn().
					Err(err).
					Msg("cannot ping core")

				sentry.CaptureException(err)
				defer sentry.Flush(5 * time.Second)

				time.Sleep(10 * time.Second)
				continue
			}

			time.Sleep(5 * time.Second)

		}
	}()

	return nil

}
