//go:build dev

package transport

import "os"

// isInsecureTLSAllowed returns true if TLS certificate verification should be skipped.
// This function is only available in dev builds (go build -tags dev).
// In production builds, TLS certificate verification is always required.
func isInsecureTLSAllowed() bool {
	return os.Getenv("SLACK_MCP_SERVER_CA_INSECURE") != ""
}
