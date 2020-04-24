package main

import (
	"encoding/json"
	"fmt"
	"project/util"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go listen(&wg)
	wg.Wait()
}

/**
 * 监听redis
 */
func listen(wg *sync.WaitGroup)  {
	client := &util.RedisClient{}
	client.ConnectRedis()
	t := time.Second * 59

	queue := util.GetCallBackConfig()
	client.OnQueue(wg,queue,t, func(result []string, e error) {
		if len(result) > 0 {
			util.Info(fmt.Sprintf("取出需要完结的打包任务数据：%v",result));
			var task util.Task
			err := json.Unmarshal([]byte(result[1]),&task)
			if err != nil {
				//解析json错误
				util.Error(fmt.Sprintf("onQueue:解析json失败：%v",result[1]));
			}else{
				err = util.DealWith(&task)
				if err != nil {
					//出错重新推入队列
					var tsk []byte
					tsk,err = json.Marshal(task)
					client.Push(queue,string(tsk))
				}
			}
		}
	})
}

