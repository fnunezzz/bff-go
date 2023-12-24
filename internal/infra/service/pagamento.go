package infra

import interfaces "github.com/fnunezzz/bff_go/internal/domain/service"



type pagamentoService struct {
}

func NewPagamentoService() interfaces.PagamentoService {
	return &pagamentoService{}
}

func (p *pagamentoService) EfetuarPagamento() {
	// TODO buscar dados externos das aplicacoes
}

