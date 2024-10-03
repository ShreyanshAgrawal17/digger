// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models_generated

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/diggerhq/digger/ee/drift/model"
)

func newGithubAppInstallationLink(db *gorm.DB, opts ...gen.DOOption) githubAppInstallationLink {
	_githubAppInstallationLink := githubAppInstallationLink{}

	_githubAppInstallationLink.githubAppInstallationLinkDo.UseDB(db, opts...)
	_githubAppInstallationLink.githubAppInstallationLinkDo.UseModel(&model.GithubAppInstallationLink{})

	tableName := _githubAppInstallationLink.githubAppInstallationLinkDo.TableName()
	_githubAppInstallationLink.ALL = field.NewAsterisk(tableName)
	_githubAppInstallationLink.ID = field.NewString(tableName, "id")
	_githubAppInstallationLink.CreatedAt = field.NewTime(tableName, "created_at")
	_githubAppInstallationLink.UpdatedAt = field.NewTime(tableName, "updated_at")
	_githubAppInstallationLink.DeletedAt = field.NewField(tableName, "deleted_at")
	_githubAppInstallationLink.GithubInstallationID = field.NewString(tableName, "github_installation_id")
	_githubAppInstallationLink.OrganisationID = field.NewString(tableName, "organisation_id")
	_githubAppInstallationLink.Status = field.NewString(tableName, "status")

	_githubAppInstallationLink.fillFieldMap()

	return _githubAppInstallationLink
}

type githubAppInstallationLink struct {
	githubAppInstallationLinkDo

	ALL                  field.Asterisk
	ID                   field.String
	CreatedAt            field.Time
	UpdatedAt            field.Time
	DeletedAt            field.Field
	GithubInstallationID field.String
	OrganisationID       field.String
	Status               field.String

	fieldMap map[string]field.Expr
}

func (g githubAppInstallationLink) Table(newTableName string) *githubAppInstallationLink {
	g.githubAppInstallationLinkDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g githubAppInstallationLink) As(alias string) *githubAppInstallationLink {
	g.githubAppInstallationLinkDo.DO = *(g.githubAppInstallationLinkDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *githubAppInstallationLink) updateTableName(table string) *githubAppInstallationLink {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewString(table, "id")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")
	g.DeletedAt = field.NewField(table, "deleted_at")
	g.GithubInstallationID = field.NewString(table, "github_installation_id")
	g.OrganisationID = field.NewString(table, "organisation_id")
	g.Status = field.NewString(table, "status")

	g.fillFieldMap()

	return g
}

func (g *githubAppInstallationLink) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *githubAppInstallationLink) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["deleted_at"] = g.DeletedAt
	g.fieldMap["github_installation_id"] = g.GithubInstallationID
	g.fieldMap["organisation_id"] = g.OrganisationID
	g.fieldMap["status"] = g.Status
}

func (g githubAppInstallationLink) clone(db *gorm.DB) githubAppInstallationLink {
	g.githubAppInstallationLinkDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g githubAppInstallationLink) replaceDB(db *gorm.DB) githubAppInstallationLink {
	g.githubAppInstallationLinkDo.ReplaceDB(db)
	return g
}

type githubAppInstallationLinkDo struct{ gen.DO }

type IGithubAppInstallationLinkDo interface {
	gen.SubQuery
	Debug() IGithubAppInstallationLinkDo
	WithContext(ctx context.Context) IGithubAppInstallationLinkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGithubAppInstallationLinkDo
	WriteDB() IGithubAppInstallationLinkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGithubAppInstallationLinkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGithubAppInstallationLinkDo
	Not(conds ...gen.Condition) IGithubAppInstallationLinkDo
	Or(conds ...gen.Condition) IGithubAppInstallationLinkDo
	Select(conds ...field.Expr) IGithubAppInstallationLinkDo
	Where(conds ...gen.Condition) IGithubAppInstallationLinkDo
	Order(conds ...field.Expr) IGithubAppInstallationLinkDo
	Distinct(cols ...field.Expr) IGithubAppInstallationLinkDo
	Omit(cols ...field.Expr) IGithubAppInstallationLinkDo
	Join(table schema.Tabler, on ...field.Expr) IGithubAppInstallationLinkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGithubAppInstallationLinkDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGithubAppInstallationLinkDo
	Group(cols ...field.Expr) IGithubAppInstallationLinkDo
	Having(conds ...gen.Condition) IGithubAppInstallationLinkDo
	Limit(limit int) IGithubAppInstallationLinkDo
	Offset(offset int) IGithubAppInstallationLinkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGithubAppInstallationLinkDo
	Unscoped() IGithubAppInstallationLinkDo
	Create(values ...*model.GithubAppInstallationLink) error
	CreateInBatches(values []*model.GithubAppInstallationLink, batchSize int) error
	Save(values ...*model.GithubAppInstallationLink) error
	First() (*model.GithubAppInstallationLink, error)
	Take() (*model.GithubAppInstallationLink, error)
	Last() (*model.GithubAppInstallationLink, error)
	Find() ([]*model.GithubAppInstallationLink, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GithubAppInstallationLink, err error)
	FindInBatches(result *[]*model.GithubAppInstallationLink, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GithubAppInstallationLink) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGithubAppInstallationLinkDo
	Assign(attrs ...field.AssignExpr) IGithubAppInstallationLinkDo
	Joins(fields ...field.RelationField) IGithubAppInstallationLinkDo
	Preload(fields ...field.RelationField) IGithubAppInstallationLinkDo
	FirstOrInit() (*model.GithubAppInstallationLink, error)
	FirstOrCreate() (*model.GithubAppInstallationLink, error)
	FindByPage(offset int, limit int) (result []*model.GithubAppInstallationLink, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGithubAppInstallationLinkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g githubAppInstallationLinkDo) Debug() IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Debug())
}

