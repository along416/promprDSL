package main
// // 1, Line
// // 2, if/else {
// // 	Line
// // 	if/else
// // 	for 
// // }
// // 3, for 

// type InnerLine struct { 
// 	Kind // 0 = line, 1, if, 2, for
// 	Content // 0
// 	IfCond string // kind = 1
// 	IfSubLines []*InnerLine
// 	ForVar string //
// 	ForSubLines []*InnerLine
// }

// type PromptParsedResult struct { 
// 	input {fieldList}
// 	output {fieldList }
// 	after string
// 	system []...
// 	user []...
// }


// // http://localhost:8080/api/genGuide
// func genGuide(http.ResponseWriter, *http.Request) {
// 	type internal_In struct {
// // input ...
// 	}	

// 	type internal_Out struct {
// // output ...
// 	}

// 	client := openai.NewClient("YOUR_API_KEY")
// 	client.CreateChatCompletion(
// 		context.Background(),
// 		openai.ChatCompletionRequest{
// 			Model: openai.GPT3Dot5Turbo,
// 			Messages: []openai.ChatCompletionMessage{
// 				{
// 					Role:    openai.ChatMessageRoleSystem,
// 					Content: func() string {
// 						systm +
// 						return 
// 					}(),
// 				},
// 				{
// 					Role:    openai.ChatMessageRoleUser,
// 					Content: func() string {
// 						user +
// 						return "What is the meaning of life?"
// 					}(),
// 				}
// 			}
// 		}

// 	)

// }