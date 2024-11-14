package orm

import "github.com/LoveCatdd/util/pkg/lib/core/viper"

func Init(userGroupPolicy UserGroupPolicy) {

	viper.Yaml(ormConf)
	if ormConf.Orm.Enable == "xorm" {
		setEngine(ormConf, userGroupPolicy)
	}
}

func DefaultInit() {
	Init(nil)
}
