// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package direct_connect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of direct connect connections
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/direct_connect"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := direct_connect.Connections(ctx, &direct_connect.ConnectionsArgs{
//				DirectConnectConnectionName: pulumi.StringRef("tf_test"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Connections(ctx *pulumi.Context, args *ConnectionsArgs, opts ...pulumi.InvokeOption) (*ConnectionsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv ConnectionsResult
	err := ctx.Invoke("volcengine:direct_connect/connections:Connections", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Connections.
type ConnectionsArgs struct {
	// The connection type of physical leased line,valid value contains `SharedConnection`,`DedicatedConnection`.
	ConnectionType *string `pulumi:"connectionType"`
	// The ID of the physical leased line access point.
	DirectConnectAccessPointId *string `pulumi:"directConnectAccessPointId"`
	// The name of directi connect connection.
	DirectConnectConnectionName *string `pulumi:"directConnectConnectionName"`
	// A list of IDs.
	Ids []string `pulumi:"ids"`
	// The operator of the physical leased line,valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.
	LineOperator *string `pulumi:"lineOperator"`
	// A Name Regex of Resource.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The peer access point of the physical leased line.
	PeerLocation *string `pulumi:"peerLocation"`
	// The filter tag of direct connect.
	TagFilters []ConnectionsTagFilter `pulumi:"tagFilters"`
}

// A collection of values returned by Connections.
type ConnectionsResult struct {
	// The connection type of direct connect.
	ConnectionType *string `pulumi:"connectionType"`
	// The access point id of direct connect.
	DirectConnectAccessPointId *string `pulumi:"directConnectAccessPointId"`
	// The name of direct connect connection.
	DirectConnectConnectionName *string `pulumi:"directConnectConnectionName"`
	// The collection of query.
	DirectConnectConnections []ConnectionsDirectConnectConnection `pulumi:"directConnectConnections"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The operator of physical leased line.
	LineOperator *string `pulumi:"lineOperator"`
	NameRegex    *string `pulumi:"nameRegex"`
	OutputFile   *string `pulumi:"outputFile"`
	// The peer access point of the physical leased line.
	PeerLocation *string                `pulumi:"peerLocation"`
	TagFilters   []ConnectionsTagFilter `pulumi:"tagFilters"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func ConnectionsOutput(ctx *pulumi.Context, args ConnectionsOutputArgs, opts ...pulumi.InvokeOption) ConnectionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ConnectionsResult, error) {
			args := v.(ConnectionsArgs)
			r, err := Connections(ctx, &args, opts...)
			var s ConnectionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ConnectionsResultOutput)
}

// A collection of arguments for invoking Connections.
type ConnectionsOutputArgs struct {
	// The connection type of physical leased line,valid value contains `SharedConnection`,`DedicatedConnection`.
	ConnectionType pulumi.StringPtrInput `pulumi:"connectionType"`
	// The ID of the physical leased line access point.
	DirectConnectAccessPointId pulumi.StringPtrInput `pulumi:"directConnectAccessPointId"`
	// The name of directi connect connection.
	DirectConnectConnectionName pulumi.StringPtrInput `pulumi:"directConnectConnectionName"`
	// A list of IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The operator of the physical leased line,valid value contains `ChinaTelecom`,`ChinaMobile`,`ChinaUnicom`,`ChinaOther`.
	LineOperator pulumi.StringPtrInput `pulumi:"lineOperator"`
	// A Name Regex of Resource.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The peer access point of the physical leased line.
	PeerLocation pulumi.StringPtrInput `pulumi:"peerLocation"`
	// The filter tag of direct connect.
	TagFilters ConnectionsTagFilterArrayInput `pulumi:"tagFilters"`
}

func (ConnectionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionsArgs)(nil)).Elem()
}

// A collection of values returned by Connections.
type ConnectionsResultOutput struct{ *pulumi.OutputState }

func (ConnectionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionsResult)(nil)).Elem()
}

func (o ConnectionsResultOutput) ToConnectionsResultOutput() ConnectionsResultOutput {
	return o
}

func (o ConnectionsResultOutput) ToConnectionsResultOutputWithContext(ctx context.Context) ConnectionsResultOutput {
	return o
}

// The connection type of direct connect.
func (o ConnectionsResultOutput) ConnectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.ConnectionType }).(pulumi.StringPtrOutput)
}

// The access point id of direct connect.
func (o ConnectionsResultOutput) DirectConnectAccessPointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.DirectConnectAccessPointId }).(pulumi.StringPtrOutput)
}

// The name of direct connect connection.
func (o ConnectionsResultOutput) DirectConnectConnectionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.DirectConnectConnectionName }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o ConnectionsResultOutput) DirectConnectConnections() ConnectionsDirectConnectConnectionArrayOutput {
	return o.ApplyT(func(v ConnectionsResult) []ConnectionsDirectConnectConnection { return v.DirectConnectConnections }).(ConnectionsDirectConnectConnectionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o ConnectionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ConnectionsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConnectionsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The operator of physical leased line.
func (o ConnectionsResultOutput) LineOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.LineOperator }).(pulumi.StringPtrOutput)
}

func (o ConnectionsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o ConnectionsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The peer access point of the physical leased line.
func (o ConnectionsResultOutput) PeerLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionsResult) *string { return v.PeerLocation }).(pulumi.StringPtrOutput)
}

func (o ConnectionsResultOutput) TagFilters() ConnectionsTagFilterArrayOutput {
	return o.ApplyT(func(v ConnectionsResult) []ConnectionsTagFilter { return v.TagFilters }).(ConnectionsTagFilterArrayOutput)
}

// The total count of query.
func (o ConnectionsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ConnectionsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionsResultOutput{})
}