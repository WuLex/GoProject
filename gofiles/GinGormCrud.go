package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// Supplier 是供应商模型
type Supplier struct {
	SupplierID  int    `gorm:"primary_key"`
	CompanyName string `gorm:"size:100"`
	ContactName string `gorm:"size:50"`
	Address     string `gorm:"size:200"`
	City        string `gorm:"size:50"`
	State       string `gorm:"size:50"`
	CountryID   int
	PostalCode  string `gorm:"size:20"`
	Phone       string `gorm:"size:20"`
	Email       string `gorm:"size:100"`
}

var (
	db  *gorm.DB
	err error
)

func main() {
	// 连接数据库
	db, err = gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=your_database_name")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建数据表
	db.AutoMigrate(&Supplier{})

	// 初始化 Gin
	r := gin.Default()

	// 路由处理程序
	r.GET("/suppliers", getSuppliers)
	r.GET("/suppliers/:id", getSupplier)
	r.POST("/suppliers", createSupplier)
	r.PUT("/suppliers/:id", updateSupplier)
	r.DELETE("/suppliers/:id", deleteSupplier)

	// 启动服务器
	r.Run(":8080")
	fmt.Printf("启动服务器\n")
}

// 获取所有供应商
func getSuppliers(c *gin.Context) {
	var suppliers []Supplier
	db.Find(&suppliers)
	c.JSON(http.StatusOK, suppliers)
}

// 根据ID获取供应商
func getSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier Supplier
	db.First(&supplier, id)
	c.JSON(http.StatusOK, supplier)
}

// 创建供应商
func createSupplier(c *gin.Context) {
	var supplier Supplier
	c.BindJSON(&supplier)
	db.Create(&supplier)
	c.JSON(http.StatusCreated, supplier)
}

// 更新供应商
func updateSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier Supplier
	db.First(&supplier, id)
	c.BindJSON(&supplier)
	db.Save(&supplier)
	c.JSON(http.StatusOK, supplier)
}

// 删除供应商
func deleteSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier Supplier
	db.First(&supplier, id)
	db.Delete(&supplier)
	c.JSON(http.StatusNoContent, gin.H{"message": "Supplier deleted successfully"})
}
