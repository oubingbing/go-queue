package util

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
)

type PackageMgr struct {
	Id int `xorm:"'id'"`
	ChannelConfigId int `xorm:"'channel_config_id'"`
	GameId int `xorm:"'game_id'"`
	GameVersionId int `xorm:"'game_version_id'"`
	AccountId int `xorm:"'account_id'"`
	BasePath string `xorm:"'base_path'"`
	ApkName string `xorm:"'apk_name'"`
	ServerType int `xorm:"'server_type'"`
	PackagedAt string `xorm:"packed_at"`
	CalledAt string `xorm:"called_at"`
	CreatedAt string `xorm:"created_at"`
	UpdatedAt string `xorm:"updated_at"`
	DeletedAt string `xorm:"deleted_at"`
}

type PackageQueue struct {
	Id int `xorm:"'id'"`
	PackageId int `xorm:"'package_id'"`
	ChannelVersionId int `xorm:"'channel_version_id'"`
	Message string `xorm:"'message'"`
	ApkName string `xorm:"'apk_name'"`
	Status int `xorm:"'status'"`
	TaskId string `xorm:"'task_id'"`
	CreatedAt string `xorm:"created_at"`
	UpdatedAt string `xorm:"updated_at"`
	DeletedAt string `xorm:"deleted_at"`
}

type Db struct {
	Host string
	Port int
	Driver string
	Database string
	Username string
	Password string
	engine *xorm.Engine
}

/**
 * 连接数据库
 */
func (db *Db) Connect()  {
	var err error
	GetMysqlConfig(db)
	db.engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",db.Username,db.Password,db.Host,db.Port,db.Database))
	if err != nil {
		Error(fmt.Sprintf("数据库连接失败-%v",err.Error()));
	}
}

/**
 * 查找package queue
 */
func (db *Db) FindQueueByTaskId(taskId string) (*PackageQueue,bool,error) {
	packageQueue := PackageQueue{}
	result , err := db.engine.Where("task_id = ?",taskId).Get(&packageQueue)
	if err != nil {
		Error(fmt.Sprintf("FindQueueByTaskId error -%v",err.Error()));
	}
	return &packageQueue,result,err
}

/**
 * 查找package
 */
func (db *Db) FindPackage(packageId int) (*PackageMgr,bool,error) {
	packageObj := PackageMgr{Id:packageId}
	result , err := db.engine.Get(&packageObj)
	if err != nil {
		Error(fmt.Sprintf("FindPackage error -%v",err.Error()));
	}
	return &packageObj,result,err
}

/**
 * 查找渠道配置
 */
func (db *Db) FindChannelConfig(channelConfigId int) (bool,error,int) {
	var channelId int
	result,err := db.engine.
		Table("channel_configs").
		Where("id = ?", channelConfigId).
		Cols("channel_id").
		Get(&channelId)
	if err != nil {
		Error(fmt.Sprintf("FindChannelConfig error -%v",err.Error()));
	}
	return result,err,channelId
}

/**
 * 查找渠道
 */
func (db *Db) FindChannel(channelId int) (bool,error,map[string]string) {
	//var channelName string
	mp := make(map[string]string)
	mp["channel_name"] = ""
	mp["type"] = ""
	result,err := db.engine.
		Table("channel_bases").
		Where("id = ?", channelId).
		Cols("channel_name","type").
		Get(&mp)
	if err != nil {
		Error(fmt.Sprintf("FindChannel error -%v",err.Error()));
	}
	return result,err,mp
}