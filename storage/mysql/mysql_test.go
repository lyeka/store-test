package mysql

import (
	"github.com/lyeka/store-test/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMysqlEngine(t *testing.T) {
	db, err := NewStorage()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	// todo db清理工作
	db.truncateTables(TableList)

	user := &model.User{
		Name:   "yeka",
		Mobile: "18866668888",
	}

	uid, err := db.CreateUser(user)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, int64(1), uid)

	post := &model.Post{
		Uid:     uid,
		Title:   "post t1",
		Content: "post c1",
	}
	pid, err := db.CreatePost(post)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, int64(1), pid)

	comment := &model.Comment{
		Pid:     pid,
		Uid:     uid,
		Content: "comment 1",
	}
	cid, err := db.CreateComment(comment)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, int64(1), cid)

	comment2 := &model.Comment{
		Pid:     pid,
		Cid:     cid,
		Uid:     uid,
		Content: "comment comment 1",
	}
	cid2, err := db.CreateComment(comment2)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.Equal(t, int64(2), cid2)
}
