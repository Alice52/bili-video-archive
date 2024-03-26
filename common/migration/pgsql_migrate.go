package migration

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/alice52/archive/common/global"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

// InitializePgsql 初始化函数, 在项目启动时调用
func InitializePgsql(db *gorm.DB) {
	mp := global.CONFIG.Pgsql.MigrationPath

	if len(mp) == 0 {
		return
	}

	if s, err := db.DB(); err != nil {
		panic(err)
	} else if err := MigratePgsql(s, mp); err != nil { // 执行数据库迁移
		panic(err)
	}
}

// MigratePgsql 执行数据库迁移
func MigratePgsql(db *sql.DB, mp string) error {
	// 创建迁移实例
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	defer func(driver database.Driver) {
		err := driver.Close()
		if err != nil {
			panic(err)
		}
	}(driver)

	m, err := migrate.NewWithDatabaseInstance(mp, "postgres", driver)
	if err != nil {
		return err
	}

	// 执行迁移
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("Database migration successful!")
	return nil
}
