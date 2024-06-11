package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required" equal:"password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1" `
}

// query params about getting the post list
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`
	Page        int64  `json:"page" form:"page" example:"1"`
	Size        int64  `json:"size" form:"size" example:"10"`
	Order       string `json:"order" form:"order" example:"score"`
}
