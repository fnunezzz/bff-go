package interfaces

type Midia struct {
	Data []struct {
		Sku string `json:"sku"`
		Imagens []Imagem `json:"imagens"`
	} `json:"data"`
}

type Preco struct {
	Data []struct {
		Sku string `json:"sku"`
		Preco float32 `json:"preco"`
	} `json:"data"`
}

type Imagem struct {
	Url string `json:"url"`
	Principal bool `json:"principal"`
	Orderm int `json:"ordem"`
}

type Produto struct {
	Sku string `json:"sku"`
	Preco float32 `json:"preco"`
	Imagens []Imagem `json:"imagens"`
}

type BuscarDadosProdutoResponse struct {
	Data []Produto `json:"data"`
}

type ProdutoService interface {
	BuscarDadosProduto(skuList []string) (*BuscarDadosProdutoResponse, error)
}
