package main

import (
	"log"

	"github.com/fnunezzz/bff_go/internal/controller"
	infra "github.com/fnunezzz/bff_go/internal/infra/service"
	"github.com/fnunezzz/bff_go/internal/shared/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
	// Server
    e := echo.New()

	// Middlewares
	e.Use(middleware.AddTraceId)
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())
	
	// Prefix
	g := e.Group("/sos-qq")

	// Services
	produtoService := infra.NewProdutoService()
	imagemService := infra.NewPagamentoService()

	// Routes
	controller.NewProdutoController(g, produtoService, imagemService)
	
    err := e.Start(":4000")
	if err != nil {
		log.Fatalf("Erro ao subir servidor %v", err)
	}
}