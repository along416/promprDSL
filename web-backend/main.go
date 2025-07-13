package main

import (
	"log"
	"net/http"
)

func main() {
	// 注册路由
	mux := http.NewServeMux()
	RegisterRoutes(mux)

	log.Println("🚀 HTTP Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("启动失败:", err)
	}
}
