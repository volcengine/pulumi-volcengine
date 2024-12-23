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

// Provides a resource to manage transit router route entry
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/transit_router"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooCustomerGateway, err := vpn.NewCustomerGateway(ctx, "fooCustomerGateway", &vpn.CustomerGatewayArgs{
//				IpAddress:           pulumi.String("192.0.1.3"),
//				CustomerGatewayName: pulumi.String("acc-test"),
//				Description:         pulumi.String("acc-test"),
//			})
//			if err != nil {
//				return err
//			}
//			fooConnection, err := vpn.NewConnection(ctx, "fooConnection", &vpn.ConnectionArgs{
//				VpnConnectionName: pulumi.String("acc-tf-test"),
//				Description:       pulumi.String("acc-tf-test"),
//				AttachType:        pulumi.String("TransitRouter"),
//				CustomerGatewayId: fooCustomerGateway.ID(),
//				LocalSubnets: pulumi.StringArray{
//					pulumi.String("192.168.0.0/22"),
//				},
//				RemoteSubnets: pulumi.StringArray{
//					pulumi.String("192.161.0.0/20"),
//				},
//				DpdAction:           pulumi.String("none"),
//				NatTraversal:        pulumi.Bool(true),
//				IkeConfigPsk:        pulumi.String("acctest@!3"),
//				IkeConfigVersion:    pulumi.String("ikev1"),
//				IkeConfigMode:       pulumi.String("main"),
//				IkeConfigEncAlg:     pulumi.String("aes"),
//				IkeConfigAuthAlg:    pulumi.String("md5"),
//				IkeConfigDhGroup:    pulumi.String("group2"),
//				IkeConfigLifetime:   pulumi.Int(9000),
//				IkeConfigLocalId:    pulumi.String("acc_test"),
//				IkeConfigRemoteId:   pulumi.String("acc_test"),
//				IpsecConfigEncAlg:   pulumi.String("aes"),
//				IpsecConfigAuthAlg:  pulumi.String("sha256"),
//				IpsecConfigDhGroup:  pulumi.String("group2"),
//				IpsecConfigLifetime: pulumi.Int(9000),
//				LogEnabled:          pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			fooTransitRouter, err := transit_router.NewTransitRouter(ctx, "fooTransitRouter", &transit_router.TransitRouterArgs{
//				TransitRouterName: pulumi.String("test-tf-acc"),
//				Description:       pulumi.String("test-tf-acc"),
//			})
//			if err != nil {
//				return err
//			}
//			fooVpnAttachment, err := transit_router.NewVpnAttachment(ctx, "fooVpnAttachment", &transit_router.VpnAttachmentArgs{
//				ZoneId:                      pulumi.String("cn-beijing-a"),
//				TransitRouterAttachmentName: pulumi.String("tf-test-acc"),
//				Description:                 pulumi.String("tf-test-acc-desc"),
//				TransitRouterId:             fooTransitRouter.ID(),
//				VpnConnectionId:             fooConnection.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooRouteTable, err := transit_router.NewRouteTable(ctx, "fooRouteTable", &transit_router.RouteTableArgs{
//				Description:                 pulumi.String("tf-test-acc-description-route-route-table"),
//				TransitRouterRouteTableName: pulumi.String("tf-table-test-acc"),
//				TransitRouterId:             fooTransitRouter.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = transit_router.NewRouteEntry(ctx, "fooRouteEntry", &transit_router.RouteEntryArgs{
//				Description:                        pulumi.String("tf-test-acc-description-entry"),
//				TransitRouterRouteEntryName:        pulumi.String("tf-acc-test-entry"),
//				DestinationCidrBlock:               pulumi.String("192.168.0.0/24"),
//				TransitRouterRouteEntryNextHopType: pulumi.String("Attachment"),
//				TransitRouterRouteTableId:          fooRouteTable.TransitRouterRouteTableId,
//				TransitRouterRouteEntryNextHopId:   fooVpnAttachment.TransitRouterAttachmentId,
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
// transit router route entry can be imported using the table and entry id, e.g.
//
// ```sh
// $ pulumi import volcengine:transit_router/routeEntry:RouteEntry default tr-rtb-12b7qd3fmzf2817q7y2jkbd55:tr-rte-1i5i8khf9m58gae5kcx6***
// ```
type RouteEntry struct {
	pulumi.CustomResourceState

	// The as path of the route entry.
	AsPath pulumi.StringOutput `pulumi:"asPath"`
	// The creation time of the route entry.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Description of the transit router route entry.
	Description pulumi.StringOutput `pulumi:"description"`
	// The target network segment of the route entry.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// The status of the route entry.
	Status pulumi.StringOutput `pulumi:"status"`
	// The id of the route entry.
	TransitRouterRouteEntryId pulumi.StringOutput `pulumi:"transitRouterRouteEntryId"`
	// The name of the route entry.
	TransitRouterRouteEntryName pulumi.StringOutput `pulumi:"transitRouterRouteEntryName"`
	// The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
	TransitRouterRouteEntryNextHopId pulumi.StringPtrOutput `pulumi:"transitRouterRouteEntryNextHopId"`
	// The next hop type of the routing entry. The value can be Attachment or BlackHole.
	TransitRouterRouteEntryNextHopType pulumi.StringOutput `pulumi:"transitRouterRouteEntryNextHopType"`
	// The type of the route entry.
	TransitRouterRouteEntryType pulumi.StringOutput `pulumi:"transitRouterRouteEntryType"`
	// The id of the route table.
	TransitRouterRouteTableId pulumi.StringOutput `pulumi:"transitRouterRouteTableId"`
	// The update time of the route entry.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRouteEntry registers a new resource with the given unique name, arguments, and options.
func NewRouteEntry(ctx *pulumi.Context,
	name string, args *RouteEntryArgs, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationCidrBlock == nil {
		return nil, errors.New("invalid value for required argument 'DestinationCidrBlock'")
	}
	if args.TransitRouterRouteEntryNextHopType == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterRouteEntryNextHopType'")
	}
	if args.TransitRouterRouteTableId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterRouteTableId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouteEntry
	err := ctx.RegisterResource("volcengine:transit_router/routeEntry:RouteEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteEntry gets an existing RouteEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteEntryState, opts ...pulumi.ResourceOption) (*RouteEntry, error) {
	var resource RouteEntry
	err := ctx.ReadResource("volcengine:transit_router/routeEntry:RouteEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteEntry resources.
type routeEntryState struct {
	// The as path of the route entry.
	AsPath *string `pulumi:"asPath"`
	// The creation time of the route entry.
	CreationTime *string `pulumi:"creationTime"`
	// Description of the transit router route entry.
	Description *string `pulumi:"description"`
	// The target network segment of the route entry.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The status of the route entry.
	Status *string `pulumi:"status"`
	// The id of the route entry.
	TransitRouterRouteEntryId *string `pulumi:"transitRouterRouteEntryId"`
	// The name of the route entry.
	TransitRouterRouteEntryName *string `pulumi:"transitRouterRouteEntryName"`
	// The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
	TransitRouterRouteEntryNextHopId *string `pulumi:"transitRouterRouteEntryNextHopId"`
	// The next hop type of the routing entry. The value can be Attachment or BlackHole.
	TransitRouterRouteEntryNextHopType *string `pulumi:"transitRouterRouteEntryNextHopType"`
	// The type of the route entry.
	TransitRouterRouteEntryType *string `pulumi:"transitRouterRouteEntryType"`
	// The id of the route table.
	TransitRouterRouteTableId *string `pulumi:"transitRouterRouteTableId"`
	// The update time of the route entry.
	UpdateTime *string `pulumi:"updateTime"`
}

type RouteEntryState struct {
	// The as path of the route entry.
	AsPath pulumi.StringPtrInput
	// The creation time of the route entry.
	CreationTime pulumi.StringPtrInput
	// Description of the transit router route entry.
	Description pulumi.StringPtrInput
	// The target network segment of the route entry.
	DestinationCidrBlock pulumi.StringPtrInput
	// The status of the route entry.
	Status pulumi.StringPtrInput
	// The id of the route entry.
	TransitRouterRouteEntryId pulumi.StringPtrInput
	// The name of the route entry.
	TransitRouterRouteEntryName pulumi.StringPtrInput
	// The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
	TransitRouterRouteEntryNextHopId pulumi.StringPtrInput
	// The next hop type of the routing entry. The value can be Attachment or BlackHole.
	TransitRouterRouteEntryNextHopType pulumi.StringPtrInput
	// The type of the route entry.
	TransitRouterRouteEntryType pulumi.StringPtrInput
	// The id of the route table.
	TransitRouterRouteTableId pulumi.StringPtrInput
	// The update time of the route entry.
	UpdateTime pulumi.StringPtrInput
}

func (RouteEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryState)(nil)).Elem()
}

type routeEntryArgs struct {
	// Description of the transit router route entry.
	Description *string `pulumi:"description"`
	// The target network segment of the route entry.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// The name of the route entry.
	TransitRouterRouteEntryName *string `pulumi:"transitRouterRouteEntryName"`
	// The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
	TransitRouterRouteEntryNextHopId *string `pulumi:"transitRouterRouteEntryNextHopId"`
	// The next hop type of the routing entry. The value can be Attachment or BlackHole.
	TransitRouterRouteEntryNextHopType string `pulumi:"transitRouterRouteEntryNextHopType"`
	// The id of the route table.
	TransitRouterRouteTableId string `pulumi:"transitRouterRouteTableId"`
}

// The set of arguments for constructing a RouteEntry resource.
type RouteEntryArgs struct {
	// Description of the transit router route entry.
	Description pulumi.StringPtrInput
	// The target network segment of the route entry.
	DestinationCidrBlock pulumi.StringInput
	// The name of the route entry.
	TransitRouterRouteEntryName pulumi.StringPtrInput
	// The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
	TransitRouterRouteEntryNextHopId pulumi.StringPtrInput
	// The next hop type of the routing entry. The value can be Attachment or BlackHole.
	TransitRouterRouteEntryNextHopType pulumi.StringInput
	// The id of the route table.
	TransitRouterRouteTableId pulumi.StringInput
}

func (RouteEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeEntryArgs)(nil)).Elem()
}

