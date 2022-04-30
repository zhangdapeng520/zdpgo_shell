package zdpgo_shell

import "github.com/zhangdapeng520/zdpgo_shell/encoding/simplifiedchinese"

/*
@Time : 2022/4/30 8:36
@Author : 张大鹏
@File : encoding
@Software: Goland2021.3.1
@Description: 解决中文乱码等问题
*/

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
