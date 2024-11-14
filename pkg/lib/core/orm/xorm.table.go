package orm

import (
	"time"

	"github.com/LoveCatdd/util/pkg/lib/core/log"
	"xorm.io/xorm"
)

type JsonTime time.Time

// 序列化时间格式
func (jt JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(jt).Format("2006-01-02 15:04:05") + `"`), nil
}

// Model Auto 自动处理
type ModelAuto struct {
	CreatedAt JsonTime `xorm:"created notnull comment('创建时间')" json:"createdAt" yaml:"createdAt" mapstructure:"createdAt"`     // Insert()或InsertOne()方法被调用时, 自动更新为当前时间或者
	UpdatedAt JsonTime `xorm:"updated notnull comment('更新时间')" json:"updatedAt" yaml:"updatedAt" mapstructure:"updatedAt"`     // Insert(), InsertOne(), Update()方法被调用时, 自动更新为当前时间
	DeletedAt JsonTime `xorm:"deleted notnull comment('删除时间')" json:"deletedAt" yaml:"deletedAt" mapstructure:"deletedAt"`     // 在Delete()时, 自动更新为当前时间而不是去删除该条记录 软删除, 后续对软删除的数据操作 需要启用Unscoped
	CreatedBy string   `xorm:"Varchar(255) notnull comment('创建人')" json:"createdBy" yaml:"createdBy" mapstructure:"createdBy"` // 创建人 格式：username#userId
	UpdatedBy string   `xorm:"Varchar(255) notnull comment('更新人')" json:"updatedBy" yaml:"updatedBy" mapstructure:"updatedBy"` // 更新人 格式：username#userId
	DeletedBy string   `xorm:"Varchar(255) notnull comment('删除人')" json:"deletedBy" yaml:"deletedBy" mapstructure:"deletedBy"` // 删除人 格式：username#userId

}

// Model Base model 基础字段
type ModelBase struct {
	Id string `xorm:"Varchar(255) notnull pk unique comment('主键')" json:"id"` // Id 主键

	ModelAuto `xorm:"extends"`
}

type Table interface {
	TableName() string
}

type tableList []Table

var tables = make(tableList, 0)

// 注册表
func TableRegister(table Table) {
	tables = append(tables, table)
}

// 创建表结构
func InitCreateTable(engine *xorm.Engine) {

	for _, table := range tables {

		if engine.DriverName() == MYSQL {
			err := engine.StoreEngine("InnoDB").Sync(table)
			if err != nil {
				log.Fatalf("create tables fatal err: %v", err)
			}
		} else {
			err := engine.Sync(table)
			if err != nil {
				log.Fatalf("create tables fatal err: %v", err)
			}

		}
	}

}
