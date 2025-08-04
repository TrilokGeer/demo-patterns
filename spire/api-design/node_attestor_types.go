package v1alpha1

const (
	NodeAttestorTypeK8SPSAT  = "k8s_psat"
	NodeAttestorTypeAWSIID   = "aws_iid"
	NodeAttestorTypeGCPIIT   = "gcp_iit"
	NodeAttestorTypeAzureMSI = "azure_msi"
)

type NodeAttestor struct {
	// type specifies the type of node attestor to use.
	// +kubebuilder:validation:Enum=k8s_psat;k8s_sds;aws_iam;aws_iam_oidc;aws_ec2_s3;aws_ec2_iam;aws_ec2_iam_oidc;aws_ec2_iam_oidc_s3;aws_ec2_iam_oidc_s3_kms;aws_ec2_iam_oidc_s3_kms_disk;aws_ec2_iam_oidc_s3_kms_disk_memory;aws_ec2_iam_oidc_s3_kms_disk_memory_disk;aws_ec2_iam_oidc_s3_kms_disk_memory_disk_memory;aws_ec2_iam_oidc_s3_kms_disk_memory_disk_memory_disk;aws_ec2_iam_oidc_s3_kms_disk_memory_disk_memory_disk_memory;aws_ec2_iam_oidc_s3_kms_disk_memory_disk_memory_disk_memory_disk;aws_ec2_iam_oidc_s3_kms_disk_memory_disk_memory_disk_memory_disk
	// +kubebuilder:default:=k8s_psat
	Type string `json:"type"`

	// pluginConfig has the config required for the spire server node attestor.
	// +kubebuilder:validation:Optional
	PluginConfig map[string]interface{} `json:"pluginConfig,omitempty"`
}

type NodeAttestorAWSIID struct {
	// accessKeyID is the AWS access key ID.
	// +kubebuilder:validation:Optional
	AccessKeyID string `json:"accessKeyID,omitempty"`

	// secretAccessKey is the AWS secret access key.
	// +kubebuilder:validation:Optional
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// skipBlockDevice is a flag to skip block device attestation.
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	SkipBlockDevice bool `json:"skipBlockDevice,omitempty"`

	// disableInstanceProfileSelector is a flag to disable instance profile selector.
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	DisableInstanceProfileSelector bool `json:"disableInstanceProfileSelector,omitempty"`

	// assumeRole is the AWS role to assume.
	// +kubebuilder:validation:Optional
	AssumeRole string `json:"assumeRole,omitempty"`

	// partition is the AWS partition to use.
	// +kubebuilder:validation:Optional
	Partition string `json:"partition,omitempty"`

	// verifyOrganization is the organization to verify.
	// +kubebuilder:validation:Optional
	VerifyOrganization *VerifyOrganization `json:"verifyOrganization,omitempty"`
}

type VerifyOrganization struct {
	// managementAccountId is the management account ID to verify.
	// +kubebuilder:validation:Optional
	ManagementAccountId string `json:"managementAccountId,omitempty"`

	// managementAccountRegion is the management account region to verify.
	// +kubebuilder:validation:Optional
	ManagementAccountRegion string `json:"managementAccountRegion,omitempty"`

	// managementAccountRole is the management account role to verify.
	// +kubebuilder:validation:Optional
	AssumeOrganizationRole string `json:"assumeOrganizationRole,omitempty"`

	// organizationAccountMapTTL defines the TTL to cache the list of accounts for particular time.
	// Should be >= 1 minute. Defaults to 3 minutes.
	// +kubebuilder:default:=3m
	// +kubebuilder:validation:Minimum:=1m
	// +kubebuilder:validation:Optional
	OrganizationAccountMapTTL string `json:"organizationAccountMapTTL,omitempty"`
}

type NodeAttestorAzureMSI struct {
	// tenants is the map of tenants to configure.
	// +kubebuilder:validation:Optional
	Tenants map[string]*AzureTenantConfig `json:"tenants"`

	// agentPathTemplate is the path template for the agent.
	// +kubebuilder:validation:Optional
	AgentPathTemplate string `json:"agentPathTemplate,omitempty"`
}

type AzureTenantConfig struct {
	// resourceID is the resource ID to use.
	ResourceID string `json:"resourceID"`

	// subscriptionID is the subscription ID to use.
	SubscriptionID string `json:"subscriptionID"`

	// appID is the app ID to use.
	AppID string `json:"appID"`

	// appSecret is the app secret to use.
	AppSecret string `json:"appSecret"`
}

type NodeAttestorGCPIIT struct {
	// projectIDAllowList is the list of project IDs to allow.
	// +kubebuilder:validation:Optional
	ProjectIDAllowList []string `json:"projectIDAllowList"`

	// agentPathTemplate is the path template for the agent.
	// +kubebuilder:validation:Optional
	AgentPathTemplate string `json:"agentPathTemplate"`

	// useInstanceMetadata is a flag to use instance metadata.
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	UseInstanceMetadata bool `json:"useInstanceMetadata"`

	// allowedLabelKeys is the list of allowed label keys.
	// +kubebuilder:validation:Optional
	AllowedLabelKeys []string `json:"allowedLabelKeys"`

	// allowedMetadataKeys is the list of allowed metadata keys.
	// +kubebuilder:validation:Optional
	AllowedMetadataKeys []string `json:"allowedMetadataKeys"`

	// maxMetadataValueSize is the maximum metadata value size.
	// +kubebuilder:validation:Optional
	MaxMetadataValueSize int `json:"maxMetadataValueSize"`

	// serviceAccountFile is the path to the service account file.
	// +kubebuilder:validation:Optional
	ServiceAccountFile string `json:"serviceAccountFile"`
}

type NodeAttestorK8SPSAT struct {
	// clusters is the map of clusters to configure.
	// +kubebuilder:validation:Optional
	Clusters map[string]*NodeAttenstorCluster `json:"clusters"`
}

type NodeAttenstorCluster struct {
	// allowedServiceAccounts is the list of allowed service accounts.
	// +kubebuilder:validation:Optional
	AllowedServiceAccounts map[string]bool `json:"allowedServiceAccounts"`

	// audience is the list of audiences.
	// +kubebuilder:validation:Optional
	Audience []string `json:"audience"`

	// kubeconfig is the path to the kubeconfig file.
	// +kubebuilder:validation:Optional
	Kubeconfig string `json:"kubeconfig"`

	// allowedNodeLabelKeys is the list of allowed node label keys.
	// +kubebuilder:validation:Optional
	AllowedNodeLabelKeys map[string]bool `json:"allowedNodeLabelKeys"`

	// allowedPodLabelKeys is the list of allowed pod label keys.
	// +kubebuilder:validation:Optional
	AllowedPodLabelKeys map[string]bool `json:"allowedPodLabelKeys"`
}
