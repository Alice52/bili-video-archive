// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/alice52/archive/bilibili/source/gen/model"
)

func newArchivedViewHistory(db *gorm.DB, opts ...gen.DOOption) archivedViewHistory {
	_archivedViewHistory := archivedViewHistory{}

	_archivedViewHistory.archivedViewHistoryDo.UseDB(db, opts...)
	_archivedViewHistory.archivedViewHistoryDo.UseModel(&model.ArchivedViewHistory{})

	tableName := _archivedViewHistory.archivedViewHistoryDo.TableName()
	_archivedViewHistory.ALL = field.NewAsterisk(tableName)
	_archivedViewHistory.ID = field.NewInt64(tableName, "id")
	_archivedViewHistory.CreateTime = field.NewInt64(tableName, "create_time")
	_archivedViewHistory.UpdateTime = field.NewInt64(tableName, "update_time")
	_archivedViewHistory.DeleteTime = field.NewField(tableName, "delete_time")

	_archivedViewHistory.fillFieldMap()

	return _archivedViewHistory
}

// archivedViewHistory 浏览历史记录
type archivedViewHistory struct {
	archivedViewHistoryDo

	ALL        field.Asterisk
	ID         field.Int64
	CreateTime field.Int64
	UpdateTime field.Int64
	DeleteTime field.Field

	fieldMap map[string]field.Expr
}

func (a archivedViewHistory) Table(newTableName string) *archivedViewHistory {
	a.archivedViewHistoryDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a archivedViewHistory) As(alias string) *archivedViewHistory {
	a.archivedViewHistoryDo.DO = *(a.archivedViewHistoryDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *archivedViewHistory) updateTableName(table string) *archivedViewHistory {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CreateTime = field.NewInt64(table, "create_time")
	a.UpdateTime = field.NewInt64(table, "update_time")
	a.DeleteTime = field.NewField(table, "delete_time")

	a.fillFieldMap()

	return a
}

func (a *archivedViewHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *archivedViewHistory) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["id"] = a.ID
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["delete_time"] = a.DeleteTime
}

func (a archivedViewHistory) clone(db *gorm.DB) archivedViewHistory {
	a.archivedViewHistoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a archivedViewHistory) replaceDB(db *gorm.DB) archivedViewHistory {
	a.archivedViewHistoryDo.ReplaceDB(db)
	return a
}

type archivedViewHistoryDo struct{ gen.DO }

type IArchivedViewHistoryDo interface {
	gen.SubQuery
	Debug() IArchivedViewHistoryDo
	WithContext(ctx context.Context) IArchivedViewHistoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArchivedViewHistoryDo
	WriteDB() IArchivedViewHistoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArchivedViewHistoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArchivedViewHistoryDo
	Not(conds ...gen.Condition) IArchivedViewHistoryDo
	Or(conds ...gen.Condition) IArchivedViewHistoryDo
	Select(conds ...field.Expr) IArchivedViewHistoryDo
	Where(conds ...gen.Condition) IArchivedViewHistoryDo
	Order(conds ...field.Expr) IArchivedViewHistoryDo
	Distinct(cols ...field.Expr) IArchivedViewHistoryDo
	Omit(cols ...field.Expr) IArchivedViewHistoryDo
	Join(table schema.Tabler, on ...field.Expr) IArchivedViewHistoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArchivedViewHistoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArchivedViewHistoryDo
	Group(cols ...field.Expr) IArchivedViewHistoryDo
	Having(conds ...gen.Condition) IArchivedViewHistoryDo
	Limit(limit int) IArchivedViewHistoryDo
	Offset(offset int) IArchivedViewHistoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArchivedViewHistoryDo
	Unscoped() IArchivedViewHistoryDo
	Create(values ...*model.ArchivedViewHistory) error
	CreateInBatches(values []*model.ArchivedViewHistory, batchSize int) error
	Save(values ...*model.ArchivedViewHistory) error
	First() (*model.ArchivedViewHistory, error)
	Take() (*model.ArchivedViewHistory, error)
	Last() (*model.ArchivedViewHistory, error)
	Find() ([]*model.ArchivedViewHistory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArchivedViewHistory, err error)
	FindInBatches(result *[]*model.ArchivedViewHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ArchivedViewHistory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArchivedViewHistoryDo
	Assign(attrs ...field.AssignExpr) IArchivedViewHistoryDo
	Joins(fields ...field.RelationField) IArchivedViewHistoryDo
	Preload(fields ...field.RelationField) IArchivedViewHistoryDo
	FirstOrInit() (*model.ArchivedViewHistory, error)
	FirstOrCreate() (*model.ArchivedViewHistory, error)
	FindByPage(offset int, limit int) (result []*model.ArchivedViewHistory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArchivedViewHistoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a archivedViewHistoryDo) Debug() IArchivedViewHistoryDo {
	return a.withDO(a.DO.Debug())
}

func (a archivedViewHistoryDo) WithContext(ctx context.Context) IArchivedViewHistoryDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a archivedViewHistoryDo) ReadDB() IArchivedViewHistoryDo {
	return a.Clauses(dbresolver.Read)
}

func (a archivedViewHistoryDo) WriteDB() IArchivedViewHistoryDo {
	return a.Clauses(dbresolver.Write)
}

func (a archivedViewHistoryDo) Session(config *gorm.Session) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Session(config))
}

