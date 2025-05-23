// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage privatelink vpc endpoint connection
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/privatelink"
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
//				SecurityGroupName: pulumi.String("acc-test-security-group"),
//				VpcId:             fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooClb, err := clb.NewClb(ctx, "fooClb", &clb.ClbArgs{
//				Type:                    pulumi.String("public"),
//				SubnetId:                fooSubnet.ID(),
//				LoadBalancerSpec:        pulumi.String("small_1"),
//				Description:             pulumi.String("acc-test-demo"),
//				LoadBalancerName:        pulumi.String("acc-test-clb"),
//				LoadBalancerBillingType: pulumi.String("PostPaid"),
//				EipBillingConfig: &clb.ClbEipBillingConfigArgs{
//					Isp:            pulumi.String("BGP"),
//					EipBillingType: pulumi.String("PostPaidByBandwidth"),
//					Bandwidth:      pulumi.Int(1),
//				},
//				Tags: clb.ClbTagArray{
//					&clb.ClbTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooVpcEndpointService, err := privatelink.NewVpcEndpointService(ctx, "fooVpcEndpointService", &privatelink.VpcEndpointServiceArgs{
//				Resources: privatelink.VpcEndpointServiceResourceTypeArray{
//					&privatelink.VpcEndpointServiceResourceTypeArgs{
//						ResourceId:   fooClb.ID(),
//						ResourceType: pulumi.String("CLB"),
//					},
//				},
//				Description: pulumi.String("acc-test"),
//			})
//			if err != nil {
//				return err
//			}
//			fooVpcEndpoint, err := privatelink.NewVpcEndpoint(ctx, "fooVpcEndpoint", &privatelink.VpcEndpointArgs{
//				SecurityGroupIds: pulumi.StringArray{
//					fooSecurityGroup.ID(),
//				},
//				ServiceId:    fooVpcEndpointService.ID(),
//				EndpointName: pulumi.String("acc-test-ep"),
//				Description:  pulumi.String("acc-test"),
//			})
//			if err != nil {
//				return err
//			}
//			fooVpcEndpointZone, err := privatelink.NewVpcEndpointZone(ctx, "fooVpcEndpointZone", &privatelink.VpcEndpointZoneArgs{
//				EndpointId:       fooVpcEndpoint.ID(),
//				SubnetId:         fooSubnet.ID(),
//				PrivateIpAddress: pulumi.String("172.16.0.251"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = privatelink.NewVpcEndpointConnection(ctx, "fooVpcEndpointConnection", &privatelink.VpcEndpointConnectionArgs{
//				EndpointId: fooVpcEndpoint.ID(),
//				ServiceId:  fooVpcEndpointService.ID(),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				fooVpcEndpointZone,
//			}))
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
// PrivateLink Vpc Endpoint Connection Service can be imported using the endpoint id and service id, e.g.
//
// ```sh
// $ pulumi import volcengine:privatelink/vpcEndpointConnection:VpcEndpointConnection default ep-3rel74u229dz45zsk2i6l69qa:epsvc-2byz5mykk9y4g2dx0efs4aqz3
// ```
type VpcEndpointConnection struct {
	pulumi.CustomResourceState

	// The status of the connection.
	ConnectionStatus pulumi.StringOutput `pulumi:"connectionStatus"`
	// The create time of the connection.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The id of the endpoint.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// The account id of the vpc endpoint.
	EndpointOwnerAccountId pulumi.StringOutput `pulumi:"endpointOwnerAccountId"`
	// The vpc id of the vpc endpoint.
	EndpointVpcId pulumi.StringOutput `pulumi:"endpointVpcId"`
	// The id of the security group.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The update time of the connection.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The available zones.
	Zones VpcEndpointConnectionZoneArrayOutput `pulumi:"zones"`
}

// NewVpcEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointConnection(ctx *pulumi.Context,
	name string, args *VpcEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*VpcEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpcEndpointConnection
	err := ctx.RegisterResource("volcengine:privatelink/vpcEndpointConnection:VpcEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointConnection gets an existing VpcEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointConnectionState, opts ...pulumi.ResourceOption) (*VpcEndpointConnection, error) {
	var resource VpcEndpointConnection
	err := ctx.ReadResource("volcengine:privatelink/vpcEndpointConnection:VpcEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointConnection resources.
type vpcEndpointConnectionState struct {
	// The status of the connection.
	ConnectionStatus *string `pulumi:"connectionStatus"`
	// The create time of the connection.
	CreationTime *string `pulumi:"creationTime"`
	// The id of the endpoint.
	EndpointId *string `pulumi:"endpointId"`
	// The account id of the vpc endpoint.
	EndpointOwnerAccountId *string `pulumi:"endpointOwnerAccountId"`
	// The vpc id of the vpc endpoint.
	EndpointVpcId *string `pulumi:"endpointVpcId"`
	// The id of the security group.
	ServiceId *string `pulumi:"serviceId"`
	// The update time of the connection.
	UpdateTime *string `pulumi:"updateTime"`
	// The available zones.
	Zones []VpcEndpointConnectionZone `pulumi:"zones"`
}

type VpcEndpointConnectionState struct {
	// The status of the connection.
	ConnectionStatus pulumi.StringPtrInput
	// The create time of the connection.
	CreationTime pulumi.StringPtrInput
	// The id of the endpoint.
	EndpointId pulumi.StringPtrInput
	// The account id of the vpc endpoint.
	EndpointOwnerAccountId pulumi.StringPtrInput
	// The vpc id of the vpc endpoint.
	EndpointVpcId pulumi.StringPtrInput
	// The id of the security group.
	ServiceId pulumi.StringPtrInput
	// The update time of the connection.
	UpdateTime pulumi.StringPtrInput
	// The available zones.
	Zones VpcEndpointConnectionZoneArrayInput
}

func (VpcEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointConnectionState)(nil)).Elem()
}

type vpcEndpointConnectionArgs struct {
	// The id of the endpoint.
	EndpointId string `pulumi:"endpointId"`
	// The id of the security group.
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a VpcEndpointConnection resource.
type VpcEndpointConnectionArgs struct {
	// The id of the endpoint.
	EndpointId pulumi.StringInput
	// The id of the security group.
	ServiceId pulumi.StringInput
}

func (VpcEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointConnectionArgs)(nil)).Elem()
}

type VpcEndpointConnectionInput interface {
	pulumi.Input

	ToVpcEndpointConnectionOutput() VpcEndpointConnectionOutput
	ToVpcEndpointConnectionOutputWithContext(ctx context.Context) VpcEndpointConnectionOutput
}

func (*VpcEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointConnection)(nil)).Elem()
}

func (i *VpcEndpointConnection) ToVpcEndpointConnectionOutput() VpcEndpointConnectionOutput {
	return i.ToVpcEndpointConnectionOutputWithContext(context.Background())
}

func (i *VpcEndpointConnection) ToVpcEndpointConnectionOutputWithContext(ctx context.Context) VpcEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionOutput)
}

// VpcEndpointConnectionArrayInput is an input type that accepts VpcEndpointConnectionArray and VpcEndpointConnectionArrayOutput values.
// You can construct a concrete instance of `VpcEndpointConnectionArrayInput` via:
//
//	VpcEndpointConnectionArray{ VpcEndpointConnectionArgs{...} }
type VpcEndpointConnectionArrayInput interface {
	pulumi.Input

	ToVpcEndpointConnectionArrayOutput() VpcEndpointConnectionArrayOutput
	ToVpcEndpointConnectionArrayOutputWithContext(context.Context) VpcEndpointConnectionArrayOutput
}

type VpcEndpointConnectionArray []VpcEndpointConnectionInput

func (VpcEndpointConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointConnection)(nil)).Elem()
}

func (i VpcEndpointConnectionArray) ToVpcEndpointConnectionArrayOutput() VpcEndpointConnectionArrayOutput {
	return i.ToVpcEndpointConnectionArrayOutputWithContext(context.Background())
}

func (i VpcEndpointConnectionArray) ToVpcEndpointConnectionArrayOutputWithContext(ctx context.Context) VpcEndpointConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionArrayOutput)
}

