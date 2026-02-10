package util

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

func PtrInt32(v int32) *int32                 { return &v }
func PtrInt64(v int64) *int64                 { return &v }
func PtrFloat32(v float32) *float32           { return &v }
func PtrFloat64(v float64) *float64           { return &v }
func PtrBool(v bool) *bool                    { return &v }
func PtrString(v string) *string              { return &v }
func PtrTime(v time.Time) *time.Time          { return &v }
func PtrUUID(v uuid.UUID) *uuid.UUID          { return &v }
func PtrInterface(v interface{}) *interface{} { return &v }

func Int32Value(p *int32) string {
	if p == nil {
		return ""
	}
	return strconv.Itoa(int(*p))
}

func Int64Value(p *int64) string {
	if p == nil {
		return ""
	}
	return strconv.FormatInt(*p, 10)
}

func Float64Value(p *float64) string {
	if p == nil {
		return ""
	}
	return strconv.FormatFloat(*p, 'f', -1, 64)
}

func BoolValue(p *bool) string {
	if p == nil {
		return ""
	}
	return strconv.FormatBool(*p)
}

func StringValue(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func Float32Value(p *float32) string {
	if p == nil {
		return ""
	}
	return strconv.FormatFloat(float64(*p), 'f', -1, 32)
}

func PtrSliceToStringSlice(src []*string) []string {
	dst := make([]string, 0, len(src))
	for _, v := range src {
		if v != nil {
			dst = append(dst, *v)
		}
	}
	return dst
}
