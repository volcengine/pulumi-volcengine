// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cr repositories
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
//			_, err := cr.Repositories(ctx, &cr.RepositoriesArgs{
//				Names: []string{
//					"repo*",
//				},
//				Registry: "tf-1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Repositories(ctx *pulumi.Context, args *RepositoriesArgs, opts ...pulumi.InvokeOption) (*RepositoriesResult, error) {
	var rv RepositoriesResult
	err := ctx.Invoke("volcengine:cr/repositories:Repositories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Repositories.
type RepositoriesArgs struct {
	// The list of instance access level.
	AccessLevels []string `pulumi:"accessLevels"`
	// The list of instance names.
	Names []string `pulumi:"names"`
	// The list of instance namespace.
	Namespaces []string `pulumi:"namespaces"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The CR instance name.
	Registry string `pulumi:"registry"`
}

// A collection of values returned by Repositories.
type RepositoriesResult struct {
	AccessLevels []string `pulumi:"accessLevels"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Names      []string `pulumi:"names"`
	Namespaces []string `pulumi:"namespaces"`
	OutputFile *string  `pulumi:"outputFile"`
	Registry   string   `pulumi:"registry"`
	// The collection of repository query.
	Repositories []RepositoriesRepository `pulumi:"repositories"`
	// The total count of instance query.
	TotalCount int `pulumi:"totalCount"`
}

func RepositoriesOutput(ctx *pulumi.Context, args RepositoriesOutputArgs, opts ...pulumi.InvokeOption) RepositoriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (RepositoriesResult, error) {
			args := v.(RepositoriesArgs)
			r, err := Repositories(ctx, &args, opts...)
			var s RepositoriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(RepositoriesResultOutput)
}

// A collection of arguments for invoking Repositories.
type RepositoriesOutputArgs struct {
	// The list of instance access level.
	AccessLevels pulumi.StringArrayInput `pulumi:"accessLevels"`
	// The list of instance names.
	Names pulumi.StringArrayInput `pulumi:"names"`
	// The list of instance namespace.
	Namespaces pulumi.StringArrayInput `pulumi:"namespaces"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The CR instance name.
	Registry pulumi.StringInput `pulumi:"registry"`
}

func (RepositoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoriesArgs)(nil)).Elem()
}

// A collection of values returned by Repositories.
type RepositoriesResultOutput struct{ *pulumi.OutputState }

func (RepositoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoriesResult)(nil)).Elem()
}

func (o RepositoriesResultOutput) ToRepositoriesResultOutput() RepositoriesResultOutput {
	return o
}

func (o RepositoriesResultOutput) ToRepositoriesResultOutputWithContext(ctx context.Context) RepositoriesResultOutput {
	return o
}

func (o RepositoriesResultOutput) AccessLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RepositoriesResult) []string { return v.AccessLevels }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o RepositoriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RepositoriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o RepositoriesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RepositoriesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o RepositoriesResultOutput) Namespaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RepositoriesResult) []string { return v.Namespaces }).(pulumi.StringArrayOutput)
}

func (o RepositoriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o RepositoriesResultOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v RepositoriesResult) string { return v.Registry }).(pulumi.StringOutput)
}

// The collection of repository query.
func (o RepositoriesResultOutput) Repositories() RepositoriesRepositoryArrayOutput {
	return o.ApplyT(func(v RepositoriesResult) []RepositoriesRepository { return v.Repositories }).(RepositoriesRepositoryArrayOutput)
}

// The total count of instance query.
func (o RepositoriesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v RepositoriesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(RepositoriesResultOutput{})
}