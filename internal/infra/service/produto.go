package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	interfaces "github.com/fnunezzz/bff_go/internal/domain/service"
)


type produtoService struct {
}

func NewProdutoService() interfaces.ProdutoService {
	return &produtoService{}
}

// Requisitando o serviço de preco
func (p *produtoService) buscarPreco(skuList []string, channel chan any, waitGroup *sync.WaitGroup) interface{} {
	defer waitGroup.Done()
	requestURL := fmt.Sprintf("%s/core/preco", os.Getenv("PRECO_URL"))
	// Chamada ao preco
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		channel <- err
		return err
	}
	res, err := http.DefaultClient.Do(req)
	
	if err != nil {
		channel <- err
		return err
	}
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		channel <- err
		return err
	}
	

	precos := &interfaces.Preco{}
	err = json.Unmarshal(body, &precos)
	if err != nil {
		channel <- err
		return err
	}

	if len(precos.Data) == 0 {
		err = errors.New("Nenhum produto encontrado")
		channel <- err
		return err
	}
	channel <- precos
	return precos
}

// Requisitando o serviço de midia
func (p *produtoService) buscarMidia(skuList []string, channel chan any, waitGroup *sync.WaitGroup) interface{} {
	defer waitGroup.Done()
	// chamada ao midia
	requestURL := fmt.Sprintf("%s/core/imagem", os.Getenv("MIDIA_URL"))
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		channel <- err
		return err
	}
	
	res, err := http.DefaultClient.Do(req)
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		channel <- err
		return err
	}
	midias := &interfaces.Midia{}
	err = json.Unmarshal(body, &midias)
	if err != nil {
		channel <- err
		return err
	}
	channel <- midias
	return midias
}

func (p *produtoService) BuscarDadosProduto(skuList []string) (*interfaces.BuscarDadosProdutoResponse, error) {
	response := &interfaces.BuscarDadosProdutoResponse{}
	precos := &interfaces.Preco{}
	midias := &interfaces.Midia{}
	channel := make(chan any, 2) // Criando um canal de dois processos em paralelo
	waitGroup := &sync.WaitGroup{} // Criando um grupo de espera
	waitGroup.Add(2) // Adicionando dois processos ao grupo de espera

	// paralelizando as chamadas aos serviços externos
	go p.buscarPreco(skuList, channel, waitGroup)
	go p.buscarMidia(skuList, channel, waitGroup)
	waitGroup.Wait() // block ate duas chamadas termianrem
	close(channel)

	// Buscando a resposta das chamadas (preco e midia)
	// Equivalente a promise.all do Node.js
	// Inferindo o tipo de resposta para atribuir a variavel correta
	for resp := range channel {
		if p, ok := resp.(*interfaces.Preco); ok {
			precos = p
		} else if m, ok := resp.(*interfaces.Midia); ok {
			midias = m
		}
	}

	
	var produtos []interfaces.Produto
	for _, preco := range precos.Data {
		produto := &interfaces.Produto{}
		produto.Sku = preco.Sku
		produto.Preco = preco.Preco
		produtos = append(produtos, *produto)
	}
	
	response.Data = produtos
	
	
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

