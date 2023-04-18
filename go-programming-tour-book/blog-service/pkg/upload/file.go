package upload

import (
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type FileType int

const TypeImage FileType = iota + 1

//获取文件名，通过文件名后缀帅选出原始文件名进行MD5加密，最后返回经过加密处理后的文件
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext) //TrimSuffix返回不带后缀字符串的s。也就是说返回example.txt 的 example
	fileName = util.EncodeMD5(fileName)       //对文件名加密

	return fileName + ext //返回一个文件名加密的文件名
}

//包装了一个获取文件名扩展名的函数
func GetFileExt(name string) string {
	//path.Ext(name)函数是用来获取文件路径名中的扩展名的。它会返回最后一个点号后面的所有字符串，
	//如果没有点号就返回空字符串。例如，如果文件名为“example.txt”，
	//则path.Ext(name)函数将返回“.txt”。如果文件名为“example”，则该函数将返回空字符串。
	return path.Ext(name)
}

//获取文件保存地址，因为我们已经在配置文件中写入了，所以返回配置中的文件保存目录即可，
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//一下开始编写监察文件的相关方法，因为需要确保在文件写入时它已经达到了必备条件，否则要给出相应的标准错误提示，继续在文件内新增代码

//检查保存路径是否存在，通过os.stat方法获取文件的描述信息FileInfo,并用os.IsNotExist 方法进行判断，
//其原理是利用 os.Stat 方法所返回的 error 值与系统中所定义的 oserror.ErrNotExist 进行判断，
//以此达到校验效果。
func CheckSavePath(dst string) bool {
	//os.Stat 函数，返回dst指定的文件或目录的FileInfo对象和一个错误对象err。
	_, err := os.Stat(dst)
	//os.IsNotExist 函数判断err是否为一个“不存在”的错误，如果是则返回true，否则返回false。
	//ErrNotExist   = fs.ErrNotExist   // "file does not exist"
	return os.IsNotExist(err)
}

//CheckContainExt：检查文件后缀是否包含在约定的后缀配置项中，
//需要的是所上传的文件的后缀有可能是大写、小写、大小写等，
//因此我们需要调用 strings.ToUpper 方法统一转为大写（固定的格式）来进行匹配。
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	//将扩展名转换为大写形式：,将获取到的扩展名字符串全部转换为大写形式，避免大小写问题导致误判。
	ext = strings.ToUpper(ext)

	switch t { //根据函数传入的文件类型 t，执行与之对应的 case 子句内的操作。
	case TypeImage: //照片,图片
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			//代码使用了一个 for 循环，在允许的扩展名列表中逐个比较实际扩展名是否相等
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		} //如果循环结束后，还没有找到匹配的扩展名，说明不符合要求，返回 false。
	}
	return false
}

//检查文件大小是否超出最大大小限制
func CheckMaxSize(t FileType, f multipart.File) bool {
	//将multipart.File里的内容读取到一个byte数组中，存在content变量中
	contant, _ := ioutil.ReadAll(f)
	//计算文件的大小
	size := len(contant)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

//检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

//文件的检查代码部分结束，开始进行文件写入/创建的操作

//创建在上传文件时所使用的保存目录，在方法内部调用：os.MkdirAll ，传入os.FileMode权限位 去递归创建所需要的的所有目录结构，
//若涉及的目录结构已存在，则不会进行任何操作，直接返回nil
func CreateSavePath(dst string, perm os.FileMode) error { //第二个参数：传入一个权限位
	err := os.MkdirAll(dst, perm) //MkdirAll创建一个名为第一个参数的目录，以及任何必要的父目录，并返回nil，否则返回一个错误。
	if err != nil {
		return err
	}
	return nil
}

//SaveFile：保存所上传的文件，该方法主要是通过调用 os.Create 方法创建目标地址的文件，
//再通过 file.Open 方法打开源地址的文件，结合 io.Copy 方法实现两者之间的文件内容拷贝。
func SaveFile(file *multipart.FileHeader, dst string) error {
	//打开源文件，如果打开出错，则返回错误；如果成功，则在函数执行完后关闭文件。
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//创建一个目标文件，如果出错则返回错误；如果成功则在函数执行完后关闭文件。
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	//将源文件内容拷贝到目标文件中，并返回错误（如果有的话）。
	_, err = io.Copy(out, src)
	return err
}
