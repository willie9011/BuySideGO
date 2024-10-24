package connSqlserver

import (
	"database/sql"
	"fmt"
)

func ConnSqlserver() {
	connString := "sever=LOHAHA;user id=sa;password=lo850608;port=1433;database=Test01"
	// 建立連線字串
	db, err := sql.Open("mssql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Employees")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	//輸出結果，確認連線成功
	for rows.Next() {
		var Name string
		var Phone string
		var Salary int
		// ... 其他欄位

		err := rows.Scan(&Name, &Phone, &Salary /* ... */)
		if err != nil {
			panic(err)
		}

		fmt.Println(Name, Phone, Salary) // 輸出結果
	}

}
