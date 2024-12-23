// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

var _ = internal.GetEnvOrDefault

type AssumeRole struct {
	// The session name to use when making the AssumeRole call.
	AssumeRoleSessionName string `pulumi:"assumeRoleSessionName"`
	// The TRN of the role to assume.
	AssumeRoleTrn string `pulumi:"assumeRoleTrn"`
	// The duration of the session when making the AssumeRole call. Its value ranges from 900 to 43200(seconds), and default is 3600 seconds.
	DurationSeconds int `pulumi:"durationSeconds"`
	// A more restrictive policy when making the AssumeRole call.
	Policy *string `pulumi:"policy"`
}

// AssumeRoleInput is an input type that accepts AssumeRoleArgs and AssumeRoleOutput values.
// You can construct a concrete instance of `AssumeRoleInput` via:
//
//	AssumeRoleArgs{...}
type AssumeRoleInput interface {
	pulumi.Input

	ToAssumeRoleOutput() AssumeRoleOutput
	ToAssumeRoleOutputWithContext(context.Context) AssumeRoleOutput
}

type AssumeRoleArgs struct {
	// The session name to use when making the AssumeRole call.
	AssumeRoleSessionName pulumi.StringInput `pulumi:"assumeRoleSessionName"`
	// The TRN of the role to assume.
	AssumeRoleTrn pulumi.StringInput `pulumi:"assumeRoleTrn"`
	// The duration of the session when making the AssumeRole call. Its value ranges from 900 to 43200(seconds), and default is 3600 seconds.
	DurationSeconds pulumi.IntInput `pulumi:"durationSeconds"`
	// A more restrictive policy when making the AssumeRole call.
	Policy pulumi.StringPtrInput `pulumi:"policy"`
}

func (AssumeRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssumeRole)(nil)).Elem()
}

func (i AssumeRoleArgs) ToAssumeRoleOutput() AssumeRoleOutput {
	return i.ToAssumeRoleOutputWithContext(context.Background())
}

func (i AssumeRoleArgs) ToAssumeRoleOutputWithContext(ctx context.Context) AssumeRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssumeRoleOutput)
}

type AssumeRoleOutput struct{ *pulumi.OutputState }

func (AssumeRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssumeRole)(nil)).Elem()
}

func (o AssumeRoleOutput) ToAssumeRoleOutput() AssumeRoleOutput {
	return o
}

func (o AssumeRoleOutput) ToAssumeRoleOutputWithContext(ctx context.Context) AssumeRoleOutput {
	return o
}

// The session name to use when making the AssumeRole call.
func (o AssumeRoleOutput) AssumeRoleSessionName() pulumi.StringOutput {
	return o.ApplyT(func(v AssumeRole) string { return v.AssumeRoleSessionName }).(pulumi.StringOutput)
}

// The TRN of the role to assume.
func (o AssumeRoleOutput) AssumeRoleTrn() pulumi.StringOutput {
	return o.ApplyT(func(v AssumeRole) string { return v.AssumeRoleTrn }).(pulumi.StringOutput)
}

// The duration of the session when making the AssumeRole call. Its value ranges from 900 to 43200(seconds), and default is 3600 seconds.
func (o AssumeRoleOutput) DurationSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v AssumeRole) int { return v.DurationSeconds }).(pulumi.IntOutput)
}

// A more restrictive policy when making the AssumeRole call.
func (o AssumeRoleOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AssumeRole) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssumeRoleInput)(nil)).Elem(), AssumeRoleArgs{})
	pulumi.RegisterOutputType(AssumeRoleOutput{})
}
