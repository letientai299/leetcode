package main

import (
	"sort"
	"strconv"
)

// Display Table of Food Orders in a Restaurant
//
// Medium
//
// https://leetcode.com/problems/display-table-of-food-orders-in-a-restaurant/
//
// Given the array `orders`, which represents the orders that customers have
// done in a restaurant. More
// specifically `orders[i]=[customerNamei,tableNumberi,foodItemi]` where
// `customerNamei` is the name of the customer, `tableNumberi` is the table
// customer sit at, and `foodItemi` is the item customer orders.
//
// _Return the restaurant's “ **display table**”_. The “ **display table**” is a
// table whose row entries denote how many of each food item each table ordered.
// The first column is the table number and the remaining columns correspond to
// each food item in alphabetical order. The first row should be a header whose
// first column is “Table”, followed by the names of the food items. Note that
// the customer names are not part of the table. Additionally, the rows should
// be sorted in numerically increasing order.
//
// **Example 1:**
//
// ```
// Input: orders = [["David","3","Ceviche"],["Corina","10","Beef
// Burrito"],["David","3","Fried
// Chicken"],["Carla","5","Water"],["Carla","5","Ceviche"],["Rous","3","Ceviche"]]
// Output: [["Table","Beef Burrito","Ceviche","Fried
// Chicken","Water"],["3","0","2","1","0"],["5","0","1","0","1"],["10","1","0","0","0"]]
// Explanation:
// The displaying table looks like:
// Table,Beef Burrito,Ceviche,Fried Chicken,Water
// 3    ,0           ,2      ,1            ,0
// 5    ,0           ,1      ,0            ,1
// 10   ,1           ,0      ,0            ,0
// For the table 3: David orders "Ceviche" and "Fried Chicken", and Rous orders
// "Ceviche".
// For the table 5: Carla orders "Water" and "Ceviche".
// For the table 10: Corina orders "Beef Burrito".
//
// ```
//
// **Example 2:**
//
// ```
// Input: orders = [["James","12","Fried Chicken"],["Ratesh","12","Fried
// Chicken"],["Amadeus","12","Fried Chicken"],["Adam","1","Canadian
// Waffles"],["Brianna","1","Canadian Waffles"]]
// Output: [["Table","Canadian Waffles","Fried
// Chicken"],["1","2","0"],["12","0","3"]]
// Explanation:
// For the table 1: Adam and Brianna order "Canadian Waffles".
// For the table 12: James, Ratesh and Amadeus order "Fried Chicken".
//
// ```
//
// **Example 3:**
//
// ```
// Input: orders = [["Laura","2","Bean Burrito"],["Jhon","2","Beef
// Burrito"],["Melissa","2","Soda"]]
// Output: [["Table","Bean Burrito","Beef Burrito","Soda"],["2","1","1","1"]]
//
// ```
//
// **Constraints:**
//
// - `1 <= orders.length <= 5 * 10^4`
// - `orders[i].length == 3`
// - `1 <= customerNamei.length, foodItemi.length <= 20`
// - `customerNamei` and `foodItemi` consist of lowercase and uppercase English
// letters and the space character.
// - `tableNumberi `is a valid integer between `1` and `500`.
func displayTable(orders [][]string) [][]string {
	var counts [][]int

	foods := make(map[string]int)
	var foodNames []string
	tables := make(map[string]int)
	var tableNum []string

	for _, o := range orders {
		table := o[1]
		food := o[2]
		fid, ok := foods[food]

		if !ok {
			fid = len(foods)
			foods[food] = fid
			foodNames = append(foodNames, food)
		}

		tid, ok := tables[table]
		if !ok {
			tid = len(tables)
			tables[table] = tid
			tableNum = append(tableNum, table)
		}

		for tid >= len(counts) {
			counts = append(counts, make([]int, len(foods)))
		}

		for fid >= len(counts[tid]) {
			counts[tid] = append(counts[tid], 0)
		}

		counts[tid][fid]++
	}

	sort.Strings(foodNames)
	sort.Slice(tableNum, func(i, j int) bool {
		a, b := tableNum[i], tableNum[j]
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)
		return x < y
	})

	var res [][]string
	res = append(res, []string{"Table"})

	for _, f := range foodNames {
		res[0] = append(res[0], f)
	}

	for i, t := range tableNum {
		res = append(res, make([]string, len(foodNames)+1))
		res[i+1][0] = t

		for j, f := range foodNames {
			tid := tables[t]
			fid := foods[f]
			c := 0
			if fid < len(counts[tid]) {
				c = counts[tid][fid]
			}
			res[i+1][j+1] = strconv.Itoa(c)
		}
	}

	return res
}
