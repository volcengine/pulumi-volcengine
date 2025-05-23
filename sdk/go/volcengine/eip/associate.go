// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eip

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage eip associate
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/eip"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := ecs.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooImages, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
//				OsType:         pulumi.StringRef("Linux"),
//				Visibility:     pulumi.StringRef("public"),
//				InstanceTypeId: pulumi.StringRef("ecs.g1.large"),
//			}, nil)
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
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooSecurityGroup, err := vpc.NewSecurityGroup(ctx, "fooSecurityGroup", &vpc.SecurityGroupArgs{
//				VpcId:             fooVpc.ID(),
//				SecurityGroupName: pulumi.String("acc-test-security-group"),
//			})
//			if err != nil {
//				return err
//			}
//			fooInstance, err := ecs.NewInstance(ctx, "fooInstance", &ecs.InstanceArgs{
//				ImageId:            pulumi.String(fooImages.Images[0].ImageId),
//				InstanceType:       pulumi.String("ecs.g1.large"),
//				InstanceName:       pulumi.String("acc-test-ecs-name"),
//				Password:           pulumi.String("93f0cb0614Aab12"),
//				InstanceChargeType: pulumi.String("PostPaid"),
//				SystemVolumeType:   pulumi.String("ESSD_PL0"),
//				SystemVolumeSize:   pulumi.Int(40),
//				SubnetId:           fooSubnet.ID(),
//				SecurityGroupIds: pulumi.StringArray{
//					fooSecurityGroup.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooAddress, err := eip.NewAddress(ctx, "fooAddress", &eip.AddressArgs{
//				BillingType: pulumi.String("PostPaidByTraffic"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = eip.NewAssociate(ctx, "fooAssociate", &eip.AssociateArgs{
//				AllocationId: fooAddress.ID(),
//				InstanceId:   fooInstance.ID(),
//				InstanceType: pulumi.String("EcsInstance"),
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
// Eip associate can be imported using the eip allocation_id:instance_id, e.g.
//
// ```sh
// $ pulumi import volcengine:eip/associate:Associate default eip-274oj9a8rs9a87fap8sf9515b:i-cm9t9ug9lggu79yr5tcw
// ```
type Associate struct {
	pulumi.CustomResourceState

	// The allocation id of the EIP.
	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// The instance id which be associated to the EIP.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The private IP address of the instance will be associated to the EIP.
	PrivateIpAddress pulumi.StringOutput `pulumi:"privateIpAddress"`
}

// NewAssociate registers a new resource with the given unique name, arguments, and options.
func NewAssociate(ctx *pulumi.Context,
	name string, args *AssociateArgs, opts ...pulumi.ResourceOption) (*Associate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllocationId == nil {
		return nil, errors.New("invalid value for required argument 'AllocationId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Associate
	err := ctx.RegisterResource("volcengine:eip/associate:Associate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssociate gets an existing Associate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssociateState, opts ...pulumi.ResourceOption) (*Associate, error) {
	var resource Associate
	err := ctx.ReadResource("volcengine:eip/associate:Associate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Associate resources.
type associateState struct {
	// The allocation id of the EIP.
	AllocationId *string `pulumi:"allocationId"`
	// The instance id which be associated to the EIP.
	InstanceId *string `pulumi:"instanceId"`
	// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
	InstanceType *string `pulumi:"instanceType"`
	// The private IP address of the instance will be associated to the EIP.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
}

type AssociateState struct {
	// The allocation id of the EIP.
	AllocationId pulumi.StringPtrInput
	// The instance id which be associated to the EIP.
	InstanceId pulumi.StringPtrInput
	// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
	InstanceType pulumi.StringPtrInput
	// The private IP address of the instance will be associated to the EIP.
	PrivateIpAddress pulumi.StringPtrInput
}

func (AssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*associateState)(nil)).Elem()
}

type associateArgs struct {
	// The allocation id of the EIP.
	AllocationId string `pulumi:"allocationId"`
	// The instance id which be associated to the EIP.
	InstanceId string `pulumi:"instanceId"`
	// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
	InstanceType string `pulumi:"instanceType"`
	// The private IP address of the instance will be associated to the EIP.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
}

// The set of arguments for constructing a Associate resource.
type AssociateArgs struct {
	// The allocation id of the EIP.
	AllocationId pulumi.StringInput
	// The instance id which be associated to the EIP.
	InstanceId pulumi.StringInput
	// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
	InstanceType pulumi.StringInput
	// The private IP address of the instance will be associated to the EIP.
	PrivateIpAddress pulumi.StringPtrInput
}

func (AssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*associateArgs)(nil)).Elem()
}

type AssociateInput interface {
	pulumi.Input

	ToAssociateOutput() AssociateOutput
	ToAssociateOutputWithContext(ctx context.Context) AssociateOutput
}

func (*Associate) ElementType() reflect.Type {
	return reflect.TypeOf((**Associate)(nil)).Elem()
}

func (i *Associate) ToAssociateOutput() AssociateOutput {
	return i.ToAssociateOutputWithContext(context.Background())
}

func (i *Associate) ToAssociateOutputWithContext(ctx context.Context) AssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociateOutput)
}

// AssociateArrayInput is an input type that accepts AssociateArray and AssociateArrayOutput values.
// You can construct a concrete instance of `AssociateArrayInput` via:
//
//	AssociateArray{ AssociateArgs{...} }
type AssociateArrayInput interface {
	pulumi.Input

	ToAssociateArrayOutput() AssociateArrayOutput
	ToAssociateArrayOutputWithContext(context.Context) AssociateArrayOutput
}

type AssociateArray []AssociateInput

func (AssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Associate)(nil)).Elem()
}

