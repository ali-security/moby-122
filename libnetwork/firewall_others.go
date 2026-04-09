//go:build !linux

package libnetwork

func setupPlatformFirewall(c *Controller) {}
func arrangeUserFilterRule()              {}
func setupUserChain(ipVersion any) error  { return nil }
