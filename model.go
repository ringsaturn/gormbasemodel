package gormbasemodel

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type MyTime struct {
	time.Time
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("%v", t.Unix())
	return []byte(formatted), nil
}

func (t *MyTime) UnmarshalJSON(data []byte) error {
	intData := binary.BigEndian.Uint64(data)
	t.Time = time.Unix(int64(intData), 0)
	return nil
}

func (t MyTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *MyTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = MyTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt *MyTime        `json:"createdAt"`
	UpdatedAt *MyTime        `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
