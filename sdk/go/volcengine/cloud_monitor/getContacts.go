// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud_monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of cloud monitor contacts
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cloud_monitor"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloud_monitor.GetContacts(ctx, &cloud_monitor.GetContactsArgs{
//				Ids: []string{
//					"17******516",
//					"1712**********0",
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
func GetContacts(ctx *pulumi.Context, args *GetContactsArgs, opts ...pulumi.InvokeOption) (*GetContactsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetContactsResult
	err := ctx.Invoke("volcengine:cloud_monitor/getContacts:getContacts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContacts.
type GetContactsArgs struct {
	// A list of Contact IDs.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by getContacts.
type GetContactsResult struct {
	// The collection of query.
	Contacts []GetContactsContact `pulumi:"contacts"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func GetContactsOutput(ctx *pulumi.Context, args GetContactsOutputArgs, opts ...pulumi.InvokeOption) GetContactsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetContactsResult, error) {
			args := v.(GetContactsArgs)
			r, err := GetContacts(ctx, &args, opts...)
			var s GetContactsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetContactsResultOutput)
}

// A collection of arguments for invoking getContacts.
type GetContactsOutputArgs struct {
	// A list of Contact IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (GetContactsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContactsArgs)(nil)).Elem()
}

// A collection of values returned by getContacts.
type GetContactsResultOutput struct{ *pulumi.OutputState }

func (GetContactsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContactsResult)(nil)).Elem()
}

func (o GetContactsResultOutput) ToGetContactsResultOutput() GetContactsResultOutput {
	return o
}

func (o GetContactsResultOutput) ToGetContactsResultOutputWithContext(ctx context.Context) GetContactsResultOutput {
	return o
}

// The collection of query.
func (o GetContactsResultOutput) Contacts() GetContactsContactArrayOutput {
	return o.ApplyT(func(v GetContactsResult) []GetContactsContact { return v.Contacts }).(GetContactsContactArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetContactsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContactsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetContactsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetContactsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetContactsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContactsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o GetContactsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetContactsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContactsResultOutput{})
}
