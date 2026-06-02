package collector

import (
	"context"
	"monitorUtility/pkg/models"
	"testing"
	"time"
)

func TestCollectRAM(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan models.MetricResult, 1)

	go CollectRAM(ctx, ch)
	select {
	case result := <-ch:
		if result.Name != "RAM" {
			t.Errorf("RAM collector returned unexpected result: %s", result.Name)
		}
		if result.Err != nil {
			t.Errorf("RAM collector returned unexpected error: %s", result.Err)
		}
		if result.Value < 0 {
			t.Errorf("RAM collector returned unexpected result: %.1f", result.Value)
		}
		t.Logf("RAM collector succeeded.")
	case <-ctx.Done():
		t.Fatal("timed out waiting for RAM collector")
	}
}
