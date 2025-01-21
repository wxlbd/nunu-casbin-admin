package repository

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/internal/service"
)

type dictTypeRepository struct {
	query *Query
}

type dictDataRepository struct {
	query *Query
}

func NewDictTypeRepository(query *Query) service.DictTypeRepository {
	return &dictTypeRepository{query: query}
}

func NewDictDataRepository(query *Query) service.DictDataRepository {
	return &dictDataRepository{query: query}
}

// DictType Repository 实现

func (r *dictTypeRepository) Create(ctx context.Context, dict *model.DictType) error {
	return r.query.WithContext(ctx).DictType.Create(dict)
}

func (r *dictTypeRepository) Update(ctx context.Context, dict *model.DictType) error {
	return r.query.WithContext(ctx).DictType.Save(dict)
}

func (r *dictTypeRepository) Delete(ctx context.Context, ids ...int64) error {
	_, err := r.query.WithContext(ctx).DictType.Where(r.query.DictType.ID.In(ids...)).Delete()
	return err
}

func (r *dictTypeRepository) FindByID(ctx context.Context, id int64) (*model.DictType, error) {
	return r.query.WithContext(ctx).DictType.Where(r.query.DictType.ID.Eq(id)).First()
}

func (r *dictTypeRepository) FindByCode(ctx context.Context, code string) (*model.DictType, error) {
	return r.query.WithContext(ctx).DictType.Where(r.query.DictType.Code.Eq(code)).First()
}

func (r *dictTypeRepository) List(ctx context.Context, query *model.DictTypeQuery) ([]*model.DictType, int64, error) {
	q := r.query.WithContext(ctx).DictType
	if query.Name != "" {
		q = q.Where(r.query.DictType.Name.Like("%" + query.Name + "%"))
	}
	if query.Code != "" {
		q = q.Where(r.query.DictType.Code.Like("%" + query.Code + "%"))
	}
	if query.Status != 0 {
		q = q.Where(r.query.DictType.Status.Eq(query.Status))
	}

	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	list, err := q.Order(r.query.DictType.Sort).Offset(offset).Limit(query.PageSize).Find()
	return list, total, err
}

// DictData Repository 实现

func (r *dictDataRepository) Create(ctx context.Context, data *model.DictDatum) error {
	return r.query.WithContext(ctx).DictDatum.Create(data)
}

func (r *dictDataRepository) Update(ctx context.Context, data *model.DictDatum) error {
	return r.query.WithContext(ctx).DictDatum.Save(data)
}

func (r *dictDataRepository) Delete(ctx context.Context, ids ...int64) error {
	_, err := r.query.WithContext(ctx).DictDatum.Where(r.query.DictDatum.ID.In(ids...)).Delete()
	return err
}

func (r *dictDataRepository) FindByID(ctx context.Context, id int64) (*model.DictDatum, error) {
	return r.query.WithContext(ctx).DictDatum.Where(r.query.DictDatum.ID.Eq(id)).First()
}

func (r *dictDataRepository) FindByTypeCode(ctx context.Context, typeCode string) ([]*model.DictDatum, error) {
	return r.query.WithContext(ctx).DictDatum.Where(r.query.DictDatum.TypeCode.Eq(typeCode)).Order(r.query.DictDatum.Sort).Find()
}

func (r *dictDataRepository) List(ctx context.Context, query *model.DictDataQuery) ([]*model.DictDatum, int64, error) {
	q := r.query.WithContext(ctx).DictDatum
	if query.TypeCode != "" {
		q = q.Where(r.query.DictDatum.TypeCode.Eq(query.TypeCode))
	}
	if query.Label != "" {
		q = q.Where(r.query.DictDatum.Label.Like("%" + query.Label + "%"))
	}
	if query.Status != 0 {
		q = q.Where(r.query.DictDatum.Status.Eq(query.Status))
	}

	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}

	offset := (query.Page - 1) * query.PageSize
	list, err := q.Order(r.query.DictDatum.Sort).Offset(offset).Limit(query.PageSize).Find()
	return list, total, err
}
