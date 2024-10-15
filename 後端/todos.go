package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Todo 結構
type Todo struct {
	Id   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

// 全域變數用於儲存 TODO 列表
var globalTodos []Todo

func main() {
	r := gin.Default()
	api := r.Group("/api")

	// POST 新增任務（固定為「吃飯」）
	api.POST("/todo", func(c *gin.Context) {
		newTodo := Todo{
			Id:   len(globalTodos) + 1,
			Task: "吃飯",
			Done: false,
		}
		globalTodos = append(globalTodos, newTodo)
		c.JSON(http.StatusCreated, newTodo) // 回傳新增的任務
	})

	// GET 獲取所有任務
	api.GET("/read", func(c *gin.Context) {
		c.JSON(http.StatusOK, globalTodos) // 回傳所有任務
	})

	// PUT 更新任務狀態
	api.PUT("/mark", func(c *gin.Context) {
		taskId := c.Query("id")   // 取得 task ID
		state := c.Query("state") // 取得狀態

		// 將 taskId 轉換為 int
		id, err := strconv.Atoi(taskId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		// 更新狀態
		for i := range globalTodos {
			if globalTodos[i].Id == id { // 比較 ID
				if state == "done" {
					globalTodos[i].Done = true
				} else if state == "pending" {
					globalTodos[i].Done = false
				}
				c.JSON(http.StatusOK, globalTodos[i]) // 回傳更新後的任務
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"}) // 找不到任務
	})

	// DELETE 刪除任務
	api.DELETE("/remove", func(c *gin.Context) {
		taskId := c.Query("id") // 取得 task ID

		// 將 taskId 轉換為 int
		id, err := strconv.Atoi(taskId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		for i := range globalTodos {
			if globalTodos[i].Id == id {
				globalTodos = append(globalTodos[:i], globalTodos[i+1:]...) // 移除任務
				c.JSON(http.StatusOK, gin.H{"message": "Todo removed"})     // 回傳成功消息
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"}) // 找不到任務
	})

	// 啟動伺服器
	r.Run(":8008") // listen and serve on 0.0.0.0:8080
}
