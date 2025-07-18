// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apig

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage apig gateway
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/apig"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooZones, err := ecs.GetZones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			fooVpc, err := vpc.NewVpc(ctx, "fooVpc", &vpc.VpcArgs{
//				VpcName:   pulumi.String("acc-test-vpc"),
//				CidrBlock: pulumi.String("172.16.0.0/16"),
//			})
//			if err != nil {
//				return err
//			}
//			foo1, err := vpc.NewSubnet(ctx, "foo1", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			foo2, err := vpc.NewSubnet(ctx, "foo2", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.1.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[1].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = apig.NewApigGateway(ctx, "fooApigGateway", &apig.ApigGatewayArgs{
//				Type:        pulumi.String("standard"),
//				Comments:    pulumi.String("acc-test"),
//				ProjectName: pulumi.String("default"),
//				Tags: apig.ApigGatewayTagArray{
//					&apig.ApigGatewayTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//				NetworkSpec: &apig.ApigGatewayNetworkSpecArgs{
//					VpcId: fooVpc.ID(),
//					SubnetIds: pulumi.StringArray{
//						foo1.ID(),
//						foo2.ID(),
//					},
//				},
//				ResourceSpec: &apig.ApigGatewayResourceSpecArgs{
//					Replicas:                 pulumi.Int(2),
//					InstanceSpecCode:         pulumi.String("1c2g"),
//					ClbSpecCode:              pulumi.String("small_1"),
//					PublicNetworkBillingType: pulumi.String("bandwidth"),
//					PublicNetworkBandwidth:   pulumi.Int(1),
//					NetworkType: &apig.ApigGatewayResourceSpecNetworkTypeArgs{
//						EnablePublicNetwork:  pulumi.Bool(true),
//						EnablePrivateNetwork: pulumi.Bool(true),
//					},
//				},
//				LogSpec: &apig.ApigGatewayLogSpecArgs{
//					Enable:    pulumi.Bool(true),
//					ProjectId: pulumi.String("d3cb87c0-faeb-4074-b1ee-9bd747865a76"),
//					TopicId:   pulumi.String("d339482e-d86d-4bd8-a9bb-f270417f00a1"),
//				},
//				MonitorSpec: &apig.ApigGatewayMonitorSpecArgs{
//					Enable:      pulumi.Bool(true),
//					WorkspaceId: pulumi.String("4ed1caf3-279d-4c5f-8301-87ea38e92ffc"),
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
// ApigGateway can be imported using the id, e.g.
//
// ```sh
// $ pulumi import volcengine:apig/apigGateway:ApigGateway default resource_id
// ```
type ApigGateway struct {
	pulumi.CustomResourceState

	// The backend spec of the api gateway.
	BackendSpec ApigGatewayBackendSpecOutput `pulumi:"backendSpec"`
	// The comments of the api gateway.
	Comments pulumi.StringPtrOutput `pulumi:"comments"`
	// The create time of the api gateway.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The log spec of the api gateway.
	LogSpec ApigGatewayLogSpecOutput `pulumi:"logSpec"`
	// The error message of the api gateway.
	Message pulumi.StringOutput `pulumi:"message"`
	// The monitor spec of the api gateway.
	MonitorSpec ApigGatewayMonitorSpecOutput `pulumi:"monitorSpec"`
	// The name of the api gateway.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network spec of the api gateway.
	NetworkSpec ApigGatewayNetworkSpecOutput `pulumi:"networkSpec"`
	// The project name of the api gateway.
	ProjectName pulumi.StringOutput `pulumi:"projectName"`
	// The resource spec of the api gateway.
	ResourceSpec ApigGatewayResourceSpecOutput `pulumi:"resourceSpec"`
	// The status of the api gateway.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags.
	Tags ApigGatewayTagArrayOutput `pulumi:"tags"`
	// The type of the api gateway. Valid values: `standard`, `serverless`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The version of the api gateway.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewApigGateway registers a new resource with the given unique name, arguments, and options.
func NewApigGateway(ctx *pulumi.Context,
	name string, args *ApigGatewayArgs, opts ...pulumi.ResourceOption) (*ApigGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkSpec == nil {
		return nil, errors.New("invalid value for required argument 'NetworkSpec'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigGateway
	err := ctx.RegisterResource("volcengine:apig/apigGateway:ApigGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigGateway gets an existing ApigGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigGatewayState, opts ...pulumi.ResourceOption) (*ApigGateway, error) {
	var resource ApigGateway
	err := ctx.ReadResource("volcengine:apig/apigGateway:ApigGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigGateway resources.
type apigGatewayState struct {
	// The backend spec of the api gateway.
	BackendSpec *ApigGatewayBackendSpec `pulumi:"backendSpec"`
	// The comments of the api gateway.
	Comments *string `pulumi:"comments"`
	// The create time of the api gateway.
	CreateTime *string `pulumi:"createTime"`
	// The log spec of the api gateway.
	LogSpec *ApigGatewayLogSpec `pulumi:"logSpec"`
	// The error message of the api gateway.
	Message *string `pulumi:"message"`
	// The monitor spec of the api gateway.
	MonitorSpec *ApigGatewayMonitorSpec `pulumi:"monitorSpec"`
	// The name of the api gateway.
	Name *string `pulumi:"name"`
	// The network spec of the api gateway.
	NetworkSpec *ApigGatewayNetworkSpec `pulumi:"networkSpec"`
	// The project name of the api gateway.
	ProjectName *string `pulumi:"projectName"`
	// The resource spec of the api gateway.
	ResourceSpec *ApigGatewayResourceSpec `pulumi:"resourceSpec"`
	// The status of the api gateway.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []ApigGatewayTag `pulumi:"tags"`
	// The type of the api gateway. Valid values: `standard`, `serverless`.
	Type *string `pulumi:"type"`
	// The version of the api gateway.
	Version *string `pulumi:"version"`
}

type ApigGatewayState struct {
	// The backend spec of the api gateway.
	BackendSpec ApigGatewayBackendSpecPtrInput
	// The comments of the api gateway.
	Comments pulumi.StringPtrInput
	// The create time of the api gateway.
	CreateTime pulumi.StringPtrInput
	// The log spec of the api gateway.
	LogSpec ApigGatewayLogSpecPtrInput
	// The error message of the api gateway.
	Message pulumi.StringPtrInput
	// The monitor spec of the api gateway.
	MonitorSpec ApigGatewayMonitorSpecPtrInput
	// The name of the api gateway.
	Name pulumi.StringPtrInput
	// The network spec of the api gateway.
	NetworkSpec ApigGatewayNetworkSpecPtrInput
	// The project name of the api gateway.
	ProjectName pulumi.StringPtrInput
	// The resource spec of the api gateway.
	ResourceSpec ApigGatewayResourceSpecPtrInput
	// The status of the api gateway.
	Status pulumi.StringPtrInput
	// Tags.
	Tags ApigGatewayTagArrayInput
	// The type of the api gateway. Valid values: `standard`, `serverless`.
	Type pulumi.StringPtrInput
	// The version of the api gateway.
	Version pulumi.StringPtrInput
}

func (ApigGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigGatewayState)(nil)).Elem()
}

type apigGatewayArgs struct {
	// The backend spec of the api gateway.
	BackendSpec *ApigGatewayBackendSpec `pulumi:"backendSpec"`
	// The comments of the api gateway.
	Comments *string `pulumi:"comments"`
	// The log spec of the api gateway.
	LogSpec *ApigGatewayLogSpec `pulumi:"logSpec"`
	// The monitor spec of the api gateway.
	MonitorSpec *ApigGatewayMonitorSpec `pulumi:"monitorSpec"`
	// The name of the api gateway.
	Name *string `pulumi:"name"`
	// The network spec of the api gateway.
	NetworkSpec ApigGatewayNetworkSpec `pulumi:"networkSpec"`
	// The project name of the api gateway.
	ProjectName *string `pulumi:"projectName"`
	// The resource spec of the api gateway.
	ResourceSpec *ApigGatewayResourceSpec `pulumi:"resourceSpec"`
	// Tags.
	Tags []ApigGatewayTag `pulumi:"tags"`
	// The type of the api gateway. Valid values: `standard`, `serverless`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ApigGateway resource.
type ApigGatewayArgs struct {
	// The backend spec of the api gateway.
	BackendSpec ApigGatewayBackendSpecPtrInput
	// The comments of the api gateway.
	Comments pulumi.StringPtrInput
	// The log spec of the api gateway.
	LogSpec ApigGatewayLogSpecPtrInput
	// The monitor spec of the api gateway.
	MonitorSpec ApigGatewayMonitorSpecPtrInput
	// The name of the api gateway.
	Name pulumi.StringPtrInput
	// The network spec of the api gateway.
	NetworkSpec ApigGatewayNetworkSpecInput
	// The project name of the api gateway.
	ProjectName pulumi.StringPtrInput
	// The resource spec of the api gateway.
	ResourceSpec ApigGatewayResourceSpecPtrInput
	// Tags.
	Tags ApigGatewayTagArrayInput
	// The type of the api gateway. Valid values: `standard`, `serverless`.
	Type pulumi.StringPtrInput
}

func (ApigGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigGatewayArgs)(nil)).Elem()
}

type ApigGatewayInput interface {
	pulumi.Input

	ToApigGatewayOutput() ApigGatewayOutput
	ToApigGatewayOutputWithContext(ctx context.Context) ApigGatewayOutput
}

func (*ApigGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigGateway)(nil)).Elem()
}

func (i *ApigGateway) ToApigGatewayOutput() ApigGatewayOutput {
	return i.ToApigGatewayOutputWithContext(context.Background())
}

func (i *ApigGateway) ToApigGatewayOutputWithContext(ctx context.Context) ApigGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigGatewayOutput)
}

