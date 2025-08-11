package v1alpha1

const (
	DataStoreTypeSQL = "sql"
)

type DataStore struct {
	// datastoreType specifies the type of datastore to use.
	// +kubebuilder:validation:Enum=sql;sqlite3;postgres;mysql;aws_postgresql;aws_mysql
	// +kubebuilder:default:=sqlite3
	DatastoreType string `json:"datastoreType"`

	// pluginConfig has the config required for the spire server datastore.
	// +kubebuilder:validation:Optional
	PluginConfig map[string]interface{} `json:"pluginConfig,omitempty"`
}

type DataStoreSQL struct {

	// connectionString contain connection credentials required for spire server Datastore.
	// +kubebuilder:default:=/run/spire/data/datastore.sqlite3
	ConnectionString string `json:"connectionString"`

	// RootConnectionString contain connection credentials required for spire server Datastore.
	// +kubebuilder:default:=""
	RootConnectionString string `json:"rootConnectionString,omitempty"`

	// MySQL TLS options.
	// +kubebuilder:default:=""
	RootCAPath     string `json:"rootCAPath,omitempty"`
	ClientCertPath string `json:"clientCertPath,omitempty"`
	ClientKeyPath  string `json:"clientKeyPath,omitempty"`

	// databaseType specifies type of database to use.
	// +kubebuilder:validation:Enum=sql;sqlite3;postgres;mysql;aws_postgresql;aws_mysql
	// +kubebuilder:default:=sqlite3
	DatabaseTypeConfig map[string]interface{} `json:"databaseTypeConfig,omitempty"`

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

	// DatabaseConfig contains the config for the database.
	// +kubebuilder:validation:Optional
	DatabaseType map[string]interface{} `json:"databaseType,omitempty"`
}

type DatabaseTypeAWSPostgres struct {
	AWSAccessConfig
}

type DatabaseTypeAWSMySQL struct {
	AWSAccessConfig
}

type AWSAccessConfig struct {
	// Region specifies the region of the AWS database.
	// +kubebuilder:validation:Optional
	Region string `json:"region"`

	// AccessKeyID specifies the access key ID of the AWS database.
	// +kubebuilder:validation:Optional
	AccessKeyID string `json:"accessKeyID,omitempty"`

	// SecretAccessKey specifies the secret access key of the AWS database.
	// +kubebuilder:validation:Optional
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
}

// DataStore configures the Spire SQL datastore backend.
type DataStore struct {
	// databaseType specifies type of database to use.
	// +kubebuilder:validation:Enum=sql;sqlite3;postgres;mysql;aws_postgresql;aws_mysql
	// +kubebuilder:default:=sqlite3
	DatabaseType string `json:"databaseType"`

	// options specifies extra DB options.
	// +kubebuilder:validation:optional
	// +kubebuilder:default:={}
	Options []string `json:"options,omitempty"`
}
