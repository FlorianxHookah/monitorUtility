package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"monitorUtility/internal/collector"
	"monitorUtility/pkg/models"
)

func main() {
	fmt.Println("Starting CPU utility (Ctrl+C to exit...)")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	metrics := make(chan models.MetricResult, 10)

	go collector.CollectCPU(ctx, metrics)
	go collector.CollectRAM(ctx, metrics)
	go collector.CollectDisk(ctx, metrics)

loop:
	for {
		select {
		case result := <-metrics:
			if result.Err != nil {
				fmt.Println("Error:", result.Err)
			} else {
				fmt.Printf("%s load %.1f%%\n", result.Name, result.Value)
			}
		case <-ctx.Done():
			fmt.Println("Exiting...")
			time.Sleep(150 * time.Millisecond)
			break loop
		}
	}
	fmt.Println("Program executed...")
}
