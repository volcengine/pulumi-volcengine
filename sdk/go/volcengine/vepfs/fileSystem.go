// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vepfs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage vepfs file system
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vepfs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:   pulumi.String("acc-test-vpc"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String("cn-beijing-a"),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vepfs.NewFileSystem(ctx, "fooFileSystem", &vepfs.FileSystemArgs{
//				FileSystemName: pulumi.String("acc-test-file-system"),
//				SubnetId:       fooSubnet.ID(),
//				StoreType:      pulumi.String("Advance_100"),
//				Description:    pulumi.String("tf-test"),
//				Capacity:       pulumi.Int(12),
//				Project:        pulumi.String("default"),
//				EnableRestripe: pulumi.Bool(false),
//				Tags: vepfs.FileSystemTagArray{
//					&vepfs.FileSystemTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
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
// VepfsFileSystem can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:vepfs/fileSystem:FileSystem default resource_id
// ```
type FileSystem struct {
	pulumi.CustomResourceState

	// The id of the account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The bandwidth info of the vepfs file system.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The capacity of the vepfs file system.
	Capacity pulumi.IntOutput `pulumi:"capacity"`
	// The charge status of the vepfs file system.
	ChargeStatus pulumi.StringOutput `pulumi:"chargeStatus"`
	// The charge type of the vepfs file system.
	ChargeType pulumi.StringOutput `pulumi:"chargeType"`
	// The create time of the vepfs file system.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description info of the vepfs file system.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
	EnableRestripe pulumi.BoolPtrOutput `pulumi:"enableRestripe"`
	// The expire time of the vepfs file system.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The name of the vepfs file system.
	FileSystemName pulumi.StringOutput `pulumi:"fileSystemName"`
	// The type of the vepfs file system.
	FileSystemType pulumi.StringOutput `pulumi:"fileSystemType"`
	// The free time of the vepfs file system.
	FreeTime pulumi.StringOutput `pulumi:"freeTime"`
	// The last modify time of the vepfs file system.
	LastModifyTime pulumi.StringOutput `pulumi:"lastModifyTime"`
	// The project of the vepfs file system.
	Project pulumi.StringOutput `pulumi:"project"`
	// The protocol type of the vepfs file system.
	ProtocolType pulumi.StringOutput `pulumi:"protocolType"`
	// The id of the region.
	RegionId pulumi.StringOutput `pulumi:"regionId"`
	// The status of the vepfs file system.
	Status pulumi.StringOutput `pulumi:"status"`
	// The stop service time of the vepfs file system.
	StopServiceTime pulumi.StringOutput `pulumi:"stopServiceTime"`
	// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
	StoreType pulumi.StringOutput `pulumi:"storeType"`
	// The store type cn name of the vepfs file system.
	StoreTypeCn pulumi.StringOutput `pulumi:"storeTypeCn"`
	// The subnet id of the vepfs file system.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Tags.
	Tags FileSystemTagArrayOutput `pulumi:"tags"`
	// The version info of the vepfs file system.
	Version pulumi.StringOutput `pulumi:"version"`
	// The id of the zone.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
	// The name of the zone.
	ZoneName pulumi.StringOutput `pulumi:"zoneName"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Capacity == nil {
		return nil, errors.New("invalid value for required argument 'Capacity'")
	}
	if args.FileSystemName == nil {
		return nil, errors.New("invalid value for required argument 'FileSystemName'")
	}
	if args.StoreType == nil {
		return nil, errors.New("invalid value for required argument 'StoreType'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FileSystem
	err := ctx.RegisterResource("volcengine:vepfs/fileSystem:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("volcengine:vepfs/fileSystem:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
	// The id of the account.
	AccountId *string `pulumi:"accountId"`
	// The bandwidth info of the vepfs file system.
	Bandwidth *int `pulumi:"bandwidth"`
	// The capacity of the vepfs file system.
	Capacity *int `pulumi:"capacity"`
	// The charge status of the vepfs file system.
	ChargeStatus *string `pulumi:"chargeStatus"`
	// The charge type of the vepfs file system.
	ChargeType *string `pulumi:"chargeType"`
	// The create time of the vepfs file system.
	CreateTime *string `pulumi:"createTime"`
	// The description info of the vepfs file system.
	Description *string `pulumi:"description"`
	// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
	EnableRestripe *bool `pulumi:"enableRestripe"`
	// The expire time of the vepfs file system.
	ExpireTime *string `pulumi:"expireTime"`
	// The name of the vepfs file system.
	FileSystemName *string `pulumi:"fileSystemName"`
	// The type of the vepfs file system.
	FileSystemType *string `pulumi:"fileSystemType"`
	// The free time of the vepfs file system.
	FreeTime *string `pulumi:"freeTime"`
	// The last modify time of the vepfs file system.
	LastModifyTime *string `pulumi:"lastModifyTime"`
	// The project of the vepfs file system.
	Project *string `pulumi:"project"`
	// The protocol type of the vepfs file system.
	ProtocolType *string `pulumi:"protocolType"`
	// The id of the region.
	RegionId *string `pulumi:"regionId"`
	// The status of the vepfs file system.
	Status *string `pulumi:"status"`
	// The stop service time of the vepfs file system.
	StopServiceTime *string `pulumi:"stopServiceTime"`
	// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
	StoreType *string `pulumi:"storeType"`
	// The store type cn name of the vepfs file system.
	StoreTypeCn *string `pulumi:"storeTypeCn"`
	// The subnet id of the vepfs file system.
	SubnetId *string `pulumi:"subnetId"`
	// Tags.
	Tags []FileSystemTag `pulumi:"tags"`
	// The version info of the vepfs file system.
	Version *string `pulumi:"version"`
	// The id of the zone.
	ZoneId *string `pulumi:"zoneId"`
	// The name of the zone.
	ZoneName *string `pulumi:"zoneName"`
}

type FileSystemState struct {
	// The id of the account.
	AccountId pulumi.StringPtrInput
	// The bandwidth info of the vepfs file system.
	Bandwidth pulumi.IntPtrInput
	// The capacity of the vepfs file system.
	Capacity pulumi.IntPtrInput
	// The charge status of the vepfs file system.
	ChargeStatus pulumi.StringPtrInput
	// The charge type of the vepfs file system.
	ChargeType pulumi.StringPtrInput
	// The create time of the vepfs file system.
	CreateTime pulumi.StringPtrInput
	// The description info of the vepfs file system.
	Description pulumi.StringPtrInput
	// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
	EnableRestripe pulumi.BoolPtrInput
	// The expire time of the vepfs file system.
	ExpireTime pulumi.StringPtrInput
	// The name of the vepfs file system.
	FileSystemName pulumi.StringPtrInput
	// The type of the vepfs file system.
	FileSystemType pulumi.StringPtrInput
	// The free time of the vepfs file system.
	FreeTime pulumi.StringPtrInput
	// The last modify time of the vepfs file system.
	LastModifyTime pulumi.StringPtrInput
	// The project of the vepfs file system.
	Project pulumi.StringPtrInput
	// The protocol type of the vepfs file system.
	ProtocolType pulumi.StringPtrInput
	// The id of the region.
	RegionId pulumi.StringPtrInput
	// The status of the vepfs file system.
	Status pulumi.StringPtrInput
	// The stop service time of the vepfs file system.
	StopServiceTime pulumi.StringPtrInput
	// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
	StoreType pulumi.StringPtrInput
	// The store type cn name of the vepfs file system.
	StoreTypeCn pulumi.StringPtrInput
	// The subnet id of the vepfs file system.
	SubnetId pulumi.StringPtrInput
	// Tags.
	Tags FileSystemTagArrayInput
	// The version info of the vepfs file system.
	Version pulumi.StringPtrInput
	// The id of the zone.
	ZoneId pulumi.StringPtrInput
	// The name of the zone.
	ZoneName pulumi.StringPtrInput
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	// The capacity of the vepfs file system.
	Capacity int `pulumi:"capacity"`
	// The description info of the vepfs file system.
	Description *string `pulumi:"description"`
	// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
	EnableRestripe *bool `pulumi:"enableRestripe"`
	// The name of the vepfs file system.
	FileSystemName string `pulumi:"fileSystemName"`
	// The project of the vepfs file system.
	Project *string `pulumi:"project"`
	// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
	StoreType string `pulumi:"storeType"`
	// The subnet id of the vepfs file system.
	SubnetId string `pulumi:"subnetId"`
	// Tags.
	Tags []FileSystemTag `pulumi:"tags"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	// The capacity of the vepfs file system.
	Capacity pulumi.IntInput
	// The description info of the vepfs file system.
	Description pulumi.StringPtrInput
	// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
	EnableRestripe pulumi.BoolPtrInput
	// The name of the vepfs file system.
	FileSystemName pulumi.StringInput
	// The project of the vepfs file system.
	Project pulumi.StringPtrInput
	// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
	StoreType pulumi.StringInput
	// The subnet id of the vepfs file system.
	SubnetId pulumi.StringInput
	// Tags.
	Tags FileSystemTagArrayInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}

type FileSystemInput interface {
	pulumi.Input

	ToFileSystemOutput() FileSystemOutput
	ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput
}

func (*FileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil)).Elem()
}

func (i *FileSystem) ToFileSystemOutput() FileSystemOutput {
	return i.ToFileSystemOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemOutput)
}