// ApigGatewayArrayInput is an input type that accepts ApigGatewayArray and ApigGatewayArrayOutput values.
// You can construct a concrete instance of `ApigGatewayArrayInput` via:
//
//	ApigGatewayArray{ ApigGatewayArgs{...} }
type ApigGatewayArrayInput interface {
	pulumi.Input

	ToApigGatewayArrayOutput() ApigGatewayArrayOutput
	ToApigGatewayArrayOutputWithContext(context.Context) ApigGatewayArrayOutput
}

type ApigGatewayArray []ApigGatewayInput

func (ApigGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigGateway)(nil)).Elem()
}

func (i ApigGatewayArray) ToApigGatewayArrayOutput() ApigGatewayArrayOutput {
	return i.ToApigGatewayArrayOutputWithContext(context.Background())
}

func (i ApigGatewayArray) ToApigGatewayArrayOutputWithContext(ctx context.Context) ApigGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigGatewayArrayOutput)
}

// ApigGatewayMapInput is an input type that accepts ApigGatewayMap and ApigGatewayMapOutput values.
// You can construct a concrete instance of `ApigGatewayMapInput` via:
//
//	ApigGatewayMap{ "key": ApigGatewayArgs{...} }
type ApigGatewayMapInput interface {
	pulumi.Input

	ToApigGatewayMapOutput() ApigGatewayMapOutput
	ToApigGatewayMapOutputWithContext(context.Context) ApigGatewayMapOutput
}

