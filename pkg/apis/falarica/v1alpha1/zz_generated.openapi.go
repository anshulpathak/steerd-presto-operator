// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.AutoscalingSpec": schema_pkg_apis_falarica_v1alpha1_AutoscalingSpec(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogList":     schema_pkg_apis_falarica_v1alpha1_CatalogList(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogSecret":   schema_pkg_apis_falarica_v1alpha1_CatalogSecret(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogSpec":     schema_pkg_apis_falarica_v1alpha1_CatalogSpec(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CoordinatorSpec": schema_pkg_apis_falarica_v1alpha1_CoordinatorSpec(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.HMSSpec":         schema_pkg_apis_falarica_v1alpha1_HMSSpec(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.ImageSpec":       schema_pkg_apis_falarica_v1alpha1_ImageSpec(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.Presto":          schema_pkg_apis_falarica_v1alpha1_Presto(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.PrestoSpec":      schema_pkg_apis_falarica_v1alpha1_PrestoSpec(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.PrestoStatus":    schema_pkg_apis_falarica_v1alpha1_PrestoStatus(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.ServiceSpec":     schema_pkg_apis_falarica_v1alpha1_ServiceSpec(ref),
		"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.WorkerSpec":      schema_pkg_apis_falarica_v1alpha1_WorkerSpec(ref),
	}
}

func schema_pkg_apis_falarica_v1alpha1_AutoscalingSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"minReplicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"maxReplicas": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"targetCPUUtilizationPercentage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_falarica_v1alpha1_CatalogList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"catalogSecrets": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret names in the same namespace",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogSecret"),
									},
								},
							},
						},
					},
					"catalogSpec": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogSpec"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogSecret", "github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogSpec"},
	}
}

func schema_pkg_apis_falarica_v1alpha1_CatalogSecret(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "one has to create catalogs as secret using the following command in order to use this kubectl create secret generic secretName --from-literal=secretKey1='connector.name=tpch' --from-literal=secretKey2='connector.name=tpcds'",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"secretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secretKey": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_falarica_v1alpha1_CatalogSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"content": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"name", "content"},
			},
		},
	}
}

func schema_pkg_apis_falarica_v1alpha1_CoordinatorSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"memoryLimit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"additionalJVMConfig": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"additionalProps": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"cpuLimit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"cpuRequest": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"httpsEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"httpsKeyPairSecretName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"httpsKeyPairSecretKey": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"httpsKeyPairPassword": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"memoryLimit", "cpuLimit"},
			},
		},
	}
}

func schema_pkg_apis_falarica_v1alpha1_HMSSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_falarica_v1alpha1_ImageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"prestoPath": {
						SchemaProps: spec.SchemaProps{
							Description: "\n\t+kubebuilder:validation:Optional",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "prestoPath"},
			},
		},
	}
}

func schema_pkg_apis_falarica_v1alpha1_Presto(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Presto is the Schema for the prestos API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.PrestoSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.PrestoStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.PrestoSpec", "github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.PrestoStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_falarica_v1alpha1_PrestoSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PrestoSpec defines the desired state of Presto",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"coordinator": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html\n\t+kubebuilder:validation:Required",
							Ref:         ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CoordinatorSpec"),
						},
					},
					"worker": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.WorkerSpec"),
						},
					},
					"catalogs": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogList"),
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.ServiceSpec"),
						},
					},
					"internalHiveMetaStore": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.HMSSpec"),
						},
					},
					"imageDetails": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.ImageSpec"),
						},
					},
					"additionalPrestoPropFiles": {
						SchemaProps: spec.SchemaProps{
							Description: "additionalPrestoPropFiles:\n  access-control.properties: |\n   access-control.name=read-only\n event-listener.properties: |\n   event-listener.name=event-logger\n   jdbc.url=jdbc:postgresql://example.com:5432/eventlog\n   jdbc.user=myuser\n   jdbc.password=mypassword",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"volumes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.PrestoVolumeSpec"),
									},
								},
							},
						},
					},
				},
				Required: []string{"coordinator", "worker", "internalHiveMetaStore", "imageDetails", "additionalPrestoPropFiles"},
			},
		},
		Dependencies: []string{
			"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CatalogList", "github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.CoordinatorSpec", "github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.HMSSpec", "github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.ImageSpec", "github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.PrestoVolumeSpec", "github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.ServiceSpec", "github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.WorkerSpec"},
	}
}

