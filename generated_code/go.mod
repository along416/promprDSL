module workflow

go 1.24.4

require service v0.0.0

require (
	github.com/sashabaranov/go-openai v1.40.5 // indirect
	github.com/wangii/littlepool v1.0.4 // indirect
)

replace service => ../service
