package ChatGPT

import (
	gpt3 "github.com/PullRequestInc/go-gpt3"
	"golang.org/x/net/context"
	"log"
	"strings"
)

func GetResponse(client gpt3.Client, ctx context.Context, quesiton string) string {
	var res strings.Builder
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(4000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		res.WriteString(resp.Choices[0].Text)
		//fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		res.WriteString(err.Error())
		//fmt.Println(err)
	}
	//fmt.Printf("\n")
	//fmt.Println(res.String())
	return res.String()
}

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func Callgpt(q string) string {
	log.SetOutput(new(NullWriter))
	apiKey := ""
	if apiKey == "" {
		panic("Missing API KEY")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	//rootCmd := &cobra.Command{
	//	Use:   "chatgpt",
	//	Short: "Chat with ChatGPT in console.",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		GetResponse(client, ctx, q)
	//	},
	//}
	//rootCmd.SetOut(os.Stdout)
	//log.Fatal(rootCmd.Execute())
	return GetResponse(client, ctx, q)
}

func validateQuestion(question string) string {
	quest := strings.Trim(question, " ")
	keywords := []string{"", "loop", "break", "continue", "cls", "exit", "block"}
	for _, x := range keywords {
		if quest == x {
			return ""
		}
	}
	return quest
}
