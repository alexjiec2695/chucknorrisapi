package server

import "github.com/gofiber/fiber/v2"

type server struct {
	app *fiber.App
}

type Server interface {
	Start() error
	Joker(handler func(c *fiber.Ctx) error)
}

func NewServer() Server {
	app := fiber.New()
	return &server{
		app: app,
	}
}

func (s *server) Start() error {
	err := s.app.Listen(":3000")
	if err != nil {
		return err
	}
	return nil
}

func (s *server) Joker(handler func(c *fiber.Ctx) error) {
	s.app.Get("/api/jokes", handler)
}
