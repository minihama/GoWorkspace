package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/enriquebris/goconcurrentqueue"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Order struct {
		UserId   int       `json:"user_id"`
		OrderId  int       `json:"order_id"`
		Products []Product `json:"products"`
	}
	Product struct {
		ProductId int     `json:"product_id"`
		Price     float64 `json:"price"`
		Quantity  int     `json:"quantity"`
	}
)

func main() {
	queue := goconcurrentqueue.NewFIFO()
	e := echo.New()
	e.PUT("/create-order", func(c echo.Context) error {
		order := new(Order)
		if err := c.Bind(order); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := queue.Enqueue(order); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.String(http.StatusOK, "OK")
	})
	go func() {
		for {
			val, err := queue.DequeueOrWaitForNextElement()
			if err != nil {
				e.Logger.Error(err)
			}
			spew.Dump(val)
		}
	}()
	e.Logger.Fatal(e.Start(":1234"))
}
