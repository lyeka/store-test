package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/lyeka/store-test/model"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var TableList = []string{
	"user",
	"post",
	"comment",
}

type Storage struct {
	db *sqlx.DB
}

func NewStorage() (*Storage, error) {
	mysqlDsn := os.Getenv("MYSQL_DSN")
	db, err := sqlx.Connect("mysql", mysqlDsn)
	return &Storage{
		db: db,
	}, err
}

func (s *Storage) Init() error {
	return s.db.Ping()
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) CreateUser(u *model.User) (int64, error) {
	res, err := s.db.Exec("insert into user (name, mobile) values (?, ?)", u.Name, u.Mobile)
	if err != nil {
		return 0, err
	}

	uid, _ := res.LastInsertId()
	return uid, nil
}

func (s *Storage) CreatePost(u *model.Post) (int64, error) {
	res, err := s.db.Exec("insert into post (uid, title, content) values(?, ?, ?)", u.Uid, u.Title, u.Content)
	if err != nil {
		return 0, err
	}

	pid, _ := res.LastInsertId()
	return pid, nil
}

func (s *Storage) CreateComment(c *model.Comment) (int64, error) {
	res, err := s.db.Exec("insert into comment (pid, cid, uid, content) values (?, ?, ?, ?)", c.Pid, c.Cid, c.Uid, c.Content)
	if err != nil {
		return 0, nil
	}

	cid, _ := res.LastInsertId()
	return cid, nil
}

func (s *Storage) GetUser(uid int64) (*model.User, error) {
	user := model.User{}
	err := s.db.Select(&user, "select * from user where id = ?", uid)
	return &user, err
}

func (s *Storage) GetPost(pid int64) (*model.Post, error) {
	post := model.Post{}
	err := s.db.Select(&post, "select * from post where id = ?", pid)
	return &post, err
}

func (s *Storage) GetPostComment(pid int64) ([]*model.Comment, error) {
	var postComments []*model.Comment
	err := s.db.Select(&postComments, "select * from comment where pid = ?", pid)
	return postComments, err
}

func (s *Storage) FindUser(u *model.User) ([]*model.User, error) {

	return nil, nil
}

func (s *Storage) FindPost(p *model.Post) ([]*model.Post, error) {
	return nil, nil
}

// 清空指定数据表
// 仅用于测试时前置调用清理数据
func (s *Storage) truncateTables(tables []string) {
	for _, table := range tables {
		s.db.MustExec("truncate table " + table)
	}
}
