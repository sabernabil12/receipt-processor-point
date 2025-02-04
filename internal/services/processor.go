package services

import (
	"math"
	"receipt-processor-point/internal/models"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	points += calculateRoundDollar(receipt.Total)

	// Rule 3: 25 points if the total is a multiple of 0.25.
	amount, _ := strconv.ParseFloat(receipt.Total, 64)

	// Check if amount is a multiple of 0.25.
	if int(amount*100)%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range receipt.Items {
		descLength := len(strings.TrimSpace(item.ShortDescription))
		if descLength%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}

func calculateRoundDollar(receiptTotal string) int {
	points := 0

	// Check if there is a decimal point.
	if strings.Contains(receiptTotal, ".") {

		// Split the number into a whole and decimal parts.
		parts := strings.Split(receiptTotal, ".")
		if len(parts) != 2 {
			return points // Invalid format.
		}

		wholePart, decimalPart := parts[0], parts[1]

		// Convert wholePart to an integer
		wholeNum, err := strconv.Atoi(wholePart)
		if err != nil || wholeNum <= 0 {
			return points // Skip non-positive values.
		}

		// Ensure the decimal part is exactly "0" or "00".
		if decimalPart == "0" || decimalPart == "00" {
			points += 50
		}
	} else {
		// Convert the whole number into an int and check if it is positive.
		number, err := strconv.Atoi(receiptTotal)
		if err != nil && number > 0 {
			points += 50
		}
	}

	return points
}
