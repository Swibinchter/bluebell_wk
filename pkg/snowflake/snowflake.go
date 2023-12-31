package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

// 声明一个snowflake节点变量，节点用于区分分布式系统中不同的机器节点
var node *sf.Node

// 根据给定的起始时间和机器id生成节点
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	// 将给定的字符串转换成go的时间类型，即time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	// 将st转换成纳秒级别的时间戳，再除1000000转换成毫秒级别的时间戳
	// 赋值给Epoch，作为算法的起始时间
	sf.Epoch = st.UnixNano() / 1000000
	// 创建一个新节点
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	// 从节点生成一个int64类型的雪花算法的id
	return node.Generate().Int64()
}
