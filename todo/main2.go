package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:ljqdemysql@tcp(127.0.0.1:3306)/demo1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	db.AutoMigrate(&UserTodo{})
}
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createtodo)
		v1.GET("/", fetchAlltodo)
		v1.GET("/:id", fetchSingletodo)
		v1.PUT("/:id", updatetodo)
		v1.DELETE("/:id", deletetodo)
	}
	router.Run()
}

type (
	UserTodo struct {
		gorm.Model
		UserID              string    `gorm:"userid"`
		UserTodoID          string    `gorm:"usertodoid"`
		UserTodoTitle       string    `gorm:"usertodotitle"`
		UserTodoDescription string    `gorm:"usertododescription"`
		UserTodoDueTime     time.Time `gorm:"duetime"`
		UserTodoRemindTime  time.Time `gorm:"remindtime"`
		Status              bool      `gorm:"status"`
	}
	// Base struct {
	// 	ID        uint      `json:"id" gorm:"primary_key"` //什么意思？grom和json都用？不可以只用一种？
	// 	CreatedAt time.Time `json:"createdat" gorm:"DEFAULT CURRENT_TIMESTAMP"`
	// 	UpdateAt  time.Time `json:"updateat" gorm:"DEFAULT CURRENT_YIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	// }
)

func createtodo(c *gin.Context) {
	const shortForm = "2006-jan-02"
	duetime, _ := time.Parse(shortForm, c.PostForm("duetime"))
	remindtime, _ := time.Parse(shortForm, c.PostForm("remindtime"))
	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		panic("状态输入错误")
	}
	todo := UserTodo{UserID: c.PostForm("userid"), UserTodoID: c.PostForm("usertodoid"), UserTodoTitle: c.PostForm("usertodotitle"),
		UserTodoDescription: c.PostForm("usertododescription"), UserTodoDueTime: duetime, UserTodoRemindTime: remindtime, Status: status}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created sucessful!", "resourceID": todo.ID})
}
func fetchAlltodo(c *gin.Context) {
	var todos []UserTodo
	//var _todos []UserTodo
	db.Find(&todos)
	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	//数据库数据共享，不用重新定义todos吧？
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todos})
}
func fetchSingletodo(c *gin.Context) {
	var todo UserTodo
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}
func updatetodo(c *gin.Context) {
	var todo UserTodo
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Model(&todo).Update("UserID", c.PostForm("userid"))
	db.Model(&todo).Update("UserTodoID", c.PostForm("usertodoid"))
	db.Model(&todo).Update("UserTodoTitle", c.PostForm("usertodotitle"))
	db.Model(&todo).Update("UserTodoDescription", c.PostForm("usertododescription"))
	const shortForm = "2006-jan-02"
	duetime, _ := time.Parse(shortForm, c.PostForm("duetime"))
	remindtime, _ := time.Parse(shortForm, c.PostForm("remindtime"))
	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		panic("状态输入错误")
	}
	db.Model(&todo).Update("UserTodoDueTime", duetime)
	db.Model(&todo).Update("UserTodoRemindTime", remindtime)
	db.Model(&todo).Update("Status", status)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}
func deletetodo(c *gin.Context) {
	var todo UserTodo
	todoID := c.Param("id")
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