type RouteEntryInput interface {
	pulumi.Input

	ToRouteEntryOutput() RouteEntryOutput
	ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput
}

func (*RouteEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (i *RouteEntry) ToRouteEntryOutput() RouteEntryOutput {
	return i.ToRouteEntryOutputWithContext(context.Background())
}

func (i *RouteEntry) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryOutput)
}

// RouteEntryArrayInput is an input type that accepts RouteEntryArray and RouteEntryArrayOutput values.
// You can construct a concrete instance of `RouteEntryArrayInput` via:
//
//	RouteEntryArray{ RouteEntryArgs{...} }
type RouteEntryArrayInput interface {
	pulumi.Input

	ToRouteEntryArrayOutput() RouteEntryArrayOutput
	ToRouteEntryArrayOutputWithContext(context.Context) RouteEntryArrayOutput
}

type RouteEntryArray []RouteEntryInput

func (RouteEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryArray) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return i.ToRouteEntryArrayOutputWithContext(context.Background())
}

func (i RouteEntryArray) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryArrayOutput)
}

// RouteEntryMapInput is an input type that accepts RouteEntryMap and RouteEntryMapOutput values.
// You can construct a concrete instance of `RouteEntryMapInput` via:
//
//	RouteEntryMap{ "key": RouteEntryArgs{...} }
type RouteEntryMapInput interface {
	pulumi.Input

	ToRouteEntryMapOutput() RouteEntryMapOutput
	ToRouteEntryMapOutputWithContext(context.Context) RouteEntryMapOutput
}

