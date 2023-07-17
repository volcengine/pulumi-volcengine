// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage privatelink vpc endpoint service resource
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
//			_, err := privatelink.NewVpcEndpointServiceResource(ctx, "foo", &privatelink.VpcEndpointServiceResourceArgs{
//				ResourceId: pulumi.String("clb-3reii8qfbp7gg5zsk2hsrbe3c"),
//				ServiceId:  pulumi.String("epsvc-3rel73uf2ewao5zsk2j2l58ro"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = privatelink.NewVpcEndpointServiceResource(ctx, "foo1", &privatelink.VpcEndpointServiceResourceArgs{
//				ResourceId: pulumi.String("clb-2d6sfye98rzls58ozfducee1o"),
//				ServiceId:  pulumi.String("epsvc-3rel73uf2ewao5zsk2j2l58ro"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = privatelink.NewVpcEndpointServiceResource(ctx, "foo2", &privatelink.VpcEndpointServiceResourceArgs{
//				ResourceId: pulumi.String("clb-3refkvae02gow5zsk2ilaev5y"),
//				ServiceId:  pulumi.String("epsvc-3rel73uf2ewao5zsk2j2l58ro"),
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
// VpcEndpointServiceResource can be imported using the serviceId:resourceId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource default epsvc-2fe630gurkl37k5gfuy33****:clb-bp1o94dp5i6ea****
//
// ```
type VpcEndpointServiceResource struct {
	pulumi.CustomResourceState

	// The id of resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The id of service.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
}

// NewVpcEndpointServiceResource registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointServiceResource(ctx *pulumi.Context,
	name string, args *VpcEndpointServiceResourceArgs, opts ...pulumi.ResourceOption) (*VpcEndpointServiceResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	var resource VpcEndpointServiceResource
	err := ctx.RegisterResource("volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointServiceResource gets an existing VpcEndpointServiceResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointServiceResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointServiceResourceState, opts ...pulumi.ResourceOption) (*VpcEndpointServiceResource, error) {
	var resource VpcEndpointServiceResource
	err := ctx.ReadResource("volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointServiceResource resources.
type vpcEndpointServiceResourceState struct {
	// The id of resource.
	ResourceId *string `pulumi:"resourceId"`
	// The id of service.
	ServiceId *string `pulumi:"serviceId"`
}

type VpcEndpointServiceResourceState struct {
	// The id of resource.
	ResourceId pulumi.StringPtrInput
	// The id of service.
	ServiceId pulumi.StringPtrInput
}

func (VpcEndpointServiceResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceResourceState)(nil)).Elem()
}

type vpcEndpointServiceResourceArgs struct {
	// The id of resource.
	ResourceId string `pulumi:"resourceId"`
	// The id of service.
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a VpcEndpointServiceResource resource.
type VpcEndpointServiceResourceArgs struct {
	// The id of resource.
	ResourceId pulumi.StringInput
	// The id of service.
	ServiceId pulumi.StringInput
}

func (VpcEndpointServiceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceResourceArgs)(nil)).Elem()
}

type VpcEndpointServiceResourceInput interface {
	pulumi.Input

	ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput
	ToVpcEndpointServiceResourceOutputWithContext(ctx context.Context) VpcEndpointServiceResourceOutput
}

func (*VpcEndpointServiceResource) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceResource)(nil)).Elem()
}

func (i *VpcEndpointServiceResource) ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput {
	return i.ToVpcEndpointServiceResourceOutputWithContext(context.Background())
}

func (i *VpcEndpointServiceResource) ToVpcEndpointServiceResourceOutputWithContext(ctx context.Context) VpcEndpointServiceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceResourceOutput)
}

// VpcEndpointServiceResourceArrayInput is an input type that accepts VpcEndpointServiceResourceArray and VpcEndpointServiceResourceArrayOutput values.
// You can construct a concrete instance of `VpcEndpointServiceResourceArrayInput` via:
//
//	VpcEndpointServiceResourceArray{ VpcEndpointServiceResourceArgs{...} }
type VpcEndpointServiceResourceArrayInput interface {
	pulumi.Input

	ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput
	ToVpcEndpointServiceResourceArrayOutputWithContext(context.Context) VpcEndpointServiceResourceArrayOutput
}

