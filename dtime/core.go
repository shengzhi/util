// Package dtime 工具包 时间辅助操作
package dtime

import (
	"fmt"
	"time"
)

// JSONTime JSON 时间， 时间戳
type JSONTime int64

func (t JSONTime) MarshalJSON() ([]byte, error) {
	if int64(t) <= 0 {
		return []byte(`""`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, time.Unix(int64(t), 0).Format("2006-01-02 15:04:05"))), nil
}

func (t JSONTime) Time() time.Time {
	return time.Unix(int64(t), 0)
}
func Now() JSONTime { return JSONTime(time.Now().Unix()) }

func Today() JSONTime {
	y, m, d := time.Now().Date()
	return JSONTime(time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix())
}
