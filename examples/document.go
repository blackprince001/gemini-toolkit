package examples

import (
	"context"
	"fmt"

	"github.com/blackprince001/gemini-toolkit/tools"
	"github.com/blackprince001/gemini-toolkit/types"
)

func example3() {
	conf := types.ToolKitConfig{
		ApiKey:    "Hello",
		BaseModel: "Hello",
	}

	document := tools.NewToolKitDocumentService(conf)

	ctx := context.Background()

	documentSummary, err := document.GetDocumentSummary(ctx, "fileUrl")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(documentSummary)
}
