package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
)

// 声明一个snowflake节点变量，节点用于区分分布式系统中不同的机器节点
var node *snowflake.Node

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
	snowflake.Epoch = st.UnixNano() / 1000000
	// 创建一个新节点
	node, err = snowflake.NewNode(machineID)
	return
}

func GenID() int64 {
	// 从节点生成一个int64类型的雪花算法的id
	return node.Generate().Int64()
}

func main() {
	err := Init("2023-01-01", 1)
	if err != nil {
		fmt.Printf("Snowflake init failed,err:%v\n", err)
		return
	}
	// 生成分布式ID
	id := GenID()
	fmt.Printf("生成的id是%d\n", id)
}
