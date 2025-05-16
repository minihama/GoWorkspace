package mdware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"xorm.io/xorm"
)

func AuthValidator() middleware.BasicAuthValidator {
	return func(id string, pw string, c echo.Context) (bool, error) {
		if id == "bloomfilter" && pw == "very_long_password" {
			return true, nil
		}
		return false, echo.NewHTTPError(401, "권한이 없습니다.")
	}
}

func DbMiddleware(db *xorm.Engine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
