// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage alb customized cfg
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
//			_, err := alb.NewCustomizedCfg(ctx, "foo", &alb.CustomizedCfgArgs{
//				CustomizedCfgContent: pulumi.String("proxy_connect_timeout 4s;proxy_request_buffering on;"),
//				CustomizedCfgName:    pulumi.String("acc-test-cfg1"),
//				Description:          pulumi.String("This is a test modify"),
//				ProjectName:          pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// AlbCustomizedCfg can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:alb/customizedCfg:CustomizedCfg default ccfg-3cj44nv0jhhxc6c6rrtet****
// ```
type CustomizedCfg struct {
	pulumi.CustomResourceState

	// The content of CustomizedCfg. The length cannot exceed 4096 characters. Spaces and semicolons need to be escaped. Currently supported configuration items are `sslProtocols`, `sslCiphers`, `clientMaxBodySize`, `keepaliveTimeout`, `proxyRequestBuffering` and `proxyConnectTimeout`.
	CustomizedCfgContent pulumi.StringOutput `pulumi:"customizedCfgContent"`
	// The name of CustomizedCfg.
	CustomizedCfgName pulumi.StringOutput `pulumi:"customizedCfgName"`
	// The description of CustomizedCfg.
	Description pulumi.StringOutput `pulumi:"description"`
	// The project name of the CustomizedCfg.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
}

// NewCustomizedCfg registers a new resource with the given unique name, arguments, and options.
func NewCustomizedCfg(ctx *pulumi.Context,
	name string, args *CustomizedCfgArgs, opts ...pulumi.ResourceOption) (*CustomizedCfg, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomizedCfgContent == nil {
		return nil, errors.New("invalid value for required argument 'CustomizedCfgContent'")
	}
	if args.CustomizedCfgName == nil {
		return nil, errors.New("invalid value for required argument 'CustomizedCfgName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomizedCfg
	err := ctx.RegisterResource("volcengine:alb/customizedCfg:CustomizedCfg", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomizedCfg gets an existing CustomizedCfg resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomizedCfg(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomizedCfgState, opts ...pulumi.ResourceOption) (*CustomizedCfg, error) {
	var resource CustomizedCfg
	err := ctx.ReadResource("volcengine:alb/customizedCfg:CustomizedCfg", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomizedCfg resources.
type customizedCfgState struct {
	// The content of CustomizedCfg. The length cannot exceed 4096 characters. Spaces and semicolons need to be escaped. Currently supported configuration items are `sslProtocols`, `sslCiphers`, `clientMaxBodySize`, `keepaliveTimeout`, `proxyRequestBuffering` and `proxyConnectTimeout`.
	CustomizedCfgContent *string `pulumi:"customizedCfgContent"`
	// The name of CustomizedCfg.
	CustomizedCfgName *string `pulumi:"customizedCfgName"`
	// The description of CustomizedCfg.
	Description *string `pulumi:"description"`
	// The project name of the CustomizedCfg.
	ProjectName *string `pulumi:"projectName"`
}

type CustomizedCfgState struct {
	// The content of CustomizedCfg. The length cannot exceed 4096 characters. Spaces and semicolons need to be escaped. Currently supported configuration items are `sslProtocols`, `sslCiphers`, `clientMaxBodySize`, `keepaliveTimeout`, `proxyRequestBuffering` and `proxyConnectTimeout`.
	CustomizedCfgContent pulumi.StringPtrInput
	// The name of CustomizedCfg.
	CustomizedCfgName pulumi.StringPtrInput
	// The description of CustomizedCfg.
	Description pulumi.StringPtrInput
	// The project name of the CustomizedCfg.
	ProjectName pulumi.StringPtrInput
}

func (CustomizedCfgState) ElementType() reflect.Type {
	return reflect.TypeOf((*customizedCfgState)(nil)).Elem()
}

type customizedCfgArgs struct {
	// The content of CustomizedCfg. The length cannot exceed 4096 characters. Spaces and semicolons need to be escaped. Currently supported configuration items are `sslProtocols`, `sslCiphers`, `clientMaxBodySize`, `keepaliveTimeout`, `proxyRequestBuffering` and `proxyConnectTimeout`.
	CustomizedCfgContent string `pulumi:"customizedCfgContent"`
	// The name of CustomizedCfg.
	CustomizedCfgName string `pulumi:"customizedCfgName"`
	// The description of CustomizedCfg.
	Description *string `pulumi:"description"`
	// The project name of the CustomizedCfg.
	ProjectName *string `pulumi:"projectName"`
}

// The set of arguments for constructing a CustomizedCfg resource.
type CustomizedCfgArgs struct {
	// The content of CustomizedCfg. The length cannot exceed 4096 characters. Spaces and semicolons need to be escaped. Currently supported configuration items are `sslProtocols`, `sslCiphers`, `clientMaxBodySize`, `keepaliveTimeout`, `proxyRequestBuffering` and `proxyConnectTimeout`.
	CustomizedCfgContent pulumi.StringInput
	// The name of CustomizedCfg.
	CustomizedCfgName pulumi.StringInput
	// The description of CustomizedCfg.
	Description pulumi.StringPtrInput
	// The project name of the CustomizedCfg.
	ProjectName pulumi.StringPtrInput
}

func (CustomizedCfgArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customizedCfgArgs)(nil)).Elem()
}

type CustomizedCfgInput interface {
	pulumi.Input

	ToCustomizedCfgOutput() CustomizedCfgOutput
	ToCustomizedCfgOutputWithContext(ctx context.Context) CustomizedCfgOutput
}

func (*CustomizedCfg) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizedCfg)(nil)).Elem()
}

func (i *CustomizedCfg) ToCustomizedCfgOutput() CustomizedCfgOutput {
	return i.ToCustomizedCfgOutputWithContext(context.Background())
}

func (i *CustomizedCfg) ToCustomizedCfgOutputWithContext(ctx context.Context) CustomizedCfgOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizedCfgOutput)
}

