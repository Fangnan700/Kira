package model

type ReceiveData struct {
	ID      string        `json:"id"`
	Object  string        `json:"object"`
	Created int           `json:"created"`
	Model   string        `json:"model"`
	Choices []DataChoices `json:"choices"`
}

type DataChoices struct {
	Delta        ChoicesDelta `json:"delta"`
	Index        int          `json:"index"`
	FinishReason string       `json:"finish_reason"`
}

type ChoicesDelta struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
