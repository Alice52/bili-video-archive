package migration

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/alice52/archive/common/global"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// InitializeMysql 初始化函数, 在项目启动时调用
func InitializeMysql(db *gorm.DB) {
	mp := global.CONFIG.Mysql.MigrationPath

	if len(mp) == 0 {
		return
	}

	if s, err := db.DB(); err != nil {
		panic(err)
	} else if err := MigrateMysql(s, mp); err != nil { // 执行数据库迁移
		panic(err)
	}
}

// MigrateMysql 执行数据库迁移
func MigrateMysql(db *sql.DB, mp string) error {
	// Create migration instance
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}
	defer func(driver database.Driver) {
		err := driver.Close()
		if err != nil {
			panic(err)
		}
	}(driver)

	m, err := migrate.NewWithDatabaseInstance(mp, "mysql", driver)
	if err != nil {
		return err
	}

	// Perform migration
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("Database migration successful!")
	return nil
}
