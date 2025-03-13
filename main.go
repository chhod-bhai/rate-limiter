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
	// API handler for rate limiter
	rateLimitAlgo := algorithms.New(model.TokenBucket)
	resp, err := rateLimitAlgo.IsVaidRequest()
	if err != nil {
		println(fmt.Sprintf("ERROR: handling rate limit %s", err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(fmt.Appendf(nil, "ERROR: during rate limit middleware: %s", err))
		return
	}
	if resp.LimitExpired {
		w.WriteHeader(http.StatusTooManyRequests)
		println("INFO: Request limit reached")
		return
	}
	b, err := json.Marshal(resp)
	if err != nil {
		println(fmt.Sprintf("ERROR: marshalling json %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func replenishLimit() {
	// Just handles token bucket algo for now
	// Need to work on algo switchability
	println("INFO: refueling bucket")
	rateLimitAlgo := algorithms.New(model.TokenBucket)
	err := rateLimitAlgo.Replenish()
	if err != nil {
		println(fmt.Sprintf("ERROR: replenishing bucket %v", err))
	}
}

func main() {
	// Replenish limit first time on application start
	replenishLimit()

	// setup crons to replenish limits periodically
	c := cron.New()
	defer c.Stop()
	c.AddFunc("@every 1m", replenishLimit)
	c.Start()

	// API endpoint to limit for now -> for testing only
	http.HandleFunc("/", rootHandler)

	// REPL
	println(fmt.Sprintf("INFO: Listening on port %d", 8080))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		println(fmt.Sprintf("ERROR: starting server %v", err))
		return
	}

}
