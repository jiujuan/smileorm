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
	bld.where = append(bld.where, args)
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

func (bld *Builder) AndWhere() (string, error) {
	return bld.parseWhere(" AND ")
}

func (bld *Builder) OrWhere() (string, error) {
	return bld.parseWhere(" OR ")
}

// Insert data
// data{id: xx, name: xxx}
func (bld *Builder) Insert(data interface{}) (int64, error) {
	fieldSlice, valueSlice := mapToSlice(parseParamData(data))

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

// Update data
// conn.Table(table_name).Where({id: xxx}).Update({name: xxx, age:xxx})
func (bld *Builder) Update(data interface{}) (int64, error) {
	stringWhere, _ := bld.AndWhere()

	mapdata := parseParamData(data)
	if len(mapdata) > 0 {
		var setstr strings.Builder
		valSlice := make([]interface{}, 0)
		for key, val := range mapdata {
			setstr.WriteString(fmt.Sprintf("%s=?,", key))
			valSlice = append(valSlice, val)
		}
		updateSQL := fmt.Sprintf(getUpdateStr(), bld.table, strings.TrimRight(setstr.String(), ","),
			" WHERE "+stringWhere)
		bld.infoSql = updateSQL

		result, err := bld.conn.UpdateRaw(updateSQL, valSlice...)
		return result, err
	}
	return 0, errors.New("update error")
}

func (bld *Builder) Delete() (int64, error) {
	stringWhere, _ := bld.AndWhere()

	deleteSQL := fmt.Sprintf(getDeleteStr(), bld.table, " WHERE "+stringWhere)
	bld.infoSql = deleteSQL

	result, err := bld.conn.DeleteRaw(deleteSQL)
	return result, err
}

func getOpSqlStr(optype string) string {
	var sql string
	switch optype {
	case "insert":
		sql = getInsertStr()
	case "update":
		sql = getUpdateStr()
	case "delete":
		sql = getDeleteStr()
	}
	return sql
}

func getInsertStr() string {
	return "INSERT INTO %s (%s) VALUES (%s)"
}

func getUpdateStr() string {
	return "UPDATE %s SET %s %s"
}

func getDeleteStr() string {
	//delete table where id=xx
	return "DELETE %s %s"
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
			sql = fmt.Sprintf(" %s=%v ", wheres[0], wheres[1])
		}
		sqlSlice = append(sqlSlice, sql)
	case reflect.Slice:
		for _, where := range wheres {
			if len(where) == 2 {
				sql = fmt.Sprintf(" %s=%v ", where[0], where[1])
			} else if len(where) == 3 {
				sql = fmt.Sprintf(" %s%s%v ", where[0], where[1], where[2])
			}
			sqlSlice = append(sqlSlice, sql)
		}
	default:
		sqlSlice = []string{}
	}

	return strings.Join(sqlSlice, sep), nil
}

func (bld *Builder) parseWhereslice()
func parseParamData(data interface{}) map[string]interface{} {
	valOf := reflect.ValueOf(data)
	tpeOf := reflect.TypeOf(data)
	val := reflect.Indirect(valOf)

	mapString := make(map[string]interface{})

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			//name tpeOf.Field(i).Name
			//value valOf.Field(i).Interface()
			// fieldSlice = append(fieldSlice, tpeOf.Field(i).Name)
			// valueSlice = append(valueSlice, valOf.Field(i).Interface())
			mapString[tpeOf.Field(i).Name] = valOf.Field(i).Interface()
		}
	case reflect.Slice:
		num := val.Len()
		for i := 0; i < num; i++ {
			item := val.Index(i)

		}
	}
	return mapString
}

func mapToSlice(data map[string]interface{}) (fields []string, values []interface{}) {
	fieldSlice := make([]string, 0)
	valueSlice := make([]interface{}, 0)
	if len(data) > 0 {
		for k, v := range data {
			fieldSlice = append(fieldSlice, k)
			valueSlice = append(valueSlice, v)
		}
	}
	return fieldSlice, valueSlice
}
