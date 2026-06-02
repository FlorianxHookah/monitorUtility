package collector

import (
	"context"
	"monitorUtility/pkg/models"
	"testing"
	"time"
)

func TestCollectCPU(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan models.MetricResult, 1)

	go CollectCPU(ctx, ch)
	select {
	case result := <-ch:
		if result.Name != "CPU" {
			t.Errorf("CPU collector returned unexpected result: %s", result.Name)
		}
		if result.Err != nil {
			t.Errorf("CPU collector returned unexpected error: %s", result.Err)
		}
		if result.Value < 0 {
			t.Errorf("CPU collector returned unexpected result: %.1f", result.Value)
		}
		t.Logf("CPU collector succeeded.")
	case <-ctx.Done():
		t.Fatal("timed out waiting for CPU collector")
	}
}
