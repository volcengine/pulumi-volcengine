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

__all__ = ['TrafficMirrorTargetArgs', 'TrafficMirrorTarget']

@pulumi.input_type
class TrafficMirrorTargetArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 instance_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TrafficMirrorTargetTagArgs']]]] = None,
                 traffic_mirror_target_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TrafficMirrorTarget resource.
        :param pulumi.Input[str] instance_id: The instance id of traffic mirror target.
        :param pulumi.Input[str] instance_type: The instance type of traffic mirror target. Valid values: `NetworkInterface`, `ClbInstance`.
        :param pulumi.Input[str] description: The description of traffic mirror target.
        :param pulumi.Input[str] project_name: The project name of traffic mirror target.
        :param pulumi.Input[Sequence[pulumi.Input['TrafficMirrorTargetTagArgs']]] tags: Tags.
        :param pulumi.Input[str] traffic_mirror_target_name: The name of traffic mirror target.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_type", instance_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if traffic_mirror_target_name is not None:
            pulumi.set(__self__, "traffic_mirror_target_name", traffic_mirror_target_name)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The instance id of traffic mirror target.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        The instance type of traffic mirror target. Valid values: `NetworkInterface`, `ClbInstance`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of traffic mirror target.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of traffic mirror target.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrafficMirrorTargetTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrafficMirrorTargetTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trafficMirrorTargetName")
    def traffic_mirror_target_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of traffic mirror target.
        """
        return pulumi.get(self, "traffic_mirror_target_name")

    @traffic_mirror_target_name.setter
    def traffic_mirror_target_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_mirror_target_name", value)


@pulumi.input_type
class _TrafficMirrorTargetState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TrafficMirrorTargetTagArgs']]]] = None,
                 traffic_mirror_target_name: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TrafficMirrorTarget resources.
        :param pulumi.Input[str] created_at: The create time of traffic mirror target.
        :param pulumi.Input[str] description: The description of traffic mirror target.
        :param pulumi.Input[str] instance_id: The instance id of traffic mirror target.
        :param pulumi.Input[str] instance_type: The instance type of traffic mirror target. Valid values: `NetworkInterface`, `ClbInstance`.
        :param pulumi.Input[str] project_name: The project name of traffic mirror target.
        :param pulumi.Input[str] status: The status of traffic mirror target.
        :param pulumi.Input[Sequence[pulumi.Input['TrafficMirrorTargetTagArgs']]] tags: Tags.
        :param pulumi.Input[str] traffic_mirror_target_name: The name of traffic mirror target.
        :param pulumi.Input[str] updated_at: The update time of traffic mirror target.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if traffic_mirror_target_name is not None:
            pulumi.set(__self__, "traffic_mirror_target_name", traffic_mirror_target_name)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of traffic mirror target.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of traffic mirror target.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance id of traffic mirror target.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The instance type of traffic mirror target. Valid values: `NetworkInterface`, `ClbInstance`.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of traffic mirror target.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of traffic mirror target.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TrafficMirrorTargetTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TrafficMirrorTargetTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="trafficMirrorTargetName")
    def traffic_mirror_target_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of traffic mirror target.
        """
        return pulumi.get(self, "traffic_mirror_target_name")

    @traffic_mirror_target_name.setter
    def traffic_mirror_target_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "traffic_mirror_target_name", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of traffic mirror target.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class TrafficMirrorTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrafficMirrorTargetTagArgs']]]]] = None,
                 traffic_mirror_target_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage traffic mirror target
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
        foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
            security_group_name="acc-test-sg",
            vpc_id=foo_vpc.id)
        foo_images = volcengine.ecs.get_images(os_type="Linux",
            visibility="public",
            instance_type_id="ecs.g3il.large")
        foo_instance = volcengine.ecs.Instance("fooInstance",
            instance_name="acc-test-ecs",
            description="acc-test",
            host_name="tf-acc-test",
            image_id=foo_images.images[0].image_id,
            instance_type="ecs.g3il.large",
            password="93f0cb0614Aab12",
            instance_charge_type="PostPaid",
            system_volume_type="ESSD_PL0",
            system_volume_size=40,
            subnet_id=foo_subnet.id,
            security_group_ids=[foo_security_group.id],
            project_name="default",
            tags=[volcengine.ecs.InstanceTagArgs(
                key="k1",
                value="v1",
            )])
        foo_network_interface = volcengine.vpc.NetworkInterface("fooNetworkInterface",
            network_interface_name="acc-test-eni",
            description="acc-test",
            subnet_id=foo_subnet.id,
            security_group_ids=[foo_security_group.id],
            primary_ip_address="172.16.0.253",
            port_security_enabled=False,
            private_ip_addresses=["172.16.0.2"],
            project_name="default",
            tags=[volcengine.vpc.NetworkInterfaceTagArgs(
                key="k1",
                value="v1",
            )])
        foo_network_interface_attach = volcengine.vpc.NetworkInterfaceAttach("fooNetworkInterfaceAttach",
            network_interface_id=foo_network_interface.id,
            instance_id=foo_instance.id)
        foo_traffic_mirror_target = volcengine.vpc.TrafficMirrorTarget("fooTrafficMirrorTarget",
            instance_type="NetworkInterface",
            instance_id=foo_network_interface.id,
            traffic_mirror_target_name="acc-test-traffic-mirror-target",
            description="acc-test",
            project_name="default",
            tags=[volcengine.vpc.TrafficMirrorTargetTagArgs(
                key="k1",
                value="v1",
            )],
            opts=pulumi.ResourceOptions(depends_on=[foo_network_interface_attach]))
        ```

        ## Import

        TrafficMirrorTarget can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:vpc/trafficMirrorTarget:TrafficMirrorTarget default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of traffic mirror target.
        :param pulumi.Input[str] instance_id: The instance id of traffic mirror target.
        :param pulumi.Input[str] instance_type: The instance type of traffic mirror target. Valid values: `NetworkInterface`, `ClbInstance`.
        :param pulumi.Input[str] project_name: The project name of traffic mirror target.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrafficMirrorTargetTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] traffic_mirror_target_name: The name of traffic mirror target.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TrafficMirrorTargetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage traffic mirror target
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
        foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
            security_group_name="acc-test-sg",
            vpc_id=foo_vpc.id)
        foo_images = volcengine.ecs.get_images(os_type="Linux",
            visibility="public",
            instance_type_id="ecs.g3il.large")
        foo_instance = volcengine.ecs.Instance("fooInstance",
            instance_name="acc-test-ecs",
            description="acc-test",
            host_name="tf-acc-test",
            image_id=foo_images.images[0].image_id,
            instance_type="ecs.g3il.large",
            password="93f0cb0614Aab12",
            instance_charge_type="PostPaid",
            system_volume_type="ESSD_PL0",
            system_volume_size=40,
            subnet_id=foo_subnet.id,
            security_group_ids=[foo_security_group.id],
            project_name="default",
            tags=[volcengine.ecs.InstanceTagArgs(
                key="k1",
                value="v1",
            )])
        foo_network_interface = volcengine.vpc.NetworkInterface("fooNetworkInterface",
            network_interface_name="acc-test-eni",
            description="acc-test",
            subnet_id=foo_subnet.id,
            security_group_ids=[foo_security_group.id],
            primary_ip_address="172.16.0.253",
            port_security_enabled=False,
            private_ip_addresses=["172.16.0.2"],
            project_name="default",
            tags=[volcengine.vpc.NetworkInterfaceTagArgs(
                key="k1",
                value="v1",
            )])
        foo_network_interface_attach = volcengine.vpc.NetworkInterfaceAttach("fooNetworkInterfaceAttach",
            network_interface_id=foo_network_interface.id,
            instance_id=foo_instance.id)
        foo_traffic_mirror_target = volcengine.vpc.TrafficMirrorTarget("fooTrafficMirrorTarget",
            instance_type="NetworkInterface",
            instance_id=foo_network_interface.id,
            traffic_mirror_target_name="acc-test-traffic-mirror-target",
            description="acc-test",
            project_name="default",
            tags=[volcengine.vpc.TrafficMirrorTargetTagArgs(
                key="k1",
                value="v1",
            )],
            opts=pulumi.ResourceOptions(depends_on=[foo_network_interface_attach]))
        ```

        ## Import

        TrafficMirrorTarget can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:vpc/trafficMirrorTarget:TrafficMirrorTarget default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param TrafficMirrorTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficMirrorTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrafficMirrorTargetTagArgs']]]]] = None,
                 traffic_mirror_target_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TrafficMirrorTargetArgs.__new__(TrafficMirrorTargetArgs)

            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["traffic_mirror_target_name"] = traffic_mirror_target_name
            __props__.__dict__["created_at"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
        super(TrafficMirrorTarget, __self__).__init__(
            'volcengine:vpc/trafficMirrorTarget:TrafficMirrorTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrafficMirrorTargetTagArgs']]]]] = None,
            traffic_mirror_target_name: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'TrafficMirrorTarget':
        """
        Get an existing TrafficMirrorTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The create time of traffic mirror target.
        :param pulumi.Input[str] description: The description of traffic mirror target.
        :param pulumi.Input[str] instance_id: The instance id of traffic mirror target.
        :param pulumi.Input[str] instance_type: The instance type of traffic mirror target. Valid values: `NetworkInterface`, `ClbInstance`.
        :param pulumi.Input[str] project_name: The project name of traffic mirror target.
        :param pulumi.Input[str] status: The status of traffic mirror target.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TrafficMirrorTargetTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] traffic_mirror_target_name: The name of traffic mirror target.
        :param pulumi.Input[str] updated_at: The update time of traffic mirror target.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TrafficMirrorTargetState.__new__(_TrafficMirrorTargetState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_type"] = instance_type
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["traffic_mirror_target_name"] = traffic_mirror_target_name
        __props__.__dict__["updated_at"] = updated_at
        return TrafficMirrorTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The create time of traffic mirror target.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of traffic mirror target.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The instance id of traffic mirror target.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        The instance type of traffic mirror target. Valid values: `NetworkInterface`, `ClbInstance`.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The project name of traffic mirror target.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of traffic mirror target.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.TrafficMirrorTargetTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficMirrorTargetName")
    def traffic_mirror_target_name(self) -> pulumi.Output[str]:
        """
        The name of traffic mirror target.
        """
        return pulumi.get(self, "traffic_mirror_target_name")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The update time of traffic mirror target.
        """
        return pulumi.get(self, "updated_at")

