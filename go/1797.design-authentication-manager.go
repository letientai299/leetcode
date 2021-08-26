package main

import "sort"

// Design Authentication Manager
//
// Medium
//
// https://leetcode.com/problems/design-authentication-manager/
//
// There is an authentication system that works with authentication tokens. For
// each session, the user will receive a new authentication token that will
// expire `timeToLive` seconds after the `currentTime`. If the token is renewed,
// the expiry time will be **extended** to expire `timeToLive` seconds after the
// (potentially different) `currentTime`.
//
// Implement the `AuthenticationManager` class:
//
// - `AuthenticationManager(int timeToLive)` constructs the
// `AuthenticationManager` and sets the `timeToLive`.
// - `generate(string tokenId, int currentTime)` generates a new token with the
// given `tokenId` at the given `currentTime` in seconds.
// - `renew(string tokenId, int currentTime)` renews the **unexpired** token
// with the given `tokenId` at the given `currentTime` in seconds. If there are
// no unexpired tokens with the given `tokenId`, the request is ignored, and
// nothing happens.
// - `countUnexpiredTokens(int currentTime)` returns the number of **unexpired**
// tokens at the given currentTime.
//
// Note that if a token expires at time `t`, and another action happens on time
// `t` (`renew` or `countUnexpiredTokens`), the expiration takes place
// **before** the other actions.
//
// **Example 1:**
//
// ![](https://assets.leetcode.com/uploads/2021/02/25/copy-of-pc68_q2.png)
//
// ```
// renewcountUnexpiredTokensrenewrenewcountUnexpiredTokenstimeToLiverenewcountUnexpiredTokensrenewrenewrenewrenewcountUnexpiredTokens
// ```
//
// **Constraints:**
//
// - `1 <= timeToLive <= 108`
// - `1 <= currentTime <= 108`
// - `1 <= tokenId.length <= 5`
// - `tokenId` consists only of lowercase letters.
// - All calls to `generate` will contain unique values of `tokenId`.
// - The values of `currentTime` across all the function calls will be
// **strictly increasing**.
// - At most `2000` calls will be made to all functions combined.

type AuthenticationManager struct {
	ttl     int
	expires []int
	m       map[string]int
}

func Constructor1797(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		ttl: timeToLive,
		m:   make(map[string]int),
	}
}

func (am *AuthenticationManager) Generate(tokenId string, currentTime int) {
	exp := currentTime + am.ttl
	am.m[tokenId] = exp
	am.expires = append(am.expires, exp)
}

func (am *AuthenticationManager) Renew(tokenId string, currentTime int) {
	exp, ok := am.m[tokenId]
	if !ok || exp <= currentTime {
		return
	}

	i := sort.Search(len(am.expires), func(i int) bool {
		return exp <= am.expires[i]
	})

	exp = currentTime + am.ttl
	am.m[tokenId] = exp
	am.expires = append(am.expires[:i], append(am.expires[i+1:], exp)...)
}

func (am *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	i := sort.Search(len(am.expires), func(i int) bool {
		return currentTime < am.expires[i]
	})
	am.expires = am.expires[i:]
	return len(am.expires)
}
