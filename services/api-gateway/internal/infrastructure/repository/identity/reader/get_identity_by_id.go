package reader

import (
	"context"
	"fmt"

	"moufube.com/m/internal/modules/identity"
	"moufube.com/m/internal/modules/identity/apperr"
	"moufube.com/m/internal/modules/identity/constant"
)

func (ir *IdentityReaderImpl) GetIdentityByID(ctx context.Context, visitorID string) (identity.Identity, error) {
	var identityData identity.Identity
	key := fmt.Sprintf("%s:%s", constant.VisitorKey, visitorID)

	val, err := ir.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return identity.Identity{}, err
	}

	if len(val) == 0 {
		return identity.Identity{}, apperr.IdentityNotFound(key)
	}

	return identityData, nil
}
