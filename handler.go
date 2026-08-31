package main

import (
	"context"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/handle"
)

// SystemServiceImpl implements the last service interface defined in the IDL.
type SystemServiceImpl struct{}

// QuerySystemDictHeaderPage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemDictHeaderPage(ctx context.Context, systemDictHeaderRequest *system.SystemDictHeaderRequest) (resp *system.SystemDictHeaderPageResponse, err error) {
	return handle.QueryDictHeaderPage(ctx, systemDictHeaderRequest)
}

// QuerySystemDictHeaderList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemDictHeaderList(ctx context.Context, systemDictHeaderRequest *system.SystemDictHeaderRequest) (resp []*system.SystemDictHeaderResponse, err error) {
	return handle.QueryDictHeaderList(ctx, systemDictHeaderRequest)
}

// CreateSystemDictHeader implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateSystemDictHeader(ctx context.Context, systemDictHeaderRequest *system.SystemDictHeaderRequest) (resp bool, err error) {
	return handle.CreateDictHeader(ctx, systemDictHeaderRequest)
}

// UpdateSystemDictHeader implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemDictHeader(ctx context.Context, systemDictHeaderRequest *system.SystemDictHeaderRequest) (resp bool, err error) {
	return handle.UpdateDictHeader(ctx, systemDictHeaderRequest)
}

// DeleteSystemDictHeader implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemDictHeader(ctx context.Context, systemDictHeaderRequest *system.SystemDictHeaderRequest) (resp bool, err error) {
	return handle.DeleteDictHeader(ctx, systemDictHeaderRequest)
}

// QuerySystemDictLinePage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemDictLinePage(ctx context.Context, systemDictLineRequest *system.SystemDictLineRequest) (resp *system.SystemDictLinePageResponse, err error) {
	return handle.QueryDictLinePage(ctx, systemDictLineRequest)
}

// QuerySystemDictLineList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemDictLineList(ctx context.Context, systemDictLineRequest *system.SystemDictLineRequest) (resp []*system.SystemDictLineResponse, err error) {
	return handle.QueryDictLineList(ctx, systemDictLineRequest)
}

// CreateSystemDictLine implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CreateSystemDictLine(ctx context.Context, systemDictLineRequest *system.SystemDictLineRequest) (resp bool, err error) {
	return handle.CreateDictLine(ctx, systemDictLineRequest)
}

// UpdateSystemDictLine implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemDictLine(ctx context.Context, systemDictLineRequest *system.SystemDictLineRequest) (resp bool, err error) {
	return handle.UpdateDictLine(ctx, systemDictLineRequest)
}

// DeleteSystemDictLine implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemDictLine(ctx context.Context, systemDictLineRequest *system.SystemDictLineRequest) (resp bool, err error) {
	return handle.DeleteDictLine(ctx, systemDictLineRequest)
}

// QuerySystemInterfacePage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemInterfacePage(ctx context.Context, systemInterfaceRequest *system.SystemInterfaceRequest) (resp *system.SystemInterfacePageResponse, err error) {
	return handle.QueryInterfacePage(ctx, systemInterfaceRequest)
}

// QuerySystemInterfaceList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemInterfaceList(ctx context.Context, systemInterfaceRequest *system.SystemInterfaceRequest) (resp []*system.SystemInterfaceResponse, err error) {
	return handle.QueryInterfaceList(ctx, systemInterfaceRequest)
}

// SaveSystemInterface implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemInterface(ctx context.Context, systemInterfaceRequest *system.SystemInterfaceRequest) (resp bool, err error) {
	return handle.SaveInterface(ctx, systemInterfaceRequest)
}

// UpdateSystemInterface implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemInterface(ctx context.Context, systemInterfaceRequest *system.SystemInterfaceRequest) (resp bool, err error) {
	return handle.UpdateInterface(ctx, systemInterfaceRequest)
}

// DeleteSystemInterface implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemInterface(ctx context.Context, systemInterfaceRequest *system.SystemInterfaceRequest) (resp bool, err error) {
	return handle.DeleteInterface(ctx, systemInterfaceRequest)
}

// CleanCacheSystemInterface implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CleanCacheSystemInterface(ctx context.Context, systemInterfaceRequest *system.SystemInterfaceRequest) (resp bool, err error) {
	return handle.CleanCacheInterface(ctx, systemInterfaceRequest)
}

// QuerySystemMenuPage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemMenuPage(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp *system.SystemMenuPageResponse, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemMenuList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemMenuList(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemMenuAllTree implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemMenuAllTree(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemMenuTree implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemMenuTree(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemMenu(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp *system.SystemMenuResponse, err error) {
	// TODO: Your code here...
	return
}

// SaveSystemMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemMenu(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// UpdateSystemMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemMenu(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// DeleteSystemMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemMenu(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemRolePage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemRolePage(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp *system.SystemRolePageResponse, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemRole(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp *system.SystemRoleResponse, err error) {
	// TODO: Your code here...
	return
}

// SaveSystemRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemRole(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// UpdateSystemRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemRole(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// DeleteSystemRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemRole(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// SaveSystemRoleMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemRoleMenu(ctx context.Context, systemRoleMenuRequest *system.SystemRoleMenuRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// UpdateSystemRoleMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemRoleMenu(ctx context.Context, systemRoleMenuRequest *system.SystemRoleMenuRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemRoleUserPage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemRoleUserPage(ctx context.Context, systemRoleUserRequest *system.SystemRoleUserRequest) (resp *system.SystemRoleUserPageResponse, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemRoleUserList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemRoleUserList(ctx context.Context, systemRoleUserRequest *system.SystemRoleUserRequest) (resp []*system.SystemRoleUserResponse, err error) {
	// TODO: Your code here...
	return
}

// SaveSystemRoleUser implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemRoleUser(ctx context.Context, systemRoleUserRequest *system.SystemRoleUserRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// CancelSystemRoleUser implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CancelSystemRoleUser(ctx context.Context, systemRoleUserRequest *system.SystemRoleUserRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// QuerySystemTemplatePage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemTemplatePage(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp *system.SystemTemplatePageResponse, err error) {
	// TODO: Your code here...
	return
}

// SaveSystemTemplate implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemTemplate(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// UpdateSystemTemplate implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemTemplate(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// DeleteSystemTemplate implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemTemplate(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}
