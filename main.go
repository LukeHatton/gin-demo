package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Album struct (Model)
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums 获取所有的albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums) // 将slice转换为JSON并设置到Context对象的response body中
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id") // 获取URL中的参数。这个参数必须与URL中的占位符一致
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album of id 【" + id + "】 is not found"})
}

// postAlbums 添加一个新的album。请求的数据已经包含在了Context对象中
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil { // 将请求的JSON数据绑定到newAlbum结构体指针
		log.Println(err.Error())
	}

	albums = append(albums, newAlbum)            // 将新的album添加到slice中
	c.IndentedJSON(http.StatusCreated, newAlbum) // 状态码201表示资源已经成功创建
}

// 初始化路由配置
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	/* ================ 路由配置：albums ================= */
	/*
	 将getAlbums处理器函数与GET请求映射到/albums endpoint
	 注意这里的getAlbums是一个函数名，而不是一个函数调用结果。如果传入函数调用结果，参数就会是getAlbums()
	*/
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumsById) // :id是一个参数占位符，可以通过c.Param("id")来获取参数值
	r.POST("/albums", postAlbums)       // 将postAlbums处理器函数与POST请求映射到/albums endpoint

	return r
}

func main() {
	router := setupRouter()
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("an error occurred", err.Error())
	}
}
