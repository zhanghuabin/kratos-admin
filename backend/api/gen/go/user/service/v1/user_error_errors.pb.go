// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package servicev1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// common error
func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

// common error
func ErrorMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, UserErrorReason_METHOD_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

// 408
func IsRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_REQUEST_TIMEOUT.String() && e.Code == 408
}

// 408
func ErrorRequestTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(408, UserErrorReason_REQUEST_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 500
func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INTERNAL_SERVER_ERROR.String() && e.Code == 500
}

// 500
func ErrorInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, UserErrorReason_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

// 501
func IsNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NOT_IMPLEMENTED.String() && e.Code == 501
}

// 501
func ErrorNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, UserErrorReason_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

// 502
func IsNetworkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NETWORK_ERROR.String() && e.Code == 502
}

// 502
func ErrorNetworkError(format string, args ...interface{}) *errors.Error {
	return errors.New(502, UserErrorReason_NETWORK_ERROR.String(), fmt.Sprintf(format, args...))
}

// 503
func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

// 503
func ErrorServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return errors.New(503, UserErrorReason_SERVICE_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}

// 504
func IsNetworkTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NETWORK_TIMEOUT.String() && e.Code == 504
}

// 504
func ErrorNetworkTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, UserErrorReason_NETWORK_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 505
func IsRequestNotSupport(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_REQUEST_NOT_SUPPORT.String() && e.Code == 505
}

// 505
func ErrorRequestNotSupport(format string, args ...interface{}) *errors.Error {
	return errors.New(505, UserErrorReason_REQUEST_NOT_SUPPORT.String(), fmt.Sprintf(format, args...))
}

// 400
func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_BAD_REQUEST.String() && e.Code == 400
}

// 400
func ErrorBadRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserErrorReason_BAD_REQUEST.String(), fmt.Sprintf(format, args...))
}

// 用户ID无效
func IsInvalidUserid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INVALID_USERID.String() && e.Code == 400
}

// 用户ID无效
func ErrorInvalidUserid(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserErrorReason_INVALID_USERID.String(), fmt.Sprintf(format, args...))
}

// 密码无效
func IsInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INVALID_PASSWORD.String() && e.Code == 400
}

// 密码无效
func ErrorInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(400, UserErrorReason_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

// 404
func IsResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_RESOURCE_NOT_FOUND.String() && e.Code == 404
}

// 404
func ErrorResourceNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_RESOURCE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 用户不存在
func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

// 用户不存在
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 角色不存在
func IsRoleNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_ROLE_NOT_FOUND.String() && e.Code == 404
}

// 角色不存在
func ErrorRoleNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_ROLE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 部门不存在
func IsDepartmentNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_DEPARTMENT_NOT_FOUND.String() && e.Code == 404
}

// 部门不存在
func ErrorDepartmentNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_DEPARTMENT_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 组织不存在
func IsOrganizationNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_ORGANIZATION_NOT_FOUND.String() && e.Code == 404
}

// 组织不存在
func ErrorOrganizationNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_ORGANIZATION_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 职位不存在
func IsPositionNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_POSITION_NOT_FOUND.String() && e.Code == 404
}

// 职位不存在
func ErrorPositionNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_POSITION_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 401
func IsNotLoggedIn(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NOT_LOGGED_IN.String() && e.Code == 401
}

// 401
func ErrorNotLoggedIn(format string, args ...interface{}) *errors.Error {
	return errors.New(401, UserErrorReason_NOT_LOGGED_IN.String(), fmt.Sprintf(format, args...))
}

// 用户被冻结
func IsUserFreeze(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_USER_FREEZE.String() && e.Code == 401
}

// 用户被冻结
func ErrorUserFreeze(format string, args ...interface{}) *errors.Error {
	return errors.New(401, UserErrorReason_USER_FREEZE.String(), fmt.Sprintf(format, args...))
}

// 密码错误
func IsIncorrectPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INCORRECT_PASSWORD.String() && e.Code == 401
}

// 密码错误
func ErrorIncorrectPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(401, UserErrorReason_INCORRECT_PASSWORD.String(), fmt.Sprintf(format, args...))
}

// 403
func IsAccessForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_ACCESS_FORBIDDEN.String() && e.Code == 403
}

// 403
func ErrorAccessForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, UserErrorReason_ACCESS_FORBIDDEN.String(), fmt.Sprintf(format, args...))
}
