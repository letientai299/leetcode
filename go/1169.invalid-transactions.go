package main

import (
	"strconv"
	"strings"
)

// Invalid Transactions
//
// Medium
//
// https://leetcode.com/problems/invalid-transactions/
//
// A transaction is possibly invalid if:
//
// - the amount exceeds `$1000`, or;
// - if it occurs within (and including) `60` minutes of another transaction
// with the **same name** in a **different city**.
//
// You are given an array of strings `transaction` where `transactions[i]`
// consists of comma-separated values representing the name, time (in minutes),
// amount, and city of the transaction.
//
// Return a list of `transactions` that are possibly invalid. You may return the
// answer in **any order**.
//
// **Example 1:**
//
// ```
// Input: transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
// Output: ["alice,20,800,mtv","alice,50,100,beijing"]
// Explanation: The first transaction is invalid because the second transaction
// occurs within a difference of 60 minutes, have the same name and is in a
// different city. Similarly the second one is invalid too.
// ```
//
// **Example 2:**
//
// ```
// Input: transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
// Output: ["alice,50,1200,mtv"]
//
// ```
//
// **Example 3:**
//
// ```
// Input: transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
// Output: ["bob,50,1200,mtv"]
//
// ```
//
// **Constraints:**
//
// - `transactions.length <= 1000`
// - Each `transactions[i]` takes the form `"{name},{time},{amount},{city}"`
// - Each `{name}` and `{city}` consist of lowercase English letters, and have
// lengths between `1` and `10`.
// - Each `{time}` consist of digits, and represent an integer between `0` and
// `1000`.
// - Each `{amount}` consist of digits, and represent an integer between `0` and
// `2000`.

// One of the stupidest problems I have seen on LeetCode!

func invalidTransactions(transactions []string) []string {
	type tx struct {
		name   string
		time   int
		amount int
		city   string
		src    string
		id     int
	}

	parse := func(i int, s string) tx {
		ss := strings.Split(s, ",")
		time, _ := strconv.Atoi(ss[1])
		amount, _ := strconv.Atoi(ss[2])
		return tx{
			id:     i,
			name:   ss[0],
			time:   time,
			amount: amount,
			city:   ss[3],
			src:    s,
		}
	}

	logs := make(map[string]map[string][]tx)
	invalids := make(map[string]int)
	seen := make(map[string]int)

	for i, t := range transactions {
		seen[t]++
		x := parse(i, t)

		if x.amount > 1000 {
			invalids[t]++
		}

		cities := logs[x.name]
		if len(cities) == 0 {
			logs[x.name] = make(map[string][]tx)
		}

		for city, times := range cities {
			if x.city == city {
				continue
			}

			marked := false
			for _, tx := range times {
				if tx.time-60 <= x.time && tx.time+60 >= x.time {
					invalids[tx.src]++
					if !marked {
						invalids[x.src]++
						marked = true
					}
				}
			}
		}

		logs[x.name][x.city] = append(logs[x.name][x.city], x)
	}

	var res []string
	for k, v := range invalids {
		if v > seen[k] {
			v = seen[k]
		}
		for i := 0; i < v; i++ {
			res = append(res, k)
		}
	}
	return res
}
