package tools

import (
	"context"

	"github.com/blackprince001/gemini-toolkit/base"
	"github.com/blackprince001/gemini-toolkit/types"
	"github.com/blackprince001/gemini-toolkit/utils"
)

type ToolKitDocumentService struct {
	config types.ToolKitConfig
	base   base.ToolKitBaseService
}

func NewToolKitDocumentService(cfg types.ToolKitConfig) *ToolKitDocumentService {
	return &ToolKitDocumentService{
		config: cfg,
	}
}

func (tool *ToolKitDocumentService) GetDocumentSummary(ctx context.Context, fileUrl string) (string, error) {
	client, err := tool.base.GetToolKitClient(ctx)
	if err != nil {
		return "", err
	}

	response, err := tool.base.GenerateContentWithFilesText(
		ctx,
		client,
		fileUrl,
		"Can you generate a summary of this document in less than 100 words?",
	)
	if err != nil {
		return "", err
	}

	txt := utils.ResponseToPartString(response)

	return txt, nil
}

func (tool *ToolKitDocumentService) GetDocumentType(ctx context.Context, fileUrl string) (string, error) {
	client, err := tool.base.GetToolKitClient(ctx)
	if err != nil {
		return "", err
	}

	response, err := tool.base.GenerateContentWithFilesText(
		ctx,
		client,
		fileUrl,
		"Can you give me a category word to describe this document?",
	)
	if err != nil {
		return "", err
	}

	txt := utils.ResponseToPartString(response)

	return txt, nil
}
