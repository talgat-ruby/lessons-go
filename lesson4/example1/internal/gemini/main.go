package gemini

type Gemini struct {
	apiKey string
}

func NewGemini(apiKey string) *Gemini {
	return &Gemini{apiKey}
}