// VpcEndpointConnectionMapInput is an input type that accepts VpcEndpointConnectionMap and VpcEndpointConnectionMapOutput values.
// You can construct a concrete instance of `VpcEndpointConnectionMapInput` via:
//
//	VpcEndpointConnectionMap{ "key": VpcEndpointConnectionArgs{...} }
type VpcEndpointConnectionMapInput interface {
	pulumi.Input

	ToVpcEndpointConnectionMapOutput() VpcEndpointConnectionMapOutput
	ToVpcEndpointConnectionMapOutputWithContext(context.Context) VpcEndpointConnectionMapOutput
}

type VpcEndpointConnectionMap map[string]VpcEndpointConnectionInput

func (VpcEndpointConnectionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointConnection)(nil)).Elem()
}

func (i VpcEndpointConnectionMap) ToVpcEndpointConnectionMapOutput() VpcEndpointConnectionMapOutput {
	return i.ToVpcEndpointConnectionMapOutputWithContext(context.Background())
}

func (i VpcEndpointConnectionMap) ToVpcEndpointConnectionMapOutputWithContext(ctx context.Context) VpcEndpointConnectionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointConnectionMapOutput)
}

type VpcEndpointConnectionOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointConnection)(nil)).Elem()
}

func (o VpcEndpointConnectionOutput) ToVpcEndpointConnectionOutput() VpcEndpointConnectionOutput {
	return o
}

func (o VpcEndpointConnectionOutput) ToVpcEndpointConnectionOutputWithContext(ctx context.Context) VpcEndpointConnectionOutput {
	return o
}

// The status of the connection.
func (o VpcEndpointConnectionOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnection) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// The create time of the connection.
func (o VpcEndpointConnectionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnection) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The id of the endpoint.
func (o VpcEndpointConnectionOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnection) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// The account id of the vpc endpoint.
func (o VpcEndpointConnectionOutput) EndpointOwnerAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnection) pulumi.StringOutput { return v.EndpointOwnerAccountId }).(pulumi.StringOutput)
}

// The vpc id of the vpc endpoint.
func (o VpcEndpointConnectionOutput) EndpointVpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnection) pulumi.StringOutput { return v.EndpointVpcId }).(pulumi.StringOutput)
}

// The id of the security group.
func (o VpcEndpointConnectionOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnection) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The update time of the connection.
func (o VpcEndpointConnectionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointConnection) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The available zones.
func (o VpcEndpointConnectionOutput) Zones() VpcEndpointConnectionZoneArrayOutput {
	return o.ApplyT(func(v *VpcEndpointConnection) VpcEndpointConnectionZoneArrayOutput { return v.Zones }).(VpcEndpointConnectionZoneArrayOutput)
}

type VpcEndpointConnectionArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointConnection)(nil)).Elem()
}

func (o VpcEndpointConnectionArrayOutput) ToVpcEndpointConnectionArrayOutput() VpcEndpointConnectionArrayOutput {
	return o
}

func (o VpcEndpointConnectionArrayOutput) ToVpcEndpointConnectionArrayOutputWithContext(ctx context.Context) VpcEndpointConnectionArrayOutput {
	return o
}

func (o VpcEndpointConnectionArrayOutput) Index(i pulumi.IntInput) VpcEndpointConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointConnection {
		return vs[0].([]*VpcEndpointConnection)[vs[1].(int)]
	}).(VpcEndpointConnectionOutput)
}

type VpcEndpointConnectionMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointConnectionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointConnection)(nil)).Elem()
}

func (o VpcEndpointConnectionMapOutput) ToVpcEndpointConnectionMapOutput() VpcEndpointConnectionMapOutput {
	return o
}

func (o VpcEndpointConnectionMapOutput) ToVpcEndpointConnectionMapOutputWithContext(ctx context.Context) VpcEndpointConnectionMapOutput {
	return o
}

func (o VpcEndpointConnectionMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointConnectionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointConnection {
		return vs[0].(map[string]*VpcEndpointConnection)[vs[1].(string)]
	}).(VpcEndpointConnectionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionInput)(nil)).Elem(), &VpcEndpointConnection{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionArrayInput)(nil)).Elem(), VpcEndpointConnectionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointConnectionMapInput)(nil)).Elem(), VpcEndpointConnectionMap{})
	pulumi.RegisterOutputType(VpcEndpointConnectionOutput{})
	pulumi.RegisterOutputType(VpcEndpointConnectionArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointConnectionMapOutput{})
}
