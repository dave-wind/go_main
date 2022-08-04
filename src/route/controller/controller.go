package controller

/*
这边写一个controller 文件
大写 是可以导出的
还要写注释
*/

// RegisterRoutes ...
func RegisterRoutes() {

	// static

	// 相同package 模块 可以直接使用
	registerHomeRoutes()
	registerAboutRoutes()
	registerCompanyRoutes()
}
