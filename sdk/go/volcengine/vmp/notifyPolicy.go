// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vmp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage vmp notify policy
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vmp"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooContact, err := vmp.NewContact(ctx, "fooContact", &vmp.ContactArgs{
//				Email: pulumi.String("acctest1@tftest.com"),
//				Webhook: &vmp.ContactWebhookArgs{
//					Address: pulumi.String("https://www.acctest1.com"),
//				},
//				LarkBotWebhook: &vmp.ContactLarkBotWebhookArgs{
//					Address: pulumi.String("https://www.acctest1.com"),
//				},
//				DingTalkBotWebhook: &vmp.ContactDingTalkBotWebhookArgs{
//					Address: pulumi.String("https://www.dingacctest1.com"),
//					AtMobiles: pulumi.StringArray{
//						pulumi.String("18046891812"),
//					},
//				},
//				PhoneNumber: &vmp.ContactPhoneNumberArgs{
//					CountryCode: pulumi.String("+86"),
//					Number:      pulumi.String("18310101010"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			foo1Contact, err := vmp.NewContact(ctx, "foo1Contact", &vmp.ContactArgs{
//				Email: pulumi.String("acctest2@tftest.com"),
//				Webhook: &vmp.ContactWebhookArgs{
//					Address: pulumi.String("https://www.acctest2.com"),
//				},
//				LarkBotWebhook: &vmp.ContactLarkBotWebhookArgs{
//					Address: pulumi.String("https://www.acctest2.com"),
//				},
//				DingTalkBotWebhook: &vmp.ContactDingTalkBotWebhookArgs{
//					Address: pulumi.String("https://www.dingacctest2.com"),
//					AtMobiles: pulumi.StringArray{
//						pulumi.String("18046891813"),
//					},
//				},
//				PhoneNumber: &vmp.ContactPhoneNumberArgs{
//					CountryCode: pulumi.String("+86"),
//					Number:      pulumi.String("18310101011"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			fooContactGroup, err := vmp.NewContactGroup(ctx, "fooContactGroup", &vmp.ContactGroupArgs{
//				ContactIds: pulumi.StringArray{
//					fooContact.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			foo1ContactGroup, err := vmp.NewContactGroup(ctx, "foo1ContactGroup", &vmp.ContactGroupArgs{
//				ContactIds: pulumi.StringArray{
//					foo1Contact.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vmp.NewNotifyPolicy(ctx, "fooNotifyPolicy", &vmp.NotifyPolicyArgs{
//				Description: pulumi.String("acc-test-1"),
//				Levels: vmp.NotifyPolicyLevelArray{
//					&vmp.NotifyPolicyLevelArgs{
//						Level: pulumi.String("P1"),
//						ContactGroupIds: pulumi.StringArray{
//							fooContactGroup.ID(),
//						},
//						Channels: pulumi.StringArray{
//							pulumi.String("Email"),
//							pulumi.String("Webhook"),
//						},
//					},
//					&vmp.NotifyPolicyLevelArgs{
//						Level: pulumi.String("P0"),
//						ContactGroupIds: pulumi.StringArray{
//							foo1ContactGroup.ID(),
//						},
//						Channels: pulumi.StringArray{
//							pulumi.String("LarkBotWebhook"),
//						},
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
// VMP Notify Policy can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:vmp/notifyPolicy:NotifyPolicy default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
// ```
type NotifyPolicy struct {
	pulumi.CustomResourceState

	// The channel notify template for the alarm notification policy.
	ChannelNotifyTemplateIds pulumi.StringArrayOutput `pulumi:"channelNotifyTemplateIds"`
	// The description of the notify policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The levels of the notify policy.
	Levels NotifyPolicyLevelArrayOutput `pulumi:"levels"`
	// The name of the notify policy.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewNotifyPolicy registers a new resource with the given unique name, arguments, and options.
func NewNotifyPolicy(ctx *pulumi.Context,
	name string, args *NotifyPolicyArgs, opts ...pulumi.ResourceOption) (*NotifyPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Levels == nil {
		return nil, errors.New("invalid value for required argument 'Levels'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NotifyPolicy
	err := ctx.RegisterResource("volcengine:vmp/notifyPolicy:NotifyPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotifyPolicy gets an existing NotifyPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotifyPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotifyPolicyState, opts ...pulumi.ResourceOption) (*NotifyPolicy, error) {
	var resource NotifyPolicy
	err := ctx.ReadResource("volcengine:vmp/notifyPolicy:NotifyPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotifyPolicy resources.
type notifyPolicyState struct {
	// The channel notify template for the alarm notification policy.
	ChannelNotifyTemplateIds []string `pulumi:"channelNotifyTemplateIds"`
	// The description of the notify policy.
	Description *string `pulumi:"description"`
	// The levels of the notify policy.
	Levels []NotifyPolicyLevel `pulumi:"levels"`
	// The name of the notify policy.
	Name *string `pulumi:"name"`
}

type NotifyPolicyState struct {
	// The channel notify template for the alarm notification policy.
	ChannelNotifyTemplateIds pulumi.StringArrayInput
	// The description of the notify policy.
	Description pulumi.StringPtrInput
	// The levels of the notify policy.
	Levels NotifyPolicyLevelArrayInput
	// The name of the notify policy.
	Name pulumi.StringPtrInput
}

func (NotifyPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*notifyPolicyState)(nil)).Elem()
}

type notifyPolicyArgs struct {
	// The channel notify template for the alarm notification policy.
	ChannelNotifyTemplateIds []string `pulumi:"channelNotifyTemplateIds"`
	// The description of the notify policy.
	Description *string `pulumi:"description"`
	// The levels of the notify policy.
	Levels []NotifyPolicyLevel `pulumi:"levels"`
	// The name of the notify policy.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a NotifyPolicy resource.
type NotifyPolicyArgs struct {
	// The channel notify template for the alarm notification policy.
	ChannelNotifyTemplateIds pulumi.StringArrayInput
	// The description of the notify policy.
	Description pulumi.StringPtrInput
	// The levels of the notify policy.
	Levels NotifyPolicyLevelArrayInput
	// The name of the notify policy.
	Name pulumi.StringPtrInput
}

func (NotifyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notifyPolicyArgs)(nil)).Elem()
}

type NotifyPolicyInput interface {
	pulumi.Input

	ToNotifyPolicyOutput() NotifyPolicyOutput
	ToNotifyPolicyOutputWithContext(ctx context.Context) NotifyPolicyOutput
}

func (*NotifyPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**NotifyPolicy)(nil)).Elem()
}

func (i *NotifyPolicy) ToNotifyPolicyOutput() NotifyPolicyOutput {
	return i.ToNotifyPolicyOutputWithContext(context.Background())
}

func (i *NotifyPolicy) ToNotifyPolicyOutputWithContext(ctx context.Context) NotifyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifyPolicyOutput)
}

