package collector

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
	"monitorUtility/pkg/models"
)

func CollectDisk(ctx context.Context, ch chan<- models.MetricResult) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			usage, err := disk.Usage("/")
			if err != nil {
				ch <- models.MetricResult{Name: "Disk", Err: err}
				continue
			}
			ch <- models.MetricResult{Name: "Disk", Value: usage.UsedPercent}
		}
	}
}
