package utils

import "time"

func GetExcelFileName() string {
	now := time.Now()
	currTime := now.Format("2006-01-02_15-04-05")
	return "reports_" + currTime + ".xlsx"
}
