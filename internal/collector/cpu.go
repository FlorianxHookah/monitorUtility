package collector

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"monitorUtility/pkg/models"
)

func CollectCPU(ctx context.Context, ch chan<- models.MetricResult) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			percentage, err := cpu.Percent(0, false)
			if err != nil {
				ch <- models.MetricResult{Name: "CPU", Err: err}
				continue
			}
			if len(percentage) > 0 {
				ch <- models.MetricResult{Name: "CPU", Value: percentage[0]}
			}
		}
	}
}
