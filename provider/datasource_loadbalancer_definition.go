package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/justtrackio/terraform-provider-gosoline/builder"
)

type LoadbalancerDefinitionData struct {
	Project              types.String `tfsdk:"project"`
	Environment          types.String `tfsdk:"environment"`
	Family               types.String `tfsdk:"family"`
	Application          types.String `tfsdk:"application"`
	LoadbalancerShortArn types.String `tfsdk:"loadbalancerShortArn"`
}

func (d LoadbalancerDefinitionData) AppId() builder.AppId {
	return builder.AppId{
		Project:     d.Project.Value,
		Environment: d.Environment.Value,
		Family:      d.Family.Value,
		Application: d.Application.Value,
	}
}

type LoadbalancerDefinitionDatasourceType struct {
}

func (a *LoadbalancerDefinitionDatasourceType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"project": {
				Type:     types.StringType,
				Required: true,
			},
			"environment": {
				Type:     types.StringType,
				Required: true,
			},
			"family": {
				Type:     types.StringType,
				Required: true,
			},
			"application": {
				Type:     types.StringType,
				Required: true,
			},
			"loadbalancerShortArn": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}, nil
}

func (a *LoadbalancerDefinitionDatasourceType) NewDataSource(_ context.Context, provider tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return &LoadbalancerDefinitionDatasource{
		metadataReader: provider.(*GosolineProvider).metadataReader,
	}, nil
}

type LoadbalancerDefinitionDatasource struct {
	metadataReader *builder.MetadataReader
}

func (a *LoadbalancerDefinitionDatasource) Read(ctx context.Context, request tfsdk.ReadDataSourceRequest, response *tfsdk.ReadDataSourceResponse) {
	state := &LoadbalancerDefinitionData{}

	diags := request.Config.Get(ctx, state)
	response.Diagnostics.Append(diags...)

	var err error
	var ecsClient *builder.EcsClient
	var targetGroups []builder.ElbTargetGroup

	if ecsClient, err = builder.NewEcsClient(ctx, state.AppId()); err != nil {
		response.Diagnostics.AddError("can not get ecs client", err.Error())
		return
	}

	if targetGroups, err = ecsClient.GetElbTargetGroups(ctx); err != nil {
		response.Diagnostics.AddError("can not get target groups", err.Error())
		return
	}

	if len(targetGroups) == 0 {
		response.Diagnostics.AddWarning("failed to get target group", "there was no target group in the ecs service")
		return
	}

	state.LoadbalancerShortArn = types.String{
		Value: targetGroups[0].LoadBalancer,
	}

	diags = response.State.Set(ctx, state)
	response.Diagnostics.Append(diags...)
}
