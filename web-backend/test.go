package main

import (
	"fmt"
	"log"
	"os"

	"promptdslcore" // 根据实际模块路径改写
)

func main() {
	fileContent, err := os.ReadFile("../promptdsl-core/example.prompt") // 文件扩展名你自己定义
	if err != nil {
		panic(fmt.Errorf("读取 DSL 文件失败: %v", err))
	}
	inputDSL := string(fileContent)

	prompt, err := promptdslcore.RunPromptDSL(inputDSL)
	if err != nil {
		log.Fatalf("RunPromptDSL error: %v", err)
	}
	fmt.Println("生成的 Prompt:\n", prompt)
}
