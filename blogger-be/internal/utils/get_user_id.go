package utils

import (
	"context"

	"github.com/Akash-Manikandan/blogger-be/internal/middleware"
)

// Function to retrieve userID in service logic
func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(middleware.UserIDKey).(string)
	return userID, ok
}
