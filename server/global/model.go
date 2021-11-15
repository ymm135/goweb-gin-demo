package global

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type GLOBAL_MODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt FormatTime     // 创建时间
	UpdatedAt FormatTime     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

//FormatTime 自定义时间
type FormatTime struct {
	time.Time
}

const timeLayout = "2006-01-02 15:04:05"

func (t *FormatTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(timeLayout, timeStr)
	*t = FormatTime{t1}
	return err
}

func (t FormatTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", t.Time.Format(timeLayout))
	return []byte(formatted), nil
}

func (t FormatTime) Value() (driver.Value, error) {
	// FormatTime 转换成 time.Time 类型
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return t.Time.Format(timeLayout), nil
}

func (t *FormatTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = FormatTime{vt}
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *FormatTime) String() string {
	return t.Time.Format(timeLayout)
}
