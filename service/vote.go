package service

import "RedBubble/models"

// 投票功能：
// 1. 用户投票的数据
//

// 投票算法：
// 投一票加432分

/*
direction=1时，两种情况：
之前可能投过反对票，现在改赞成票
之前没有投过票，现在投赞成票

direction=0时，两种情况：
之前可能投过反对票，现在取消
之前可能投过赞成票，现在取消

direction=-1时，两种情况：
之前可能投过赞成票，现在改反对票
之前没有投过票，现在投反对票

投票限制：
帖子自发表之日起一周内允许投票，过期不可以投票
	1.到期之后把redis中存的赞成和反对票持久化到mysql
	2.到期后删除KeyPostVotedZSetPrefix

*/

// 给帖子投票
func VoteForPost(userID int64, p *models.ParamVoteData) {


}