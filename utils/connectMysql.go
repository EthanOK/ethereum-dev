package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetMysqlDB() *sql.DB {
	// 数据库连接信息
	db, err := sql.Open("mysql", "root:root@tcp(192.168.0.173:3306)/aggregator_ethan")

	// 检查是否连接成功
	if err != nil {
		log.Fatal(err)
	}

	// 确保在程序退出时关闭数据库连接
	// defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Query(db *sql.DB, sqlStatement string) {

	// 执行查询操作
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个存储数据的切片，每个元素都是 sql.RawBytes 类型
	values := make([]sql.RawBytes, len(columns))

	// 为每个列的数据创建一个指针
	pointers := make([]interface{}, len(columns))
	for i := range values {
		pointers[i] = &values[i]
	}

	// 遍历查询结果
	for rows.Next() {
		if err := rows.Scan(pointers...); err != nil {
			log.Fatal(err)
		}

		// 打印每一行的数据
		for i, col := range values {
			if col == nil {
				fmt.Printf("%s: NULL ", columns[i])
			} else {
				fmt.Printf("%s: %s ", columns[i], col)
			}
		}
		fmt.Println()
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// 返回自增长的Id
func Insert(db *sql.DB, sqlStatement string, args ...interface{}) int64 {
	// 准备插入语句
	insertStmt, err := db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()

	// 插入数据
	result, err := insertStmt.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}

	// 获取最新插入的自增ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return lastInsertID
}

// 只插入不返回数据
func InsertOnly(db *sql.DB, sqlStatement string, args ...interface{}) {
	// 准备插入语句
	insertStmt, err := db.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()

	// 插入数据
	result, err := insertStmt.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}
	result.RowsAffected()
}
