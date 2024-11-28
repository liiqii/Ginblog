package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// string转成int：
func StringToInt(str string) (res int, err error) {
	res, err = strconv.Atoi(str)
	return
}

// string转成int64：
func StringToInt64(str string) (res int64, err error) {
	res, err = strconv.ParseInt(str, 10, 64)
	return
}

// string转float
func StringToFloat(str string) (res float64, err error) {
	res, err = strconv.ParseFloat(str, 64)
	return
}

// string转bool
func StringToBool(str string) (res bool, err error) {
	res, err = strconv.ParseBool(str)
	return
}

// int转成string：
func IntToString(i int) (res string) {
	res = strconv.Itoa(i)
	return
}

// int64转成string：
func Int64ToString(i int64) (res string) {
	res = strconv.FormatInt(i, 10)
	return
}

// float转int
func FloatToInt(str string) (res int, err error) {
	res, err = strconv.Atoi(fmt.Sprintf("%1.0f", str))
	return
}

// 字符串 转 []byte
func StringToByte(str string) (res []byte) {
	res = []byte(str)
	return
}

// []byte 转 字符串
func ByteToString(str []byte) (res string) {
	res = string(str)
	return
}

// datetime转换成时间字符串
func DatetimeToTime() (res string) {
	now := time.Now()                       // 当前 datetime 时间
	res = now.Format("2006-01-02 15:04:05") // 把当前 datetime 时间转换成时间字符串
	return
}

// datetime转换成时间戳
func DatetimeToTimestamp() (res int64) {
	now := time.Now()
	res = now.Unix()
	return
}

// 时间戳转换成时间字符串
func TimestampToTime(str int64) (res string) {
	res = time.Unix(str, 0).Format("2006-01-02 15:04:05")
	return
}

// 时间字符串转换成时间戳
func TimeToTimestamp(str string) (res int64) {
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	res = t1.Unix()
	return
}

// 生成一个随机数
func GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	ok1, _ := regexp.MatchString(".[1|2|3|4|5|6|7|8|9]", string(result))
	ok2, _ := regexp.MatchString(".[Z|X|C|V|B|N|M|A|S|D|F|G|H|J|K|L|Q|W|E|R|T|Y|U|I|P]", string(result))
	if ok1 && ok2 {
		return string(result)
	} else {
		return GetRandomString(l)
	}
}

// go获取主机IP
func GetOutboundIP(ip string) (res string, err error) {
	// conn, err := net.Dial("udp", "8.8.8.8:80")
	conn, err := net.Dial("udp", ip)
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	res = strings.Split(localAddr.IP.String(), ":")[0]
	return
}

// go 如何判断一个字符串是否在切片列表里
func StrIsExistInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// 如何在切片中查找指定参数
func SliceIsExistInStr(s []string, t string) (int, bool) {
	// go中使用 sort.searchXXX 方法，在排序好的切片中查找指定的方法，但是其返回是对应的查找元素不存在时，待插入的位置下标(元素插入在返回下标前)。
	iIndex := sort.SearchStrings(s, t)
	bExist := iIndex != len(s) && s[iIndex] == t

	return iIndex, bExist
}

// go 执行shell脚本 在指定目录下执行指定脚本
func ExecShell(directory, script string) bool {
	// directory := "/root/a"
	// script := "./aa"

	// 创建一个执行命令的对象
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("cd %s && %s", directory, script))

	// 设置命令的工作目录
	cmd.Dir = directory

	// 将命令的输出连接到当前进程的输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
	// 检查命令执行结果
	if cmd.ProcessState.Success() {
		return true
	} else {
		return false
	}
}

// go 判断一个文件或文件夹是否存在
func PathExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// go InArray
func InArray(arr interface{}, ele interface{}) bool {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(ele, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}

// curl post请求 data 为请求参数 map[string]interface{}
func PostCurl(url string, data interface{}) (map[string]interface{}, error) {
	postBody, _ := json.Marshal(data)
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("JSON 解析失败: %w", err)
	}

	return result, nil
}

// curl get请求
func GetCurl(url string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("JSON 解析失败: %w", err)
	}

	return result, nil
}

// curl 下载文件
func GetFile(c *gin.Context, url, filename string) error {
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 发送GET请求到外部API
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching file from API: %w", err)
	}
	defer resp.Body.Close()

	// 检查外部API的响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	// 设置HTTP响应头，提示浏览器这是一个文件下载
	// 注意：这里我们假设文件是PDF类型，但根据实际情况，您可能需要从响应头中获取正确的MIME类型
	c.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	c.Header("Content-Type", "application/pdf")

	// 将外部API的响应内容写入到HTTP响应中
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		return fmt.Errorf("error sending file to client: %w", err)
	}

	return nil
}
