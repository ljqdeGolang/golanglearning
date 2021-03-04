// package main

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// //var db *gorm.DB

// // func init() {
// // 	var err error
// // 	db, err = gorm.Open("mysql", "root:ljqdemysql@tcp(127.0.0.1:3306)/demo1?charset=utf8&parseTime=True&loc=Local")
// // 	if err != nil {
// // 		fmt.Println(err.Error())
// // 		panic("failed to connect database")
// // 	}
// // 	db.AutoMigrate(&todoModel{})
// // }

// // func main() {
// // 	router := gin.Default()
// // 	v1 := router.Group("/api/v1/todos")
// // 	{
// // 		v1.POST("/", createTodo)
// // 		v1.GET("/", fetchAllTodo)
// // 		v1.GET("/:id", fetchSingleTodo)
// // 		v1.PUT("/:id", updateTodo)
// // 		v1.DELETE("/:id", deleteTodo)
// // 	}
// // 	router.Run()
// // }

// // type (
// // 	todoModel struct {
// // 		gorm.Model
// // 		Title     string `json:"title"`
// // 		Completed int    `json:"completed"`
// // 	}
// // 	transformedTodo struct {
// // 		ID        uint   `json:"id"`
// // 		Title     string `json:"title"`
// // 		Completed bool   `json:"completed"`
// // 	}
// // )

// func createTodo(c *gin.Context) {
// 	completed, _ := strconv.Atoi(c.PostForm("completed")) //为什么要加_？返回值有两个？
// 	todo := todoModel{Title: c.PostForm("title"), Completed: completed}
// 	db.Save(&todo)
// 	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created sucessful!", "resourceID": todo.ID})
// }

// //fetchAllTodo中不是每次都会输出No todo found吗？
// func fetchAllTodo(c *gin.Context) {
// 	var todos []todoModel
// 	var _todos []transformedTodo
// 	db.Find(&todos) //这个Find的目的是？找到数据库所有的字段？
// 	if len(todos) <= 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}
// 	for _, item := range todos {
// 		completed := false
// 		if item.Completed == 1 {
// 			completed = true
// 		} else {
// 			completed = false
// 		}
// 		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
// 	} //这个for循环的目的是在外界输入的completed为1的情况下，将输入打印出来？
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
// }
// func fetchSingleTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")
// 	db.First(&todo, todoID)
// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}
// 	completed := false
// 	if todo.Completed == 1 {
// 		completed = true
// 	} else {
// 		completed = false
// 	}
// 	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
// }
// func updateTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")
// 	db.First(&todo, todoID)
// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}
// 	db.Model(&todo).Update("title", c.PostForm("title"))
// 	completed, _ := strconv.Atoi(c.PostForm("completed"))
// 	db.Model(&todo).Update("completed", completed)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
// }
// func deleteTodo(c *gin.Context) {
// 	var todo todoModel
// 	todoID := c.Param("id")
// 	db.First(&todo, todoID)
// 	if todo.ID == 0 {
// 		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
// 		return
// 	}
// 	db.Delete(&todo)
// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
// }
