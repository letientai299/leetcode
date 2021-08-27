package main

import "sort"

// Tweet Counts Per Frequency
//
// Medium
//
// https://leetcode.com/problems/tweet-counts-per-frequency/
//
// A social media company is trying to monitor activity on their site by
// analyzing the number of tweets that occur in select periods of time. These
// periods can be partitioned into smaller **time chunks** based on a certain
// frequency (every **minute**, **hour**, or **day**).
//
// For example, the period `[10, 10000]` (in **seconds**) would be partitioned
// into the following **time chunks** with these frequencies:
//
// - Every **minute** (60-second chunks): `[10,69]`, `[70,129]`, `[130,189]`,
// `...`, `[9970,10000]`
// - Every **hour** (3600-second chunks): `[10,3609]`, `[3610,7209]`,
// `[7210,10000]`
// - Every **day** (86400-second chunks): `[10,10000]`
//
// Notice that the last chunk may be shorter than the specified frequency's
// chunk size and will always end with the end time of the period (`10000` in
// the above example).
//
// Design and implement an API to help the company with their analysis.
//
// Implement the `TweetCounts` class:
//
// - `TweetCounts()` Initializes the `TweetCounts` object.
// - `void recordTweet(String tweetName, int time)` Stores the `tweetName` at
// the recorded `time` (in **seconds**).
// - `List<Integer> getTweetCountsPerFrequency(String freq, String tweetName,
// int startTime, int endTime)` Returns a list of integers representing the
// number of tweets with `tweetName` in each **time chunk** for the given period
// of time `[startTime, endTime]` (in **seconds**) and frequency `freq`.
//
//     - `freq` is one of `"minute"`, `"hour"`, or `"day"` representing a
// frequency of every **minute**, **hour**, or **day** respectively.
//
// **Example:**
//
// ```
// Input
// ["TweetCounts","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","getTweetCountsPerFrequency","recordTweet","getTweetCountsPerFrequency"]
// [[],["tweet3",0],["tweet3",60],["tweet3",10],["minute","tweet3",0,59],["minute","tweet3",0,60],["tweet3",120],["hour","tweet3",0,210]]
//
// Output
// [null,null,null,null,[2],[2,1],null,[4]]
//
// Explanation
// TweetCounts tweetCounts = new TweetCounts();
// tweetCounts.recordTweet("tweet3", 0);                              // New
// tweet "tweet3" at time 0
// tweetCounts.recordTweet("tweet3", 60);                             // New
// tweet "tweet3" at time 60
// tweetCounts.recordTweet("tweet3", 10);                             // New
// tweet "tweet3" at time 10
// tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 59); // return
// [2]; chunk [0,59] had 2 tweets
// tweetCounts.getTweetCountsPerFrequency("minute", "tweet3", 0, 60); // return
// [2,1]; chunk [0,59] had 2 tweets, chunk [60,60] had 1 tweet
// tweetCounts.recordTweet("tweet3", 120);                            // New
// tweet "tweet3" at time 120
// tweetCounts.getTweetCountsPerFrequency("hour", "tweet3", 0, 210);  // return
// [4]; chunk [0,210] had 4 tweets
//
// ```
//
// **Constraints:**
//
// - `0 <= time, startTime, endTime <= 109`
// - `0 <= endTime - startTime <= 104`
// - There will be at most `104` calls **in total** to `recordTweet` and
// `getTweetCountsPerFrequency`.

type TweetCounts struct {
	m map[string][]int
}

func Constructor1348() TweetCounts {
	return TweetCounts{
		m: make(map[string][]int),
	}
}

func (tw *TweetCounts) RecordTweet(s string, time int) {
	if tw.m[s] == nil {
		tw.m[s] = append(tw.m[s], time)
		return
	}

	times := tw.m[s]
	i := sort.SearchInts(times, time)
	times = append(times, 0)
	copy(times[i+1:], times[i:])
	times[i] = time
	tw.m[s] = times
}

func (tw *TweetCounts) GetTweetCountsPerFrequency(freq string, s string, startTime int, endTime int) []int {
	seconds := tw.secondOf(freq)
	counts := make([]int, (endTime-startTime)/seconds+1)
	start := sort.SearchInts(tw.m[s], startTime)
	end := sort.SearchInts(tw.m[s], endTime+1)
	for i := start; i < end; i++ {
		idx := (tw.m[s][i] - startTime) / seconds
		counts[idx]++
	}
	return counts
}

func (tw *TweetCounts) secondOf(freq string) int {
	switch freq {
	case "minute":
		return 60
	case "hour":
		return 3600
	default: // day
		return 86400
	}
}
