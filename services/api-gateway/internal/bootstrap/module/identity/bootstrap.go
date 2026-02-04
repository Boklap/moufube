package identity

import (
	"github.com/redis/go-redis/v9"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/infrastructure/repository/identity/reader"
	"moufube.com/m/internal/infrastructure/repository/identity/writer"
	handleratelimit "moufube.com/m/internal/modules/identity/handle_rate_limit"
)

type Module struct {
	HandleVisitor   *HandleVisitorFeature
	HandleRateLimit *HandleRateLimitFeature
}

type HandleRateLimitFeature struct {
	Service    *handleratelimit.Service
	Controller *handleratelimit.Controller
}

func BootstrapingIdentityModule(
	rdb *redis.Client,
	cfg *config.Config,
	logger *logger.AppLogger,
) *Module {
	reader := reader.NewIdentityReaderImpl(rdb)
	writer := writer.NewIdentityWriterImpl(rdb, cfg)

	handleVisitor := bootstrapingHandleFisitorFeature(reader, writer, cfg, logger)
	handleRateLimit := bootstrapingHandleRateLimitFeature(reader, writer, cfg, logger)

	return &Module{
		HandleVisitor:   handleVisitor,
		HandleRateLimit: handleRateLimit,
	}
}
