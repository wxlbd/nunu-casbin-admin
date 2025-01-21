package service

import (
	"context"

	"github.com/wxlbd/gin-casbin-admin/internal/handler"

	"github.com/wxlbd/gin-casbin-admin/internal/dto"
	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"github.com/wxlbd/gin-casbin-admin/pkg/log"
)

type dictService struct {
	log      *log.Logger
	repo     Repository
	typeRepo DictTypeRepository
	dataRepo DictDataRepository
}

func NewDictService(logger *log.Logger, repo Repository) handler.DictService {
	return &dictService{
		log:      logger,
		repo:     repo,
		typeRepo: repo.DictType(),
		dataRepo: repo.DictData(),
	}
}

// DictType 实现

// CreateDictType DictType
func (s *dictService) CreateDictType(ctx context.Context, req *dto.DictTypeRequest) error {
	// 检查编码是否存在
	exist, err := s.typeRepo.FindByCode(ctx, req.Code)
	if err != nil {
		return err
	}
	if exist != nil {
		return errors.WithMsg(errors.AlreadyExists, "字典类型编码已存在")
	}

	dictType := &model.DictType{
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
		Sort:   req.Sort,
		Remark: req.Remark,
	}

	return s.typeRepo.Create(ctx, dictType)
}

func (s *dictService) UpdateDictType(ctx context.Context, req *dto.DictTypeRequest) error {
	// 检查是否存在
	exist, err := s.typeRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.WithMsg(errors.NotFound, "字典类型不存在")
	}

	// 如果修改了编码，检查新编码是否存在
	if req.Code != exist.Code {
		if existCode, _ := s.typeRepo.FindByCode(ctx, req.Code); existCode != nil {
			return errors.WithMsg(errors.AlreadyExists, "字典类型编码已存在")
		}
	}

	dictType := &model.DictType{
		ID:     req.ID,
		Name:   req.Name,
		Code:   req.Code,
		Status: req.Status,
		Sort:   req.Sort,
		Remark: req.Remark,
	}

	return s.typeRepo.Update(ctx, dictType)
}

func (s *dictService) DeleteDictType(ctx context.Context, ids ...int64) error {
	return s.typeRepo.Delete(ctx, ids...)
}

func (s *dictService) GetDictType(ctx context.Context, id int64) (*model.DictType, error) {
	return s.typeRepo.FindByID(ctx, id)
}

func (s *dictService) ListDictType(ctx context.Context, query *model.DictTypeQuery) ([]*model.DictType, int64, error) {
	return s.typeRepo.List(ctx, query)
}

// DictData 实现
func (s *dictService) CreateDictData(ctx context.Context, req *dto.DictDataRequest) error {
	// 检查字典类型是否存在
	dictType, err := s.typeRepo.FindByCode(ctx, req.TypeCode)
	if err != nil {
		return err
	}
	if dictType == nil {
		return errors.WithMsg(errors.NotFound, "字典类型不存在")
	}

	dictData := &model.DictDatum{
		TypeCode: req.TypeCode,
		Label:    req.Label,
		Value:    req.Value,
		Status:   req.Status,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}

	return s.dataRepo.Create(ctx, dictData)
}

func (s *dictService) UpdateDictData(ctx context.Context, req *dto.DictDataRequest) error {
	// 检查是否存在
	exist, err := s.dataRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if exist == nil {
		return errors.WithMsg(errors.NotFound, "字典数据不存在")
	}

	// 如果修改了类型，检查新类型是否存在
	if req.TypeCode != exist.TypeCode {
		dictType, err := s.typeRepo.FindByCode(ctx, req.TypeCode)
		if err != nil {
			return err
		}
		if dictType == nil {
			return errors.WithMsg(errors.NotFound, "字典类型不存在")
		}
	}

	dictData := &model.DictDatum{
		ID:       req.ID,
		TypeCode: req.TypeCode,
		Label:    req.Label,
		Value:    req.Value,
		Status:   req.Status,
		Sort:     req.Sort,
		Remark:   req.Remark,
	}

	return s.dataRepo.Update(ctx, dictData)
}

func (s *dictService) DeleteDictData(ctx context.Context, ids ...int64) error {
	return s.dataRepo.Delete(ctx, ids...)
}

func (s *dictService) GetDictData(ctx context.Context, id int64) (*model.DictDatum, error) {
	return s.dataRepo.FindByID(ctx, id)
}

func (s *dictService) ListDictData(ctx context.Context, query *model.DictDataQuery) ([]*model.DictDatum, int64, error) {
	return s.dataRepo.List(ctx, query)
}

func (s *dictService) GetDictDataByType(ctx context.Context, typeCode string) ([]*model.DictDatum, error) {
	return s.dataRepo.FindByTypeCode(ctx, typeCode)
}
