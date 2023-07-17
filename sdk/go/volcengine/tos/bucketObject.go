// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tos

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage tos object
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/tos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tos.NewBucketObject(ctx, "default", &tos.BucketObjectArgs{
//				AccountAcls: tos.BucketObjectAccountAclArray{
//					&tos.BucketObjectAccountAclArgs{
//						AccountId:  pulumi.String("1"),
//						Permission: pulumi.String("READ"),
//					},
//					&tos.BucketObjectAccountAclArgs{
//						AccountId:  pulumi.String("2001"),
//						Permission: pulumi.String("WRITE_ACP"),
//					},
//				},
//				BucketName: pulumi.String("test-xym-1"),
//				Encryption: pulumi.String("AES256"),
//				FilePath:   pulumi.String("/Users/bytedance/Work/Go/build/test.txt"),
//				ObjectName: pulumi.String("demo_xym"),
//				PublicAcl:  pulumi.String("private"),
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
// TOS Object can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:tos/bucketObject:BucketObject default bucketName:objectName
//
// ```
type BucketObject struct {
	pulumi.CustomResourceState

	// The user set of grant full control.
	AccountAcls BucketObjectAccountAclArrayOutput `pulumi:"accountAcls"`
	// The name of the bucket.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// The file md5 sum (32-bit hexadecimal string) for upload.
	ContentMd5 pulumi.StringPtrOutput `pulumi:"contentMd5"`
	// The content type of the object.
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// The flag of enable tos version.
	EnableVersion pulumi.BoolOutput `pulumi:"enableVersion"`
	// The encryption of the object.Valid value is AES256.
	Encryption pulumi.StringPtrOutput `pulumi:"encryption"`
	// The file path for upload.
	FilePath pulumi.StringOutput `pulumi:"filePath"`
	// The name of the object.
	ObjectName pulumi.StringOutput `pulumi:"objectName"`
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl pulumi.StringPtrOutput `pulumi:"publicAcl"`
	// The storage type of the object.Valid value is STANDARD|IA.
	StorageClass pulumi.StringPtrOutput `pulumi:"storageClass"`
	// The version ids of the object if exist.
	VersionIds pulumi.StringArrayOutput `pulumi:"versionIds"`
}

// NewBucketObject registers a new resource with the given unique name, arguments, and options.
func NewBucketObject(ctx *pulumi.Context,
	name string, args *BucketObjectArgs, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	if args.FilePath == nil {
		return nil, errors.New("invalid value for required argument 'FilePath'")
	}
	if args.ObjectName == nil {
		return nil, errors.New("invalid value for required argument 'ObjectName'")
	}
	var resource BucketObject
	err := ctx.RegisterResource("volcengine:tos/bucketObject:BucketObject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketObject gets an existing BucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketObjectState, opts ...pulumi.ResourceOption) (*BucketObject, error) {
	var resource BucketObject
	err := ctx.ReadResource("volcengine:tos/bucketObject:BucketObject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketObject resources.
type bucketObjectState struct {
	// The user set of grant full control.
	AccountAcls []BucketObjectAccountAcl `pulumi:"accountAcls"`
	// The name of the bucket.
	BucketName *string `pulumi:"bucketName"`
	// The file md5 sum (32-bit hexadecimal string) for upload.
	ContentMd5 *string `pulumi:"contentMd5"`
	// The content type of the object.
	ContentType *string `pulumi:"contentType"`
	// The flag of enable tos version.
	EnableVersion *bool `pulumi:"enableVersion"`
	// The encryption of the object.Valid value is AES256.
	Encryption *string `pulumi:"encryption"`
	// The file path for upload.
	FilePath *string `pulumi:"filePath"`
	// The name of the object.
	ObjectName *string `pulumi:"objectName"`
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl *string `pulumi:"publicAcl"`
	// The storage type of the object.Valid value is STANDARD|IA.
	StorageClass *string `pulumi:"storageClass"`
	// The version ids of the object if exist.
	VersionIds []string `pulumi:"versionIds"`
}

type BucketObjectState struct {
	// The user set of grant full control.
	AccountAcls BucketObjectAccountAclArrayInput
	// The name of the bucket.
	BucketName pulumi.StringPtrInput
	// The file md5 sum (32-bit hexadecimal string) for upload.
	ContentMd5 pulumi.StringPtrInput
	// The content type of the object.
	ContentType pulumi.StringPtrInput
	// The flag of enable tos version.
	EnableVersion pulumi.BoolPtrInput
	// The encryption of the object.Valid value is AES256.
	Encryption pulumi.StringPtrInput
	// The file path for upload.
	FilePath pulumi.StringPtrInput
	// The name of the object.
	ObjectName pulumi.StringPtrInput
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl pulumi.StringPtrInput
	// The storage type of the object.Valid value is STANDARD|IA.
	StorageClass pulumi.StringPtrInput
	// The version ids of the object if exist.
	VersionIds pulumi.StringArrayInput
}

func (BucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectState)(nil)).Elem()
}

type bucketObjectArgs struct {
	// The user set of grant full control.
	AccountAcls []BucketObjectAccountAcl `pulumi:"accountAcls"`
	// The name of the bucket.
	BucketName string `pulumi:"bucketName"`
	// The file md5 sum (32-bit hexadecimal string) for upload.
	ContentMd5 *string `pulumi:"contentMd5"`
	// The content type of the object.
	ContentType *string `pulumi:"contentType"`
	// The encryption of the object.Valid value is AES256.
	Encryption *string `pulumi:"encryption"`
	// The file path for upload.
	FilePath string `pulumi:"filePath"`
	// The name of the object.
	ObjectName string `pulumi:"objectName"`
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl *string `pulumi:"publicAcl"`
	// The storage type of the object.Valid value is STANDARD|IA.
	StorageClass *string `pulumi:"storageClass"`
}

// The set of arguments for constructing a BucketObject resource.
type BucketObjectArgs struct {
	// The user set of grant full control.
	AccountAcls BucketObjectAccountAclArrayInput
	// The name of the bucket.
	BucketName pulumi.StringInput
	// The file md5 sum (32-bit hexadecimal string) for upload.
	ContentMd5 pulumi.StringPtrInput
	// The content type of the object.
	ContentType pulumi.StringPtrInput
	// The encryption of the object.Valid value is AES256.
	Encryption pulumi.StringPtrInput
	// The file path for upload.
	FilePath pulumi.StringInput
	// The name of the object.
	ObjectName pulumi.StringInput
	// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
	PublicAcl pulumi.StringPtrInput
	// The storage type of the object.Valid value is STANDARD|IA.
	StorageClass pulumi.StringPtrInput
}

func (BucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketObjectArgs)(nil)).Elem()
}

type BucketObjectInput interface {
	pulumi.Input

	ToBucketObjectOutput() BucketObjectOutput
	ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput
}

func (*BucketObject) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObject)(nil)).Elem()
}

