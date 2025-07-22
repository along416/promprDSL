module web-backend

go 1.24.4

require (
	promptdslcore v0.0.0
	service v0.0.0
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/dlclark/regexp2 v1.11.4 // indirect
	github.com/dop251/goja v0.0.0-20250630131328-58d95d85e994 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/google/pprof v0.0.0-20230207041349-798e818bf904 // indirect
	github.com/qdxiao/llmjsonrepair v0.0.0-20250628144450-b0ae449984bd // indirect
	github.com/sashabaranov/go-openai v1.40.5 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/text v0.3.8 // indirect
)

replace promptdslcore => ../promptdsl-core

replace service => ../service
