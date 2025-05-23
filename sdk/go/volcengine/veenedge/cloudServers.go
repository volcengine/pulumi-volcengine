// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package veenedge

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of veenedge cloud servers
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/veenedge"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := veenedge.GetCloudServers(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Deprecated: volcengine.veenedge.CloudServers has been deprecated in favor of volcengine.veenedge.getCloudServers
func CloudServers(ctx *pulumi.Context, args *CloudServersArgs, opts ...pulumi.InvokeOption) (*CloudServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv CloudServersResult
	err := ctx.Invoke("volcengine:veenedge/cloudServers:CloudServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking CloudServers.
type CloudServersArgs struct {
	// A list of cloud server IDs.
	Ids []string `pulumi:"ids"`
	// A Name Regex of Cloud Server.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by CloudServers.
type CloudServersResult struct {
	// The collection of cloud servers query.
	CloudServers []CloudServersCloudServer `pulumi:"cloudServers"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The total count of cloud servers query.
	TotalCount int `pulumi:"totalCount"`
}

func CloudServersOutput(ctx *pulumi.Context, args CloudServersOutputArgs, opts ...pulumi.InvokeOption) CloudServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CloudServersResult, error) {
			args := v.(CloudServersArgs)
			r, err := CloudServers(ctx, &args, opts...)
			var s CloudServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CloudServersResultOutput)
}

// A collection of arguments for invoking CloudServers.
type CloudServersOutputArgs struct {
	// A list of cloud server IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of Cloud Server.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (CloudServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServersArgs)(nil)).Elem()
}

// A collection of values returned by CloudServers.
type CloudServersResultOutput struct{ *pulumi.OutputState }

func (CloudServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudServersResult)(nil)).Elem()
}

func (o CloudServersResultOutput) ToCloudServersResultOutput() CloudServersResultOutput {
	return o
}

func (o CloudServersResultOutput) ToCloudServersResultOutputWithContext(ctx context.Context) CloudServersResultOutput {
	return o
}

// The collection of cloud servers query.
func (o CloudServersResultOutput) CloudServers() CloudServersCloudServerArrayOutput {
	return o.ApplyT(func(v CloudServersResult) []CloudServersCloudServer { return v.CloudServers }).(CloudServersCloudServerArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o CloudServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CloudServersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o CloudServersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CloudServersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o CloudServersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o CloudServersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudServersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of cloud servers query.
func (o CloudServersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v CloudServersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(CloudServersResultOutput{})
}
