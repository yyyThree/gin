package output

import (
	"encoding/json"
	"errors"
	"gin/output/code"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type baseResponse struct {
	State int    `json:"state"`
	Msg   string `json:"msg"`
}

type SucResponse struct {
	*baseResponse
	Data interface{} `json:"data,omitempty"`
}

type ListResponse struct {
	*SucResponse
	Total int `json:"total"`
}

type ErrResponse struct {
	*baseResponse
	ErrorData interface{} `json:"error_data,omitempty"`
}

type ResponseI interface {
	setBaseInfo(*Status)
}

func (resp *baseResponse) setBaseInfo(status *Status) {
	resp.State = status.GetCode()
	resp.Msg = status.GetMessage()
}

func (resp *SucResponse) setBaseInfo(status *Status) {
	if resp.baseResponse == nil {
		resp.baseResponse = &baseResponse{}
	}
	resp.State = status.GetCode()
	resp.Msg = status.GetMessage()
}

func (resp *ListResponse) setBaseInfo(status *Status) {
	if resp.SucResponse == nil {
		resp.SucResponse = &SucResponse{
			baseResponse: &baseResponse{},
			Data:         nil,
		}
	}
	if resp.SucResponse.baseResponse == nil {
		resp.SucResponse.baseResponse = &baseResponse{}
	}

	resp.State = status.GetCode()
	resp.Msg = status.GetMessage()
}

func (resp *ErrResponse) setBaseInfo(status *Status) {
	if resp.baseResponse == nil {
		resp.baseResponse = &baseResponse{}
	}

	if len(status.GetDetails()) > 0 {
		resp.ErrorData = status.GetDetails()
	}
	resp.State = status.GetCode()
	resp.Msg = status.GetMessage()
}

// 统一输出方法
func Response(c *gin.Context, resp ResponseI, err error) {
	// 成功
	if err == nil {
		resp.setBaseInfo(Error(code.OK))
		c.JSON(http.StatusOK, resp)
		return
	}

	// err != nil
	resp = &ErrResponse{}
	status := &Status{}
	switch v := err.(type) {
	case *Status:
		status = v
	case *json.UnmarshalTypeError, *strconv.NumError:
		status = Error(code.ParamBindErr).WithDetails(v.Error())
	case *mysql.MySQLError:
		status = Error(code.MySqlErr).WithDetails(err)
	default:
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = Error(code.RecordNotFound)
			break
		}
		status = Error(code.ServerErr).WithDetails(err)
	}
	resp.setBaseInfo(status)

	// 输出不同的http状态码
	switch code.Code(status.GetCode()) {
	case code.ApiNotFound: // 404
		c.JSON(http.StatusNotFound, resp)
	case code.ServerErr: // 500
		c.JSON(http.StatusInternalServerError, resp)
	case code.ParamBindErr, code.IllegalParams: // 参数非法
		c.JSON(http.StatusBadRequest, resp)
	case code.NoAuthorization, code.AuthorizationErr: // Authorization非法
		c.JSON(http.StatusUnauthorized, resp)
	case code.TokenNotFound, code.TokenMalformed, code.TokenExpired, code.TokenNotValidYet, code.TokenNotValid: // token非法
		c.JSON(http.StatusUnauthorized, resp)
	default:
		c.JSON(http.StatusOK, resp)
	}
	return
}
