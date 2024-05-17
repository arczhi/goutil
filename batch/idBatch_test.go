package batch

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestIdBatch(t *testing.T) {
	idBatch := NewIdBatch(1000, 100, 6)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := idBatch.Exec(func(roundIndex int) error {
		func(ctx context.Context) {
			fmt.Println("[DOING TASK]", "roundIndex", roundIndex)
		}(ctx)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}

// experiment
func TestIdBatchAsync(t *testing.T) {
	idBatch := NewIdBatch(1000, 100, 6)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := idBatch.GoExec(func(roundIndex int) error {
		func(ctx context.Context) {
			fmt.Println("[DOING TASK]", "roundIndex", roundIndex)
		}(ctx)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
