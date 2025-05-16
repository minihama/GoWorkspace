package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"xorm.io/xorm"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func UserFindByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, ok := c.Get("db").(*xorm.Engine)
		if !ok {
			return echo.NewHTTPError(500, "DB 커넥션을 가져올 수 없습니다.")
		}
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return echo.NewHTTPError(400, "잘못된 인자값입니다.")
		}
		user := User{Id: id}
		has, err := db.Table("users").Get(&user)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
		if !has {
			return echo.NewHTTPError(404, "없는 아이디입니다.")
		}
		return c.JSON(200, user)
	}
}

func UserListHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, ok := c.Get("db").(*xorm.Engine)
		if !ok {
			return echo.NewHTTPError(500, "DB 커넥션을 가져올 수 없습니다.")
		}
		name := c.QueryParam("name")
		var users []User
		sess := db.Table("users")
		if len(name) > 0 {
			sess.Where("name like ?", name+"%")
		}
		err := sess.Find(&users)
		if err != nil {
			return echo.NewHTTPError(500, err)
		}
		return c.JSON(200, users)
	}
}
