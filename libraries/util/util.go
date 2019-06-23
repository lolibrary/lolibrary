package util

import (
	"time"

	"github.com/monzo/terrors"
)

// TimeToProto turns a Go time.Time into a string for protoc.
func TimeToProto(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Truncate(time.Millisecond).Format(time.RFC3339Nano)
}

// ProtoToTime converts a protoc datetime to a time.Time.
func ProtoToTime(date string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		return time.Time{}, terrors.Wrap(err, nil)
	}

	return t, nil
}

// ForceProtoToTime is like ProtoToTime, except it returns a zero date on error.
func ForceProtoToTime(date string) time.Time {
	if t, err := ProtoToTime(date); err == nil {
		return t
	}

	return time.Time{}
}