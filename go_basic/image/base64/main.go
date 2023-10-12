package base64

import (
	"fmt"
	"strings"
)

//获取base64图片大小，返回 Byte 数
func getImageSize(base64Data []byte) int {
	fmt.Println("origin size:", len(base64Data))
	//把头部去掉
	var str = strings.ReplaceAll(string(base64Data), "data:image/png;base64,", "")
	// 找到等号，把等号也去掉
	var equalIndex = strings.Index(str, "=")
	if equalIndex > 0 {
		str = str[0:equalIndex]
	}
	// 原来的字符流大小，单位为字节
	var strLength = len(str)
	// 计算后得到的文件流大小，单位为字节
	var fileLength = strLength - (strLength/8)*2
	// 由字节转换为kb
	//var size = ""
	//size = (fileLength / 1024).toFixed(2)
	//var sizeStr = size + ""              //转成字符串
	//var index = sizeStr.indexOf(".")     //获取小数点处的索引
	//var dou = sizeStr.substr(index+1, 2) //获取小数点后两位的值
	//if dou == "00" {                     //判断后两位是否为00，如果是则删除00
	//	return sizeStr.substring(0, index) + sizeStr.substr(index+3, 2)
	//}
	return fileLength
}
