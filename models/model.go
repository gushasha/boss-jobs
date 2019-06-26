package models

import (
	"fmt"
	"github.com/gushasha/boss-jobs/conf"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB

	username    string = conf.DB_USERNAME
	password    string = conf.DB_PASSWORD
	dbName      string = conf.DB_NAME
	tablePrefix string = conf.DB_TABLE_PREFIX
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreateTime int `json:"create_time"`
}

func init() {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}

	DB.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
	}
}
