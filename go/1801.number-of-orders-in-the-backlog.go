package main

import "container/heap"

// Number of Orders in the Backlog
//
// Medium
//
// https://leetcode.com/problems/number-of-orders-in-the-backlog/
//
// You are given a 2D integer array `orders`, where each `orders[i] = [pricei,
// amounti, orderTypei]` denotes that `amounti`orders have been placed of type
// `orderTypei` at the price `pricei`. The `orderTypei` is:
//
// - `0` if it is a batch of `buy` orders, or
// - `1` if it is a batch of `sell` orders.
//
// Note that `orders[i]` represents a batch of `amounti` independent orders with
// the same price and order type. All orders represented by `orders[i]` will be
// placed before all orders represented by `orders[i+1]` for all valid `i`.
//
// There is a **backlog** that consists of orders that have not been executed.
// The backlog is initially empty. When an order is placed, the following
// happens:
//
// - If the order is a `buy` order, you look at the `sell` order with the
// **smallest** price in the backlog. If that `sell` order's price is **smaller
// than or equal to** the current `buy` order's price, they will match and be
// executed, and that `sell` order will be removed from the backlog. Else, the
// `buy` order is added to the backlog.
// - Vice versa, if the order is a `sell` order, you look at the `buy` order
// with the **largest** price in the backlog. If that `buy` order's price is
// **larger than or equal to** the current `sell` order's price, they will match
// and be executed, and that `buy` order will be removed from the backlog. Else,
// the `sell` order is added to the backlog.
//
// Return _the total **amount** of orders in the backlog after placing all the
// orders from the input_. Since this number can be large, return it **modulo**
// `109 + 7`.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/03/11/ex1.png)
//
// ```
// Input: orders = [[10,5,0],[15,2,1],[25,1,1],[30,4,0]]
// Output: 6
// Explanation: Here is what happens with the orders:
// - 5 orders of type buy with price 10 are placed. There are no sell orders, so
// the 5 orders are added to the backlog.
// - 2 orders of type sell with price 15 are placed. There are no buy orders
// with prices larger than or equal to 15, so the 2 orders are added to the
// backlog.
// - 1 order of type sell with price 25 is placed. There are no buy orders with
// prices larger than or equal to 25 in the backlog, so this order is added to
// the backlog.
// - 4 orders of type buy with price 30 are placed. The first 2 orders are
// matched with the 2 sell orders of the least price, which is 15 and these 2
// sell orders are removed from the backlog. The 3rd order is matched with the
// sell order of the least price, which is 25 and this sell order is removed
// from the backlog. Then, there are no more sell orders in the backlog, so the
// 4th order is added to the backlog.
// Finally, the backlog has 5 buy orders with price 10, and 1 buy order with
// price 30. So the total number of orders in the backlog is 6.
//
// ```
//
// **Example 2:**
//
// ![](https://assets.leetcode.com/uploads/2021/03/11/ex2.png)
//
// ```
// Input: orders = [[7,1000000000,1],[15,3,0],[5,999999995,0],[5,1,1]]
// Output: 999999984
// Explanation: Here is what happens with the orders:
// - 109 orders of type sell with price 7 are placed. There are no buy orders,
// so the 109 orders are added to the backlog.
// - 3 orders of type buy with price 15 are placed. They are matched with the 3
// sell orders with the least price which is 7, and those 3 sell orders are
// removed from the backlog.
// - 999999995 orders of type buy with price 5 are placed. The least price of a
// sell order is 7, so the 999999995 orders are added to the backlog.
// - 1 order of type sell with price 5 is placed. It is matched with the buy
// order of the highest price, which is 5, and that buy order is removed from
// the backlog.
// Finally, the backlog has (1000000000-3) sell orders with price 7, and
// (999999995-1) buy orders with price 5. So the total number of orders =
// 1999999991, which is equal to 999999984 % (109 + 7).
//
// ```
//
// **Constraints:**
//
// - `1 <= orders.length <= 105`
// - `orders[i].length == 3`
// - `1 <= pricei, amounti <= 109`
// - `orderTypei` is either `0` or `1`.
func getNumberOfBacklogOrders(orders [][]int) int {
	var buys buyOrderHeap1801
	var sells sellOrderHeap1801

	for _, o := range orders {
		if o[2] == 0 {
			// buy
			b := buyOrder1801{price: o[0], amount: o[1]}
			for b.amount > 0 && sells.Len() > 0 && sells[0].price <= b.price {
				if sells[0].amount <= b.amount {
					b.amount -= sells[0].amount
					heap.Pop(&sells)
				} else {
					sells[0].amount -= b.amount
					b.amount = 0
				}
			}

			if b.amount > 0 {
				heap.Push(&buys, b)
			}
			continue
		}

		// buy
		s := sellOrder1801{price: o[0], amount: o[1]}
		for s.amount > 0 && buys.Len() > 0 && buys[0].price >= s.price {
			if buys[0].amount <= s.amount {
				s.amount -= buys[0].amount
				heap.Pop(&buys)
			} else {
				buys[0].amount -= s.amount
				s.amount = 0
			}
		}

		if s.amount > 0 {
			heap.Push(&sells, s)
		}
	}

	r := 0
	const mod = 1e9 + 7
	for _, b := range buys {
		r = (r + b.amount) % mod
	}

	for _, b := range sells {
		r = (r + b.amount) % mod
	}
	return r
}

type buyOrder1801 struct {
	price  int
	amount int
}

type buyOrderHeap1801 []buyOrder1801

func (o buyOrderHeap1801) Len() int            { return len(o) }
func (o buyOrderHeap1801) Less(i, j int) bool  { return o[i].price > o[j].price }
func (o buyOrderHeap1801) Swap(i, j int)       { o[i], o[j] = o[j], o[i] }
func (o *buyOrderHeap1801) Push(x interface{}) { *o = append(*o, x.(buyOrder1801)) }
func (o *buyOrderHeap1801) Pop() interface{} {
	n := len(*o) - 1
	v := (*o)[n]
	*o = (*o)[:n]
	return v
}

type sellOrder1801 struct {
	price  int
	amount int
}

type sellOrderHeap1801 []sellOrder1801

func (o sellOrderHeap1801) Len() int            { return len(o) }
func (o sellOrderHeap1801) Less(i, j int) bool  { return o[i].price < o[j].price }
func (o sellOrderHeap1801) Swap(i, j int)       { o[i], o[j] = o[j], o[i] }
func (o *sellOrderHeap1801) Push(x interface{}) { *o = append(*o, x.(sellOrder1801)) }
func (o *sellOrderHeap1801) Pop() interface{} {
	n := len(*o) - 1
	v := (*o)[n]
	*o = (*o)[:n]
	return v
}
