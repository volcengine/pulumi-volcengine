// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage cloudfs access
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cloudfs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudfs.NewAccess(ctx, "foo1", &cloudfs.AccessArgs{
//				FsName:          pulumi.String("tftest2"),
//				SecurityGroupId: pulumi.String("sg-rrv1klfg5s00v0x578mx14m"),
//				SubnetId:        pulumi.String("subnet-13fca1crr5d6o3n6nu46cyb5m"),
//				VpcRouteEnabled: pulumi.Bool(false),
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
// CloudFs Access can be imported using the FsName:AccessId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:cloudfs/access:Access default tfname:access-**rdgmedx3fow
//
// ```
type Access struct {
	pulumi.CustomResourceState

	// The account id of access.
	AccessAccountId pulumi.IntOutput `pulumi:"accessAccountId"`
	// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
	AccessIamRole pulumi.StringPtrOutput `pulumi:"accessIamRole"`
	// The id of access.
	AccessId pulumi.StringOutput `pulumi:"accessId"`
	// The service name of access.
	AccessServiceName pulumi.StringOutput `pulumi:"accessServiceName"`
	// The creation time.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The name of file system.
	FsName pulumi.StringOutput `pulumi:"fsName"`
	// Whether is default access.
	IsDefault pulumi.BoolOutput `pulumi:"isDefault"`
	// The id of security group.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// Status of access.
	Status pulumi.StringOutput `pulumi:"status"`
	// The id of subnet.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The id of vpc.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Whether enable all vpc route.
	VpcRouteEnabled pulumi.BoolPtrOutput `pulumi:"vpcRouteEnabled"`
}

// NewAccess registers a new resource with the given unique name, arguments, and options.
func NewAccess(ctx *pulumi.Context,
	name string, args *AccessArgs, opts ...pulumi.ResourceOption) (*Access, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FsName == nil {
		return nil, errors.New("invalid value for required argument 'FsName'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Access
	err := ctx.RegisterResource("volcengine:cloudfs/access:Access", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccess gets an existing Access resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessState, opts ...pulumi.ResourceOption) (*Access, error) {
	var resource Access
	err := ctx.ReadResource("volcengine:cloudfs/access:Access", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Access resources.
type accessState struct {
	// The account id of access.
	AccessAccountId *int `pulumi:"accessAccountId"`
	// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
	AccessIamRole *string `pulumi:"accessIamRole"`
	// The id of access.
	AccessId *string `pulumi:"accessId"`
	// The service name of access.
	AccessServiceName *string `pulumi:"accessServiceName"`
	// The creation time.
	CreatedTime *string `pulumi:"createdTime"`
	// The name of file system.
	FsName *string `pulumi:"fsName"`
	// Whether is default access.
	IsDefault *bool `pulumi:"isDefault"`
	// The id of security group.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Status of access.
	Status *string `pulumi:"status"`
	// The id of subnet.
	SubnetId *string `pulumi:"subnetId"`
	// The id of vpc.
	VpcId *string `pulumi:"vpcId"`
	// Whether enable all vpc route.
	VpcRouteEnabled *bool `pulumi:"vpcRouteEnabled"`
}

type AccessState struct {
	// The account id of access.
	AccessAccountId pulumi.IntPtrInput
	// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
	AccessIamRole pulumi.StringPtrInput
	// The id of access.
	AccessId pulumi.StringPtrInput
	// The service name of access.
	AccessServiceName pulumi.StringPtrInput
	// The creation time.
	CreatedTime pulumi.StringPtrInput
	// The name of file system.
	FsName pulumi.StringPtrInput
	// Whether is default access.
	IsDefault pulumi.BoolPtrInput
	// The id of security group.
	SecurityGroupId pulumi.StringPtrInput
	// Status of access.
	Status pulumi.StringPtrInput
	// The id of subnet.
	SubnetId pulumi.StringPtrInput
	// The id of vpc.
	VpcId pulumi.StringPtrInput
	// Whether enable all vpc route.
	VpcRouteEnabled pulumi.BoolPtrInput
}

func (AccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessState)(nil)).Elem()
}

type accessArgs struct {
	// The account id of access.
	AccessAccountId *int `pulumi:"accessAccountId"`
	// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
	AccessIamRole *string `pulumi:"accessIamRole"`
	// The name of file system.
	FsName string `pulumi:"fsName"`
	// The id of security group.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The id of subnet.
	SubnetId string `pulumi:"subnetId"`
	// Whether enable all vpc route.
	VpcRouteEnabled *bool `pulumi:"vpcRouteEnabled"`
}

// The set of arguments for constructing a Access resource.
type AccessArgs struct {
	// The account id of access.
	AccessAccountId pulumi.IntPtrInput
	// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
	AccessIamRole pulumi.StringPtrInput
	// The name of file system.
	FsName pulumi.StringInput
	// The id of security group.
	SecurityGroupId pulumi.StringInput
	// The id of subnet.
	SubnetId pulumi.StringInput
	// Whether enable all vpc route.
	VpcRouteEnabled pulumi.BoolPtrInput
}

func (AccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessArgs)(nil)).Elem()
}

type AccessInput interface {
	pulumi.Input

	ToAccessOutput() AccessOutput
	ToAccessOutputWithContext(ctx context.Context) AccessOutput
}

func (*Access) ElementType() reflect.Type {
	return reflect.TypeOf((**Access)(nil)).Elem()
}

func (i *Access) ToAccessOutput() AccessOutput {
	return i.ToAccessOutputWithContext(context.Background())
}

func (i *Access) ToAccessOutputWithContext(ctx context.Context) AccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessOutput)
}

// AccessArrayInput is an input type that accepts AccessArray and AccessArrayOutput values.
// You can construct a concrete instance of `AccessArrayInput` via:
//
//	AccessArray{ AccessArgs{...} }
type AccessArrayInput interface {
	pulumi.Input

	ToAccessArrayOutput() AccessArrayOutput
	ToAccessArrayOutputWithContext(context.Context) AccessArrayOutput
}

type AccessArray []AccessInput

func (AccessArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Access)(nil)).Elem()
}

