package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)

	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.CreateCategory(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// @Summary 获取分类列表
// @Produce  json
// @Param name query string false "分类名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Category "成功"
// @Router /api/v1/categorys [get]
func GetCategorys(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	cates := model.GetCategorys(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cates,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改分类
func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	_ = c.ShouldBindJSON(&data)

	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		model.EditCategory(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})

}

// 删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteCategory(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
