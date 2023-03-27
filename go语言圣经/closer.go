package main

import "os"

func main() {
	//	当文件 studygolang.txt 不存在或找不到时，
	//	file.Close() 会 panic，因为 file 是 nil。
	//	因此，应该将 defer file.Close() 放在错误检查之后。
	file, err := os.Open("studygolang.txt")
	if err != nil {
		return
	}
	defer file.Close()
}
