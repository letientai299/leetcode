package main

import (
	"testing"
)

func TestTweetCounts(t *testing.T) {
	tw := Constructor1348()
	const s = "a"
	tw.RecordTweet(s, 857105800)
	//tw.RecordTweet(s, 89)
	//tw.RecordTweet(s, 10)
	r := tw.GetTweetCountsPerFrequency("minute", s, 255428777, 255438544)
	t.Log(r)
	//r = tw.GetTweetCountsPerFrequency("minute", s, 0, 60)
	//t.Log(r)

	// Test case 1:
	// ["TweetCounts","recordTweet","recordTweet","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","recordTweet","recordTweet","recordTweet","recordTweet"]
	// [[],           ["tweet0",33],["tweet1",89],["tweet2",99],["tweet3",53],["tweet4",3], ["hour","tweet0",89,3045],["tweet0",28],["tweet0",91],["tweet0",9],["tweet1",6]]

	// Test case 2:
	// ["TweetCounts","recordTweet","recordTweet","recordTweet","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","getTweetCountsPerFrequency"]
	// [[],           ["tweet0",13],["tweet1",16],["tweet2",12],["tweet3",18],["tweet4",82],["tweet3",89],["day","tweet0",89,9471],["hour","tweet3",13,4024]]

	// Test case 3:
	// ["TweetCounts","recordTweet","recordTweet","recordTweet","recordTweet","recordTweet","getTweetCountsPerFrequency","getTweetCountsPerFrequency","getTweetCountsPerFrequency","getTweetCountsPerFrequency"]
	// [[],           ["tweet0",857105800],["tweet1",255428777],["tweet2",13881700],["tweet3",82366032],["tweet4",334311127],["minute","tweet0",255428777,255438544],["day","tweet2",857105800,857108372],["minute","tweet4",334311127,334316350],["hour","tweet3",82366032,82370980]]
}
