package user_api

import (
	"context"

	"efishery_test/user_api/entity"
)

type ContextID struct {
	name string
}

var ContextKeyName = &ContextID{name: "user_api"}

func MetaFromContext(ctx context.Context) entity.ResourceClaims {
	return ctx.Value(ContextKeyName).(entity.ResourceClaims)
}

func GetUserID(ctx context.Context) int64 {
	return MetaFromContext(ctx).ID
}
