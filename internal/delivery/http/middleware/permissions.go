package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/motixo/goat-api/internal/delivery/http/response"
	"github.com/motixo/goat-api/internal/domain/entity"
	"github.com/motixo/goat-api/internal/domain/service"
	"github.com/motixo/goat-api/internal/domain/usecase/user"
	"github.com/motixo/goat-api/internal/domain/valueobject"
)

type PermMiddleware struct {
	userUC    user.UseCase
	permCache service.PermCacheService
	userCache service.UserCacheService
}

func NewPermMiddleware(userUC user.UseCase, permCache service.PermCacheService, userCache service.UserCacheService) *PermMiddleware {
	return &PermMiddleware{
		userUC:    userUC,
		permCache: permCache,
		userCache: userCache,
	}
}

func (m *PermMiddleware) Require(requiredPerm valueobject.Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDVal, exists := c.Get(string(UserIDKey))
		if !exists {
			response.Unauthorized(c, "authentication required")
			c.Abort()
			return
		}
		userID, ok := userIDVal.(string)
		if !ok || userID == "" {
			response.Unauthorized(c, "invalid user context")
			c.Abort()
			return
		}

		roleID, err := m.userCache.GetUserRole(c.Request.Context(), userID)
		if err != nil {
			response.Internal(c)
			c.Abort()
			return
		}
		if roleID == valueobject.RoleUnknown {
			response.Unauthorized(c, "someting went wrong, contact support.")
			c.Abort()
			return
		}

		perms, err := m.permCache.GetRolePermissions(c.Request.Context(), roleID)
		if err != nil {
			response.Internal(c)
			c.Abort()
			return
		}

		if !hasPermission(perms, requiredPerm) {
			response.Forbidden(c, "insufficient permissions")
			c.Abort()
			return
		}

		c.Next()
	}
}

func hasPermission(perms []*entity.Permission, required valueobject.Permission) bool {
	requiredStr := string(required)
	fullAccessStr := string(valueobject.PermFullAccess)

	for _, p := range perms {
		if p.Action == requiredStr || p.Action == fullAccessStr {
			return true
		}
	}
	return false
}
