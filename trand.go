// Package trand provides utilities for running deterministic randomized tests.
package trand

import (
	"flag"
	"fmt"
	"math/rand"
	"testing"
)

var flagSeed int64

func init() {
	flag.Int64Var(
		&flagSeed,
		"trand.seed",
		0,
		"If specified, run trand tests only with the given seed. Useful for re-running particular "+
			"failures.",
	)
}

// Runs the given test f n times with a different seed each time. f should be deterministic, so that
// re-running a failed seed with --trand.seed produces the same results.
//
// If the --trand.seed flag is set, runs the test just once with --trand.seed.
func RandomN(t *testing.T, n int, f func(t *testing.T, r *rand.Rand)) {
	if flagSeed != 0 {
		run(t, flagSeed, f)
		return
	}

	for i := 0; i < n; i++ {
		run(t, newSeed(), f)
	}
}

func run(t *testing.T, seed int64, f func(t *testing.T, r *rand.Rand)) {
	// Use --trand.seed=[seed] as the test name, so it's obvious how to re-run failed tests.
	t.Run(fmt.Sprintf("--trand.seed=%d", seed), func(t *testing.T) {
		f(t, rand.New(rand.NewSource(seed)))
	})
}

func newSeed() int64 {
	for {
		seed := rand.Int63()
		// Just to be safe, since 0 means not set for flagSeed.
		if seed != 0 {
			return seed
		}
	}
}
