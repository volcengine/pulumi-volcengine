// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package private_zone

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of private zones
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/private_zone"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := private_zone.GetPrivateZones(ctx, &private_zone.GetPrivateZonesArgs{
//				LineMode:      pulumi.IntRef(3),
//				RecursionMode: pulumi.BoolRef(true),
//				SearchMode:    pulumi.StringRef("EXACT"),
//				Zid:           pulumi.IntRef(770000),
//				ZoneName:      pulumi.StringRef("volces.com"),
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
// Deprecated: volcengine.private_zone.PrivateZones has been deprecated in favor of volcengine.private_zone.getPrivateZones
func PrivateZones(ctx *pulumi.Context, args *PrivateZonesArgs, opts ...pulumi.InvokeOption) (*PrivateZonesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv PrivateZonesResult
	err := ctx.Invoke("volcengine:private_zone/privateZones:PrivateZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking PrivateZones.
type PrivateZonesArgs struct {
	// The keyword of zone name.
	KeyWord *string `pulumi:"keyWord"`
	// The line mode of Private Zone, specified whether the intelligent mode and the load balance function is enabled.
	LineMode *int `pulumi:"lineMode"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The project name of the private zone.
	ProjectName *string `pulumi:"projectName"`
	// Whether the recursion mode of Private Zone is enabled.
	RecursionMode *bool `pulumi:"recursionMode"`
	// The region of Private Zone.
	Region *string `pulumi:"region"`
	// The search mode of query. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
	SearchMode *string `pulumi:"searchMode"`
	// List of tag filters.
	TagFilters []PrivateZonesTagFilter `pulumi:"tagFilters"`
	// The vpc id associated to Private Zone.
	VpcId *string `pulumi:"vpcId"`
	// The zid of Private Zone.
	Zid *int `pulumi:"zid"`
	// The name of Private Zone.
	ZoneName *string `pulumi:"zoneName"`
}

// A collection of values returned by PrivateZones.
type PrivateZonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	KeyWord *string `pulumi:"keyWord"`
	// The line mode of the private zone, specified whether the intelligent mode and the load balance function is enabled.
	LineMode   *int    `pulumi:"lineMode"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of query.
	PrivateZones []PrivateZonesPrivateZone `pulumi:"privateZones"`
	// The project name of the private zone.
	ProjectName *string `pulumi:"projectName"`
	// Whether the recursion mode of the private zone is enabled.
	RecursionMode *bool `pulumi:"recursionMode"`
	// The region of the private zone.
	Region     *string                 `pulumi:"region"`
	SearchMode *string                 `pulumi:"searchMode"`
	TagFilters []PrivateZonesTagFilter `pulumi:"tagFilters"`
	// The total count of query.
	TotalCount int     `pulumi:"totalCount"`
	VpcId      *string `pulumi:"vpcId"`
	// The id of the private zone.
	Zid *int `pulumi:"zid"`
	// The id of the private zone.
	ZoneName *string `pulumi:"zoneName"`
}

func PrivateZonesOutput(ctx *pulumi.Context, args PrivateZonesOutputArgs, opts ...pulumi.InvokeOption) PrivateZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (PrivateZonesResult, error) {
			args := v.(PrivateZonesArgs)
			r, err := PrivateZones(ctx, &args, opts...)
			var s PrivateZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(PrivateZonesResultOutput)
}

// A collection of arguments for invoking PrivateZones.
type PrivateZonesOutputArgs struct {
	// The keyword of zone name.
	KeyWord pulumi.StringPtrInput `pulumi:"keyWord"`
	// The line mode of Private Zone, specified whether the intelligent mode and the load balance function is enabled.
	LineMode pulumi.IntPtrInput `pulumi:"lineMode"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The project name of the private zone.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Whether the recursion mode of Private Zone is enabled.
	RecursionMode pulumi.BoolPtrInput `pulumi:"recursionMode"`
	// The region of Private Zone.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The search mode of query. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
	SearchMode pulumi.StringPtrInput `pulumi:"searchMode"`
	// List of tag filters.
	TagFilters PrivateZonesTagFilterArrayInput `pulumi:"tagFilters"`
	// The vpc id associated to Private Zone.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// The zid of Private Zone.
	Zid pulumi.IntPtrInput `pulumi:"zid"`
	// The name of Private Zone.
	ZoneName pulumi.StringPtrInput `pulumi:"zoneName"`
}

func (PrivateZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateZonesArgs)(nil)).Elem()
}

// A collection of values returned by PrivateZones.
type PrivateZonesResultOutput struct{ *pulumi.OutputState }

func (PrivateZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateZonesResult)(nil)).Elem()
}

func (o PrivateZonesResultOutput) ToPrivateZonesResultOutput() PrivateZonesResultOutput {
	return o
}

func (o PrivateZonesResultOutput) ToPrivateZonesResultOutputWithContext(ctx context.Context) PrivateZonesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o PrivateZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateZonesResultOutput) KeyWord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *string { return v.KeyWord }).(pulumi.StringPtrOutput)
}

// The line mode of the private zone, specified whether the intelligent mode and the load balance function is enabled.
func (o PrivateZonesResultOutput) LineMode() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *int { return v.LineMode }).(pulumi.IntPtrOutput)
}

func (o PrivateZonesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o PrivateZonesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o PrivateZonesResultOutput) PrivateZones() PrivateZonesPrivateZoneArrayOutput {
	return o.ApplyT(func(v PrivateZonesResult) []PrivateZonesPrivateZone { return v.PrivateZones }).(PrivateZonesPrivateZoneArrayOutput)
}

// The project name of the private zone.
func (o PrivateZonesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Whether the recursion mode of the private zone is enabled.
func (o PrivateZonesResultOutput) RecursionMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *bool { return v.RecursionMode }).(pulumi.BoolPtrOutput)
}

// The region of the private zone.
func (o PrivateZonesResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o PrivateZonesResultOutput) SearchMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *string { return v.SearchMode }).(pulumi.StringPtrOutput)
}

func (o PrivateZonesResultOutput) TagFilters() PrivateZonesTagFilterArrayOutput {
	return o.ApplyT(func(v PrivateZonesResult) []PrivateZonesTagFilter { return v.TagFilters }).(PrivateZonesTagFilterArrayOutput)
}

// The total count of query.
func (o PrivateZonesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v PrivateZonesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o PrivateZonesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// The id of the private zone.
func (o PrivateZonesResultOutput) Zid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *int { return v.Zid }).(pulumi.IntPtrOutput)
}

// The id of the private zone.
func (o PrivateZonesResultOutput) ZoneName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateZonesResult) *string { return v.ZoneName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateZonesResultOutput{})
}
