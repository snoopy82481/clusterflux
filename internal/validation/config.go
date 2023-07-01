package validation

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	"github.com/gookit/validate"
	"github.com/snoopy82481/clusterflux/internal/config"
	"github.com/snoopy82481/clusterflux/internal/logger"
)

func ValidateConfig(config *config.Config) error {
	v := validate.Struct(config)

	// custom validation
	v.AddValidator("loadBalancerRange", validateLoadBalancerRange)
	v.AddValidator("ansibleIPNotInLoadBalancerRange", validateIPNotInLoadBalancerRange)

	if !v.Validate() {
		for _, err := range v.Errors {
			logger.LogWarn(err.String())
		}

		return fmt.Errorf("Config validation failed")
	}

	return nil
}

func validateLoadBalancerRange(field, val string) (string, bool) {
	_, _, err := net.ParseCIDR(val)
	if err == nil || isValidIPRange(val) {
		return "", true
	}

	return fmt.Sprintf("%s must be a valid IP range or CIDR", field), false
}

func isValidIPRange(val string) bool {
	parts := strings.Split(val, "=")
	if len(parts) == 2 {
		ip1 := net.ParseIP(strings.TrimSpace(parts[0]))
		ip2 := net.ParseIP(strings.TrimSpace(parts[1]))

		if ip1 != nil && ip2 != nil {
			return true
		}
	}

	return false
}

func validateIPNotInLoadBalancerRange(field, val string, config *config.Config) (string, bool) {
	for _, host := range config.Ansible.Hosts {
		ip := net.ParseIP(host.IPAddress)
		if ip != nil {
			// Check if IP falls within the CIDR range
			_, cidr, err := net.ParseCIDR(config.Network.LoadBalancerRange)
			if err != nil {
				// If parsing as CIDR fails, check as IP range
				startIP, endIP, err := parseIPRange(config.Network.LoadBalancerRange)
				if err != nil {
					return fmt.Sprintf("Invalid loadBalancerRange: %s", config.Network.LoadBalancerRange), false
				}
				if isIPInRange(ip, startIP, endIP) {
					return fmt.Sprintf("%s IP address is in the loadBalancerRange", field), false
				}
			} else {
				if cidr.Contains(ip) {
					return fmt.Sprintf("%s IP address is in the loadBalancerRange", field), false
				}
			}
		}
	}
	return "", true
}

// Helper function to parse IP range
func parseIPRange(ipRange string) (net.IP, net.IP, error) {
	// Split the IP range string into start and end IP
	ipRangeParts := strings.Split(ipRange, "-")
	if len(ipRangeParts) != 2 {
		return nil, nil, fmt.Errorf("invalid IP range format")
	}
	startIP := net.ParseIP(strings.TrimSpace(ipRangeParts[0]))
	endIP := net.ParseIP(strings.TrimSpace(ipRangeParts[1]))
	if startIP == nil || endIP == nil {
		return nil, nil, fmt.Errorf("invalid IP range format")
	}
	return startIP, endIP, nil
}

// Helper function to check if an IP is within a range
func isIPInRange(ip, startIP, endIP net.IP) bool {
	return bytes.Compare(ip, startIP) >= 0 && bytes.Compare(ip, endIP) <= 0
}
