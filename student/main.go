// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"strconv"
// )

// func calculateRange(numbers []float64) (float64, float64) {
// 	if len(numbers) < 2 {
// 		return float64(int(numbers[0]) - 1), float64(int(numbers[0]) + 1)
// 	} else {
// 		x := make([]float64, len(numbers))
// 		for i := 0; i < len(numbers); i++ {
// 			x[i] = float64(i)
// 		}

// 		var sumX, sumY, sumXY, sumX2 float64

// 		for i := 0; i < len(numbers); i++ {
// 			sumX += x[i]
// 			sumY += numbers[i]
// 			sumXY += x[i] * numbers[i]
// 			sumX2 += x[i] * x[i]
// 		}

// 		n := float64(len(numbers))

// 		b := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
// 		a := (sumY - b*sumX) / n

// 		nextX := float64(len(numbers))
// 		lowerLimit := a + b*(nextX) - 1.5
// 		upperLimit := a + b*(nextX) + 0.5

// 		return lowerLimit, upperLimit
// 	}
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	var numbers []float64
// 	var oldUpperLimit float64

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		number, err := strconv.ParseFloat(line, 64)
// 		if err != nil {
// 			fmt.Println("Invalid input. Please enter a valid number.")
// 			continue
// 		}
// 		if number > 0 {

// 			numbers = append(numbers, number)
// 			lowerLimit, upperLimit := calculateRange(numbers)
// 			if upperLimit > oldUpperLimit {
// 				oldUpperLimit = upperLimit
// 				// oldLowerLomit = lowerLimit
// 				fmt.Printf("%d %d\n", int(math.Round(lowerLimit)), int(math.Round(upperLimit)))
// 			} else {
// 				fmt.Printf("%d %d\n", int(math.Round(number-10)), int(math.Round(number+10)))
// 			}

// 		}
// 	}

//		if err := scanner.Err(); err != nil {
//			fmt.Println("Error reading input:", err)
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Check if a file path argument is provided
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run your-program.go")
		return
	}

	// Initialize variables for calculation
	var x, y, xy, xSquared, ySquared float64
	count := 0
	var yValues []float64

	// Create a scanner to read data from standard input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter data values (one per line, empty line to finish):")

	// Read data from standard input
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // Stop reading on empty line
		}

		value := float64(count)
		yValue, err := float64FromString(line)
		if err != nil {
			log.Fatalf("Error parsing data: %v", err)
		}

		// Update variables for calculations
		x += value
		y += yValue
		xy += value * yValue
		xSquared += value * value
		ySquared += yValue * yValue
		yValues = append(yValues, yValue)

		count++
	}

	// Calculate Linear Regression Line
	m := (float64(count)*xy - x*y) / (float64(count)*xSquared - x*x)
	b := (y - m*x) / float64(count)

	// Calculate Pearson Correlation Coefficient
	r := (float64(count)*xy - x*y) / (math.Sqrt((float64(count)*xSquared - x*x) * (float64(count)*ySquared - y*y)))

	// Calculate the standard error of the estimate (SE)
	// SE = sqrt((1-r^2)*SSE/(n-2)), where SSE is the sum of squared errors
	SSE := ySquared - (m*xy + b*y - m*x*x)
	SE := math.Sqrt((1 - r*r) * SSE / float64(count-2))

	// Calculate the prediction range for the next value (95% confidence interval)
	confidenceLevel := 0.95
	t := criticalTValue(confidenceLevel, float64(count-2))
	predictionRange := t * SE

	// Print the results
	fmt.Printf("Linear Regression Line:\n")
	fmt.Printf("y = %.6fx + %.6f\n", m, b)
	fmt.Printf("Pearson Correlation Coefficient (r): %.10f\n", r)
	fmt.Printf("Prediction Range for Next Value (95%% Confidence Interval): ±%.6f\n", predictionRange)

	// Sort yValues for prediction range calculation
	sort.Float64s(yValues)

	// Calculate the predicted value for the next data point
	nextX := float64(count)
	nextPredictedY := m*nextX + b

	// Calculate the lower and upper bounds of the prediction range
	lowerBoundIndex := int((1.0 - confidenceLevel) * float64(count))
	upperBoundIndex := int(confidenceLevel * float64(count))
	lowerBound := yValues[lowerBoundIndex]
	upperBound := yValues[upperBoundIndex]

	fmt.Printf("Predicted Value for Next Data Point: %.6f\n", nextPredictedY)
	fmt.Printf("Prediction Range for Next Value (95%% Confidence Interval): [%.6f, %.6f]\n", lowerBound, upperBound)
}

// Helper function to convert a string to a float64
func float64FromString(s string) (float64, error) {
	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// Calculate the critical t-value for a given confidence level and degrees of freedom
func criticalTValue(confidenceLevel float64, df float64) float64 {
	// Using Student's t-distribution to calculate the t-value
	t := math.Abs(math.Round(math.Sqrt(df) * math.Abs(math.Tanh(math.Acos(confidenceLevel)))))
	return t
}

