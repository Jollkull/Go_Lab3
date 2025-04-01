package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
)

// Input struct for JSON decoding
type Input struct {
	Values []float64 `json:"values"`
}

func integrateGaussian(a, b float64, n int, power, error float64) float64 {
	// Implementation of Gaussian quadrature using the midpoint rule
	step := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x := a + (float64(i)+0.5)*step
		sum += math.Exp(-math.Pow((x-power)/error, 2)/2) / (error * math.Sqrt(2*math.Pi))
	}

	return sum * step
}

func calculateTask1(power, errorBefore, errorAfter, price float64) string {

	a := power - errorAfter // Lower integration limit
	b := power + errorAfter // Upper integration limit
	n := 1000               // Number of partitions (higher = more accurate)

	shareWithoutImbalancesBefore := integrateGaussian(a, b, n, power, errorBefore)
	profitBefore := power * 24 * shareWithoutImbalancesBefore * price
	fineBefore := power * 24 * (1 - shareWithoutImbalancesBefore) * price

	shareWithoutImbalancesAfter := integrateGaussian(a, b, n, power, errorAfter)
	profitAfter := power * 24 * shareWithoutImbalancesAfter * price
	fineAfter := power * 24 * (1 - shareWithoutImbalancesAfter) * price

	output := fmt.Sprintf(
		"Прибуток до вдосконалення: %.0f тис. грн \n"+
			"Штраф до вдосконалення: %.0f тис. грн\n"+
			"Виручка до вдосконалення: %.0f тис. грн\n"+
			"Прибуток після вдосконалення: %.0f тис. грн\n"+
			"Штраф після вдосконалення: %.0f тис. грн\n"+
			"Виручка після вдосконалення: %.0f тис. грн\n",
		math.Round(profitBefore),
		math.Round(fineBefore),
		math.Round(profitBefore-fineBefore),
		math.Round(profitAfter),
		math.Round(fineAfter),
		math.Round(profitAfter-fineAfter),
	)

	return output
}

func calculator1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input Input
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if len(input.Values) != 4 {
		http.Error(w, "Invalid number of inputs", http.StatusBadRequest)
		return
	}
	result := calculateTask1(input.Values[0], input.Values[1], input.Values[2], input.Values[3])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"result": result})
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/calculator1", calculator1Handler)

	fmt.Println("Server running at http://localhost:8083")
	http.ListenAndServe(":8083", nil)
}
