package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	interfaces "github.com/fnunezzz/bff_go/internal/domain/service"
	"github.com/labstack/echo/v4"
)

type produtoController struct {
	produtoService interfaces.ProdutoService
	imagemService interfaces.PagamentoService
}

type healthDto struct {
	Time time.Time `json:"time"`
	Status bool `json:"status"`
}



func NewProdutoController(g *echo.Group, produtoService interfaces.ProdutoService, imagemService interfaces.PagamentoService) {
	controller := &produtoController{produtoService: produtoService, imagemService: imagemService}
	g.GET("/status", controller.Health)
	g.GET("/produto/informacoes", controller.BuscarInformacoes)
}

func (pc *produtoController) BuscarInformacoes(c echo.Context) error {
	sku := c.QueryParam("sku")
	if sku == "" {
		return c.JSON(http.StatusBadRequest, "sku não informado")
	}
	arr := strings.Split(sku, ",")
	
	res, _ := pc.produtoService.BuscarDadosProduto(arr)
	fmt.Println(res)
	// dto := &healthDto{Time: time.Now(), Status: true}
	return c.JSON(http.StatusOK, res)
	// return json.NewEncoder(c.Response()).Encode(json.RawMessage(res))
}

func (pc *produtoController) Health(c echo.Context) error {
		dto := &healthDto{Time: time.Now(), Status: true}
		return c.JSON(http.StatusOK, dto)

}