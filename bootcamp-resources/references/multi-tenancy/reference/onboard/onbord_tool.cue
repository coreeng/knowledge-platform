package kube

import (
	"encoding/yaml"
	"tool/exec"
	"tool/cli"
)

manifestObjects: [ for v in manifestObjectSets for x in v {x}]

manifestObjectSets: [
	parent_namespace,
	subnamespace,
	service_account,
	read_write_role,
	role_binding,
]

networkPolicyObjects: [ for v in networkPolicyObjectSets for x in v {x}]

networkPolicyObjectSets: [
	default_deny_policy
]

allObjects:  manifestObjects + networkPolicyObjects

command: onboard: {
	task: namespaces: exec.Run & {
		cmd:    "kubectl apply -f -"
		stdin:  yaml.MarshalStream(manifestObjects)
		stdout: string
	}

	task: displayOutputNamespaces: cli.Print & {
		text: task.namespaces.stdout
	}

	task: netwokPolicies: exec.Run & {
		cmd:    "kubectl apply -f -"
		stdin:  yaml.MarshalStream(networkPolicyObjects)
		$dep: task.namespaces.$done //  explicit dependency on the namespaces task
		stdout: string
	}

	task: displayOutputNetworkPolicies: cli.Print & {
		text: task.netwokPolicies.stdout
	}

}

command: monitoring: {
	task: kube: exec.Run & {
		cmd:    "./onboard/monitoring/install.sh"
		stdout: string
	}

	task: display: cli.Print & {
		text: task.kube.stdout
	}
}

command: print: {
	task: print: cli.Print & {
		text: yaml.MarshalStream(allObjects)
	}
}