func (i *BucketObject) ToBucketObjectOutput() BucketObjectOutput {
	return i.ToBucketObjectOutputWithContext(context.Background())
}

func (i *BucketObject) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectOutput)
}

// BucketObjectArrayInput is an input type that accepts BucketObjectArray and BucketObjectArrayOutput values.
// You can construct a concrete instance of `BucketObjectArrayInput` via:
//
//	BucketObjectArray{ BucketObjectArgs{...} }
type BucketObjectArrayInput interface {
	pulumi.Input

	ToBucketObjectArrayOutput() BucketObjectArrayOutput
	ToBucketObjectArrayOutputWithContext(context.Context) BucketObjectArrayOutput
}

type BucketObjectArray []BucketObjectInput

func (BucketObjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObject)(nil)).Elem()
}

func (i BucketObjectArray) ToBucketObjectArrayOutput() BucketObjectArrayOutput {
	return i.ToBucketObjectArrayOutputWithContext(context.Background())
}

func (i BucketObjectArray) ToBucketObjectArrayOutputWithContext(ctx context.Context) BucketObjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectArrayOutput)
}

// BucketObjectMapInput is an input type that accepts BucketObjectMap and BucketObjectMapOutput values.
// You can construct a concrete instance of `BucketObjectMapInput` via:
//
//	BucketObjectMap{ "key": BucketObjectArgs{...} }
type BucketObjectMapInput interface {
	pulumi.Input

	ToBucketObjectMapOutput() BucketObjectMapOutput
	ToBucketObjectMapOutputWithContext(context.Context) BucketObjectMapOutput
}

