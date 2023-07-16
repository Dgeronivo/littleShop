package apiserver

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sirupsen/logrus"
)

const viewPath = "project/view"
const imgPath = "project/img"
const stylePath = "project/style"

// found and log 404
// catalog page
// product page
// cart
// checkout page
// Договір публічної оферти
// Авторське право
type ApiServer struct {
	config *Config
	logger *logrus.Logger
	app    *fiber.App
}

func New(config *Config) *ApiServer {
	return &ApiServer{
		config: config,
		logger: logrus.New(),
		app:    fiber.New(fiber.Config{Views: createHtmlEngine()}),
	}
}

func (s *ApiServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting server")
	s.configureRouter()

	log.Fatal(s.listen())

	return nil
}

func (s *ApiServer) configureRouter() {
	s.app.Get("/", handleIndex)
	s.app.Get("/delivery", handleDelivery)
	s.app.Get("/about-us", handleAboutUs)
	s.app.Get("/privacy-policy", handlePolicy)
	s.app.Static("/img", imgPath)
	s.app.Static("/style", stylePath)
}

func (s *ApiServer) configureLogger() error {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		log.Fatal(err)
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *ApiServer) listen() error {
	return s.app.Listen(s.config.BindAddr)
}

func createHtmlEngine() *html.Engine {
	return html.New(viewPath, ".html")
}
