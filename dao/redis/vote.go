package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// 本项目使用简化版的投票分数
// 投一票加60*60*24/200=432分，即投200票能给帖子延续一整天的热度

/*
>>>投票的限制
每个帖子自发表之日起一个星期内允许运用投票，超过一个星期就不允许在投票了
1.到期之后将redis保存的赞成票数和反对票数存储到mysql中
2.到期之后删除对应Key

>>>投票请求分析
direction=1时
	>原来是反对票-1，现在是赞成票1，票数=1-(-1)=2
	>原来是没有投票0，现在是赞成票1，票数=1-0=1
direction=0时
	>原来是反对票-1，现在是取消投票0，票数=0-(-1)=1
	>原来是赞成票1，现在是取消投票0，票数=0-(1)=-1
direction=1时
	>原来是赞成票1，现在是反对票-1，票数=(-1)-1=-2
	>原来是没有投票0，现在是反对票-1，票数=(-1)-0=-1
总结，票数的变动就是direction新值减去旧值
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600 // 每一周多少秒
	scorePerVote     = 432           // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

// 创建帖子同时在redis中创建
func CreatePost(postId, communityId int64) error {
	// 创建连接，设置五秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// pipeline可以一次性发送多个命令给redis，执行成功后一起返回，减少单个命令传输过程中的时间损耗
	pipeline := client.TxPipeline()

	// 帖子与时间的Zset，key名字是bluebell:post:time，元素是帖子id，分数是时间戳
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	// 帖子与分数的Zset，key名字是bluebell:post:score，元素是帖子id，分数是根据投票数计算的分数
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()), //?初始分数是当前时间，后续分数在这基础上增加折算的秒数，可以延续有效时间
		Member: postId,
	})

	// 将社区id加到社区的set中
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityId)))
	pipeline.SAdd(ctx, cKey, postId)

	// 执行pipeline命令
	c, err := pipeline.Exec(ctx)
	fmt.Printf("pipeline执行后的结果是%v\n", c)
	return err
}

// VoteForPost
func VoteForPost(userId, postId string, value float64) (err error) {
	// 创建连接，设置五秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 判断帖子时间是否过期
	postTime := client.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postId).Val()
	// 超过一周的话提示过期
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	// 先查之前的投票记录
	ov := client.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+postId), userId).Val()
	fmt.Printf("之前投票记录的查询结果是%v\n", ov)
	// 如果本次跟之前是一样的票，提示不能重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	// 计算帖子的分数变动
	diff := value - ov
	fmt.Printf("计算出来的投票差值diff是%v\n", diff)
	// 更新帖子的分数
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), diff*scorePerVote, postId)

	// 记录本次用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postId), userId)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postId), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userId,
		})
	}
	_, err = pipeline.Exec(ctx)
	return
}
