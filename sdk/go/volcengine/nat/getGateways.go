// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of nat gateways
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/nat"
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
// var fooGateway []*nat.Gateway
//
//	for index := 0; index < 3; index++ {
//	    key0 := index
//	    val0 := index
//
// __res, err := nat.NewGateway(ctx, fmt.Sprintf("fooGateway-%v", key0), &nat.GatewayArgs{
// VpcId: fooVpc.ID(),
// SubnetId: fooSubnet.ID(),
// Spec: pulumi.String("Small"),
// NatGatewayName: pulumi.String(fmt.Sprintf("acc-test-ng-%v", val0)),
// Description: pulumi.String("acc-test"),
// BillingType: pulumi.String("PostPaid"),
// ProjectName: pulumi.String("default"),
// Tags: nat.GatewayTagArray{
// &nat.GatewayTagArgs{
// Key: pulumi.String("k1"),
// Value: pulumi.String("v1"),
// },
// },
// })
// if err != nil {
// return err
// }
// fooGateway = append(fooGateway, __res)
// }
// _ = nat.GetGatewaysOutput(ctx, nat.GetGatewaysOutputArgs{
// Ids: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-functions-volcengine:nat-getGateways:getGateways.pp:30,9-25),
// }, nil);
// return nil
// })
// }
// ```
func GetGateways(ctx *pulumi.Context, args *GetGatewaysArgs, opts ...pulumi.InvokeOption) (*GetGatewaysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetGatewaysResult
	err := ctx.Invoke("volcengine:nat/getGateways:getGateways", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGateways.
type GetGatewaysArgs struct {
	// The description of the NatGateway.
	Description *string `pulumi:"description"`
	// The list of NatGateway IDs.
	Ids []string `pulumi:"ids"`
	// The Name Regex of NatGateway.
	NameRegex *string `pulumi:"nameRegex"`
	// The name of the NatGateway.
	NatGatewayName *string `pulumi:"natGatewayName"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The specification of the NatGateway.
	Spec *string `pulumi:"spec"`
	// The id of the Subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Tags.
	Tags []GetGatewaysTag `pulumi:"tags"`
	// The id of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getGateways.
type GetGatewaysResult struct {
	// The description of the NatGateway.
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id        string   `pulumi:"id"`
	Ids       []string `pulumi:"ids"`
	NameRegex *string  `pulumi:"nameRegex"`
	// The name of the NatGateway.
	NatGatewayName *string `pulumi:"natGatewayName"`
	// The collection of NatGateway query.
	NatGateways []GetGatewaysNatGateway `pulumi:"natGateways"`
	OutputFile  *string                 `pulumi:"outputFile"`
	// The specification of the NatGateway.
	Spec *string `pulumi:"spec"`
	// The ID of the Subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Tags.
	Tags []GetGatewaysTag `pulumi:"tags"`
	// The total count of NatGateway query.
	TotalCount int `pulumi:"totalCount"`
	// The ID of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

func GetGatewaysOutput(ctx *pulumi.Context, args GetGatewaysOutputArgs, opts ...pulumi.InvokeOption) GetGatewaysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetGatewaysResult, error) {
			args := v.(GetGatewaysArgs)
			r, err := GetGateways(ctx, &args, opts...)
			var s GetGatewaysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetGatewaysResultOutput)
}

// A collection of arguments for invoking getGateways.
type GetGatewaysOutputArgs struct {
	// The description of the NatGateway.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The list of NatGateway IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Name Regex of NatGateway.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The name of the NatGateway.
	NatGatewayName pulumi.StringPtrInput `pulumi:"natGatewayName"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The specification of the NatGateway.
	Spec pulumi.StringPtrInput `pulumi:"spec"`
	// The id of the Subnet.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Tags.
	Tags GetGatewaysTagArrayInput `pulumi:"tags"`
	// The id of the VPC.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (GetGatewaysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaysArgs)(nil)).Elem()
}

// A collection of values returned by getGateways.
type GetGatewaysResultOutput struct{ *pulumi.OutputState }

func (GetGatewaysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetGatewaysResult)(nil)).Elem()
}

func (o GetGatewaysResultOutput) ToGetGatewaysResultOutput() GetGatewaysResultOutput {
	return o
}

func (o GetGatewaysResultOutput) ToGetGatewaysResultOutputWithContext(ctx context.Context) GetGatewaysResultOutput {
	return o
}

// The description of the NatGateway.
func (o GetGatewaysResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetGatewaysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetGatewaysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetGatewaysResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetGatewaysResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetGatewaysResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

// The name of the NatGateway.
func (o GetGatewaysResultOutput) NatGatewayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.NatGatewayName }).(pulumi.StringPtrOutput)
}

// The collection of NatGateway query.
func (o GetGatewaysResultOutput) NatGateways() GetGatewaysNatGatewayArrayOutput {
	return o.ApplyT(func(v GetGatewaysResult) []GetGatewaysNatGateway { return v.NatGateways }).(GetGatewaysNatGatewayArrayOutput)
}

func (o GetGatewaysResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The specification of the NatGateway.
func (o GetGatewaysResultOutput) Spec() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.Spec }).(pulumi.StringPtrOutput)
}

// The ID of the Subnet.
func (o GetGatewaysResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Tags.
func (o GetGatewaysResultOutput) Tags() GetGatewaysTagArrayOutput {
	return o.ApplyT(func(v GetGatewaysResult) []GetGatewaysTag { return v.Tags }).(GetGatewaysTagArrayOutput)
}

// The total count of NatGateway query.
func (o GetGatewaysResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetGatewaysResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The ID of the VPC.
func (o GetGatewaysResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetGatewaysResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetGatewaysResultOutput{})
}
