package main

import (
	"fmt"
	"testing"
)

func TestTwitter(t *testing.T) {
	tw := Constructor355()
	tw.PostTweet(1, 5)
	fmt.Println(tw.GetNewsFeed(1))
	tw.Follow(1, 2)
	tw.PostTweet(2, 6)
	fmt.Println(tw.GetNewsFeed(1))
}
