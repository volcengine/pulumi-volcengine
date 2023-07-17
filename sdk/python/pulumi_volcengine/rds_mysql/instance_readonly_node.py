# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['InstanceReadonlyNodeArgs', 'InstanceReadonlyNode']

@pulumi.input_type
class InstanceReadonlyNodeArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 node_spec: pulumi.Input[str],
                 zone_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a InstanceReadonlyNode resource.
        :param pulumi.Input[str] instance_id: The RDS mysql instance id of the readonly node.
        :param pulumi.Input[str] node_spec: The specification of readonly node.
        :param pulumi.Input[str] zone_id: The available zone of readonly node.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "node_spec", node_spec)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The RDS mysql instance id of the readonly node.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="nodeSpec")
    def node_spec(self) -> pulumi.Input[str]:
        """
        The specification of readonly node.
        """
        return pulumi.get(self, "node_spec")

    @node_spec.setter
    def node_spec(self, value: pulumi.Input[str]):
        pulumi.set(self, "node_spec", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The available zone of readonly node.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class _InstanceReadonlyNodeState:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 node_spec: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering InstanceReadonlyNode resources.
        :param pulumi.Input[str] instance_id: The RDS mysql instance id of the readonly node.
        :param pulumi.Input[str] node_id: The id of the readonly node.
        :param pulumi.Input[str] node_spec: The specification of readonly node.
        :param pulumi.Input[str] zone_id: The available zone of readonly node.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)
        if node_spec is not None:
            pulumi.set(__self__, "node_spec", node_spec)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The RDS mysql instance id of the readonly node.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the readonly node.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)

    @property
    @pulumi.getter(name="nodeSpec")
    def node_spec(self) -> Optional[pulumi.Input[str]]:
        """
        The specification of readonly node.
        """
        return pulumi.get(self, "node_spec")

    @node_spec.setter
    def node_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_spec", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The available zone of readonly node.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


class InstanceReadonlyNode(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 node_spec: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage rds mysql instance readonly node
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.rds_mysql.InstanceReadonlyNode("foo",
            instance_id="mysql-b3fca7f571d6",
            node_spec="rds.mysql.1c2g",
            zone_id="cn-guilin-b")
        ```

        ## Import

        Rds Mysql Instance Readonly Node can be imported using the instance_id:node_id, e.g.

        ```sh
         $ pulumi import volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode default mysql-72da4258c2c7:mysql-72da4258c2c7-r7f93
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The RDS mysql instance id of the readonly node.
        :param pulumi.Input[str] node_spec: The specification of readonly node.
        :param pulumi.Input[str] zone_id: The available zone of readonly node.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceReadonlyNodeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage rds mysql instance readonly node
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.rds_mysql.InstanceReadonlyNode("foo",
            instance_id="mysql-b3fca7f571d6",
            node_spec="rds.mysql.1c2g",
            zone_id="cn-guilin-b")
        ```

        ## Import

        Rds Mysql Instance Readonly Node can be imported using the instance_id:node_id, e.g.

        ```sh
         $ pulumi import volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode default mysql-72da4258c2c7:mysql-72da4258c2c7-r7f93
        ```

        :param str resource_name: The name of the resource.
        :param InstanceReadonlyNodeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceReadonlyNodeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 node_spec: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InstanceReadonlyNodeArgs.__new__(InstanceReadonlyNodeArgs)

            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if node_spec is None and not opts.urn:
                raise TypeError("Missing required property 'node_spec'")
            __props__.__dict__["node_spec"] = node_spec
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["node_id"] = None
        super(InstanceReadonlyNode, __self__).__init__(
            'volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            node_id: Optional[pulumi.Input[str]] = None,
            node_spec: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'InstanceReadonlyNode':
        """
        Get an existing InstanceReadonlyNode resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] instance_id: The RDS mysql instance id of the readonly node.
        :param pulumi.Input[str] node_id: The id of the readonly node.
        :param pulumi.Input[str] node_spec: The specification of readonly node.
        :param pulumi.Input[str] zone_id: The available zone of readonly node.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InstanceReadonlyNodeState.__new__(_InstanceReadonlyNodeState)

        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["node_id"] = node_id
        __props__.__dict__["node_spec"] = node_spec
        __props__.__dict__["zone_id"] = zone_id
        return InstanceReadonlyNode(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The RDS mysql instance id of the readonly node.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> pulumi.Output[str]:
        """
        The id of the readonly node.
        """
        return pulumi.get(self, "node_id")

    @property
    @pulumi.getter(name="nodeSpec")
    def node_spec(self) -> pulumi.Output[str]:
        """
        The specification of readonly node.
        """
        return pulumi.get(self, "node_spec")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The available zone of readonly node.
        """
        return pulumi.get(self, "zone_id")
