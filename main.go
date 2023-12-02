package main

import (
	"log"
	"runtime"
	"time"
)

var storage [][]int

func main() {
	done := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				usage := make([]int, 0, 999999)
				storage = append(storage, usage)
			}
		}

	}()
	i := 0
	for i <= 1000 {
		select {
		case <-done:
			return
		case <-ticker.C:
			m := MemoryAsMB()
			log.Printf("alloc: %d mb", m.Alloc)
			log.Printf("Totalalloc: %d mb", m.Total)
			i++
		}
	}
	done <- true
}

type Mem struct {
	Total uint64
	Alloc uint64
}

// MemoryAsMB returns a Mem instance containing Alloc and TotalAlloc in megabyte int representation
func MemoryAsMB() Mem {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return Mem{Total: toMb(m.TotalAlloc), Alloc: toMb(m.Alloc)}
}

// toMb converts bytes to megabytes as uint64
func toMb(b uint64) uint64 {
	return b / 1024 / 1024
}
