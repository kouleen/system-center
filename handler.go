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

// QuerySystemDictHeader implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemDictHeader(ctx context.Context, systemDictHeaderRequest *system.SystemDictHeaderRequest) (resp *system.SystemDictHeaderResponse, err error) {
	return handle.QueryDictHeader(ctx, systemDictHeaderRequest)
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

// QuerySystemDictLine implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemDictLine(ctx context.Context, systemDictLineRequest *system.SystemDictLineRequest) (resp *system.SystemDictLineResponse, err error) {
	return handle.QueryDictLine(ctx, systemDictLineRequest)
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

// QuerySystemInterface implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemInterface(ctx context.Context, systemInterfaceRequest *system.SystemInterfaceRequest) (resp *system.SystemInterfaceResponse, err error) {
	return handle.QueryInterface(ctx, systemInterfaceRequest)
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

// QuerySystemMenuTree implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemMenuTree(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	return handle.QueryMenuTree(ctx, systemMenuRequest)
}

// QuerySystemMenuList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemMenuList(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp []*system.SystemMenuResponse, err error) {
	return handle.QueryMenuList(ctx, systemMenuRequest)
}

// QuerySystemMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemMenu(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp *system.SystemMenuResponse, err error) {
	return handle.QueryMenu(ctx, systemMenuRequest)
}

// SaveSystemMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemMenu(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp bool, err error) {
	return handle.SaveMenu(ctx, systemMenuRequest)
}

// UpdateSystemMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemMenu(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp bool, err error) {
	return handle.UpdateMenu(ctx, systemMenuRequest)
}

// DeleteSystemMenu implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemMenu(ctx context.Context, systemMenuRequest *system.SystemMenuRequest) (resp bool, err error) {
	return handle.DeleteMenu(ctx, systemMenuRequest)
}

// QuerySystemRolePage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemRolePage(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp *system.SystemRolePageResponse, err error) {
	return handle.QueryRolePage(ctx, systemRoleRequest)
}

// QuerySystemRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemRole(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp *system.SystemRoleResponse, err error) {
	return handle.QueryRole(ctx, systemRoleRequest)
}

// SaveSystemRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemRole(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp bool, err error) {
	return handle.SaveRole(ctx, systemRoleRequest)
}

// UpdateSystemRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemRole(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp bool, err error) {
	return handle.UpdateRole(ctx, systemRoleRequest)
}

func (s *SystemServiceImpl) UpdateSystemRoleStatus(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp bool, err error) {
	return handle.UpdateRoleStatus(ctx, systemRoleRequest)
}

// DeleteSystemRole implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemRole(ctx context.Context, systemRoleRequest *system.SystemRoleRequest) (resp bool, err error) {
	return handle.DeleteRole(ctx, systemRoleRequest)
}

// QuerySystemRoleUserPage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemRoleUserPage(ctx context.Context, systemRoleUserRequest *system.SystemRoleUserRequest) (resp *system.SystemRoleUserPageResponse, err error) {
	return handle.QueryRoleUserPage(ctx, systemRoleUserRequest)
}

// QuerySystemRoleUserList implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemRoleUserList(ctx context.Context, systemRoleUserRequest *system.SystemRoleUserRequest) (resp []*system.SystemRoleUserResponse, err error) {
	return handle.QueryRoleUserList(ctx, systemRoleUserRequest)
}

// SaveSystemRoleUser implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemRoleUser(ctx context.Context, systemRoleUserRequest *system.SystemRoleUserRequest) (resp bool, err error) {
	return handle.SaveRoleUser(ctx, systemRoleUserRequest)
}

// CancelSystemRoleUser implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) CancelSystemRoleUser(ctx context.Context, systemRoleUserRequest *system.SystemRoleUserRequest) (resp bool, err error) {
	return handle.CancelRoleUser(ctx, systemRoleUserRequest)
}

// QuerySystemTemplatePage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemTemplatePage(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp *system.SystemTemplatePageResponse, err error) {
	return handle.QueryTemplatePage(ctx, systemTemplateRequest)
}

// QuerySystemTemplate implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemTemplate(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp *system.SystemTemplateResponse, err error) {
	return handle.QueryTemplate(ctx, systemTemplateRequest)
}

// SaveSystemTemplate implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemTemplate(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp bool, err error) {
	return handle.SaveTemplate(ctx, systemTemplateRequest)
}

// UpdateSystemTemplate implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemTemplate(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp bool, err error) {
	return handle.UpdateTemplate(ctx, systemTemplateRequest)
}

// DeleteSystemTemplate implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemTemplate(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp bool, err error) {
	return handle.DeleteTemplate(ctx, systemTemplateRequest)
}

// QuerySystemBulletinPage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemBulletinPage(ctx context.Context, systemBulletinRequest *system.SystemBulletinRequest) (resp *system.SystemBulletinPageResponse, err error) {
	return handle.QueryBulletinPage(ctx, systemBulletinRequest)
}

// QuerySystemBulletin implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemBulletin(ctx context.Context, systemBulletinRequest *system.SystemBulletinRequest) (resp *system.SystemBulletinResponse, err error) {
	return handle.QueryBulletin(ctx, systemBulletinRequest)
}

// SaveSystemBulletin implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemBulletin(ctx context.Context, systemBulletinRequest *system.SystemBulletinRequest) (resp bool, err error) {
	return handle.SaveBulletin(ctx, systemBulletinRequest)
}

// UpdateSystemBulletin implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemBulletin(ctx context.Context, systemBulletinRequest *system.SystemBulletinRequest) (resp bool, err error) {
	return handle.UpdateBulletin(ctx, systemBulletinRequest)
}

// DeleteSystemBulletin implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemBulletin(ctx context.Context, systemBulletinRequest *system.SystemBulletinRequest) (resp bool, err error) {
	return handle.DeleteBulletin(ctx, systemBulletinRequest)
}

// QuerySystemEducationPage implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemEducationPage(ctx context.Context, systemEducationRequest *system.SystemEducationRequest) (resp *system.SystemEducationPageResponse, err error) {
	return handle.QueryEducationPage(ctx, systemEducationRequest)
}

// QuerySystemEducation implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) QuerySystemEducation(ctx context.Context, systemEducationRequest *system.SystemEducationRequest) (resp *system.SystemEducationResponse, err error) {
	return handle.QueryEducation(ctx, systemEducationRequest)
}

// SaveSystemEducation implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) SaveSystemEducation(ctx context.Context, systemEducationRequest *system.SystemEducationRequest) (resp bool, err error) {
	return handle.SaveEducation(ctx, systemEducationRequest)
}

// UpdateSystemEducation implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) UpdateSystemEducation(ctx context.Context, systemEducationRequest *system.SystemEducationRequest) (resp bool, err error) {
	return handle.UpdateEducation(ctx, systemEducationRequest)
}

// DeleteSystemEducation implements the SystemServiceImpl interface.
func (s *SystemServiceImpl) DeleteSystemEducation(ctx context.Context, systemEducationRequest *system.SystemEducationRequest) (resp bool, err error) {
	return handle.DeleteEducation(ctx, systemEducationRequest)
}
