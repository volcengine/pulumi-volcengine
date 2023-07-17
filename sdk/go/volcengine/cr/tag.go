// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage cr tag
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/cr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cr.NewTag(ctx, "default", &cr.TagArgs{
//				Namespace:  pulumi.String("langyu"),
//				Registry:   pulumi.String("enterprise-1"),
//				Repository: pulumi.String("repo"),
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
// CR tags can be imported using the registry:namespace:repository:tag, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:cr/tag:Tag default cr-basic:namespace-1:repo-1:v1
//
// ```
type Tag struct {
	pulumi.CustomResourceState

	// The chart attribute,valid when tag type is Chart.
	ChartAttributes TagChartAttributeArrayOutput `pulumi:"chartAttributes"`
	// The digest of image.
	Digest pulumi.StringOutput `pulumi:"digest"`
	// The list of image attributes,valid when tag type is Image.
	ImageAttributes TagImageAttributeArrayOutput `pulumi:"imageAttributes"`
	// The name of OCI product.
	Name pulumi.StringOutput `pulumi:"name"`
	// The target namespace name.
	Namespace pulumi.StringOutput `pulumi:"namespace"`
	// The last push time of OCI product.
	PushTime pulumi.StringOutput `pulumi:"pushTime"`
	// The CrRegistry name.
	Registry pulumi.StringOutput `pulumi:"registry"`
	// The name of repository.
	Repository pulumi.StringOutput `pulumi:"repository"`
	// The size of OCI product.
	Size pulumi.IntOutput `pulumi:"size"`
	// The type of OCI product tag.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTag registers a new resource with the given unique name, arguments, and options.
func NewTag(ctx *pulumi.Context,
	name string, args *TagArgs, opts ...pulumi.ResourceOption) (*Tag, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	if args.Registry == nil {
		return nil, errors.New("invalid value for required argument 'Registry'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
	}
	var resource Tag
	err := ctx.RegisterResource("volcengine:cr/tag:Tag", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTag gets an existing Tag resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTag(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagState, opts ...pulumi.ResourceOption) (*Tag, error) {
	var resource Tag
	err := ctx.ReadResource("volcengine:cr/tag:Tag", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Tag resources.
type tagState struct {
	// The chart attribute,valid when tag type is Chart.
	ChartAttributes []TagChartAttribute `pulumi:"chartAttributes"`
	// The digest of image.
	Digest *string `pulumi:"digest"`
	// The list of image attributes,valid when tag type is Image.
	ImageAttributes []TagImageAttribute `pulumi:"imageAttributes"`
	// The name of OCI product.
	Name *string `pulumi:"name"`
	// The target namespace name.
	Namespace *string `pulumi:"namespace"`
	// The last push time of OCI product.
	PushTime *string `pulumi:"pushTime"`
	// The CrRegistry name.
	Registry *string `pulumi:"registry"`
	// The name of repository.
	Repository *string `pulumi:"repository"`
	// The size of OCI product.
	Size *int `pulumi:"size"`
	// The type of OCI product tag.
	Type *string `pulumi:"type"`
}

type TagState struct {
	// The chart attribute,valid when tag type is Chart.
	ChartAttributes TagChartAttributeArrayInput
	// The digest of image.
	Digest pulumi.StringPtrInput
	// The list of image attributes,valid when tag type is Image.
	ImageAttributes TagImageAttributeArrayInput
	// The name of OCI product.
	Name pulumi.StringPtrInput
	// The target namespace name.
	Namespace pulumi.StringPtrInput
	// The last push time of OCI product.
	PushTime pulumi.StringPtrInput
	// The CrRegistry name.
	Registry pulumi.StringPtrInput
	// The name of repository.
	Repository pulumi.StringPtrInput
	// The size of OCI product.
	Size pulumi.IntPtrInput
	// The type of OCI product tag.
	Type pulumi.StringPtrInput
}

func (TagState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagState)(nil)).Elem()
}

type tagArgs struct {
	// The name of OCI product.
	Name *string `pulumi:"name"`
	// The target namespace name.
	Namespace string `pulumi:"namespace"`
	// The CrRegistry name.
	Registry string `pulumi:"registry"`
	// The name of repository.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a Tag resource.
type TagArgs struct {
	// The name of OCI product.
	Name pulumi.StringPtrInput
	// The target namespace name.
	Namespace pulumi.StringInput
	// The CrRegistry name.
	Registry pulumi.StringInput
	// The name of repository.
	Repository pulumi.StringInput
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagArgs)(nil)).Elem()
}

type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(ctx context.Context) TagOutput
}

func (*Tag) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (i *Tag) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i *Tag) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}

// TagArrayInput is an input type that accepts TagArray and TagArrayOutput values.
// You can construct a concrete instance of `TagArrayInput` via:
//
//	TagArray{ TagArgs{...} }
type TagArrayInput interface {
	pulumi.Input

	ToTagArrayOutput() TagArrayOutput
	ToTagArrayOutputWithContext(context.Context) TagArrayOutput
}

type TagArray []TagInput

func (TagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tag)(nil)).Elem()
}

