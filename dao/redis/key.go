package redis

// 存放跟Redis key相关的内容

// key的命名建议使用这种“项目前缀:元素:分数来源”的形式，方便查询和分割
// 命名时使用冒号，会让redis工具自动归类
const (
	// 待补充这是啥的解释？？？只是为了更规范的命令key
	// 项目前缀
	Prefix = "bluebell:"
	// zset类型的key，元素是帖子，分数是时间戳
	KeyPostTimeZSet = "post:time"
	// zset类型的key，元素是帖子，分数是投票数
	KeyPostScoreZSet = "post:score"
	// zset类型的key(需要补充postid)，元素是帖子，分数是投的赞成票1还是反对票-1，PF是prefix待补充的意思
	KeyPostVotedZSetPF = "post:voted:"
	// set类型的key，保存每个分区下帖子的id
	KeyCommunitySetPF = "community:"
)

// 给redis的key加上前缀，规范命名
func getRedisKey(key string) string {
	return Prefix + key
}
