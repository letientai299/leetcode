package main

import "container/heap"

// Design Twitter
//
// Medium
//
// https://leetcode.com/problems/design-twitter/
//
// Design a simplified version of Twitter where users can post tweets,
// follow/unfollow another user, and is able to see the `10` most recent tweets
// in the user's news feed.
//
// Implement the `Twitter` class:
//
// - `Twitter()` Initializes your twitter object.
// - `void postTweet(int userId, int tweetId)` Composes a new tweet with ID
// `tweetId` by the user `userId`. Each call to this function will be made with
// a unique `tweetId`.
// - `List<Integer> getNewsFeed(int userId)` Retrieves the `10` most recent
// tweet IDs in the user's news feed. Each item in the news feed must be posted
// by users who the user followed or by the user themself. Tweets must be
// **ordered from most recent to least recent**.
// - `void follow(int followerId, int followeeId)` The user with ID `followerId`
// started following the user with ID `followeeId`.
// - `void unfollow(int followerId, int followeeId)` The user with ID
// `followerId` started unfollowing the user with ID `followeeId`.
//
// **Example 1:**
//
// ```
// Input
// ["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed",
// "unfollow", "getNewsFeed"]
// [[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
// Output
// [null, null, [5], null, null, [6, 5], null, [5]]
//
// Explanation
// Twitter twitter = new Twitter();
// twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1
// tweet id -> [5]. return [5]
// twitter.follow(1, 2);    // User 1 follows user 2.
// twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2
// tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is
// posted after tweet id 5.
// twitter.unfollow(1, 2);  // User 1 unfollows user 2.
// twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1
// tweet id -> [5], since user 1 is no longer following user 2.
//
// ```
//
// **Constraints:**
//
// - `1 <= userId, followerId, followeeId <= 500`
// - `0 <= tweetId <= 104`
// - All the tweets have **unique** IDs.
// - At most `3 * 104` calls will be made to `postTweet`, `getNewsFeed`,
// `follow`, and `unfollow`.

type Twitter struct {
	time  int // simulate posting time
	users map[int]*twitterUser
}

type twitterUser struct {
	following map[int]*twitterUser
	tweets    *tweet
}

type tweet struct {
	id   int
	time int
	next *tweet
}

func Constructor() Twitter {
	return Twitter{
		users: make(map[int]*twitterUser),
	}
}

func (tw *Twitter) PostTweet(userId int, tweetId int) {
	u, ok := tw.users[userId]
	if !ok {
		u = &twitterUser{following: make(map[int]*twitterUser)}
	}

	tw.time++
	t := &tweet{
		id:   tweetId,
		time: tw.time,
		next: u.tweets,
	}
	u.tweets = t
	tw.users[userId] = u
}

func (tw *Twitter) Follow(a int, b int) {
	x, ok := tw.users[a]
	if !ok {
		x = &twitterUser{following: make(map[int]*twitterUser)}
	}

	y, ok := tw.users[b]
	if !ok {
		y = &twitterUser{following: make(map[int]*twitterUser)}
	}

	x.following[b] = y
	tw.users[a] = x
	tw.users[b] = y
}

func (tw *Twitter) Unfollow(a int, b int) {
	x, ok := tw.users[a]
	if !ok {
		x = &twitterUser{following: make(map[int]*twitterUser)}
	}

	y, ok := tw.users[b]
	if !ok {
		y = &twitterUser{following: make(map[int]*twitterUser)}
	}

	delete(x.following, b)
	tw.users[a] = x
	tw.users[b] = y
}

func (tw *Twitter) GetNewsFeed(userId int) []int {
	u, ok := tw.users[userId]
	if !ok {
		return nil
	}

	res := make([]int, 0, 10)

	// merge k-sorted list (problem 23)
	// A medium problem that requires solution of a hard problem to solve!
	lists := make([]*tweet, 0, 1+len(u.following))
	if u.tweets != nil {
		lists = append(lists, u.tweets)
	}
	for _, o := range u.following {
		if o.tweets != nil {
			lists = append(lists, o.tweets)
		}
	}

	h := tweetHeap(lists)
	heap.Init(&h)
	for h.Len() > 0 && len(res) < 10 {
		v := heap.Pop(&h).(*tweet)
		res = append(res, v.id)
		if v.next != nil {
			heap.Push(&h, v.next)
		}
	}

	return res
}

type tweetHeap []*tweet

func (t tweetHeap) Len() int           { return len(t) }
func (t tweetHeap) Less(i, j int) bool { return t[i].time > t[j].time }
func (t *tweetHeap) Swap(i, j int)     { (*t)[i], (*t)[j] = (*t)[j], (*t)[i] }

func (t *tweetHeap) Push(x interface{}) {
	v := x.(*tweet)
	*t = append(*t, v)
}

func (t *tweetHeap) Pop() interface{} {
	v := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return v
}
