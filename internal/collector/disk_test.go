package collector

import (
	"context"
	"monitorUtility/pkg/models"
	"testing"
	"time"
)

func TestCollectDisk(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ch := make(chan models.MetricResult, 1)

	go CollectDisk(ctx, ch)
	select {
	case result := <-ch:
		if result.Name != "Disk" {
			t.Errorf("Disk collector returned unexpected result: %s", result.Name)
		}
		if result.Err != nil {
			t.Errorf("Disk collector returned unexpected error: %s", result.Err)
		}
		if result.Value < 0 {
			t.Errorf("Disk collector returned unexpected result: %.1f", result.Value)
		}
		t.Logf("Disk collector succeeded.")
	case <-ctx.Done():
		t.Fatal("timed out waiting for Disk collector")
	}
}
