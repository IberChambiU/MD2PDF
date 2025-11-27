package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	sr "github.com/IberChambiU/MD2PDF/api/navigation"
	"github.com/IberChambiU/MD2PDF/env"
	"github.com/IberChambiU/MD2PDF/worker"
	"github.com/gofiber/fiber/v2"
)

func init() {
	env.Init()
	log.Println("Environment variables loaded.")
}

func main() {

	e := env.Env()

	worker.StartCleanupWorker()

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	sr.ComplementRouter(app)
	sr.ApiRouter(app)
	sr.PublicRouter(app)

	port := ":" + e.GetPort()

	if e.GetSecure() {
		go func() {
			log.Printf(`Running with TLS in https://localhost%v`, port)
			if err := app.ListenTLS(
				port,
				"./certs/127.0.0.1.pem",
				"./certs/127.0.0.1-key.pem",
			); err != nil {
				log.Panic(err)
			}
		}()
	} else {
		go func() {
			log.Printf(`Running in http://localhost%v`, port)
			if err := app.Listen(port); err != nil {
				log.Panic(err)
			}
		}()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	log.Println("Fiber was successful shutdown.")

}
