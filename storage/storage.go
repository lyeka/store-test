package storage

import (
	"errors"
	"github.com/lyeka/store-test/model"
	"github.com/lyeka/store-test/storage/mysql"
	"log"
	"os"
)

type Storage interface {
	Init() error
	Close() error

	CreateUser(u *model.User) (int64, error)
	CreatePost(p *model.Post) (int64, error)
	CreateComment(c *model.Comment) (int64, error)

	GetUser(uid int64) (*model.User, error)
	GetPost(pid int64) (*model.Post, error)
	GetPostComment(pid int64) ([]*model.Comment, error)
	FindUser(u *model.User) ([]*model.User, error)
	FindPost(p *model.Post) ([]*model.Post, error)
}

var Instance Storage

func Init() error {
	var storageInstance Storage
	var err error

	storeDrive := os.Getenv("STORAGE_DRIVE")
	switch storeDrive {
	case "mysql":
		storageInstance, err = mysql.NewStorage()
	case "redis":
	case "mongo":
	default:
		return errors.New("wrong storage drive")
	}

	if err != nil {
		return err
	}

	Instance = storageInstance

	err = Instance.Init()
	if err != nil {
		return err
	}

	log.Println("init storage success.")
	return nil
}
