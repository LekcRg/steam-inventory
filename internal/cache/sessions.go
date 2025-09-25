package cache

import (
	"context"
	"time"
)

const (
	SessionExpiration = 30 * 24 * time.Hour
	SessionPrefix     = "session:"
)

func (c *Cache) SetSession(
	ctx context.Context, session, steamID string,
) error {
	return c.client.Set(
		ctx,
		SessionPrefix+session,
		steamID,
		SessionExpiration,
	).Err()
}

func (c *Cache) DelSession(ctx context.Context, session string) error {
	return c.client.Del(ctx, SessionPrefix+session).Err()
}

func (c *Cache) GetSession(
	ctx context.Context, session string,
) (string, error) {
	return c.client.Get(ctx, SessionPrefix+session).Result()
}
