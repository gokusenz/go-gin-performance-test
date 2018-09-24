package helpers

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// ContextHelper is a grpc context helper
type ContextHelper struct{}

// GetData function is to get data from grpc context
func (r *ContextHelper) GetData(ctx context.Context, key string) string {
	data := ""
	// Get user id from context
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values := md.Get(key)
		if len(values) >= 1 {
			data = values[0]
		}
	}
	return data
}
