package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync/atomic"
)

var counter uint64 = 0

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		newCounter := atomic.AddUint64(&counter, 1)
		if newCounter%3 == 0 {
			return c.String(http.StatusOK, fmt.Sprintf("OK: %d\n", newCounter))
		}
		return c.String(http.StatusInternalServerError, fmt.Sprintf("BAD: %d\n", newCounter))
	})
	e.Logger.Fatal(e.Start(":1234"))
}
