// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage iam role attachment
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/iam"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := ecs.Zones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:   pulumi.String("acc-test-vpc"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     *pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooSecurityGroup, err := vpc.NewSecurityGroup(ctx, "fooSecurityGroup", &vpc.SecurityGroupArgs{
//				SecurityGroupName: pulumi.String("acc-test-security-group"),
//				VpcId:             fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooImages, err := ecs.Images(ctx, &ecs.ImagesArgs{
//				OsType:         pulumi.StringRef("Linux"),
//				Visibility:     pulumi.StringRef("public"),
//				InstanceTypeId: pulumi.StringRef("ecs.g1ie.large"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			fooInstance, err := ecs.NewInstance(ctx, "fooInstance", &ecs.InstanceArgs{
//				InstanceName:       pulumi.String("acc-test-ecs"),
//				Description:        pulumi.String("acc-test"),
//				HostName:           pulumi.String("tf-acc-test"),
//				ImageId:            *pulumi.String(fooImages.Images[0].ImageId),
//				InstanceType:       pulumi.String("ecs.g1ie.large"),
//				Password:           pulumi.String("93f0cb0614Aab12"),
//				InstanceChargeType: pulumi.String("PostPaid"),
//				SystemVolumeType:   pulumi.String("ESSD_PL0"),
//				SystemVolumeSize:   pulumi.Int(40),
//				DataVolumes: ecs.InstanceDataVolumeArray{
//					&ecs.InstanceDataVolumeArgs{
//						VolumeType:         pulumi.String("ESSD_PL0"),
//						Size:               pulumi.Int(50),
//						DeleteWithInstance: pulumi.Bool(true),
//					},
//				},
//				SubnetId: fooSubnet.ID(),
//				SecurityGroupIds: pulumi.StringArray{
//					fooSecurityGroup.ID(),
//				},
//				ProjectName: pulumi.String("default"),
//				Tags: ecs.InstanceTagArray{
//					&ecs.InstanceTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooRole, err := iam.NewRole(ctx, "fooRole", &iam.RoleArgs{
//				RoleName:            pulumi.String("acc-test-role"),
//				DisplayName:         pulumi.String("acc-test"),
//				Description:         pulumi.String("acc-test"),
//				TrustPolicyDocument: pulumi.String("{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"ecs\"]}}]}"),
//				MaxSessionDuration:  pulumi.Int(36000),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ecs.NewIamRoleAttachment(ctx, "fooIamRoleAttachment", &ecs.IamRoleAttachmentArgs{
//				IamRoleName: fooRole.ID(),
//				InstanceId:  fooInstance.ID(),
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
// IamRoleAttachment can be imported using the iam_role_name:instance_id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:ecs/iamRoleAttachment:IamRoleAttachment default role_name:instance_id
//
// ```
type IamRoleAttachment struct {
	pulumi.CustomResourceState

	// The name of the iam role.
	IamRoleName pulumi.StringOutput `pulumi:"iamRoleName"`
	// The id of the ecs instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewIamRoleAttachment registers a new resource with the given unique name, arguments, and options.
func NewIamRoleAttachment(ctx *pulumi.Context,
	name string, args *IamRoleAttachmentArgs, opts ...pulumi.ResourceOption) (*IamRoleAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IamRoleName == nil {
		return nil, errors.New("invalid value for required argument 'IamRoleName'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IamRoleAttachment
	err := ctx.RegisterResource("volcengine:ecs/iamRoleAttachment:IamRoleAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamRoleAttachment gets an existing IamRoleAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamRoleAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamRoleAttachmentState, opts ...pulumi.ResourceOption) (*IamRoleAttachment, error) {
	var resource IamRoleAttachment
	err := ctx.ReadResource("volcengine:ecs/iamRoleAttachment:IamRoleAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamRoleAttachment resources.
type iamRoleAttachmentState struct {
	// The name of the iam role.
	IamRoleName *string `pulumi:"iamRoleName"`
	// The id of the ecs instance.
	InstanceId *string `pulumi:"instanceId"`
}

type IamRoleAttachmentState struct {
	// The name of the iam role.
	IamRoleName pulumi.StringPtrInput
	// The id of the ecs instance.
	InstanceId pulumi.StringPtrInput
}

func (IamRoleAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamRoleAttachmentState)(nil)).Elem()
}

type iamRoleAttachmentArgs struct {
	// The name of the iam role.
	IamRoleName string `pulumi:"iamRoleName"`
	// The id of the ecs instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a IamRoleAttachment resource.
type IamRoleAttachmentArgs struct {
	// The name of the iam role.
	IamRoleName pulumi.StringInput
	// The id of the ecs instance.
	InstanceId pulumi.StringInput
}

func (IamRoleAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamRoleAttachmentArgs)(nil)).Elem()
}

type IamRoleAttachmentInput interface {
	pulumi.Input

	ToIamRoleAttachmentOutput() IamRoleAttachmentOutput
	ToIamRoleAttachmentOutputWithContext(ctx context.Context) IamRoleAttachmentOutput
}

func (*IamRoleAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**IamRoleAttachment)(nil)).Elem()
}

func (i *IamRoleAttachment) ToIamRoleAttachmentOutput() IamRoleAttachmentOutput {
	return i.ToIamRoleAttachmentOutputWithContext(context.Background())
}

func (i *IamRoleAttachment) ToIamRoleAttachmentOutputWithContext(ctx context.Context) IamRoleAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamRoleAttachmentOutput)
}

