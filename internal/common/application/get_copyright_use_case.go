package application

import "github.com/paguerre3/blockchain/configs"

var (
	config = configs.Instance()
)

type PageData struct {
	Year int `json:"year"`
}

type GetCopyrightUseCase interface {
	GetCopyright() *PageData
}

type getCopyrightUseCase struct {
}

func NewGetCopyrightUseCase() GetCopyrightUseCase {
	return &getCopyrightUseCase{}
}

func (g *getCopyrightUseCase) GetCopyright() *PageData {
	return &PageData{Year: config.Wallet().CopyrightYear()}
}
