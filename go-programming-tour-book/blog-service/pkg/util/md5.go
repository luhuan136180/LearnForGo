package util

import (
	"crypto/md5"
	"encoding/hex"
)

//该方法用于针对上传后的文件名格式化，简单来讲，将文件名 MD5 后再进行写入，防止直接把原始名称就暴露出去了。
func EncodeMD5(value string) string {
	m := md5.New() //返回一个可以进行 MD5 散列运算的 hash.Hash 接口实例m
	//下面这哈那个代码将一个字符串或字节切片作为输入数据写入到MD5哈希算法的实例m中进行计算和处理。这样可以对字符串或字节切片进行MD5哈希计算，生成一个唯一的128位的散列值。
	//Write() 方法可以将字符串或字节切片写入 MD5 实例中
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil)) //// EncodeToString返回src的十六进制编码。
}
