package user

import (
	. "rest-api/handler"
	"rest-api/pkg/errno"
	"rest-api/service"

	"github.com/gin-gonic/gin"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})

}
