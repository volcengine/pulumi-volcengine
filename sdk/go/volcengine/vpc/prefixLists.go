// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of vpc prefix lists
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooPrefixList, err := vpc.NewPrefixList(ctx, "fooPrefixList", &vpc.PrefixListArgs{
//				PrefixListName: pulumi.String("acc-test-prefix"),
//				MaxEntries:     pulumi.Int(3),
//				Description:    pulumi.String("acc test description"),
//				IpVersion:      pulumi.String("IPv4"),
//				PrefixListEntries: vpc.PrefixListPrefixListEntryArray{
//					&vpc.PrefixListPrefixListEntryArgs{
//						Cidr:        pulumi.String("192.168.4.0/28"),
//						Description: pulumi.String("acc-test-1"),
//					},
//					&vpc.PrefixListPrefixListEntryArgs{
//						Cidr:        pulumi.String("192.168.5.0/28"),
//						Description: pulumi.String("acc-test-2"),
//					},
//				},
//				Tags: vpc.PrefixListTagArray{
//					&vpc.PrefixListTagArgs{
//						Key:   pulumi.String("tf-key1"),
//						Value: pulumi.String("tf-value1"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_ = vpc.PrefixListsOutput(ctx, vpc.PrefixListsOutputArgs{
//				Ids: pulumi.StringArray{
//					fooPrefixList.ID(),
//				},
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func PrefixLists(ctx *pulumi.Context, args *PrefixListsArgs, opts ...pulumi.InvokeOption) (*PrefixListsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv PrefixListsResult
	err := ctx.Invoke("volcengine:vpc/prefixLists:PrefixLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking PrefixLists.
type PrefixListsArgs struct {
	// A list of prefix list ids.
	Ids []string `pulumi:"ids"`
	// IP version of prefix list.
	IpVersion *string `pulumi:"ipVersion"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// A Name of prefix list.
	PrefixListName *string `pulumi:"prefixListName"`
	// List of tag filters.
	TagFilters []PrefixListsTagFilter `pulumi:"tagFilters"`
}

// A collection of values returned by PrefixLists.
type PrefixListsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The ip version of the prefix list.
	IpVersion  *string `pulumi:"ipVersion"`
	OutputFile *string `pulumi:"outputFile"`
	// The prefix list name.
	PrefixListName *string `pulumi:"prefixListName"`
	// The collection of query.
	PrefixLists []PrefixListsPrefixList `pulumi:"prefixLists"`
	TagFilters  []PrefixListsTagFilter  `pulumi:"tagFilters"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func PrefixListsOutput(ctx *pulumi.Context, args PrefixListsOutputArgs, opts ...pulumi.InvokeOption) PrefixListsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (PrefixListsResult, error) {
			args := v.(PrefixListsArgs)
			r, err := PrefixLists(ctx, &args, opts...)
			var s PrefixListsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(PrefixListsResultOutput)
}

// A collection of arguments for invoking PrefixLists.
type PrefixListsOutputArgs struct {
	// A list of prefix list ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// IP version of prefix list.
	IpVersion pulumi.StringPtrInput `pulumi:"ipVersion"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A Name of prefix list.
	PrefixListName pulumi.StringPtrInput `pulumi:"prefixListName"`
	// List of tag filters.
	TagFilters PrefixListsTagFilterArrayInput `pulumi:"tagFilters"`
}

func (PrefixListsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrefixListsArgs)(nil)).Elem()
}

// A collection of values returned by PrefixLists.
type PrefixListsResultOutput struct{ *pulumi.OutputState }

func (PrefixListsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrefixListsResult)(nil)).Elem()
}

func (o PrefixListsResultOutput) ToPrefixListsResultOutput() PrefixListsResultOutput {
	return o
}

func (o PrefixListsResultOutput) ToPrefixListsResultOutputWithContext(ctx context.Context) PrefixListsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o PrefixListsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrefixListsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrefixListsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrefixListsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The ip version of the prefix list.
func (o PrefixListsResultOutput) IpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrefixListsResult) *string { return v.IpVersion }).(pulumi.StringPtrOutput)
}

func (o PrefixListsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrefixListsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The prefix list name.
func (o PrefixListsResultOutput) PrefixListName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrefixListsResult) *string { return v.PrefixListName }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o PrefixListsResultOutput) PrefixLists() PrefixListsPrefixListArrayOutput {
	return o.ApplyT(func(v PrefixListsResult) []PrefixListsPrefixList { return v.PrefixLists }).(PrefixListsPrefixListArrayOutput)
}

func (o PrefixListsResultOutput) TagFilters() PrefixListsTagFilterArrayOutput {
	return o.ApplyT(func(v PrefixListsResult) []PrefixListsTagFilter { return v.TagFilters }).(PrefixListsTagFilterArrayOutput)
}

// The total count of query.
func (o PrefixListsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v PrefixListsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(PrefixListsResultOutput{})
}