package handleRequest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type messageRequest struct {
	Message string `json:"message"`
}

// 定义JSON对应的结构体
func HandlePostJSON(w http.ResponseWriter, r *http.Request) {
	url := r.Host + r.URL.Path
	fmt.Println(url)

	// 检查 Post
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// 检查 Content-Type 是否为 application/json
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
		return
	}
	// 解析请求体中的 JSON 数

	var data messageRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// message := data.Message
	// fmt.Println(message)

	// 使用 json.Marshal 将结构体转换为 JSON 字符串
	jsonData_, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	// 将 JSON 数据转换为字符串
	jsonString := string(jsonData_)
	// 打印 JSON 字符串
	// fmt.Println(jsonString)
	// {"message":"Hi Go"}

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": jsonString})
}

// map[string]interface{}
func HandlePostJSON_Map(w http.ResponseWriter, r *http.Request) {
	url := r.Host + r.URL.Path
	fmt.Println(url)
	// 检查 Post
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// 检查 Content-Type 是否为 application/json
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
		return
	}

	// 解析请求体中的 JSON
	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// message := data["message"]
	// fmt.Println(message)

	// 使用 json.Marshal 将结构体转换为 JSON 字符串
	jsonData_, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	// 将 JSON 数据转换为字符串
	jsonString := string(jsonData_)
	// 打印 JSON 字符串
	// fmt.Println(jsonString)
	// {"message":"Hi Go"}

	// 返回响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"response": jsonString})
}