type ApigGatewayMap map[string]ApigGatewayInput

func (ApigGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigGateway)(nil)).Elem()
}

func (i ApigGatewayMap) ToApigGatewayMapOutput() ApigGatewayMapOutput {
	return i.ToApigGatewayMapOutputWithContext(context.Background())
}

func (i ApigGatewayMap) ToApigGatewayMapOutputWithContext(ctx context.Context) ApigGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigGatewayMapOutput)
}

type ApigGatewayOutput struct{ *pulumi.OutputState }

func (ApigGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigGateway)(nil)).Elem()
}

func (o ApigGatewayOutput) ToApigGatewayOutput() ApigGatewayOutput {
	return o
}

func (o ApigGatewayOutput) ToApigGatewayOutputWithContext(ctx context.Context) ApigGatewayOutput {
	return o
}

// The backend spec of the api gateway.
func (o ApigGatewayOutput) BackendSpec() ApigGatewayBackendSpecOutput {
	return o.ApplyT(func(v *ApigGateway) ApigGatewayBackendSpecOutput { return v.BackendSpec }).(ApigGatewayBackendSpecOutput)
}

// The comments of the api gateway.
func (o ApigGatewayOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigGateway) pulumi.StringPtrOutput { return v.Comments }).(pulumi.StringPtrOutput)
}

