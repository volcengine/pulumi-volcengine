// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Provides a resource to manage redis allow list associate
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/redis"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooAllowList, err := redis.NewAllowList(ctx, "fooAllowList", &redis.AllowListArgs{
//				AllowLists: pulumi.StringArray{
//					pulumi.String("192.168.0.0/24"),
//				},
//				AllowListName: pulumi.String("acc-test-allowlist"),
//			})
//			if err != nil {
//				return err
//			}
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
//			fooSubnet, err := vpc.NewSubnet(ctx, "fooSubnet", &vpc.SubnetArgs{
//				SubnetName: pulumi.String("acc-test-subnet"),
//				CidrBlock:  pulumi.String("172.16.0.0/24"),
//				ZoneId:     pulumi.String(fooZones.Zones[0].Id),
//				VpcId:      fooVpc.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			fooInstance, err := redis.NewInstance(ctx, "fooInstance", &redis.InstanceArgs{
//				ZoneIds: pulumi.StringArray{
//					pulumi.String(fooZones.Zones[0].Id),
//				},
//				InstanceName:       pulumi.String("acc-test-tf-redis"),
//				ShardedCluster:     pulumi.Int(1),
//				Password:           pulumi.String("1qaz!QAZ12"),
//				NodeNumber:         pulumi.Int(2),
//				ShardCapacity:      pulumi.Int(1024),
//				ShardNumber:        pulumi.Int(2),
//				EngineVersion:      pulumi.String("5.0"),
//				SubnetId:           fooSubnet.ID(),
//				DeletionProtection: pulumi.String("disabled"),
//				VpcAuthMode:        pulumi.String("close"),
//				ChargeType:         pulumi.String("PostPaid"),
//				Port:               pulumi.Int(6381),
//				ProjectName:        pulumi.String("default"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = redis.NewAllowListAssociate(ctx, "fooAllowListAssociate", &redis.AllowListAssociateArgs{
//				AllowListId: fooAllowList.ID(),
//				InstanceId:  fooInstance.ID(),
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
// Redis AllowList Association can be imported using the instanceId:allowListId, e.g.
//
// ```sh
// $ pulumi import volcengine:redis/allowListAssociate:AllowListAssociate default redis-asdljioeixxxx:acl-cn03wk541s55c376xxxx
// ```
type AllowListAssociate struct {
	pulumi.CustomResourceState

	// Id of allow list to associate.
	AllowListId pulumi.StringOutput `pulumi:"allowListId"`
	// Id of instance to associate.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewAllowListAssociate registers a new resource with the given unique name, arguments, and options.
func NewAllowListAssociate(ctx *pulumi.Context,
	name string, args *AllowListAssociateArgs, opts ...pulumi.ResourceOption) (*AllowListAssociate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowListId == nil {
		return nil, errors.New("invalid value for required argument 'AllowListId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AllowListAssociate
	err := ctx.RegisterResource("volcengine:redis/allowListAssociate:AllowListAssociate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAllowListAssociate gets an existing AllowListAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAllowListAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AllowListAssociateState, opts ...pulumi.ResourceOption) (*AllowListAssociate, error) {
	var resource AllowListAssociate
	err := ctx.ReadResource("volcengine:redis/allowListAssociate:AllowListAssociate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AllowListAssociate resources.
type allowListAssociateState struct {
	// Id of allow list to associate.
	AllowListId *string `pulumi:"allowListId"`
	// Id of instance to associate.
	InstanceId *string `pulumi:"instanceId"`
}

type AllowListAssociateState struct {
	// Id of allow list to associate.
	AllowListId pulumi.StringPtrInput
	// Id of instance to associate.
	InstanceId pulumi.StringPtrInput
}

func (AllowListAssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*allowListAssociateState)(nil)).Elem()
}

type allowListAssociateArgs struct {
	// Id of allow list to associate.
	AllowListId string `pulumi:"allowListId"`
	// Id of instance to associate.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a AllowListAssociate resource.
type AllowListAssociateArgs struct {
	// Id of allow list to associate.
	AllowListId pulumi.StringInput
	// Id of instance to associate.
	InstanceId pulumi.StringInput
}

func (AllowListAssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*allowListAssociateArgs)(nil)).Elem()
}

type AllowListAssociateInput interface {
	pulumi.Input

	ToAllowListAssociateOutput() AllowListAssociateOutput
	ToAllowListAssociateOutputWithContext(ctx context.Context) AllowListAssociateOutput
}

func (*AllowListAssociate) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowListAssociate)(nil)).Elem()
}

func (i *AllowListAssociate) ToAllowListAssociateOutput() AllowListAssociateOutput {
	return i.ToAllowListAssociateOutputWithContext(context.Background())
}

