package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type FormatTime time.Time

type Base struct {
	ID         string `gorm:"primaryKey,type:nvarchar,size:50,not null,default:''"`
	CreateTime FormatTime
	DeleteTime FormatTime
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 FormatTime
func (f *FormatTime) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal FormatTime value:", value))
	}

	result := time.Time{}
	err := json.Unmarshal(bytes, &result)
	*f = FormatTime(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 FormatTime value
func (f FormatTime) Value() (driver.Value, error) {
	if time.Time(f).Unix() != 0 {
		return time.Now(), nil
	}
	return time.Time(f), nil
}
