package main

import (
	"github.com/astaxie/beego"
)

// UserController 是一个示例的控制器
type UserController struct {
	beego.Controller
}

// Get 方法处理 GET 请求
func (c *UserController) Get() {
	c.Ctx.WriteString("Hello, World!")
}

func main() {
	// 注册路由
	beego.Router("/user", &UserController{})

	// 启动应用程序
	beego.Run()
}
