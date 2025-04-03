package types

import "google.golang.org/genai"

func GetGenerateConfigJSON(responseSchema *genai.Schema) *genai.GenerateContentConfig {
	return &genai.GenerateContentConfig{
		Temperature:      genai.Ptr[float32](0.2),
		ResponseMIMEType: "application/json",
		StopSequences:    []string{"\n\n"},
		Seed:             genai.Ptr[int32](420),
		ResponseSchema:   responseSchema,
	}
}

func GetGenerateConfigText() *genai.GenerateContentConfig {
	return &genai.GenerateContentConfig{
		Temperature:       genai.Ptr[float32](0.5),
		TopP:              genai.Ptr[float32](0.5),
		TopK:              genai.Ptr[float32](2.0),
		ResponseMIMEType:  "text/plain",
		SystemInstruction: &genai.Content{Parts: []*genai.Part{{Text: "You are a helpful assistant."}}},
	}
}

func GetGenerateConfigSearch() *genai.GenerateContentConfig {
	return &genai.GenerateContentConfig{
		Temperature:      genai.Ptr[float32](0.2),
		TopP:             genai.Ptr[float32](0.9),
		TopK:             genai.Ptr[float32](2.0),
		ResponseMIMEType: "text/plain",
		Tools:            []*genai.Tool{{GoogleSearch: &genai.GoogleSearch{}}},
	}
}