func (i AccessArray) ToAccessArrayOutput() AccessArrayOutput {
	return i.ToAccessArrayOutputWithContext(context.Background())
}

func (i AccessArray) ToAccessArrayOutputWithContext(ctx context.Context) AccessArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessArrayOutput)
}

// AccessMapInput is an input type that accepts AccessMap and AccessMapOutput values.
// You can construct a concrete instance of `AccessMapInput` via:
//
//	AccessMap{ "key": AccessArgs{...} }
type AccessMapInput interface {
	pulumi.Input

	ToAccessMapOutput() AccessMapOutput
	ToAccessMapOutputWithContext(context.Context) AccessMapOutput
}

type AccessMap map[string]AccessInput

func (AccessMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Access)(nil)).Elem()
}

func (i AccessMap) ToAccessMapOutput() AccessMapOutput {
	return i.ToAccessMapOutputWithContext(context.Background())
}

func (i AccessMap) ToAccessMapOutputWithContext(ctx context.Context) AccessMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessMapOutput)
}

type AccessOutput struct{ *pulumi.OutputState }

func (AccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Access)(nil)).Elem()
}

func (o AccessOutput) ToAccessOutput() AccessOutput {
	return o
}

func (o AccessOutput) ToAccessOutputWithContext(ctx context.Context) AccessOutput {
	return o
}

// The account id of access.
func (o AccessOutput) AccessAccountId() pulumi.IntOutput {
	return o.ApplyT(func(v *Access) pulumi.IntOutput { return v.AccessAccountId }).(pulumi.IntOutput)
}

// The iam role of access. If the VPC of another account is attached, the other account needs to create a role with CFSCacheAccess permission, and enter the role name as a parameter.
func (o AccessOutput) AccessIamRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Access) pulumi.StringPtrOutput { return v.AccessIamRole }).(pulumi.StringPtrOutput)
}

// The id of access.
func (o AccessOutput) AccessId() pulumi.StringOutput {
	return o.ApplyT(func(v *Access) pulumi.StringOutput { return v.AccessId }).(pulumi.StringOutput)
}

// The service name of access.
func (o AccessOutput) AccessServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Access) pulumi.StringOutput { return v.AccessServiceName }).(pulumi.StringOutput)
}

// The creation time.
func (o AccessOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Access) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The name of file system.
func (o AccessOutput) FsName() pulumi.StringOutput {
	return o.ApplyT(func(v *Access) pulumi.StringOutput { return v.FsName }).(pulumi.StringOutput)
}

// Whether is default access.
func (o AccessOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v *Access) pulumi.BoolOutput { return v.IsDefault }).(pulumi.BoolOutput)
}

// The id of security group.
func (o AccessOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Access) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Status of access.
func (o AccessOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Access) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The id of subnet.
func (o AccessOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Access) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The id of vpc.
func (o AccessOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Access) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Whether enable all vpc route.
func (o AccessOutput) VpcRouteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Access) pulumi.BoolPtrOutput { return v.VpcRouteEnabled }).(pulumi.BoolPtrOutput)
}

type AccessArrayOutput struct{ *pulumi.OutputState }

func (AccessArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Access)(nil)).Elem()
}

func (o AccessArrayOutput) ToAccessArrayOutput() AccessArrayOutput {
	return o
}

func (o AccessArrayOutput) ToAccessArrayOutputWithContext(ctx context.Context) AccessArrayOutput {
	return o
}

func (o AccessArrayOutput) Index(i pulumi.IntInput) AccessOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Access {
		return vs[0].([]*Access)[vs[1].(int)]
	}).(AccessOutput)
}

type AccessMapOutput struct{ *pulumi.OutputState }

func (AccessMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Access)(nil)).Elem()
}

func (o AccessMapOutput) ToAccessMapOutput() AccessMapOutput {
	return o
}

func (o AccessMapOutput) ToAccessMapOutputWithContext(ctx context.Context) AccessMapOutput {
	return o
}

func (o AccessMapOutput) MapIndex(k pulumi.StringInput) AccessOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Access {
		return vs[0].(map[string]*Access)[vs[1].(string)]
	}).(AccessOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessInput)(nil)).Elem(), &Access{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessArrayInput)(nil)).Elem(), AccessArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessMapInput)(nil)).Elem(), AccessMap{})
	pulumi.RegisterOutputType(AccessOutput{})
	pulumi.RegisterOutputType(AccessArrayOutput{})
	pulumi.RegisterOutputType(AccessMapOutput{})
}