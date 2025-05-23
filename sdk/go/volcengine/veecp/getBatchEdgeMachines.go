// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package veecp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of veecp batch edge machines
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/veecp"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooBatchEdgeMachine, err := veecp.NewBatchEdgeMachine(ctx, "fooBatchEdgeMachine", &veecp.BatchEdgeMachineArgs{
//				ClusterId:  pulumi.String("ccvd7mte6t101fno98u60"),
//				NodePoolId: pulumi.String("pcvd90uacnsr73g6bjic0"),
//				TtlHours:   pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_ = veecp.GetBatchEdgeMachinesOutput(ctx, veecp.GetBatchEdgeMachinesOutputArgs{
//				ClusterIds: pulumi.StringArray{
//					fooBatchEdgeMachine.ClusterId,
//				},
//				Ids: pulumi.StringArray{
//					fooBatchEdgeMachine.ID(),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetBatchEdgeMachines(ctx *pulumi.Context, args *GetBatchEdgeMachinesArgs, opts ...pulumi.InvokeOption) (*GetBatchEdgeMachinesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetBatchEdgeMachinesResult
	err := ctx.Invoke("volcengine:veecp/getBatchEdgeMachines:getBatchEdgeMachines", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBatchEdgeMachines.
type GetBatchEdgeMachinesArgs struct {
	// The ClusterIds of NodePool IDs.
	ClusterIds []string `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken *string `pulumi:"createClientToken"`
	// A list of IDs.
	Ids []string `pulumi:"ids"`
	// The IPs.
	Ips []string `pulumi:"ips"`
	// The Name of NodePool.
	Name *string `pulumi:"name"`
	// Whether it is necessary to query the node management script.
	NeedBootstrapScript *string `pulumi:"needBootstrapScript"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The Status of NodePool.
	Statuses []GetBatchEdgeMachinesStatus `pulumi:"statuses"`
	// The Zone Ids.
	ZoneIds []string `pulumi:"zoneIds"`
}

// A collection of values returned by getBatchEdgeMachines.
type GetBatchEdgeMachinesResult struct {
	ClusterIds []string `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	Ips []string `pulumi:"ips"`
	// The collection of query.
	Machines []GetBatchEdgeMachinesMachine `pulumi:"machines"`
	// The Name of NodePool.
	Name                *string                      `pulumi:"name"`
	NeedBootstrapScript *string                      `pulumi:"needBootstrapScript"`
	OutputFile          *string                      `pulumi:"outputFile"`
	Statuses            []GetBatchEdgeMachinesStatus `pulumi:"statuses"`
	// The total count of query.
	TotalCount int      `pulumi:"totalCount"`
	ZoneIds    []string `pulumi:"zoneIds"`
}

func GetBatchEdgeMachinesOutput(ctx *pulumi.Context, args GetBatchEdgeMachinesOutputArgs, opts ...pulumi.InvokeOption) GetBatchEdgeMachinesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetBatchEdgeMachinesResult, error) {
			args := v.(GetBatchEdgeMachinesArgs)
			r, err := GetBatchEdgeMachines(ctx, &args, opts...)
			var s GetBatchEdgeMachinesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetBatchEdgeMachinesResultOutput)
}

// A collection of arguments for invoking getBatchEdgeMachines.
type GetBatchEdgeMachinesOutputArgs struct {
	// The ClusterIds of NodePool IDs.
	ClusterIds pulumi.StringArrayInput `pulumi:"clusterIds"`
	// The ClientToken when successfully created.
	CreateClientToken pulumi.StringPtrInput `pulumi:"createClientToken"`
	// A list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The IPs.
	Ips pulumi.StringArrayInput `pulumi:"ips"`
	// The Name of NodePool.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Whether it is necessary to query the node management script.
	NeedBootstrapScript pulumi.StringPtrInput `pulumi:"needBootstrapScript"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Status of NodePool.
	Statuses GetBatchEdgeMachinesStatusArrayInput `pulumi:"statuses"`
	// The Zone Ids.
	ZoneIds pulumi.StringArrayInput `pulumi:"zoneIds"`
}

func (GetBatchEdgeMachinesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBatchEdgeMachinesArgs)(nil)).Elem()
}

// A collection of values returned by getBatchEdgeMachines.
type GetBatchEdgeMachinesResultOutput struct{ *pulumi.OutputState }

func (GetBatchEdgeMachinesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBatchEdgeMachinesResult)(nil)).Elem()
}

func (o GetBatchEdgeMachinesResultOutput) ToGetBatchEdgeMachinesResultOutput() GetBatchEdgeMachinesResultOutput {
	return o
}

func (o GetBatchEdgeMachinesResultOutput) ToGetBatchEdgeMachinesResultOutputWithContext(ctx context.Context) GetBatchEdgeMachinesResultOutput {
	return o
}

func (o GetBatchEdgeMachinesResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

// The ClientToken when successfully created.
func (o GetBatchEdgeMachinesResultOutput) CreateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) *string { return v.CreateClientToken }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetBatchEdgeMachinesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBatchEdgeMachinesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetBatchEdgeMachinesResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

// The collection of query.
func (o GetBatchEdgeMachinesResultOutput) Machines() GetBatchEdgeMachinesMachineArrayOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) []GetBatchEdgeMachinesMachine { return v.Machines }).(GetBatchEdgeMachinesMachineArrayOutput)
}

// The Name of NodePool.
func (o GetBatchEdgeMachinesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetBatchEdgeMachinesResultOutput) NeedBootstrapScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) *string { return v.NeedBootstrapScript }).(pulumi.StringPtrOutput)
}

func (o GetBatchEdgeMachinesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetBatchEdgeMachinesResultOutput) Statuses() GetBatchEdgeMachinesStatusArrayOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) []GetBatchEdgeMachinesStatus { return v.Statuses }).(GetBatchEdgeMachinesStatusArrayOutput)
}

// The total count of query.
func (o GetBatchEdgeMachinesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o GetBatchEdgeMachinesResultOutput) ZoneIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBatchEdgeMachinesResult) []string { return v.ZoneIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBatchEdgeMachinesResultOutput{})
}
