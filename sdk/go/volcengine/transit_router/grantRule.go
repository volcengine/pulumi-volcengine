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

// Provides a resource to manage transit router grant rule
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/transit_router"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := transit_router.NewGrantRule(ctx, "foo", &transit_router.GrantRuleArgs{
//				Description:     pulumi.String("tf-test"),
//				GrantAccountId:  pulumi.String("200000xxxx"),
//				TransitRouterId: pulumi.String("tr-2bzy39uy6u3282dx0efxiqyq0"),
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
// TransitRouterGrantRule can be imported using the transit router id and accountId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:transit_router/grantRule:GrantRule default trId:accountId
//
// ```
type GrantRule struct {
	pulumi.CustomResourceState

	// The description of the rule.
	Description pulumi.StringOutput `pulumi:"description"`
	// Account ID awaiting authorization for intermediate router instance.
	GrantAccountId pulumi.StringOutput `pulumi:"grantAccountId"`
	// The id of the transit router.
	TransitRouterId pulumi.StringOutput `pulumi:"transitRouterId"`
}

// NewGrantRule registers a new resource with the given unique name, arguments, and options.
func NewGrantRule(ctx *pulumi.Context,
	name string, args *GrantRuleArgs, opts ...pulumi.ResourceOption) (*GrantRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GrantAccountId == nil {
		return nil, errors.New("invalid value for required argument 'GrantAccountId'")
	}
	if args.TransitRouterId == nil {
		return nil, errors.New("invalid value for required argument 'TransitRouterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GrantRule
	err := ctx.RegisterResource("volcengine:transit_router/grantRule:GrantRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrantRule gets an existing GrantRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrantRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrantRuleState, opts ...pulumi.ResourceOption) (*GrantRule, error) {
	var resource GrantRule
	err := ctx.ReadResource("volcengine:transit_router/grantRule:GrantRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GrantRule resources.
type grantRuleState struct {
	// The description of the rule.
	Description *string `pulumi:"description"`
	// Account ID awaiting authorization for intermediate router instance.
	GrantAccountId *string `pulumi:"grantAccountId"`
	// The id of the transit router.
	TransitRouterId *string `pulumi:"transitRouterId"`
}

type GrantRuleState struct {
	// The description of the rule.
	Description pulumi.StringPtrInput
	// Account ID awaiting authorization for intermediate router instance.
	GrantAccountId pulumi.StringPtrInput
	// The id of the transit router.
	TransitRouterId pulumi.StringPtrInput
}

func (GrantRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*grantRuleState)(nil)).Elem()
}

type grantRuleArgs struct {
	// The description of the rule.
	Description *string `pulumi:"description"`
	// Account ID awaiting authorization for intermediate router instance.
	GrantAccountId string `pulumi:"grantAccountId"`
	// The id of the transit router.
	TransitRouterId string `pulumi:"transitRouterId"`
}

// The set of arguments for constructing a GrantRule resource.
type GrantRuleArgs struct {
	// The description of the rule.
	Description pulumi.StringPtrInput
	// Account ID awaiting authorization for intermediate router instance.
	GrantAccountId pulumi.StringInput
	// The id of the transit router.
	TransitRouterId pulumi.StringInput
}

func (GrantRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grantRuleArgs)(nil)).Elem()
}

type GrantRuleInput interface {
	pulumi.Input

	ToGrantRuleOutput() GrantRuleOutput
	ToGrantRuleOutputWithContext(ctx context.Context) GrantRuleOutput
}

func (*GrantRule) ElementType() reflect.Type {
	return reflect.TypeOf((**GrantRule)(nil)).Elem()
}

func (i *GrantRule) ToGrantRuleOutput() GrantRuleOutput {
	return i.ToGrantRuleOutputWithContext(context.Background())
}

func (i *GrantRule) ToGrantRuleOutputWithContext(ctx context.Context) GrantRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantRuleOutput)
}

