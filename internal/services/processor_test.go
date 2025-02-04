package services_test

import (
	"receipt-processor-point/internal/models"
	"receipt-processor-point/internal/services"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "Case 1: Example Target",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			name: "Case 2: Round Dollar amount",
			receipt: models.Receipt{Retailer: "Walmart",
				PurchaseDate: "2023-10-16",
				PurchaseTime: "12:30",
				Items: []models.Item{
					{ShortDescription: "Coca Cola 12PK", Price: "5.00"},
				},
				Total: "5.00"},
			expected: 82,
			//  7 points for alphanumeric + 50 points for round dollar amount +
			// 25 points for multiple of 0.25.
		},
		{
			name: "Case 3: Odd purchase date",
			receipt: models.Receipt{
				Retailer:     "Best Buy",
				PurchaseDate: "2023-10-15", // 15 is odd
				PurchaseTime: "14:25",
				Items: []models.Item{
					{ShortDescription: "USB Cable", Price: "10.25"},
				},
				Total: "10.25",
			},
			expected: 51,
			// 7 points for alphanumeric + 25 points for multiple of 0.25 +
			// 3 points for length multipe of 3 + 6 points for odd date +
			// 10 points for the time.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points := services.CalculatePoints(tt.receipt)
			if points != tt.expected {
				t.Errorf("Test case '%s' failed: expected %d points, got %d", tt.name, tt.expected, points)
			}
		})
	}
}
