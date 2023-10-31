package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

const (
	healthy   = true
	unhealthy = false
)

var newHealthCommand = func() *cobra.Command {
	return &cobra.Command{
		Use:   "health",
		Short: "Calls the health endpoint of a Sablier instance",
		Run: func(cmd *cobra.Command, args []string) {
			details, healthy := Health(cmd.Flag("url").Value.String())

			if healthy {
				fmt.Fprintf(os.Stderr, "healthy: %v\n", details)
				os.Exit(0)
			} else {
				fmt.Fprintf(os.Stderr, "unhealthy: %v\n", details)
				os.Exit(1)
			}
		},
	}
}

func Health(url string) (string, bool) {
	resp, err := http.Get(url)

	if err != nil {
		return err.Error(), unhealthy
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err.Error(), unhealthy
	}

	if resp.StatusCode >= 400 {
		return string(body), unhealthy
	}

	return string(body), healthy
}
