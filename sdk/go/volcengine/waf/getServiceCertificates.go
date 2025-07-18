// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of waf service certificates
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/waf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := waf.GetServiceCertificates(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetServiceCertificates(ctx *pulumi.Context, args *GetServiceCertificatesArgs, opts ...pulumi.InvokeOption) (*GetServiceCertificatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceCertificatesResult
	err := ctx.Invoke("volcengine:waf/getServiceCertificates:getServiceCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceCertificates.
type GetServiceCertificatesArgs struct {
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getServiceCertificates.
type GetServiceCertificatesResult struct {
	// The Information of the certificate.
	Datas []GetServiceCertificatesData `pulumi:"datas"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetServiceCertificatesOutput(ctx *pulumi.Context, args GetServiceCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetServiceCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServiceCertificatesResult, error) {
			args := v.(GetServiceCertificatesArgs)
			r, err := GetServiceCertificates(ctx, &args, opts...)
			var s GetServiceCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServiceCertificatesResultOutput)
}

// A collection of arguments for invoking getServiceCertificates.
type GetServiceCertificatesOutputArgs struct {
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetServiceCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getServiceCertificates.
type GetServiceCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetServiceCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceCertificatesResult)(nil)).Elem()
}

func (o GetServiceCertificatesResultOutput) ToGetServiceCertificatesResultOutput() GetServiceCertificatesResultOutput {
	return o
}

func (o GetServiceCertificatesResultOutput) ToGetServiceCertificatesResultOutputWithContext(ctx context.Context) GetServiceCertificatesResultOutput {
	return o
}

// The Information of the certificate.
func (o GetServiceCertificatesResultOutput) Datas() GetServiceCertificatesDataArrayOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) []GetServiceCertificatesData { return v.Datas }).(GetServiceCertificatesDataArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServiceCertificatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetServiceCertificatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetServiceCertificatesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetServiceCertificatesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceCertificatesResultOutput{})
}
