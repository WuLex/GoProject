package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // SQL Server 驱动程序
)

type Employee struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// 连接字符串
	//connString := "server=.;user id=sa;password=*****;database=ChipDb"
    connString := "Data Source=.;Initial Catalog=ChipDb;Integrated Security=True"
	// 连接到 SQL Server
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("数据库连接失败：", err.Error())
	}
	defer db.Close()

	// 检查数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接失败：", err.Error())
	}

	fmt.Println("成功连接到数据库")

	// 检查表是否存在
	if !isTableExists(db, "Employees") {
		// 表不存在，创建表
		createTable(db)
	} else {
		fmt.Println("表已存在，无需创建")
	}

	// 插入数据
	employee := Employee{ID: 1, Name: "John", Age: 30}
	insertEmployee(db, employee)

	// 查询数据
	employees, err := getEmployees(db)
	if err != nil {
		log.Fatal("查询数据失败：", err.Error())
	}

	// 打印查询结果
	for _, emp := range employees {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", emp.ID, emp.Name, emp.Age)
	}

	// 更新数据
	updatedEmployee := Employee{ID: 1, Name: "John Doe", Age: 35}
	updateEmployee(db, updatedEmployee)

	// 删除数据
	deleteEmployee(db, 1)
}

// 检查表是否存在
func isTableExists(db *sql.DB, tableName string) bool {
	query := `
		SELECT COUNT(*) 
		FROM INFORMATION_SCHEMA.TABLES 
		WHERE TABLE_NAME = @TableName
	`

	var count int
	err := db.QueryRow(query, sql.Named("TableName", tableName)).Scan(&count)
	if err != nil {
		log.Fatal("检查表存在性失败：", err.Error())
	}

	return count > 0
}

// 创建表
func createTable(db *sql.DB) {
	query := `
		CREATE TABLE Employees (
			ID INT PRIMARY KEY,
			Name VARCHAR(50),
			Age INT
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("创建表失败：", err.Error())
	}

	fmt.Println("表创建成功")
}

// 插入数据
func insertEmployee(db *sql.DB, employee Employee) {
	query := `
		INSERT INTO Employees (ID, Name, Age)
		VALUES (@ID, @Name, @Age)
	`

	_, err := db.Exec(query, sql.Named("ID", employee.ID), sql.Named("Name", employee.Name), sql.Named("Age", employee.Age))
	if err != nil {
		log.Fatal("插入数据失败：", err.Error())
	}

	fmt.Println("数据插入成功")
}

// 查询数据
func getEmployees(db *sql.DB) ([]Employee, error) {
	query := "SELECT ID, Name, Age FROM Employees"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Age)
		if err != nil {
		return nil, err
	    } 
	employees = append(employees, emp)
} 
return employees, nil
}

// 更新数据
func updateEmployee(db *sql.DB, employee Employee) {
query := "UPDATE Employees SET Name = @Name, Age = @Age WHERE ID = @ID"

_, err := db.Exec(query, sql.Named("ID", employee.ID), sql.Named("Name", employee.Name), sql.Named("Age", employee.Age))
if err != nil {
	log.Fatal("更新数据失败：", err.Error())
}

fmt.Println("数据更新成功")
}

// 删除数据
func deleteEmployee(db *sql.DB, employeeID int) {
query := "DELETE FROM Employees WHERE ID = @ID"

_, err := db.Exec(query, sql.Named("ID", employeeID))
if err != nil {
	log.Fatal("删除数据失败：", err.Error())
}

fmt.Println("数据删除成功")
}