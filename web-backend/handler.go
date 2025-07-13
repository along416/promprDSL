package main

import (
	"encoding/json"
	"io"
	"net/http"
	"promptdslcore" // 引用核心模块
)

// 请求结构
type GenGuideRequest struct {
	Input string `json:"input"`
}

// 响应结构
type GenGuideResponse struct {
	Prompt string `json:"prompt"`
}

// POST /api/genGuide
func HandleGenGuide(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "读取请求体失败", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req GenGuideRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "JSON 解析失败", http.StatusBadRequest)
		return
	}

	// 调用核心模块处理 DSL
	prompt, err := promptdslcore.RunPromptDSL(req.Input)
	if err != nil {
		http.Error(w, "生成失败: "+err.Error(), http.StatusInternalServerError)
		return
	}

	resp := GenGuideResponse{Prompt: prompt}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
