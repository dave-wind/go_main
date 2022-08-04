package model

import (
	"crud/common"
)

type Company struct {
	ID       string
	Name     string
	NickName string
}

// 获取所有公司 ⚠️:companies 在返回值类型声明,语法糖 无需 return xx
func GetAllCompanies() (companies []Company, err error) {
	sql := "SELECT id, name, nickname FROM company"
	rows, err := common.Db.Query(sql)
	if err != nil {
		return
	}
	for rows.Next() {
		c := Company{}
		err = rows.Scan(&c.ID, &c.Name, &c.NickName)
		if err != nil {
			return
		}
		companies = append(companies, c)
	}
	return
}

// GetCompany ..
func GetCompany(id string) (company Company, err error) {
	// sql :=  "SELECT id, name, nickname FROM company WHERE id=$1"
	sql := "SELECT id,name,nickname FROM company WHERE id = ?"
	err = common.Db.QueryRow(sql, id).Scan(&company.ID, &company.Name, &company.NickName)
	return
}

// Insert ...
func (company *Company) Insert() (err error) {
	// sql := "INSERT INTO company (id, name, nickname) VALUES ($1, $2, $3)"
	sql := "INSERT INTO company (id, name, nickname) VALUES (?, ?, ?)"
	stmt, err := common.Db.Prepare(sql)
	if err != nil {
		return
	}
	_, err = stmt.Exec(company.ID, company.Name, company.NickName)
	if err != nil {
		return
	}
	return
}

// Update ...
func (company *Company) Update() (err error) {
	// sql := "UPDATE company set name=$1, nickname=$2 WHERE id=$3"
	sql := "UPDATE company set name = ?, nickname = ? WHERE id = ?"
	_, err = common.Db.Exec(sql, company.Name, company.NickName, company.ID)
	return
}

// DeleteCompany ...
func DeleteCompany(id string) (err error) {
	sql := "DELETE FROM company WHERE id = ?"
	_, err = common.Db.Exec(sql, id)
	return
}
