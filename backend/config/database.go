package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(cfg DatabaseConfig) (*gorm.DB, error) {
	//mysql連線
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&clientFoundRows=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	//建立gorm db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 取得 GORM 底層的 *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	//確認是否順利連上db
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
