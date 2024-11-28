package v1

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/wejectchen/ginblog/utils"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr: utils.RedisAddr + ":" + utils.RedisPort,
	})
}

type ScanDetailType struct {
	Account string `json:"account" form:"account"`
	Pwd     string `json:"pwd" form:"pwd"`
	TaskId  string `json:"task_id" form:"task_id"`
	Risk    string `json:"risk" form:"risk"`
	Level   string `json:"level" form:"level"`
}

func ScanUpdate(c *gin.Context) {
	var apiUrl string = utils.Url
	var account string = utils.Account
	var pwd string = utils.Pwd

	assettype := c.PostForm("assettype")
	ftypename := c.PostForm("ftypename")
	url := c.PostForm("url")
	scan := c.DefaultPostForm("scan", "true")
	compliance := c.DefaultPostForm("compliance", "false")
	appname := c.PostForm("appname")
	appversion := c.PostForm("appversion")
	apptype := c.PostForm("apptype")
	licensed := c.DefaultPostForm("licensed", "true")

	body := map[string]string{
		"account":    account,
		"pwd":        pwd,
		"assettype":  assettype,
		"ftypename":  ftypename,
		"url":        url,
		"scan":       scan,
		"compliance": compliance,
		"appname":    appname,
		"appversion": appversion,
		"apptype":    apptype,
		"licensed":   licensed,
	}

	res, err := utils.PostCurl(apiUrl+"/system/api/upload", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func ScanDetail(c *gin.Context) {
	var apiUrl string = utils.Url
	var account string = utils.Account
	var pwd string = utils.Pwd

	// var scanDetailType ScanDetailType
	// _ = c.Bind(&scanDetailType)

	task_id := c.PostForm("task_id")
	risk := c.DefaultPostForm("risk", "false")
	level := c.DefaultPostForm("level", "1,2,3")

	body := map[string]string{
		"account": account,
		"pwd":     pwd,
		"task_id": task_id,
		"risk":    risk,
		"level":   level,
	}

	res, err := utils.PostCurl(apiUrl+"/system/api/detail", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func ScanStatus(c *gin.Context) {
	var apiUrl string = utils.Url
	var account string = utils.Account
	var pwd string = utils.Pwd

	task_id := c.PostForm("task_id")

	body := map[string]string{
		"account": account,
		"pwd":     pwd,
		"task_id": task_id,
	}

	res, err := utils.PostCurl(apiUrl+"/system/api/status", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func ScanExport(c *gin.Context) {
	var apiUrl string = utils.Url
	var account string = utils.Account
	var pwd string = utils.Pwd

	task_id := c.Query("task_id")
	app_name := c.Query("app_name")
	appversion := c.Query("appversion")

	body := map[string]string{
		"account":    account,
		"pwd":        pwd,
		"task_id":    task_id,
		"app_name":   app_name,
		"appversion": appversion,
	}
	// 将map转换为查询字符串
	var queryStrings []string
	for key, value := range body {
		queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(value)))
	}
	queryString := strings.Join(queryStrings, "&")

	fullUrl := apiUrl + "/system/api/export?" + queryString
	// 调用GetFile方法来处理文件下载
	err := utils.GetFile(c, fullUrl, "aaa.pdf")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 如果一切顺利，返回成功响应（虽然在这个例子中我们不需要额外的响应数据）
	c.JSON(http.StatusOK, gin.H{})
}

// getTokenFromRedis 从 Redis 中获取 token
func GetTokenFromRedis() (string, error) {
	ctx := context.Background()
	val, err := rdb.Get(ctx, "token_key").Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

// saveTokenToRedis 将 token 存储到 Redis 中，并设置过期时间
func SaveTokenToRedis(token string) error {
	ctx := context.Background()
	err := rdb.Set(ctx, "token_key", token, 2*time.Hour).Err() // 设置 2 小时过期时间
	return err
}