// GrantRuleArrayInput is an input type that accepts GrantRuleArray and GrantRuleArrayOutput values.
// You can construct a concrete instance of `GrantRuleArrayInput` via:
//
//	GrantRuleArray{ GrantRuleArgs{...} }
type GrantRuleArrayInput interface {
	pulumi.Input

	ToGrantRuleArrayOutput() GrantRuleArrayOutput
	ToGrantRuleArrayOutputWithContext(context.Context) GrantRuleArrayOutput
}

type GrantRuleArray []GrantRuleInput

func (GrantRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrantRule)(nil)).Elem()
}

func (i GrantRuleArray) ToGrantRuleArrayOutput() GrantRuleArrayOutput {
	return i.ToGrantRuleArrayOutputWithContext(context.Background())
}

func (i GrantRuleArray) ToGrantRuleArrayOutputWithContext(ctx context.Context) GrantRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantRuleArrayOutput)
}

// GrantRuleMapInput is an input type that accepts GrantRuleMap and GrantRuleMapOutput values.
// You can construct a concrete instance of `GrantRuleMapInput` via:
//
//	GrantRuleMap{ "key": GrantRuleArgs{...} }
type GrantRuleMapInput interface {
	pulumi.Input

	ToGrantRuleMapOutput() GrantRuleMapOutput
	ToGrantRuleMapOutputWithContext(context.Context) GrantRuleMapOutput
}

type GrantRuleMap map[string]GrantRuleInput

func (GrantRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrantRule)(nil)).Elem()
}

func (i GrantRuleMap) ToGrantRuleMapOutput() GrantRuleMapOutput {
	return i.ToGrantRuleMapOutputWithContext(context.Background())
}

func (i GrantRuleMap) ToGrantRuleMapOutputWithContext(ctx context.Context) GrantRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantRuleMapOutput)
}

type GrantRuleOutput struct{ *pulumi.OutputState }

func (GrantRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrantRule)(nil)).Elem()
}

func (o GrantRuleOutput) ToGrantRuleOutput() GrantRuleOutput {
	return o
}

func (o GrantRuleOutput) ToGrantRuleOutputWithContext(ctx context.Context) GrantRuleOutput {
	return o
}

// The description of the rule.
func (o GrantRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Account ID awaiting authorization for intermediate router instance.
func (o GrantRuleOutput) GrantAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantRule) pulumi.StringOutput { return v.GrantAccountId }).(pulumi.StringOutput)
}

// The id of the transit router.
func (o GrantRuleOutput) TransitRouterId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantRule) pulumi.StringOutput { return v.TransitRouterId }).(pulumi.StringOutput)
}

type GrantRuleArrayOutput struct{ *pulumi.OutputState }

func (GrantRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrantRule)(nil)).Elem()
}

func (o GrantRuleArrayOutput) ToGrantRuleArrayOutput() GrantRuleArrayOutput {
	return o
}

func (o GrantRuleArrayOutput) ToGrantRuleArrayOutputWithContext(ctx context.Context) GrantRuleArrayOutput {
	return o
}

func (o GrantRuleArrayOutput) Index(i pulumi.IntInput) GrantRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GrantRule {
		return vs[0].([]*GrantRule)[vs[1].(int)]
	}).(GrantRuleOutput)
}

type GrantRuleMapOutput struct{ *pulumi.OutputState }

func (GrantRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrantRule)(nil)).Elem()
}

func (o GrantRuleMapOutput) ToGrantRuleMapOutput() GrantRuleMapOutput {
	return o
}

func (o GrantRuleMapOutput) ToGrantRuleMapOutputWithContext(ctx context.Context) GrantRuleMapOutput {
	return o
}

func (o GrantRuleMapOutput) MapIndex(k pulumi.StringInput) GrantRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GrantRule {
		return vs[0].(map[string]*GrantRule)[vs[1].(string)]
	}).(GrantRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GrantRuleInput)(nil)).Elem(), &GrantRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrantRuleArrayInput)(nil)).Elem(), GrantRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrantRuleMapInput)(nil)).Elem(), GrantRuleMap{})
	pulumi.RegisterOutputType(GrantRuleOutput{})
	pulumi.RegisterOutputType(GrantRuleArrayOutput{})
	pulumi.RegisterOutputType(GrantRuleMapOutput{})
}