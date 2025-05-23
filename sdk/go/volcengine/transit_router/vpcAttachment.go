// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package transit_router

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage transit router vpc attachment
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/transit_router"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooTransitRouter, err := transit_router.NewTransitRouter(ctx, "fooTransitRouter", &transit_router.TransitRouterArgs{
//				TransitRouterName: pulumi.String("test-tf-acc"),
//				Description:       pulumi.String("test-tf-acc"),
//				Asn:               pulumi.Int(4294967293),
//			})
//			if err != nil {
//				return err
//			}
//			fooZones, err := ecs.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:   pulumi.String("acc-test-vpc-acc"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				VpcId:      fooVpc.ID(),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				SubnetName: pulumi.String("acc-test-subnet"),
//			})
//			if err != nil {
//				return err
//			}
//			foo2, err := vpc.NewSubnet(ctx, "foo2", &vpc.SubnetArgs{
//				VpcId:      fooVpc.ID(),
//				CidrBlock:  pulumi.String("172.16.255.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[1].Id),
//				SubnetName: pulumi.String("acc-test-subnet2"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = transit_router.NewVpcAttachment(ctx, "fooVpcAttachment", &transit_router.VpcAttachmentArgs{
//				TransitRouterId: fooTransitRouter.ID(),
//				VpcId:           fooVpc.ID(),
//				AttachPoints: transit_router.VpcAttachmentAttachPointArray{
//					&transit_router.VpcAttachmentAttachPointArgs{
//						SubnetId: fooSubnet.ID(),
//						ZoneId:   pulumi.String("cn-beijing-a"),
//					},
//					&transit_router.VpcAttachmentAttachPointArgs{
//						SubnetId: foo2.ID(),
//						ZoneId:   pulumi.String("cn-beijing-b"),
//					},
//				},
//				TransitRouterAttachmentName: pulumi.String("tf-test-acc-vpc-attach"),
//				Description:                 pulumi.String("tf-test-acc-description"),
//				AutoPublishRouteEnabled:     pulumi.Bool(true),
//				Tags: transit_router.VpcAttachmentTagArray{
//					&transit_router.VpcAttachmentTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
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
// TransitRouterVpcAttachment can be imported using the transitRouterId:attachmentId, e.g.
//
// ```sh
// $ pulumi import volcengine:transit_router/vpcAttachment:VpcAttachment default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
// ```
type VpcAttachment struct {
	pulumi.CustomResourceState

	// The attach points of transit router vpc attachment.
	AttachPoints VpcAttachmentAttachPointArrayOutput `pulumi:"attachPoints"`
	// Whether to auto publish route of the transit router to vpc instance. Default is false.
	AutoPublishRouteEnabled pulumi.BoolPtrOutput `pulumi:"autoPublishRouteEnabled"`
	// The create time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The description of the transit router vpc attachment.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The status of the transit router.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags.
	Tags VpcAttachmentTagArrayOutput `pulumi:"tags"`
	// The id of the transit router attachment.
	TransitRouterAttachmentId pulumi.StringOutput `pulumi:"transitRouterAttachmentId"`
	// The name of the transit router vpc attachment.
	TransitRouterAttachmentName pulumi.StringOutput `pulumi:"transitRouterAttachmentName"`
	// The id of the transit router.
	TransitRouterId pulumi.StringOutput `pulumi:"transitRouterId"`
	// The update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The ID of vpc.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVpcAttachment registers a new resource with the given unique name, arguments, and options.
func NewVpcAttachment(ctx *pulumi.Context,
	name string, args *VpcAttachmentArgs, opts ...pulumi.ResourceOption) (*VpcAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttachPoints == nil {
		return nil, errors.New("invalid value for required argument 'AttachPoints'")
	}
	if args.TransitRouterId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcAttachment
	err := ctx.RegisterResource("volcengine:transit_router/vpcAttachment:VpcAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcAttachment gets an existing VpcAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcAttachmentState, opts ...pulumi.ResourceOption) (*VpcAttachment, error) {
	var resource VpcAttachment
	err := ctx.ReadResource("volcengine:transit_router/vpcAttachment:VpcAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcAttachment resources.
type vpcAttachmentState struct {
	// The attach points of transit router vpc attachment.
	AttachPoints []VpcAttachmentAttachPoint `pulumi:"attachPoints"`
	// Whether to auto publish route of the transit router to vpc instance. Default is false.
	AutoPublishRouteEnabled *bool `pulumi:"autoPublishRouteEnabled"`
	// The create time.
	CreationTime *string `pulumi:"creationTime"`
	// The description of the transit router vpc attachment.
	Description *string `pulumi:"description"`
	// The status of the transit router.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []VpcAttachmentTag `pulumi:"tags"`
	// The id of the transit router attachment.
	TransitRouterAttachmentId *string `pulumi:"transitRouterAttachmentId"`
	// The name of the transit router vpc attachment.
	TransitRouterAttachmentName *string `pulumi:"transitRouterAttachmentName"`
	// The id of the transit router.
	TransitRouterId *string `pulumi:"transitRouterId"`
	// The update time.
	UpdateTime *string `pulumi:"updateTime"`
	// The ID of vpc.
	VpcId *string `pulumi:"vpcId"`
}

type VpcAttachmentState struct {
	// The attach points of transit router vpc attachment.
	AttachPoints VpcAttachmentAttachPointArrayInput
	// Whether to auto publish route of the transit router to vpc instance. Default is false.
	AutoPublishRouteEnabled pulumi.BoolPtrInput
	// The create time.
	CreationTime pulumi.StringPtrInput
	// The description of the transit router vpc attachment.
	Description pulumi.StringPtrInput
	// The status of the transit router.
	Status pulumi.StringPtrInput
	// Tags.
	Tags VpcAttachmentTagArrayInput
	// The id of the transit router attachment.
	TransitRouterAttachmentId pulumi.StringPtrInput
	// The name of the transit router vpc attachment.
	TransitRouterAttachmentName pulumi.StringPtrInput
	// The id of the transit router.
	TransitRouterId pulumi.StringPtrInput
	// The update time.
	UpdateTime pulumi.StringPtrInput
	// The ID of vpc.
	VpcId pulumi.StringPtrInput
}

func (VpcAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAttachmentState)(nil)).Elem()
}

type vpcAttachmentArgs struct {
	// The attach points of transit router vpc attachment.
	AttachPoints []VpcAttachmentAttachPoint `pulumi:"attachPoints"`
	// Whether to auto publish route of the transit router to vpc instance. Default is false.
	AutoPublishRouteEnabled *bool `pulumi:"autoPublishRouteEnabled"`
	// The description of the transit router vpc attachment.
	Description *string `pulumi:"description"`
	// Tags.
	Tags []VpcAttachmentTag `pulumi:"tags"`
	// The name of the transit router vpc attachment.
	TransitRouterAttachmentName *string `pulumi:"transitRouterAttachmentName"`
	// The id of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
	// The ID of vpc.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VpcAttachment resource.
type VpcAttachmentArgs struct {
	// The attach points of transit router vpc attachment.
	AttachPoints VpcAttachmentAttachPointArrayInput
	// Whether to auto publish route of the transit router to vpc instance. Default is false.
	AutoPublishRouteEnabled pulumi.BoolPtrInput
	// The description of the transit router vpc attachment.
	Description pulumi.StringPtrInput
	// Tags.
	Tags VpcAttachmentTagArrayInput
	// The name of the transit router vpc attachment.
	TransitRouterAttachmentName pulumi.StringPtrInput
	// The id of the transit router.
	TransitRouterId pulumi.StringInput
	// The ID of vpc.
	VpcId pulumi.StringInput
}

func (VpcAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAttachmentArgs)(nil)).Elem()
}

type VpcAttachmentInput interface {
	pulumi.Input

	ToVpcAttachmentOutput() VpcAttachmentOutput
	ToVpcAttachmentOutputWithContext(ctx context.Context) VpcAttachmentOutput
}

func (*VpcAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcAttachment)(nil)).Elem()
}

