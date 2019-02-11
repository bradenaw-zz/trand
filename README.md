# `trand`

[![GoDoc](https://godoc.org/github.com/bradenaw/trand?status.svg)](https://godoc.org/github.com/bradenaw/trand)

Provides utilities for deterministic randomized tests.

At the moment, only has one function (`RandomN()`), used like this:

```
import (
    "math/rand"
    "sort"

    "github.com/bradenaw/trand"
)

func TestSort(t *testing.T) {
    trand.RandomN(t, 1000, func(t *testing.T, r *rand.Rand) {
        a := make([]int, r.Intn(2000))
        for i := range a {
            a[i] = r.Int()
        }

        sort.Slice(a, func(i, j) int { return a[i] < a[j] })

        for i := 1; i < len(a); i++ {
            if a[i] < a[i-1] {
                t.Fatal("list not in sorted order")
            }
        }
    })
}
```

This runs the test 1000 times with different seeds each time. Each seed is run as a subtest using
`t.Run()`. If any of the invocations fail, the name of the subtest contains a `--trand.seed` flag
that can be used to re-run just the failing seed, which makes it easy to (for example) locally debug
failures found in CI.
