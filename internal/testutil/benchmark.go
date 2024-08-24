package testutil

import (
	"fmt"
	"time"
)

type Benchmark struct {
	name      string
	startTime time.Time
}

func StartBenchmark(name string) *Benchmark {
	return &Benchmark{
		name:      name,
		startTime: time.Now(),
	}
}

func (b *Benchmark) Print() {
	elapsed := time.Since(b.startTime)
	fmt.Printf("=====================================\n")
	fmt.Printf("> \"%s\" took %dms\n", b.name, elapsed.Milliseconds())
	fmt.Printf("=====================================\n")
}
