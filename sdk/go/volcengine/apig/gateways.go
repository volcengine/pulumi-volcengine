// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of apig gateways
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/apig"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := apig.GetGateways(ctx, &apig.GetGatewaysArgs{
//				Ids: []string{
//					"gd13d8c6eq1emkiunq6p0",
//					"gd07fq7pte3scmnaj1b1g",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.apig.Gateways has been deprecated in favor of volcengine.apig.getGateways
func Gateways(ctx *pulumi.Context, args *GatewaysArgs, opts ...pulumi.InvokeOption) (*GatewaysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GatewaysResult
	err := ctx.Invoke("volcengine:apig/gateways:Gateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Gateways.
type GatewaysArgs struct {
	// A list of api gateway IDs.
	Ids []string `pulumi:"ids"`
	// The name of api gateway. This field support fuzzy query.
	Name *string `pulumi:"name"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The project name of api gateway.
	ProjectName *string `pulumi:"projectName"`
	// The status of api gateway.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []GatewaysTag `pulumi:"tags"`
	// The type of api gateway.
	Type *string `pulumi:"type"`
	// A list of vpc IDs.
	VpcIds []string `pulumi:"vpcIds"`
}

// A collection of values returned by Gateways.
type GatewaysResult struct {
	// The collection of query.
	Gateways []GatewaysGateway `pulumi:"gateways"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The name of the api gateway.
	Name       *string `pulumi:"name"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The project name of the api gateway.
	ProjectName *string `pulumi:"projectName"`
	// The status of the api gateway.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []GatewaysTag `pulumi:"tags"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
	// The type of the api gateway.
	Type   *string  `pulumi:"type"`
	VpcIds []string `pulumi:"vpcIds"`
}

func GatewaysOutput(ctx *pulumi.Context, args GatewaysOutputArgs, opts ...pulumi.InvokeOption) GatewaysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GatewaysResult, error) {
			args := v.(GatewaysArgs)
			r, err := Gateways(ctx, &args, opts...)
			var s GatewaysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GatewaysResultOutput)
}

// A collection of arguments for invoking Gateways.
type GatewaysOutputArgs struct {
	// A list of api gateway IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The name of api gateway. This field support fuzzy query.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The project name of api gateway.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// The status of api gateway.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Tags.
	Tags GatewaysTagArrayInput `pulumi:"tags"`
	// The type of api gateway.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// A list of vpc IDs.
	VpcIds pulumi.StringArrayInput `pulumi:"vpcIds"`
}

func (GatewaysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewaysArgs)(nil)).Elem()
}

// A collection of values returned by Gateways.
type GatewaysResultOutput struct{ *pulumi.OutputState }

func (GatewaysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewaysResult)(nil)).Elem()
}

func (o GatewaysResultOutput) ToGatewaysResultOutput() GatewaysResultOutput {
	return o
}

func (o GatewaysResultOutput) ToGatewaysResultOutputWithContext(ctx context.Context) GatewaysResultOutput {
	return o
}

// The collection of query.
func (o GatewaysResultOutput) Gateways() GatewaysGatewayArrayOutput {
	return o.ApplyT(func(v GatewaysResult) []GatewaysGateway { return v.Gateways }).(GatewaysGatewayArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GatewaysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GatewaysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GatewaysResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewaysResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The name of the api gateway.
func (o GatewaysResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaysResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GatewaysResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaysResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GatewaysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The project name of the api gateway.
func (o GatewaysResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaysResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The status of the api gateway.
func (o GatewaysResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaysResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Tags.
func (o GatewaysResultOutput) Tags() GatewaysTagArrayOutput {
	return o.ApplyT(func(v GatewaysResult) []GatewaysTag { return v.Tags }).(GatewaysTagArrayOutput)
}

// The total count of query.
func (o GatewaysResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GatewaysResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The type of the api gateway.
func (o GatewaysResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GatewaysResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GatewaysResultOutput) VpcIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GatewaysResult) []string { return v.VpcIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewaysResultOutput{})
}
