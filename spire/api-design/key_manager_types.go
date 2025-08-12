package v1alpha1

const (
	KeyManagerTypeMemory = "memory"
	KeyManagerTypeDisk   = "disk"
	KeyManagerTypeAWSKMS = "aws_kms"
)

type KeyManager struct {
	// type specifies the type of key manager to use.
	// +kubebuilder:validation:Enum=memory;disk;aws_kms
	// +kubebuilder:default:=memory
	Type string `json:"type"`

	// KeyManagerMemory contains the config for the spire server memory key manager.
	// +kubebuilder:validation:Optional
	KeyManagerMemory *KeyManagerMemory `json:"keyManagerMemory,omitempty"`

	// KeyManagerDisk contains the config for the spire server disk key manager.
	// +kubebuilder:validation:Optional
	KeyManagerDisk *KeyManagerDisk `json:"keyManagerDisk,omitempty"`

	// KeyManagerAWSKMS contains the config for the spire server AWS KMS key manager.
	// +kubebuilder:validation:Optional
	KeyManagerAWSKMS *KeyManagerAWSKMS `json:"keyManagerAWSKMS,omitempty"`
}

// KeyManagerMemory has the config required for the spire server memory key manager.
type KeyManagerMemory struct {
	// Memory key manager has no configuration parameters
	// This struct exists for API consistency and future extensibility
}

// KeyManagerDisk has the config required for the spire server disk key manager.
type KeyManagerDisk struct {
	// keys path is the path to the keys directory.
	KeysPath string `json:"keysPath"`
}

// AWSKMSConfig has the config required for the spire server AWS KMS key manager.
type KeyManagerAWSKMS struct {
	AWSAccessConfig

	// keyIdentifier is the AWS KMS key identifier.
	// +kubebuilder:validation:Optional
	KeyIdentifier string `json:"keyIdentifier,omitempty"`

	// key policy file path.
	// +kubebuilder:validation:Optional
	KeyPolicyFilePath string `json:"keyPolicyFilePath,omitempty"`
}
