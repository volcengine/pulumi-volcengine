// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage image import
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewImageImport(ctx, "foo", &ecs.ImageImportArgs{
//				BootMode:    pulumi.String("UEFI"),
//				Description: pulumi.String("acc-test"),
//				ImageName:   pulumi.String("acc-test-image"),
//				Platform:    pulumi.String("CentOS"),
//				ProjectName: pulumi.String("default"),
//				Tags: ecs.ImageImportTagArray{
//					&ecs.ImageImportTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//				Url: pulumi.String("https://*****_system.qcow2"),
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
// ImageImport can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:ecs/imageImport:ImageImport default resource_id
// ```
type ImageImport struct {
	pulumi.CustomResourceState

	// The architecture of the custom image. Valid values: `amd64`, `arm64`.
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
	BootMode pulumi.StringOutput `pulumi:"bootMode"`
	// The create time of Image.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the custom image.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the custom image.
	ImageName pulumi.StringOutput `pulumi:"imageName"`
	// Whether the Image support cloud-init.
	IsSupportCloudInit pulumi.BoolOutput `pulumi:"isSupportCloudInit"`
	// The license type of the custom image. Valid values: `VolcanoEngine`.
	LicenseType pulumi.StringOutput `pulumi:"licenseType"`
	// The name of Image operating system.
	OsName pulumi.StringOutput `pulumi:"osName"`
	// The os type of the custom image. Valid values: `linux`, `Windows`.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
	Platform pulumi.StringOutput `pulumi:"platform"`
	// The platform version of the custom image.
	PlatformVersion pulumi.StringOutput `pulumi:"platformVersion"`
	// The project name of the custom image.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The share mode of Image.
	ShareStatus pulumi.StringOutput `pulumi:"shareStatus"`
	// The size(GiB) of Image.
	Size pulumi.IntOutput `pulumi:"size"`
	// The status of Image.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags.
	Tags ImageImportTagArrayOutput `pulumi:"tags"`
	// The update time of Image.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Url pulumi.StringOutput `pulumi:"url"`
	// The visibility of Image.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewImageImport registers a new resource with the given unique name, arguments, and options.
func NewImageImport(ctx *pulumi.Context,
	name string, args *ImageImportArgs, opts ...pulumi.ResourceOption) (*ImageImport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageName == nil {
		return nil, errors.New("invalid value for required argument 'ImageName'")
	}
	if args.Platform == nil {
		return nil, errors.New("invalid value for required argument 'Platform'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ImageImport
	err := ctx.RegisterResource("volcengine:ecs/imageImport:ImageImport", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageImport gets an existing ImageImport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageImport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageImportState, opts ...pulumi.ResourceOption) (*ImageImport, error) {
	var resource ImageImport
	err := ctx.ReadResource("volcengine:ecs/imageImport:ImageImport", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageImport resources.
type imageImportState struct {
	// The architecture of the custom image. Valid values: `amd64`, `arm64`.
	Architecture *string `pulumi:"architecture"`
	// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
	BootMode *string `pulumi:"bootMode"`
	// The create time of Image.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the custom image.
	Description *string `pulumi:"description"`
	// The name of the custom image.
	ImageName *string `pulumi:"imageName"`
	// Whether the Image support cloud-init.
	IsSupportCloudInit *bool `pulumi:"isSupportCloudInit"`
	// The license type of the custom image. Valid values: `VolcanoEngine`.
	LicenseType *string `pulumi:"licenseType"`
	// The name of Image operating system.
	OsName *string `pulumi:"osName"`
	// The os type of the custom image. Valid values: `linux`, `Windows`.
	OsType *string `pulumi:"osType"`
	// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
	Platform *string `pulumi:"platform"`
	// The platform version of the custom image.
	PlatformVersion *string `pulumi:"platformVersion"`
	// The project name of the custom image.
	ProjectName *string `pulumi:"projectName"`
	// The share mode of Image.
	ShareStatus *string `pulumi:"shareStatus"`
	// The size(GiB) of Image.
	Size *int `pulumi:"size"`
	// The status of Image.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []ImageImportTag `pulumi:"tags"`
	// The update time of Image.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Url *string `pulumi:"url"`
	// The visibility of Image.
	Visibility *string `pulumi:"visibility"`
}

type ImageImportState struct {
	// The architecture of the custom image. Valid values: `amd64`, `arm64`.
	Architecture pulumi.StringPtrInput
	// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
	BootMode pulumi.StringPtrInput
	// The create time of Image.
	CreatedAt pulumi.StringPtrInput
	// The description of the custom image.
	Description pulumi.StringPtrInput
	// The name of the custom image.
	ImageName pulumi.StringPtrInput
	// Whether the Image support cloud-init.
	IsSupportCloudInit pulumi.BoolPtrInput
	// The license type of the custom image. Valid values: `VolcanoEngine`.
	LicenseType pulumi.StringPtrInput
	// The name of Image operating system.
	OsName pulumi.StringPtrInput
	// The os type of the custom image. Valid values: `linux`, `Windows`.
	OsType pulumi.StringPtrInput
	// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
	Platform pulumi.StringPtrInput
	// The platform version of the custom image.
	PlatformVersion pulumi.StringPtrInput
	// The project name of the custom image.
	ProjectName pulumi.StringPtrInput
	// The share mode of Image.
	ShareStatus pulumi.StringPtrInput
	// The size(GiB) of Image.
	Size pulumi.IntPtrInput
	// The status of Image.
	Status pulumi.StringPtrInput
	// Tags.
	Tags ImageImportTagArrayInput
	// The update time of Image.
	UpdatedAt pulumi.StringPtrInput
	// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Url pulumi.StringPtrInput
	// The visibility of Image.
	Visibility pulumi.StringPtrInput
}

func (ImageImportState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageImportState)(nil)).Elem()
}

type imageImportArgs struct {
	// The architecture of the custom image. Valid values: `amd64`, `arm64`.
	Architecture *string `pulumi:"architecture"`
	// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
	BootMode *string `pulumi:"bootMode"`
	// The description of the custom image.
	Description *string `pulumi:"description"`
	// The name of the custom image.
	ImageName string `pulumi:"imageName"`
	// The license type of the custom image. Valid values: `VolcanoEngine`.
	LicenseType *string `pulumi:"licenseType"`
	// The os type of the custom image. Valid values: `linux`, `Windows`.
	OsType *string `pulumi:"osType"`
	// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
	Platform string `pulumi:"platform"`
	// The platform version of the custom image.
	PlatformVersion *string `pulumi:"platformVersion"`
	// The project name of the custom image.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []ImageImportTag `pulumi:"tags"`
	// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ImageImport resource.
type ImageImportArgs struct {
	// The architecture of the custom image. Valid values: `amd64`, `arm64`.
	Architecture pulumi.StringPtrInput
	// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
	BootMode pulumi.StringPtrInput
	// The description of the custom image.
	Description pulumi.StringPtrInput
	// The name of the custom image.
	ImageName pulumi.StringInput
	// The license type of the custom image. Valid values: `VolcanoEngine`.
	LicenseType pulumi.StringPtrInput
	// The os type of the custom image. Valid values: `linux`, `Windows`.
	OsType pulumi.StringPtrInput
	// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
	Platform pulumi.StringInput
	// The platform version of the custom image.
	PlatformVersion pulumi.StringPtrInput
	// The project name of the custom image.
	ProjectName pulumi.StringPtrInput
	// Tags.
	Tags ImageImportTagArrayInput
	// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Url pulumi.StringInput
}

func (ImageImportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageImportArgs)(nil)).Elem()
}

type ImageImportInput interface {
	pulumi.Input

	ToImageImportOutput() ImageImportOutput
	ToImageImportOutputWithContext(ctx context.Context) ImageImportOutput
}

func (*ImageImport) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageImport)(nil)).Elem()
}

func (i *ImageImport) ToImageImportOutput() ImageImportOutput {
	return i.ToImageImportOutputWithContext(context.Background())
}

func (i *ImageImport) ToImageImportOutputWithContext(ctx context.Context) ImageImportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageImportOutput)
}

