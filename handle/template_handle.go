package handle

import (
	"context"
	"errors"

	"github.com/kouleen/idl/kitex_gen/system"
	"github.com/kouleen/system-center/service"
)

func QueryTemplatePage(ctx context.Context, req *system.SystemTemplateRequest) (*system.SystemTemplatePageResponse, error) {
	return service.QueryTemplatePage(ctx, req)
}

func QueryTemplate(ctx context.Context, systemTemplateRequest *system.SystemTemplateRequest) (resp *system.SystemTemplateResponse, err error) {
	return service.QueryTemplate(ctx, systemTemplateRequest)
}

func SaveTemplate(ctx context.Context, req *system.SystemTemplateRequest) (bool, error) {
	if req.TemplateName == "" {
		return false, errors.New("empty template name")
	}
	if req.TemplateType == "" {
		return false, errors.New("empty template type")
	}
	if req.TemplateContent == "" {
		return false, errors.New("empty template content")
	}
	return service.CreateTemplate(ctx, req)
}

func UpdateTemplate(ctx context.Context, req *system.SystemTemplateRequest) (resp bool, err error) {
	if req.Id != nil {
		return false, errors.New("template id should not be nil")
	}
	if req.TemplateCode == "" {
		return false, errors.New("empty template code")
	}
	if req.TemplateName == "" {
		return false, errors.New("empty template name")
	}
	if req.TemplateType == "" {
		return false, errors.New("empty template type")
	}
	if req.TemplateContent == "" {
		return false, errors.New("empty template content")
	}
	return service.UpdateTemplate(ctx, req)
}

func DeleteTemplate(ctx context.Context, req *system.SystemTemplateRequest) (resp bool, err error) {
	if req.Id == nil {
		return false, errors.New("template id should not be nil")
	}
	if req.IsDelete == nil {
		return false, errors.New("template isDelete should not be nil")
	}
	return service.DeleteTemplate(ctx, req)
}
