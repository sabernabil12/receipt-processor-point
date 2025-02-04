package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor-point/internal/models"
	"receipt-processor-point/internal/services"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var receipts = make(map[string]int)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid receipt", http.StatusBadRequest)
		return
	}

	receipt.ID = uuid.New().String()
	points := services.CalculatePoints(receipt)
	receipts[receipt.ID] = points

	response := map[string]string{"id": receipt.ID}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := receipts[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response := models.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
