//go:build !dev

package transport

// isInsecureTLSAllowed always returns false in production builds.
// TLS certificate verification cannot be skipped in production.
// This setting is only available when building with -tags dev.
func isInsecureTLSAllowed() bool {
	return false
}
