//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewAuthenticationService,
	NewUserService,
	NewMenuService,
	NewRouterService,
	NewTaskService,
	NewRoleService,
	NewOrganizationService,
	NewDepartmentService,
	NewPositionService,
	NewDictService,
	NewAdminLoginLogService,
	NewAdminOperationLogService,
	NewFileService,
	NewUEditorService,
)
