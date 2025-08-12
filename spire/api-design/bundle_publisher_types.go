package v1alpha1

const (
	BundlePublisherTypeAWSS3 = "aws_s3"
	BundlePublisherTypeGCS   = "gcs"
)

type BundlePublisher struct {
	// type specifies the type of bundle publisher to use.
	// +kubebuilder:validation:Enum=aws_s3;gcs;aws_rolesanywhere_trustanchor
	// +kubebuilder:default:=aws_s3
	Type string `json:"type"`

	// BundlePublisherAWSS3 contains the config for the AWS S3 bundle publisher.
	// +kubebuilder:validation:Optional
	BundlePublisherAWSS3 *BundlePublisherAWSS3 `json:"awsS3,omitempty"`

	// BundlePublisherGCS contains the config for the GCS bundle publisher.
	// +kubebuilder:validation:Optional
	BundlePublisherGCS *BundlePublisherGCS `json:"gcs,omitempty"`
}

type BundlePublisherAWSS3 struct {
	AWSAccessConfig

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
