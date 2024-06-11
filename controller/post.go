package controller

import (
	"acana/logic"
	"acana/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	// step1: get params and check params
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON() error", zap.Any("err", err))
		zap.L().Error("create post with vaild parameters")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// get the userID from context `c`
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID

	// step2: create post
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CrearePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// step3: return response
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	// step1: get params
	post_id_str := c.Param("id")
	post_id, err := strconv.ParseInt(post_id_str, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// step2: search post data by post id
	data, err := logic.GetPostByID(post_id)
	if err != nil {
		zap.L().Error("logic.GetPostByID() falied", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// step3: return
	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// get post list by the params which trans from frontend
// oder by create time or vote numbers
func GetOrderPostListHandler(c *gin.Context) {
	// step1: get params
	// (GET request params: /api/v1/orderedposts?page=1&s=10&order=time)
	p := new(models.ParamPostList)
	if err := c.ShouldBindQuery(&p); err != nil {
		zap.L().Error("GetOrderPostListHandler get the invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// step2: quert by list of id in redis
	data, err := logic.GetPostListNew(p)
	// step3: query the post's detail data in mysql db by list of id from redis

	if err != nil {
		zap.L().Error("logic.GetPostListNew() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// // get post list by community
// func GetCommunityPostListHandler(c *gin.Context) {
// 	// step1: get params
// 	// (GET request params: /api/v1/orderedposts?page=1&s=10&order=time)
// 	p := new(models.ParamPostList)
// 	if err := c.ShouldBindQuery(&p); err != nil {
// 		zap.L().Error("GetCommunityPostListHandler get the invalid params", zap.Error(err))
// 		ResponseError(c, CodeInvalidParam)
// 		return
// 	}
// 	// step2: quert by list of id in redis
// 	data, err := logic.GetCommunityPostList(p)
// 	// step3: query the post's detail data in mysql db by list of id from redis

// 	if err != nil {
// 		zap.L().Error("logic.GetOrderPostList() failed", zap.Error(err))
// 		ResponseError(c, CodeServerBusy)
// 		return
// 	}
// 	ResponseSuccess(c, data)
// }
