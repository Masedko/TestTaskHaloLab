package dtos

import "time"

type TimeIntervalUnixTimestamp struct {
	StarTime string `validate:"required" query:"fromDateTime"`
	EndTime  string `validate:"required" query:"untilDateTime"`
}

type TimeInterval struct {
	StarTime time.Time `validate:"required" query:"fromDateTime"`
	EndTime  time.Time `validate:"required" query:"untilDateTime"`
}
