package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
	"github.com/nicholasjackson/env"
)

var version = "dev"

type imageKey struct{}
type postResponse struct{}

var baseURI = env.String("BASE_URI", false, "http://localhost:18080", "base URI for requests")
var timeout = env.Duration("TIMEOUT", false, 60*time.Second, "timeout for each scenario")
var duration = env.Duration("DURATION", false, 5*time.Second, "test duration")
var users = env.Int("USERS", false, 5, "concurrent users")
var showProgress = env.Bool("SHOW_PROGRESS", false, true, "show graphical progress")

func main() {
	// Parse the config env vars
	err := env.Parse()
	if err != nil {
		fmt.Println(env.Help())
		os.Exit(1)
	}

	fmt.Println("Benchmarking application")

	b := bench.New(*showProgress, *users, *duration, 0*time.Second, *timeout)
	b.AddOutput(0*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.AddOutput(0*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)

	b.RunBenchmarks(HashiCupsFlow)
}

// HashiCupsFlow defines a typical user flow for the HashiCups application
func HashiCupsFlow() error {

	graphQLQuery := `{"operationName":null,"variables":{},"query":"{\n  coffees {\n    id\n    name\n  }\n}\n"}`

	resp, err := http.Post(fmt.Sprintf("%s/api", *baseURI), "application/json", bytes.NewBufferString(graphQLQuery))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected response code 200 got %d", resp.StatusCode)
	}

	return nil
}