// CustomizedCfgArrayInput is an input type that accepts CustomizedCfgArray and CustomizedCfgArrayOutput values.
// You can construct a concrete instance of `CustomizedCfgArrayInput` via:
//
//	CustomizedCfgArray{ CustomizedCfgArgs{...} }
type CustomizedCfgArrayInput interface {
	pulumi.Input

	ToCustomizedCfgArrayOutput() CustomizedCfgArrayOutput
	ToCustomizedCfgArrayOutputWithContext(context.Context) CustomizedCfgArrayOutput
}

type CustomizedCfgArray []CustomizedCfgInput

func (CustomizedCfgArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomizedCfg)(nil)).Elem()
}

func (i CustomizedCfgArray) ToCustomizedCfgArrayOutput() CustomizedCfgArrayOutput {
	return i.ToCustomizedCfgArrayOutputWithContext(context.Background())
}

func (i CustomizedCfgArray) ToCustomizedCfgArrayOutputWithContext(ctx context.Context) CustomizedCfgArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizedCfgArrayOutput)
}

// CustomizedCfgMapInput is an input type that accepts CustomizedCfgMap and CustomizedCfgMapOutput values.
// You can construct a concrete instance of `CustomizedCfgMapInput` via:
//
//	CustomizedCfgMap{ "key": CustomizedCfgArgs{...} }
type CustomizedCfgMapInput interface {
	pulumi.Input

	ToCustomizedCfgMapOutput() CustomizedCfgMapOutput
	ToCustomizedCfgMapOutputWithContext(context.Context) CustomizedCfgMapOutput
}

type CustomizedCfgMap map[string]CustomizedCfgInput

func (CustomizedCfgMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomizedCfg)(nil)).Elem()
}

func (i CustomizedCfgMap) ToCustomizedCfgMapOutput() CustomizedCfgMapOutput {
	return i.ToCustomizedCfgMapOutputWithContext(context.Background())
}

func (i CustomizedCfgMap) ToCustomizedCfgMapOutputWithContext(ctx context.Context) CustomizedCfgMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomizedCfgMapOutput)
}

type CustomizedCfgOutput struct{ *pulumi.OutputState }

func (CustomizedCfgOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomizedCfg)(nil)).Elem()
}

func (o CustomizedCfgOutput) ToCustomizedCfgOutput() CustomizedCfgOutput {
	return o
}

func (o CustomizedCfgOutput) ToCustomizedCfgOutputWithContext(ctx context.Context) CustomizedCfgOutput {
	return o
}

// The content of CustomizedCfg. The length cannot exceed 4096 characters. Spaces and semicolons need to be escaped. Currently supported configuration items are `sslProtocols`, `sslCiphers`, `clientMaxBodySize`, `keepaliveTimeout`, `proxyRequestBuffering` and `proxyConnectTimeout`.
func (o CustomizedCfgOutput) CustomizedCfgContent() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedCfg) pulumi.StringOutput { return v.CustomizedCfgContent }).(pulumi.StringOutput)
}

// The name of CustomizedCfg.
func (o CustomizedCfgOutput) CustomizedCfgName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedCfg) pulumi.StringOutput { return v.CustomizedCfgName }).(pulumi.StringOutput)
}

// The description of CustomizedCfg.
func (o CustomizedCfgOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedCfg) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The project name of the CustomizedCfg.
func (o CustomizedCfgOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomizedCfg) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

type CustomizedCfgArrayOutput struct{ *pulumi.OutputState }

func (CustomizedCfgArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomizedCfg)(nil)).Elem()
}

func (o CustomizedCfgArrayOutput) ToCustomizedCfgArrayOutput() CustomizedCfgArrayOutput {
	return o
}

func (o CustomizedCfgArrayOutput) ToCustomizedCfgArrayOutputWithContext(ctx context.Context) CustomizedCfgArrayOutput {
	return o
}

func (o CustomizedCfgArrayOutput) Index(i pulumi.IntInput) CustomizedCfgOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomizedCfg {
		return vs[0].([]*CustomizedCfg)[vs[1].(int)]
	}).(CustomizedCfgOutput)
}

type CustomizedCfgMapOutput struct{ *pulumi.OutputState }

func (CustomizedCfgMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomizedCfg)(nil)).Elem()
}

func (o CustomizedCfgMapOutput) ToCustomizedCfgMapOutput() CustomizedCfgMapOutput {
	return o
}

func (o CustomizedCfgMapOutput) ToCustomizedCfgMapOutputWithContext(ctx context.Context) CustomizedCfgMapOutput {
	return o
}

func (o CustomizedCfgMapOutput) MapIndex(k pulumi.StringInput) CustomizedCfgOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomizedCfg {
		return vs[0].(map[string]*CustomizedCfg)[vs[1].(string)]
	}).(CustomizedCfgOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomizedCfgInput)(nil)).Elem(), &CustomizedCfg{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomizedCfgArrayInput)(nil)).Elem(), CustomizedCfgArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomizedCfgMapInput)(nil)).Elem(), CustomizedCfgMap{})
	pulumi.RegisterOutputType(CustomizedCfgOutput{})
	pulumi.RegisterOutputType(CustomizedCfgArrayOutput{})
	pulumi.RegisterOutputType(CustomizedCfgMapOutput{})
}
