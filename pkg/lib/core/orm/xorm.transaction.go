package orm

import (
	"github.com/LoveCatdd/util/pkg/lib/core/log"
	"xorm.io/xorm"
)

type transactionFunc func(*xorm.Session) error

type engineWrapper struct {
	engine *xorm.Engine
}

var ew engineWrapper

// TransactionOps 接受多个 transactionFunc 函数并在事务中依次执行。
// 如果任何一个 transactionFunc 出错，则回滚整个事务。
func (ew engineWrapper) TransactionOps(bodies ...transactionFunc) error {
	session := ew.engine.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	for _, body := range bodies {
		if err := body(session); err != nil {
			_ = session.Rollback() // 如果任一函数出错则回滚
			log.Errorf("Transaction had rollback err: %v", err)
			return err
		}
	}

	if err := session.Commit(); err != nil {
		_ = session.Rollback() // 提交失败时回滚事务
		log.Errorf("Transaction commit failed and had rollback err: %v", err)
		return err
	}

	return nil
}

func SetEngineWrapper(e *xorm.Engine) *xorm.Engine {
	ew = engineWrapper{
		engine: e,
	}
	return e
}

// 注意如果您使用的是 mysql，数据库引擎为 innodb 事务才有效，myisam 引擎是不支持事务的。
func (x Xorm) MasterEngineWrapper() engineWrapper {
	x.Master()
	return ew
}

// 注意如果您使用的是 mysql，数据库引擎为 innodb 事务才有效，myisam 引擎是不支持事务的。
func (x Xorm) SlaveEngineWrapper() engineWrapper {
	x.Slave()
	return ew
}
