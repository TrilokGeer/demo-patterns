package v1alpha1

const (
	UpstreamAuthorityTypeDisk      = "disk"
	UpstreamAuthorityTypeAWSPCA    = "aws_pca"
	UpstreamAuthorityTypeAWSSecret = "aws_secret"
	UpstreamAuthorityTypeGCPCAS    = "gcp_cas"
	UpstreamAuthorityTypeVault     = "vault"
)

// UpstreamCA has the config required for the spire server upstream CA.
type UpstreamCA struct {
	// Type is the SPIRE UpstreamAuthority plugin type (e.g., "disk", "aws_pca")
	Type string `json:"type"`

	// PluginConfig holds plugin-specific configuration as appendable JSON
	PluginConfig map[string]interface{} `json:"pluginConfig,omitempty"`
}

type UpstreamAuthorityDisk struct {
	// CertFilePath is the path to the CA certificate file
	CertFilePath string `json:"certFilePath"`

	// KeyFilePath is the path to the CA key file
	KeyFilePath string `json:"keyFilePath"`

	// BundleFilePath is the path to the CA bundle file
	BundleFilePath string `json:"bundleFilePath"`
}

type UpstreamAuthorityAWSPCA struct {
	// Region is the AWS region
	Region string `json:"region"`

	// CertificateAuthorityARN is the AWS Certificate Authority ARN
	CertificateAuthorityARN string `json:"certificateAuthorityARN"`

	// CertificateSigningTemplateARN is the AWS Certificate Signing Template ARN
	CertificateSigningTemplateARN string `json:"certificateSigningTemplateARN"`

	// SigningAlgorithm is the signing algorithm to use
	SigningAlgorithm string `json:"signingAlgorithm"`

	// AssumeroleARN is the AWS Assumerole ARN
	AssumeroleARN string `json:"assumeroleARN"`

	// SupplementalBundlePath is the path to the supplemental bundle file
	SupplementalBundlePath string `json:"supplementalBundlePath"`
}

type UpstreamAuthorityAWSSecret struct {
	// Region is the AWS region
	Region string `json:"region"`

	// CertificateFileARN is the AWS Certificate File ARN
	CertificateFileARN string `json:"certificateFileARN"`

	// KeyFileARN is the AWS Key File ARN
	KeyFileARN string `json:"keyFileARN"`

	// BundleFileARN is the AWS Bundle File ARN
	BundleFileARN string `json:"bundleFileARN"`

	// AccessKeyID is the AWS Access Key ID
	AccessKeyID string `json:"accessKeyID"`

	// SecretAccessKey is the AWS Secret Access Key
	SecretAccessKey string `json:"secretAccessKey"`

	// SecretToken is the AWS Secret Token
	SecretToken string `json:"secretToken"`

	// AssumeRoleARN is the AWS Assume Role ARN
	AssumeRoleARN string `json:"assumeRoleARN"`
}

type UpstreamGCPCAS struct {
	// ProjectName is the GCP Project Name
	ProjectName string `json:"projectName"`

	// RegionName is the GCP Region Name
	RegionName string `json:"regionName"`

	// CAPoolName is the GCP CA Pool Name
	CAPoolName string `json:"caPoolName"`

	// LabelKey is the GCP Label Key
	LabelKey string `json:"labelKey"`

	// LabelValue is the GCP Label Value
	LabelValue string `json:"labelValue"`
}

type UpstreamAuthorityVault struct {
	// Address is the Vault Address
	Address string `json:"address"`

	// Namespace is the Vault Namespace
	Namespace string `json:"namespace"`

	// PKIMountName is the Vault PKI Mount Name
	PKIMountName string `json:"pkiMountName"`

	// CACertificatePath is the Vault CA Certificate Path
	CACertificatePath string `json:"caCertificatePath"`

	// InsecureSkipVerify is the Vault Insecure Skip Verify
	InsecureSkipVerify bool `json:"insecureSkipVerify"`

	// ClientCertficateAuth is the Vault Client Certificate Auth
	ClientCertficateAuth *VaultClientCertficateAuth `json:"clientCertficateAuth"`

	TokenAuth *VaultTokenAuth `json:"vaultTokenAuth"`

	// AppRoleAuth is the Vault App Role Auth
	AppRoleAuth *VaultAppRoleAuth `json:"appRoleAuth"`
}

type VaultAppRoleAuth struct {
	// RoleID is the Vault Role ID
	RoleID string `json:"roleID"`

	// SecretID is the Vault Secret ID
	SecretID string `json:"secretID"`

	// MountPoint is the Vault Mount Point
	MountPoint string `json:"mountPoint"`
}

type VaultTokenAuth struct {
	// Token is the Vault Token
	Token string `json:"token"`
}

type VaultClientCertficateAuth struct {
	// ClientCertPath is the Vault Client Cert Path
	ClientCertPath string `json:"clientCertPath"`

	// ClientKeyPath is the Vault Client Key Path
	ClientKeyPath string `json:"clientKeyPath"`

	// ClientAuthRoleName is the Vault Client Auth Role Name
	ClientAuthRoleName string `json:"clientAuthRoleName"`

	// ClientAuthMountPath is the Vault Client Auth Mount Path
	ClientAuthMountPath string `json:"clientAuthMountPath"`
}

type UpstreamAuthoritySpire struct {
	// ServerAddress is the Spire Server Address
	ServerAddress string `json:"serverAddress"`

	// ServerPort is the Spire Server Port
	ServerPort int `json:"serverPort"`

	// WorkloadAPISocketPath is the Spire Workload API Socket Path
	WorkloadAPISocketPath string `json:"workloadAPISocketPath"`
}

type UpstreamAuthorityCertManager struct {
	// KubeconfigPath is the Kubeconfig Path
	KubeconfigPath string `json:"kubeconfigPath"`

	// Namespace is the Namespace
	Namespace string `json:"namespace"`

	// IssuerName is the Issuer Name
	IssuerName string `json:"issuerName"`

	// IssuerKind is the Issuer Kind
	IssuerKind string `json:"issuerKind"`

	// IssuerGroup is the Issuer Group
	IssuerGroup string `json:"issuerGroup"`
}
