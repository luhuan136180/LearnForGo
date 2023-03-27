package error

type Error interface {
	Caller() []CallerInfo
	Wraped() []error
	Code() int
	error

	private()
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine int
}

//定义一系列辅助函数
//func New(msg string) error
//func NewWithCode(code int, msg string) error
//
//func Wrap(err error, msg string) error
//func WrapWithCode(code int, err error, msg string) error
//
//func FromJson(json string) (Error, error)
//func ToJson(err error) string