func (i *AllowListAssociate) ToAllowListAssociateOutputWithContext(ctx context.Context) AllowListAssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowListAssociateOutput)
}

// AllowListAssociateArrayInput is an input type that accepts AllowListAssociateArray and AllowListAssociateArrayOutput values.
// You can construct a concrete instance of `AllowListAssociateArrayInput` via:
//
//	AllowListAssociateArray{ AllowListAssociateArgs{...} }
type AllowListAssociateArrayInput interface {
	pulumi.Input

	ToAllowListAssociateArrayOutput() AllowListAssociateArrayOutput
	ToAllowListAssociateArrayOutputWithContext(context.Context) AllowListAssociateArrayOutput
}

type AllowListAssociateArray []AllowListAssociateInput

func (AllowListAssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AllowListAssociate)(nil)).Elem()
}

func (i AllowListAssociateArray) ToAllowListAssociateArrayOutput() AllowListAssociateArrayOutput {
	return i.ToAllowListAssociateArrayOutputWithContext(context.Background())
}

func (i AllowListAssociateArray) ToAllowListAssociateArrayOutputWithContext(ctx context.Context) AllowListAssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowListAssociateArrayOutput)
}

// AllowListAssociateMapInput is an input type that accepts AllowListAssociateMap and AllowListAssociateMapOutput values.
// You can construct a concrete instance of `AllowListAssociateMapInput` via:
//
//	AllowListAssociateMap{ "key": AllowListAssociateArgs{...} }
type AllowListAssociateMapInput interface {
	pulumi.Input

	ToAllowListAssociateMapOutput() AllowListAssociateMapOutput
	ToAllowListAssociateMapOutputWithContext(context.Context) AllowListAssociateMapOutput
}

type AllowListAssociateMap map[string]AllowListAssociateInput

func (AllowListAssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AllowListAssociate)(nil)).Elem()
}

func (i AllowListAssociateMap) ToAllowListAssociateMapOutput() AllowListAssociateMapOutput {
	return i.ToAllowListAssociateMapOutputWithContext(context.Background())
}

func (i AllowListAssociateMap) ToAllowListAssociateMapOutputWithContext(ctx context.Context) AllowListAssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowListAssociateMapOutput)
}

type AllowListAssociateOutput struct{ *pulumi.OutputState }

func (AllowListAssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowListAssociate)(nil)).Elem()
}

func (o AllowListAssociateOutput) ToAllowListAssociateOutput() AllowListAssociateOutput {
	return o
}

func (o AllowListAssociateOutput) ToAllowListAssociateOutputWithContext(ctx context.Context) AllowListAssociateOutput {
	return o
}

// Id of allow list to associate.
func (o AllowListAssociateOutput) AllowListId() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowListAssociate) pulumi.StringOutput { return v.AllowListId }).(pulumi.StringOutput)
}

// Id of instance to associate.
func (o AllowListAssociateOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowListAssociate) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type AllowListAssociateArrayOutput struct{ *pulumi.OutputState }

func (AllowListAssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AllowListAssociate)(nil)).Elem()
}

func (o AllowListAssociateArrayOutput) ToAllowListAssociateArrayOutput() AllowListAssociateArrayOutput {
	return o
}

func (o AllowListAssociateArrayOutput) ToAllowListAssociateArrayOutputWithContext(ctx context.Context) AllowListAssociateArrayOutput {
	return o
}

func (o AllowListAssociateArrayOutput) Index(i pulumi.IntInput) AllowListAssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AllowListAssociate {
		return vs[0].([]*AllowListAssociate)[vs[1].(int)]
	}).(AllowListAssociateOutput)
}

type AllowListAssociateMapOutput struct{ *pulumi.OutputState }

func (AllowListAssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AllowListAssociate)(nil)).Elem()
}

func (o AllowListAssociateMapOutput) ToAllowListAssociateMapOutput() AllowListAssociateMapOutput {
	return o
}

func (o AllowListAssociateMapOutput) ToAllowListAssociateMapOutputWithContext(ctx context.Context) AllowListAssociateMapOutput {
	return o
}

func (o AllowListAssociateMapOutput) MapIndex(k pulumi.StringInput) AllowListAssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AllowListAssociate {
		return vs[0].(map[string]*AllowListAssociate)[vs[1].(string)]
	}).(AllowListAssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AllowListAssociateInput)(nil)).Elem(), &AllowListAssociate{})
	pulumi.RegisterInputType(reflect.TypeOf((*AllowListAssociateArrayInput)(nil)).Elem(), AllowListAssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AllowListAssociateMapInput)(nil)).Elem(), AllowListAssociateMap{})
	pulumi.RegisterOutputType(AllowListAssociateOutput{})
	pulumi.RegisterOutputType(AllowListAssociateArrayOutput{})
	pulumi.RegisterOutputType(AllowListAssociateMapOutput{})
}
