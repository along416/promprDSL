package main

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"service"
	"strings"

	// "net/http"
	"os"

	"promptdslcore" // 根据实际模块路径改写
)

// var llm *service.LLMClient
func main() {
	// 打开或创建日志文件
	logFile, err := os.OpenFile("llm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	// 设置日志输出到文件
	log.SetOutput(logFile)
	// 可选: 添加时间戳、文件名等信息
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	fileContent, err := os.ReadFile("../promptdsl-core/example copy.prompt") // 文件扩展名你自己定义
	if err != nil {
		panic(fmt.Errorf("读取 DSL 文件失败: %v", err))
	}
	inputDSL := string(fileContent)

	prompt, err := promptdslcore.RunPromptDSL(inputDSL)
	if err != nil {
		log.Fatalf("RunPromptDSL error: %v", err)
	}
	fmt.Println("生成的 Prompt:\n", prompt)
	// 步骤 2：提取 system / user 块
	promptcontent := *prompt
	systemPart := promptcontent.Sys
	userPart := promptcontent.User
	systemText := strings.Join(systemPart, "\n")
	userText := strings.Join(userPart, "\n")
	// fmt.Println("systemText:", systemText)
	apiKey := "sk-02e496929ecc485796d29bd94e7ce371"
	llm := service.NewLLMClient(apiKey)
	result, err := llm.GeneratePromptResponse(systemText, userText)

	// if err != nil {
	// 	http.Error(w, "调用大模型失败: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	log.Println("模型输出:\n", result)
	// 假设模型输出是 JSON 数组字符串（或你需要先解析处理）

	final, err := CallUserPostProcessor(result)
	if err != nil {
		log.Fatalf("调用 AfterProcess 失败: %v", err)
	}
	log.Println("后处理结果:", final)

	//goja
	// //处理输出，调用生成的函数执行
	// // FixCode 是 DSL 中生成的 JS 代码，比如含有 function FixProcess(...)
	// script := `
	// function FixProcess(response) {
	//     try {
	//         var output = JSON.parse(response);
	//         for (var i = 0; i < output.length; i++) {
	//             var item = output[i];
	//             if (!item.KnowledgePoint || item.KnowledgePoint.trim() === "") {
	//                 item.KnowledgePoint = "默认知识点";
	//             }
	//             if (Array.isArray(item.Conditions)) {
	//                 item.Conditions = item.Conditions.map(function(cond) {
	//                     return cond.trim();
	//                 });
	//             }
	//             var delta = Math.abs(item.KnowledgePoint.length - 5);
	//             if (delta > 10) {
	//                 console.warn("知识点长度过长:", item.KnowledgePoint);
	//             }
	//         }
	//         return JSON.stringify(output);
	//     } catch (e) {
	//         console.error("解析或处理失败:", e);
	//         return response;
	//     }
	// };
	// `

	// script=prompt.Fix
	// vm := goja.New()
	// // 自定义一个简单的console对象，映射到Go的log.Println
	// console := map[string]func(goja.FunctionCall) goja.Value{
	// 	"log": func(fc goja.FunctionCall) goja.Value {
	// 		args := []interface{}{}
	// 		for _, arg := range fc.Arguments {
	// 			args = append(args, arg.String())
	// 		}
	// 		log.Println(args...)
	// 		return goja.Undefined()
	// 	},
	// 	"warn": func(fc goja.FunctionCall) goja.Value {
	// 		args := []interface{}{}
	// 		for _, arg := range fc.Arguments {
	// 			args = append(args, arg.String())
	// 		}
	// 		// log.Println("WARN:", args...)
	// 		return goja.Undefined()
	// 	},
	// 	"error": func(fc goja.FunctionCall) goja.Value {
	// 		args := []interface{}{}
	// 		for _, arg := range fc.Arguments {
	// 			args = append(args, arg.String())
	// 		}
	// 		// log.Println("ERROR:", args...)
	// 		return goja.Undefined()
	// 	},
	// }

	// vm.Set("console", console)
	// // 加载 script 脚本，注册 FixProcess 到 VM
	// _, err = vm.RunString(script)
	// if err != nil {
	// 	log.Fatalf("JS 运行错误: %v", err)
	// }

	// // 从 JS VM 中取出函数 FixProcess
	// fn, ok := goja.AssertFunction(vm.Get("FixProcess"))
	// if !ok {
	// 	log.Fatal("FixProcess 不是一个函数")
	// }

	// // 执行 JS 函数 FixProcess
	// finalresult, err := fn(goja.Undefined(), vm.ToValue(result))
	// if err != nil {
	// 	log.Fatalf("FixProcess 执行错误: %v", err)
	// }

	// fmt.Println("结果:", finalresult.String()) // Hello World（去空格）
}

// 假设 result 是 string，包含 JSON 数组（即模型返回结果）
func CallUserPostProcessor(jsonStr string) ([]map[string]interface{}, error) {
	fmt.Println("CallUserPostProcessor进入前")
	cmd := exec.Command("go", "run", "code_gen/code_gen.go")
	cmd.Stderr = os.Stderr
	// cmd.Stdout = os.Stdout // 可选，也能打印标准输出（调试用）
	stdin, _ := cmd.StdinPipe()
	// if err := cmd.Start(); err != nil {
	// 	fmt.Println("cmd.Start error:", err)
	// 	return nil, err
	// }

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("2Error:", err)
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("3Error:", err)
		return nil, err
	}

	// 传入 JSON 输入
	_, err = stdin.Write([]byte(jsonStr))
	if err != nil {
		return nil, err
	}
	stdin.Close()

	// 读取输出
	outputBytes, err := io.ReadAll(stdout)
	if err != nil {
		return nil, err
	}
	fmt.Println("outputBytes:",string(outputBytes))
	cmd.Wait()

	var finalOutput []map[string]interface{}
	err = json.Unmarshal(outputBytes, &finalOutput)
	if err != nil {
		return nil, fmt.Errorf("无法解析 AfterProcess 输出: %v", err)
	}
	return finalOutput, nil
}
