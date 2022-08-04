package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

func registerCompanyRoutes() {
	http.HandleFunc("/companies", HandleCompanies)
	http.HandleFunc("/companies/", HandleCompany)
}

func HandleCompanies(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "companies.html")
	t.ExecuteTemplate(w, "layout", "nil")
}

/*
 路由参数:

带参数路由: 根据路由参数,创建一族不同的页面

*/
func HandleCompany(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html", "company.html")

	pattern, _ := regexp.Compile(`/companies/(\d+)`)
	// 获取字符串, 返回切片, 返回值1:原字符串, 2:后面的数字
	matches := pattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		fmt.Println(matches[0])
		companyID, _ := strconv.Atoi(matches[1])
		t.ExecuteTemplate(w, "layout", companyID)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
