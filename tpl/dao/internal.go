// ==========================================================================
// Code generated by Tekin. DO NOT EDIT. {TplDatetimeStr}
// @author TekinTian <tekintian@gmail.com>
// ==========================================================================

package internal

import (
	"context"
	"database/sql"
	"{TplImportPrefix}/app/system/model/entity"
	"time"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// {TplTableNameCamelCase}Dao is the data access object for table {TplTableName}.
type {TplTableNameCamelCase}Dao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns {TplTableNameCamelCase}Columns // columns contains all the column names of Table for convenient usage.
}

// {TplTableNameCamelCase}Columns defines and stores column names for table {TplTableName}.
type {TplTableNameCamelCase}Columns struct {
	{TplColumnDefine}
}

//  {TplTableNameCamelLowerCase}Columns holds the columns for table {TplTableName}.
var {TplTableNameCamelLowerCase}Columns = {TplTableNameCamelCase}Columns{
	{TplColumnNames}
}

// New{TplTableNameCamelCase}Dao creates and returns a new DAO object for table data access.
func New{TplTableNameCamelCase}Dao() *{TplTableNameCamelCase}Dao {
	return &{TplTableNameCamelCase}Dao{
		group:   "{TplGroupName}",
		table:   "{TplTableName}",
		columns: {TplTableNameCamelLowerCase}Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *{TplTableNameCamelCase}Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *{TplTableNameCamelCase}Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *{TplTableNameCamelCase}Dao) Columns() {TplTableNameCamelCase}Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *{TplTableNameCamelCase}Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *{TplTableNameCamelCase}Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *{TplTableNameCamelCase}Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}


// TX sets the transaction for current operation.
func (dao *{TplTableNameCamelCase}Dao) TX(tx *gdb.TX) *gdb.Model {
	var model *entity.{TplTableNameCamelCase}
	return tx.Model(model)
}

// As sets an alias name for current table.
func (dao *{TplTableNameCamelCase}Dao) As(ctx context.Context,as string) *gdb.Model {
	return dao.Ctx(ctx).As(as)
}

