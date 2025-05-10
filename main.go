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
	if size == 0 {
		return []int{}
	}

	res := make([]int, size)

	for i := range size {
		res[i] = rand.Intn(SIZE) + 1 // +1, чтобы избежать нуля, который я возьму за ошибку
	}

	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	mx := data[0]

	for _, v := range data {
		mx = max(mx, v)
	}

	return mx
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < 8 {
		return maximum(data)
	}

	maxes := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS

	wg := sync.WaitGroup{}
	wg.Add(CHUNKS)

	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		go func(start, end int) {
			defer wg.Done()

			maxes[i] = maximum(data[start:end])
		}(start, end)
	}

	wg.Wait()

	return maximum(maxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	nums := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	timeStart := time.Now()
	max := maximum(nums)
	elapsed := time.Since(timeStart).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	timeStart = time.Now()
	max = maxChunks(nums)
	elapsed = time.Since(timeStart).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
