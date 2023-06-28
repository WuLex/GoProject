package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	// 连接到数据库
	//connString := "server=hostname;user id=username;password=password;database=ChipDb"
	connString := "server=.;user id=sa;password=*******;database=ChipDb"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 执行查询
	rows, err := db.Query("SELECT ManufacturerID, ManufacturerName FROM Manufacturer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 创建Excel文件
	f := excelize.NewFile()
	sheetName := "Sheet1"

	// 写入表头
	headers := []string{"ManufacturerID", "ManufacturerName"}
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 写入数据行
	row := 2 // 从第2行开始写入数据
	for rows.Next() {
		var manufacturerID int
		var manufacturerName string
		err := rows.Scan(&manufacturerID, &manufacturerName)
		if err != nil {
			log.Fatal(err)
		}

		data := []interface{}{manufacturerID, manufacturerName}
		for col, value := range data {
			cell, _ := excelize.CoordinatesToCellName(col+1, row)
			f.SetCellValue(sheetName, cell, value)
		}

		row++
	}

	// 保存Excel文件
	err = f.SaveAs("output.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("数据导出成功！")
}