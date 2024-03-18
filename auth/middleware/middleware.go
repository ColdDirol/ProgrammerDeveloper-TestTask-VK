package middleware

import (
	"ProgrammerDeveloper-TestTask-VK/auth/jwt"
	"ProgrammerDeveloper-TestTask-VK/utils"
	"ProgrammerDeveloper-TestTask-VK/utils/logger"
	"errors"
	"net/http"
)

const (
	user  = "user"
	admin = "admin"
)

const (
	Secure   = "secure"
	Unsecure = "unsecure"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch jwt.SecureMode {
		case Secure:
			token := r.Header.Get("Authorization")

			claims, err := jwt.VerifyToken(token)
			if err != nil {
				utils.LOG.Error("unauthorized user - wrong jwt", logger.Err(err))
				http.Error(w, "unauthorized user - wrong jwt", http.StatusUnauthorized)
				return
			}

			switch claims.Role {
			case user:
				if r.Method != http.MethodGet {
					utils.LOG.Error("not enough rights", logger.Err(errors.New("not enough rights")))
					http.Error(w, "not enough rights", http.StatusForbidden)
				}
			case admin:
			default:
				utils.LOG.Error("not enough rights", logger.Err(err))
				http.Error(w, "not enough rights", http.StatusForbidden)
			}
		case Unsecure:
			// just skip
		}

		next(w, r)
	}
}
