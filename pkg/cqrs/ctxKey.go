package cqrs

import (
	"context"
	"time"
)

const (
	ctxKeyId = iota
	ctxKeyTimestamp
)

func GetIdFromCtx(ctx context.Context) string {
	s, _ := ctx.Value(ctxKeyId).(string)
	return s
}

func GetTimestampFromCtx(ctx context.Context) time.Time {
	s, _ := ctx.Value(ctxKeyTimestamp).(time.Time)
	return s
}
