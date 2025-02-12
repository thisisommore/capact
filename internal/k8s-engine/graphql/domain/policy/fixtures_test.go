package policy_test

import (
	"capact.io/capact/internal/ptr"
	"capact.io/capact/pkg/engine/api/graphql"
	"capact.io/capact/pkg/engine/k8s/policy"
	"capact.io/capact/pkg/sdk/apis/0.0.1/types"
)

func fixGQLInput() graphql.PolicyInput {
	return graphql.PolicyInput{
		Rules: []*graphql.RulesForInterfaceInput{
			{
				Interface: &graphql.ManifestReferenceInput{
					Path:     "cap.interface.database.postgresql.install",
					Revision: ptr.String("0.1.0"),
				},
				OneOf: []*graphql.PolicyRuleInput{
					{
						ImplementationConstraints: &graphql.PolicyRuleImplementationConstraintsInput{
							Requires: []*graphql.ManifestReferenceInput{
								{
									Path:     "cap.type.gcp.auth.service-account",
									Revision: ptr.String("0.1.0"),
								},
							},
							Attributes: []*graphql.ManifestReferenceInput{
								{
									Path: "cap.attribute.cloud.provider.gcp",
								},
							},
						},
						Inject: &graphql.PolicyRuleInjectDataInput{
							RequiredTypeInstances: []*graphql.RequiredTypeInstanceReferenceInput{
								{
									ID:          "c268d3f5-8834-434b-bea2-b677793611c5",
									Description: ptr.String("Sample description"),
								},
							},
							AdditionalParameters: []*graphql.AdditionalParameterInput{
								{
									Name: "additional-parameters",
									Value: map[string]interface{}{
										"key1": "boom",
									},
								},
							},
							AdditionalTypeInstances: []*graphql.AdditionalTypeInstanceReferenceInput{
								{
									Name: "sample",
									ID:   "0b6dba9a-d111-419d-b236-357cf0e8603a",
								},
							},
						},
					},
					{
						ImplementationConstraints: &graphql.PolicyRuleImplementationConstraintsInput{
							Path: ptr.String("cap.implementation.bitnami.postgresql.install"),
						},
					},
				},
			},
			{
				Interface: &graphql.ManifestReferenceInput{
					Path: "cap.*",
				},
				OneOf: []*graphql.PolicyRuleInput{
					{
						ImplementationConstraints: &graphql.PolicyRuleImplementationConstraintsInput{},
					},
				},
			},
		},
	}
}

func fixGQL() graphql.Policy {
	return graphql.Policy{
		Rules: []*graphql.RulesForInterface{
			{
				Interface: &graphql.ManifestReferenceWithOptionalRevision{
					Path:     "cap.interface.database.postgresql.install",
					Revision: ptr.String("0.1.0"),
				},
				OneOf: []*graphql.PolicyRule{
					{
						ImplementationConstraints: &graphql.PolicyRuleImplementationConstraints{
							Requires: []*graphql.ManifestReferenceWithOptionalRevision{
								{
									Path:     "cap.type.gcp.auth.service-account",
									Revision: ptr.String("0.1.0"),
								},
							},
							Attributes: []*graphql.ManifestReferenceWithOptionalRevision{
								{
									Path: "cap.attribute.cloud.provider.gcp",
								},
							},
						},
						Inject: &graphql.PolicyRuleInjectData{
							RequiredTypeInstances: []*graphql.RequiredTypeInstanceReference{
								{
									ID:          "c268d3f5-8834-434b-bea2-b677793611c5",
									Description: ptr.String("Sample description"),
								},
							},
							AdditionalParameters: []*graphql.AdditionalParameter{
								{
									Name: "additional-parameters",
									Value: map[string]interface{}{
										"key1": "boom",
									},
								},
							},
							AdditionalTypeInstances: []*graphql.AdditionalTypeInstanceReference{
								{
									Name: "sample",
									ID:   "0b6dba9a-d111-419d-b236-357cf0e8603a",
								},
							},
						},
					},
					{
						ImplementationConstraints: &graphql.PolicyRuleImplementationConstraints{
							Path: ptr.String("cap.implementation.bitnami.postgresql.install"),
						},
					},
				},
			},
			{
				Interface: &graphql.ManifestReferenceWithOptionalRevision{
					Path: "cap.*",
				},
				OneOf: []*graphql.PolicyRule{
					{
						ImplementationConstraints: &graphql.PolicyRuleImplementationConstraints{},
					},
				},
			},
		},
	}
}

func fixModel() policy.Policy {
	return policy.Policy{
		Rules: policy.RulesList{
			{
				Interface: types.ManifestRefWithOptRevision{
					Path:     "cap.interface.database.postgresql.install",
					Revision: ptr.String("0.1.0"),
				},
				OneOf: []policy.Rule{
					{
						ImplementationConstraints: policy.ImplementationConstraints{
							Requires: &[]types.ManifestRefWithOptRevision{
								{
									Path:     "cap.type.gcp.auth.service-account",
									Revision: ptr.String("0.1.0"),
								},
							},
							Attributes: &[]types.ManifestRefWithOptRevision{
								{
									Path: "cap.attribute.cloud.provider.gcp",
								},
							},
						},
						Inject: &policy.InjectData{
							RequiredTypeInstances: []policy.RequiredTypeInstanceToInject{
								{
									RequiredTypeInstanceReference: policy.RequiredTypeInstanceReference{
										ID:          "c268d3f5-8834-434b-bea2-b677793611c5",
										Description: ptr.String("Sample description"),
									},
								},
							},
							AdditionalParameters: []policy.AdditionalParametersToInject{
								{
									Name: "additional-parameters",
									Value: map[string]interface{}{
										"key1": "boom",
									},
								},
							},
							AdditionalTypeInstances: []policy.AdditionalTypeInstanceToInject{
								{
									AdditionalTypeInstanceReference: policy.AdditionalTypeInstanceReference{
										Name: "sample",
										ID:   "0b6dba9a-d111-419d-b236-357cf0e8603a",
									},
								},
							},
						},
					},
					{
						ImplementationConstraints: policy.ImplementationConstraints{
							Path: ptr.String("cap.implementation.bitnami.postgresql.install"),
						},
					},
				},
			},
			{
				Interface: types.ManifestRefWithOptRevision{
					Path: "cap.*",
				},
				OneOf: []policy.Rule{
					{
						ImplementationConstraints: policy.ImplementationConstraints{},
					},
				},
			},
		},
	}
}
