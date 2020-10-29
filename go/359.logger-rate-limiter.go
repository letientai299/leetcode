package main

/*
 * @lc app=leetcode id=359 lang=golang
 *
 * [359] Logger Rate Limiter
 *
 * https://leetcode.com/problems/logger-rate-limiter/description/
 *
 * algorithms
 * Easy (67.55%)
 * Total Accepted:    118.1K
 * Total Submissions: 165.5K
 * Testcase Example:  '["Logger","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage","shouldPrintMessage"]\n' +
  '[[],[1,"foo"],[2,"bar"],[3,"foo"],[8,"bar"],[10,"foo"],[11,"foo"]]'
 *
 * Design a logger system that receive stream of messages along with its
 * timestamps, each message should be printed if and only if it is not printed
 * in the last 10 seconds.
 *
 * Given a message and a timestamp (in seconds granularity), return true if the
 * message should be printed in the given timestamp, otherwise returns false.
 *
 * It is possible that several messages arrive roughly at the same time.
 *
 * Example:
 *
 *
 * Logger logger = new Logger();
 *
 * // logging string "foo" at timestamp 1
 * logger.shouldPrintMessage(1, "foo"); returns true;
 *
 * // logging string "bar" at timestamp 2
 * logger.shouldPrintMessage(2,"bar"); returns true;
 *
 * // logging string "foo" at timestamp 3
 * logger.shouldPrintMessage(3,"foo"); returns false;
 *
 * // logging string "bar" at timestamp 8
 * logger.shouldPrintMessage(8,"bar"); returns false;
 *
 * // logging string "foo" at timestamp 10
 * logger.shouldPrintMessage(10,"foo"); returns false;
 *
 * // logging string "foo" at timestamp 11
 * logger.shouldPrintMessage(11,"foo"); returns true;
 *
*/
type Logger struct {
	m map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{
		m: make(map[string]int),
	}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (g *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	ts, ok := g.m[message]
	if !ok {
		g.m[message] = timestamp
		return true
	}

	if timestamp-ts < 10 {
		return false
	}

	g.m[message] = timestamp
	return true
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
