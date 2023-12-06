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

// Initialize 初始化函数, 在项目启动时调用
func Initialize(db *gorm.DB) {
	mp := global.CONFIG.Pgsql.MigrationPath
	if len(mp) == 0 {
		return
	}

	s, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 执行数据库迁移
	if err := MigrateDB(s, mp); err != nil {
		panic(err)
	}
}

// MigrateDB 执行数据库迁移
func MigrateDB(db *sql.DB, mp string) error {
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
