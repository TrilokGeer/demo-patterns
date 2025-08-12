package v1alpha1

const (
	DataStoreTypeSQL = "sql"
)

const (
	DatabaseTypeSQLite3     = "sqlite3"
	DatabaseTypePostgres    = "postgres"
	DatabaseTypeMySQL       = "mysql"
	DatabaseTypeAWSPostgres = "aws_postgresql"
	DatabaseTypeAWSMySQL    = "aws_mysql"
)

type DataStore struct {
	// datastoreType specifies the type of datastore to use.
	// +kubebuilder:validation:Enum=sql
	// +kubebuilder:default:=sql
	DatastoreType string `json:"datastoreType"`

	// DataStoreSQL contains the config for the SQL datastore.
	// +kubebuilder:validation:Optional
	DataStoreSQL *DataStoreSQL `json:"dataStoreSQL,omitempty"`
}

type DataStoreSQL struct {
	// DatabaseType specifies type of database to use.
	// +kubebuilder:validation:Enum=sql;sqlite3;postgres;mysql;aws_postgresql;aws_mysql
	// +kubebuilder:default:=sqlite3
	DatabaseType string `json:"databaseType"`

	// DatabaseConfig contains the config for the database.
	// +kubebuilder:validation:Optional
	DatabaseConfig map[string]interface{} `json:"databaseConfig,omitempty"`
	// connectionString contain connection credentials required for spire server Datastore.
	// +kubebuilder:default:=/run/spire/data/datastore.sqlite3
	ConnectionString string `json:"connectionString"`

	// roConnectionString contain read-only connection credentials required for spire server Datastore.
	// +kubebuilder:validation:Optional
	RoConnectionString string `json:"roConnectionString,omitempty"`

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
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	DisableMigration bool `json:"disableMigration,omitempty"`

	// logSQL enables SQL query logging for debugging
	// +kubebuilder:default:=false
	// +kubebuilder:validation:Optional
	LogSQL bool `json:"logSQL,omitempty"`

	// AWSAccessConfig contains the config for the AWS database.
	// +kubebuilder:validation:Optional
	AWSAccessConfig AWSAccessConfig `json:"awsAccessConfig,omitempty"`
}
