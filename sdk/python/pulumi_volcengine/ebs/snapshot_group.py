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

__all__ = ['SnapshotGroupArgs', 'SnapshotGroup']

@pulumi.input_type
class SnapshotGroupArgs:
    def __init__(__self__, *,
                 volume_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotGroupTagArgs']]]] = None):
        """
        The set of arguments for constructing a SnapshotGroup resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] volume_ids: The volume id of the snapshot group. The status of the volume must be `attached`.If multiple volumes are specified, they need to be attached to the same ECS instance.
        :param pulumi.Input[str] description: The instance id of the snapshot group.
        :param pulumi.Input[str] instance_id: The instance id of the snapshot group.
        :param pulumi.Input[str] name: The name of the snapshot group.
        :param pulumi.Input[str] project_name: The project name of the snapshot group.
        :param pulumi.Input[Sequence[pulumi.Input['SnapshotGroupTagArgs']]] tags: Tags.
        """
        pulumi.set(__self__, "volume_ids", volume_ids)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="volumeIds")
    def volume_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The volume id of the snapshot group. The status of the volume must be `attached`.If multiple volumes are specified, they need to be attached to the same ECS instance.
        """
        return pulumi.get(self, "volume_ids")

    @volume_ids.setter
    def volume_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "volume_ids", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The instance id of the snapshot group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance id of the snapshot group.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the snapshot group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of the snapshot group.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotGroupTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotGroupTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SnapshotGroupState:
    def __init__(__self__, *,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotGroupTagArgs']]]] = None,
                 volume_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SnapshotGroup resources.
        :param pulumi.Input[str] creation_time: The creation time of the snapshot group.
        :param pulumi.Input[str] description: The instance id of the snapshot group.
        :param pulumi.Input[str] image_id: The image id of the snapshot group.
        :param pulumi.Input[str] instance_id: The instance id of the snapshot group.
        :param pulumi.Input[str] name: The name of the snapshot group.
        :param pulumi.Input[str] project_name: The project name of the snapshot group.
        :param pulumi.Input[str] status: The status of the snapshot group.
        :param pulumi.Input[Sequence[pulumi.Input['SnapshotGroupTagArgs']]] tags: Tags.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] volume_ids: The volume id of the snapshot group. The status of the volume must be `attached`.If multiple volumes are specified, they need to be attached to the same ECS instance.
        """
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if volume_ids is not None:
            pulumi.set(__self__, "volume_ids", volume_ids)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the snapshot group.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The instance id of the snapshot group.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The image id of the snapshot group.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance id of the snapshot group.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the snapshot group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of the snapshot group.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the snapshot group.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotGroupTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SnapshotGroupTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="volumeIds")
    def volume_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The volume id of the snapshot group. The status of the volume must be `attached`.If multiple volumes are specified, they need to be attached to the same ECS instance.
        """
        return pulumi.get(self, "volume_ids")

    @volume_ids.setter
    def volume_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "volume_ids", value)


class SnapshotGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotGroupTagArgs']]]]] = None,
                 volume_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage ebs snapshot group
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
            security_group_name="acc-test-security-group",
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
        foo_volume = volcengine.ebs.Volume("fooVolume",
            volume_name="acc-test-volume",
            volume_type="ESSD_PL0",
            description="acc-test",
            kind="data",
            size=500,
            zone_id=foo_zones.zones[0].id,
            volume_charge_type="PostPaid",
            project_name="default")
        foo_volume_attach = volcengine.ebs.VolumeAttach("fooVolumeAttach",
            instance_id=foo_instance.id,
            volume_id=foo_volume.id)
        foo_snapshot_group = volcengine.ebs.SnapshotGroup("fooSnapshotGroup",
            volume_ids=[
                foo_instance.system_volume_id,
                foo_volume.id,
            ],
            instance_id=foo_instance.id,
            description="acc-test",
            project_name="default",
            tags=[volcengine.ebs.SnapshotGroupTagArgs(
                key="k1",
                value="v1",
            )],
            opts=pulumi.ResourceOptions(depends_on=[foo_volume_attach]))
        ```

        ## Import

        EbsSnapshotGroup can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:ebs/snapshotGroup:SnapshotGroup default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The instance id of the snapshot group.
        :param pulumi.Input[str] instance_id: The instance id of the snapshot group.
        :param pulumi.Input[str] name: The name of the snapshot group.
        :param pulumi.Input[str] project_name: The project name of the snapshot group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotGroupTagArgs']]]] tags: Tags.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] volume_ids: The volume id of the snapshot group. The status of the volume must be `attached`.If multiple volumes are specified, they need to be attached to the same ECS instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SnapshotGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage ebs snapshot group
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
            security_group_name="acc-test-security-group",
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
        foo_volume = volcengine.ebs.Volume("fooVolume",
            volume_name="acc-test-volume",
            volume_type="ESSD_PL0",
            description="acc-test",
            kind="data",
            size=500,
            zone_id=foo_zones.zones[0].id,
            volume_charge_type="PostPaid",
            project_name="default")
        foo_volume_attach = volcengine.ebs.VolumeAttach("fooVolumeAttach",
            instance_id=foo_instance.id,
            volume_id=foo_volume.id)
        foo_snapshot_group = volcengine.ebs.SnapshotGroup("fooSnapshotGroup",
            volume_ids=[
                foo_instance.system_volume_id,
                foo_volume.id,
            ],
            instance_id=foo_instance.id,
            description="acc-test",
            project_name="default",
            tags=[volcengine.ebs.SnapshotGroupTagArgs(
                key="k1",
                value="v1",
            )],
            opts=pulumi.ResourceOptions(depends_on=[foo_volume_attach]))
        ```

        ## Import

        EbsSnapshotGroup can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:ebs/snapshotGroup:SnapshotGroup default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param SnapshotGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SnapshotGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotGroupTagArgs']]]]] = None,
                 volume_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SnapshotGroupArgs.__new__(SnapshotGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["name"] = name
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["tags"] = tags
            if volume_ids is None and not opts.urn:
                raise TypeError("Missing required property 'volume_ids'")
            __props__.__dict__["volume_ids"] = volume_ids
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["image_id"] = None
            __props__.__dict__["status"] = None
        super(SnapshotGroup, __self__).__init__(
            'volcengine:ebs/snapshotGroup:SnapshotGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotGroupTagArgs']]]]] = None,
            volume_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SnapshotGroup':
        """
        Get an existing SnapshotGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_time: The creation time of the snapshot group.
        :param pulumi.Input[str] description: The instance id of the snapshot group.
        :param pulumi.Input[str] image_id: The image id of the snapshot group.
        :param pulumi.Input[str] instance_id: The instance id of the snapshot group.
        :param pulumi.Input[str] name: The name of the snapshot group.
        :param pulumi.Input[str] project_name: The project name of the snapshot group.
        :param pulumi.Input[str] status: The status of the snapshot group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SnapshotGroupTagArgs']]]] tags: Tags.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] volume_ids: The volume id of the snapshot group. The status of the volume must be `attached`.If multiple volumes are specified, they need to be attached to the same ECS instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SnapshotGroupState.__new__(_SnapshotGroupState)

        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["name"] = name
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["volume_ids"] = volume_ids
        return SnapshotGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The creation time of the snapshot group.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The instance id of the snapshot group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        """
        The image id of the snapshot group.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The instance id of the snapshot group.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the snapshot group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The project name of the snapshot group.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the snapshot group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.SnapshotGroupTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeIds")
    def volume_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The volume id of the snapshot group. The status of the volume must be `attached`.If multiple volumes are specified, they need to be attached to the same ECS instance.
        """
        return pulumi.get(self, "volume_ids")