// The create time of the api gateway.
func (o ApigGatewayOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGateway) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The log spec of the api gateway.
func (o ApigGatewayOutput) LogSpec() ApigGatewayLogSpecOutput {
	return o.ApplyT(func(v *ApigGateway) ApigGatewayLogSpecOutput { return v.LogSpec }).(ApigGatewayLogSpecOutput)
}

// The error message of the api gateway.
func (o ApigGatewayOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGateway) pulumi.StringOutput { return v.Message }).(pulumi.StringOutput)
}

// The monitor spec of the api gateway.
func (o ApigGatewayOutput) MonitorSpec() ApigGatewayMonitorSpecOutput {
	return o.ApplyT(func(v *ApigGateway) ApigGatewayMonitorSpecOutput { return v.MonitorSpec }).(ApigGatewayMonitorSpecOutput)
}

// The name of the api gateway.
func (o ApigGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network spec of the api gateway.
func (o ApigGatewayOutput) NetworkSpec() ApigGatewayNetworkSpecOutput {
	return o.ApplyT(func(v *ApigGateway) ApigGatewayNetworkSpecOutput { return v.NetworkSpec }).(ApigGatewayNetworkSpecOutput)
}

// The project name of the api gateway.
func (o ApigGatewayOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGateway) pulumi.StringOutput { return v.ProjectName }).(pulumi.StringOutput)
}

// The resource spec of the api gateway.
func (o ApigGatewayOutput) ResourceSpec() ApigGatewayResourceSpecOutput {
	return o.ApplyT(func(v *ApigGateway) ApigGatewayResourceSpecOutput { return v.ResourceSpec }).(ApigGatewayResourceSpecOutput)
}

// The status of the api gateway.
func (o ApigGatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags.
func (o ApigGatewayOutput) Tags() ApigGatewayTagArrayOutput {
	return o.ApplyT(func(v *ApigGateway) ApigGatewayTagArrayOutput { return v.Tags }).(ApigGatewayTagArrayOutput)
}

// The type of the api gateway. Valid values: `standard`, `serverless`.
func (o ApigGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The version of the api gateway.
func (o ApigGatewayOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigGateway) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type ApigGatewayArrayOutput struct{ *pulumi.OutputState }

func (ApigGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigGateway)(nil)).Elem()
}

func (o ApigGatewayArrayOutput) ToApigGatewayArrayOutput() ApigGatewayArrayOutput {
	return o
}

func (o ApigGatewayArrayOutput) ToApigGatewayArrayOutputWithContext(ctx context.Context) ApigGatewayArrayOutput {
	return o
}

func (o ApigGatewayArrayOutput) Index(i pulumi.IntInput) ApigGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigGateway {
		return vs[0].([]*ApigGateway)[vs[1].(int)]
	}).(ApigGatewayOutput)
}

type ApigGatewayMapOutput struct{ *pulumi.OutputState }

func (ApigGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigGateway)(nil)).Elem()
}

func (o ApigGatewayMapOutput) ToApigGatewayMapOutput() ApigGatewayMapOutput {
	return o
}

func (o ApigGatewayMapOutput) ToApigGatewayMapOutputWithContext(ctx context.Context) ApigGatewayMapOutput {
	return o
}

func (o ApigGatewayMapOutput) MapIndex(k pulumi.StringInput) ApigGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigGateway {
		return vs[0].(map[string]*ApigGateway)[vs[1].(string)]
	}).(ApigGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigGatewayInput)(nil)).Elem(), &ApigGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigGatewayArrayInput)(nil)).Elem(), ApigGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigGatewayMapInput)(nil)).Elem(), ApigGatewayMap{})
	pulumi.RegisterOutputType(ApigGatewayOutput{})
	pulumi.RegisterOutputType(ApigGatewayArrayOutput{})
	pulumi.RegisterOutputType(ApigGatewayMapOutput{})
}
