package reader

import (
	"context"
	"fmt"

	"moufube.com/m/internal/modules/identity"
)

func (ir *IdentityReaderImpl) GetIdentityByID(ctx context.Context, visitorID string) (map[string]string, error) {
	key := fmt.Sprintf("%s:%s", identity.VisitorKey, visitorID)

	identityMap, err := ir.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]string{}, err
	}

	if len(identityMap) == 0 {
		return map[string]string{}, identity.NewIdentityError(
			identity.ErrIdentityNotFound,
			fmt.Sprintf("Identity not found for VisitorID: %s", visitorID),
		)
	}

	return identityMap, nil
}
