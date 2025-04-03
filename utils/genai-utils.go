package utils

import (
	"encoding/json"
	"iter"

	"google.golang.org/genai"
)

func ResponseToPartString(response *genai.GenerateContentResponse) string {
	var raw string = ""
	for _, cand := range response.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				raw += part.Text
			}
		}
	}
	return raw
}

func IterResponseToString(response iter.Seq2[*genai.GenerateContentResponse, error]) string {
	next, stop := iter.Pull2(response)
	defer stop()

	var data string = ""
	for {
		resp, err, ok := next()
		if !ok {
			break
		}

		if err != nil {
			break
		}

		data += ResponseToPartString(resp)
	}

	return data
}

func ResponseToStructure(response *genai.GenerateContentResponse, object any) error {
	content := ResponseToPartString(response)

	err := json.Unmarshal([]byte(content), &object)
	if err != nil {
		return err
	}

	return nil
}
