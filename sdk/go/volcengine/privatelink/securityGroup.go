// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage privatelink security group
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/privatelink"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := privatelink.NewSecurityGroup(ctx, "foo", &privatelink.SecurityGroupArgs{
//				EndpointId:      pulumi.String("ep-2byz5npiuu1hc2dx0efkv7ehc"),
//				SecurityGroupId: pulumi.String("sg-2d6722jpp55og58ozfd1sqtdb"),
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
// PrivateLink Security Group Service can be imported using the endpoint id and security group id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:privatelink/securityGroup:SecurityGroup default ep-2fe630gurkl37k5gfuy33****:sg-xxxxx
//
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// The id of the endpoint.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// The id of the security group.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("volcengine:privatelink/securityGroup:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("volcengine:privatelink/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// The id of the endpoint.
	EndpointId *string `pulumi:"endpointId"`
	// The id of the security group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
}

type SecurityGroupState struct {
	// The id of the endpoint.
	EndpointId pulumi.StringPtrInput
	// The id of the security group.
	SecurityGroupId pulumi.StringPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// The id of the endpoint.
	EndpointId string `pulumi:"endpointId"`
	// The id of the security group.
	SecurityGroupId string `pulumi:"securityGroupId"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// The id of the endpoint.
	EndpointId pulumi.StringInput
	// The id of the security group.
	SecurityGroupId pulumi.StringInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

type SecurityGroupInput interface {
	pulumi.Input

	ToSecurityGroupOutput() SecurityGroupOutput
	ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput
}

func (*SecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (i *SecurityGroup) ToSecurityGroupOutput() SecurityGroupOutput {
	return i.ToSecurityGroupOutputWithContext(context.Background())
}

func (i *SecurityGroup) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupOutput)
}

// SecurityGroupArrayInput is an input type that accepts SecurityGroupArray and SecurityGroupArrayOutput values.
// You can construct a concrete instance of `SecurityGroupArrayInput` via:
//
//	SecurityGroupArray{ SecurityGroupArgs{...} }
type SecurityGroupArrayInput interface {
	pulumi.Input

	ToSecurityGroupArrayOutput() SecurityGroupArrayOutput
	ToSecurityGroupArrayOutputWithContext(context.Context) SecurityGroupArrayOutput
}

type SecurityGroupArray []SecurityGroupInput

func (SecurityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return i.ToSecurityGroupArrayOutputWithContext(context.Background())
}

func (i SecurityGroupArray) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupArrayOutput)
}

// SecurityGroupMapInput is an input type that accepts SecurityGroupMap and SecurityGroupMapOutput values.
// You can construct a concrete instance of `SecurityGroupMapInput` via:
//
//	SecurityGroupMap{ "key": SecurityGroupArgs{...} }
type SecurityGroupMapInput interface {
	pulumi.Input

	ToSecurityGroupMapOutput() SecurityGroupMapOutput
	ToSecurityGroupMapOutputWithContext(context.Context) SecurityGroupMapOutput
}

type SecurityGroupMap map[string]SecurityGroupInput

func (SecurityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroupMap) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return i.ToSecurityGroupMapOutputWithContext(context.Background())
}

func (i SecurityGroupMap) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupMapOutput)
}

type SecurityGroupOutput struct{ *pulumi.OutputState }

func (SecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupOutput) ToSecurityGroupOutput() SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return o
}

// The id of the endpoint.
func (o SecurityGroupOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// The id of the security group.
func (o SecurityGroupOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

type SecurityGroupArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutput() SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) ToSecurityGroupArrayOutputWithContext(ctx context.Context) SecurityGroupArrayOutput {
	return o
}

func (o SecurityGroupArrayOutput) Index(i pulumi.IntInput) SecurityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityGroup {
		return vs[0].([]*SecurityGroup)[vs[1].(int)]
	}).(SecurityGroupOutput)
}

type SecurityGroupMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroup)(nil)).Elem()
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutput() SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) ToSecurityGroupMapOutputWithContext(ctx context.Context) SecurityGroupMapOutput {
	return o
}

func (o SecurityGroupMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityGroup {
		return vs[0].(map[string]*SecurityGroup)[vs[1].(string)]
	}).(SecurityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupInput)(nil)).Elem(), &SecurityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupArrayInput)(nil)).Elem(), SecurityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupMapInput)(nil)).Elem(), SecurityGroupMap{})
	pulumi.RegisterOutputType(SecurityGroupOutput{})
	pulumi.RegisterOutputType(SecurityGroupArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupMapOutput{})
}