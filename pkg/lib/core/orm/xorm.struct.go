package orm

import (
	"github.com/LoveCatdd/util/pkg/lib/core/viper"
	"xorm.io/xorm"
)

type ORMConfig struct {
	Orm struct {
		Enable string `mapstructure:"enable"`
		Xorm   struct {
			Engine []ConnectInfo `mapstructure:"engine"`
			Policy string        `mapstructure:"groupPolicy"`
			Weight []int         `mapstructure:"weight"`
			Show   bool          `mapstructure:"showLog"`
			Level  string        `mapstructure:"level"`
		} `mapstructure:"xorm"`
		Gorm struct {
			Engine []ConnectInfo `mapstructure:"engine"`
		} `mapstructure:"gorm"`
	} `mapstructure:"orm"`
}

type ConnectInfo struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
	DBName   string `mapstructure:"dbname"`
}

func (*ORMConfig) FileType() string {
	return viper.VIPER_YAML
}

var ormConf = new(ORMConfig)

// dbname constant
const (
	MYSQL      = "mysql"
	SQLITEGO   = "sqlite"
	MSSQL      = "mssql"
	DAMENG     = "dm"
	ORCALE     = "oracle"
	TIDB       = "tidb"
	POSTGRESQL = "pgx"

	// Deprecated: Use MYSQL instead.
	// MYMYSQL    = "mymysql"

	// Deprecated: Use sqlite instead.
	// SQLITE3    = "sqlite3"

	// Deprecated: Use oracle instead.
	// ORCALEOCI8 = "oci8"
)

// slave db 负载均衡 constant
const (
	RandomPolicy           = "RandomPolicy"           // 随机访问负载策略
	WeightRandomPolicy     = "WeightRandomPolicy"     // 权重随机访问负载策略 []int{2, 3}
	RoundRobinPolicy       = "RoundRobinPolicy"       // 轮询访问负载策略
	WeightRoundRobinPolicy = "WeightRoundRobinPolicy" // 权重轮询访问负载策略  []int{2, 3}
	LeastConnPolicy        = "LeastConnPolicy"        // 最小连接数访问负载策略
	UserDefined            = "UserDefined"            // 自定义负载策略
)

type UserGroupPolicy interface {
	// 自定义负载均衡
	UserDefined() xorm.GroupPolicyHandler
}

// log level constant
const (
	LOG_DEBUG   = "debug"
	LOG_INFO    = "info"
	LOG_WARNING = "warn"
)
