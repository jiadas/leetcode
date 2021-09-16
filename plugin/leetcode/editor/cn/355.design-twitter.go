package golang

import (
	"container/list"
)

// leetcode submit region begin(Prohibit modification and deletion)

type Twitter struct {
	recentMax     int
	users         map[int]*User
	time          int
	timeByTweetID map[int]int
}

type User struct {
	follows map[int]struct{}
	tweets  *list.List
}

func NewUser() *User {
	return &User{
		follows: make(map[int]struct{}),
		tweets:  list.New(),
	}
}

// Constructor initialize your data structure here.
func Constructor() Twitter {
	return Twitter{
		recentMax:     10,
		users:         make(map[int]*User),
		timeByTweetID: make(map[int]int),
	}
}

// PostTweet compose a new tweet.
func (t *Twitter) PostTweet(userId int, tweetId int) {
	user, ok := t.users[userId]
	if !ok {
		user = NewUser()
		t.users[userId] = user
	}

	if user.tweets.Len() == t.recentMax {
		user.tweets.Remove(user.tweets.Back())
	}

	t.time++
	t.timeByTweetID[tweetId] = t.time
	user.tweets.PushFront(tweetId)
}

// GetNewsFeed retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
func (t *Twitter) GetNewsFeed(userId int) []int {
	user, ok := t.users[userId]
	if !ok {
		return nil
	}
	var ans []int
	for e := user.tweets.Front(); e != nil; e = e.Next() {
		ans = append(ans, e.Value.(int))
	}
	for followeeID := range user.follows {
		if followeeID == userId {
			continue
		}

		followee, ok := t.users[followeeID]
		if !ok {
			continue
		}

		res := make([]int, 0, t.recentMax)
		i := 0
		it := followee.tweets.Front()
		for i < len(ans) && it != nil {
			if t.timeByTweetID[it.Value.(int)] > t.timeByTweetID[ans[i]] {
				res = append(res, it.Value.(int))
				it = it.Next()
			} else {
				res = append(res, ans[i])
				i++
			}
			if len(res) == t.recentMax {
				break
			}
		}

		for ; i < len(ans) && len(res) < t.recentMax; i++ {
			res = append(res, ans[i])
		}
		for ; it != nil && len(res) < t.recentMax; it = it.Next() {
			res = append(res, it.Value.(int))
		}

		ans = res
	}
	return ans
}

// Follow follower follows a followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Follow(followerId int, followeeId int) {
	user, ok := t.users[followerId]
	if !ok {
		user = NewUser()
		t.users[followerId] = user
	}
	user.follows[followeeId] = struct{}{}
}

// Unfollow Follower unfollows a followee. If the operation is invalid, it should be a no-op.
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	user, ok := t.users[followerId]
	if !ok {
		return
	}
	delete(user.follows, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// leetcode submit region end(Prohibit modification and deletion)
