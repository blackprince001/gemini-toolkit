package examples

import (
	"context"
	"fmt"

	"github.com/blackprince001/gemini-toolkit/tools"
	"github.com/blackprince001/gemini-toolkit/types"
	"github.com/blackprince001/gemini-toolkit/utils"
)

func example1() {
	conf := types.ToolKitConfig{
		ApiKey:    "Hello",
		BaseModel: "Hello",
	}

	chat := tools.NewToolKitChatService(conf)

	ctx := context.Background()

	chatResponse, err := chat.GetChatResponse(ctx, "Which AI model are you and what can you do for me?")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(utils.ResponseToPartString(chatResponse))
}
