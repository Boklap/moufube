package writer

import (
	"context"
	"fmt"

	"moufube.com/m/internal/modules/identity"
	"moufube.com/m/internal/modules/identity/constant"
)

func (iw *IdentityWriterImpl) SetIdentity(ctx context.Context, visitorID string, data identity.Identity) error {
	err := iw.rdb.HSet(ctx, fmt.Sprintf("%s:%s", constant.VisitorKey, visitorID), data).Err()
	if err != nil {
		return err
	}

	return nil
}
