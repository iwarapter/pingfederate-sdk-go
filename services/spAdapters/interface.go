package spAdapters

import (
	"context"
	"net/http"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
)

type SpAdaptersAPI interface {
	GetSpAdapterDescriptors() (output *models.SpAdapterDescriptors, resp *http.Response, err error)
	GetSpAdapterDescriptorsWithContext(ctx context.Context) (output *models.SpAdapterDescriptors, resp *http.Response, err error)

	GetSpAdapterDescriptorsById(input *GetSpAdapterDescriptorsByIdInput) (output *models.SpAdapterDescriptor, resp *http.Response, err error)
	GetSpAdapterDescriptorsByIdWithContext(ctx context.Context, input *GetSpAdapterDescriptorsByIdInput) (output *models.SpAdapterDescriptor, resp *http.Response, err error)

	GetSpAdapters(input *GetSpAdaptersInput) (output *models.SpAdapters, resp *http.Response, err error)
	GetSpAdaptersWithContext(ctx context.Context, input *GetSpAdaptersInput) (output *models.SpAdapters, resp *http.Response, err error)

	CreateSpAdapter(input *CreateSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error)
	CreateSpAdapterWithContext(ctx context.Context, input *CreateSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error)

	GetSpAdapter(input *GetSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error)
	GetSpAdapterWithContext(ctx context.Context, input *GetSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error)

	UpdateSpAdapter(input *UpdateSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error)
	UpdateSpAdapterWithContext(ctx context.Context, input *UpdateSpAdapterInput) (output *models.SpAdapter, resp *http.Response, err error)

	DeleteSpAdapter(input *DeleteSpAdapterInput) (output *models.ApiResult, resp *http.Response, err error)
	DeleteSpAdapterWithContext(ctx context.Context, input *DeleteSpAdapterInput) (output *models.ApiResult, resp *http.Response, err error)

	GetActions(input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)
	GetActionsWithContext(ctx context.Context, input *GetActionsInput) (output *models.Actions, resp *http.Response, err error)

	GetAction(input *GetActionInput) (output *models.Action, resp *http.Response, err error)
	GetActionWithContext(ctx context.Context, input *GetActionInput) (output *models.Action, resp *http.Response, err error)

	InvokeAction(input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)
	InvokeActionWithContext(ctx context.Context, input *InvokeActionInput) (output *models.ActionResult, resp *http.Response, err error)

	GetUrlMappings() (output *models.SpAdapterUrlMappings, resp *http.Response, err error)
	GetUrlMappingsWithContext(ctx context.Context) (output *models.SpAdapterUrlMappings, resp *http.Response, err error)

	UpdateUrlMappings(input *UpdateUrlMappingsInput) (output *models.SpAdapterUrlMappings, resp *http.Response, err error)
	UpdateUrlMappingsWithContext(ctx context.Context, input *UpdateUrlMappingsInput) (output *models.SpAdapterUrlMappings, resp *http.Response, err error)
}
