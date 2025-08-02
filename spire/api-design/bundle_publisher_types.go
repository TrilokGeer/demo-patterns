package v1alpha1

type BundlePublisher struct {
	// type specifies the type of bundle publisher to use.
	// +kubebuilder:validation:Enum=aws_s3;gcs;
	// +kubebuilder:default:=aws_s3
	Type string `json:"type"`

	// pluginConfig has the config required for the spire server bundle publisher.
	// +kubebuilder:validation:Optional
	PluginConfig map[string]interface{} `json:"pluginConfig,omitempty"`
}

type BundlePublisherAWSS3 struct {
	// region is the AWS region.
	// +kubebuilder:validation:Optional
	Region string `json:"region,omitempty"`

	// accessKeyID is the AWS access key ID.
	// +kubebuilder:validation:Optional
	AccessKeyID string `json:"accessKeyID,omitempty"`

	// secretAccessKey is the AWS secret access key.
	// +kubebuilder:validation:Optional
	SecretAccessKey string `json:"secretAccessKey,omitempty"`

	// bucketName is the name of the S3 bucket.
	// +kubebuilder:validation:Optional
	BucketName string `json:"bucketName,omitempty"`

	// objectKey is the key of the S3 object.
	// +kubebuilder:validation:Optional
	ObjectKey string `json:"objectKey,omitempty"`

	// objectFormat is the format of the S3 object.
	// +kubebuilder:validation:Enum=json;yaml
	// +kubebuilder:default:=json
	ObjectFormat string `json:"objectFormat,omitempty"`

	// endpoint is the S3 endpoint.
	// +kubebuilder:validation:Optional
	Endpoint string `json:"endpoint,omitempty"`
}

type BundlePublisherGCS struct {
	// credentialsFile is the path to the GCS credentials file.
	// +kubebuilder:validation:Optional
	CredentialsFile string `json:"credentialsFile,omitempty"`

	// bucketName is the name of the GCS bucket.
	// +kubebuilder:validation:Optional
	BucketName string `json:"bucketName"`

	// objectName is the name of the GCS object.
	// +kubebuilder:validation:Optional
	ObjectName string `json:"objectName"`

	// objectFormat is the format of the GCS object.
	// +kubebuilder:validation:Enum=json;yaml
	// +kubebuilder:default:=json
	ObjectFormat string `json:"objectFormat"`
}