// NotifyPolicyArrayInput is an input type that accepts NotifyPolicyArray and NotifyPolicyArrayOutput values.
// You can construct a concrete instance of `NotifyPolicyArrayInput` via:
//
//	NotifyPolicyArray{ NotifyPolicyArgs{...} }
type NotifyPolicyArrayInput interface {
	pulumi.Input

	ToNotifyPolicyArrayOutput() NotifyPolicyArrayOutput
	ToNotifyPolicyArrayOutputWithContext(context.Context) NotifyPolicyArrayOutput
}

type NotifyPolicyArray []NotifyPolicyInput

func (NotifyPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotifyPolicy)(nil)).Elem()
}

func (i NotifyPolicyArray) ToNotifyPolicyArrayOutput() NotifyPolicyArrayOutput {
	return i.ToNotifyPolicyArrayOutputWithContext(context.Background())
}

func (i NotifyPolicyArray) ToNotifyPolicyArrayOutputWithContext(ctx context.Context) NotifyPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifyPolicyArrayOutput)
}

// NotifyPolicyMapInput is an input type that accepts NotifyPolicyMap and NotifyPolicyMapOutput values.
// You can construct a concrete instance of `NotifyPolicyMapInput` via:
//
//	NotifyPolicyMap{ "key": NotifyPolicyArgs{...} }
type NotifyPolicyMapInput interface {
	pulumi.Input

	ToNotifyPolicyMapOutput() NotifyPolicyMapOutput
	ToNotifyPolicyMapOutputWithContext(context.Context) NotifyPolicyMapOutput
}

