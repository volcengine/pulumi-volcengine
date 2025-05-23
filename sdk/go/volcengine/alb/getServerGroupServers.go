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
func GetServerGroupServers(ctx *pulumi.Context, args *GetServerGroupServersArgs, opts ...pulumi.InvokeOption) (*GetServerGroupServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServerGroupServersResult
	err := ctx.Invoke("volcengine:alb/getServerGroupServers:getServerGroupServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerGroupServers.
type GetServerGroupServersArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the ServerGroup.
	ServerGroupId string `pulumi:"serverGroupId"`
}

// A collection of values returned by getServerGroupServers.
type GetServerGroupServersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	OutputFile    *string `pulumi:"outputFile"`
	ServerGroupId string  `pulumi:"serverGroupId"`
	// The server list of ServerGroup.
	Servers []GetServerGroupServersServer `pulumi:"servers"`
	// The total count of ServerGroupServer query.
	TotalCount int `pulumi:"totalCount"`
}

func GetServerGroupServersOutput(ctx *pulumi.Context, args GetServerGroupServersOutputArgs, opts ...pulumi.InvokeOption) GetServerGroupServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerGroupServersResult, error) {
			args := v.(GetServerGroupServersArgs)
			r, err := GetServerGroupServers(ctx, &args, opts...)
			var s GetServerGroupServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerGroupServersResultOutput)
}

// A collection of arguments for invoking getServerGroupServers.
type GetServerGroupServersOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringInput `pulumi:"serverGroupId"`
}

func (GetServerGroupServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerGroupServersArgs)(nil)).Elem()
}

// A collection of values returned by getServerGroupServers.
type GetServerGroupServersResultOutput struct{ *pulumi.OutputState }

func (GetServerGroupServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerGroupServersResult)(nil)).Elem()
}

func (o GetServerGroupServersResultOutput) ToGetServerGroupServersResultOutput() GetServerGroupServersResultOutput {
	return o
}

func (o GetServerGroupServersResultOutput) ToGetServerGroupServersResultOutputWithContext(ctx context.Context) GetServerGroupServersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerGroupServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerGroupServersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServerGroupServersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerGroupServersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o GetServerGroupServersResultOutput) ServerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerGroupServersResult) string { return v.ServerGroupId }).(pulumi.StringOutput)
}

// The server list of ServerGroup.
func (o GetServerGroupServersResultOutput) Servers() GetServerGroupServersServerArrayOutput {
	return o.ApplyT(func(v GetServerGroupServersResult) []GetServerGroupServersServer { return v.Servers }).(GetServerGroupServersServerArrayOutput)
}

// The total count of ServerGroupServer query.
func (o GetServerGroupServersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetServerGroupServersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerGroupServersResultOutput{})
}
