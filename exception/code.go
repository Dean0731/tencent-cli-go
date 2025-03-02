package exception

var (
	UnknownError       = &CustomError{"UnknownError", "未知的错误"}
	FileNotFound       = &CustomError{"FileNotFound", "文件%s未找到"}
	FileFormatError    = &CustomError{"FileFormatError", "文件%s格式错误"}
	InvalidParam       = &CustomError{"InvalidParam", "无效的参数%s值：%v"}
	ParamIsBlank       = &CustomError{"ParamIsBlank", "参数%s值为空"}
	JsonLoadError      = &CustomError{"JsonLoadError", "JSON字符串序列化失败：%s"}
	JsonDumpError      = &CustomError{"JsonLoadError", "JSON字符串反序列化失败"}
	MethodNotImplement = &CustomError{"MethodNotImplement", "方法%s暂未实现"}
)
