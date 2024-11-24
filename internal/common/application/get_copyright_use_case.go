package application

const (
	WALLET_COPYRIGHT_YEAR = 2022
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
	return &PageData{Year: WALLET_COPYRIGHT_YEAR}
}
