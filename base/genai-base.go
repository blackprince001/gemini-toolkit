package base

import (
	"context"
	"io"
	"iter"
	"net/http"

	"github.com/blackprince001/gemini-toolkit/types"

	"github.com/pkg/errors"
	"google.golang.org/genai"
)

type ToolKitBaseService struct {
	Config types.ToolKitConfig
}

func NewToolKitBaseService(cfg types.ToolKitConfig) *ToolKitBaseService {
	return &ToolKitBaseService{
		Config: cfg,
	}
}

func (tool *ToolKitBaseService) GetToolKitClient(ctx context.Context) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendGeminiAPI,
		APIKey: tool.Config.ApiKey,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (tool *ToolKitBaseService) GenerateContentStreamText(
	ctx context.Context,
	client *genai.Client,
	prompt string,
) iter.Seq2[*genai.GenerateContentResponse, error] {
	parts := []*genai.Part{
		{Text: prompt},
	}
	contents := []*genai.Content{{Parts: parts}}

	config := types.GetGenerateConfigText()
	result := client.Models.GenerateContentStream(ctx, tool.Config.BaseModel, contents, config)

	return result
}

func (tool *ToolKitBaseService) GenerateContentText(
	ctx context.Context,
	client *genai.Client,
	prompt string,
) (*genai.GenerateContentResponse, error) {
	parts := []*genai.Part{
		{Text: prompt},
	}
	contents := []*genai.Content{{Parts: parts}}

	config := types.GetGenerateConfigText()
	result, err := client.Models.GenerateContent(ctx, tool.Config.BaseModel, contents, config)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (tool *ToolKitBaseService) GenerateContentWithFilesText(
	ctx context.Context,
	client *genai.Client,
	fileUrl, prompt string,
) (*genai.GenerateContentResponse, error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return nil, errors.Errorf("Error fetching Document: %s", err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error failed loading Document content: %s", err)
	}

	parts := []*genai.Part{
		{Text: prompt},
		{InlineData: &genai.Blob{Data: data, MIMEType: "application/pdf"}},
	}
	contents := []*genai.Content{{Parts: parts}}

	config := types.GetGenerateConfigText()
	result, err := client.Models.GenerateContent(ctx, tool.Config.BaseModel, contents, config)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (tool *ToolKitBaseService) GenerateContentWithFilesJSON(
	ctx context.Context,
	client *genai.Client,
	fileUrl, prompt string,
	responseSchema *genai.Schema,
) (*genai.GenerateContentResponse, error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return nil, errors.Errorf("Error fetching Document: %s", err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error failed loading Document content: %s", err)
	}

	parts := []*genai.Part{
		{Text: prompt},
		{InlineData: &genai.Blob{Data: data, MIMEType: "application/json"}},
	}
	contents := []*genai.Content{{Parts: parts}}

	config := types.GetGenerateConfigText()
	result, err := client.Models.GenerateContent(ctx, tool.Config.BaseModel, contents, config)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (tool *ToolKitBaseService) GenerateContentStreamWithFilesJSON(
	ctx context.Context,
	client *genai.Client,
	fileUrl, prompt string,
	responseSchema *genai.Schema,
) (iter.Seq2[*genai.GenerateContentResponse, error], error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return nil, errors.Errorf("Error fetching Document: %s", err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error failed loading Document content: %s", err)
	}

	parts := []*genai.Part{
		{Text: prompt},
		{InlineData: &genai.Blob{Data: data, MIMEType: "application/pdf"}},
	}
	contents := []*genai.Content{{Parts: parts}}

	config := types.GetGenerateConfigJSON(responseSchema)
	result := client.Models.GenerateContentStream(ctx, tool.Config.BaseModel, contents, config)

	return result, nil
}

func (tool *ToolKitBaseService) GenerateContentWithImagesText(
	ctx context.Context,
	client *genai.Client,
	imageUrl, prompt string,
) (*genai.GenerateContentResponse, error) {
	resp, err := http.Get(imageUrl)
	if err != nil {
		return nil, errors.Errorf("Error fetching Image: %s", err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error failed loading Image content: %s", err)
	}

	parts := []*genai.Part{
		{Text: prompt},
		{InlineData: &genai.Blob{Data: data, MIMEType: "image/jpeg"}},
	}
	contents := []*genai.Content{{Parts: parts}}

	config := types.GetGenerateConfigText()
	result, err := client.Models.GenerateContent(ctx, tool.Config.BaseModel, contents, config)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (tool *ToolKitBaseService) GenerateContentWithImagesJSON(
	ctx context.Context,
	client *genai.Client,
	imageUrl, prompt string,
	responseSchema *genai.Schema,
) (*genai.GenerateContentResponse, error) {
	resp, err := http.Get(imageUrl)
	if err != nil {
		return nil, errors.Errorf("Error fetching Image: %s", err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("Error failed loading Image content: %s", err)
	}

	parts := []*genai.Part{
		{Text: prompt},
		{InlineData: &genai.Blob{Data: data, MIMEType: "image/jpeg"}},
	}
	contents := []*genai.Content{{Parts: parts}}

	config := types.GetGenerateConfigJSON(responseSchema)
	result, err := client.Models.GenerateContent(ctx, tool.Config.BaseModel, contents, config)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (tool *ToolKitBaseService) Search(
	ctx context.Context,
	client *genai.Client,
	prompt string,
) (*genai.GenerateContentResponse, error) {
	parts := []*genai.Part{
		{Text: prompt},
	}
	contents := []*genai.Content{{Parts: parts}}

	config := types.GetGenerateConfigSearch()

	result, err := client.Models.GenerateContent(ctx, tool.Config.BaseModel, contents, config)
	if err != nil {
		return nil, err
	}

	return result, nil
}
