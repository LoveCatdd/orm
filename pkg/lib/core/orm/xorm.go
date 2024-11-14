package orm

import (
	"fmt"

	// _ "gitee.com/travelliu/dm"           // dameng-driver
	// _ "github.com/denisenkom/go-mssqldb" // sqlserver-driver
	// _ "github.com/go-sql-driver/mysql"   // mysql-driver
	// _ "github.com/godror/godror"         // oracle-driver
	// _ "github.com/jackc/pgx/v5" 			// postgresql-driver
	// _ "modernc.org/sqlite"               // sqlite-driver

	"github.com/LoveCatdd/util/pkg/lib/core/log"
	"xorm.io/xorm"
)

var _egroup *xorm.EngineGroup

// =======================================================//
// Out
type Xorm struct{}

func (Xorm) Master() *xorm.Engine {
	return SetEngineWrapper(_egroup.Master())
}

func (Xorm) Slave() *xorm.Engine {
	return SetEngineWrapper(_egroup.Slave())
}

// =====================================================//

func setEngine(ormConf *ORMConfig, userGroupPolicy UserGroupPolicy) {

	var engines []*xorm.Engine

	// Engine
	for _, connect := range ormConf.Orm.Xorm.Engine {

		dbName := connect.DBName

		if url := withUrl(dbName, connect); url != "" {
			engine, err := xorm.NewEngine(dbName, url)
			if err != nil {
				log.Errorf("new engine error %v", err)
				return
			}
			engines = append(engines, engine)
		}
	}

	var (
		eg  *xorm.EngineGroup
		err error
	)

	if len(engines) > 1 {
		eg, err = xorm.NewEngineGroup(engines[0], engines[1:],
			withPolicy(ormConf.Orm.Xorm.Policy, ormConf.Orm.Xorm.Weight, userGroupPolicy))
	} else {
		eg, err = xorm.NewEngineGroup(engines[0], []*xorm.Engine{})
	}

	if err != nil {
		log.Errorf("new engine group error %v", err)
		return
	}

	if ormConf.Orm.Xorm.Show {
		xormLogger := NewXormLogger(withLevel(ormConf.Orm.Xorm.Level))
		xormLogger.ShowSQL(ormConf.Orm.Xorm.Show)
		eg.SetLogger(xormLogger)
	}

	// // 和SnakeMapper很类似，但是对于特定词支持更好，比如ID会翻译成id而不是i_d
	// eg.SetMapper(names.GonicMapper{})

	_egroup = eg

	log.Infof("app Init ORM: XORM, successful run [master: 1, slaves: %v, groupPolicy: %v, showLog: %v]",
		len(_egroup.Slaves()), ormConf.Orm.Xorm.Policy, ormConf.Orm.Xorm.Show)

}

func withPolicy(policy string, weight []int, userGroupPolicy UserGroupPolicy) xorm.GroupPolicy {
	switch policy {
	case RandomPolicy:
		return xorm.RandomPolicy()
	case WeightRandomPolicy:
		return xorm.WeightRandomPolicy(weight)
	case RoundRobinPolicy:
		return xorm.RoundRobinPolicy()
	case WeightRoundRobinPolicy:
		return xorm.WeightRoundRobinPolicy(weight)
	case LeastConnPolicy:
		return xorm.LeastConnPolicy()
	case UserDefined:
		return userGroupPolicy.UserDefined()
	default:
		return xorm.RandomPolicy()
	}
}

func withUrl(dbname string, connect ConnectInfo) string {
	switch dbname {
	case MYSQL:
		return _MySQL(connect.User, connect.Password, connect.Url)
	case TIDB:
		return _TiDB(connect.User, connect.Password, connect.Url)
	case POSTGRESQL:
		return _PostgreSQLX(connect.User, connect.Password, connect.Url)
	case SQLITEGO:
		return _SQLite(connect.Url)
	case MSSQL:
		return _MsSql(connect.User, connect.Password, connect.Url)
	case DAMENG:
		return _Dameng(connect.User, connect.Password, connect.Url)
	case ORCALE:
		return _Oracle(connect.User, connect.Password, connect.Url)
	default:
		log.Errorf("dbname isnot unfound %v", dbname)
		return ""
	}

	// case MYMYSQL:
	// 	// Deprecated: Use MYSQL instead.
	// 	return _MyMySQL(connect.User, connect.Password, connect.Url)

	// case ORCALEOCI8:
	// 	// Deprecated: Use ORCALE instead.
	// 	return _Oracle(connect.User, connect.Password, connect.Url)

	// case SQLITE3:
	// 	// Deprecated: Use SQLITE instead.
	// 	return _SQLite(connect.Url)

}

func _MySQL(user, password, url string) string {

	// "user:password@tcp(127.0.0.1:3306)/dbname"
	return fmt.
		Sprintf("%v:%v@%v", user, password, url)
}

func _TiDB(user, password, url string) string {

	// "user:password@tcp(127.0.0.1:4000)/dbname"
	return fmt.
		Sprintf("%v:%v@%v", user, password, url)
}

func _PostgreSQLX(user, password, url string) string {
	// "postgresql://username:password@localhost:5432/dbname?sslmode=require"
	return fmt.
		Sprintf("postgresql://%v:%v@%v", user, password, url)
}

func _SQLite(url string) string {
	// "./test.db"
	return fmt.
		Sprintf("./%v", url)
}

func _MsSql(user, password, url string) string {

	// "sqlserver://username:password@localhost:1433?database=dbname"
	return fmt.
		Sprintf("sqlserver://%v:%v@%v", user, password, url)
}

func _Dameng(user, password, url string) string {
	// "dm://username:password@host:port?schema=dbname"
	return fmt.
		Sprintf("dm://%v:%v@%v", user, password, url)
}

func _Oracle(user, password, url string) string {

	// "user/password@localhost:1521/dbname"
	return fmt.
		Sprintf("%v/%v@%v", user, password, url)
}

// Deprecated: Use MYSQL instead.
// func _MyMySQL(user, password, url string) string {

// 	// "tcp:127.0.0.1:3306*dbname/user/password"
// 	return fmt.
// 		Sprintf("%v/%v/%v", url, user, password)
// }
