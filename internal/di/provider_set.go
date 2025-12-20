package di

import (
	"context"
	"reflect"

	"github.com/google/wire"

	"github.com/motixo/goat-api/internal/config"
	"github.com/motixo/goat-api/internal/cron"
	"github.com/motixo/goat-api/internal/pkg"

	// Delivery layer
	"github.com/motixo/goat-api/internal/delivery/http"
	"github.com/motixo/goat-api/internal/delivery/http/handlers"
	"github.com/motixo/goat-api/internal/delivery/http/middleware"

	// Domain layer
	domainEvent "github.com/motixo/goat-api/internal/domain/event"
	"github.com/motixo/goat-api/internal/domain/service"

	// Usecase layer
	"github.com/motixo/goat-api/internal/usecase/auth"
	"github.com/motixo/goat-api/internal/usecase/permission"
	"github.com/motixo/goat-api/internal/usecase/session"
	"github.com/motixo/goat-api/internal/usecase/user"

	// infra layer
	authInfra "github.com/motixo/goat-api/internal/infra/auth"
	permcache "github.com/motixo/goat-api/internal/infra/cache/permission"
	usercache "github.com/motixo/goat-api/internal/infra/cache/user"
	"github.com/motixo/goat-api/internal/infra/database/postgres"
	postgresPermission "github.com/motixo/goat-api/internal/infra/database/postgres/permission"
	postgresUser "github.com/motixo/goat-api/internal/infra/database/postgres/user"
	"github.com/motixo/goat-api/internal/infra/event"
	"github.com/motixo/goat-api/internal/infra/logger"
	"github.com/motixo/goat-api/internal/infra/metrics"
	"github.com/motixo/goat-api/internal/infra/ratelimiter"
	"github.com/motixo/goat-api/internal/infra/storage/redis"
	redisSession "github.com/motixo/goat-api/internal/infra/storage/redis/session"
)

// infra providers
var infraSet = wire.NewSet(
	postgres.NewDatabase,
	redis.NewClient,
	usercache.NewCache,
	permcache.NewCache,
	ProvideConfiguredEventBus,
	wire.Bind(new(domainEvent.Publisher), new(*event.InMemoryPublisher)),
)

// Repository providers
var RepositorySet = wire.NewSet(
	postgresUser.NewRepository,
	postgresPermission.NewRepository,
	redisSession.NewRepository,
	usercache.NewCachedRepository,
	permcache.NewCachedRepository,
)

// Service providers
var ServiceSet = wire.NewSet(
	NewJWTManager,
	wire.Bind(new(service.JWTService), new(*authInfra.JWTManager)),
	metrics.NewPrometheusMetrics,
	wire.Bind(new(service.MetricsService), new(*metrics.PrometheusMetrics)),
	ratelimiter.NewRedisRateLimiter,
	authInfra.NewPasswordService,
)

// Logger providers
var LoggerSet = wire.NewSet(
	logger.NewZapLogger,
	wire.Bind(new(pkg.Logger), new(*logger.ZapLogger)),
)

// Configuration providers
var ConfigSet = wire.NewSet(
	config.Load,
	ProvideAccessTTL,
	ProvideRefreshTTL,
	ProvideSessionTTL,
	ProvideRateLimit,
)

// UseCase provider
var UseCaseSet = wire.NewSet(
	auth.NewUsecase,
	session.NewUsecase,
	user.NewUsecase,
	permission.NewUsecase,
)

// HTTP providers
var HTTPSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewUserHandler,
	handlers.NewSessionHandler,
	handlers.NewPermissionHandler,
	middleware.NewAuthMiddleware,
	middleware.NewPermMiddleware,
	middleware.NewMetricsMiddleware,
	middleware.NewRateLimitMiddleware,
	http.NewServer,
)

// Cron providers
var CronSet = wire.NewSet(
	cron.NewSessionCleaner,
)

// ProviderSet bundles everything
var ProviderSet = wire.NewSet(
	ConfigSet,
	LoggerSet,
	ServiceSet,
	infraSet,
	RepositorySet,
	UseCaseSet,
	HTTPSet,
	CronSet,
)

// Token config
func ProvideAccessTTL(cfg *config.Config) auth.AccessTTL {
	return auth.AccessTTL(cfg.JWTExpiration)
}

func ProvideRefreshTTL(cfg *config.Config) auth.RefreshTTL {
	return auth.RefreshTTL(cfg.RefreshTokenExpiration)
}

func ProvideSessionTTL(cfg *config.Config) auth.SessionTTL {
	return auth.SessionTTL(cfg.SessionExpiration)
}

func NewJWTManager(cfg *config.Config) *authInfra.JWTManager {
	return authInfra.NewJWTManager(cfg.JWTSecret)
}

func ProvideRateLimit(cfg *config.Config) middleware.RateLimitConfig {
	return middleware.RateLimitConfig{
		Auth: middleware.RateLimit{
			Limit:  cfg.RateLimitAuthLimit,
			Window: cfg.RateLimitAuthWindow,
		},
		Public: middleware.RateLimit{
			Limit:  cfg.RateLimitPublicLimit,
			Window: cfg.RateLimitPublicWindow,
		},
		Private: middleware.RateLimit{
			Limit:  cfg.RateLimitPrivateLimit,
			Window: cfg.RateLimitPrivateWindow,
		},
	}
}

// EventBus onfig
func ProvideConfiguredEventBus(
	logger pkg.Logger,
	userCacheRepo service.UserCacheService,
	permCacheRepo service.PermCacheService,
) (*event.InMemoryPublisher, error) {
	bus := event.NewInMemoryPublisher(logger)

	// Register UserUpdatedEvent handler
	bus.RegisterHandler(
		reflect.TypeOf(domainEvent.UserUpdatedEvent{}),
		func(ctx context.Context, e any) error {
			evt, ok := e.(domainEvent.UserUpdatedEvent)
			if !ok {
				return nil // or log
			}
			return userCacheRepo.ClearCache(ctx, evt.UserID)
		},
	)

	// Register PermissionUpdatedEvent handler
	bus.RegisterHandler(
		reflect.TypeOf(domainEvent.PermissionUpdatedEvent{}),
		func(ctx context.Context, e any) error {
			evt, ok := e.(domainEvent.PermissionUpdatedEvent)
			if !ok {
				return nil
			}
			return permCacheRepo.ClearCache(ctx, evt.Role)
		},
	)

	return bus, nil
}
