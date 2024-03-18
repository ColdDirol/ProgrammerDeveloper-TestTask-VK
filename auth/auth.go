package auth

import (
	"ProgrammerDeveloper-TestTask-VK/auth/jwt"
	"ProgrammerDeveloper-TestTask-VK/auth/middleware"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"log/slog"
)

func InitAuth(AuthConfig *utils.Auth) {
	jwt.SecretKey = []byte(AuthConfig.SecretKey)
	jwt.Salt = AuthConfig.Salt
	switch AuthConfig.SecureMode {
	case middleware.Secure:
		jwt.SecureMode = middleware.Secure
	case middleware.Unsecure:
		jwt.SecureMode = middleware.Unsecure
	default:
		utils.LOG.Error("failed to setup SecureMode", slog.String("secure mode:", AuthConfig.SecureMode))
	}

	initRegistrationHandlers()
	initLoginHandlers()
}
