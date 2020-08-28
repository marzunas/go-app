package handler

import (
	"context"
	"testing"
)

//Test class for ShutdownHandler
//TODO: Additional test cases with mock HTTP objects to test handler function
func Test_ShutdownHandler(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ShutdownHandler(cancel).ServeHTTP(nil, nil)
	if ctx.Err() == nil {
		t.Error("should have been cancelled")
	}
}