type NotifyPolicyMap map[string]NotifyPolicyInput

func (NotifyPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotifyPolicy)(nil)).Elem()
}

func (i NotifyPolicyMap) ToNotifyPolicyMapOutput() NotifyPolicyMapOutput {
	return i.ToNotifyPolicyMapOutputWithContext(context.Background())
}

func (i NotifyPolicyMap) ToNotifyPolicyMapOutputWithContext(ctx context.Context) NotifyPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotifyPolicyMapOutput)
}

type NotifyPolicyOutput struct{ *pulumi.OutputState }

func (NotifyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotifyPolicy)(nil)).Elem()
}

func (o NotifyPolicyOutput) ToNotifyPolicyOutput() NotifyPolicyOutput {
	return o
}

func (o NotifyPolicyOutput) ToNotifyPolicyOutputWithContext(ctx context.Context) NotifyPolicyOutput {
	return o
}

// The channel notify template for the alarm notification policy.
func (o NotifyPolicyOutput) ChannelNotifyTemplateIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NotifyPolicy) pulumi.StringArrayOutput { return v.ChannelNotifyTemplateIds }).(pulumi.StringArrayOutput)
}

// The description of the notify policy.
func (o NotifyPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NotifyPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The levels of the notify policy.
func (o NotifyPolicyOutput) Levels() NotifyPolicyLevelArrayOutput {
	return o.ApplyT(func(v *NotifyPolicy) NotifyPolicyLevelArrayOutput { return v.Levels }).(NotifyPolicyLevelArrayOutput)
}

// The name of the notify policy.
func (o NotifyPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NotifyPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type NotifyPolicyArrayOutput struct{ *pulumi.OutputState }

func (NotifyPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NotifyPolicy)(nil)).Elem()
}

func (o NotifyPolicyArrayOutput) ToNotifyPolicyArrayOutput() NotifyPolicyArrayOutput {
	return o
}

func (o NotifyPolicyArrayOutput) ToNotifyPolicyArrayOutputWithContext(ctx context.Context) NotifyPolicyArrayOutput {
	return o
}

func (o NotifyPolicyArrayOutput) Index(i pulumi.IntInput) NotifyPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NotifyPolicy {
		return vs[0].([]*NotifyPolicy)[vs[1].(int)]
	}).(NotifyPolicyOutput)
}

type NotifyPolicyMapOutput struct{ *pulumi.OutputState }

func (NotifyPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NotifyPolicy)(nil)).Elem()
}

func (o NotifyPolicyMapOutput) ToNotifyPolicyMapOutput() NotifyPolicyMapOutput {
	return o
}

func (o NotifyPolicyMapOutput) ToNotifyPolicyMapOutputWithContext(ctx context.Context) NotifyPolicyMapOutput {
	return o
}

func (o NotifyPolicyMapOutput) MapIndex(k pulumi.StringInput) NotifyPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NotifyPolicy {
		return vs[0].(map[string]*NotifyPolicy)[vs[1].(string)]
	}).(NotifyPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotifyPolicyInput)(nil)).Elem(), &NotifyPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotifyPolicyArrayInput)(nil)).Elem(), NotifyPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotifyPolicyMapInput)(nil)).Elem(), NotifyPolicyMap{})
	pulumi.RegisterOutputType(NotifyPolicyOutput{})
	pulumi.RegisterOutputType(NotifyPolicyArrayOutput{})
	pulumi.RegisterOutputType(NotifyPolicyMapOutput{})
}
