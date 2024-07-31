package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Result int `json:"result"`
}

type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}

func calSum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var req Request
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	sum := req.A + req.B
	response := Response{Result: sum}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/sum", calSum)
	fmt.Println("Listening on port 1408")
	err := http.ListenAndServe(":1408", nil)
	if err != nil {
		fmt.Println("Error :", err)
	}
}
