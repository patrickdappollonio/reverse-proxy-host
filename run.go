package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/spf13/cobra"
)

const loopback = "0.0.0.0"

func run() *cobra.Command {
	var hostname string
	var listenPort int
	var destinationPort int

	rootCmd := &cobra.Command{
		Use:   "reverse-proxy-host",
		Short: "A reverse proxy that sets the Host header",
		RunE: func(cmd *cobra.Command, args []string) error {
			stdout, stderr := cmd.OutOrStdout(), cmd.OutOrStderr()
			return proxy(stdout, stderr, hostname, destinationPort, listenPort)
		},
	}

	rootCmd.Flags().StringVarP(&hostname, "host", "H", "", "the hostname to used when mocking requests")
	rootCmd.Flags().IntVarP(&listenPort, "listen-port", "l", 8080, "the port to listen on")
	rootCmd.Flags().IntVarP(&destinationPort, "destination-port", "d", 443, "the port to forward to")

	rootCmd.MarkFlagRequired("host")
	rootCmd.MarkFlagRequired("listen-port")
	rootCmd.MarkFlagRequired("destination-port")

	return rootCmd
}

func proxy(stdout, stderr io.Writer, hostname string, destinationPort, listenPort int) error {
	proxy := &httputil.ReverseProxy{
		Transport: roundTripper(rt),
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = generateHostPort(loopback, destinationPort)
			req.Host = hostname
			req.Header.Set("Host", hostname)
		},
	}

	log.Printf("Starting reverse proxy to \"%s:%d\", listening for requests on port %d", loopback, destinationPort, listenPort)
	if err := http.ListenAndServe(generateHostPort(loopback, listenPort), proxy); err != nil {
		return err
	}

	return nil
}

func generateHostPort(hostname string, port int) string {
	switch port {
	case 80, 443:
		return hostname
	default:
		return fmt.Sprintf("%s:%d", hostname, port)
	}
}

func rt(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	log.Printf("%s %s (%s)", req.Method, req.URL.String(), resp.Status)

	return resp, nil
}
