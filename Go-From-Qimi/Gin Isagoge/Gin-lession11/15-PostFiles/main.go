package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB

	r.LoadHTMLFiles("15-PostFiles/index.html")

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/post", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		} else {
			//dst :=fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./", f.Filename)
			//将读取的文件保存到本地
			c.SaveUploadedFile(f, dst)

			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": fmt.Sprintf("`%s` uploaded!", f.Filename),
			})
		}
		//将读取的文件保存

	})
	r.Run(":8080")
}
