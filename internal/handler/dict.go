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
// @Summary 创建字典类型
// @Description 创建一个新的字典类型
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param data body dto.DictTypeRequest true "字典类型信息"
// @Success 200 {object} ginx.Response{data=dto.DictTypeResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/type [post]
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
// @Summary 更新字典类型
// @Description 更新指定ID的字典类型
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典类型ID"
// @Param data body dto.DictTypeRequest true "字典类型信息"
// @Success 200 {object} ginx.Response{data=dto.DictTypeResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "字典类型不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/type/{id} [put]
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
// @Summary 删除字典类型
// @Description 删除指定ID的字典类型
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param ids path string true "字典类型ID列表(多个用逗号分隔)"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/type/{ids} [delete]
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
// @Summary 获取字典类型详情
// @Description 获取指定ID的字典类型详情
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典类型ID"
// @Success 200 {object} ginx.Response{data=dto.DictTypeResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "字典类型不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/type/{id} [get]
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
// @Summary 获取字典类型列表
// @Description 分页获取字典类型列表
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param name query string false "字典类型名称"
// @Param code query string false "字典类型编码"
// @Param status query int false "状态(1:正常 2:禁用)"
// @Success 200 {object} ginx.Response{data=ginx.ListData{list=[]dto.DictTypeResponse,total=int64}} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/type [get]
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

	ginx.Success(c, ginx.ListData{
		List:  resp,
		Total: total,
	})
}

// CreateDictData 创建字典数据
// @Summary 创建字典数据
// @Description 创建一个新的字典数据
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param data body dto.DictDataRequest true "字典数据信息"
// @Success 200 {object} ginx.Response{data=dto.DictDataResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/data [post]
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
// @Summary 更新字典数据
// @Description 更新指定ID的字典数据
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典数据ID"
// @Param data body dto.DictDataRequest true "字典数据信息"
// @Success 200 {object} ginx.Response{data=dto.DictDataResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "字典数据不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/data/{id} [put]
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
// @Summary 删除字典数据
// @Description 删除指定ID的字典数据
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param ids path string true "字典数据ID列表(多个用逗号分隔)"
// @Success 200 {object} ginx.Response "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/data/{ids} [delete]
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
// @Summary 获取字典数据详情
// @Description 获取指定ID的字典数据详情
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param id path int true "字典数据ID"
// @Success 200 {object} ginx.Response{data=dto.DictDataResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 404 {object} ginx.Response "字典数据不存在"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/data/{id} [get]
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
// @Summary 获取字典数据列表
// @Description 分页获取字典数据列表
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Param type_code query string false "字典类型编码"
// @Param label query string false "字典标签"
// @Param status query int false "状态(1:正常 2:禁用)"
// @Success 200 {object} ginx.Response{data=ginx.ListData{list=[]dto.DictDataResponse,total=int64}} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/data [get]
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

	ginx.Success(c, ginx.ListData{
		List:  resp,
		Total: total,
	})
}

// GetDictDataByType 根据字典类型获取字典数据
// @Summary 根据字典类型获取字典数据
// @Description 根据字典类型编码获取字典数据列表
// @Tags 字典管理
// @Accept json
// @Produce json
// @Param type path string true "字典类型编码"
// @Success 200 {object} ginx.Response{data=[]dto.DictDataResponse} "成功"
// @Failure 400 {object} ginx.Response "请求参数错误"
// @Failure 500 {object} ginx.Response "服务器内部错误"
// @Security Bearer
// @Router /system/dict/data/type/{type} [get]
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
