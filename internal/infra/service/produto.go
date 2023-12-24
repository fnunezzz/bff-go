package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	interfaces "github.com/fnunezzz/bff_go/internal/domain/service"
)


type produtoService struct {
}

func NewProdutoService() interfaces.ProdutoService {
	return &produtoService{}
}

func (p *produtoService) BuscarDadosProduto(skuList []string) (*interfaces.BuscarDadosProdutoResponse, error) {
	requestURL := fmt.Sprintf("%s/core/preco", os.Getenv("PRECO_URL"))
	// Chamada ao preco
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	
	if err != nil {
		return nil, err
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	

	precos := &interfaces.Preco{}
	json.Unmarshal(body, &precos)
	
	if len(precos.Data) == 0 {
		return nil, errors.New("Nenhum produto encontrado")
	}

	response := &interfaces.BuscarDadosProdutoResponse{}

	var produtos []interfaces.Produto
	for _, preco := range precos.Data {
		produto := &interfaces.Produto{}
		produto.Sku = preco.Sku
		produto.Preco = preco.Preco
		produtos = append(produtos, *produto)
	}

	response.Data = produtos

	// chamada ao midia
	requestURL = fmt.Sprintf("%s/core/imagem", os.Getenv("MIDIA_URL"))
	req, err = http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return response, err // todo posso ter retornos parciais
	}
	
	res, err = http.DefaultClient.Do(req)

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return response, err // todo posso ter retornos parciais
	}
	midias := &interfaces.Midia{}
	json.Unmarshal(body, &midias)

	for _, midia := range midias.Data {
		exists := false
		for i, produto := range response.Data {
			if produto.Sku == midia.Sku {
				response.Data[i].Imagens = midia.Imagens
				exists = true
			}
		}
		if !exists {
			produto := &interfaces.Produto{}
			produto.Sku = midia.Sku
			produto.Imagens = midia.Imagens
			produto.Preco = 0
			response.Data = append(response.Data, *produto)
		}
	}
	
	return response, nil
	
}

