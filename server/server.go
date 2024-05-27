package server

import (
	"context"
	"fmt"
	"hatchet-go-quickstart/internal/events"
	"net/http"
	"os"

	"github.com/hatchet-dev/hatchet/pkg/client"
	"github.com/labstack/echo"
)

func Run() error {
	token := os.Getenv("HATCHET_CLIENT_TOKEN")
	if token == "" {
		return fmt.Errorf("HATCHET_CLIENT_TOKEN environment variable is not set")
	}

	hatchetClient, err := client.New(
		client.WithToken(token),
	)

	if err != nil {
		return err
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

	return nil
}
