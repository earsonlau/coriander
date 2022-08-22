package pageInfo

import (
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetPageInfo(c *gin.Context) (int, int) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var(
		page int
		size int
		err error
	)
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	} 
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	return page,size
}
