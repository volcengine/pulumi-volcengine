// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of snat entries
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/nat"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := nat.SnatEntries(ctx, &nat.SnatEntriesArgs{
//				Ids: []string{
//					"snat-274zl8b1kxzb47fap8u35uune",
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
func SnatEntries(ctx *pulumi.Context, args *SnatEntriesArgs, opts ...pulumi.InvokeOption) (*SnatEntriesResult, error) {
	var rv SnatEntriesResult
	err := ctx.Invoke("volcengine:nat/snatEntries:SnatEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking SnatEntries.
type SnatEntriesArgs struct {
	// An id of the public ip address used by the SNAT entry.
	EipId *string `pulumi:"eipId"`
	// A list of SNAT entry ids.
	Ids []string `pulumi:"ids"`
	// An id of the nat gateway to which the entry belongs.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// A name of SNAT entry.
	SnatEntryName *string `pulumi:"snatEntryName"`
	// The SourceCidr of SNAT entry.
	SourceCidr *string `pulumi:"sourceCidr"`
	// An id of the subnet that is required to access the Internet.
	SubnetId *string `pulumi:"subnetId"`
}

// A collection of values returned by SnatEntries.
type SnatEntriesResult struct {
	// The id of the public ip address used by the SNAT entry.
	EipId *string `pulumi:"eipId"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId *string `pulumi:"natGatewayId"`
	OutputFile   *string `pulumi:"outputFile"`
	// The collection of snat entries.
	SnatEntries []SnatEntriesSnatEntry `pulumi:"snatEntries"`
	// The name of the SNAT entry.
	SnatEntryName *string `pulumi:"snatEntryName"`
	// The SourceCidr of the SNAT entry.
	SourceCidr *string `pulumi:"sourceCidr"`
	// The id of the subnet that is required to access the internet.
	SubnetId *string `pulumi:"subnetId"`
	// The total count of snat entries query.
	TotalCount int `pulumi:"totalCount"`
}

func SnatEntriesOutput(ctx *pulumi.Context, args SnatEntriesOutputArgs, opts ...pulumi.InvokeOption) SnatEntriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SnatEntriesResult, error) {
			args := v.(SnatEntriesArgs)
			r, err := SnatEntries(ctx, &args, opts...)
			var s SnatEntriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(SnatEntriesResultOutput)
}

// A collection of arguments for invoking SnatEntries.
type SnatEntriesOutputArgs struct {
	// An id of the public ip address used by the SNAT entry.
	EipId pulumi.StringPtrInput `pulumi:"eipId"`
	// A list of SNAT entry ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// An id of the nat gateway to which the entry belongs.
	NatGatewayId pulumi.StringPtrInput `pulumi:"natGatewayId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A name of SNAT entry.
	SnatEntryName pulumi.StringPtrInput `pulumi:"snatEntryName"`
	// The SourceCidr of SNAT entry.
	SourceCidr pulumi.StringPtrInput `pulumi:"sourceCidr"`
	// An id of the subnet that is required to access the Internet.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
}

func (SnatEntriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SnatEntriesArgs)(nil)).Elem()
}

// A collection of values returned by SnatEntries.
type SnatEntriesResultOutput struct{ *pulumi.OutputState }

func (SnatEntriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnatEntriesResult)(nil)).Elem()
}

func (o SnatEntriesResultOutput) ToSnatEntriesResultOutput() SnatEntriesResultOutput {
	return o
}

func (o SnatEntriesResultOutput) ToSnatEntriesResultOutputWithContext(ctx context.Context) SnatEntriesResultOutput {
	return o
}

// The id of the public ip address used by the SNAT entry.
func (o SnatEntriesResultOutput) EipId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnatEntriesResult) *string { return v.EipId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o SnatEntriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SnatEntriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o SnatEntriesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SnatEntriesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The id of the nat gateway to which the entry belongs.
func (o SnatEntriesResultOutput) NatGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnatEntriesResult) *string { return v.NatGatewayId }).(pulumi.StringPtrOutput)
}

func (o SnatEntriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnatEntriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of snat entries.
func (o SnatEntriesResultOutput) SnatEntries() SnatEntriesSnatEntryArrayOutput {
	return o.ApplyT(func(v SnatEntriesResult) []SnatEntriesSnatEntry { return v.SnatEntries }).(SnatEntriesSnatEntryArrayOutput)
}

// The name of the SNAT entry.
func (o SnatEntriesResultOutput) SnatEntryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnatEntriesResult) *string { return v.SnatEntryName }).(pulumi.StringPtrOutput)
}

// The SourceCidr of the SNAT entry.
func (o SnatEntriesResultOutput) SourceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnatEntriesResult) *string { return v.SourceCidr }).(pulumi.StringPtrOutput)
}

// The id of the subnet that is required to access the internet.
func (o SnatEntriesResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnatEntriesResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// The total count of snat entries query.
func (o SnatEntriesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v SnatEntriesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(SnatEntriesResultOutput{})
}