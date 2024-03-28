package main

import (
	"fmt"
	"github.com/alice52/jasypt-go"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var MySQLDSN string

func init() {
	if v, err := jasypt.New().Decrypt("wJ0ZxlPJVsTm/iaFWxAIa1M5jF+eYR9JvAFlnCdxzCbWJ+TQY093XQ4OKiqqmjBKzHxFIvOoiZ7/eMzuHBaU7zZ2UzzZyHj0/jbDoLIfqK6qu20k4ibz8BXR7bzUAFykoxSTNW00g1kWUPj5yiBWql/LrtkeKMmCusoreXsRNwII+DvPIEVI9JKIB2ynYkyT"); err != nil {
		panic(err)
	} else {
		MySQLDSN = v
	}
}

func main() {

	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./source/gen/dal",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	}
	g.WithDataTypeMap(dataMap)

	autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		return map[string][]string{"column": {"update_time"}, "type": {"datetime(3)"}, "autoUpdateTime": {}}
	})
	autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		return map[string][]string{"column": {"create_time"}, "type": {"datetime(3)"}, "autoCreateTime": {}}
	})
	softDeleteField := gen.FieldType("delete_time", "gorm.DeletedAt")
	fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteField}

	// 这里创建个别模型仅仅是为了拿到`*generate.QueryStructMeta`类型对象用于后面的模型关联操作中
	UpTag := g.GenerateModel("archived_ups_tag")
	FavFolder := g.GenerateModel("archived_fav_folders")
	Video := g.GenerateModel("archived_video")

	allModel := g.GenerateAllTable(fieldOpts...)

	// 创建有关联关系的模型文件
	Upper := g.GenerateModel("archived_ups",
		append(
			fieldOpts,
			gen.FieldRelate(field.HasOne, "UpTag", UpTag,
				&field.RelateConfig{GORMTag: map[string][]string{"foreignKey": {"tag_id"}}}),
		)...,
	)

	// 创建有关联关系的模型文件
	Fav := g.GenerateModel("archived_fav",
		append(
			fieldOpts,
			gen.FieldRelate(field.HasOne, "FavFolder", FavFolder,
				&field.RelateConfig{GORMTag: map[string][]string{"foreignKey": {"fid"}}}),
			gen.FieldRelate(field.HasOne, "VideoInfo", Video,
				&field.RelateConfig{GORMTag: map[string][]string{"foreignKey": {"bvid"}}}),
		)...,
	)

	Coin := g.GenerateModel("archived_coin",
		append(
			fieldOpts,
			gen.FieldRelate(field.HasOne, "VideoInfo", Video,
				&field.RelateConfig{GORMTag: map[string][]string{"foreignKey": {"bvid"}}}),
		)...,
	)
	Like := g.GenerateModel("archived_like",
		append(
			fieldOpts,
			gen.FieldRelate(field.HasOne, "VideoInfo", Video,
				&field.RelateConfig{GORMTag: map[string][]string{"foreignKey": {"bvid"}}}),
		)...,
	)
	History := g.GenerateModel("archived_view_history",
		append(
			fieldOpts,
			gen.FieldRelate(field.HasOne, "VideoInfo", Video,
				&field.RelateConfig{GORMTag: map[string][]string{"foreignKey": {"bvid"}}}),
		)...,
	)
	// 创建模型的方法,生成文件在 query 目录; 先创建结果不会被后创建的覆盖
	g.ApplyBasic(Video, Upper, UpTag, Fav, Like, Coin, History)
	g.ApplyBasic(allModel...)
	// g.ApplyInterface(func(UserInterface) {}, g.GenerateModel("user"))

	g.Execute()
}
