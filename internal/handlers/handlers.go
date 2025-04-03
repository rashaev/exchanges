package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Request struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}

type Response struct {
	Exchanges [][]int `json:"exchanges"`
}

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if logger, ok := r.Context().Value("logger").(*slog.Logger); ok {
		logger.Info("Получен HTTP запрос", "url", r.URL.Path)
	}

	// Вызов функции для расчета размена
	exchanges := calculateExchanges(req.Amount, req.Banknotes)

	// Формирование ответа
	resp := Response{Exchanges: exchanges}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func calculateExchanges(amount int, banknotes []int) [][]int {
	var result [][]int
	var current []int

	// Рекурсивная функция для поиска всех комбинаций
	var backtrack func(int, int)
	backtrack = func(start, remaining int) {
		if remaining == 0 {
			// Добавляем найденную комбинацию в результат
			result = append(result, append([]int{}, current...))
			return
		}

		for i := start; i < len(banknotes); i++ {
			if banknotes[i] > remaining {
				continue
			}
			current = append(current, banknotes[i])
			backtrack(i, remaining-banknotes[i])
			current = current[:len(current)-1] // Откат
		}
	}

	backtrack(0, amount)
	return result
}
