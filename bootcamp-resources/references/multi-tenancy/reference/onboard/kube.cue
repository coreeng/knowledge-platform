package kube

for orgConfig in orgs {

	// create namespace
	parent_namespace: "\(orgConfig.name)": {
		metadata: {
			name: "\(orgConfig.name)"
		}
	}

	default_deny_policy: "\(orgConfig.name)": {}

	for tenantConfig in orgConfig.tenants {

		// create a subnamespace for each onboard in the org namespace
		subnamespace: "\(tenantConfig.name)": {
			metadata: {
				namespace: "\(orgConfig.name)"
				name:      "\(tenantConfig.name)"
			}
			spec: labels: tenantConfig.labels
		}

		// for tenants that opt in for monitoring, provide a monitoring namespace with the monitoring stack
		if tenantConfig.monitoring {
			// create a monitoring namespace for each onboard that requires it
			subnamespace: "\(tenantConfig.name)-monitoring": {
				metadata: {
					namespace: "\(tenantConfig.name)"
					name:      "\(tenantConfig.name)-monitoring"
				}
				spec: labels: tenantConfig.labels
			}
		}

		for appName in tenantConfig.apps {

			// create a subnamespace for every app in the configuration
			subnamespace: "\(appName)": {
				metadata: {
					namespace: "\(tenantConfig.name)"
					name:      "\(appName)"
				}
				spec: labels: tenantConfig.labels
			}
		}

		// create a service account for each onboard
		service_account: "\(tenantConfig.name)-sa": {
			metadata: {
				name:      "\(tenantConfig.name)-sa"
				namespace: "\(tenantConfig.name)"
			}
		}

		// create read-write role
		read_write_role: "\(tenantConfig.name)-read-write-role": {
			apiVersion: "rbac.authorization.k8s.io/v1"
			kind:       "Role"
			metadata: {
				name:      "\(tenantConfig.name)-read-write-role"
				namespace: "\(tenantConfig.name)"
			}
		}

		// create role binding for each onboard
		role_binding: "\(tenantConfig.name)-rb": {
			metadata: {
				name:      "\(tenantConfig.name)-rb"
				namespace: "\(tenantConfig.name)"
			}
			roleRef: {
				name: "\(tenantConfig.name)-read-write-role"
			}
			subjects: [{
				name:      "\(tenantConfig.name)-sa"
				namespace: "\(tenantConfig.name)"
			}]
		}
	}
}
