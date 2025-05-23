// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage rds mysql parameter template
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds_mysql.NewParameterTemplate(ctx, "foo", &rds_mysql.ParameterTemplateArgs{
//				TemplateDesc: pulumi.String("test"),
//				TemplateName: pulumi.String("test"),
//				TemplateParams: rds_mysql.ParameterTemplateTemplateParamArray{
//					&rds_mysql.ParameterTemplateTemplateParamArgs{
//						Name:         pulumi.String("auto_increment_increment"),
//						RunningValue: pulumi.String("1"),
//					},
//				},
//				TemplateType:        pulumi.String("Mysql"),
//				TemplateTypeVersion: pulumi.String("MySQL_8_0"),
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
// RdsMysqlParameterTemplate can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:rds_mysql/parameterTemplate:ParameterTemplate default resource_id
// ```
type ParameterTemplate struct {
	pulumi.CustomResourceState

	// Parameter template description.
	TemplateDesc pulumi.StringPtrOutput `pulumi:"templateDesc"`
	// Parameter template name.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
	// Parameters contained in the parameter template.
	TemplateParams ParameterTemplateTemplateParamArrayOutput `pulumi:"templateParams"`
	// Database type of parameter template. The default value is Mysql.
	TemplateType pulumi.StringOutput `pulumi:"templateType"`
	// Database version of parameter template. Value range:
	// MySQL_5_7: Default value. MySQL 5.7 version.
	// MySQL_8_0: MySQL 8.0 version.
	TemplateTypeVersion pulumi.StringOutput `pulumi:"templateTypeVersion"`
}

// NewParameterTemplate registers a new resource with the given unique name, arguments, and options.
func NewParameterTemplate(ctx *pulumi.Context,
	name string, args *ParameterTemplateArgs, opts ...pulumi.ResourceOption) (*ParameterTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	if args.TemplateParams == nil {
		return nil, errors.New("invalid value for required argument 'TemplateParams'")
	}
	if args.TemplateType == nil {
		return nil, errors.New("invalid value for required argument 'TemplateType'")
	}
	if args.TemplateTypeVersion == nil {
		return nil, errors.New("invalid value for required argument 'TemplateTypeVersion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ParameterTemplate
	err := ctx.RegisterResource("volcengine:rds_mysql/parameterTemplate:ParameterTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterTemplate gets an existing ParameterTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterTemplateState, opts ...pulumi.ResourceOption) (*ParameterTemplate, error) {
	var resource ParameterTemplate
	err := ctx.ReadResource("volcengine:rds_mysql/parameterTemplate:ParameterTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterTemplate resources.
type parameterTemplateState struct {
	// Parameter template description.
	TemplateDesc *string `pulumi:"templateDesc"`
	// Parameter template name.
	TemplateName *string `pulumi:"templateName"`
	// Parameters contained in the parameter template.
	TemplateParams []ParameterTemplateTemplateParam `pulumi:"templateParams"`
	// Database type of parameter template. The default value is Mysql.
	TemplateType *string `pulumi:"templateType"`
	// Database version of parameter template. Value range:
	// MySQL_5_7: Default value. MySQL 5.7 version.
	// MySQL_8_0: MySQL 8.0 version.
	TemplateTypeVersion *string `pulumi:"templateTypeVersion"`
}

type ParameterTemplateState struct {
	// Parameter template description.
	TemplateDesc pulumi.StringPtrInput
	// Parameter template name.
	TemplateName pulumi.StringPtrInput
	// Parameters contained in the parameter template.
	TemplateParams ParameterTemplateTemplateParamArrayInput
	// Database type of parameter template. The default value is Mysql.
	TemplateType pulumi.StringPtrInput
	// Database version of parameter template. Value range:
	// MySQL_5_7: Default value. MySQL 5.7 version.
	// MySQL_8_0: MySQL 8.0 version.
	TemplateTypeVersion pulumi.StringPtrInput
}

func (ParameterTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterTemplateState)(nil)).Elem()
}

type parameterTemplateArgs struct {
	// Parameter template description.
	TemplateDesc *string `pulumi:"templateDesc"`
	// Parameter template name.
	TemplateName string `pulumi:"templateName"`
	// Parameters contained in the parameter template.
	TemplateParams []ParameterTemplateTemplateParam `pulumi:"templateParams"`
	// Database type of parameter template. The default value is Mysql.
	TemplateType string `pulumi:"templateType"`
	// Database version of parameter template. Value range:
	// MySQL_5_7: Default value. MySQL 5.7 version.
	// MySQL_8_0: MySQL 8.0 version.
	TemplateTypeVersion string `pulumi:"templateTypeVersion"`
}

// The set of arguments for constructing a ParameterTemplate resource.
type ParameterTemplateArgs struct {
	// Parameter template description.
	TemplateDesc pulumi.StringPtrInput
	// Parameter template name.
	TemplateName pulumi.StringInput
	// Parameters contained in the parameter template.
	TemplateParams ParameterTemplateTemplateParamArrayInput
	// Database type of parameter template. The default value is Mysql.
	TemplateType pulumi.StringInput
	// Database version of parameter template. Value range:
	// MySQL_5_7: Default value. MySQL 5.7 version.
	// MySQL_8_0: MySQL 8.0 version.
	TemplateTypeVersion pulumi.StringInput
}

func (ParameterTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterTemplateArgs)(nil)).Elem()
}

type ParameterTemplateInput interface {
	pulumi.Input

	ToParameterTemplateOutput() ParameterTemplateOutput
	ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput
}

func (*ParameterTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterTemplate)(nil)).Elem()
}

func (i *ParameterTemplate) ToParameterTemplateOutput() ParameterTemplateOutput {
	return i.ToParameterTemplateOutputWithContext(context.Background())
}

func (i *ParameterTemplate) ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateOutput)
}

