//go:build dev

package provider

import "os"

// isDemoMode returns true if demo credentials are configured.
// This function is only available in dev builds (go build -tags dev).
func isDemoMode() bool {
	return os.Getenv("SLACK_MCP_XOXP_TOKEN") == "demo" ||
		(os.Getenv("SLACK_MCP_XOXC_TOKEN") == "demo" && os.Getenv("SLACK_MCP_XOXD_TOKEN") == "demo")
}
