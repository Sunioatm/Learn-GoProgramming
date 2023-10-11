package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/api", api)

	app.Listen("localhost:8001")
}

func init() {
	hystrix.ConfigureCommand("api", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 1,
		ErrorPercentThreshold:  100,
		SleepWindow:            15000,
	})
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe("localhost:8002", hystrixStreamHandler)

}

func api(c *fiber.Ctx) error {

	output := make(chan string, 1)
	hystrix.Go("api", func() error {
		res, err := http.Get("http://localhost:8000/api")
		if err != nil {
			return err
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		msg := string(data)
		fmt.Println(msg)

		output <- msg

		return nil
	}, func(err error) error {
		fmt.Println(err)
		return nil
	})

	out := <-output

	return c.SendString(out)

}