// ParameterTemplateArrayInput is an input type that accepts ParameterTemplateArray and ParameterTemplateArrayOutput values.
// You can construct a concrete instance of `ParameterTemplateArrayInput` via:
//
//	ParameterTemplateArray{ ParameterTemplateArgs{...} }
type ParameterTemplateArrayInput interface {
	pulumi.Input

	ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput
	ToParameterTemplateArrayOutputWithContext(context.Context) ParameterTemplateArrayOutput
}

type ParameterTemplateArray []ParameterTemplateInput

func (ParameterTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ParameterTemplate)(nil)).Elem()
}

func (i ParameterTemplateArray) ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput {
	return i.ToParameterTemplateArrayOutputWithContext(context.Background())
}

func (i ParameterTemplateArray) ToParameterTemplateArrayOutputWithContext(ctx context.Context) ParameterTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateArrayOutput)
}

// ParameterTemplateMapInput is an input type that accepts ParameterTemplateMap and ParameterTemplateMapOutput values.
// You can construct a concrete instance of `ParameterTemplateMapInput` via:
//
//	ParameterTemplateMap{ "key": ParameterTemplateArgs{...} }
type ParameterTemplateMapInput interface {
	pulumi.Input

	ToParameterTemplateMapOutput() ParameterTemplateMapOutput
	ToParameterTemplateMapOutputWithContext(context.Context) ParameterTemplateMapOutput
}

type ParameterTemplateMap map[string]ParameterTemplateInput

func (ParameterTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ParameterTemplate)(nil)).Elem()
}

func (i ParameterTemplateMap) ToParameterTemplateMapOutput() ParameterTemplateMapOutput {
	return i.ToParameterTemplateMapOutputWithContext(context.Background())
}

func (i ParameterTemplateMap) ToParameterTemplateMapOutputWithContext(ctx context.Context) ParameterTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateMapOutput)
}

type ParameterTemplateOutput struct{ *pulumi.OutputState }

func (ParameterTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateOutput) ToParameterTemplateOutput() ParameterTemplateOutput {
	return o
}

func (o ParameterTemplateOutput) ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput {
	return o
}

// Parameter template description.
func (o ParameterTemplateOutput) TemplateDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringPtrOutput { return v.TemplateDesc }).(pulumi.StringPtrOutput)
}

// Parameter template name.
func (o ParameterTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

// Parameters contained in the parameter template.
func (o ParameterTemplateOutput) TemplateParams() ParameterTemplateTemplateParamArrayOutput {
	return o.ApplyT(func(v *ParameterTemplate) ParameterTemplateTemplateParamArrayOutput { return v.TemplateParams }).(ParameterTemplateTemplateParamArrayOutput)
}

// Database type of parameter template. The default value is Mysql.
func (o ParameterTemplateOutput) TemplateType() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringOutput { return v.TemplateType }).(pulumi.StringOutput)
}

// Database version of parameter template. Value range:
// MySQL_5_7: Default value. MySQL 5.7 version.
// MySQL_8_0: MySQL 8.0 version.
func (o ParameterTemplateOutput) TemplateTypeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringOutput { return v.TemplateTypeVersion }).(pulumi.StringOutput)
}

type ParameterTemplateArrayOutput struct{ *pulumi.OutputState }

func (ParameterTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateArrayOutput) ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput {
	return o
}

func (o ParameterTemplateArrayOutput) ToParameterTemplateArrayOutputWithContext(ctx context.Context) ParameterTemplateArrayOutput {
	return o
}

func (o ParameterTemplateArrayOutput) Index(i pulumi.IntInput) ParameterTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ParameterTemplate {
		return vs[0].([]*ParameterTemplate)[vs[1].(int)]
	}).(ParameterTemplateOutput)
}

type ParameterTemplateMapOutput struct{ *pulumi.OutputState }

func (ParameterTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateMapOutput) ToParameterTemplateMapOutput() ParameterTemplateMapOutput {
	return o
}

func (o ParameterTemplateMapOutput) ToParameterTemplateMapOutputWithContext(ctx context.Context) ParameterTemplateMapOutput {
	return o
}

func (o ParameterTemplateMapOutput) MapIndex(k pulumi.StringInput) ParameterTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ParameterTemplate {
		return vs[0].(map[string]*ParameterTemplate)[vs[1].(string)]
	}).(ParameterTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateInput)(nil)).Elem(), &ParameterTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateArrayInput)(nil)).Elem(), ParameterTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateMapInput)(nil)).Elem(), ParameterTemplateMap{})
	pulumi.RegisterOutputType(ParameterTemplateOutput{})
	pulumi.RegisterOutputType(ParameterTemplateArrayOutput{})
	pulumi.RegisterOutputType(ParameterTemplateMapOutput{})
}
