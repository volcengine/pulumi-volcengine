// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package private_zone

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage private zone resolver rule
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/private_zone"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := private_zone.NewResolverRule(ctx, "foo", &private_zone.ResolverRuleArgs{
//				EndpointId: pulumi.Int(346),
//				ForwardIps: private_zone.ResolverRuleForwardIpArray{
//					&private_zone.ResolverRuleForwardIpArgs{
//						Ip:   pulumi.String("10.199.38.19"),
//						Port: pulumi.Int(33),
//					},
//				},
//				Type: pulumi.String("OUTBOUND"),
//				Vpcs: private_zone.ResolverRuleVpcArray{
//					&private_zone.ResolverRuleVpcArgs{
//						Region: pulumi.String("cn-beijing"),
//						VpcId:  pulumi.String("vpc-13f9uuuqfdjb43n6nu5p1****"),
//					},
//				},
//				ZoneNames: pulumi.StringArray{
//					pulumi.String("www.baidu.com"),
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
// PrivateZoneResolverRule can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:private_zone/resolverRule:ResolverRule default resource_id
// ```
type ResolverRule struct {
	pulumi.CustomResourceState

	// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
	EndpointId pulumi.IntPtrOutput `pulumi:"endpointId"`
	// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ForwardIps ResolverRuleForwardIpArrayOutput `pulumi:"forwardIps"`
	// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
	Line pulumi.IntPtrOutput `pulumi:"line"`
	// The name of the rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
	Type pulumi.StringOutput `pulumi:"type"`
	// The parameter name <region> is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
	Vpcs ResolverRuleVpcArrayOutput `pulumi:"vpcs"`
	// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ZoneNames pulumi.StringArrayOutput `pulumi:"zoneNames"`
}

// NewResolverRule registers a new resource with the given unique name, arguments, and options.
func NewResolverRule(ctx *pulumi.Context,
	name string, args *ResolverRuleArgs, opts ...pulumi.ResourceOption) (*ResolverRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Vpcs == nil {
		return nil, errors.New("invalid value for required argument 'Vpcs'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ResolverRule
	err := ctx.RegisterResource("volcengine:private_zone/resolverRule:ResolverRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverRule gets an existing ResolverRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverRuleState, opts ...pulumi.ResourceOption) (*ResolverRule, error) {
	var resource ResolverRule
	err := ctx.ReadResource("volcengine:private_zone/resolverRule:ResolverRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverRule resources.
type resolverRuleState struct {
	// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
	EndpointId *int `pulumi:"endpointId"`
	// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ForwardIps []ResolverRuleForwardIp `pulumi:"forwardIps"`
	// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
	Line *int `pulumi:"line"`
	// The name of the rule.
	Name *string `pulumi:"name"`
	// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
	Type *string `pulumi:"type"`
	// The parameter name <region> is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
	Vpcs []ResolverRuleVpc `pulumi:"vpcs"`
	// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ZoneNames []string `pulumi:"zoneNames"`
}

type ResolverRuleState struct {
	// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
	EndpointId pulumi.IntPtrInput
	// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ForwardIps ResolverRuleForwardIpArrayInput
	// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
	Line pulumi.IntPtrInput
	// The name of the rule.
	Name pulumi.StringPtrInput
	// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
	Type pulumi.StringPtrInput
	// The parameter name <region> is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
	Vpcs ResolverRuleVpcArrayInput
	// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ZoneNames pulumi.StringArrayInput
}

func (ResolverRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleState)(nil)).Elem()
}

type resolverRuleArgs struct {
	// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
	EndpointId *int `pulumi:"endpointId"`
	// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ForwardIps []ResolverRuleForwardIp `pulumi:"forwardIps"`
	// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
	Line *int `pulumi:"line"`
	// The name of the rule.
	Name *string `pulumi:"name"`
	// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
	Type string `pulumi:"type"`
	// The parameter name <region> is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
	Vpcs []ResolverRuleVpc `pulumi:"vpcs"`
	// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ZoneNames []string `pulumi:"zoneNames"`
}

// The set of arguments for constructing a ResolverRule resource.
type ResolverRuleArgs struct {
	// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
	EndpointId pulumi.IntPtrInput
	// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ForwardIps ResolverRuleForwardIpArrayInput
	// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
	Line pulumi.IntPtrInput
	// The name of the rule.
	Name pulumi.StringPtrInput
	// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
	Type pulumi.StringInput
	// The parameter name <region> is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
	Vpcs ResolverRuleVpcArrayInput
	// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
	ZoneNames pulumi.StringArrayInput
}

func (ResolverRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleArgs)(nil)).Elem()
}

type ResolverRuleInput interface {
	pulumi.Input

	ToResolverRuleOutput() ResolverRuleOutput
	ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput
}

func (*ResolverRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverRule)(nil)).Elem()
}

func (i *ResolverRule) ToResolverRuleOutput() ResolverRuleOutput {
	return i.ToResolverRuleOutputWithContext(context.Background())
}

func (i *ResolverRule) ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleOutput)
}

// ResolverRuleArrayInput is an input type that accepts ResolverRuleArray and ResolverRuleArrayOutput values.
// You can construct a concrete instance of `ResolverRuleArrayInput` via:
//
//	ResolverRuleArray{ ResolverRuleArgs{...} }
type ResolverRuleArrayInput interface {
	pulumi.Input

	ToResolverRuleArrayOutput() ResolverRuleArrayOutput
	ToResolverRuleArrayOutputWithContext(context.Context) ResolverRuleArrayOutput
}

type ResolverRuleArray []ResolverRuleInput

func (ResolverRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverRule)(nil)).Elem()
}