func schema_pkg_apis_falarica_v1alpha1_PrestoStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PrestoStatus defines the observed state of Presto",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"uuid": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"desiredWorkers": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"currentWorkers": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"headlessService": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"coordinatorAddress": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"catalogConfig": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"coordinatorConfig": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"workerConfig": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"workerReplicaset": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"coordinatorReplicaset": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"hpaName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"clusterState": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"errorReason": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"modificationTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"coordinatorCPU": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"workerCPU": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"uuid", "desiredWorkers", "currentWorkers", "headlessService", "service", "coordinatorAddress", "catalogConfig", "coordinatorConfig", "workerConfig", "workerReplicaset", "coordinatorReplicaset", "hpaName", "clusterState", "errorReason"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_falarica_v1alpha1_ServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServiceSpec describes the attributes that a user creates on a service. Following is a copy of v1.ServiceSpec except that Ports is an optional field and Selectors field is removed.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodePort": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"clusterIP": {
						SchemaProps: spec.SchemaProps{
							Description: "clusterIP is the IP address of the service and is usually assigned randomly by the master. If an address is specified manually and is not in use by others, it will be allocated to the service; otherwise, creation of the service will fail. This field can not be changed through updates. Valid values are \"None\", empty string (\"\"), or a valid IP address. \"None\" can be specified for headless services when proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. \"ExternalName\" maps to the specified externalName. \"ClusterIP\" allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object. If clusterIP is \"None\", no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a stable IP. \"NodePort\" builds on ClusterIP and allocates a port on every node which routes to the clusterIP. \"LoadBalancer\" builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the clusterIP. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"externalIPs": {
						SchemaProps: spec.SchemaProps{
							Description: "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service.  These IPs are not managed by Kubernetes.  The user is responsible for ensuring that traffic arrives at a node with this IP.  A common example is external load-balancers that are not part of the Kubernetes system.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"sessionAffinity": {
						SchemaProps: spec.SchemaProps{
							Description: "Supports \"ClientIP\" and \"None\". Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"loadBalancerIP": {
						SchemaProps: spec.SchemaProps{
							Description: "Only applies to Service Type: LoadBalancer LoadBalancer will get created with the IP specified in this field. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"loadBalancerSourceRanges": {
						SchemaProps: spec.SchemaProps{
							Description: "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.\" More info: https://kubernetes.io/docs/tasks/access-application-cluster/configure-cloud-provider-firewall/",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"externalName": {
						SchemaProps: spec.SchemaProps{
							Description: "externalName is the external reference that kubedns or equivalent will return as a CNAME record for this service. No proxying will be involved. Must be a valid RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires Type to be ExternalName.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"externalTrafficPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints. \"Local\" preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. \"Cluster\" obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"healthCheckNodePort": {
						SchemaProps: spec.SchemaProps{
							Description: "healthCheckNodePort specifies the healthcheck nodePort for the service. If not specified, HealthCheckNodePort is created by the service api backend with the allocated nodePort. Will use user-specified nodePort value if specified by the client. Only effects when Type is set to LoadBalancer and ExternalTrafficPolicy is set to Local.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"publishNotReadyAddresses": {
						SchemaProps: spec.SchemaProps{
							Description: "publishNotReadyAddresses, when set to true, indicates that DNS implementations must publish the notReadyAddresses of subsets for the Endpoints associated with the Service. The default value is false. The primary use case for setting this field is to use a StatefulSet's Headless Service to propagate SRV records for its Pods without respect to their readiness for purpose of peer discovery.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"sessionAffinityConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "sessionAffinityConfig contains the configurations of session affinity.",
							Ref:         ref("k8s.io/api/core/v1.SessionAffinityConfig"),
						},
					},
					"ipFamily": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SessionAffinityConfig"},
	}
}

func schema_pkg_apis_falarica_v1alpha1_WorkerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"memoryLimit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"additionalJVMConfig": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"additionalProps": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"cpuLimit": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"cpuRequest": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"terminationGracePeriodSeconds": {
						SchemaProps: spec.SchemaProps{
							Description: "Optional duration in seconds the pod needs to terminate gracefully. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 7200 seconds.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"count": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"autoscaling": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.AutoscalingSpec"),
						},
					},
				},
				Required: []string{"memoryLimit", "cpuLimit", "count"},
			},
		},
		Dependencies: []string{
			"github.com/falarica/steerd-presto-operator/pkg/apis/falarica/v1alpha1.AutoscalingSpec"},
	}
}