# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceParameterArgs', 'InstanceParameter']

@pulumi.input_type
class InstanceParameterArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 parameter_name: pulumi.Input[str],
                 parameter_role: pulumi.Input[str],
                 parameter_value: pulumi.Input[str]):
        """
        The set of arguments for constructing a InstanceParameter resource.
        :param pulumi.Input[str] instance_id: The instance ID.
        :param pulumi.Input[str] parameter_name: The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        :param pulumi.Input[str] parameter_role: The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        :param pulumi.Input[str] parameter_value: The value of parameter.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "parameter_name", parameter_name)
        pulumi.set(__self__, "parameter_role", parameter_role)
        pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="parameterName")
    def parameter_name(self) -> pulumi.Input[str]:
        """
        The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        """
        return pulumi.get(self, "parameter_name")

    @parameter_name.setter
    def parameter_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_name", value)

    @property
    @pulumi.getter(name="parameterRole")
    def parameter_role(self) -> pulumi.Input[str]:
        """
        The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        """
        return pulumi.get(self, "parameter_role")

    @parameter_role.setter
    def parameter_role(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_role", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> pulumi.Input[str]:
        """
        The value of parameter.
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "parameter_value", value)


@pulumi.input_type
class _InstanceParameterState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 parameter_name: Optional[pulumi.Input[str]] = None,
                 parameter_role: Optional[pulumi.Input[str]] = None,
                 parameter_value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceParameter resources.
        :param pulumi.Input[str] instance_id: The instance ID.
        :param pulumi.Input[str] parameter_name: The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        :param pulumi.Input[str] parameter_role: The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        :param pulumi.Input[str] parameter_value: The value of parameter.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if parameter_name is not None:
            pulumi.set(__self__, "parameter_name", parameter_name)
        if parameter_role is not None:
            pulumi.set(__self__, "parameter_role", parameter_role)
        if parameter_value is not None:
            pulumi.set(__self__, "parameter_value", parameter_value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="parameterName")
    def parameter_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        """
        return pulumi.get(self, "parameter_name")

    @parameter_name.setter
    def parameter_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_name", value)

    @property
    @pulumi.getter(name="parameterRole")
    def parameter_role(self) -> Optional[pulumi.Input[str]]:
        """
        The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        """
        return pulumi.get(self, "parameter_role")

    @parameter_role.setter
    def parameter_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_role", value)

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> Optional[pulumi.Input[str]]:
        """
        The value of parameter.
        """
        return pulumi.get(self, "parameter_value")

    @parameter_value.setter
    def parameter_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameter_value", value)


class InstanceParameter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 parameter_name: Optional[pulumi.Input[str]] = None,
                 parameter_role: Optional[pulumi.Input[str]] = None,
                 parameter_value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage mongodb instance parameter
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.get_zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        foo_instance = volcengine.mongodb.Instance("fooInstance",
            db_engine_version="MongoDB_4_0",
            instance_type="ReplicaSet",
            super_account_password="@acc-test-123",
            node_spec="mongo.2c4g",
            mongos_node_spec="mongo.mongos.2c4g",
            instance_name="acc-test-mongo-replica",
            charge_type="PostPaid",
            project_name="default",
            mongos_node_number=32,
            shard_number=3,
            storage_space_gb=20,
            subnet_id=foo_subnet.id,
            zone_id=foo_zones.zones[0].id,
            tags=[volcengine.mongodb.InstanceTagArgs(
                key="k1",
                value="v1",
            )])
        foo_instance_parameter = volcengine.mongodb.InstanceParameter("fooInstanceParameter",
            instance_id=foo_instance.id,
            parameter_name="cursorTimeoutMillis",
            parameter_role="Node",
            parameter_value="600111")
        ```

        ## Import

        mongodb parameter can be imported using the param:instanceId:parameterName:parameterRole, e.g.

        ```sh
        $ pulumi import volcengine:mongodb/instanceParameter:InstanceParameter default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The instance ID.
        :param pulumi.Input[str] parameter_name: The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        :param pulumi.Input[str] parameter_role: The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        :param pulumi.Input[str] parameter_value: The value of parameter.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceParameterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage mongodb instance parameter
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.get_zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        foo_instance = volcengine.mongodb.Instance("fooInstance",
            db_engine_version="MongoDB_4_0",
            instance_type="ReplicaSet",
            super_account_password="@acc-test-123",
            node_spec="mongo.2c4g",
            mongos_node_spec="mongo.mongos.2c4g",
            instance_name="acc-test-mongo-replica",
            charge_type="PostPaid",
            project_name="default",
            mongos_node_number=32,
            shard_number=3,
            storage_space_gb=20,
            subnet_id=foo_subnet.id,
            zone_id=foo_zones.zones[0].id,
            tags=[volcengine.mongodb.InstanceTagArgs(
                key="k1",
                value="v1",
            )])
        foo_instance_parameter = volcengine.mongodb.InstanceParameter("fooInstanceParameter",
            instance_id=foo_instance.id,
            parameter_name="cursorTimeoutMillis",
            parameter_role="Node",
            parameter_value="600111")
        ```

        ## Import

        mongodb parameter can be imported using the param:instanceId:parameterName:parameterRole, e.g.

        ```sh
        $ pulumi import volcengine:mongodb/instanceParameter:InstanceParameter default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
        ```

        :param str resource_name: The name of the resource.
        :param InstanceParameterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceParameterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 parameter_name: Optional[pulumi.Input[str]] = None,
                 parameter_role: Optional[pulumi.Input[str]] = None,
                 parameter_value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceParameterArgs.__new__(InstanceParameterArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if parameter_name is None and not opts.urn:
                raise TypeError("Missing required property 'parameter_name'")
            __props__.__dict__["parameter_name"] = parameter_name
            if parameter_role is None and not opts.urn:
                raise TypeError("Missing required property 'parameter_role'")
            __props__.__dict__["parameter_role"] = parameter_role
            if parameter_value is None and not opts.urn:
                raise TypeError("Missing required property 'parameter_value'")
            __props__.__dict__["parameter_value"] = parameter_value
        super(InstanceParameter, __self__).__init__(
            'volcengine:mongodb/instanceParameter:InstanceParameter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            parameter_name: Optional[pulumi.Input[str]] = None,
            parameter_role: Optional[pulumi.Input[str]] = None,
            parameter_value: Optional[pulumi.Input[str]] = None) -> 'InstanceParameter':
        """
        Get an existing InstanceParameter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The instance ID.
        :param pulumi.Input[str] parameter_name: The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        :param pulumi.Input[str] parameter_role: The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        :param pulumi.Input[str] parameter_value: The value of parameter.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceParameterState.__new__(_InstanceParameterState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["parameter_name"] = parameter_name
        __props__.__dict__["parameter_role"] = parameter_role
        __props__.__dict__["parameter_value"] = parameter_value
        return InstanceParameter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The instance ID.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="parameterName")
    def parameter_name(self) -> pulumi.Output[str]:
        """
        The name of parameter. The parameter resource can only be added or modified, deleting this resource will not actually execute any operation.
        """
        return pulumi.get(self, "parameter_name")

    @property
    @pulumi.getter(name="parameterRole")
    def parameter_role(self) -> pulumi.Output[str]:
        """
        The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
        """
        return pulumi.get(self, "parameter_role")

    @property
    @pulumi.getter(name="parameterValue")
    def parameter_value(self) -> pulumi.Output[str]:
        """
        The value of parameter.
        """
        return pulumi.get(self, "parameter_value")

