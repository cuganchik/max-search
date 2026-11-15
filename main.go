package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for index := range data {
		data[index] = rand.Intn(1_000_000) + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValues := make([]int, CHUNKS)
	hasValue := make([]bool, CHUNKS)
	chunkSize := len(data) / CHUNKS
	remainder := len(data) % CHUNKS

	var waitGroup sync.WaitGroup
	start := 0
	for chunkIndex := 0; chunkIndex < CHUNKS; chunkIndex++ {
		end := start + chunkSize
		if chunkIndex < remainder {
			end++
		}

		chunk := data[start:end]
		waitGroup.Add(1)
		go func(index int, values []int) {
			defer waitGroup.Done()
			if len(values) > 0 {
				maxValues[index] = maximum(values)
				hasValue[index] = true
			}
		}(chunkIndex, chunk)

		start = end
	}

	waitGroup.Wait()
	max := 0
	foundValue := false
	for index, value := range maxValues {
		if hasValue[index] && (!foundValue || value > max) {
			max = value
			foundValue = true
		}
	}

	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
