// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of waf domains
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
//			_, err := waf.GetDomains(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.waf.Domains has been deprecated in favor of volcengine.waf.getDomains
func Domains(ctx *pulumi.Context, args *DomainsArgs, opts ...pulumi.InvokeOption) (*DomainsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv DomainsResult
	err := ctx.Invoke("volcengine:waf/domains:Domains", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Domains.
type DomainsArgs struct {
	// Matching mode.
	AccurateQuery *int `pulumi:"accurateQuery"`
	// The domain name of the protected website that needs to be queried.
	Domain *string `pulumi:"domain"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Domains.
type DomainsResult struct {
	AccurateQuery *int `pulumi:"accurateQuery"`
	// The collection of query.
	Datas []DomainsData `pulumi:"datas"`
	// domain names that need to be protected by WAF.
	Domain *string `pulumi:"domain"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func DomainsOutput(ctx *pulumi.Context, args DomainsOutputArgs, opts ...pulumi.InvokeOption) DomainsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (DomainsResult, error) {
			args := v.(DomainsArgs)
			r, err := Domains(ctx, &args, opts...)
			var s DomainsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(DomainsResultOutput)
}

// A collection of arguments for invoking Domains.
type DomainsOutputArgs struct {
	// Matching mode.
	AccurateQuery pulumi.IntPtrInput `pulumi:"accurateQuery"`
	// The domain name of the protected website that needs to be queried.
	Domain pulumi.StringPtrInput `pulumi:"domain"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (DomainsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainsArgs)(nil)).Elem()
}

// A collection of values returned by Domains.
type DomainsResultOutput struct{ *pulumi.OutputState }

func (DomainsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainsResult)(nil)).Elem()
}

func (o DomainsResultOutput) ToDomainsResultOutput() DomainsResultOutput {
	return o
}

func (o DomainsResultOutput) ToDomainsResultOutputWithContext(ctx context.Context) DomainsResultOutput {
	return o
}

func (o DomainsResultOutput) AccurateQuery() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DomainsResult) *int { return v.AccurateQuery }).(pulumi.IntPtrOutput)
}

// The collection of query.
func (o DomainsResultOutput) Datas() DomainsDataArrayOutput {
	return o.ApplyT(func(v DomainsResult) []DomainsData { return v.Datas }).(DomainsDataArrayOutput)
}

// domain names that need to be protected by WAF.
func (o DomainsResultOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainsResult) *string { return v.Domain }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o DomainsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DomainsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o DomainsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o DomainsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DomainsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o DomainsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v DomainsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainsResultOutput{})
}
