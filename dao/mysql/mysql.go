package mysql

import (
	"acana/models"
	"acana/setting"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *setting.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	fmt.Printf("%v\n", dsn)

	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed:", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	// test for connect mysql

	// communities, _ := GetCommunityList()
	// for _, community := range communities {
	// 	fmt.Printf("ID: %d, Name: %s\n", community.ID, community.Name)
	// }

	return
}

func GetCommunityList_() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

func Close() {
	_ = db.Close()
}
