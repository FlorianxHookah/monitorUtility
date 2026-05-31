package collector

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
	"monitorUtility/pkg/models"
)

func CollectRAM(ctx context.Context, ch chan<- models.MetricResult) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			value, err := mem.VirtualMemory()
			if err != nil {
				ch <- models.MetricResult{Name: "RAM", Err: err}
				continue
			}
			ch <- models.MetricResult{Name: "RAM", Value: value.UsedPercent}
		}
	}
}
