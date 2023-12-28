package main

import (
	"context"
	"hatchet-go-quickstart/internal/events"
	"net/http"

	"github.com/hatchet-dev/hatchet/pkg/client"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	hatchetClient, err := client.New(
		client.WithHostPort("127.0.0.1", 7077),
	)

	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/test", func(c echo.Context) error {
		err = hatchetClient.Event().Push(
			context.Background(),
			"test-called",
			&events.TestEvent{
				Name: "testing",
			},
		)

		if err != nil {
			panic(err)
		}

		return c.String(http.StatusOK, "OK")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
