// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage waf bot analyse protect rule
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/waf"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := waf.NewBotAnalyseProtectRule(ctx, "foo", &waf.BotAnalyseProtectRuleArgs{
//				AccurateGroup: &waf.BotAnalyseProtectRuleAccurateGroupArgs{
//					AccurateRules: waf.BotAnalyseProtectRuleAccurateGroupAccurateRuleArray{
//						&waf.BotAnalyseProtectRuleAccurateGroupAccurateRuleArgs{
//							HttpObj:     pulumi.String("request.uri"),
//							ObjType:     pulumi.Int(1),
//							Opretar:     pulumi.Int(2),
//							Property:    pulumi.Int(0),
//							ValueString: pulumi.String("tf"),
//						},
//						&waf.BotAnalyseProtectRuleAccurateGroupAccurateRuleArgs{
//							HttpObj:     pulumi.String("request.schema"),
//							ObjType:     pulumi.Int(0),
//							Opretar:     pulumi.Int(2),
//							Property:    pulumi.Int(0),
//							ValueString: pulumi.String("tf-2"),
//						},
//					},
//					Logic: pulumi.Int(2),
//				},
//				ActionAfterVerification: pulumi.Int(1),
//				ActionType:              pulumi.Int(1),
//				EffectTime:              pulumi.Int(1000),
//				Enable:                  pulumi.Int(1),
//				ExemptionTime:           pulumi.Int(60),
//				Field:                   pulumi.String("HEADER:User-Agent"),
//				Host:                    pulumi.String("www.tf-test.com"),
//				Path:                    pulumi.String("/mod"),
//				PathThreshold:           pulumi.Int(1000),
//				ProjectName:             pulumi.String("default"),
//				RulePriority:            pulumi.Int(3),
//				SingleProportion:        pulumi.Float64(0.25),
//				SingleThreshold:         pulumi.Int(100),
//				StatisticalDuration:     pulumi.Int(50),
//				StatisticalType:         pulumi.Int(2),
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
// WafBotAnalyseProtectRule can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:waf/botAnalyseProtectRule:BotAnalyseProtectRule default resource_id:bot_space:host
// ```
type BotAnalyseProtectRule struct {
	pulumi.CustomResourceState

	// Advanced conditions.
	AccurateGroup BotAnalyseProtectRuleAccurateGroupPtrOutput `pulumi:"accurateGroup"`
	// Perform the action after verification/challenge.
	ActionAfterVerification pulumi.IntPtrOutput `pulumi:"actionAfterVerification"`
	// perform the action.
	ActionType pulumi.IntOutput `pulumi:"actionType"`
	// Limit the duration.
	EffectTime pulumi.IntOutput `pulumi:"effectTime"`
	// Whether to enable the rules.
	Enable pulumi.IntOutput `pulumi:"enable"`
	// The number of statistical protection rules enabled under the current domain name.
	EnableCount pulumi.IntOutput `pulumi:"enableCount"`
	// Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
	ExemptionTime pulumi.IntPtrOutput `pulumi:"exemptionTime"`
	// Statistical objects, with multiple objects separated by commas.
	Field pulumi.StringOutput `pulumi:"field"`
	// Website domain names that require the setting of protection rules.
	Host pulumi.StringOutput `pulumi:"host"`
	// The name of rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The requested path.
	Path pulumi.StringOutput `pulumi:"path"`
	// The path access frequency threshold is enabled when StatisticalType=1.
	PathThreshold pulumi.IntPtrOutput `pulumi:"pathThreshold"`
	// The Name of the affiliated project resource.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// Details of the rule group.
	RuleGroups BotAnalyseProtectRuleRuleGroupArrayOutput `pulumi:"ruleGroups"`
	// Priority of rule effectiveness.
	RulePriority pulumi.IntOutput `pulumi:"rulePriority"`
	// The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
	SingleProportion pulumi.Float64PtrOutput `pulumi:"singleProportion"`
	// The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
	SingleThreshold pulumi.IntOutput `pulumi:"singleThreshold"`
	// The duration of statistics.
	StatisticalDuration pulumi.IntOutput `pulumi:"statisticalDuration"`
	// Statistical content and methods.
	StatisticalType pulumi.IntOutput `pulumi:"statisticalType"`
	// The total number of statistical protection rules under the current domain name.
	TotalCount pulumi.IntOutput `pulumi:"totalCount"`
}

