package utils

import (
	"context"

	"github.com/google/uuid"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

func GetRequestID(ctx context.Context) string {
	val := ctx.Value(constants.RequestID)

	if val != "" {
		return val.(string)
	}

	return uuid.UUID{}.String()
}
