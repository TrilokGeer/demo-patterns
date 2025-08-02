package v1alpha1

type Notifier struct {
	// type specifies the type of notifier to use.
	// +kubebuilder:validation:Enum=spire-server;spire-agent
	// +kubebuilder:default:=spire-server
	Type string `json:"type"`

	// pluginConfig has the config required for the spire server notifier.
	// +kubebuilder:validation:Optional
	PluginConfig map[string]interface{} `json:"pluginConfig,omitempty"`
}

type NotifierGCSBundle struct {
	// bucketName is the name of the GCS bucket.
	// +kubebuilder:validation:Optional
	BucketName string `json:"bucketName,omitempty"`

	// objectPath is the path to the GCS object.
	// +kubebuilder:validation:Optional
	ObjectPath string `json:"objectPath,omitempty"`

	// credentialsFile is the path to the GCS credentials file.
	// +kubebuilder:validation:Optional
	CredentialsFile string `json:"credentialsFile,omitempty"`
}

type NotifierK8SBundle struct {
	// namespace is the namespace of the K8S secret.
	// +kubebuilder:validation:Optional
	Namespace string `json:"namespace,omitempty"`

	// configMapName is the name of the K8S config map.
	// +kubebuilder:validation:Optional
	ConfigMap string `json:"configMap,omitempty"`

	// configMapKey is the key of the K8S config map.
	// +kubebuilder:validation:Optional
	ConfigMapKey string `json:"configMapKey,omitempty"`

	// kubeconfigPath is the path to the K8S kubeconfig file.
	// +kubebuilder:validation:Optional
	KubeconfigPath string `json:"kubeconfigPath,omitempty"`

	// apiserviceLabel is a boolean to enable the APIService label.
	// +kubebuilder:validation:Optional
	APIServiceLabel bool `json:"apiserviceLabel,omitempty"`

	// webhookLabel is a boolean to enable the Webhook label.
	// +kubebuilder:validation:Optional
	WebhookLabel bool `json:"webhookLabel,omitempty"`

	// clusters is the list of clusters to watch.
	// +kubebuilder:validation:Optional
	Clusters []ClusterConfig `json:"clusters,omitempty"`
}

type ClusterConfig struct {
	// kubeconfigPath is the path to the K8S kubeconfig file.
	// +kubebuilder:validation:Optional
	KubeconfigPath string `json:"kubeconfigPath,omitempty"`
}
