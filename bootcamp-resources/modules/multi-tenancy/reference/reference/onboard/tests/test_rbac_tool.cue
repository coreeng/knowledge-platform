package kube

import (
	"tool/exec"
	"tool/cli"
)


command: testRbac: {
	for orgConfig in orgs {
		for tenantConfig in orgConfig.tenants {

			task: "\(tenantConfig.name)": {
				kube: exec.Run & {
					cmd:    "./onboard/tests/test_rbac.sh " + "\(tenantConfig.name):\(tenantConfig.name)-sa"
					stdout: string
				}
				display: cli.Print & {
					text: kube.stdout
				}
			}
		}

	}
}

command: testNetworkIsolation: {
	task: test: {
		kube: exec.Run & {
			cmd:    "./onboard/tests/test_network_isolation.sh"
			stdout: string
		}
		display: cli.Print & {
			text: kube.stdout
		}
	}
}