func (a archivedViewHistoryDo) Clauses(conds ...clause.Expression) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a archivedViewHistoryDo) Returning(value interface{}, columns ...string) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a archivedViewHistoryDo) Not(conds ...gen.Condition) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a archivedViewHistoryDo) Or(conds ...gen.Condition) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a archivedViewHistoryDo) Select(conds ...field.Expr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a archivedViewHistoryDo) Where(conds ...gen.Condition) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a archivedViewHistoryDo) Order(conds ...field.Expr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a archivedViewHistoryDo) Distinct(cols ...field.Expr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a archivedViewHistoryDo) Omit(cols ...field.Expr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a archivedViewHistoryDo) Join(table schema.Tabler, on ...field.Expr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a archivedViewHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a archivedViewHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a archivedViewHistoryDo) Group(cols ...field.Expr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a archivedViewHistoryDo) Having(conds ...gen.Condition) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a archivedViewHistoryDo) Limit(limit int) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a archivedViewHistoryDo) Offset(offset int) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a archivedViewHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a archivedViewHistoryDo) Unscoped() IArchivedViewHistoryDo {
	return a.withDO(a.DO.Unscoped())
}

func (a archivedViewHistoryDo) Create(values ...*model.ArchivedViewHistory) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a archivedViewHistoryDo) CreateInBatches(values []*model.ArchivedViewHistory, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a archivedViewHistoryDo) Save(values ...*model.ArchivedViewHistory) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a archivedViewHistoryDo) First() (*model.ArchivedViewHistory, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedViewHistory), nil
	}
}

func (a archivedViewHistoryDo) Take() (*model.ArchivedViewHistory, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedViewHistory), nil
	}
}

func (a archivedViewHistoryDo) Last() (*model.ArchivedViewHistory, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedViewHistory), nil
	}
}

func (a archivedViewHistoryDo) Find() ([]*model.ArchivedViewHistory, error) {
	result, err := a.DO.Find()
	return result.([]*model.ArchivedViewHistory), err
}

func (a archivedViewHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ArchivedViewHistory, err error) {
	buf := make([]*model.ArchivedViewHistory, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a archivedViewHistoryDo) FindInBatches(result *[]*model.ArchivedViewHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a archivedViewHistoryDo) Attrs(attrs ...field.AssignExpr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a archivedViewHistoryDo) Assign(attrs ...field.AssignExpr) IArchivedViewHistoryDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a archivedViewHistoryDo) Joins(fields ...field.RelationField) IArchivedViewHistoryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a archivedViewHistoryDo) Preload(fields ...field.RelationField) IArchivedViewHistoryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a archivedViewHistoryDo) FirstOrInit() (*model.ArchivedViewHistory, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedViewHistory), nil
	}
}

func (a archivedViewHistoryDo) FirstOrCreate() (*model.ArchivedViewHistory, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ArchivedViewHistory), nil
	}
}

func (a archivedViewHistoryDo) FindByPage(offset int, limit int) (result []*model.ArchivedViewHistory, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a archivedViewHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a archivedViewHistoryDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a archivedViewHistoryDo) Delete(models ...*model.ArchivedViewHistory) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *archivedViewHistoryDo) withDO(do gen.Dao) *archivedViewHistoryDo {
	a.DO = *do.(*gen.DO)
	return a
}