// NewBotAnalyseProtectRule registers a new resource with the given unique name, arguments, and options.
func NewBotAnalyseProtectRule(ctx *pulumi.Context,
	name string, args *BotAnalyseProtectRuleArgs, opts ...pulumi.ResourceOption) (*BotAnalyseProtectRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ActionType == nil {
		return nil, errors.New("invalid value for required argument 'ActionType'")
	}
	if args.EffectTime == nil {
		return nil, errors.New("invalid value for required argument 'EffectTime'")
	}
	if args.Enable == nil {
		return nil, errors.New("invalid value for required argument 'Enable'")
	}
	if args.Field == nil {
		return nil, errors.New("invalid value for required argument 'Field'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.RulePriority == nil {
		return nil, errors.New("invalid value for required argument 'RulePriority'")
	}
	if args.SingleThreshold == nil {
		return nil, errors.New("invalid value for required argument 'SingleThreshold'")
	}
	if args.StatisticalDuration == nil {
		return nil, errors.New("invalid value for required argument 'StatisticalDuration'")
	}
	if args.StatisticalType == nil {
		return nil, errors.New("invalid value for required argument 'StatisticalType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BotAnalyseProtectRule
	err := ctx.RegisterResource("volcengine:waf/botAnalyseProtectRule:BotAnalyseProtectRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBotAnalyseProtectRule gets an existing BotAnalyseProtectRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBotAnalyseProtectRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BotAnalyseProtectRuleState, opts ...pulumi.ResourceOption) (*BotAnalyseProtectRule, error) {
	var resource BotAnalyseProtectRule
	err := ctx.ReadResource("volcengine:waf/botAnalyseProtectRule:BotAnalyseProtectRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BotAnalyseProtectRule resources.
type botAnalyseProtectRuleState struct {
	// Advanced conditions.
	AccurateGroup *BotAnalyseProtectRuleAccurateGroup `pulumi:"accurateGroup"`
	// Perform the action after verification/challenge.
	ActionAfterVerification *int `pulumi:"actionAfterVerification"`
	// perform the action.
	ActionType *int `pulumi:"actionType"`
	// Limit the duration.
	EffectTime *int `pulumi:"effectTime"`
	// Whether to enable the rules.
	Enable *int `pulumi:"enable"`
	// The number of statistical protection rules enabled under the current domain name.
	EnableCount *int `pulumi:"enableCount"`
	// Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
	ExemptionTime *int `pulumi:"exemptionTime"`
	// Statistical objects, with multiple objects separated by commas.
	Field *string `pulumi:"field"`
	// Website domain names that require the setting of protection rules.
	Host *string `pulumi:"host"`
	// The name of rule.
	Name *string `pulumi:"name"`
	// The requested path.
	Path *string `pulumi:"path"`
	// The path access frequency threshold is enabled when StatisticalType=1.
	PathThreshold *int `pulumi:"pathThreshold"`
	// The Name of the affiliated project resource.
	ProjectName *string `pulumi:"projectName"`
	// Details of the rule group.
	RuleGroups []BotAnalyseProtectRuleRuleGroup `pulumi:"ruleGroups"`
	// Priority of rule effectiveness.
	RulePriority *int `pulumi:"rulePriority"`
	// The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
	SingleProportion *float64 `pulumi:"singleProportion"`
	// The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
	SingleThreshold *int `pulumi:"singleThreshold"`
	// The duration of statistics.
	StatisticalDuration *int `pulumi:"statisticalDuration"`
	// Statistical content and methods.
	StatisticalType *int `pulumi:"statisticalType"`
	// The total number of statistical protection rules under the current domain name.
	TotalCount *int `pulumi:"totalCount"`
}

type BotAnalyseProtectRuleState struct {
	// Advanced conditions.
	AccurateGroup BotAnalyseProtectRuleAccurateGroupPtrInput
	// Perform the action after verification/challenge.
	ActionAfterVerification pulumi.IntPtrInput
	// perform the action.
	ActionType pulumi.IntPtrInput
	// Limit the duration.
	EffectTime pulumi.IntPtrInput
	// Whether to enable the rules.
	Enable pulumi.IntPtrInput
	// The number of statistical protection rules enabled under the current domain name.
	EnableCount pulumi.IntPtrInput
	// Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
	ExemptionTime pulumi.IntPtrInput
	// Statistical objects, with multiple objects separated by commas.
	Field pulumi.StringPtrInput
	// Website domain names that require the setting of protection rules.
	Host pulumi.StringPtrInput
	// The name of rule.
	Name pulumi.StringPtrInput
	// The requested path.
	Path pulumi.StringPtrInput
	// The path access frequency threshold is enabled when StatisticalType=1.
	PathThreshold pulumi.IntPtrInput
	// The Name of the affiliated project resource.
	ProjectName pulumi.StringPtrInput
	// Details of the rule group.
	RuleGroups BotAnalyseProtectRuleRuleGroupArrayInput
	// Priority of rule effectiveness.
	RulePriority pulumi.IntPtrInput
	// The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
	SingleProportion pulumi.Float64PtrInput
	// The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
	SingleThreshold pulumi.IntPtrInput
	// The duration of statistics.
	StatisticalDuration pulumi.IntPtrInput
	// Statistical content and methods.
	StatisticalType pulumi.IntPtrInput
	// The total number of statistical protection rules under the current domain name.
	TotalCount pulumi.IntPtrInput
}

func (BotAnalyseProtectRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*botAnalyseProtectRuleState)(nil)).Elem()
}

type botAnalyseProtectRuleArgs struct {
	// Advanced conditions.
	AccurateGroup *BotAnalyseProtectRuleAccurateGroup `pulumi:"accurateGroup"`
	// Perform the action after verification/challenge.
	ActionAfterVerification *int `pulumi:"actionAfterVerification"`
	// perform the action.
	ActionType int `pulumi:"actionType"`
	// Limit the duration.
	EffectTime int `pulumi:"effectTime"`
	// Whether to enable the rules.
	Enable int `pulumi:"enable"`
	// Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
	ExemptionTime *int `pulumi:"exemptionTime"`
	// Statistical objects, with multiple objects separated by commas.
	Field string `pulumi:"field"`
	// Website domain names that require the setting of protection rules.
	Host string `pulumi:"host"`
	// The name of rule.
	Name *string `pulumi:"name"`
	// The requested path.
	Path string `pulumi:"path"`
	// The path access frequency threshold is enabled when StatisticalType=1.
	PathThreshold *int `pulumi:"pathThreshold"`
	// The Name of the affiliated project resource.
	ProjectName *string `pulumi:"projectName"`
	// Priority of rule effectiveness.
	RulePriority int `pulumi:"rulePriority"`
	// The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
	SingleProportion *float64 `pulumi:"singleProportion"`
	// The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
	SingleThreshold int `pulumi:"singleThreshold"`
	// The duration of statistics.
	StatisticalDuration int `pulumi:"statisticalDuration"`
	// Statistical content and methods.
	StatisticalType int `pulumi:"statisticalType"`
}

// The set of arguments for constructing a BotAnalyseProtectRule resource.
type BotAnalyseProtectRuleArgs struct {
	// Advanced conditions.
	AccurateGroup BotAnalyseProtectRuleAccurateGroupPtrInput
	// Perform the action after verification/challenge.
	ActionAfterVerification pulumi.IntPtrInput
	// perform the action.
	ActionType pulumi.IntInput
	// Limit the duration.
	EffectTime pulumi.IntInput
	// Whether to enable the rules.
	Enable pulumi.IntInput
	// Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
	ExemptionTime pulumi.IntPtrInput
	// Statistical objects, with multiple objects separated by commas.
	Field pulumi.StringInput
	// Website domain names that require the setting of protection rules.
	Host pulumi.StringInput
	// The name of rule.
	Name pulumi.StringPtrInput
	// The requested path.
	Path pulumi.StringInput
	// The path access frequency threshold is enabled when StatisticalType=1.
	PathThreshold pulumi.IntPtrInput
	// The Name of the affiliated project resource.
	ProjectName pulumi.StringPtrInput
	// Priority of rule effectiveness.
	RulePriority pulumi.IntInput
	// The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
	SingleProportion pulumi.Float64PtrInput
	// The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
	SingleThreshold pulumi.IntInput
	// The duration of statistics.
	StatisticalDuration pulumi.IntInput
	// Statistical content and methods.
	StatisticalType pulumi.IntInput
}

func (BotAnalyseProtectRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*botAnalyseProtectRuleArgs)(nil)).Elem()
}

type BotAnalyseProtectRuleInput interface {
	pulumi.Input

	ToBotAnalyseProtectRuleOutput() BotAnalyseProtectRuleOutput
	ToBotAnalyseProtectRuleOutputWithContext(ctx context.Context) BotAnalyseProtectRuleOutput
}

func (*BotAnalyseProtectRule) ElementType() reflect.Type {
	return reflect.TypeOf((**BotAnalyseProtectRule)(nil)).Elem()
}

func (i *BotAnalyseProtectRule) ToBotAnalyseProtectRuleOutput() BotAnalyseProtectRuleOutput {
	return i.ToBotAnalyseProtectRuleOutputWithContext(context.Background())
}

func (i *BotAnalyseProtectRule) ToBotAnalyseProtectRuleOutputWithContext(ctx context.Context) BotAnalyseProtectRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotAnalyseProtectRuleOutput)
}

