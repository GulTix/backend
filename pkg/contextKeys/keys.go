package contextkeys

import "context"

type contextKey string

const (
	UserInfoKey contextKey = "userInfo"
)

func GetValue[T any](ctx context.Context, key contextKey) (T, bool) {
	val, ok := ctx.Value(key).(T)

	return val, ok
}
