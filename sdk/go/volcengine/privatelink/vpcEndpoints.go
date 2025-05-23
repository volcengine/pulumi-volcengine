// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of privatelink vpc endpoints
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/privatelink"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// fooZones, err := ecs.GetZones(ctx, nil, nil);
// if err != nil {
// return err
// }
// fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
// VpcName: pulumi.String("acc-test-vpc"),
// CidrBlock: pulumi.String("172.16.0.0/16"),
// })
// if err != nil {
// return err
// }
// fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
// SubnetName: pulumi.String("acc-test-subnet"),
// CidrBlock: pulumi.String("172.16.0.0/24"),
// ZoneId: pulumi.String(fooZones.Zones[0].Id),
// VpcId: fooVpc.ID(),
// })
// if err != nil {
// return err
// }
// fooSecurityGroup, err := vpc.NewSecurityGroup(ctx, "fooSecurityGroup", &vpc.SecurityGroupArgs{
// SecurityGroupName: pulumi.String("acc-test-security-group"),
// VpcId: fooVpc.ID(),
// })
// if err != nil {
// return err
// }
// fooClb, err := clb.NewClb(ctx, "fooClb", &clb.ClbArgs{
// Type: pulumi.String("public"),
// SubnetId: fooSubnet.ID(),
// LoadBalancerSpec: pulumi.String("small_1"),
// Description: pulumi.String("acc-test-demo"),
// LoadBalancerName: pulumi.String("acc-test-clb"),
// LoadBalancerBillingType: pulumi.String("PostPaid"),
// EipBillingConfig: &clb.ClbEipBillingConfigArgs{
// Isp: pulumi.String("BGP"),
// EipBillingType: pulumi.String("PostPaidByBandwidth"),
// Bandwidth: pulumi.Int(1),
// },
// Tags: clb.ClbTagArray{
// &clb.ClbTagArgs{
// Key: pulumi.String("k1"),
// Value: pulumi.String("v1"),
// },
// },
// })
// if err != nil {
// return err
// }
// fooVpcEndpointService, err := privatelink.NewVpcEndpointService(ctx, "fooVpcEndpointService", &privatelink.VpcEndpointServiceArgs{
// Resources: privatelink.VpcEndpointServiceResourceTypeArray{
// &privatelink.VpcEndpointServiceResourceTypeArgs{
// ResourceId: fooClb.ID(),
// ResourceType: pulumi.String("CLB"),
// },
// },
// Description: pulumi.String("acc-test"),
// AutoAcceptEnabled: pulumi.Bool(true),
// })
// if err != nil {
// return err
// }
// var fooVpcEndpoint []*privatelink.VpcEndpoint
//
//	for index := 0; index < 2; index++ {
//	    key0 := index
//	    _ := index
//
// __res, err := privatelink.NewVpcEndpoint(ctx, fmt.Sprintf("fooVpcEndpoint-%v", key0), &privatelink.VpcEndpointArgs{
// SecurityGroupIds: pulumi.StringArray{
// fooSecurityGroup.ID(),
// },
// ServiceId: fooVpcEndpointService.ID(),
// EndpointName: pulumi.String("acc-test-ep"),
// Description: pulumi.String("acc-test"),
// })
// if err != nil {
// return err
// }
// fooVpcEndpoint = append(fooVpcEndpoint, __res)
// }
// _ = privatelink.GetVpcEndpointsOutput(ctx, privatelink.GetVpcEndpointsOutputArgs{
// Ids: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-functions-volcengine:privatelink-vpcEndpoints:VpcEndpoints.pp:52,9-29),
// }, nil);
// return nil
// })
// }
// ```
//
// Deprecated: volcengine.privatelink.VpcEndpoints has been deprecated in favor of volcengine.privatelink.getVpcEndpoints
func VpcEndpoints(ctx *pulumi.Context, args *VpcEndpointsArgs, opts ...pulumi.InvokeOption) (*VpcEndpointsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv VpcEndpointsResult
	err := ctx.Invoke("volcengine:privatelink/vpcEndpoints:VpcEndpoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking VpcEndpoints.
type VpcEndpointsArgs struct {
	// The name of vpc endpoint.
	EndpointName *string `pulumi:"endpointName"`
	// The IDs of vpc endpoint.
	Ids []string `pulumi:"ids"`
	// A Name Regex of vpc endpoint.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of vpc endpoint service.
	ServiceName *string `pulumi:"serviceName"`
	// The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
	Status *string `pulumi:"status"`
	// The vpc id of vpc endpoint.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by VpcEndpoints.
type VpcEndpointsResult struct {
	// The name of vpc endpoint.
	EndpointName *string `pulumi:"endpointName"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The name of vpc endpoint service.
	ServiceName *string `pulumi:"serviceName"`
	// The status of vpc endpoint.
	Status *string `pulumi:"status"`
	// Returns the total amount of the data list.
	TotalCount int `pulumi:"totalCount"`
	// The collection of query.
	VpcEndpoints []VpcEndpointsVpcEndpoint `pulumi:"vpcEndpoints"`
	// The vpc id of vpc endpoint.
	VpcId *string `pulumi:"vpcId"`
}

func VpcEndpointsOutput(ctx *pulumi.Context, args VpcEndpointsOutputArgs, opts ...pulumi.InvokeOption) VpcEndpointsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VpcEndpointsResult, error) {
			args := v.(VpcEndpointsArgs)
			r, err := VpcEndpoints(ctx, &args, opts...)
			var s VpcEndpointsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VpcEndpointsResultOutput)
}

// A collection of arguments for invoking VpcEndpoints.
type VpcEndpointsOutputArgs struct {
	// The name of vpc endpoint.
	EndpointName pulumi.StringPtrInput `pulumi:"endpointName"`
	// The IDs of vpc endpoint.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of vpc endpoint.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of vpc endpoint service.
	ServiceName pulumi.StringPtrInput `pulumi:"serviceName"`
	// The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// The vpc id of vpc endpoint.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (VpcEndpointsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointsArgs)(nil)).Elem()
}

// A collection of values returned by VpcEndpoints.
type VpcEndpointsResultOutput struct{ *pulumi.OutputState }

func (VpcEndpointsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointsResult)(nil)).Elem()
}

func (o VpcEndpointsResultOutput) ToVpcEndpointsResultOutput() VpcEndpointsResultOutput {
	return o
}

func (o VpcEndpointsResultOutput) ToVpcEndpointsResultOutputWithContext(ctx context.Context) VpcEndpointsResultOutput {
	return o
}

// The name of vpc endpoint.
func (o VpcEndpointsResultOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointsResult) *string { return v.EndpointName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o VpcEndpointsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VpcEndpointsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VpcEndpointsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpcEndpointsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o VpcEndpointsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o VpcEndpointsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of vpc endpoint service.
func (o VpcEndpointsResultOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointsResult) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

// The status of vpc endpoint.
func (o VpcEndpointsResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointsResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Returns the total amount of the data list.
func (o VpcEndpointsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v VpcEndpointsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The collection of query.
func (o VpcEndpointsResultOutput) VpcEndpoints() VpcEndpointsVpcEndpointArrayOutput {
	return o.ApplyT(func(v VpcEndpointsResult) []VpcEndpointsVpcEndpoint { return v.VpcEndpoints }).(VpcEndpointsVpcEndpointArrayOutput)
}

// The vpc id of vpc endpoint.
func (o VpcEndpointsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcEndpointsResultOutput{})
}
