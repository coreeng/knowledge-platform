package kube

service_account: [sa_name=_]: {
	apiVersion: "v1"
	kind:       "ServiceAccount"
	metadata: {
		name:      sa_name
		namespace: string
	}
}

read_write_role: [role_name=_]: {
	apiVersion: "rbac.authorization.k8s.io/v1"
	kind:       "Role"
	metadata: {
		name:      role_name
		namespace: string
	}
	rules: [
		{
				apiGroups: [""]
				resources: ["*"]
				verbs: [
					"get",
					"list",
					"watch",
					"create",
					"update",
					"patch",
					"delete",
				]
		},
		{
				apiGroups: ["cilium.io"]
				resources: ["ciliumnetworkpolicies"]
				verbs: [
					"get",
					"list",
					"watch",
					"create",
					"update",
					"delete",
				]
		},
	]
}

role_binding: [role_binding_name=_]: {
	apiVersion: "rbac.authorization.k8s.io/v1"
	kind:       "RoleBinding"
	metadata: {
		name:      role_binding_name
		namespace: string
	}
	roleRef: {
		apiGroup: "rbac.authorization.k8s.io"
		kind:     "Role"
		name:     string
	}
	subjects: [{
		name:      string
		namespace: string
		kind:      "ServiceAccount"
	}]
}
