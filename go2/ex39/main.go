package main

import (
	"ex39/pkg/queue"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func main() {
	q := queue.New()
	e := echo.New()
	e.PUT("/create-order", func(c echo.Context) error {
		order := new(queue.Order)
		if err := c.Bind(order); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := q.Enqueue(order); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "OK")
	})
	go func() {
		for {
			if val, err := q.Dequeue(); err == nil {
				spew.Dump(val)
			}
			time.Sleep(time.Second)
		}
	}()
	e.Logger.Fatal(e.Start(":1234"))
}
