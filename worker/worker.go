package worker

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/242617/synapse-core/api"

	"github.com/242617/synapse-crawler/config"
	"github.com/242617/synapse-crawler/log"
)

var logger log.Logger

func Init(base log.Logger) error {
	logger = base

	// Check core online
	go func() {
		for {
			err := ping()
			if err != nil {
				logger.Error().
					Err(err).
					Msg("ping error")
				time.Sleep(10 * time.Second)
				continue
			}
			time.Sleep(5 * time.Second)
		}
	}()

	return nil
}

func ping() error {

	conn, err := grpc.Dial(config.Cfg.Core.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := api.NewSystemClient(conn)
	_, err = client.Info(ctx, &api.Void{})
	if err != nil {
		return err
	}

	return nil

}
