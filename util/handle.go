package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	BUILD_SUCCESS = 2 //打包成功
	BUILD_FAIL 	  = 3 	  //打包失败
)

type Task struct {
	Task_id string
	Status int
	Apk_name string
	Token string
}

/**
 * websocket推送消息
 */
type Message struct {
	data int
	typ string
	message string
}

/**
 * websocket服务返回信息
 */
type Response struct {
	Code int
	Message string
	Data interface{}
	Contact_email string
}

/**
 * 处理打包回调
 */
func DealWith(task *Task) error {
	db := &Db{}
	db.Connect()

	packageQueue,result,err := db.FindQueueByTaskId(task.Task_id);
	if err != nil {
		return err
	}

	if !result {
		//数据不存在
		Error(fmt.Sprintf("package queue不存在 %v",task));
		return errors.New("package queue不存在")
	}

	//转换打包状态
	if task.Status == 1 {
		packageQueue.Status = BUILD_SUCCESS
	}else{
		packageQueue.Status = BUILD_FAIL
	}
	packageQueue.ApkName = task.Apk_name
	packageQueue.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	//更新数据库
	var affected int64
	affected,err = db.engine.Where("id = ?", packageQueue.Id).Update(packageQueue)
	if err != nil {
		//数据库查询错误
		Error(fmt.Sprintf("更新package queue 失败,packageQueue: %v",packageQueue));
		return err
	}
	if affected <= 0 {
		//更新失败
		Error(fmt.Sprintf("更新package queue 更新失败,packageQueue: %v",packageQueue));
		return err
	}
	Info(fmt.Sprintf("打包回调成功,packageQueue: %v",packageQueue))

	var packageObj *PackageMgr
	packageObj,result,err = db.FindPackage(packageQueue.PackageId)
	if err != nil {
		//查找package失败
		Error(fmt.Sprintf("查找package失败,packageQueue: %v",packageQueue));
		return nil
	}
	if !result {
		Error(fmt.Sprintf("pacakge 不存在,packageQueue: %v",packageQueue));
		return nil
	}

	var channelId int
	result,err,channelId = db.FindChannelConfig(packageObj.ChannelConfigId)
	if err != nil {
		Error(fmt.Sprintf("查找channel_config失败,packageQueue: %v",packageQueue));
		return nil
	}
	if !result {
		Error(fmt.Sprintf("FindChannelConfig 不存在,packageQueue: -%v",packageQueue));
	}

	var channel map[string]string
	result,err,channel = db.FindChannel(channelId)
	if err != nil {
		//查找channel失败
		Error(fmt.Sprintf("查找channel失败,packageQueue: %v",packageQueue));
		return nil
	}
	if !result {
		Error(fmt.Sprintf("FindChannel 不存在 ,packageQueue: -%v",packageQueue));
		return nil
	}

	//通知
	var msg string
	if channel["type"] == "2" {
		msg = "官方包的打包任务完成了，快去下载吧"
	}else{
		msg = channel["channel_name"]+"的打包任务完成了，快去下载吧"
	}
	message := Message{packageQueue.PackageId,"package",msg}
	go Notify(&message,task.Token)

	return nil
}

/**
 * 通知长连接
 */
func Notify(message *Message,token string)  {
	var err error
	data := url.Values{"message": {fmt.Sprintf(`{"data":%v,"type":"%v","message":"%v"}`,message.data,message.typ,message.message)}}
	urlString := GetSocketConfig()+"/push?token="+token;

	var client HttpClient
	err = client.Post(urlString,data, nil, func(resp *http.Response) {
		response := Response{}
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body,&response)
		if err != nil {
			Error(fmt.Sprintf("解析错误: -%v",err));
		}

		Info(fmt.Sprintf("消息推送结果: -%v",response));
	})

	if err != nil {
		fmt.Println(err)
		Info(fmt.Sprintf("post请求出错: -%v",err));
	}
}

