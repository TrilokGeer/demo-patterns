package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:validation:XValidation:rule="self.metadata.name == 'cluster'",message="SpireServer is a singleton, .metadata.name must be 'cluster'"
// +operator-sdk:csv:customresourcedefinitions:displayName="SpireServer"

// SpireServer defines the configuration for the SPIRE Server managed by zero trust workload identity manager.
// This includes details related to trust domain, data storage, plugins
// and other configs required for workload authentication.
type SpireServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpireServerSpec   `json:"spec,omitempty"`
	Status            SpireServerStatus `json:"status,omitempty"`
}

// SpireServerSpec will have specifications for configuration related to the spire server.
type SpireServerSpec struct {

	// jwtIssuer is the JWT issuer domain. Defaults to oidc-discovery.$trustDomain if unset
	// +kubebuilder:validation:Optional
	JwtIssuer string `json:"jwtIssuer,omitempty"`

	// keyManager has configs for the spire server key manager.
	// +kubebuilder:validation:Optional
	KeyManager *KeyManager `json:"keyManager,omitempty"`

	// caSubject contains subject information for the Spire CA.
	// +kubebuilder:validation:Optional
	CASubject *CASubject `json:"caSubject,omitempty"`

	// spireSQLConfig has the config required for the spire server SQL DataStore.
	// +kubebuilder:validation:Optional
	Datastore *DataStoreConfig `json:"datastore,omitempty"`

	// upstreamCAConfig has the config required for the spire server upstream CA.
	// +kubebuilder:validation:Optional
	UpstreamCAConfig *UpstreamCAConfig `json:"upstreamCAConfig,omitempty"`

	// notifierConfig has the config required for the spire server notifier.
	// +kubebuilder:validation:Optional
	NotifierConfig *NotifierConfig `json:"notifierConfig,omitempty"`

	// bundlePublisherConfig has the config required for the spire server bundle publisher.
	// aws_anywhere is a special case where the bundle publisher is configured to publish bundles to AWS Anywhere
	// and is not supported for now.
	// +kubebuilder:validation:Optional
	BundlePublisherConfig *BundlePublisherConfig `json:"bundlePublisherConfig,omitempty"`

	// deploymentConfig has the config required for the spire server deployment.
	// +kubebuilder:validation:Optional
	Deployment `json:",inline"`
}

// DataStoreConfig configures the Spire SQL datastore backend.
type DataStoreConfig struct {
	// databaseType specifies type of database to use.
	// +kubebuilder:validation:Enum=sql;sqlite3;postgres;mysql;aws_postgresql;aws_mysql
	// +kubebuilder:default:=sqlite3
	DatabaseType string `json:"databaseType"`

	// connectionString contain connection credentials required for spire server Datastore.
	// +kubebuilder:default:=/run/spire/data/datastore.sqlite3
	ConnectionString string `json:"connectionString"`

	// options specifies extra DB options.
	// +kubebuilder:validation:optional
	// +kubebuilder:default:={}
	Options []string `json:"options,omitempty"`

	// MySQL TLS options.
	// +kubebuilder:default:=""
	RootCAPath     string `json:"rootCAPath,omitempty"`
	ClientCertPath string `json:"clientCertPath,omitempty"`
	ClientKeyPath  string `json:"clientKeyPath,omitempty"`

	// DB pool config
	// maxOpenConns will specify the maximum connections for the DB pool.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default:=100
	MaxOpenConns int `json:"maxOpenConns"`

	// maxIdleConns specifies the maximum idle connection to be configured.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default:=2
	MaxIdleConns int `json:"maxIdleConns"`

	// connMaxLifetime will specify maximum lifetime connections.
	// Max time (in seconds) a connection may live.
	// +kubebuilder:validation:Minimum=0
	ConnMaxLifetime int `json:"connMaxLifetime"`

	// disableMigration specifies the migration state
	// If true, disables DB auto-migration.
	// +kubebuilder:default:="false"
	// +kubebuilder:validation:Enum:="true";"false"
	// +kubebuilder:validation:Optional
	DisableMigration string `json:"disableMigration"`
}

type KeyManagerConfig struct {
	// type specifies the type of key manager to use.
	// +kubebuilder:validation:Enum=aws_kms;memory;disk
	// +kubebuilder:default:=memory
	Type string `json:"type"`

	// pluginConfig has the config required for the spire server key manager.
	// +kubebuilder:validation:Optional
	PluginConfig map[string]interface{} `json:"pluginConfig,omitempty"`
}

// CASubject defines the subject information for the Spire CA.
type CASubject struct {
	// country specifies the country for the CA.
	// +kubebuilder:validation:Optional
	Country string `json:"country,omitempty"`

	// organization specifies the organization for the CA.
	// +kubebuilder:validation:Optional
	Organization string `json:"organization,omitempty"`

	// commonName specifies the common name for the CA.
	// +kubebuilder:validation:Optional
	CommonName string `json:"commonName,omitempty"`
}

// SpireServerStatus defines the observed state of spire-server related reconciliation made by operator
type SpireServerStatus struct {
	// conditions holds information of the current state of the spire-server resources.
	ConditionalStatus `json:",inline,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SpireServerList contain the list of SpireServer
type SpireServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpireServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpireServer{}, &SpireServerList{})
}