func (i *VpcAttachment) ToVpcAttachmentOutput() VpcAttachmentOutput {
	return i.ToVpcAttachmentOutputWithContext(context.Background())
}

func (i *VpcAttachment) ToVpcAttachmentOutputWithContext(ctx context.Context) VpcAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAttachmentOutput)
}

// VpcAttachmentArrayInput is an input type that accepts VpcAttachmentArray and VpcAttachmentArrayOutput values.
// You can construct a concrete instance of `VpcAttachmentArrayInput` via:
//
//	VpcAttachmentArray{ VpcAttachmentArgs{...} }
type VpcAttachmentArrayInput interface {
	pulumi.Input

	ToVpcAttachmentArrayOutput() VpcAttachmentArrayOutput
	ToVpcAttachmentArrayOutputWithContext(context.Context) VpcAttachmentArrayOutput
}

type VpcAttachmentArray []VpcAttachmentInput

func (VpcAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcAttachment)(nil)).Elem()
}

func (i VpcAttachmentArray) ToVpcAttachmentArrayOutput() VpcAttachmentArrayOutput {
	return i.ToVpcAttachmentArrayOutputWithContext(context.Background())
}

func (i VpcAttachmentArray) ToVpcAttachmentArrayOutputWithContext(ctx context.Context) VpcAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAttachmentArrayOutput)
}

