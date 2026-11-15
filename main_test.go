package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "положительный размер", size: 100},
		{name: "нулевой размер", size: 0},
		{name: "отрицательный размер", size: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := generateRandomElements(test.size)
			expectedLength := max(test.size, 0)
			if len(data) != expectedLength {
				t.Fatalf("длина слайса = %d, ожидалось %d", len(data), expectedLength)
			}
			for _, value := range data {
				if value <= 0 {
					t.Fatalf("сгенерировано неположительное число: %d", value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{9, 5, 1}, want: 9},
		{name: "максимум в конце", data: []int{1, 5, 9}, want: 9},
		{name: "отрицательные значения", data: []int{-9, -5, -1}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", test.data, got, test.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := [][]int{
		{},
		{42},
		{1, 9, 5, 3, 9, 2, 7, 8, 4},
		{-9, -5, -1},
	}

	for _, data := range tests {
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("maxChunks(%v) = %d, ожидалось %d", data, got, want)
		}
	}
}
