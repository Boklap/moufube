package handleratelimit

import (
	"context"

	"moufube.com/m/internal/config"
	"moufube.com/m/internal/modules/identity"
)

type Service struct {
	writer identity.Writer
	reader identity.Reader
	cfg    *config.Config
}

func NewService(
	writer identity.Writer,
	reader identity.Reader,
	cfg *config.Config,
) *Service {
	return &Service{
		writer: writer,
		reader: reader,
		cfg:    cfg,
	}
}

func (s *Service) VisitorIDRateLimit(ctx context.Context, visitorID string) (int, bool, error) {
	count, err := s.writer.SetVisitorIDRateLimit(ctx, visitorID)
	if err != nil {
		return count, false, err
	}

	shouldBlock := count > s.cfg.RLVisitorMax

	return count, shouldBlock, nil
}