func (i AssociateArray) ToAssociateArrayOutput() AssociateArrayOutput {
	return i.ToAssociateArrayOutputWithContext(context.Background())
}

func (i AssociateArray) ToAssociateArrayOutputWithContext(ctx context.Context) AssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociateArrayOutput)
}

// AssociateMapInput is an input type that accepts AssociateMap and AssociateMapOutput values.
// You can construct a concrete instance of `AssociateMapInput` via:
//
//	AssociateMap{ "key": AssociateArgs{...} }
type AssociateMapInput interface {
	pulumi.Input

	ToAssociateMapOutput() AssociateMapOutput
	ToAssociateMapOutputWithContext(context.Context) AssociateMapOutput
}

type AssociateMap map[string]AssociateInput

func (AssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Associate)(nil)).Elem()
}

func (i AssociateMap) ToAssociateMapOutput() AssociateMapOutput {
	return i.ToAssociateMapOutputWithContext(context.Background())
}

func (i AssociateMap) ToAssociateMapOutputWithContext(ctx context.Context) AssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociateMapOutput)
}

type AssociateOutput struct{ *pulumi.OutputState }

func (AssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Associate)(nil)).Elem()
}

func (o AssociateOutput) ToAssociateOutput() AssociateOutput {
	return o
}

func (o AssociateOutput) ToAssociateOutputWithContext(ctx context.Context) AssociateOutput {
	return o
}

// The allocation id of the EIP.
func (o AssociateOutput) AllocationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Associate) pulumi.StringOutput { return v.AllocationId }).(pulumi.StringOutput)
}

// The instance id which be associated to the EIP.
func (o AssociateOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Associate) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
func (o AssociateOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Associate) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The private IP address of the instance will be associated to the EIP.
func (o AssociateOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Associate) pulumi.StringOutput { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

type AssociateArrayOutput struct{ *pulumi.OutputState }

func (AssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Associate)(nil)).Elem()
}

func (o AssociateArrayOutput) ToAssociateArrayOutput() AssociateArrayOutput {
	return o
}

func (o AssociateArrayOutput) ToAssociateArrayOutputWithContext(ctx context.Context) AssociateArrayOutput {
	return o
}

func (o AssociateArrayOutput) Index(i pulumi.IntInput) AssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Associate {
		return vs[0].([]*Associate)[vs[1].(int)]
	}).(AssociateOutput)
}

type AssociateMapOutput struct{ *pulumi.OutputState }

func (AssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Associate)(nil)).Elem()
}

func (o AssociateMapOutput) ToAssociateMapOutput() AssociateMapOutput {
	return o
}

func (o AssociateMapOutput) ToAssociateMapOutputWithContext(ctx context.Context) AssociateMapOutput {
	return o
}

func (o AssociateMapOutput) MapIndex(k pulumi.StringInput) AssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Associate {
		return vs[0].(map[string]*Associate)[vs[1].(string)]
	}).(AssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssociateInput)(nil)).Elem(), &Associate{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssociateArrayInput)(nil)).Elem(), AssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssociateMapInput)(nil)).Elem(), AssociateMap{})
	pulumi.RegisterOutputType(AssociateOutput{})
	pulumi.RegisterOutputType(AssociateArrayOutput{})
	pulumi.RegisterOutputType(AssociateMapOutput{})
}
