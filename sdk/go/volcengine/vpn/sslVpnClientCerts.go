// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of ssl vpn client certs
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpn"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// fooZones, err := ecs.Zones(ctx, nil, nil);
// if err != nil {
// return err
// }
// fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
// VpcName: pulumi.String("acc-test-vpc"),
// CidrBlock: pulumi.String("172.16.0.0/16"),
// })
// if err != nil {
// return err
// }
// fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
// SubnetName: pulumi.String("acc-test-subnet"),
// CidrBlock: pulumi.String("172.16.0.0/24"),
// ZoneId: *pulumi.String(fooZones.Zones[0].Id),
// VpcId: fooVpc.ID(),
// })
// if err != nil {
// return err
// }
// fooGateway, err := vpn.NewGateway(ctx, "fooGateway", &vpn.GatewayArgs{
// VpcId: fooVpc.ID(),
// SubnetId: fooSubnet.ID(),
// Bandwidth: pulumi.Int(5),
// VpnGatewayName: pulumi.String("acc-test1"),
// Description: pulumi.String("acc-test1"),
// Period: pulumi.Int(7),
// ProjectName: pulumi.String("default"),
// SslEnabled: pulumi.Bool(true),
// SslMaxConnections: pulumi.Int(5),
// })
// if err != nil {
// return err
// }
// fooSslVpnServer, err := vpn.NewSslVpnServer(ctx, "fooSslVpnServer", &vpn.SslVpnServerArgs{
// VpnGatewayId: fooGateway.ID(),
// LocalSubnets: pulumi.StringArray{
// fooSubnet.CidrBlock,
// },
// ClientIpPool: pulumi.String("172.16.2.0/24"),
// SslVpnServerName: pulumi.String("acc-test-ssl"),
// Description: pulumi.String("acc-test"),
// Protocol: pulumi.String("UDP"),
// Cipher: pulumi.String("AES-128-CBC"),
// Auth: pulumi.String("SHA1"),
// Compress: pulumi.Bool(true),
// })
// if err != nil {
// return err
// }
// var fooSslVpnClientCert []*vpn.SslVpnClientCert
//
//	for index := 0; index < 5; index++ {
//	    key0 := index
//	    val0 := index
//
// __res, err := vpn.NewSslVpnClientCert(ctx, fmt.Sprintf("fooSslVpnClientCert-%v", key0), &vpn.SslVpnClientCertArgs{
// SslVpnServerId: fooSslVpnServer.ID(),
// SslVpnClientCertName: pulumi.String(fmt.Sprintf("acc-test-client-cert-%v", val0)),
// Description: pulumi.String("acc-test"),
// })
// if err != nil {
// return err
// }
// fooSslVpnClientCert = append(fooSslVpnClientCert, __res)
// }
// _ = vpn.SslVpnClientCertsOutput(ctx, vpn.SslVpnClientCertsOutputArgs{
// Ids: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-functions-volcengine:vpn-sslVpnClientCerts:SslVpnClientCerts.pp:44,9-34),
// }, nil);
// return nil
// })
// }
// ```
func SslVpnClientCerts(ctx *pulumi.Context, args *SslVpnClientCertsArgs, opts ...pulumi.InvokeOption) (*SslVpnClientCertsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv SslVpnClientCertsResult
	err := ctx.Invoke("volcengine:vpn/sslVpnClientCerts:SslVpnClientCerts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking SslVpnClientCerts.
type SslVpnClientCertsArgs struct {
	// The ids list of ssl vpn client cert.
	Ids []string `pulumi:"ids"`
	// A Name Regex of ssl vpn client cert.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of the ssl vpn client cert.
	SslVpnClientCertName *string `pulumi:"sslVpnClientCertName"`
	// The id of the ssl vpn server.
	SslVpnServerId *string `pulumi:"sslVpnServerId"`
}

// A collection of values returned by SslVpnClientCerts.
type SslVpnClientCertsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The name of the ssl vpn client cert.
	SslVpnClientCertName *string `pulumi:"sslVpnClientCertName"`
	// The collection of of ssl vpn client certs.
	SslVpnClientCerts []SslVpnClientCertsSslVpnClientCert `pulumi:"sslVpnClientCerts"`
	// The id of the ssl vpn server.
	SslVpnServerId *string `pulumi:"sslVpnServerId"`
	// The total count of ssl vpn client cert query.
	TotalCount int `pulumi:"totalCount"`
}

func SslVpnClientCertsOutput(ctx *pulumi.Context, args SslVpnClientCertsOutputArgs, opts ...pulumi.InvokeOption) SslVpnClientCertsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (SslVpnClientCertsResult, error) {
			args := v.(SslVpnClientCertsArgs)
			r, err := SslVpnClientCerts(ctx, &args, opts...)
			var s SslVpnClientCertsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(SslVpnClientCertsResultOutput)
}

// A collection of arguments for invoking SslVpnClientCerts.
type SslVpnClientCertsOutputArgs struct {
	// The ids list of ssl vpn client cert.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of ssl vpn client cert.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of the ssl vpn client cert.
	SslVpnClientCertName pulumi.StringPtrInput `pulumi:"sslVpnClientCertName"`
	// The id of the ssl vpn server.
	SslVpnServerId pulumi.StringPtrInput `pulumi:"sslVpnServerId"`
}

func (SslVpnClientCertsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SslVpnClientCertsArgs)(nil)).Elem()
}

// A collection of values returned by SslVpnClientCerts.
type SslVpnClientCertsResultOutput struct{ *pulumi.OutputState }

func (SslVpnClientCertsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslVpnClientCertsResult)(nil)).Elem()
}

func (o SslVpnClientCertsResultOutput) ToSslVpnClientCertsResultOutput() SslVpnClientCertsResultOutput {
	return o
}

func (o SslVpnClientCertsResultOutput) ToSslVpnClientCertsResultOutputWithContext(ctx context.Context) SslVpnClientCertsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o SslVpnClientCertsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SslVpnClientCertsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o SslVpnClientCertsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SslVpnClientCertsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o SslVpnClientCertsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslVpnClientCertsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o SslVpnClientCertsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslVpnClientCertsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The name of the ssl vpn client cert.
func (o SslVpnClientCertsResultOutput) SslVpnClientCertName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslVpnClientCertsResult) *string { return v.SslVpnClientCertName }).(pulumi.StringPtrOutput)
}

// The collection of of ssl vpn client certs.
func (o SslVpnClientCertsResultOutput) SslVpnClientCerts() SslVpnClientCertsSslVpnClientCertArrayOutput {
	return o.ApplyT(func(v SslVpnClientCertsResult) []SslVpnClientCertsSslVpnClientCert { return v.SslVpnClientCerts }).(SslVpnClientCertsSslVpnClientCertArrayOutput)
}

// The id of the ssl vpn server.
func (o SslVpnClientCertsResultOutput) SslVpnServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SslVpnClientCertsResult) *string { return v.SslVpnServerId }).(pulumi.StringPtrOutput)
}

// The total count of ssl vpn client cert query.
func (o SslVpnClientCertsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v SslVpnClientCertsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(SslVpnClientCertsResultOutput{})
}