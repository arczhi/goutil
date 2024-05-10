package batch

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestPageBatch(t *testing.T) {
	pageBatch := NewPageBatch(1000, 100, 6)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := pageBatch.Exec(func(pageIndex int, offset int) error {
		func(ctx context.Context) {
			fmt.Println("[FETCHING DATA] ...", "pageIndex", pageIndex, "offset", offset)
		}(ctx)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
