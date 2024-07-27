package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	Result int `json:"result"`
}

func calSum(w http.ResponseWriter, r *http.Request) {
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)
	if errA != nil || errB != nil {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}
	sum := a + b
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
