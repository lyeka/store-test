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

// 清空指定数据表
// 仅用于测试时前置调用清理数据
func (s *Storage) truncateTables(tables []string) {
	for _, table := range tables {
		s.db.MustExec("truncate table " + table)
	}
}
