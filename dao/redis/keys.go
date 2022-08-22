package redis

//redis key
//尽量使用命名空间区分

const (
	KeyPrefix        = "coriander:"
	KeyPostTimeZSet  = "post:time" //zset;帖子及发帖时间
	KeyPostScoreZSet = "post:score"//zset;帖子及投票分数
	KeyPostVotedZSetPrefix     = "post:voted:"//zset;记录用户投票及类型;参数是post id
)