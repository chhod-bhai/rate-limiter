package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chhod-bhai/rate-limiter/algorithms"
	"github.com/chhod-bhai/rate-limiter/model"
	"github.com/robfig/cron/v3"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rateLimitAlgo := algorithms.New(model.TokenBucket)
	resp, err := rateLimitAlgo.IsVaidRequest()
	if err != nil {
		println(fmt.Sprintf("Error handling rate limit %v", err))
		w.Write(fmt.Appendf(nil, "ERROR: Error during rate limit middleware: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	if resp.RemainingLimit == 0 {
		w.WriteHeader(http.StatusTooManyRequests)
		println("INFO: Request limit reached")
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		println(fmt.Sprintf("Error marshalling json %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)

}

func main() {
	http.HandleFunc("/", rootHandler)

	println(fmt.Sprintf("Listening on port %d", 8080))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		println(fmt.Sprintf("Error starting server %v", err))
	}

	c := cron.New()
	defer c.Stop()
	c.AddFunc("@every 1m", func() {
		rateLimitAlgo := algorithms.New(model.TokenBucket)
		err := rateLimitAlgo.Replenish()
		if err != nil {
			println(fmt.Sprintf("ERROR: Error replenishing bucket %v", err))
		}
	})
	c.Start()
}
