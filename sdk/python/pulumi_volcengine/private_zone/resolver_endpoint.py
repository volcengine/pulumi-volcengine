# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ResolverEndpointArgs', 'ResolverEndpoint']

@pulumi.input_type
class ResolverEndpointArgs:
    def __init__(__self__, *,
                 ip_configs: pulumi.Input[Sequence[pulumi.Input['ResolverEndpointIpConfigArgs']]],
                 security_group_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 vpc_region: pulumi.Input[str],
                 direction: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ResolverEndpoint resource.
        :param pulumi.Input[Sequence[pulumi.Input['ResolverEndpointIpConfigArgs']]] ip_configs: Availability zones, subnets, and IP configurations of terminal nodes.
        :param pulumi.Input[str] security_group_id: The security group ID of the endpoint.
        :param pulumi.Input[str] vpc_id: The VPC ID of the endpoint.
        :param pulumi.Input[str] vpc_region: The VPC region of the endpoint.
        :param pulumi.Input[str] direction: DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        :param pulumi.Input[str] name: The name of the private zone resolver endpoint.
        """
        pulumi.set(__self__, "ip_configs", ip_configs)
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        pulumi.set(__self__, "vpc_region", vpc_region)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="ipConfigs")
    def ip_configs(self) -> pulumi.Input[Sequence[pulumi.Input['ResolverEndpointIpConfigArgs']]]:
        """
        Availability zones, subnets, and IP configurations of terminal nodes.
        """
        return pulumi.get(self, "ip_configs")

    @ip_configs.setter
    def ip_configs(self, value: pulumi.Input[Sequence[pulumi.Input['ResolverEndpointIpConfigArgs']]]):
        pulumi.set(self, "ip_configs", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        The security group ID of the endpoint.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The VPC ID of the endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcRegion")
    def vpc_region(self) -> pulumi.Input[str]:
        """
        The VPC region of the endpoint.
        """
        return pulumi.get(self, "vpc_region")

    @vpc_region.setter
    def vpc_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_region", value)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the private zone resolver endpoint.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ResolverEndpointState:
    def __init__(__self__, *,
                 direction: Optional[pulumi.Input[str]] = None,
                 ip_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ResolverEndpointIpConfigArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_region: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ResolverEndpoint resources.
        :param pulumi.Input[str] direction: DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        :param pulumi.Input[Sequence[pulumi.Input['ResolverEndpointIpConfigArgs']]] ip_configs: Availability zones, subnets, and IP configurations of terminal nodes.
        :param pulumi.Input[str] name: The name of the private zone resolver endpoint.
        :param pulumi.Input[str] security_group_id: The security group ID of the endpoint.
        :param pulumi.Input[str] vpc_id: The VPC ID of the endpoint.
        :param pulumi.Input[str] vpc_region: The VPC region of the endpoint.
        """
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if ip_configs is not None:
            pulumi.set(__self__, "ip_configs", ip_configs)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if security_group_id is not None:
            pulumi.set(__self__, "security_group_id", security_group_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_region is not None:
            pulumi.set(__self__, "vpc_region", vpc_region)

    @property
    @pulumi.getter
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter(name="ipConfigs")
    def ip_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ResolverEndpointIpConfigArgs']]]]:
        """
        Availability zones, subnets, and IP configurations of terminal nodes.
        """
        return pulumi.get(self, "ip_configs")

    @ip_configs.setter
    def ip_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ResolverEndpointIpConfigArgs']]]]):
        pulumi.set(self, "ip_configs", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the private zone resolver endpoint.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The security group ID of the endpoint.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC ID of the endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcRegion")
    def vpc_region(self) -> Optional[pulumi.Input[str]]:
        """
        The VPC region of the endpoint.
        """
        return pulumi.get(self, "vpc_region")

    @vpc_region.setter
    def vpc_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_region", value)


class ResolverEndpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 ip_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResolverEndpointIpConfigArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage private zone resolver endpoint
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.private_zone.ResolverEndpoint("foo",
            ip_configs=[
                volcengine.private_zone.ResolverEndpointIpConfigArgs(
                    az_id="cn-beijing-a",
                    ip="172.16.0.2",
                    subnet_id="subnet-mj2o4co2m2v45smt1bx1****",
                ),
                volcengine.private_zone.ResolverEndpointIpConfigArgs(
                    az_id="cn-beijing-a",
                    ip="172.16.0.3",
                    subnet_id="subnet-mj2o4co2m2v45smt1bx1****",
                ),
                volcengine.private_zone.ResolverEndpointIpConfigArgs(
                    az_id="cn-beijing-a",
                    ip="172.16.0.4",
                    subnet_id="subnet-mj2o4co2m2v45smt1bx1****",
                ),
                volcengine.private_zone.ResolverEndpointIpConfigArgs(
                    az_id="cn-beijing-a",
                    ip="172.16.0.5",
                    subnet_id="subnet-mj2o4co2m2v45smt1bx1****",
                ),
            ],
            security_group_id="sg-mj2nsckay29s5smt1b0d****",
            vpc_id="vpc-13f9uuuqfdjb43n6nu5p1****",
            vpc_region="cn-beijing")
        ```

        ## Import

        PrivateZoneResolverEndpoint can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:private_zone/resolverEndpoint:ResolverEndpoint default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] direction: DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResolverEndpointIpConfigArgs']]]] ip_configs: Availability zones, subnets, and IP configurations of terminal nodes.
        :param pulumi.Input[str] name: The name of the private zone resolver endpoint.
        :param pulumi.Input[str] security_group_id: The security group ID of the endpoint.
        :param pulumi.Input[str] vpc_id: The VPC ID of the endpoint.
        :param pulumi.Input[str] vpc_region: The VPC region of the endpoint.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResolverEndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage private zone resolver endpoint
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.private_zone.ResolverEndpoint("foo",
            ip_configs=[
                volcengine.private_zone.ResolverEndpointIpConfigArgs(
                    az_id="cn-beijing-a",
                    ip="172.16.0.2",
                    subnet_id="subnet-mj2o4co2m2v45smt1bx1****",
                ),
                volcengine.private_zone.ResolverEndpointIpConfigArgs(
                    az_id="cn-beijing-a",
                    ip="172.16.0.3",
                    subnet_id="subnet-mj2o4co2m2v45smt1bx1****",
                ),
                volcengine.private_zone.ResolverEndpointIpConfigArgs(
                    az_id="cn-beijing-a",
                    ip="172.16.0.4",
                    subnet_id="subnet-mj2o4co2m2v45smt1bx1****",
                ),
                volcengine.private_zone.ResolverEndpointIpConfigArgs(
                    az_id="cn-beijing-a",
                    ip="172.16.0.5",
                    subnet_id="subnet-mj2o4co2m2v45smt1bx1****",
                ),
            ],
            security_group_id="sg-mj2nsckay29s5smt1b0d****",
            vpc_id="vpc-13f9uuuqfdjb43n6nu5p1****",
            vpc_region="cn-beijing")
        ```

        ## Import

        PrivateZoneResolverEndpoint can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:private_zone/resolverEndpoint:ResolverEndpoint default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param ResolverEndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResolverEndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 ip_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResolverEndpointIpConfigArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 security_group_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_region: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResolverEndpointArgs.__new__(ResolverEndpointArgs)

            __props__.__dict__["direction"] = direction
            if ip_configs is None and not opts.urn:
                raise TypeError("Missing required property 'ip_configs'")
            __props__.__dict__["ip_configs"] = ip_configs
            __props__.__dict__["name"] = name
            if security_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_id'")
            __props__.__dict__["security_group_id"] = security_group_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            if vpc_region is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_region'")
            __props__.__dict__["vpc_region"] = vpc_region
        super(ResolverEndpoint, __self__).__init__(
            'volcengine:private_zone/resolverEndpoint:ResolverEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            direction: Optional[pulumi.Input[str]] = None,
            ip_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResolverEndpointIpConfigArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            security_group_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vpc_region: Optional[pulumi.Input[str]] = None) -> 'ResolverEndpoint':
        """
        Get an existing ResolverEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] direction: DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ResolverEndpointIpConfigArgs']]]] ip_configs: Availability zones, subnets, and IP configurations of terminal nodes.
        :param pulumi.Input[str] name: The name of the private zone resolver endpoint.
        :param pulumi.Input[str] security_group_id: The security group ID of the endpoint.
        :param pulumi.Input[str] vpc_id: The VPC ID of the endpoint.
        :param pulumi.Input[str] vpc_region: The VPC region of the endpoint.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResolverEndpointState.__new__(_ResolverEndpointState)

        __props__.__dict__["direction"] = direction
        __props__.__dict__["ip_configs"] = ip_configs
        __props__.__dict__["name"] = name
        __props__.__dict__["security_group_id"] = security_group_id
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vpc_region"] = vpc_region
        return ResolverEndpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Output[Optional[str]]:
        """
        DNS request forwarding direction for terminal nodes. OUTBOUND: (default) Outbound terminal nodes forward DNS query requests from within the VPC to external DNS servers. INBOUND: Inbound terminal nodes forward DNS query requests from external sources to resolvers.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter(name="ipConfigs")
    def ip_configs(self) -> pulumi.Output[Sequence['outputs.ResolverEndpointIpConfig']]:
        """
        Availability zones, subnets, and IP configurations of terminal nodes.
        """
        return pulumi.get(self, "ip_configs")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the private zone resolver endpoint.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Output[str]:
        """
        The security group ID of the endpoint.
        """
        return pulumi.get(self, "security_group_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC ID of the endpoint.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcRegion")
    def vpc_region(self) -> pulumi.Output[str]:
        """
        The VPC region of the endpoint.
        """
        return pulumi.get(self, "vpc_region")