// ImageImportArrayInput is an input type that accepts ImageImportArray and ImageImportArrayOutput values.
// You can construct a concrete instance of `ImageImportArrayInput` via:
//
//	ImageImportArray{ ImageImportArgs{...} }
type ImageImportArrayInput interface {
	pulumi.Input

	ToImageImportArrayOutput() ImageImportArrayOutput
	ToImageImportArrayOutputWithContext(context.Context) ImageImportArrayOutput
}

type ImageImportArray []ImageImportInput

func (ImageImportArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageImport)(nil)).Elem()
}

func (i ImageImportArray) ToImageImportArrayOutput() ImageImportArrayOutput {
	return i.ToImageImportArrayOutputWithContext(context.Background())
}

func (i ImageImportArray) ToImageImportArrayOutputWithContext(ctx context.Context) ImageImportArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageImportArrayOutput)
}

// ImageImportMapInput is an input type that accepts ImageImportMap and ImageImportMapOutput values.
// You can construct a concrete instance of `ImageImportMapInput` via:
//
//	ImageImportMap{ "key": ImageImportArgs{...} }
type ImageImportMapInput interface {
	pulumi.Input

	ToImageImportMapOutput() ImageImportMapOutput
	ToImageImportMapOutputWithContext(context.Context) ImageImportMapOutput
}

