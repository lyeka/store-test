package model

type Comment struct {
	Id      int64 // 考虑 bigint ？
	Pid     int64 // post id
	Cid     int64 // 回复的评论id
	Uid     int64 // user id
	Content string
}
