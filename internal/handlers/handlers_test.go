package handlers

import (
	"reflect"
	"testing"
)

func TestCalculateExchanges(t *testing.T) {
	var tests = []struct {
		amount    int
		banknotes []int
		result    [][]int
	}{
		{
			600,
			[]int{500, 200, 100},
			[][]int{{500, 100}, {200, 200, 200}, {200, 200, 100, 100}, {200, 100, 100, 100, 100}, {100, 100, 100, 100, 100, 100}},
		},
		{
			500,
			[]int{200, 100, 50},
			[][]int{{200, 200, 100}, {200, 200, 50, 50}, {200, 100, 100, 100}, {200, 100, 100, 50, 50}, {200, 100, 50, 50, 50, 50},
				{200, 50, 50, 50, 50, 50, 50}, {100, 100, 100, 100, 100}, {100, 100, 100, 100, 50, 50}, {100, 100, 100, 50, 50, 50, 50},
				{100, 100, 50, 50, 50, 50, 50, 50}, {100, 50, 50, 50, 50, 50, 50, 50, 50}, {50, 50, 50, 50, 50, 50, 50, 50, 50, 50}},
		},
	}

	for _, test := range tests {
		if !reflect.DeepEqual(calculateExchanges(test.amount, test.banknotes), test.result) {
			t.Errorf("func calculateExchanges returned wrong result")
		}
	}

}
