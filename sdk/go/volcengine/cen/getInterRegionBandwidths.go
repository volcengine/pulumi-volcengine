// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of cen inter region bandwidths
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cen"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooCen, err := cen.NewCen(ctx, "fooCen", &cen.CenArgs{
//				CenName:     pulumi.String("acc-test-cen"),
//				Description: pulumi.String("acc-test"),
//				ProjectName: pulumi.String("default"),
//				Tags: cen.CenTagArray{
//					&cen.CenTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooBandwidthPackage, err := cen.NewBandwidthPackage(ctx, "fooBandwidthPackage", &cen.BandwidthPackageArgs{
//				LocalGeographicRegionSetId: pulumi.String("China"),
//				PeerGeographicRegionSetId:  pulumi.String("China"),
//				Bandwidth:                  pulumi.Int(5),
//				CenBandwidthPackageName:    pulumi.String("acc-test-cen-bp"),
//				Description:                pulumi.String("acc-test"),
//				BillingType:                pulumi.String("PrePaid"),
//				PeriodUnit:                 pulumi.String("Month"),
//				Period:                     pulumi.Int(1),
//				ProjectName:                pulumi.String("default"),
//				Tags: cen.BandwidthPackageTagArray{
//					&cen.BandwidthPackageTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooBandwidthPackageAssociate, err := cen.NewBandwidthPackageAssociate(ctx, "fooBandwidthPackageAssociate", &cen.BandwidthPackageAssociateArgs{
//				CenBandwidthPackageId: fooBandwidthPackage.ID(),
//				CenId:                 fooCen.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooInterRegionBandwidth, err := cen.NewInterRegionBandwidth(ctx, "fooInterRegionBandwidth", &cen.InterRegionBandwidthArgs{
//				CenId:         fooCen.ID(),
//				LocalRegionId: pulumi.String("cn-beijing"),
//				PeerRegionId:  pulumi.String("cn-shanghai"),
//				Bandwidth:     pulumi.Int(2),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				fooBandwidthPackageAssociate,
//			}))
//			if err != nil {
//				return err
//			}
//			_ = cen.GetInterRegionBandwidthsOutput(ctx, cen.GetInterRegionBandwidthsOutputArgs{
//				Ids: pulumi.StringArray{
//					fooInterRegionBandwidth.ID(),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetInterRegionBandwidths(ctx *pulumi.Context, args *GetInterRegionBandwidthsArgs, opts ...pulumi.InvokeOption) (*GetInterRegionBandwidthsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInterRegionBandwidthsResult
	err := ctx.Invoke("volcengine:cen/getInterRegionBandwidths:getInterRegionBandwidths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInterRegionBandwidths.
type GetInterRegionBandwidthsArgs struct {
	// The ID of the cen.
	CenId *string `pulumi:"cenId"`
	// A list of cen inter region bandwidth IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getInterRegionBandwidths.
type GetInterRegionBandwidthsResult struct {
	// The cen ID of the cen inter region bandwidth.
	CenId *string `pulumi:"cenId"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The collection of cen inter region bandwidth query.
	InterRegionBandwidths []GetInterRegionBandwidthsInterRegionBandwidth `pulumi:"interRegionBandwidths"`
	OutputFile            *string                                        `pulumi:"outputFile"`
	// The total count of cen inter region bandwidth query.
	TotalCount int `pulumi:"totalCount"`
}

func GetInterRegionBandwidthsOutput(ctx *pulumi.Context, args GetInterRegionBandwidthsOutputArgs, opts ...pulumi.InvokeOption) GetInterRegionBandwidthsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInterRegionBandwidthsResult, error) {
			args := v.(GetInterRegionBandwidthsArgs)
			r, err := GetInterRegionBandwidths(ctx, &args, opts...)
			var s GetInterRegionBandwidthsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInterRegionBandwidthsResultOutput)
}

// A collection of arguments for invoking getInterRegionBandwidths.
type GetInterRegionBandwidthsOutputArgs struct {
	// The ID of the cen.
	CenId pulumi.StringPtrInput `pulumi:"cenId"`
	// A list of cen inter region bandwidth IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetInterRegionBandwidthsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInterRegionBandwidthsArgs)(nil)).Elem()
}

// A collection of values returned by getInterRegionBandwidths.
type GetInterRegionBandwidthsResultOutput struct{ *pulumi.OutputState }

func (GetInterRegionBandwidthsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInterRegionBandwidthsResult)(nil)).Elem()
}

func (o GetInterRegionBandwidthsResultOutput) ToGetInterRegionBandwidthsResultOutput() GetInterRegionBandwidthsResultOutput {
	return o
}

func (o GetInterRegionBandwidthsResultOutput) ToGetInterRegionBandwidthsResultOutputWithContext(ctx context.Context) GetInterRegionBandwidthsResultOutput {
	return o
}

// The cen ID of the cen inter region bandwidth.
func (o GetInterRegionBandwidthsResultOutput) CenId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInterRegionBandwidthsResult) *string { return v.CenId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInterRegionBandwidthsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInterRegionBandwidthsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInterRegionBandwidthsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInterRegionBandwidthsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The collection of cen inter region bandwidth query.
func (o GetInterRegionBandwidthsResultOutput) InterRegionBandwidths() GetInterRegionBandwidthsInterRegionBandwidthArrayOutput {
	return o.ApplyT(func(v GetInterRegionBandwidthsResult) []GetInterRegionBandwidthsInterRegionBandwidth {
		return v.InterRegionBandwidths
	}).(GetInterRegionBandwidthsInterRegionBandwidthArrayOutput)
}

func (o GetInterRegionBandwidthsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInterRegionBandwidthsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of cen inter region bandwidth query.
func (o GetInterRegionBandwidthsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetInterRegionBandwidthsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInterRegionBandwidthsResultOutput{})
}
