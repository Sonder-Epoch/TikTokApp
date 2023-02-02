package common

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type VideoDTO struct {
	Id            int64   `json:"id"`
	Author        UserDTO `json:"author"`
	PlayUrl       string  `json:"play_url"`
	CoverUrl      string  `json:"cover_url"`
	FavoriteCount int64   `json:"favorite_count"`
	CommentCount  int64   `json:"comment_count"`
	IsFavorite    bool    `json:"is_favorite"`
	Title         string  `json:"title"`
}

type CommentDTO struct {
	Id         int64   `json:"id,omitempty"`
	User       UserDTO `json:"user"`
	Content    string  `json:"content,omitempty"`
	CreateDate string  `json:"create_date,omitempty"`
}

type UserDTO struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
