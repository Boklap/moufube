package writer

import (
	"context"
	"fmt"
	"time"

	"moufube.com/m/internal/modules/identity"
)

func (iw *IdentityWriterImpl) SetVisitorIDRateLimit(ctx context.Context, visitorID string) (int, error) {
	key := fmt.Sprintf("%s:%s:%s", identity.RateLimitKey, identity.VisitorKey, visitorID)
	count, err := iw.rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if count == 1 {
		_ = iw.rdb.Expire(ctx, key, time.Duration(iw.cfg.RLDuration)*time.Second).Err()
	}

	return int(count), nil
}