// Args sets custom arguments for model operation.
func (dao *{TplTableNameCamelCase}Dao) Args(ctx context.Context,args ...interface{}) *gdb.Model {

    return dao.Ctx(ctx).Args(args ...)
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the entity.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (dao *{TplTableNameCamelCase}Dao) LeftJoin(ctx context.Context,table ...string) *gdb.Model {
	return dao.Ctx(ctx).LeftJoin(table...)
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the entity.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (dao *{TplTableNameCamelCase}Dao) RightJoin(ctx context.Context,table ...string) *gdb.Model {
	return dao.Ctx(ctx).RightJoin(table...)
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the entity.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (dao *{TplTableNameCamelCase}Dao) InnerJoin(ctx context.Context,table ...string) *gdb.Model {
	return dao.Ctx(ctx).InnerJoin(table...)
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (dao *{TplTableNameCamelCase}Dao) Fields(ctx context.Context,fieldNamesOrMapStruct ...interface{}) *gdb.Model {
	return dao.Ctx(ctx).Fields(fieldNamesOrMapStruct...)
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (dao *{TplTableNameCamelCase}Dao) FieldsEx(ctx context.Context,fieldNamesOrMapStruct ...interface{}) *gdb.Model {
	return dao.Ctx(ctx).FieldsEx(fieldNamesOrMapStruct...)
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (dao *{TplTableNameCamelCase}Dao) OmitEmpty(ctx context.Context) *gdb.Model {
	return dao.Ctx(ctx).OmitEmpty()
}


// Where sets the condition statement for the entity. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (dao *{TplTableNameCamelCase}Dao) Where(ctx context.Context,where interface{}, args ...interface{}) *gdb.Model {
	return dao.Ctx(ctx).Where(where, args...)
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (dao *{TplTableNameCamelCase}Dao) WherePri(ctx context.Context,where interface{}, args ...interface{}) *gdb.Model {
	return dao.Ctx(ctx).WherePri(where, args...)
}


// Or adds "OR" condition to the where statement.
func (dao *{TplTableNameCamelCase}Dao) Or(ctx context.Context,where interface{}, args ...interface{}) *gdb.Model {
	return dao.Ctx(ctx).WhereOr(where, args...)
}

// Group sets the "GROUP BY" statement for the entity.
func (dao *{TplTableNameCamelCase}Dao) GroupBy(ctx context.Context,groupBy string) *gdb.Model {
	return dao.Ctx(ctx).Group(groupBy)
}

// Order sets the "ORDER BY" statement for the entity.
func (dao *{TplTableNameCamelCase}Dao) Order(ctx context.Context,orderBy ...string) *gdb.Model {
	return dao.Ctx(ctx).Order(orderBy)
}

// Limit sets the "LIMIT" statement for the entity.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (dao *{TplTableNameCamelCase}Dao) Limit(ctx context.Context,limit ...int) *gdb.Model {
	return dao.Ctx(ctx).Limit(limit...)
}

// Offset sets the "OFFSET" statement for the entity.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (dao *{TplTableNameCamelCase}Dao) Offset(ctx context.Context,offset int) *gdb.Model {
	return dao.Ctx(ctx).Offset(offset)
}

// Page sets the paging number for the entity.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (dao *{TplTableNameCamelCase}Dao) Page(ctx context.Context,page, limit int) *gdb.Model {
	return dao.Ctx(ctx).Page(page, limit)
}

// Batch sets the batch operation number for the entity.
func (dao *{TplTableNameCamelCase}Dao) Batch(ctx context.Context,batch int) *gdb.Model {
	return dao.Ctx(ctx).Batch(batch)
}

// Cache sets the cache feature for the entity. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (dao *{TplTableNameCamelCase}Dao) Cache(ctx context.Context,duration time.Duration, name ...string) *gdb.Model {
	var option gdb.CacheOption
	if len(name)>0 {
	   option = gdb.CacheOption{ Name: gconv.String(name),Duration:  duration}
	}else{
		option = gdb.CacheOption{ Duration:  duration}
	}
    return dao.Ctx(ctx).Cache(option)
}
// Cache add Force caches the query result whatever the result is nil or not.
// It is used to avoid Cache Penetration.
func (dao *{TplTableNameCamelCase}Dao) CacheForce(ctx context.Context,duration time.Duration, isForce bool, name ...string) *gdb.Model {
	var option gdb.CacheOption
	if len(name)>0 {
	   option = gdb.CacheOption{ Name: gconv.String(name),Duration:  duration, Force: isForce}
	}else{
		option = gdb.CacheOption{ Duration:  duration, Force: isForce}
	}
    return dao.Ctx(ctx).Cache(option)
}
// Data sets the operation data for the entity.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (dao *{TplTableNameCamelCase}Dao) Data(ctx context.Context,data ...interface{}) *gdb.Model {
	return dao.Ctx(ctx).Data(data...)
}

// All does "SELECT FROM ..." statement for the entity.
// It retrieves the records from table and returns the result as []*entity.{TplTableNameCamelCase}.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (dao *{TplTableNameCamelCase}Dao) All(ctx context.Context,where ...interface{}) ([]*entity.{TplTableNameCamelCase}, error) {
	all, err := dao.Ctx(ctx).All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*entity.{TplTableNameCamelCase}
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *entity.{TplTableNameCamelCase}.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (dao *{TplTableNameCamelCase}Dao) One(ctx context.Context,where ...interface{}) (*entity.{TplTableNameCamelCase}, error) {
	one, err := dao.Ctx(ctx).One(where...)
	if err != nil {
		return nil, err
	}
	var entity *entity.{TplTableNameCamelCase}
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (dao *{TplTableNameCamelCase}Dao) FindOne(ctx context.Context,where ...interface{}) (*entity.{TplTableNameCamelCase}, error) {
	one, err := dao.Ctx(ctx).One(where...)
	if err != nil {
		return nil, err
	}
	var entity *entity.{TplTableNameCamelCase}
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (dao *{TplTableNameCamelCase}Dao) FindAll(ctx context.Context,where ...interface{}) ([]*entity.{TplTableNameCamelCase}, error) {
	all, err := dao.Ctx(ctx).All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*entity.{TplTableNameCamelCase}
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter <pointer> should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not nil.
//
// Eg:
// user := new(User)
// err  := dao.User.Where(ctx, "id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where(ctx, "id", 1).Struct(&user)
func (dao *{TplTableNameCamelCase}Dao) Struct(ctx context.Context,pointer interface{}, where ...interface{}) error {
	return dao.Ctx(ctx).Scan(pointer, where...)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter <pointer> should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not empty.
//
// Eg:
// users := ([]User)(nil)
// err   := dao.User.Structs(ctx, &users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(ctx, &users)
func (dao *{TplTableNameCamelCase}Dao) Structs(ctx context.Context,pointer interface{}, where ...interface{}) error {
	return dao.Ctx(ctx).Scan(pointer, where...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter <pointer>.
// It calls function Struct if <pointer> is type of *struct/**struct.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved and given pointer is not empty or nil.
//
// Eg:
// user  := new(User)
// err   := dao.User.Where(ctx, "id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where(ctx, "id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(ctx, &users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(ctx, &users)
func (dao *{TplTableNameCamelCase}Dao) Scan(ctx context.Context,pointer interface{}, where ...interface{}) error {
	return dao.Ctx(ctx).Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (dao *{TplTableNameCamelCase}Dao) Chunk(ctx context.Context,limit int, callback func(entities []*entity.{TplTableNameCamelCase}, err error) bool) {
	dao.Ctx(ctx).Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*entity.{TplTableNameCamelCase}
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (dao *{TplTableNameCamelCase}Dao) LockUpdate(ctx context.Context) *gdb.Model {
	return dao.Ctx(ctx).LockUpdate()
}

// LockShared sets the lock in share mode for current operation.
func (dao *{TplTableNameCamelCase}Dao) LockShared(ctx context.Context) *gdb.Model {
	return dao.Ctx(ctx).LockShared()
}

// Unscoped enables/disables the soft deleting feature.
func (dao *{TplTableNameCamelCase}Dao) Unscoped(ctx context.Context) *gdb.Model {
	return dao.Ctx(ctx).Unscoped()
}

func (dao *{TplTableNameCamelCase}Dao) Update(ctx context.Context, dataAndWhere ...interface{}) (result sql.Result, err error) {
	return dao.Ctx(ctx).Update(dataAndWhere...)
}

func (dao *{TplTableNameCamelCase}Dao) FindCount(ctx context.Context,where ...interface{}) (int, error) {
	return dao.Ctx(ctx).Count(where...)
}

func (dao *{TplTableNameCamelCase}Dao) Delete(ctx context.Context,where ...interface{}) (int64, error) {
	rs,_err := dao.Ctx(ctx).Delete(where ...)
	if _err != nil {
		return -1, _err
	}
	return rs.RowsAffected()
}

// just put by Tekin , for gotms