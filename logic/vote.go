package logic

import (
	"goWebCli/dao/redis"
	"goWebCli/model"
	"strconv"

	"go.uber.org/zap"
)

// VoteForPost 给帖子投票的业务
func VoteForPost(userId int64, p *model.ParamVote) error {
	zap.L().Debug("VoteForPost", zap.Int64("userid", userId), zap.String("post_id", p.PostId),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userId)), p.PostId, float64(p.Direction))
}
