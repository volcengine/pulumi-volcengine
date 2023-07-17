// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage clb rule
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clb.NewRule(ctx, "foo", &clb.RuleArgs{
//				Domain:        pulumi.String("test-volc123.com"),
//				ListenerId:    pulumi.String("lsn-273ywvnmiu70g7fap8u2xzg9d"),
//				ServerGroupId: pulumi.String("rsp-273yxuqfova4g7fap8tyemn6t"),
//				Url:           pulumi.String("/yyyy"),
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
// Rule can be imported using the id, e.g. NoticeresourceId is ruleId, due to the lack of describeRuleAttributes in openapi, for import resources, please use ruleId:listenerId to import. we will fix this problem later.
//
// ```sh
//
//	$ pulumi import volcengine:clb/rule:Rule foo rule-273zb9hzi1gqo7fap8u1k3utb:lsn-273ywvnmiu70g7fap8u2xzg9d
//
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The description of the Rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The domain of Rule.
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// The ID of listener.
	ListenerId pulumi.StringOutput `pulumi:"listenerId"`
	// Server Group Id.
	ServerGroupId pulumi.StringOutput `pulumi:"serverGroupId"`
	// The Url of Rule.
	Url pulumi.StringPtrOutput `pulumi:"url"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ListenerId == nil {
		return nil, errors.New("invalid value for required argument 'ListenerId'")
	}
	if args.ServerGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroupId'")
	}
	var resource Rule
	err := ctx.RegisterResource("volcengine:clb/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("volcengine:clb/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The description of the Rule.
	Description *string `pulumi:"description"`
	// The domain of Rule.
	Domain *string `pulumi:"domain"`
	// The ID of listener.
	ListenerId *string `pulumi:"listenerId"`
	// Server Group Id.
	ServerGroupId *string `pulumi:"serverGroupId"`
	// The Url of Rule.
	Url *string `pulumi:"url"`
}

type RuleState struct {
	// The description of the Rule.
	Description pulumi.StringPtrInput
	// The domain of Rule.
	Domain pulumi.StringPtrInput
	// The ID of listener.
	ListenerId pulumi.StringPtrInput
	// Server Group Id.
	ServerGroupId pulumi.StringPtrInput
	// The Url of Rule.
	Url pulumi.StringPtrInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// The description of the Rule.
	Description *string `pulumi:"description"`
	// The domain of Rule.
	Domain *string `pulumi:"domain"`
	// The ID of listener.
	ListenerId string `pulumi:"listenerId"`
	// Server Group Id.
	ServerGroupId string `pulumi:"serverGroupId"`
	// The Url of Rule.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// The description of the Rule.
	Description pulumi.StringPtrInput
	// The domain of Rule.
	Domain pulumi.StringPtrInput
	// The ID of listener.
	ListenerId pulumi.StringInput
	// Server Group Id.
	ServerGroupId pulumi.StringInput
	// The Url of Rule.
	Url pulumi.StringPtrInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

// RuleArrayInput is an input type that accepts RuleArray and RuleArrayOutput values.
// You can construct a concrete instance of `RuleArrayInput` via:
//
//	RuleArray{ RuleArgs{...} }
type RuleArrayInput interface {
	pulumi.Input

	ToRuleArrayOutput() RuleArrayOutput
	ToRuleArrayOutputWithContext(context.Context) RuleArrayOutput
}

type RuleArray []RuleInput

func (RuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (i RuleArray) ToRuleArrayOutput() RuleArrayOutput {
	return i.ToRuleArrayOutputWithContext(context.Background())
}

func (i RuleArray) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleArrayOutput)
}

// RuleMapInput is an input type that accepts RuleMap and RuleMapOutput values.
// You can construct a concrete instance of `RuleMapInput` via:
//
//	RuleMap{ "key": RuleArgs{...} }
type RuleMapInput interface {
	pulumi.Input

	ToRuleMapOutput() RuleMapOutput
	ToRuleMapOutputWithContext(context.Context) RuleMapOutput
}

type RuleMap map[string]RuleInput

func (RuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (i RuleMap) ToRuleMapOutput() RuleMapOutput {
	return i.ToRuleMapOutputWithContext(context.Background())
}

func (i RuleMap) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleMapOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rule)(nil)).Elem()
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

// The description of the Rule.
func (o RuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The domain of Rule.
func (o RuleOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// The ID of listener.
func (o RuleOutput) ListenerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.ListenerId }).(pulumi.StringOutput)
}

// Server Group Id.
func (o RuleOutput) ServerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringOutput { return v.ServerGroupId }).(pulumi.StringOutput)
}

// The Url of Rule.
func (o RuleOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Rule) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

type RuleArrayOutput struct{ *pulumi.OutputState }

func (RuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rule)(nil)).Elem()
}

func (o RuleArrayOutput) ToRuleArrayOutput() RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) ToRuleArrayOutputWithContext(ctx context.Context) RuleArrayOutput {
	return o
}

func (o RuleArrayOutput) Index(i pulumi.IntInput) RuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].([]*Rule)[vs[1].(int)]
	}).(RuleOutput)
}

type RuleMapOutput struct{ *pulumi.OutputState }

func (RuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rule)(nil)).Elem()
}

func (o RuleMapOutput) ToRuleMapOutput() RuleMapOutput {
	return o
}

func (o RuleMapOutput) ToRuleMapOutputWithContext(ctx context.Context) RuleMapOutput {
	return o
}

func (o RuleMapOutput) MapIndex(k pulumi.StringInput) RuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rule {
		return vs[0].(map[string]*Rule)[vs[1].(string)]
	}).(RuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleArrayInput)(nil)).Elem(), RuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleMapInput)(nil)).Elem(), RuleMap{})
	pulumi.RegisterOutputType(RuleOutput{})
	pulumi.RegisterOutputType(RuleArrayOutput{})
	pulumi.RegisterOutputType(RuleMapOutput{})
}