type BucketObjectMap map[string]BucketObjectInput

func (BucketObjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObject)(nil)).Elem()
}

func (i BucketObjectMap) ToBucketObjectMapOutput() BucketObjectMapOutput {
	return i.ToBucketObjectMapOutputWithContext(context.Background())
}

func (i BucketObjectMap) ToBucketObjectMapOutputWithContext(ctx context.Context) BucketObjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketObjectMapOutput)
}

type BucketObjectOutput struct{ *pulumi.OutputState }

func (BucketObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketObject)(nil)).Elem()
}

func (o BucketObjectOutput) ToBucketObjectOutput() BucketObjectOutput {
	return o
}

func (o BucketObjectOutput) ToBucketObjectOutputWithContext(ctx context.Context) BucketObjectOutput {
	return o
}

// The user set of grant full control.
func (o BucketObjectOutput) AccountAcls() BucketObjectAccountAclArrayOutput {
	return o.ApplyT(func(v *BucketObject) BucketObjectAccountAclArrayOutput { return v.AccountAcls }).(BucketObjectAccountAclArrayOutput)
}

// The name of the bucket.
func (o BucketObjectOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.BucketName }).(pulumi.StringOutput)
}

// The file md5 sum (32-bit hexadecimal string) for upload.
func (o BucketObjectOutput) ContentMd5() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.ContentMd5 }).(pulumi.StringPtrOutput)
}

// The content type of the object.
func (o BucketObjectOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// The flag of enable tos version.
func (o BucketObjectOutput) EnableVersion() pulumi.BoolOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.BoolOutput { return v.EnableVersion }).(pulumi.BoolOutput)
}

// The encryption of the object.Valid value is AES256.
func (o BucketObjectOutput) Encryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.Encryption }).(pulumi.StringPtrOutput)
}

// The file path for upload.
func (o BucketObjectOutput) FilePath() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.FilePath }).(pulumi.StringOutput)
}

// The name of the object.
func (o BucketObjectOutput) ObjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringOutput { return v.ObjectName }).(pulumi.StringOutput)
}

// The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
func (o BucketObjectOutput) PublicAcl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.PublicAcl }).(pulumi.StringPtrOutput)
}

// The storage type of the object.Valid value is STANDARD|IA.
func (o BucketObjectOutput) StorageClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringPtrOutput { return v.StorageClass }).(pulumi.StringPtrOutput)
}

// The version ids of the object if exist.
func (o BucketObjectOutput) VersionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BucketObject) pulumi.StringArrayOutput { return v.VersionIds }).(pulumi.StringArrayOutput)
}

type BucketObjectArrayOutput struct{ *pulumi.OutputState }

func (BucketObjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BucketObject)(nil)).Elem()
}

func (o BucketObjectArrayOutput) ToBucketObjectArrayOutput() BucketObjectArrayOutput {
	return o
}

func (o BucketObjectArrayOutput) ToBucketObjectArrayOutputWithContext(ctx context.Context) BucketObjectArrayOutput {
	return o
}

func (o BucketObjectArrayOutput) Index(i pulumi.IntInput) BucketObjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BucketObject {
		return vs[0].([]*BucketObject)[vs[1].(int)]
	}).(BucketObjectOutput)
}

type BucketObjectMapOutput struct{ *pulumi.OutputState }

func (BucketObjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BucketObject)(nil)).Elem()
}

func (o BucketObjectMapOutput) ToBucketObjectMapOutput() BucketObjectMapOutput {
	return o
}

func (o BucketObjectMapOutput) ToBucketObjectMapOutputWithContext(ctx context.Context) BucketObjectMapOutput {
	return o
}

func (o BucketObjectMapOutput) MapIndex(k pulumi.StringInput) BucketObjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BucketObject {
		return vs[0].(map[string]*BucketObject)[vs[1].(string)]
	}).(BucketObjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectInput)(nil)).Elem(), &BucketObject{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectArrayInput)(nil)).Elem(), BucketObjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BucketObjectMapInput)(nil)).Elem(), BucketObjectMap{})
	pulumi.RegisterOutputType(BucketObjectOutput{})
	pulumi.RegisterOutputType(BucketObjectArrayOutput{})
	pulumi.RegisterOutputType(BucketObjectMapOutput{})
}