package tools

import (
	"context"
	"iter"

	"github.com/blackprince001/gemini-toolkit/base"
	"github.com/blackprince001/gemini-toolkit/types"

	"google.golang.org/genai"
)

type ToolkitChatService struct {
	config types.ToolKitConfig
	base   base.ToolKitBaseService
}

func NewToolKitChatService(cfg types.ToolKitConfig) *ToolkitChatService {
	return &ToolkitChatService{
		config: cfg,
	}
}

func (tool *ToolkitChatService) GetChatResponseStream(ctx context.Context, prompt string) (iter.Seq2[*genai.GenerateContentResponse, error], error) {
	client, err := tool.base.GetToolKitClient(ctx)
	if err != nil {
		return nil, err
	}

	stream := tool.base.GenerateContentStreamText(ctx, client, prompt)

	return stream, nil
}

func (tool *ToolkitChatService) GetChatResponse(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	client, err := tool.base.GetToolKitClient(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := tool.base.GenerateContentText(ctx, client, prompt)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (tool *ToolkitChatService) GetSearchResponse(ctx context.Context, prompt string) (*genai.GenerateContentResponse, error) {
	client, err := tool.base.GetToolKitClient(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := tool.base.Search(ctx, client, prompt)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
