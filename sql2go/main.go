package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	//	sql := `
	//CREATE TABLE IF NOT EXISTS t_person (
	//  f_age INT(11) NULL,
	//  f_id INT(11) PRIMARY KEY AUTO_INCREMENT NOT NULL,
	//  f_name VARCHAR(30) NOT NULL,
	//  f_sex VARCHAR(2) NULL
	//  ) ENGINE=InnoDB;
	//`
	//	args := sql2go.NewConvertArgs().SetGenJson(true).
	//		SetPackageName("test").
	//		SetColPrefix("f_").
	//		SetTablePrefix("t_")
	//
	//	code, err := sql2go.FromSql(sql, args)
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}
	//	f, err := os.Create("db_struct3.go")
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}
	//	defer f.Close()
	//	f.Write(code)

	f := excelize.NewFile()
	sheetName := "表格名称"
	f.SetSheetName("Sheet1", sheetName)
	err := f.SaveAs("a.xlsx")
	if err != nil {
		fmt.Println(fmt.Sprintf("❎:%s", err))
		return
	}
}
