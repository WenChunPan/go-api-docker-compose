package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// main的function已寫在其他檔案
// func main() {
// 	hw1()
// }

type TODO struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

// python list
var globalTodo = []TODO{}

func hw1_api() {
	//post create / new task

	// 初始化 Gin router
	r := gin.Default()
	// 定義 API 群組
	api := r.Group("/api")

	// api.POST("/create", func(c *gin.Context) {
	// 	//接收資料
	// 	task := c.PostForm("task") // 接收傳統HTML表單的用法
	// 	done := c.PostForm("done") == "true" // 將字串轉換為布林值
	// 	//回傳資料
	// 	c.JSON(200, gin.H{
	// 		"task": task,
	// 		"done": done,
	// 	})
	// 	// 給todo結構類型的新任務一個變數名稱
	// 	p := TODO{
	// 		Task: task,
	// 		Done: done,
	// 	}
	// 	globalTodo = append(globalTodo, p)
	// })

	api.POST("/create", func(c *gin.Context) {
		p := TODO{}
		if err := c.ShouldBindJSON(&p); err != nil {
			// 如果資料格式無效
			c.JSON(http.StatusBadRequest, gin.H{"error": "無效的資料格式"})
			return
		}

		//自動產生ID，從1開始
		p.ID = len(globalTodo)
		globalTodo = append(globalTodo, p)

		// 回傳資料，當OK時回傳p的資料
		c.JSON(http.StatusOK, p)
	})

	//GET  read list / filter done or not
	api.GET("/read", func(c *gin.Context) {
		doneFilter := c.Query("done") //取得查詢參數

		var filteredTodos []TODO // 儲存過濾後的任務
		for _, todo := range globalTodo {
			if doneFilter == "true" && todo.Done { //如果doneFilter為true且todo.Done為true
				// 將過濾後的任務加入filteredTodes
				filteredTodos = append(filteredTodos, todo)
			} else if doneFilter == "false" && !todo.Done { //如果doneFilter為false且todo.Done為false
				// 將過濾後的任務加入filteredTodes
				filteredTodos = append(filteredTodos, todo)
			} else if doneFilter == "" {
				// 如果沒有提供過濾條件，則顯示所有任務
				filteredTodos = append(filteredTodos, todo)
			}
		}

		c.JSON(200, filteredTodos) // 回傳過濾後的任務列表

	})

	//PUT / update task
	api.PUT("/update/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index")) // 取得 URL 參數 index並轉換為整數
		if err != nil || index < 0 || index >= len(globalTodo) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "無效的索引"})
			return
		}

		var updateData struct {
			Done bool `json:"done"`
		}

		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "無效的資料格式"})
			return
		}

		// 更新任務的完成狀態
		globalTodo[index].Done = updateData.Done
		c.JSON(http.StatusOK, gin.H{"message": "更新成功", "task": globalTodo[index]})
	})
	//DELETE / delete task
	api.DELETE("/delete/:index", func(c *gin.Context) {
		index, err := strconv.Atoi(c.Param("index")) // 取得路由參數並轉換為整數
		if err != nil || index < 0 || index >= len(globalTodo) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "無效的索引"})
			return
		}

		// 刪除指定索引的任務
		globalTodo = append(globalTodo[:index], globalTodo[index+1:]...)
		c.JSON(http.StatusOK, gin.H{"message": "任務已刪除"})
	})
	//啟動伺服器
	r.Run(":8081")
}
