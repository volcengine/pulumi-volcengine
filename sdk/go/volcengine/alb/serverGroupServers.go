// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of alb server group servers
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/alb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := alb.GetServerGroupServers(ctx, &alb.GetServerGroupServersArgs{
//				ServerGroupId: "rsp-1g7317vrcx3pc2zbhq4c3i6a2",
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
// Deprecated: volcengine.alb.ServerGroupServers has been deprecated in favor of volcengine.alb.getServerGroupServers
func ServerGroupServers(ctx *pulumi.Context, args *ServerGroupServersArgs, opts ...pulumi.InvokeOption) (*ServerGroupServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ServerGroupServersResult
	err := ctx.Invoke("volcengine:alb/serverGroupServers:ServerGroupServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ServerGroupServers.
type ServerGroupServersArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the ServerGroup.
	ServerGroupId string `pulumi:"serverGroupId"`
}

// A collection of values returned by ServerGroupServers.
type ServerGroupServersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	OutputFile    *string `pulumi:"outputFile"`
	ServerGroupId string  `pulumi:"serverGroupId"`
	// The server list of ServerGroup.
	Servers []ServerGroupServersServer `pulumi:"servers"`
	// The total count of ServerGroupServer query.
	TotalCount int `pulumi:"totalCount"`
}

func ServerGroupServersOutput(ctx *pulumi.Context, args ServerGroupServersOutputArgs, opts ...pulumi.InvokeOption) ServerGroupServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ServerGroupServersResult, error) {
			args := v.(ServerGroupServersArgs)
			r, err := ServerGroupServers(ctx, &args, opts...)
			var s ServerGroupServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ServerGroupServersResultOutput)
}

// A collection of arguments for invoking ServerGroupServers.
type ServerGroupServersOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringInput `pulumi:"serverGroupId"`
}

func (ServerGroupServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupServersArgs)(nil)).Elem()
}

// A collection of values returned by ServerGroupServers.
type ServerGroupServersResultOutput struct{ *pulumi.OutputState }

func (ServerGroupServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupServersResult)(nil)).Elem()
}

func (o ServerGroupServersResultOutput) ToServerGroupServersResultOutput() ServerGroupServersResultOutput {
	return o
}

func (o ServerGroupServersResultOutput) ToServerGroupServersResultOutputWithContext(ctx context.Context) ServerGroupServersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o ServerGroupServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServerGroupServersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServerGroupServersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupServersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o ServerGroupServersResultOutput) ServerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ServerGroupServersResult) string { return v.ServerGroupId }).(pulumi.StringOutput)
}

// The server list of ServerGroup.
func (o ServerGroupServersResultOutput) Servers() ServerGroupServersServerArrayOutput {
	return o.ApplyT(func(v ServerGroupServersResult) []ServerGroupServersServer { return v.Servers }).(ServerGroupServersServerArrayOutput)
}

// The total count of ServerGroupServer query.
func (o ServerGroupServersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ServerGroupServersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerGroupServersResultOutput{})
}
