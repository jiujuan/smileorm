package smileorm

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Builder struct {
	conn    *Connection
	table   string
	fields  []string
	where   [][]interface{}
	join    [][]interface{}
	order   string
	group   string
	sql     []string
	infoSql string
}

func NewBuilder(conn *Connection, table string) *Builder {
	return &Builder{conn: conn, table: table}
}

func (bld *Builder) Table(table string) *Builder {
	bld.table = table
	return bld
}

func (bld *Builder) Fields(fields ...string) *Builder {
	bld.fields = fields
	return bld
}

func (bld *Builder) Where(args ...interface{}) *Builder {
	wh := []interface{}{" 1=1 ", args}
	bld.where = append(bld.where, wh)
	return bld
}

func (bld *Builder) WhereIn(field string, arr []interface{}) *Builder {
	return bld.Where(field, "IN", arr)
}

func (bld *Builder) GroupBy(group string) *Builder {
	bld.group = group
	return bld
}

func (bld *Builder) Order(order string) *Builder {
	bld.order = order
	return bld
}

func (bld *Builder) Join(args ...interface{}) *Builder {
	bld.join = append(bld.join, []interface{}{"INNER", args})
	return bld
}

func (bld *Builder) LeftJoin(args ...interface{}) *Builder {
	bld.join = append(bld.join, []interface{}{"LEFT", args})
	return bld
}

func (bld *Builder) RightJoin(args ...interface{}) *Builder {
	bld.join = append(bld.join, []interface{}{"RIGHT", args})
	return bld
}

func (bld *Builder) DebugSql() string {
	return bld.infoSql
}

// parse where conditions
//where("id", 2)
// where("id", ">", 2)
// where({{"id", ">", 2}, {"name", "=", "tom"}})
func (bld *Builder) parseWhere(sep string) (string, error) {
	wheres := bld.where
	var sqlSlice []string
	var sql string

	val := reflect.Indirect(reflect.ValueOf(wheres[0]))
	switch val.Kind() {
	case reflect.String:
		if len(wheres) == 3 {
			sql = fmt.Sprintf(" %s%s%v ", wheres[0], wheres[1], wheres[2])
		} else if len(wheres) == 2 {
			sql = fmt.Sprintf(" %s%v ", wheres[0], wheres[1])
		}
		sqlSlice = append(sqlSlice, sql)
	case reflect.Slice:
		for _, where := range wheres {
			sql = fmt.Sprintf(" %s%s%v ", where[0], where[1], where[2])
			sqlSlice = append(sqlSlice, sql)
		}
	default:
		sqlSlice = []string{}
	}

	return strings.Join(sqlSlice, sep), nil
}

func (bld *Builder) AndWhere() (string, error) {
	return bld.parseWhere(" AND ")
}

func (bld *Builder) OrWhere() (string, error) {
	return bld.parseWhere(" OR ")
}

// Insert data
// data{id: xx, name: xxx}
func (bld *Builder) Insert(data interface{}) (int64, error) {
	fieldSlice, valueSlice := reflectStruct(data)

	var sep strings.Builder
	if len(fieldSlice) > 0 && len(valueSlice) > 0 {
		for i := 0; i < len(fieldSlice); i++ {
			sep.WriteString("?,")
		}

		insertSQL := fmt.Sprintf(getInsertStr(), bld.table, strings.Join(fieldSlice, ", "),
			strings.TrimRight(sep.String(), ","))

		bld.infoSql = insertSQL
		result, err := bld.conn.InsertRaw(insertSQL, valueSlice...)
		return result, err
	}
	bld.infoSql = ""
	return 0, errors.New("insert error")
}

func (bld *Builder) Update(data interface{}) (int64, error) {

}

func getInsertStr() string {
	return "INSERT INTO %s (%s) VALUES (%s)"
}

func getUpdateStr() string {
	return "UPDATE %s SET %s %s"
}

func getDeleteStr() string {
	return "DELETE %s %s"
}

func reflectStruct(data interface{}) (columns []string, values []interface{}) {
	valOf := reflect.ValueOf(data)
	tpeOf := reflect.TypeOf(data)
	val := reflect.Indirect(valOf)

	fieldSlice := make([]string, 0)
	valueSlice := make([]interface{}, 0)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			//name tpeOf.Field(i).Name
			//value valOf.Field(i).Interface()
			fieldSlice = append(fieldSlice, tpeOf.Field(i).Name)
			valueSlice = append(valueSlice, valOf.Field(i).Interface())
		}
	}
	return fieldSlice, valueSlice
}
