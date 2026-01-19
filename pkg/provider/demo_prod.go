//go:build !dev

package provider

// isDemoMode always returns false in production builds.
// Demo mode is only available when building with -tags dev.
func isDemoMode() bool {
	return false
}
