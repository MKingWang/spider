package queue

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func initRedis(redisAddr string) (redis.Conn, error) {
	return redis.Dial("tcp", redisAddr)
}

/*
将节点内容写入redis队列
*/
func pushQueue(key interface{}) error {
	redisCli, err := initRedis("192.168.25.57:6379")
	if err != nil {
		fmt.Errorf("init redis err:%v", err)
		return err
	}
	defer redisCli.Close()

	_, err = redisCli.Do("LPUSH", "spiderqueue", key)
	if err != nil {
		fmt.Errorf("push queue err:%v", err)
		return err
	}

	return nil
}

/*
获取爬虫需要的节点信息
先从键盘输入，后续再调整
*/
func getInfo() {
	var k string
	fmt.Printf("Input queue key:")
	fmt.Scan(&k)
	pushQueue(k)
	fmt.Println("push queue ok")
}

func Run() {
	for {
		getInfo()
	}
}
