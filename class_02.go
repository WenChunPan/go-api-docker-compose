package main

import (
	// "encoding/json"
	// "fmt"

	"github.com/gin-gonic/gin"
)

// func class_02() {
// 	// python dict
// 	// m := map[string]string{
// 	// 	"a": "1",
// 	// 	"b": "abc",
// 	// 	"c": "true",
// 	// }

// 	// for k, v := range m {
// 	// 	fmt.Println(k, "==>", v)
// 	// }

// 	// python dict 有不同型別的value
// 	m := map[string]any{
// 		"a": "1",
// 		"b": "abc",
// 		"c": true,
// 	}
// 	// 不同的print方式:
// 	// for k := range m {
// 	// 	fmt.Println(k, "==>", m[k])
// 	// }

// 	// v, ok := m["a"]
// 	// fmt.Println(ok, v)

// 	// v = m["d"]
// 	// fmt.Println("d ==> ", v)

// 	v, ok := m["d"] //ok為true或false
// 	if ok {
// 		fmt.Println("e==>", v)
// 	} else {
// 		fmt.Println("e not found")
// 	}
// 	// python class
// 	//go沒有class，只有struct，先輸入tys會跑出下面這串定義struct的東西
// 	type Person struct {
// 		Name    string `json:"name"` //tag
// 		Age     int    `json:"年齡"`
// 		Student bool   `json:"-"` //使用-：忽略
// 		job     string
// 	}

// 	p := Person{
// 		Name:    "John",
// 		Age:     18,
// 		Student: true,
// 		job:     "student",
// 	}
// 	fmt.Println(p.Name, p.Age, p.Student)

// 	// python json dumps
// 	b, err := json.Marshal(p) //轉換成json格式看有沒有成功
// 	//沒有成功就是panic有error
// 	if err != nil {
// 		fmt.Println("轉換失敗", err)
// 	} else {
// 		// fmt.Println("json(bytes)==>", b)
// 		fmt.Println("json(string)==>", string(b))
// 	}

// 	// python json loads
// 	str := `{"Name": "Yu", "年齡": 33, "student": flase}`
// 	var p2 Person //宣告 type

// 	err = json.Unmarshal([]byte(str), &p2)
// 	if err != nil {
// 		fmt.Println("Json轉換失敗", err)
// 	} else {
// 		fmt.Printf("p2: %+v/n", p2)
// 	}

// }

func class_02_api() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!!")
	})

	api := r.Group("/api")
	api.GET("/hello", func(c *gin.Context) {
		name := c.Query("name") //Query:取得URL參數，主要是用於取得GET參數，優點是可以設定預設值，缺點是不太安全，與path param的不同是，path param可以直接修改變數
		c.String(200, "Hello, %s", name)
	})

	// 另一種資料的取得方式：path param，優點是
	api.GET("/hello/:name", func(c *gin.Context) { // :name是變數，可直接修改name這個變數
		name := c.Param("name")
		c.String(200, "Hello, %s", name)
	})

	//post
	api.POST("/register", func(c *gin.Context) {
		//接收資料
		username := c.PostForm("username")
		password := c.PostForm("password")
		avatar := c.PostForm("avatar")
		// //也可以用這種方式
		// password := c.PostForm("password")
		// username := c.PostForm("username")

		//回傳資料
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
			"avatar":   avatar,
		})
	})
	api.POST("login", func(c *gin.Context) {
		//這個是請求的資料格式
		type LoginRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		//req 變數，將用於儲存解析後的 JSON 資料
		var req LoginRequest
		err := c.ShouldBindJSON(&req) //shouldbindjson:解析請求的 JSON 資料並存入 req 變數
		//如果解析失敗，回傳錯誤
		if err != nil {
			c.JSON(200, gin.H{
				"error": "傳入參數錯誤：" + err.Error(),
			})
			return
		}

		c.JSON(200, req)
	})
	r.Run(":8080")

}
