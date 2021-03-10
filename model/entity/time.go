package entity

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type DateTime time.Time

func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = DateTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = DateTime(now)
	return
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

// 写入 mysql 时调用
func (t DateTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" || t.String() == "" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// 检出 mysql 时调用
func (t *DateTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = DateTime(tTime)
	return nil
}

// 获取数据时使用
func (t DateTime) String() string {
	dateTime := time.Time(t).Format(TimeFormat)
	if dateTime == "0001-01-01 00:00:00" {
		return ""
	}
	return dateTime
}
