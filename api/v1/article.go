package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	arts, code, total := model.GetCateArticle(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    arts,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code := model.CreateArticle(&data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 获取文章列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	arts, code, total := model.GetArticles(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    arts,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 获取单篇文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	art, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    art,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改文章
func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code := model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArticle(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
