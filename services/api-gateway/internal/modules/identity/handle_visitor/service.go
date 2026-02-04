package handlevisitor

import (
	"context"

	"moufube.com/m/internal/appctx/strings"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/modules/identity"
)

type Service struct {
	identityReader identity.Reader
	identityWriter identity.Writer
	cfg            *config.Config
}

func NewService(
	identityReader identity.Reader,
	identityWriter identity.Writer,
	cfg *config.Config,
) *Service {
	return &Service{
		identityReader: identityReader,
		identityWriter: identityWriter,
		cfg:            cfg,
	}
}

func (s *Service) CreateNewVisitor(ctx context.Context) (*identity.Identity, error) {
	visitorID, err := strings.GenerateBase64Token(s.cfg.SizeIdentityToken)
	if err != nil {
		return nil, err
	}

	idnty := &identity.Identity{
		ID:              visitorID,
		IsAuthenticated: false,
	}

	err = s.identityWriter.SetIdentity(ctx, visitorID, *idnty)
	if err != nil {
		return nil, err
	}

	return idnty, nil
}

func (s *Service) GetIdentityByID(ctx context.Context, visitorID string) (*identity.Identity, error) {
	idntyMap, err := s.identityReader.GetIdentityByID(ctx, visitorID)
	if err != nil {
		return nil, err
	}

	idnty := identity.MapToIdentity(idntyMap)
	return idnty, nil
}
