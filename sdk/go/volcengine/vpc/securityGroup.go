// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage security group
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NewSecurityGroup(ctx, "g1test1", &vpc.SecurityGroupArgs{
//				ProjectName: pulumi.String("yuwenhao"),
//				VpcId:       pulumi.String("vpc-2feppmy1ugt1c59gp688n1fld"),
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
// SecurityGroup can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:vpc/securityGroup:SecurityGroup default sg-273ycgql3ig3k7fap8t3dyvqx
//
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// Creation time of SecurityGroup.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Description of SecurityGroup.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ProjectName of SecurityGroup.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// Name of SecurityGroup.
	SecurityGroupName pulumi.StringOutput `pulumi:"securityGroupName"`
	// Status of SecurityGroup.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags.
	Tags SecurityGroupTagArrayOutput `pulumi:"tags"`
	// Id of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("volcengine:vpc/securityGroup:SecurityGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("volcengine:vpc/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// Creation time of SecurityGroup.
	CreationTime *string `pulumi:"creationTime"`
	// Description of SecurityGroup.
	Description *string `pulumi:"description"`
	// The ProjectName of SecurityGroup.
	ProjectName *string `pulumi:"projectName"`
	// Name of SecurityGroup.
	SecurityGroupName *string `pulumi:"securityGroupName"`
	// Status of SecurityGroup.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []SecurityGroupTag `pulumi:"tags"`
	// Id of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

type SecurityGroupState struct {
	// Creation time of SecurityGroup.
	CreationTime pulumi.StringPtrInput
	// Description of SecurityGroup.
	Description pulumi.StringPtrInput
	// The ProjectName of SecurityGroup.
	ProjectName pulumi.StringPtrInput
	// Name of SecurityGroup.
	SecurityGroupName pulumi.StringPtrInput
	// Status of SecurityGroup.
	Status pulumi.StringPtrInput
	// Tags.
	Tags SecurityGroupTagArrayInput
	// Id of the VPC.
	VpcId pulumi.StringPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// Description of SecurityGroup.
	Description *string `pulumi:"description"`
	// The ProjectName of SecurityGroup.
	ProjectName *string `pulumi:"projectName"`
	// Name of SecurityGroup.
	SecurityGroupName *string `pulumi:"securityGroupName"`
	// Tags.
	Tags []SecurityGroupTag `pulumi:"tags"`
	// Id of the VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// Description of SecurityGroup.
	Description pulumi.StringPtrInput
	// The ProjectName of SecurityGroup.
	ProjectName pulumi.StringPtrInput
	// Name of SecurityGroup.
	SecurityGroupName pulumi.StringPtrInput
	// Tags.
	Tags SecurityGroupTagArrayInput
	// Id of the VPC.
	VpcId pulumi.StringInput
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

// Creation time of SecurityGroup.
func (o SecurityGroupOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Description of SecurityGroup.
func (o SecurityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ProjectName of SecurityGroup.
func (o SecurityGroupOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Name of SecurityGroup.
func (o SecurityGroupOutput) SecurityGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.SecurityGroupName }).(pulumi.StringOutput)
}

// Status of SecurityGroup.
func (o SecurityGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags.
func (o SecurityGroupOutput) Tags() SecurityGroupTagArrayOutput {
	return o.ApplyT(func(v *SecurityGroup) SecurityGroupTagArrayOutput { return v.Tags }).(SecurityGroupTagArrayOutput)
}

// Id of the VPC.
func (o SecurityGroupOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroup) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
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