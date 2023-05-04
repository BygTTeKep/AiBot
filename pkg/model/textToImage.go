package model

type TextToImageImage struct {
	Base64       string `json:"base64"`
	Seed         uint32 `json:"seed"`
	FinishReason string `json:"finishReason"`
}

type TextToImageResponse struct {
	Images []TextToImageImage `json:"artifacts"`
}
