package main

import (
	"chucknorris/entities"
	"chucknorris/externalservices"
	server2 "chucknorris/server"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sync"
)

func main() {

	instanceServer := externalservices.NewChuckNorrisServices("https://api.chucknorris.io/jokes/random")

	server := server2.NewServer()
	server.Joker(func(ctx *fiber.Ctx) error {

		c := make(chan entities.ChuckNorrisItem)

		wg := sync.WaitGroup{}

		worker := make(chan int, 10)

		go func() {
			for i := 0; i < 25; i++ {
				worker <- 1
				wg.Add(1)
				go instanceServer.GetItem(c, &wg, worker)
			}
			wg.Wait()
			close(c)
		}()

		response := []entities.ChuckNorrisItem{}

		for item := range c {
			response = append(response, item)
		}

		return ctx.Status(http.StatusOK).JSON(response)

	})
	server.Start()

}
