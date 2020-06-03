package smileorm

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	DB *sql.DB
	tx *sql.Tx
}

type Rows struct {
	Rows *sql.Rows
	Err  error
}

func NewDB() *sql.DB {
	config, err := GetConfig()
	if err != nil {
		panic("Get Config error: " + err.Error())
	}
	username := config.DB.Username
	password := config.DB.Password
	protocol := config.DB.Protocol
	port := config.DB.Port
	dbname := config.DB.DB
	host := config.DB.Host
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8", username, password, protocol, host, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("conect db error : " + err.Error())
	}
	db.SetMaxOpenConns(config.DB.MaxOpenConns)
	db.SetMaxIdleConns(config.DB.MaxIdleConns)
	return db
}

func NewConnection() *Connection {
	return &Connection{DB: NewDB()}
}

func (conn *Connection) Close() error {
	return conn.DB.Close()
}

// SelectRaw, Raw select
func (conn *Connection) SelectRaw(query string, args ...interface{}) *Rows {
	rows, err := conn.queryRows(query, args...)
	if err != nil {
		return &Rows{Rows: nil, Err: err}
	}

	return &Rows{Rows: rows, Err: nil}
}

func (conn *Connection) InsertRaw(query string, args ...interface{}) (int64, error) {
	rows, err := conn.exec(query, args...)
	if err != nil {
		return 0, err
	}
	return rows.LastInsertId()
}

func (conn *Connection) UpdateRaw(query string, args ...interface{}) (int64, error) {
	return conn.OpRaw(query, args...)
}

func (conn *Connection) DeleteRaw(query string, args ...interface{}) (int64, error) {
	return conn.OpRaw(query, args...)
}

func (conn *Connection) OpRaw(query string, args ...interface{}) (int64, error) {
	rows, err := conn.exec(query, args...)
	if err != nil {
		return 0, err
	}
	return rows.RowsAffected()
}

func (conn *Connection) queryRows(query string, args ...interface{}) (rs *sql.Rows, err error) {
	rows, err := conn.DB.Query(query, args...)

	return rows, err
}

func (conn *Connection) exec(query string, args ...interface{}) (result sql.Result, err error) {
	return conn.DB.Exec(query, args...)
}

func (conn *Connection) Table(table string) *Builder {
	return NewBuilder(conn, table)
}
