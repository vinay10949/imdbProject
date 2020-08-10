package infrastructures

import (
	"database/sql"
	"fmt"
	"imdbProject/interfaces"
)

type MySQLHandler struct {
	Conn *sql.DB
}

func (handler *MySQLHandler) Execute(statement string,args ...interface{})error {
	stmt,_:=handler.Conn.Prepare(statement)
	_,err:=stmt.Exec(args)
	return  err
}

func (handler *MySQLHandler) Query(statement string) (interfaces.IRow, error) {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)

	if err != nil {
		fmt.Println(err)
		return new(MySQLRow),err
	}
	row := new(MySQLRow)
	row.Rows = rows

	return row, nil
}

type MySQLRow struct {
	Rows *sql.Rows
}

func (r MySQLRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return  nil
}

func (r MySQLRow) Next() bool {
	return r.Rows.Next()
}