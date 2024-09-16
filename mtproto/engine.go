package mtproto

import (
	"context"
	"time"

	"github.com/go-faster/errors"
)

func (c *Conn) engineClosedLoop(ctx context.Context) error {
	ticker := c.clock.Ticker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "engine closed loop")
		case <-ticker.C():
			if closed := c.rpc.IsClosed(); closed {
				return errors.Wrap(errors.New("engine closed"), "engine closed loop")
			}
		}
	}
}
