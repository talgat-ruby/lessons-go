package ai

type Handler struct {
	apiKey string
}

func New(apiKey string) *Handler {
	return &Handler{apiKey: apiKey}
}