// VpcAttachmentMapInput is an input type that accepts VpcAttachmentMap and VpcAttachmentMapOutput values.
// You can construct a concrete instance of `VpcAttachmentMapInput` via:
//
//	VpcAttachmentMap{ "key": VpcAttachmentArgs{...} }
type VpcAttachmentMapInput interface {
	pulumi.Input

	ToVpcAttachmentMapOutput() VpcAttachmentMapOutput
	ToVpcAttachmentMapOutputWithContext(context.Context) VpcAttachmentMapOutput
}

type VpcAttachmentMap map[string]VpcAttachmentInput

func (VpcAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcAttachment)(nil)).Elem()
}

func (i VpcAttachmentMap) ToVpcAttachmentMapOutput() VpcAttachmentMapOutput {
	return i.ToVpcAttachmentMapOutputWithContext(context.Background())
}

func (i VpcAttachmentMap) ToVpcAttachmentMapOutputWithContext(ctx context.Context) VpcAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcAttachmentMapOutput)
}

type VpcAttachmentOutput struct{ *pulumi.OutputState }

func (VpcAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcAttachment)(nil)).Elem()
}

func (o VpcAttachmentOutput) ToVpcAttachmentOutput() VpcAttachmentOutput {
	return o
}

func (o VpcAttachmentOutput) ToVpcAttachmentOutputWithContext(ctx context.Context) VpcAttachmentOutput {
	return o
}

// The attach points of transit router vpc attachment.
func (o VpcAttachmentOutput) AttachPoints() VpcAttachmentAttachPointArrayOutput {
	return o.ApplyT(func(v *VpcAttachment) VpcAttachmentAttachPointArrayOutput { return v.AttachPoints }).(VpcAttachmentAttachPointArrayOutput)
}

// Whether to auto publish route of the transit router to vpc instance. Default is false.
func (o VpcAttachmentOutput) AutoPublishRouteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.BoolPtrOutput { return v.AutoPublishRouteEnabled }).(pulumi.BoolPtrOutput)
}

// The create time.
func (o VpcAttachmentOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The description of the transit router vpc attachment.
func (o VpcAttachmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The status of the transit router.
func (o VpcAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags.
func (o VpcAttachmentOutput) Tags() VpcAttachmentTagArrayOutput {
	return o.ApplyT(func(v *VpcAttachment) VpcAttachmentTagArrayOutput { return v.Tags }).(VpcAttachmentTagArrayOutput)
}

// The id of the transit router attachment.
func (o VpcAttachmentOutput) TransitRouterAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.StringOutput { return v.TransitRouterAttachmentId }).(pulumi.StringOutput)
}

// The name of the transit router vpc attachment.
func (o VpcAttachmentOutput) TransitRouterAttachmentName() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.StringOutput { return v.TransitRouterAttachmentName }).(pulumi.StringOutput)
}

// The id of the transit router.
func (o VpcAttachmentOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.StringOutput { return v.TransitRouterId }).(pulumi.StringOutput)
}

// The update time.
func (o VpcAttachmentOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The ID of vpc.
func (o VpcAttachmentOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcAttachment) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type VpcAttachmentArrayOutput struct{ *pulumi.OutputState }

func (VpcAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcAttachment)(nil)).Elem()
}

func (o VpcAttachmentArrayOutput) ToVpcAttachmentArrayOutput() VpcAttachmentArrayOutput {
	return o
}

func (o VpcAttachmentArrayOutput) ToVpcAttachmentArrayOutputWithContext(ctx context.Context) VpcAttachmentArrayOutput {
	return o
}

func (o VpcAttachmentArrayOutput) Index(i pulumi.IntInput) VpcAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcAttachment {
		return vs[0].([]*VpcAttachment)[vs[1].(int)]
	}).(VpcAttachmentOutput)
}

type VpcAttachmentMapOutput struct{ *pulumi.OutputState }

func (VpcAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcAttachment)(nil)).Elem()
}

func (o VpcAttachmentMapOutput) ToVpcAttachmentMapOutput() VpcAttachmentMapOutput {
	return o
}

func (o VpcAttachmentMapOutput) ToVpcAttachmentMapOutputWithContext(ctx context.Context) VpcAttachmentMapOutput {
	return o
}

func (o VpcAttachmentMapOutput) MapIndex(k pulumi.StringInput) VpcAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcAttachment {
		return vs[0].(map[string]*VpcAttachment)[vs[1].(string)]
	}).(VpcAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAttachmentInput)(nil)).Elem(), &VpcAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAttachmentArrayInput)(nil)).Elem(), VpcAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcAttachmentMapInput)(nil)).Elem(), VpcAttachmentMap{})
	pulumi.RegisterOutputType(VpcAttachmentOutput{})
	pulumi.RegisterOutputType(VpcAttachmentArrayOutput{})
	pulumi.RegisterOutputType(VpcAttachmentMapOutput{})
}
