package log

import (
	"io"
	l "log"
	"os"

	console "github.com/mattn/go-isatty"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/242617/synapse-crawler/config"
)

type Logger = zerolog.Logger

func init() { l.SetFlags(l.Lshortfile) }

func Create() (Logger, error) {

	var log Logger
	var w io.Writer
	var o = os.Stderr

	if console.IsTerminal(o.Fd()) {
		w = zerolog.ConsoleWriter{Out: o}
	} else {
		w = o
	}

	level, err := zerolog.ParseLevel(config.Cfg.Logger.Level)
	if err != nil {
		return log, errors.Wrap(err, "failed to parse logging level from config")
	}

	log = zerolog.New(w).With().
		Timestamp().Logger().
		Level(level)

	return log, nil
}
