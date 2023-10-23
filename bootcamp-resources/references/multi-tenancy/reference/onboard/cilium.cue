package kube

default_deny_policy: [namespace_name=string]: {
		apiVersion: "cilium.io/v2"
		kind:       "CiliumNetworkPolicy"
		metadata: {
			name:      "ingress-default-deny"
			namespace: namespace_name
		}
		spec: {
			endpointSelector: {}
			ingress: [{
				fromEndpoints: [ {}]
			}]
		}
}