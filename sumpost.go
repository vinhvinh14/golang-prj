package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Định nghĩa cấu trúc dữ liệu cho phản hồi 
type Response struct {
	Result int `json:"result"`
}

// Định nghĩa cấu trúc dữ liệu cho yêu cầu 
type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

// Hàm xử lý yêu cầu tính tổng
func calSum(w http.ResponseWriter, r *http.Request) {
	// Kiểm tra phương thức yêu cầu có phải POST không
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Đọc nội dung từ yêu cầu
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	// Phân tích JSON từ yêu cầu thành cấu trúc dữ liệu
	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Tính tổng của hai số a và b
	sum := req.A + req.B
	// Tạo phản hồi với kết quả
	response := Response{Result: sum}
	// Đặt tiêu đề cho phản hồi là JSON
	w.Header().Set("Content-Type", "application/json")
	// Mã hóa phản hồi thành JSON và gửi 
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Đăng ký hàm xử lý cho đường dẫn 
	http.HandleFunc("/sum", calSum)
	fmt.Println("Listening on port 1408")
	// Lắng nghe trên cổng 1408
	err := http.ListenAndServe(":1408", nil)
	// Xử lí lỗi sever
	if err != nil {
		fmt.Println("Error :", err)
	}
}
