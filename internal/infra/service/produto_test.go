package infra

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	interfaces "github.com/fnunezzz/bff_go/internal/domain/service"
)

// Teste para verificar se o serviço está sendo instanciado corretamente
//
// Deve retornar uma instancia da interface ProdutoService
func TestSeInstanciaProdutoService(t *testing.T) {
	produtoService := NewProdutoService()
	if produtoService == nil {
		t.Fatalf("Esperava que ProdutoService não fosse nulo")
	}
	_, err := produtoService.(interfaces.ProdutoService)
	if !err {
		t.Fatalf("Esperava que ProdutoService implementasse interfaces.ProdutoService")
	}

}


// Teste para verificar se o serviço está retornando os dados corretos quando o produto não possui imagem.
//
// Deve retornar 1 produto somente com preco
func TestBuscarDadosProdutoUmProdutoSemImagem(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/core/preco" && r.URL.Path != "/core/imagem" {
			t.Fatalf("Expected to request '/core/preco' or '/core/imagem', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": [{"sku": "1515", "preco": 10.00}]}`))
	}))
	defer server.Close()

	// redireciona as chamadas para o servidor mockado
	os.Setenv("PRECO_URL", server.URL)
	os.Setenv("MIDIA_URL", server.URL)
	var a = []string{"1515", "1520"}
	produtoService := NewProdutoService()
	
	value, err := produtoService.BuscarDadosProduto(a)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if value == nil {
		t.Fatalf("Esperava que BuscarDadosProdutoResponse não fosse nulo")
	}

	if value.Data[0].Sku != "1515" {
		t.Fatalf("Esperava que sku = 1515")
	}
	if value.Data[0].Preco != 10.00 {
		t.Fatalf("Esperava que preco = 10.00")
	}
}


//Teste para verificar se o serviço está retornando os dados corretos quando o produto possui imagem
//
//Deve retornar 1 produto completo
func TestBuscarDadosProdutoUmProdutoComImagem(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/core/preco" && r.URL.Path != "/core/imagem" {
			t.Fatalf("Expected to request '/core/preco' or '/core/imagem', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		if r.URL.Path == "/core/preco" {
			w.Write([]byte(`{"data": [{"sku": "1515", "preco": 10.00}]}`))
		}
		if r.URL.Path == "/core/imagem" {
			w.Write([]byte(`{"data": [{"sku": "1515", "imagens": [{"url": "http://teste.com.br", "principal": true, "ordem": 1}]}]}`))
		}
	}))
	defer server.Close()

	// redireciona as chamadas para o servidor mockado
	os.Setenv("PRECO_URL", server.URL)
	os.Setenv("MIDIA_URL", server.URL)
	var a = []string{"1515", "1520"}
	produtoService := NewProdutoService()
	
	value, err := produtoService.BuscarDadosProduto(a)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if value == nil {
		t.Fatalf("Esperava que BuscarDadosProdutoResponse não fosse nulo")
	}

	if value.Data[0].Sku != "1515" {
		t.Fatalf("Esperava que sku = 1515")
	}
	if value.Data[0].Preco != 10.00 {
		t.Fatalf("Esperava que preco = 10.00")
	}

	if len(value.Data[0].Imagens) != 1 {
		t.Fatalf("Esperava que imagens fosse de tamanho 1")
	}

	if value.Data[0].Imagens[0].Url != "http://teste.com.br" {
		t.Fatalf("Esperava que url = http://teste.com.br")
	}
	if value.Data[0].Imagens[0].Principal != true {
		t.Fatalf("Esperava que principal = true")
	}
	if value.Data[0].Imagens[0].Orderm != 1 {
		t.Fatalf("Esperava que ordem = 1")
	}
	
}


//Teste para verificar se o serviço está retornando os dados corretos quando a imagem retornada não bate com o produto
//
//Deve retornar 2 produtos: 1 com imagem e sem preco e outro com preco sem imagem
func TestBuscarDadosProdutoUmProdutoComImagemErrada(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/core/preco" && r.URL.Path != "/core/imagem" {
			t.Fatalf("Expected to request '/core/preco' or '/core/imagem', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		if r.URL.Path == "/core/preco" {
			w.Write([]byte(`{"data": [{"sku": "1515", "preco": 10.00}]}`))
		}
		if r.URL.Path == "/core/imagem" {
			w.Write([]byte(`{"data": [{"sku": "1520", "imagens": [{"url": "http://teste.com.br", "principal": true, "ordem": 1}]}]}`))
		}
	}))
	defer server.Close()

	// redireciona as chamadas para o servidor mockado
	os.Setenv("PRECO_URL", server.URL)
	os.Setenv("MIDIA_URL", server.URL)
	var a = []string{"1515", "1520"}
	produtoService := NewProdutoService()
	
	value, err := produtoService.BuscarDadosProduto(a)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if value == nil {
		t.Fatalf("Esperava que BuscarDadosProdutoResponse não fosse nulo")
	}

	if len(value.Data) != 2 {
		t.Fatalf("Esperava que BuscarDadosProdutoResponse.Data fosse de tamanho 2")
	}

	if value.Data[0].Sku != "1515" {
		t.Fatalf("Esperava que sku = 1515")
	}
	if value.Data[0].Preco != 10.00 {
		t.Fatalf("Esperava que preco = 10.00")
	}

	if len(value.Data[0].Imagens) != 0 {
		t.Fatalf("Esperava que imagens fosse vazio")
	}

	if value.Data[1].Sku != "1520" {
		t.Fatalf("Esperava que sku = 1520")
	}

	if value.Data[1].Preco != 0 {
		t.Fatalf("Esperava que preco = 0")
	}

	if len(value.Data[1].Imagens) != 1 {
		t.Fatalf("Esperava que imagens fosse de tamanho 1")
	}

	if value.Data[1].Imagens[0].Url != "http://teste.com.br" {
		t.Fatalf("Esperava que url = http://teste.com.br")
	}

	if value.Data[1].Imagens[0].Principal != true {
		t.Fatalf("Esperava que principal = true")
	}

	if value.Data[1].Imagens[0].Orderm != 1 {
		t.Fatalf("Esperava que ordem = 1")
	}
}


//Teste para verificar se o serviço está retornando os dados corretos quando multiplos retornos diferentes
//
//Deve retornar 3 produtos: 2 com imagem e preco e 1 com preco e sem imagem
func TestBuscarDadosProduto(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/core/preco" && r.URL.Path != "/core/imagem" {
			t.Fatalf("Expected to request '/core/preco' or '/core/imagem', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		if r.URL.Path == "/core/preco" {
			w.Write([]byte(`{"data": [{"sku": "1515", "preco": 10.00}, {"sku": "1520", "preco": 20.00}, {"sku": "1530", "preco": 30.00}]}`))
		}
		if r.URL.Path == "/core/imagem" {
			w.Write([]byte(`{
				"data": [
				  {
					"sku": "1515",
					"imagens": [
					  {
						"url": "http://teste.com.br/1",
						"principal": true,
						"ordem": 1
					  },
					  {
						"url": "http://teste.com.br/2",
						"principal": false,
						"ordem": 2
					  }
					]
				  },
				  {
					"sku": "1530",
					"imagens": [
					  {
						"url": "http://teste.com.br/1",
						"principal": true,
						"ordem": 1
					  },
					  {
						"url": "http://teste.com.br/2",
						"principal": false,
						"ordem": 2
					  },
					  {
						"url": "http://teste.com.br/3",
						"principal": false,
						"ordem": 3
					  }
					]
				  }
				]
			  }`))
		}
	}))
	defer server.Close()

	// redireciona as chamadas para o servidor mockado
	os.Setenv("PRECO_URL", server.URL)
	os.Setenv("MIDIA_URL", server.URL)
	var a = []string{"1515", "1520"}
	produtoService := NewProdutoService()
	
	value, err := produtoService.BuscarDadosProduto(a)
	
	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if value == nil {
		t.Fatalf("Esperava que BuscarDadosProdutoResponse não fosse nulo")
	}

	// SKU 1515
	if value.Data[0].Sku != "1515" {
		t.Fatalf("[1515] Esperava que sku = 1515")
	}
	if value.Data[0].Preco != 10.00 {
		t.Fatalf("[1515] Esperava que preco = 10.00")
	}

	if len(value.Data[0].Imagens) != 2 {
		t.Fatalf("[1515] Esperava que imagens fosse de tamanho 2")
	}

	if value.Data[0].Imagens[0].Url != "http://teste.com.br/1" {
		t.Fatalf("[1515] Esperava que primeira imagem url = http://teste.com.br/1")
	}
	if value.Data[0].Imagens[0].Principal != true {
		t.Fatalf("[1515] Esperava que primeira imagem principal = true")
	}
	if value.Data[0].Imagens[0].Orderm != 1 {
		t.Fatalf("[1515] Esperava que primeira imagem ordem = 1")
	}
	if value.Data[0].Imagens[1].Url != "http://teste.com.br/2" {
		t.Fatalf("[1515] Esperava que segunda imagem url = http://teste.com.br/2")
	}
	if value.Data[0].Imagens[1].Principal != false {
		t.Fatalf("[1515] Esperava que segunda imagem principal = true")
	}
	if value.Data[0].Imagens[1].Orderm != 2 {
		t.Fatalf("[1515] Esperava que segunda imagem ordem = 1")
	}

	// SKU 1520
	if value.Data[1].Sku != "1520" {
		t.Fatalf("[1520] Esperava que sku = 1520")
	}
	if value.Data[1].Preco != 20.00 {
		t.Fatalf("[1520] Esperava que preco = 20.00")
	}

	if len(value.Data[1].Imagens) != 0 {
		t.Fatalf("[1520] Esperava que imagens fosse vazio")
	}


	// SKU 1530
	if value.Data[2].Sku != "1530" {
		t.Fatalf("[1530] Esperava que sku = 1530")
	}
	if value.Data[2].Preco != 30.00 {
		t.Fatalf("[1530] Esperava que preco = 30.00")
	}

	if len(value.Data[2].Imagens) != 3 {
		t.Fatalf("[1530] Esperava que houvessem 3 imagens")
	}

	if value.Data[2].Imagens[0].Url != "http://teste.com.br/1" {
		t.Fatalf("[1530] Esperava que a primeira imagem url = http://teste.com.br/1")
	}
	if value.Data[2].Imagens[0].Principal != true {
		t.Fatalf("[1530] Esperava que a primeira imagem principal = true")
	}
	if value.Data[2].Imagens[0].Orderm != 1 {
		t.Fatalf("[1530] Esperava que a primeira imagem ordem = 1")
	}
	if value.Data[2].Imagens[1].Url != "http://teste.com.br/2" {
		t.Fatalf("[1530] Esperava que a segunda imagem url = http://teste.com.br/2")
	}
	if value.Data[2].Imagens[1].Principal != false {
		t.Fatalf("[1530] Esperava que a segunda imagem principal = false")
	}
	if value.Data[2].Imagens[1].Orderm != 2 {
		t.Fatalf("[1530] Esperava que a segunda imagem ordem = 2")
	}
	if value.Data[2].Imagens[2].Url != "http://teste.com.br/3" {
		t.Fatalf("[1530] Esperava que a terceira imagem url = http://teste.com.br/3")
	}
	if value.Data[2].Imagens[2].Principal != false {
		t.Fatalf("[1530] Esperava que a terceira imagem principal = false")
	}
	if value.Data[2].Imagens[2].Orderm != 3 {
		t.Fatalf("[1530] Esperava que a terceira imagem ordem = 3")
	}
}