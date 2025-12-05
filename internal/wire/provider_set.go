// internal/wire/provider_set.go
package wire

import (
	"time"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"

	"github.com/motixo/goth-api/internal/config"
	"github.com/motixo/goth-api/internal/delivery/http"
	"github.com/motixo/goth-api/internal/delivery/http/handlers"
	"github.com/motixo/goth-api/internal/delivery/http/middleware"

	// Domain layer
	"github.com/motixo/goth-api/internal/domain/repository"
	"github.com/motixo/goth-api/internal/domain/service"
	"github.com/motixo/goth-api/internal/domain/usecase/auth"
	"github.com/motixo/goth-api/internal/domain/usecase/permission"
	"github.com/motixo/goth-api/internal/domain/usecase/session"
	"github.com/motixo/goth-api/internal/domain/usecase/user"

	// Infrastructure layer
	authInfra "github.com/motixo/goth-api/internal/infrastructure/auth"
	permissionCache "github.com/motixo/goth-api/internal/infrastructure/cache/permission"
	"github.com/motixo/goth-api/internal/infrastructure/database/postgres"
	postgresPermission "github.com/motixo/goth-api/internal/infrastructure/database/postgres/permission"
	postgresUser "github.com/motixo/goth-api/internal/infrastructure/database/postgres/user"
	redisSession "github.com/motixo/goth-api/internal/infrastructure/storage/redis/session"
	"github.com/motixo/goth-api/internal/shared"
)

// Infrastructure providers
var InfrastructureSet = wire.NewSet(
	config.Load,
	ProvideRedisClient,
	postgres.NewDatabase,
)

// Repository providers
var RepositorySet = wire.NewSet(
	postgresUser.NewRepository,
	postgresPermission.NewRepository,
	redisSession.NewRepository,
)

// Service providers
var ServiceSet = wire.NewSet(
	service.NewULIDGenerator,
	service.NewPasswordService,
	NewJWTManager,
	NewZapLogger,
	wire.Bind(new(service.JWTService), new(*authInfra.JWTManager)),
	wire.Bind(new(service.Logger), new(*shared.ZapLogger)),
)

// Configuration providers
var ConfigSet = wire.NewSet(
	ProvideAccessTTL,
	ProvideRefreshTTL,
	ProvideSessionTTL,
	ProvidePermissionCache,
	ProvideCachedPermissionRepository,
)

// UseCase providers - Wire will automatically use your constructor!
var UseCaseSet = wire.NewSet(
	auth.NewUsecase,    // Wire will match parameters by type
	session.NewUsecase, // You'll need to fix session constructor too
	user.NewUsecase,
	permission.NewUsecase,
)

// HTTP providers
var HTTPSet = wire.NewSet(
	handlers.NewAuthHandler,
	handlers.NewUserHandler,
	handlers.NewSessionHandler,
	middleware.NewAuthMiddleware,
	middleware.NewPermMiddleware,
	http.NewServer,
)

// ProviderSet bundles everything
var ProviderSet = wire.NewSet(
	InfrastructureSet,
	RepositorySet,
	ServiceSet,
	ConfigSet,
	UseCaseSet,
	HTTPSet,
)

// Infrastructure providers
func ProvideRedisClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})
}

// Configuration providers with unique types
func ProvideAccessTTL(cfg *config.Config) auth.AccessTTL {
	return auth.AccessTTL(cfg.JWTExpiration)
}

func ProvideRefreshTTL(cfg *config.Config) auth.RefreshTTL {
	return auth.RefreshTTL(cfg.RefreshTokenExpiration)
}

func ProvideSessionTTL(cfg *config.Config) auth.SessionTTL {
	return auth.SessionTTL(cfg.SessionExpiration)
}

// Service providers
func NewJWTManager(cfg *config.Config) *authInfra.JWTManager {
	return authInfra.NewJWTManager(cfg.JWTSecret)
}

func NewZapLogger() *shared.ZapLogger {
	return shared.NewZapLogger()
}

func ProvidePermissionCache(rdb *redis.Client) *permissionCache.Cache {
	return permissionCache.NewCache(rdb, 5*time.Minute)
}

func ProvideCachedPermissionRepository(
	dbRepo *postgresPermission.Repository,
	cache *permissionCache.Cache,
	logger service.Logger,
) repository.PermissionRepository {
	return permissionCache.NewCachedRepository(dbRepo, cache, logger)
}
