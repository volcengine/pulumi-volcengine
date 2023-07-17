// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage server group
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clb.NewServerGroup(ctx, "foo", &clb.ServerGroupArgs{
//				Description:     pulumi.String("hello demo11"),
//				LoadBalancerId:  pulumi.String("clb-273z7d4r8tvk07fap8tsniyfe"),
//				ServerGroupName: pulumi.String("demo-demo11"),
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
// ServerGroup can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:clb/serverGroup:ServerGroup default rsp-273yv0kir1vk07fap8tt9jtwg
//
// ```
type ServerGroup struct {
	pulumi.CustomResourceState

	// The description of ServerGroup.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ID of the Clb.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringOutput `pulumi:"serverGroupId"`
	// The name of the ServerGroup.
	ServerGroupName pulumi.StringOutput `pulumi:"serverGroupName"`
}

// NewServerGroup registers a new resource with the given unique name, arguments, and options.
func NewServerGroup(ctx *pulumi.Context,
	name string, args *ServerGroupArgs, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerId == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerId'")
	}
	var resource ServerGroup
	err := ctx.RegisterResource("volcengine:clb/serverGroup:ServerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerGroup gets an existing ServerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupState, opts ...pulumi.ResourceOption) (*ServerGroup, error) {
	var resource ServerGroup
	err := ctx.ReadResource("volcengine:clb/serverGroup:ServerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroup resources.
type serverGroupState struct {
	// The description of ServerGroup.
	Description *string `pulumi:"description"`
	// The ID of the Clb.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// The ID of the ServerGroup.
	ServerGroupId *string `pulumi:"serverGroupId"`
	// The name of the ServerGroup.
	ServerGroupName *string `pulumi:"serverGroupName"`
}

type ServerGroupState struct {
	// The description of ServerGroup.
	Description pulumi.StringPtrInput
	// The ID of the Clb.
	LoadBalancerId pulumi.StringPtrInput
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringPtrInput
	// The name of the ServerGroup.
	ServerGroupName pulumi.StringPtrInput
}

func (ServerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupState)(nil)).Elem()
}

type serverGroupArgs struct {
	// The description of ServerGroup.
	Description *string `pulumi:"description"`
	// The ID of the Clb.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// The ID of the ServerGroup.
	ServerGroupId *string `pulumi:"serverGroupId"`
	// The name of the ServerGroup.
	ServerGroupName *string `pulumi:"serverGroupName"`
}

// The set of arguments for constructing a ServerGroup resource.
type ServerGroupArgs struct {
	// The description of ServerGroup.
	Description pulumi.StringPtrInput
	// The ID of the Clb.
	LoadBalancerId pulumi.StringInput
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringPtrInput
	// The name of the ServerGroup.
	ServerGroupName pulumi.StringPtrInput
}

func (ServerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupArgs)(nil)).Elem()
}

type ServerGroupInput interface {
	pulumi.Input

	ToServerGroupOutput() ServerGroupOutput
	ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput
}

func (*ServerGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (i *ServerGroup) ToServerGroupOutput() ServerGroupOutput {
	return i.ToServerGroupOutputWithContext(context.Background())
}

func (i *ServerGroup) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupOutput)
}

// ServerGroupArrayInput is an input type that accepts ServerGroupArray and ServerGroupArrayOutput values.
// You can construct a concrete instance of `ServerGroupArrayInput` via:
//
//	ServerGroupArray{ ServerGroupArgs{...} }
type ServerGroupArrayInput interface {
	pulumi.Input

	ToServerGroupArrayOutput() ServerGroupArrayOutput
	ToServerGroupArrayOutputWithContext(context.Context) ServerGroupArrayOutput
}

type ServerGroupArray []ServerGroupInput

func (ServerGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupArray) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return i.ToServerGroupArrayOutputWithContext(context.Background())
}

func (i ServerGroupArray) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupArrayOutput)
}

// ServerGroupMapInput is an input type that accepts ServerGroupMap and ServerGroupMapOutput values.
// You can construct a concrete instance of `ServerGroupMapInput` via:
//
//	ServerGroupMap{ "key": ServerGroupArgs{...} }
type ServerGroupMapInput interface {
	pulumi.Input

	ToServerGroupMapOutput() ServerGroupMapOutput
	ToServerGroupMapOutputWithContext(context.Context) ServerGroupMapOutput
}

type ServerGroupMap map[string]ServerGroupInput

func (ServerGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (i ServerGroupMap) ToServerGroupMapOutput() ServerGroupMapOutput {
	return i.ToServerGroupMapOutputWithContext(context.Background())
}

func (i ServerGroupMap) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupMapOutput)
}

type ServerGroupOutput struct{ *pulumi.OutputState }

func (ServerGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroup)(nil)).Elem()
}

func (o ServerGroupOutput) ToServerGroupOutput() ServerGroupOutput {
	return o
}

func (o ServerGroupOutput) ToServerGroupOutputWithContext(ctx context.Context) ServerGroupOutput {
	return o
}

// The description of ServerGroup.
func (o ServerGroupOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The ID of the Clb.
func (o ServerGroupOutput) LoadBalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.LoadBalancerId }).(pulumi.StringOutput)
}

// The ID of the ServerGroup.
func (o ServerGroupOutput) ServerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ServerGroupId }).(pulumi.StringOutput)
}

// The name of the ServerGroup.
func (o ServerGroupOutput) ServerGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroup) pulumi.StringOutput { return v.ServerGroupName }).(pulumi.StringOutput)
}

type ServerGroupArrayOutput struct{ *pulumi.OutputState }

func (ServerGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutput() ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) ToServerGroupArrayOutputWithContext(ctx context.Context) ServerGroupArrayOutput {
	return o
}

func (o ServerGroupArrayOutput) Index(i pulumi.IntInput) ServerGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].([]*ServerGroup)[vs[1].(int)]
	}).(ServerGroupOutput)
}

type ServerGroupMapOutput struct{ *pulumi.OutputState }

func (ServerGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroup)(nil)).Elem()
}

func (o ServerGroupMapOutput) ToServerGroupMapOutput() ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) ToServerGroupMapOutputWithContext(ctx context.Context) ServerGroupMapOutput {
	return o
}

func (o ServerGroupMapOutput) MapIndex(k pulumi.StringInput) ServerGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerGroup {
		return vs[0].(map[string]*ServerGroup)[vs[1].(string)]
	}).(ServerGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupInput)(nil)).Elem(), &ServerGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupArrayInput)(nil)).Elem(), ServerGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupMapInput)(nil)).Elem(), ServerGroupMap{})
	pulumi.RegisterOutputType(ServerGroupOutput{})
	pulumi.RegisterOutputType(ServerGroupArrayOutput{})
	pulumi.RegisterOutputType(ServerGroupMapOutput{})
}