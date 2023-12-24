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

func (p *produtoService) buscarPreco(skuList []string) (*interfaces.Preco, error) {
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
	return precos, nil

}

func (p *produtoService) buscarMidia(skuList []string) (*interfaces.Midia, error){
	// chamada ao midia
	requestURL := fmt.Sprintf("%s/core/imagem", os.Getenv("MIDIA_URL"))
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err 
	}
	midias := &interfaces.Midia{}
	json.Unmarshal(body, &midias)
	return midias, nil
}

func (p *produtoService) BuscarDadosProduto(skuList []string) (*interfaces.BuscarDadosProdutoResponse, error) {
	
	precos, err := p.buscarPreco(skuList)

	if err != nil {
		return nil, err
	}

	var produtos []interfaces.Produto
	for _, preco := range precos.Data {
		produto := &interfaces.Produto{}
		produto.Sku = preco.Sku
		produto.Preco = preco.Preco
		produtos = append(produtos, *produto)
	}
	
	response := &interfaces.BuscarDadosProdutoResponse{}
	response.Data = produtos

	midias, err := p.buscarMidia(skuList)
	if err != nil {
		return nil, err
	}
	
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

