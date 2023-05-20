package main

//超级大的问题?????

//func main() {
//	//1
//	key := md5.Sum([]byte("hello"))
//	fmt.Println("key1:", key)
//	fmt.Printf("%x\n", key) // 输出 16 个十六进制数字，共 32 个字符
//	//2
//	h := md5.New()
//	h.Write([]byte("hello"))
//	key2 := h.Sum(nil)
//	fmt.Println("key2", key2)
//	fmt.Printf("%x\n", key2) // 输出 16 个十六进制数字，共 32 个字符
//	//3
//	m := md5.New()
//	key3 := m.Sum([]byte("hello"))
//	fmt.Println("key3:", key3)
//	fmt.Printf("%x\n", key3) // 输出 16 个十六进制数字，共 32 个字符
//	//4
//	h2 := md5.New()
//	key4 := h2.Sum([]byte("hello,world"))
//	fmt.Println("key4:", key4)
//	fmt.Printf("%x\n", key4)
//	//5
//	h3 := md5.New()
//	h3.Write([]byte("hello,world!"))
//	key5 := h3.Sum(nil)
//	fmt.Println("key5:", key5)
//	fmt.Printf("%x\n", key5) // 输出 16 个十六进制数字，共 32 个字符
//	//
//	h4 := md5.New()
//	h4.Write([]byte("hello"))     //添加需要加密的内容，不会改变得到的哈希值长度
//	key6 := h4.Sum([]byte("mhx")) //添加加密用的额外密钥，但会导致得出的哈希值长度改变
//	fmt.Println(key6)
//	//
//	h5 := md5.New()
//	h5.Write([]byte("hello,world"))
//	key7 := h5.Sum([]byte("mhx"))
//	fmt.Println(key7)
//	//
//	h6 := md5.New()
//	h6.Write([]byte("hello"))
//	key8 := h6.Sum([]byte("secret"))
//	fmt.Println(key8)
//}