// FileSystemArrayInput is an input type that accepts FileSystemArray and FileSystemArrayOutput values.
// You can construct a concrete instance of `FileSystemArrayInput` via:
//
//	FileSystemArray{ FileSystemArgs{...} }
type FileSystemArrayInput interface {
	pulumi.Input

	ToFileSystemArrayOutput() FileSystemArrayOutput
	ToFileSystemArrayOutputWithContext(context.Context) FileSystemArrayOutput
}

type FileSystemArray []FileSystemInput

func (FileSystemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FileSystem)(nil)).Elem()
}

func (i FileSystemArray) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return i.ToFileSystemArrayOutputWithContext(context.Background())
}

func (i FileSystemArray) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemArrayOutput)
}

// FileSystemMapInput is an input type that accepts FileSystemMap and FileSystemMapOutput values.
// You can construct a concrete instance of `FileSystemMapInput` via:
//
//	FileSystemMap{ "key": FileSystemArgs{...} }
type FileSystemMapInput interface {
	pulumi.Input

	ToFileSystemMapOutput() FileSystemMapOutput
	ToFileSystemMapOutputWithContext(context.Context) FileSystemMapOutput
}

type FileSystemMap map[string]FileSystemInput

func (FileSystemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FileSystem)(nil)).Elem()
}