// IamRoleAttachmentArrayInput is an input type that accepts IamRoleAttachmentArray and IamRoleAttachmentArrayOutput values.
// You can construct a concrete instance of `IamRoleAttachmentArrayInput` via:
//
//	IamRoleAttachmentArray{ IamRoleAttachmentArgs{...} }
type IamRoleAttachmentArrayInput interface {
	pulumi.Input

	ToIamRoleAttachmentArrayOutput() IamRoleAttachmentArrayOutput
	ToIamRoleAttachmentArrayOutputWithContext(context.Context) IamRoleAttachmentArrayOutput
}

type IamRoleAttachmentArray []IamRoleAttachmentInput

func (IamRoleAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamRoleAttachment)(nil)).Elem()
}

func (i IamRoleAttachmentArray) ToIamRoleAttachmentArrayOutput() IamRoleAttachmentArrayOutput {
	return i.ToIamRoleAttachmentArrayOutputWithContext(context.Background())
}

func (i IamRoleAttachmentArray) ToIamRoleAttachmentArrayOutputWithContext(ctx context.Context) IamRoleAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamRoleAttachmentArrayOutput)
}

// IamRoleAttachmentMapInput is an input type that accepts IamRoleAttachmentMap and IamRoleAttachmentMapOutput values.
// You can construct a concrete instance of `IamRoleAttachmentMapInput` via:
//
//	IamRoleAttachmentMap{ "key": IamRoleAttachmentArgs{...} }
type IamRoleAttachmentMapInput interface {
	pulumi.Input

	ToIamRoleAttachmentMapOutput() IamRoleAttachmentMapOutput
	ToIamRoleAttachmentMapOutputWithContext(context.Context) IamRoleAttachmentMapOutput
}

type IamRoleAttachmentMap map[string]IamRoleAttachmentInput

func (IamRoleAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamRoleAttachment)(nil)).Elem()
}

func (i IamRoleAttachmentMap) ToIamRoleAttachmentMapOutput() IamRoleAttachmentMapOutput {
	return i.ToIamRoleAttachmentMapOutputWithContext(context.Background())
}

func (i IamRoleAttachmentMap) ToIamRoleAttachmentMapOutputWithContext(ctx context.Context) IamRoleAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamRoleAttachmentMapOutput)
}

type IamRoleAttachmentOutput struct{ *pulumi.OutputState }

func (IamRoleAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamRoleAttachment)(nil)).Elem()
}

func (o IamRoleAttachmentOutput) ToIamRoleAttachmentOutput() IamRoleAttachmentOutput {
	return o
}

func (o IamRoleAttachmentOutput) ToIamRoleAttachmentOutputWithContext(ctx context.Context) IamRoleAttachmentOutput {
	return o
}

// The name of the iam role.
func (o IamRoleAttachmentOutput) IamRoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *IamRoleAttachment) pulumi.StringOutput { return v.IamRoleName }).(pulumi.StringOutput)
}

// The id of the ecs instance.
func (o IamRoleAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IamRoleAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type IamRoleAttachmentArrayOutput struct{ *pulumi.OutputState }

func (IamRoleAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IamRoleAttachment)(nil)).Elem()
}

func (o IamRoleAttachmentArrayOutput) ToIamRoleAttachmentArrayOutput() IamRoleAttachmentArrayOutput {
	return o
}

func (o IamRoleAttachmentArrayOutput) ToIamRoleAttachmentArrayOutputWithContext(ctx context.Context) IamRoleAttachmentArrayOutput {
	return o
}

func (o IamRoleAttachmentArrayOutput) Index(i pulumi.IntInput) IamRoleAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IamRoleAttachment {
		return vs[0].([]*IamRoleAttachment)[vs[1].(int)]
	}).(IamRoleAttachmentOutput)
}

type IamRoleAttachmentMapOutput struct{ *pulumi.OutputState }

func (IamRoleAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IamRoleAttachment)(nil)).Elem()
}

func (o IamRoleAttachmentMapOutput) ToIamRoleAttachmentMapOutput() IamRoleAttachmentMapOutput {
	return o
}

func (o IamRoleAttachmentMapOutput) ToIamRoleAttachmentMapOutputWithContext(ctx context.Context) IamRoleAttachmentMapOutput {
	return o
}

func (o IamRoleAttachmentMapOutput) MapIndex(k pulumi.StringInput) IamRoleAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IamRoleAttachment {
		return vs[0].(map[string]*IamRoleAttachment)[vs[1].(string)]
	}).(IamRoleAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IamRoleAttachmentInput)(nil)).Elem(), &IamRoleAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamRoleAttachmentArrayInput)(nil)).Elem(), IamRoleAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IamRoleAttachmentMapInput)(nil)).Elem(), IamRoleAttachmentMap{})
	pulumi.RegisterOutputType(IamRoleAttachmentOutput{})
	pulumi.RegisterOutputType(IamRoleAttachmentArrayOutput{})
	pulumi.RegisterOutputType(IamRoleAttachmentMapOutput{})
}