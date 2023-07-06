package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User 是一个示例的模型结构体
type User struct {
	Id   int    `orm:"auto"`
	Name string `orm:"size(100)"`
}

// UserController 是一个示例的控制器
type UserController struct {
	beego.Controller
}

// Create 方法处理创建用户的请求
func (c *UserController) Create() {
	// 从请求中获取用户信息
	name := c.GetString("name")

	// 创建用户对象
	user := User{Name: name}

	// 使用ORM插入用户数据
	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": "User created successfully"}
	}
	c.ServeJSON()
}

// Read 方法处理获取用户信息的请求
func (c *UserController) Read() {
	// 从请求中获取用户ID
	id, _ := c.GetInt(":id")

	// 使用ORM查询用户数据
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "User not found"}
	} else if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = user
	}
	c.ServeJSON()
}

// GetUsers 方法处理获取所有用户信息的请求
func (c *UserController) GetUsers() {
	// 使用ORM查询所有用户数据
	o := orm.NewOrm()
	var users []User
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = users
	}
	c.ServeJSON()
}

// Update 方法处理更新用户信息的请求
func (c *UserController) Update() {
	// 从请求中获取用户ID和新的用户名
	id, _ := c.GetInt(":id")
	name := c.GetString("name")

	// 使用ORM查询用户数据并更新
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "User not found"}
	} else if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		user.Name = name
		_, err := o.Update(&user)
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": err.Error()}
		} else {
			c.Data["json"] = map[string]string{"message": "User updated successfully"}
		}
	}
	c.ServeJSON()
}

// Delete 方法处理删除用户的请求
func (c *UserController) Delete() {
	// 从请求中获取用户ID
	id, _ := c.GetInt(":id")

	// 使用ORM查询用户数据并删除
	o := orm.NewOrm()
	user := User{Id: id}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "User not found"}
	} else if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		_, err := o.Delete(&user)
		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": err.Error()}
		} else {
			c.Data["json"] = map[string]string{"message": "User deleted successfully"}
		}
	}
	c.ServeJSON()
}

func init() {
	// 注册模型
	orm.RegisterModel(new(User))

	// 配置数据库连接
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:wu12345@tcp(localhost:3306)/heavendb?charset=utf8")

	// 自动创建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	// 注册路由
	beego.Router("/user", &UserController{}, "post:Create")
	beego.Router("/user/:id", &UserController{}, "get:Read")
	beego.Router("/user/:id", &UserController{}, "put:Update")
	beego.Router("/user/:id", &UserController{}, "delete:Delete")
	beego.Router("/users", &UserController{}, "get:GetUsers")
	// 启动应用程序
	beego.Run()
}
