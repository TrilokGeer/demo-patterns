package v1alpha1

const (
	KeyManagerTypeMemory = "memory"
	KeyManagerTypeDisk   = "disk"
	KeyManagerTypeAWSKMS = "aws_kms"
)

type KeyManager struct {
	// type specifies the type of bundle publisher to use.
	// +kubebuilder:validation:Enum=aws_s3;gcs;
	// +kubebuilder:default:=aws_s3
	Type string `json:"type"`

	// pluginConfig has the config required for the spire server bundle publisher.
	// +kubebuilder:validation:Optional
	PluginConfig map[string]interface{} `json:"pluginConfig,omitempty"`
}

// KeyManagerDisk has the config required for the spire server disk key manager.
type KeyManagerDisk struct {
	// keys path is the path to the keys directory.
	KeysPath string `json:"keysPath"`
}

// AWSKMSConfig has the config required for the spire server AWS KMS key manager.
type KeyManagerAWSKMS struct {
	// accessKeyID is the AWS access key ID.
	// +kubebuilder:validation:Optional
	AccessKeyID string `json:"accessKeyID,omitempty"`

	// secretAccessKey is the AWS secret access key.
	// +kubebuilder:validation:Optional
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// region is the AWS region.
	// +kubebuilder:validation:Optional
	Region string `json:"region,omitempty"`

	// keyIdentifier is the AWS KMS key identifier.
	// +kubebuilder:validation:Optional
	KeyIdentifier string `json:"keyIdentifier,omitempty"`

	// key policy file path.
	// +kubebuilder:validation:Optional
	KeyPolicyFilePath string `json:"keyPolicyFilePath,omitempty"`
}