type VpcEndpointServiceResourceArray []VpcEndpointServiceResourceInput

func (VpcEndpointServiceResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServiceResource)(nil)).Elem()
}

func (i VpcEndpointServiceResourceArray) ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput {
	return i.ToVpcEndpointServiceResourceArrayOutputWithContext(context.Background())
}

func (i VpcEndpointServiceResourceArray) ToVpcEndpointServiceResourceArrayOutputWithContext(ctx context.Context) VpcEndpointServiceResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceResourceArrayOutput)
}

// VpcEndpointServiceResourceMapInput is an input type that accepts VpcEndpointServiceResourceMap and VpcEndpointServiceResourceMapOutput values.
// You can construct a concrete instance of `VpcEndpointServiceResourceMapInput` via:
//
//	VpcEndpointServiceResourceMap{ "key": VpcEndpointServiceResourceArgs{...} }
type VpcEndpointServiceResourceMapInput interface {
	pulumi.Input

	ToVpcEndpointServiceResourceMapOutput() VpcEndpointServiceResourceMapOutput
	ToVpcEndpointServiceResourceMapOutputWithContext(context.Context) VpcEndpointServiceResourceMapOutput
}

type VpcEndpointServiceResourceMap map[string]VpcEndpointServiceResourceInput

func (VpcEndpointServiceResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServiceResource)(nil)).Elem()
}

func (i VpcEndpointServiceResourceMap) ToVpcEndpointServiceResourceMapOutput() VpcEndpointServiceResourceMapOutput {
	return i.ToVpcEndpointServiceResourceMapOutputWithContext(context.Background())
}

func (i VpcEndpointServiceResourceMap) ToVpcEndpointServiceResourceMapOutputWithContext(ctx context.Context) VpcEndpointServiceResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServiceResourceMapOutput)
}

type VpcEndpointServiceResourceOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServiceResource)(nil)).Elem()
}

func (o VpcEndpointServiceResourceOutput) ToVpcEndpointServiceResourceOutput() VpcEndpointServiceResourceOutput {
	return o
}

func (o VpcEndpointServiceResourceOutput) ToVpcEndpointServiceResourceOutputWithContext(ctx context.Context) VpcEndpointServiceResourceOutput {
	return o
}

// The id of resource.
func (o VpcEndpointServiceResourceOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceResource) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The id of service.
func (o VpcEndpointServiceResourceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServiceResource) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

type VpcEndpointServiceResourceArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServiceResource)(nil)).Elem()
}

func (o VpcEndpointServiceResourceArrayOutput) ToVpcEndpointServiceResourceArrayOutput() VpcEndpointServiceResourceArrayOutput {
	return o
}

func (o VpcEndpointServiceResourceArrayOutput) ToVpcEndpointServiceResourceArrayOutputWithContext(ctx context.Context) VpcEndpointServiceResourceArrayOutput {
	return o
}

func (o VpcEndpointServiceResourceArrayOutput) Index(i pulumi.IntInput) VpcEndpointServiceResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointServiceResource {
		return vs[0].([]*VpcEndpointServiceResource)[vs[1].(int)]
	}).(VpcEndpointServiceResourceOutput)
}

type VpcEndpointServiceResourceMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointServiceResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServiceResource)(nil)).Elem()
}

func (o VpcEndpointServiceResourceMapOutput) ToVpcEndpointServiceResourceMapOutput() VpcEndpointServiceResourceMapOutput {
	return o
}

func (o VpcEndpointServiceResourceMapOutput) ToVpcEndpointServiceResourceMapOutputWithContext(ctx context.Context) VpcEndpointServiceResourceMapOutput {
	return o
}

func (o VpcEndpointServiceResourceMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointServiceResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointServiceResource {
		return vs[0].(map[string]*VpcEndpointServiceResource)[vs[1].(string)]
	}).(VpcEndpointServiceResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceResourceInput)(nil)).Elem(), &VpcEndpointServiceResource{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceResourceArrayInput)(nil)).Elem(), VpcEndpointServiceResourceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServiceResourceMapInput)(nil)).Elem(), VpcEndpointServiceResourceMap{})
	pulumi.RegisterOutputType(VpcEndpointServiceResourceOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceResourceArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointServiceResourceMapOutput{})
}