package models

type VoteRequest struct {
	PostId int `json:"post_id" validate:"required"`
}

type Vote struct {
	VoteRequest
	UserId int `json:"user_id"`
}

func ParseToVote(v *Vote, vr VoteRequest) {
	v.PostId = vr.PostId
}
