package validation

import (
	"fmt"
	"os/exec"

	"github.com/snoopy82481/clusterflux/internal/config"
	"github.com/snoopy82481/clusterflux/internal/logger"
)

func ValidateAnsibleHosts(config *config.Config) error {
	var sshError bool

	controlNodeCount := countControlNodes(config.Ansible.Hosts)

	if controlNodeCount == 0 || controlNodeCount%2 != 1 {
		return fmt.Errorf("You must have 1, 3 or an odd number of master nodes >=3 in order for etcd to have quorum")
	}

	for _, host := range config.Ansible.Hosts {
		if host.IPAddress == config.Network.KubeVIPAddr {
			errMsg := fmt.Sprintf("The kube-vip IP '%s' should not be the same as the IP for node '%s'", config.Network.KubeVIPAddr, host.IPAddress)
			return fmt.Errorf(errMsg)
		}

		if host.IPAddress == config.Network.IngressAddr {
			errMsg := fmt.Sprintf("The ingress load balancer IP '%s' should not be the same as the IP for node '%s'", config.Network.IngressAddr, host.IPAddress)
			return fmt.Errorf(errMsg)
		}

		if host.IPAddress == config.Network.K8sGatewayAddr {
			errMsg := fmt.Sprintf("The k8s-gateway IP '%s' should not be the same as the IP for node '%s'", config.Network.KubeVIPAddr, host.IPAddress)
			return fmt.Errorf(errMsg)
		}

		cmd := exec.Command("ssh", "-q", "-o", "BatchMode=yes", "-o", "ConnectTimeout=5", host.SSHUsername+"@"+host.IPAddress, "true")
		if err := cmd.Run(); err != nil {
			errMsg := fmt.Sprintf("SSH into host '%s' with username '%s' was NOT successful, did you copy over your SSH key?", host.IPAddress, host.SSHUsername)
			logger.LogWarn(errMsg)
			sshError = true
		}
	}

	if sshError {
		return fmt.Errorf("An error occurred during SSH validation.")
	}

	return nil
}

func countControlNodes(hosts []config.AnsibleHost) int {
	controlNodeCount := 0
	for _, host := range hosts {
		if host.ControlNode {
			controlNodeCount++
		}
	}
	return controlNodeCount
}
