package kube

parent_namespace: [namespace_name=_]: {
	apiVersion: "v1"
	kind:       "Namespace"
	metadata: {
		name: namespace_name
	}
}

subnamespace: [subnamespace_name=_]: {
	apiVersion: "hnc.x-k8s.io/v1alpha2"
	kind:       "SubnamespaceAnchor"
	metadata: {
		namespace: string
		name:      subnamespace_name
	}
	spec: {}
}