func (i TagArray) ToTagArrayOutput() TagArrayOutput {
	return i.ToTagArrayOutputWithContext(context.Background())
}

func (i TagArray) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagArrayOutput)
}

// TagMapInput is an input type that accepts TagMap and TagMapOutput values.
// You can construct a concrete instance of `TagMapInput` via:
//
//	TagMap{ "key": TagArgs{...} }
type TagMapInput interface {
	pulumi.Input

	ToTagMapOutput() TagMapOutput
	ToTagMapOutputWithContext(context.Context) TagMapOutput
}

type TagMap map[string]TagInput

func (TagMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tag)(nil)).Elem()
}

func (i TagMap) ToTagMapOutput() TagMapOutput {
	return i.ToTagMapOutputWithContext(context.Background())
}

func (i TagMap) ToTagMapOutputWithContext(ctx context.Context) TagMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagMapOutput)
}

type TagOutput struct{ *pulumi.OutputState }

func (TagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Tag)(nil)).Elem()
}

func (o TagOutput) ToTagOutput() TagOutput {
	return o
}

func (o TagOutput) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return o
}

// The chart attribute,valid when tag type is Chart.
func (o TagOutput) ChartAttributes() TagChartAttributeArrayOutput {
	return o.ApplyT(func(v *Tag) TagChartAttributeArrayOutput { return v.ChartAttributes }).(TagChartAttributeArrayOutput)
}

// The digest of image.
func (o TagOutput) Digest() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.Digest }).(pulumi.StringOutput)
}

// The list of image attributes,valid when tag type is Image.
func (o TagOutput) ImageAttributes() TagImageAttributeArrayOutput {
	return o.ApplyT(func(v *Tag) TagImageAttributeArrayOutput { return v.ImageAttributes }).(TagImageAttributeArrayOutput)
}

// The name of OCI product.
func (o TagOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The target namespace name.
func (o TagOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.Namespace }).(pulumi.StringOutput)
}

// The last push time of OCI product.
func (o TagOutput) PushTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.PushTime }).(pulumi.StringOutput)
}

// The CrRegistry name.
func (o TagOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.Registry }).(pulumi.StringOutput)
}

// The name of repository.
func (o TagOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.Repository }).(pulumi.StringOutput)
}

// The size of OCI product.
func (o TagOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Tag) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The type of OCI product tag.
func (o TagOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Tag) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type TagArrayOutput struct{ *pulumi.OutputState }

func (TagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Tag)(nil)).Elem()
}

func (o TagArrayOutput) ToTagArrayOutput() TagArrayOutput {
	return o
}

func (o TagArrayOutput) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return o
}

func (o TagArrayOutput) Index(i pulumi.IntInput) TagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Tag {
		return vs[0].([]*Tag)[vs[1].(int)]
	}).(TagOutput)
}

type TagMapOutput struct{ *pulumi.OutputState }

func (TagMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Tag)(nil)).Elem()
}

func (o TagMapOutput) ToTagMapOutput() TagMapOutput {
	return o
}

func (o TagMapOutput) ToTagMapOutputWithContext(ctx context.Context) TagMapOutput {
	return o
}

func (o TagMapOutput) MapIndex(k pulumi.StringInput) TagOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Tag {
		return vs[0].(map[string]*Tag)[vs[1].(string)]
	}).(TagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TagInput)(nil)).Elem(), &Tag{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagArrayInput)(nil)).Elem(), TagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TagMapInput)(nil)).Elem(), TagMap{})
	pulumi.RegisterOutputType(TagOutput{})
	pulumi.RegisterOutputType(TagArrayOutput{})
	pulumi.RegisterOutputType(TagMapOutput{})
}