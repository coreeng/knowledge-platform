package kube

// onboard configuration per organisation
orgs: [
	{
		name: "cecg"
		tenants: [
			{
				name: "team-a"
				labels: [{
					key:   "custom.org/costCode"
					value: "12345"
				}]
				apps: [
					"app-1",
					"app-2",
				]
				monitoring: true
			},
					{
						name: "team-b"
						labels: [{
							key:   "custom.org/costCode"
							value: "5678"
						}]
						apps: [
							"app-3",
						]
						monitoring: true
					},

		]
	},
]
