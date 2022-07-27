package main

// 内置 response
/*
	NotFound 函数，包装一个 404 状态码和一个额外的信息
	ServeFile 函数，从文件系统提供文件，返回给请求者
	ServeContent 函数，它可以把实现了 io.ReadSeeker 接口的任何东西里面的内容返回给请求者
	还可以处理 Range 请求(范围请求)，如果只请求了资源的一部分内容，那么
	ServeContent 就可以如此响应。而 ServeFile 或io.Copy则不行。
	Redirect 函数，告诉客户端重定向到另一个URL
*/
func main() {

}
