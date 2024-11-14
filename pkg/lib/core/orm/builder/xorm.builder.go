package orm_builder

import "github.com/LoveCatdd/orm/pkg/lib/core/orm"

type BaseOption func(base *orm.ModelBase)

func NewModelBase(id string, ops ...BaseOption) *orm.ModelBase {
	base := &orm.ModelBase{
		Id: id,
	}

	for _, o := range ops {
		o(base)
	}

	return base
}

func WithDeletedBy(deletedBy string) BaseOption {
	return func(base *orm.ModelBase) {
		base.DeletedBy = deletedBy
	}
}

func WithCreatedBy(createdBy string) BaseOption {
	return func(base *orm.ModelBase) {
		base.CreatedBy = createdBy
	}
}

func WithUpdatedBy(updatedBy string) BaseOption {
	return func(base *orm.ModelBase) {
		base.UpdatedBy = updatedBy
	}
}