func (i FileSystemMap) ToFileSystemMapOutput() FileSystemMapOutput {
	return i.ToFileSystemMapOutputWithContext(context.Background())
}

func (i FileSystemMap) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemMapOutput)
}

type FileSystemOutput struct{ *pulumi.OutputState }

func (FileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil)).Elem()
}

func (o FileSystemOutput) ToFileSystemOutput() FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return o
}

// The id of the account.
func (o FileSystemOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The bandwidth info of the vepfs file system.
func (o FileSystemOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The capacity of the vepfs file system.
func (o FileSystemOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.IntOutput { return v.Capacity }).(pulumi.IntOutput)
}

// The charge status of the vepfs file system.
func (o FileSystemOutput) ChargeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.ChargeStatus }).(pulumi.StringOutput)
}

// The charge type of the vepfs file system.
func (o FileSystemOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.ChargeType }).(pulumi.StringOutput)
}

// The create time of the vepfs file system.
func (o FileSystemOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description info of the vepfs file system.
func (o FileSystemOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.
func (o FileSystemOutput) EnableRestripe() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.BoolPtrOutput { return v.EnableRestripe }).(pulumi.BoolPtrOutput)
}

// The expire time of the vepfs file system.
func (o FileSystemOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// The name of the vepfs file system.
func (o FileSystemOutput) FileSystemName() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.FileSystemName }).(pulumi.StringOutput)
}

// The type of the vepfs file system.
func (o FileSystemOutput) FileSystemType() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.FileSystemType }).(pulumi.StringOutput)
}

// The free time of the vepfs file system.
func (o FileSystemOutput) FreeTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.FreeTime }).(pulumi.StringOutput)
}

// The last modify time of the vepfs file system.
func (o FileSystemOutput) LastModifyTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.LastModifyTime }).(pulumi.StringOutput)
}

// The project of the vepfs file system.
func (o FileSystemOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The protocol type of the vepfs file system.
func (o FileSystemOutput) ProtocolType() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.ProtocolType }).(pulumi.StringOutput)
}

// The id of the region.
func (o FileSystemOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.RegionId }).(pulumi.StringOutput)
}

// The status of the vepfs file system.
func (o FileSystemOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The stop service time of the vepfs file system.
func (o FileSystemOutput) StopServiceTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.StopServiceTime }).(pulumi.StringOutput)
}

// The store type of the vepfs file system. Valid values: `Advance_100`, `Performance`, `Intelligent_Computing`.
func (o FileSystemOutput) StoreType() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.StoreType }).(pulumi.StringOutput)
}

// The store type cn name of the vepfs file system.
func (o FileSystemOutput) StoreTypeCn() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.StoreTypeCn }).(pulumi.StringOutput)
}

// The subnet id of the vepfs file system.
func (o FileSystemOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Tags.
func (o FileSystemOutput) Tags() FileSystemTagArrayOutput {
	return o.ApplyT(func(v *FileSystem) FileSystemTagArrayOutput { return v.Tags }).(FileSystemTagArrayOutput)
}

// The version info of the vepfs file system.
func (o FileSystemOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// The id of the zone.
func (o FileSystemOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

// The name of the zone.
func (o FileSystemOutput) ZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.ZoneName }).(pulumi.StringOutput)
}

type FileSystemArrayOutput struct{ *pulumi.OutputState }

func (FileSystemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FileSystem)(nil)).Elem()
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutput() FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) ToFileSystemArrayOutputWithContext(ctx context.Context) FileSystemArrayOutput {
	return o
}

func (o FileSystemArrayOutput) Index(i pulumi.IntInput) FileSystemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FileSystem {
		return vs[0].([]*FileSystem)[vs[1].(int)]
	}).(FileSystemOutput)
}

type FileSystemMapOutput struct{ *pulumi.OutputState }

func (FileSystemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FileSystem)(nil)).Elem()
}

func (o FileSystemMapOutput) ToFileSystemMapOutput() FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) ToFileSystemMapOutputWithContext(ctx context.Context) FileSystemMapOutput {
	return o
}

func (o FileSystemMapOutput) MapIndex(k pulumi.StringInput) FileSystemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FileSystem {
		return vs[0].(map[string]*FileSystem)[vs[1].(string)]
	}).(FileSystemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemInput)(nil)).Elem(), &FileSystem{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemArrayInput)(nil)).Elem(), FileSystemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemMapInput)(nil)).Elem(), FileSystemMap{})
	pulumi.RegisterOutputType(FileSystemOutput{})
	pulumi.RegisterOutputType(FileSystemArrayOutput{})
	pulumi.RegisterOutputType(FileSystemMapOutput{})
}