// BotAnalyseProtectRuleArrayInput is an input type that accepts BotAnalyseProtectRuleArray and BotAnalyseProtectRuleArrayOutput values.
// You can construct a concrete instance of `BotAnalyseProtectRuleArrayInput` via:
//
//	BotAnalyseProtectRuleArray{ BotAnalyseProtectRuleArgs{...} }
type BotAnalyseProtectRuleArrayInput interface {
	pulumi.Input

	ToBotAnalyseProtectRuleArrayOutput() BotAnalyseProtectRuleArrayOutput
	ToBotAnalyseProtectRuleArrayOutputWithContext(context.Context) BotAnalyseProtectRuleArrayOutput
}

type BotAnalyseProtectRuleArray []BotAnalyseProtectRuleInput

func (BotAnalyseProtectRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BotAnalyseProtectRule)(nil)).Elem()
}

func (i BotAnalyseProtectRuleArray) ToBotAnalyseProtectRuleArrayOutput() BotAnalyseProtectRuleArrayOutput {
	return i.ToBotAnalyseProtectRuleArrayOutputWithContext(context.Background())
}

func (i BotAnalyseProtectRuleArray) ToBotAnalyseProtectRuleArrayOutputWithContext(ctx context.Context) BotAnalyseProtectRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotAnalyseProtectRuleArrayOutput)
}

