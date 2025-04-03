package examples

import (
	"context"
	"fmt"
	"iter"

	"github.com/blackprince001/gemini-toolkit/tools"
	"github.com/blackprince001/gemini-toolkit/types"
	"github.com/blackprince001/gemini-toolkit/utils"
)

func example2() {
	conf := types.ToolKitConfig{
		ApiKey:    "Hello",
		BaseModel: "Hello",
	}

	chat := tools.NewToolKitChatService(conf)

	ctx := context.Background()

	chatResponse, err := chat.GetChatResponseStream(ctx, "Which AI model are you and what can you do for me?")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(utils.IterResponseToString(chatResponse)) // prints the streamed response all at once

	// print each token as you receive them
	next, stop := iter.Pull2(chatResponse)
	defer stop()

	for {
		resp, err, ok := next()
		if !ok {
			break
		}

		if err != nil {
			break
		}

		fmt.Println(utils.ResponseToPartString(resp))
	}
}
