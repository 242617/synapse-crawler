package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"

	"github.com/242617/synapse-crawler/config"
	"github.com/242617/synapse-crawler/log"
	"github.com/242617/synapse-crawler/version"
	"github.com/242617/synapse-crawler/worker"
)

var (
	configFile = flag.String("config", "config.yaml", "Application config path")
)

func main() {
	flag.Parse()

	err := config.Init(*configFile)
	if err != nil {
		fmt.Println(errors.Wrap(err, "cannot init config"))
		os.Exit(1)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:         config.Cfg.Services.Sentry.DSN,
		Environment: version.Environment,
	})
	if err != nil {
		fmt.Println(errors.Wrap(err, "cannot init sentry"))
		os.Exit(1)
	}

	base, err := log.Create()
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(5 * time.Second)
		fmt.Println(errors.Wrap(err, "cannot create logger"))
		os.Exit(1)
	}

	base.
		Info().
		Str("environment", version.Environment).
		Msgf("start %s", version.Application)

	err = worker.Init(base.With().Str("unit", "worker").Logger())
	if err != nil {
		sentry.CaptureException(err)
		defer sentry.Flush(5 * time.Second)
		base.Error().
			Err(err).
			Msg("cannot init worker")
		os.Exit(1)
	}

	select {}

}
