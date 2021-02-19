package mysql

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IniterDatabase struct {
	conn *gorm.DB
}

func NewInitDatabase() *IniterDatabase {
	conn, err := gorm.Open(mysql.Open(os.Getenv("MYSQL_ADDR")), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		log.Fatal("Create mysql connection error")
		return nil
	}
	return &IniterDatabase{
		conn: conn,
	}
}

func (i *IniterDatabase) InsertBatches(ctx context.Context, daos interface{}) error {
	if err := i.conn.WithContext(ctx).CreateInBatches(daos, 100).Error; err != nil {
		log.Fatal("Insert error")
	}

	return nil
}

func (i *IniterDatabase) Insert(ctx context.Context, daos interface{}) error {
	if err := i.conn.WithContext(ctx).Create(daos).Error; err != nil {
		log.Fatal("Insert error")
	}

	return nil
}

func (i *IniterDatabase) Update(ctx context.Context, dao interface{}) error {
	if err := i.conn.WithContext(ctx).Model(dao).Updates(dao).Error; err != nil {
		log.Fatal("Update Error")
	}

	return nil
}

func (i *IniterDatabase) Upsert(ctx context.Context, id string, dao interface{}) error {
	cpDao := dao
	cpDao, err := i.FindById(ctx, id, cpDao)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return i.Insert(ctx, dao)
	}
	if err != nil {
		return err
	}

	return i.Update(ctx, dao)
}

func (i *IniterDatabase) FindById(ctx context.Context, id string, dao interface{}) (interface{}, error) {
	if err := i.conn.WithContext(ctx).Where(map[string]interface{}{"id": id}).First(dao).Error; err != nil {
		return nil, err
	}
	return dao, nil
}

func (i *IniterDatabase) ReadJsondataFromFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read file error or empty: %w", err)
	}

	return string(data), nil
}
