package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"
	. "github.com/hunterhug/go_image"
	"github.com/astaxie/beego"
)

func UnixToDate(timestamp int) string {

	t := time.Unix(int64(timestamp), 0)

	return t.Format("2006-01-02 15:04:05")
}

//2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		beego.Info(err)
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}

func GetUnixNano() int64 {
	return time.Now().UnixNano()
}
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

func Hello(in string) (out string) {
	out = in + "world"
	return
}
func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func ResizeImage(filename string) {
	extName := path.Ext(filename)
	resizeImage := strings.Split(beego.AppConfig.String("resizeImageSize"),",")
	for i:=0;i<len(resizeImage);i++{
		width,_ := strconv.Atoi(resizeImage[i])
		saveName := filename + "_" + resizeImage[i] + "*" + resizeImage[i] + extName
		err := ThumbnailF2F(filename, saveName, width, width)
		if err != nil {
			fmt.Println("制作缩略图错误: ",err)
		}
	}
}

//格式化图片
func FormatImg(picName string) string {
	ossStatus, err := beego.AppConfig.Bool("ossStatus")
	if err != nil {
		//判断目录前面是否有/
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName
	}
	if ossStatus {
		return beego.AppConfig.String("ossDomain") + "/" + picName
	} else {
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName

	}

}