// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of route entries
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.GetRouteEntries(ctx, &vpc.GetRouteEntriesArgs{
//				Ids:          []interface{}{},
//				RouteTableId: "vtb-274e19skkuhog7fap8u4i8ird",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetRouteEntries(ctx *pulumi.Context, args *GetRouteEntriesArgs, opts ...pulumi.InvokeOption) (*GetRouteEntriesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRouteEntriesResult
	err := ctx.Invoke("volcengine:vpc/getRouteEntries:getRouteEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteEntries.
type GetRouteEntriesArgs struct {
	// A destination CIDR block of route entry.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// A list of route entry ids.
	Ids []string `pulumi:"ids"`
	// An id of next hop.
	NextHopId *string `pulumi:"nextHopId"`
	// A type of next hop, Optional choice contains `Instance`, `NetworkInterface`, `NatGW`, `VpnGW`.
	NextHopType *string `pulumi:"nextHopType"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// A name of route entry.
	RouteEntryName *string `pulumi:"routeEntryName"`
	// A type of route entry.
	RouteEntryType *string `pulumi:"routeEntryType"`
	// An id of route table.
	RouteTableId string `pulumi:"routeTableId"`
}

// A collection of values returned by getRouteEntries.
type GetRouteEntriesResult struct {
	// The destination CIDR block of the route entry.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The id of the next hop.
	NextHopId *string `pulumi:"nextHopId"`
	// The type of the next hop.
	NextHopType *string `pulumi:"nextHopType"`
	OutputFile  *string `pulumi:"outputFile"`
	// The collection of route tables.
	RouteEntries []GetRouteEntriesRouteEntry `pulumi:"routeEntries"`
	// The name of the route entry.
	RouteEntryName *string `pulumi:"routeEntryName"`
	RouteEntryType *string `pulumi:"routeEntryType"`
	// The id of the route table to which the route entry belongs.
	RouteTableId string `pulumi:"routeTableId"`
	// The total count of route entry query.
	TotalCount int `pulumi:"totalCount"`
}

func GetRouteEntriesOutput(ctx *pulumi.Context, args GetRouteEntriesOutputArgs, opts ...pulumi.InvokeOption) GetRouteEntriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetRouteEntriesResult, error) {
			args := v.(GetRouteEntriesArgs)
			r, err := GetRouteEntries(ctx, &args, opts...)
			var s GetRouteEntriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetRouteEntriesResultOutput)
}

// A collection of arguments for invoking getRouteEntries.
type GetRouteEntriesOutputArgs struct {
	// A destination CIDR block of route entry.
	DestinationCidrBlock pulumi.StringPtrInput `pulumi:"destinationCidrBlock"`
	// A list of route entry ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// An id of next hop.
	NextHopId pulumi.StringPtrInput `pulumi:"nextHopId"`
	// A type of next hop, Optional choice contains `Instance`, `NetworkInterface`, `NatGW`, `VpnGW`.
	NextHopType pulumi.StringPtrInput `pulumi:"nextHopType"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A name of route entry.
	RouteEntryName pulumi.StringPtrInput `pulumi:"routeEntryName"`
	// A type of route entry.
	RouteEntryType pulumi.StringPtrInput `pulumi:"routeEntryType"`
	// An id of route table.
	RouteTableId pulumi.StringInput `pulumi:"routeTableId"`
}

func (GetRouteEntriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouteEntriesArgs)(nil)).Elem()
}

// A collection of values returned by getRouteEntries.
type GetRouteEntriesResultOutput struct{ *pulumi.OutputState }

func (GetRouteEntriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRouteEntriesResult)(nil)).Elem()
}

func (o GetRouteEntriesResultOutput) ToGetRouteEntriesResultOutput() GetRouteEntriesResultOutput {
	return o
}

func (o GetRouteEntriesResultOutput) ToGetRouteEntriesResultOutputWithContext(ctx context.Context) GetRouteEntriesResultOutput {
	return o
}

// The destination CIDR block of the route entry.
func (o GetRouteEntriesResultOutput) DestinationCidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) *string { return v.DestinationCidrBlock }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRouteEntriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRouteEntriesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The id of the next hop.
func (o GetRouteEntriesResultOutput) NextHopId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) *string { return v.NextHopId }).(pulumi.StringPtrOutput)
}

// The type of the next hop.
func (o GetRouteEntriesResultOutput) NextHopType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) *string { return v.NextHopType }).(pulumi.StringPtrOutput)
}

func (o GetRouteEntriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of route tables.
func (o GetRouteEntriesResultOutput) RouteEntries() GetRouteEntriesRouteEntryArrayOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) []GetRouteEntriesRouteEntry { return v.RouteEntries }).(GetRouteEntriesRouteEntryArrayOutput)
}

// The name of the route entry.
func (o GetRouteEntriesResultOutput) RouteEntryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) *string { return v.RouteEntryName }).(pulumi.StringPtrOutput)
}

func (o GetRouteEntriesResultOutput) RouteEntryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) *string { return v.RouteEntryType }).(pulumi.StringPtrOutput)
}

// The id of the route table to which the route entry belongs.
func (o GetRouteEntriesResultOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) string { return v.RouteTableId }).(pulumi.StringOutput)
}

// The total count of route entry query.
func (o GetRouteEntriesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetRouteEntriesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRouteEntriesResultOutput{})
}