type ImageImportMap map[string]ImageImportInput

func (ImageImportMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageImport)(nil)).Elem()
}

func (i ImageImportMap) ToImageImportMapOutput() ImageImportMapOutput {
	return i.ToImageImportMapOutputWithContext(context.Background())
}

func (i ImageImportMap) ToImageImportMapOutputWithContext(ctx context.Context) ImageImportMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageImportMapOutput)
}

type ImageImportOutput struct{ *pulumi.OutputState }

func (ImageImportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageImport)(nil)).Elem()
}

func (o ImageImportOutput) ToImageImportOutput() ImageImportOutput {
	return o
}

func (o ImageImportOutput) ToImageImportOutputWithContext(ctx context.Context) ImageImportOutput {
	return o
}

// The architecture of the custom image. Valid values: `amd64`, `arm64`.
func (o ImageImportOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.Architecture }).(pulumi.StringOutput)
}

// The boot mode of the custom image. Valid values: `BIOS`, `UEFI`.
func (o ImageImportOutput) BootMode() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.BootMode }).(pulumi.StringOutput)
}

// The create time of Image.
func (o ImageImportOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the custom image.
func (o ImageImportOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the custom image.
func (o ImageImportOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.ImageName }).(pulumi.StringOutput)
}

// Whether the Image support cloud-init.
func (o ImageImportOutput) IsSupportCloudInit() pulumi.BoolOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.BoolOutput { return v.IsSupportCloudInit }).(pulumi.BoolOutput)
}

// The license type of the custom image. Valid values: `VolcanoEngine`.
func (o ImageImportOutput) LicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.LicenseType }).(pulumi.StringOutput)
}

// The name of Image operating system.
func (o ImageImportOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.OsName }).(pulumi.StringOutput)
}

// The os type of the custom image. Valid values: `linux`, `Windows`.
func (o ImageImportOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

// The platform of the custom image. Valid values: `CentOS`, `Debian`, `veLinux`, `Windows Server`, `Fedora`, `OpenSUSE`, `Ubuntu`, `Rocky Linux`, `AlmaLinux`.
func (o ImageImportOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

// The platform version of the custom image.
func (o ImageImportOutput) PlatformVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.PlatformVersion }).(pulumi.StringOutput)
}

// The project name of the custom image.
func (o ImageImportOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The share mode of Image.
func (o ImageImportOutput) ShareStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.ShareStatus }).(pulumi.StringOutput)
}

// The size(GiB) of Image.
func (o ImageImportOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The status of Image.
func (o ImageImportOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags.
func (o ImageImportOutput) Tags() ImageImportTagArrayOutput {
	return o.ApplyT(func(v *ImageImport) ImageImportTagArrayOutput { return v.Tags }).(ImageImportTagArrayOutput)
}

// The update time of Image.
func (o ImageImportOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The url of the custom image in tos bucket.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
func (o ImageImportOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The visibility of Image.
func (o ImageImportOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *ImageImport) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type ImageImportArrayOutput struct{ *pulumi.OutputState }

func (ImageImportArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageImport)(nil)).Elem()
}

func (o ImageImportArrayOutput) ToImageImportArrayOutput() ImageImportArrayOutput {
	return o
}

func (o ImageImportArrayOutput) ToImageImportArrayOutputWithContext(ctx context.Context) ImageImportArrayOutput {
	return o
}

func (o ImageImportArrayOutput) Index(i pulumi.IntInput) ImageImportOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageImport {
		return vs[0].([]*ImageImport)[vs[1].(int)]
	}).(ImageImportOutput)
}

type ImageImportMapOutput struct{ *pulumi.OutputState }

func (ImageImportMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageImport)(nil)).Elem()
}

func (o ImageImportMapOutput) ToImageImportMapOutput() ImageImportMapOutput {
	return o
}

func (o ImageImportMapOutput) ToImageImportMapOutputWithContext(ctx context.Context) ImageImportMapOutput {
	return o
}

func (o ImageImportMapOutput) MapIndex(k pulumi.StringInput) ImageImportOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageImport {
		return vs[0].(map[string]*ImageImport)[vs[1].(string)]
	}).(ImageImportOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageImportInput)(nil)).Elem(), &ImageImport{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageImportArrayInput)(nil)).Elem(), ImageImportArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageImportMapInput)(nil)).Elem(), ImageImportMap{})
	pulumi.RegisterOutputType(ImageImportOutput{})
	pulumi.RegisterOutputType(ImageImportArrayOutput{})
	pulumi.RegisterOutputType(ImageImportMapOutput{})
}