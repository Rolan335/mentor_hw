package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var counter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "counter",
	Help: "how many times route /add was used",
})

var gauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "Goroutines",
	Help: "how many gorouitines are working",
})

func metricGoroutines() {
	go func() {
		for {
			gauge.Set(float64(runtime.NumGoroutine()))
			time.Sleep(time.Second * 10)
		}
	}()
}

func main() {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		counter.Inc()
		w.Write([]byte("info"))
	})
	http.Handle("/metrics", promhttp.Handler())
	metricGoroutines()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
