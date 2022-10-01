package middleware

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	UserCookieName = "uc"
)

// This currently used only for test purposes to simulate logged-in user
func UserCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		userCookie, err := context.Cookie(UserCookieName)
		if err != nil {
			fmt.Println(err)
		}
		if userCookie != nil && userCookie.Value != "" {
			return next(context)
		}

		context.SetCookie(&http.Cookie{
			Name:    UserCookieName,
			Value:   generateUserCookie(),
			Expires: time.Now().Add(24 * time.Hour),
			Secure:  false,
		})
		return next(context)
	}
}

func generateUserCookie() string {
	b := make([]byte, 20)
	_, _ = rand.Read(b)
	uc := fmt.Sprintf("%x", b)
	fmt.Println("userCookie", uc)
	return uc
}