func (g githubAppInstallationLinkDo) WithContext(ctx context.Context) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g githubAppInstallationLinkDo) ReadDB() IGithubAppInstallationLinkDo {
	return g.Clauses(dbresolver.Read)
}

func (g githubAppInstallationLinkDo) WriteDB() IGithubAppInstallationLinkDo {
	return g.Clauses(dbresolver.Write)
}

func (g githubAppInstallationLinkDo) Session(config *gorm.Session) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Session(config))
}

func (g githubAppInstallationLinkDo) Clauses(conds ...clause.Expression) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g githubAppInstallationLinkDo) Returning(value interface{}, columns ...string) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g githubAppInstallationLinkDo) Not(conds ...gen.Condition) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g githubAppInstallationLinkDo) Or(conds ...gen.Condition) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g githubAppInstallationLinkDo) Select(conds ...field.Expr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g githubAppInstallationLinkDo) Where(conds ...gen.Condition) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g githubAppInstallationLinkDo) Order(conds ...field.Expr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g githubAppInstallationLinkDo) Distinct(cols ...field.Expr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g githubAppInstallationLinkDo) Omit(cols ...field.Expr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g githubAppInstallationLinkDo) Join(table schema.Tabler, on ...field.Expr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g githubAppInstallationLinkDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g githubAppInstallationLinkDo) RightJoin(table schema.Tabler, on ...field.Expr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g githubAppInstallationLinkDo) Group(cols ...field.Expr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g githubAppInstallationLinkDo) Having(conds ...gen.Condition) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g githubAppInstallationLinkDo) Limit(limit int) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g githubAppInstallationLinkDo) Offset(offset int) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g githubAppInstallationLinkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g githubAppInstallationLinkDo) Unscoped() IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Unscoped())
}

func (g githubAppInstallationLinkDo) Create(values ...*model.GithubAppInstallationLink) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g githubAppInstallationLinkDo) CreateInBatches(values []*model.GithubAppInstallationLink, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g githubAppInstallationLinkDo) Save(values ...*model.GithubAppInstallationLink) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g githubAppInstallationLinkDo) First() (*model.GithubAppInstallationLink, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubAppInstallationLink), nil
	}
}

func (g githubAppInstallationLinkDo) Take() (*model.GithubAppInstallationLink, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubAppInstallationLink), nil
	}
}

func (g githubAppInstallationLinkDo) Last() (*model.GithubAppInstallationLink, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubAppInstallationLink), nil
	}
}

func (g githubAppInstallationLinkDo) Find() ([]*model.GithubAppInstallationLink, error) {
	result, err := g.DO.Find()
	return result.([]*model.GithubAppInstallationLink), err
}

func (g githubAppInstallationLinkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GithubAppInstallationLink, err error) {
	buf := make([]*model.GithubAppInstallationLink, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g githubAppInstallationLinkDo) FindInBatches(result *[]*model.GithubAppInstallationLink, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g githubAppInstallationLinkDo) Attrs(attrs ...field.AssignExpr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g githubAppInstallationLinkDo) Assign(attrs ...field.AssignExpr) IGithubAppInstallationLinkDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g githubAppInstallationLinkDo) Joins(fields ...field.RelationField) IGithubAppInstallationLinkDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g githubAppInstallationLinkDo) Preload(fields ...field.RelationField) IGithubAppInstallationLinkDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g githubAppInstallationLinkDo) FirstOrInit() (*model.GithubAppInstallationLink, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubAppInstallationLink), nil
	}
}

func (g githubAppInstallationLinkDo) FirstOrCreate() (*model.GithubAppInstallationLink, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GithubAppInstallationLink), nil
	}
}

func (g githubAppInstallationLinkDo) FindByPage(offset int, limit int) (result []*model.GithubAppInstallationLink, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g githubAppInstallationLinkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g githubAppInstallationLinkDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g githubAppInstallationLinkDo) Delete(models ...*model.GithubAppInstallationLink) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *githubAppInstallationLinkDo) withDO(do gen.Dao) *githubAppInstallationLinkDo {
	g.DO = *do.(*gen.DO)
	return g
}