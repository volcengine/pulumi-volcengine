// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bandwidth_package

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of bandwidth packages
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/bandwidth_package"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// var fooBandwidthPackage []*bandwidth_package.BandwidthPackage
//
//	for index := 0; index < 2; index++ {
//	    key0 := index
//	    _ := index
//
// __res, err := bandwidth_package.NewBandwidthPackage(ctx, fmt.Sprintf("fooBandwidthPackage-%v", key0), &bandwidth_package.BandwidthPackageArgs{
// BandwidthPackageName: pulumi.String("acc-test-bp"),
// BillingType: pulumi.String("PostPaidByBandwidth"),
// Isp: pulumi.String("BGP"),
// Description: pulumi.String("acc-test"),
// Bandwidth: pulumi.Int(2),
// Protocol: pulumi.String("IPv4"),
// SecurityProtectionTypes: pulumi.StringArray{
// pulumi.String("AntiDDoS_Enhanced"),
// },
// Tags: bandwidth_package.BandwidthPackageTagArray{
// &bandwidth_package.BandwidthPackageTagArgs{
// Key: pulumi.String("k1"),
// Value: pulumi.String("v1"),
// },
// },
// })
// if err != nil {
// return err
// }
// fooBandwidthPackage = append(fooBandwidthPackage, __res)
// }
// _ = bandwidth_package.BandwidthPackagesOutput(ctx, bandwidth_package.BandwidthPackagesOutputArgs{
// Ids: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-functions-volcengine:bandwidth_package-bandwidthPackages:BandwidthPackages.pp:18,9-34),
// }, nil);
// return nil
// })
// }
// ```
func BandwidthPackages(ctx *pulumi.Context, args *BandwidthPackagesArgs, opts ...pulumi.InvokeOption) (*BandwidthPackagesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv BandwidthPackagesResult
	err := ctx.Invoke("volcengine:bandwidth_package/bandwidthPackages:BandwidthPackages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking BandwidthPackages.
type BandwidthPackagesArgs struct {
	// Shared bandwidth package name to be queried.
	BandwidthPackageName *string `pulumi:"bandwidthPackageName"`
	// Shared bandwidth package instance ID to be queried.
	Ids []string `pulumi:"ids"`
	// Line types for shared bandwidth packages.
	Isp *string `pulumi:"isp"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The project name of the bandwidth package to be queried.
	ProjectName *string `pulumi:"projectName"`
	// The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
	Protocol *string `pulumi:"protocol"`
	// Security protection types for shared bandwidth packages.
	SecurityProtectionEnabled *bool `pulumi:"securityProtectionEnabled"`
	// A list of tags.
	TagFilters []BandwidthPackagesTagFilter `pulumi:"tagFilters"`
}

// A collection of values returned by BandwidthPackages.
type BandwidthPackagesResult struct {
	// The name of the bandwidth package.
	BandwidthPackageName *string `pulumi:"bandwidthPackageName"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The line type.
	Isp        *string `pulumi:"isp"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of query.
	Packages []BandwidthPackagesPackage `pulumi:"packages"`
	// The project name of the bandwidth package.
	ProjectName *string `pulumi:"projectName"`
	// The protocol of the bandwidth package.
	Protocol                  *string                      `pulumi:"protocol"`
	SecurityProtectionEnabled *bool                        `pulumi:"securityProtectionEnabled"`
	TagFilters                []BandwidthPackagesTagFilter `pulumi:"tagFilters"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func BandwidthPackagesOutput(ctx *pulumi.Context, args BandwidthPackagesOutputArgs, opts ...pulumi.InvokeOption) BandwidthPackagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (BandwidthPackagesResult, error) {
			args := v.(BandwidthPackagesArgs)
			r, err := BandwidthPackages(ctx, &args, opts...)
			var s BandwidthPackagesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(BandwidthPackagesResultOutput)
}

// A collection of arguments for invoking BandwidthPackages.
type BandwidthPackagesOutputArgs struct {
	// Shared bandwidth package name to be queried.
	BandwidthPackageName pulumi.StringPtrInput `pulumi:"bandwidthPackageName"`
	// Shared bandwidth package instance ID to be queried.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Line types for shared bandwidth packages.
	Isp pulumi.StringPtrInput `pulumi:"isp"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The project name of the bandwidth package to be queried.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// Security protection types for shared bandwidth packages.
	SecurityProtectionEnabled pulumi.BoolPtrInput `pulumi:"securityProtectionEnabled"`
	// A list of tags.
	TagFilters BandwidthPackagesTagFilterArrayInput `pulumi:"tagFilters"`
}

func (BandwidthPackagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthPackagesArgs)(nil)).Elem()
}

// A collection of values returned by BandwidthPackages.
type BandwidthPackagesResultOutput struct{ *pulumi.OutputState }

func (BandwidthPackagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthPackagesResult)(nil)).Elem()
}

func (o BandwidthPackagesResultOutput) ToBandwidthPackagesResultOutput() BandwidthPackagesResultOutput {
	return o
}

func (o BandwidthPackagesResultOutput) ToBandwidthPackagesResultOutputWithContext(ctx context.Context) BandwidthPackagesResultOutput {
	return o
}

// The name of the bandwidth package.
func (o BandwidthPackagesResultOutput) BandwidthPackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.BandwidthPackageName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o BandwidthPackagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o BandwidthPackagesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The line type.
func (o BandwidthPackagesResultOutput) Isp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.Isp }).(pulumi.StringPtrOutput)
}

func (o BandwidthPackagesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o BandwidthPackagesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o BandwidthPackagesResultOutput) Packages() BandwidthPackagesPackageArrayOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) []BandwidthPackagesPackage { return v.Packages }).(BandwidthPackagesPackageArrayOutput)
}

// The project name of the bandwidth package.
func (o BandwidthPackagesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The protocol of the bandwidth package.
func (o BandwidthPackagesResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o BandwidthPackagesResultOutput) SecurityProtectionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *bool { return v.SecurityProtectionEnabled }).(pulumi.BoolPtrOutput)
}

func (o BandwidthPackagesResultOutput) TagFilters() BandwidthPackagesTagFilterArrayOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) []BandwidthPackagesTagFilter { return v.TagFilters }).(BandwidthPackagesTagFilterArrayOutput)
}

// The total count of query.
func (o BandwidthPackagesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(BandwidthPackagesResultOutput{})
}