// BotAnalyseProtectRuleMapInput is an input type that accepts BotAnalyseProtectRuleMap and BotAnalyseProtectRuleMapOutput values.
// You can construct a concrete instance of `BotAnalyseProtectRuleMapInput` via:
//
//	BotAnalyseProtectRuleMap{ "key": BotAnalyseProtectRuleArgs{...} }
type BotAnalyseProtectRuleMapInput interface {
	pulumi.Input

	ToBotAnalyseProtectRuleMapOutput() BotAnalyseProtectRuleMapOutput
	ToBotAnalyseProtectRuleMapOutputWithContext(context.Context) BotAnalyseProtectRuleMapOutput
}

type BotAnalyseProtectRuleMap map[string]BotAnalyseProtectRuleInput

func (BotAnalyseProtectRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BotAnalyseProtectRule)(nil)).Elem()
}

func (i BotAnalyseProtectRuleMap) ToBotAnalyseProtectRuleMapOutput() BotAnalyseProtectRuleMapOutput {
	return i.ToBotAnalyseProtectRuleMapOutputWithContext(context.Background())
}

func (i BotAnalyseProtectRuleMap) ToBotAnalyseProtectRuleMapOutputWithContext(ctx context.Context) BotAnalyseProtectRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotAnalyseProtectRuleMapOutput)
}

type BotAnalyseProtectRuleOutput struct{ *pulumi.OutputState }

func (BotAnalyseProtectRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BotAnalyseProtectRule)(nil)).Elem()
}

func (o BotAnalyseProtectRuleOutput) ToBotAnalyseProtectRuleOutput() BotAnalyseProtectRuleOutput {
	return o
}

func (o BotAnalyseProtectRuleOutput) ToBotAnalyseProtectRuleOutputWithContext(ctx context.Context) BotAnalyseProtectRuleOutput {
	return o
}

// Advanced conditions.
func (o BotAnalyseProtectRuleOutput) AccurateGroup() BotAnalyseProtectRuleAccurateGroupPtrOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) BotAnalyseProtectRuleAccurateGroupPtrOutput { return v.AccurateGroup }).(BotAnalyseProtectRuleAccurateGroupPtrOutput)
}

// Perform the action after verification/challenge.
func (o BotAnalyseProtectRuleOutput) ActionAfterVerification() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntPtrOutput { return v.ActionAfterVerification }).(pulumi.IntPtrOutput)
}

// perform the action.
func (o BotAnalyseProtectRuleOutput) ActionType() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.ActionType }).(pulumi.IntOutput)
}

// Limit the duration.
func (o BotAnalyseProtectRuleOutput) EffectTime() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.EffectTime }).(pulumi.IntOutput)
}

// Whether to enable the rules.
func (o BotAnalyseProtectRuleOutput) Enable() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.Enable }).(pulumi.IntOutput)
}

// The number of statistical protection rules enabled under the current domain name.
func (o BotAnalyseProtectRuleOutput) EnableCount() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.EnableCount }).(pulumi.IntOutput)
}

// Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.
func (o BotAnalyseProtectRuleOutput) ExemptionTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntPtrOutput { return v.ExemptionTime }).(pulumi.IntPtrOutput)
}

// Statistical objects, with multiple objects separated by commas.
func (o BotAnalyseProtectRuleOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.StringOutput { return v.Field }).(pulumi.StringOutput)
}

// Website domain names that require the setting of protection rules.
func (o BotAnalyseProtectRuleOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// The name of rule.
func (o BotAnalyseProtectRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The requested path.
func (o BotAnalyseProtectRuleOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// The path access frequency threshold is enabled when StatisticalType=1.
func (o BotAnalyseProtectRuleOutput) PathThreshold() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntPtrOutput { return v.PathThreshold }).(pulumi.IntPtrOutput)
}

// The Name of the affiliated project resource.
func (o BotAnalyseProtectRuleOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Details of the rule group.
func (o BotAnalyseProtectRuleOutput) RuleGroups() BotAnalyseProtectRuleRuleGroupArrayOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) BotAnalyseProtectRuleRuleGroupArrayOutput { return v.RuleGroups }).(BotAnalyseProtectRuleRuleGroupArrayOutput)
}

// Priority of rule effectiveness.
func (o BotAnalyseProtectRuleOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.RulePriority }).(pulumi.IntOutput)
}

// The IP proportion of the same statistical object needs to be configured when StatisticalType=3.
func (o BotAnalyseProtectRuleOutput) SingleProportion() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.Float64PtrOutput { return v.SingleProportion }).(pulumi.Float64PtrOutput)
}

// The maximum number of ips of the same statistical object is enabled when StatisticalType=2.
func (o BotAnalyseProtectRuleOutput) SingleThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.SingleThreshold }).(pulumi.IntOutput)
}

// The duration of statistics.
func (o BotAnalyseProtectRuleOutput) StatisticalDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.StatisticalDuration }).(pulumi.IntOutput)
}

// Statistical content and methods.
func (o BotAnalyseProtectRuleOutput) StatisticalType() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.StatisticalType }).(pulumi.IntOutput)
}

// The total number of statistical protection rules under the current domain name.
func (o BotAnalyseProtectRuleOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v *BotAnalyseProtectRule) pulumi.IntOutput { return v.TotalCount }).(pulumi.IntOutput)
}

type BotAnalyseProtectRuleArrayOutput struct{ *pulumi.OutputState }

func (BotAnalyseProtectRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BotAnalyseProtectRule)(nil)).Elem()
}

func (o BotAnalyseProtectRuleArrayOutput) ToBotAnalyseProtectRuleArrayOutput() BotAnalyseProtectRuleArrayOutput {
	return o
}

func (o BotAnalyseProtectRuleArrayOutput) ToBotAnalyseProtectRuleArrayOutputWithContext(ctx context.Context) BotAnalyseProtectRuleArrayOutput {
	return o
}

func (o BotAnalyseProtectRuleArrayOutput) Index(i pulumi.IntInput) BotAnalyseProtectRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BotAnalyseProtectRule {
		return vs[0].([]*BotAnalyseProtectRule)[vs[1].(int)]
	}).(BotAnalyseProtectRuleOutput)
}

type BotAnalyseProtectRuleMapOutput struct{ *pulumi.OutputState }

func (BotAnalyseProtectRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BotAnalyseProtectRule)(nil)).Elem()
}

func (o BotAnalyseProtectRuleMapOutput) ToBotAnalyseProtectRuleMapOutput() BotAnalyseProtectRuleMapOutput {
	return o
}

func (o BotAnalyseProtectRuleMapOutput) ToBotAnalyseProtectRuleMapOutputWithContext(ctx context.Context) BotAnalyseProtectRuleMapOutput {
	return o
}

func (o BotAnalyseProtectRuleMapOutput) MapIndex(k pulumi.StringInput) BotAnalyseProtectRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BotAnalyseProtectRule {
		return vs[0].(map[string]*BotAnalyseProtectRule)[vs[1].(string)]
	}).(BotAnalyseProtectRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BotAnalyseProtectRuleInput)(nil)).Elem(), &BotAnalyseProtectRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*BotAnalyseProtectRuleArrayInput)(nil)).Elem(), BotAnalyseProtectRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BotAnalyseProtectRuleMapInput)(nil)).Elem(), BotAnalyseProtectRuleMap{})
	pulumi.RegisterOutputType(BotAnalyseProtectRuleOutput{})
	pulumi.RegisterOutputType(BotAnalyseProtectRuleArrayOutput{})
	pulumi.RegisterOutputType(BotAnalyseProtectRuleMapOutput{})
}
