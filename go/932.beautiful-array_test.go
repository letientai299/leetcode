package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_beautifulArray(t *testing.T) {
	tests := 1000
	for i := 3; i <= tests; i++ {
		n := i
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			got := beautifulArray(n)
			seen := make(map[int]bool)
			for _, v := range got {
				seen[v] = true
			}

			if !assert.Equalf(t, n, len(seen), "not enough value: %v", got) {
				return
			}

			for i := 0; i < n-2; i++ {
				for j := i + 2; j < n; j++ {
					for k := i + 1; k < j; k++ {
						if got[i]+got[j] == 2*got[k] {
							t.Errorf("invalid triplet (i=%d:%d, j=%d:%d, k=%d:%d)", i, got[i], j, got[j], k, got[k])
							t.Log(got)
							return
						}
					}
				}
			}
		})
	}
}
