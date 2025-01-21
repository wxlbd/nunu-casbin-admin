package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wxlbd/gin-casbin-admin/internal/dto"
	"github.com/wxlbd/gin-casbin-admin/internal/model"
	"github.com/wxlbd/gin-casbin-admin/pkg/errors"
	"github.com/wxlbd/gin-casbin-admin/pkg/ginx"
)

type DictHandler struct {
	svc DictService
}

func NewDictHandler(svc DictService) *DictHandler {
	return &DictHandler{
		svc: svc,
	}
}

// CreateDictType 创建字典类型
func (h *DictHandler) CreateDictType(c *gin.Context) {
	var req dto.DictTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	if err := h.svc.CreateDictType(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// UpdateDictType 更新字典类型
func (h *DictHandler) UpdateDictType(c *gin.Context) {
	var req dto.DictTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
		return
	}
	req.ID = id

	if err := h.svc.UpdateDictType(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// DeleteDictType 删除字典类型
func (h *DictHandler) DeleteDictType(c *gin.Context) {
	ids := strings.Split(c.Param("ids"), ",")
	var idList []int64
	for _, id := range ids {
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
			return
		}
		idList = append(idList, idInt)
	}

	if err := h.svc.DeleteDictType(c, idList...); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// GetDictType 获取字典类型详情
func (h *DictHandler) GetDictType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
		return
	}

	dictType, err := h.svc.GetDictType(c, id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, dto.ToDictTypeResponse(dictType))
}

// ListDictType 获取字典类型列表
func (h *DictHandler) ListDictType(c *gin.Context) {
	var query model.DictTypeQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ginx.ParamError(c, err)
		return
	}

	list, total, err := h.svc.ListDictType(c, &query)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	var resp []*dto.DictTypeResponse
	for _, item := range list {
		resp = append(resp, dto.ToDictTypeResponse(item))
	}

	ginx.Success(c, gin.H{
		"list":  resp,
		"total": total,
	})
}

// CreateDictData 创建字典数据
func (h *DictHandler) CreateDictData(c *gin.Context) {
	var req dto.DictDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	if err := h.svc.CreateDictData(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// UpdateDictData 更新字典数据
func (h *DictHandler) UpdateDictData(c *gin.Context) {
	var req dto.DictDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ginx.ParamError(c, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
		return
	}
	req.ID = id

	if err := h.svc.UpdateDictData(c, &req); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// DeleteDictData 删除字典数据
func (h *DictHandler) DeleteDictData(c *gin.Context) {
	ids := strings.Split(c.Param("ids"), ",")
	var idList []int64
	for _, id := range ids {
		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
			return
		}
		idList = append(idList, idInt)
	}

	if err := h.svc.DeleteDictData(c, idList...); err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, nil)
}

// GetDictData 获取字典数据详情
func (h *DictHandler) GetDictData(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "无效的ID"))
		return
	}

	dictData, err := h.svc.GetDictData(c, id)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	ginx.Success(c, dto.ToDictDataResponse(dictData))
}

// ListDictData 获取字典数据列表
func (h *DictHandler) ListDictData(c *gin.Context) {
	var query model.DictDataQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		ginx.ParamError(c, err)
		return
	}

	list, total, err := h.svc.ListDictData(c, &query)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	var resp []*dto.DictDataResponse
	for _, item := range list {
		resp = append(resp, dto.ToDictDataResponse(item))
	}

	ginx.Success(c, gin.H{
		"list":  resp,
		"total": total,
	})
}

// GetDictDataByType 根据字典类型获取字典数据
func (h *DictHandler) GetDictDataByType(c *gin.Context) {
	typeCode := c.Param("type")
	if typeCode == "" {
		ginx.ParamError(c, errors.WithMsg(errors.InvalidParam, "字典类型编码不能为空"))
		return
	}

	list, err := h.svc.GetDictDataByType(c, typeCode)
	if err != nil {
		ginx.ServerError(c, err)
		return
	}

	var resp []*dto.DictDataResponse
	for _, item := range list {
		resp = append(resp, dto.ToDictDataResponse(item))
	}

	ginx.Success(c, resp)
}
