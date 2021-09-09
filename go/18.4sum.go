package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int
	for p := 0; p < n-3; {
		a := nums[p]
		q := n - 1
		for q-2 > p {
			d := nums[q]
			for j, k := p+1, q-1; j < k; {
				b := nums[j]
				c := nums[k]
				if a+b+c+d > target {
					k--
					continue
				}
				if a+b+c+d < target {
					j++
					continue
				}

				res = append(res, []int{a, b, c, d})
				for j < k && nums[j] == b {
					j++
				}

				for k > j && nums[k] == c {
					k--
				}
			}

			q--
			for q-2 > p && nums[q] == d {
				q--
			}
		}

		for p < n-3 && nums[p] == a {
			p++
		}
	}
	return res
}
