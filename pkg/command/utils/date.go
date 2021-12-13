package utils

import (
    "context"
    "go.uber.org/zap"
    "time"
    "yang-backend/pkg/ginlog"
)

func ParseTime(ctx context.Context, format string, value string) (time.Time, error) {
    time, err := time.Parse(format, value)
    if err != nil {
        ginlog.CtxLogger(ctx).Error("time parse failed",
            zap.Error(err))
        return time, err
    }
    return time, err
}

func GetYesterday() time.Time {
    curTime := time.Now()
    yesterdayTime := curTime.AddDate(0, 0, -1)
    return yesterdayTime
}