type RouteEntryMap map[string]RouteEntryInput

func (RouteEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (i RouteEntryMap) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return i.ToRouteEntryMapOutputWithContext(context.Background())
}

func (i RouteEntryMap) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteEntryMapOutput)
}

type RouteEntryOutput struct{ *pulumi.OutputState }

func (RouteEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteEntry)(nil)).Elem()
}

func (o RouteEntryOutput) ToRouteEntryOutput() RouteEntryOutput {
	return o
}

func (o RouteEntryOutput) ToRouteEntryOutputWithContext(ctx context.Context) RouteEntryOutput {
	return o
}

// The as path of the route entry.
func (o RouteEntryOutput) AsPath() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.AsPath }).(pulumi.StringOutput)
}

// The creation time of the route entry.
func (o RouteEntryOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Description of the transit router route entry.
func (o RouteEntryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The target network segment of the route entry.
func (o RouteEntryOutput) DestinationCidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.DestinationCidrBlock }).(pulumi.StringOutput)
}

// The status of the route entry.
func (o RouteEntryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The id of the route entry.
func (o RouteEntryOutput) TransitRouterRouteEntryId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.TransitRouterRouteEntryId }).(pulumi.StringOutput)
}

// The name of the route entry.
func (o RouteEntryOutput) TransitRouterRouteEntryName() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.TransitRouterRouteEntryName }).(pulumi.StringOutput)
}

// The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
func (o RouteEntryOutput) TransitRouterRouteEntryNextHopId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringPtrOutput { return v.TransitRouterRouteEntryNextHopId }).(pulumi.StringPtrOutput)
}

// The next hop type of the routing entry. The value can be Attachment or BlackHole.
func (o RouteEntryOutput) TransitRouterRouteEntryNextHopType() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.TransitRouterRouteEntryNextHopType }).(pulumi.StringOutput)
}

// The type of the route entry.
func (o RouteEntryOutput) TransitRouterRouteEntryType() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.TransitRouterRouteEntryType }).(pulumi.StringOutput)
}

// The id of the route table.
func (o RouteEntryOutput) TransitRouterRouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.TransitRouterRouteTableId }).(pulumi.StringOutput)
}

// The update time of the route entry.
func (o RouteEntryOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *RouteEntry) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type RouteEntryArrayOutput struct{ *pulumi.OutputState }

func (RouteEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutput() RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) ToRouteEntryArrayOutputWithContext(ctx context.Context) RouteEntryArrayOutput {
	return o
}

func (o RouteEntryArrayOutput) Index(i pulumi.IntInput) RouteEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].([]*RouteEntry)[vs[1].(int)]
	}).(RouteEntryOutput)
}

type RouteEntryMapOutput struct{ *pulumi.OutputState }

func (RouteEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouteEntry)(nil)).Elem()
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutput() RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) ToRouteEntryMapOutputWithContext(ctx context.Context) RouteEntryMapOutput {
	return o
}

func (o RouteEntryMapOutput) MapIndex(k pulumi.StringInput) RouteEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouteEntry {
		return vs[0].(map[string]*RouteEntry)[vs[1].(string)]
	}).(RouteEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryInput)(nil)).Elem(), &RouteEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryArrayInput)(nil)).Elem(), RouteEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteEntryMapInput)(nil)).Elem(), RouteEntryMap{})
	pulumi.RegisterOutputType(RouteEntryOutput{})
	pulumi.RegisterOutputType(RouteEntryArrayOutput{})
	pulumi.RegisterOutputType(RouteEntryMapOutput{})
}
