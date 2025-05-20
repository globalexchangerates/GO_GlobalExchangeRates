package globalexchangerates

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	
	if s == "" || s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		t, err = time.Parse(time.RFC3339, s)
		if err != nil {
			return fmt.Errorf("cannot parse date: %v", err)
		}
	}
	
	ct.Time = t
	return nil
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(ct.Time.Format("2006-01-02"))
}
