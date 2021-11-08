package utils

import "time"

// 获取本周几几点的时间, 不如5-10 就是获取本周5的10点的时间
func GetTimeFromWeek(week int, hour int) string {
	now := time.Now()
	//先计算今天是星期几
	offset := int(time.Monday - now.Weekday())
	if offset > 0 { //周天的情况
		offset = -6
	}

	//目前offset是周一的,如果计算周几的，偏移量
	offset = week - offset - 1

	date := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time.Local).
		AddDate(0, 0, offset)

	return date.Format("2006-01-02 15:04:05")
}
