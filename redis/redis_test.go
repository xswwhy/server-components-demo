package redisOper

import (
	"fmt"
	"log"
	"testing"
)

func TestOperatorRedis(t *testing.T) {
	redis, err := NewRedis("127.0.0.1:6379", "", 0)
	if err != nil {
		log.Fatal("redis创建失败")
	}
	fmt.Println("redis连接成功!")

	redis.OperRedis()
	redis.ZsetExample()
}