func (i ResolverRuleArray) ToResolverRuleArrayOutput() ResolverRuleArrayOutput {
	return i.ToResolverRuleArrayOutputWithContext(context.Background())
}

func (i ResolverRuleArray) ToResolverRuleArrayOutputWithContext(ctx context.Context) ResolverRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleArrayOutput)
}

// ResolverRuleMapInput is an input type that accepts ResolverRuleMap and ResolverRuleMapOutput values.
// You can construct a concrete instance of `ResolverRuleMapInput` via:
//
//	ResolverRuleMap{ "key": ResolverRuleArgs{...} }
type ResolverRuleMapInput interface {
	pulumi.Input

	ToResolverRuleMapOutput() ResolverRuleMapOutput
	ToResolverRuleMapOutputWithContext(context.Context) ResolverRuleMapOutput
}

type ResolverRuleMap map[string]ResolverRuleInput

func (ResolverRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverRule)(nil)).Elem()
}

func (i ResolverRuleMap) ToResolverRuleMapOutput() ResolverRuleMapOutput {
	return i.ToResolverRuleMapOutputWithContext(context.Background())
}

func (i ResolverRuleMap) ToResolverRuleMapOutputWithContext(ctx context.Context) ResolverRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleMapOutput)
}

type ResolverRuleOutput struct{ *pulumi.OutputState }

func (ResolverRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverRule)(nil)).Elem()
}

func (o ResolverRuleOutput) ToResolverRuleOutput() ResolverRuleOutput {
	return o
}

func (o ResolverRuleOutput) ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput {
	return o
}

// Terminal node ID. This parameter is only valid and required when the Type parameter is OUTBOUND.
func (o ResolverRuleOutput) EndpointId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.IntPtrOutput { return v.EndpointId }).(pulumi.IntPtrOutput)
}

// IP address and port of external DNS server. You can add up to 10 IP addresses. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
func (o ResolverRuleOutput) ForwardIps() ResolverRuleForwardIpArrayOutput {
	return o.ApplyT(func(v *ResolverRule) ResolverRuleForwardIpArrayOutput { return v.ForwardIps }).(ResolverRuleForwardIpArrayOutput)
}

// The operator of the exit IP address of the recursive DNS server. This parameter is only valid when the Type parameter is LINE and is a required parameter. MOBILE, TELECOM, UNICOM.
func (o ResolverRuleOutput) Line() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.IntPtrOutput { return v.Line }).(pulumi.IntPtrOutput)
}

// The name of the rule.
func (o ResolverRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Forwarding rule types. OUTBOUND: Forward to external DNS servers. LINE: Set the recursive DNS server used for recursive resolution to the recursive DNS server of the Volcano Engine PublicDNS, and customize the operator's exit IP address for the recursive DNS server.
func (o ResolverRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The parameter name <region> is a variable that represents the region where the VPC is located, such as cn-beijing. The parameter value can include one or more VPC IDs, such as vpc-2750bd1. For example, if you associate a VPC in the cn-beijing region with a domain name and the VPC ID is vpc-2d6si87atfh1c58ozfd0nzq8k, the parameter would be "cn-beijing":["vpc-2d6si87atfh1c58ozfd0nzq8k"]. You can add one or more regions. When the Type parameter is OUTBOUND, the VPC region must be the same as the region where the endpoint is located.
func (o ResolverRuleOutput) Vpcs() ResolverRuleVpcArrayOutput {
	return o.ApplyT(func(v *ResolverRule) ResolverRuleVpcArrayOutput { return v.Vpcs }).(ResolverRuleVpcArrayOutput)
}

// Domain names associated with forwarding rules. You can enter one or more domain names. Up to 500 domain names are supported. This parameter is only valid when the Type parameter is OUTBOUND and is a required parameter.
func (o ResolverRuleOutput) ZoneNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ResolverRule) pulumi.StringArrayOutput { return v.ZoneNames }).(pulumi.StringArrayOutput)
}

type ResolverRuleArrayOutput struct{ *pulumi.OutputState }

func (ResolverRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResolverRule)(nil)).Elem()
}

func (o ResolverRuleArrayOutput) ToResolverRuleArrayOutput() ResolverRuleArrayOutput {
	return o
}

func (o ResolverRuleArrayOutput) ToResolverRuleArrayOutputWithContext(ctx context.Context) ResolverRuleArrayOutput {
	return o
}

func (o ResolverRuleArrayOutput) Index(i pulumi.IntInput) ResolverRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResolverRule {
		return vs[0].([]*ResolverRule)[vs[1].(int)]
	}).(ResolverRuleOutput)
}

type ResolverRuleMapOutput struct{ *pulumi.OutputState }

func (ResolverRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResolverRule)(nil)).Elem()
}

func (o ResolverRuleMapOutput) ToResolverRuleMapOutput() ResolverRuleMapOutput {
	return o
}

func (o ResolverRuleMapOutput) ToResolverRuleMapOutputWithContext(ctx context.Context) ResolverRuleMapOutput {
	return o
}

func (o ResolverRuleMapOutput) MapIndex(k pulumi.StringInput) ResolverRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResolverRule {
		return vs[0].(map[string]*ResolverRule)[vs[1].(string)]
	}).(ResolverRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverRuleInput)(nil)).Elem(), &ResolverRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverRuleArrayInput)(nil)).Elem(), ResolverRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResolverRuleMapInput)(nil)).Elem(), ResolverRuleMap{})
	pulumi.RegisterOutputType(ResolverRuleOutput{})
	pulumi.RegisterOutputType(ResolverRuleArrayOutput{})
	pulumi.RegisterOutputType(ResolverRuleMapOutput{})
}