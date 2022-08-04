package controller

import (
	"crud/funcs"
	"crud/model"
	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

func registerRoutes() {
	http.HandleFunc("/", listCompanies)
	http.HandleFunc("/companies", listCompanies)
	http.HandleFunc("/companies/seed", seed)
	http.HandleFunc("/companies/add", addCompany)

	http.HandleFunc("/companies/edit/", editCompany)

	http.HandleFunc("/companies/delete/", deleteCompany)
}

func listCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := model.GetAllCompanies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		// 自定义 渲染函数map
		funcMap := template.FuncMap{"add": funcs.Add}
		// 然后构建template 传入函数
		t := template.New("companies").Funcs(funcMap)
		// 解析模版
		t, err = t.ParseFiles("./templates/_layout.html", "./templates/company/list.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		t.ExecuteTemplate(w, "layout", companies)
	}
}

func addCompany(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t := template.New("company-add")
		t, err := t.ParseFiles("./templates/_layout.html", "./templates/company/add.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			t.ExecuteTemplate(w, "layout", nil)
		}
		return
		// post 方法 验证表单
	case http.MethodPost:
		newCompany := model.Company{}
		// form post
		newCompany.ID = r.PostFormValue("id")
		newCompany.Name = r.PostFormValue("name")
		newCompany.NickName = r.PostFormValue("nickName")
		// fmt.Println(newCompany.Name)
		err := newCompany.Insert()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			// 重定向
			http.Redirect(w, r, "/companies", http.StatusSeeOther)
		}
		return
	}
}

func seed(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Ok")
}

func editCompany(w http.ResponseWriter, r *http.Request) {

	// 找到
	idPattern := regexp.MustCompile(`/companies/edit/([a-zA-Z0-9]*$)`)

	// 获取字符串, 返回切片, 返回值1:原字符串, 2:后面的数字
	matches := idPattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		id := matches[1]

		fmt.Println("edit id---", id)
		switch r.Method {
		case http.MethodGet:
			company, err := model.GetCompany(id)
			fmt.Println("find company---", company)
			if err == nil {
				t := template.New("company-edit")
				t, err := t.ParseFiles("./templates/_layout.html", "./templates/company/edit.html")
				if err == nil {
					t.ExecuteTemplate(w, "layout", company)
					return
				}
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		case http.MethodPost:
			company := &model.Company{}
			company.ID = r.PostFormValue("id")
			company.Name = r.PostFormValue("name")
			company.NickName = r.PostFormValue("nickName")
			// 更新
			err := company.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			} else {
				http.Redirect(w, r, "/companies", http.StatusSeeOther)
			}
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteCompany(w http.ResponseWriter, r *http.Request) {
	idPattern := regexp.MustCompile(`/companies/delete/([a-zA-Z0-9]*$)`)

	matches := idPattern.FindStringSubmatch(r.URL.Path)

	if len(matches) > 0 {
		id := matches[1]

		if r.Method == http.MethodDelete {
			fmt.Println("delete id---", id)
			err := model.DeleteCompany(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			http.Redirect(w, r, "/companies", http.StatusSeeOther)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
