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

__all__ = ['FileSystemArgs', 'FileSystem']

@pulumi.input_type
class FileSystemArgs:
    def __init__(__self__, *,
                 capacity: pulumi.Input[int],
                 file_system_name: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemTagArgs']]]] = None):
        """
        The set of arguments for constructing a FileSystem resource.
        :param pulumi.Input[int] capacity: The capacity of the nas file system. Unit: GiB.
        :param pulumi.Input[str] file_system_name: The name of the nas file system.
        :param pulumi.Input[str] zone_id: The zone id of the nas file system.
        :param pulumi.Input[str] description: The description of the nas file system.
        :param pulumi.Input[str] project_name: The project name of the nas file system.
        :param pulumi.Input[str] snapshot_id: The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[Sequence[pulumi.Input['FileSystemTagArgs']]] tags: Tags.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "file_system_name", file_system_name)
        pulumi.set(__self__, "zone_id", zone_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Input[int]:
        """
        The capacity of the nas file system. Unit: GiB.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: pulumi.Input[int]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> pulumi.Input[str]:
        """
        The name of the nas file system.
        """
        return pulumi.get(self, "file_system_name")

    @file_system_name.setter
    def file_system_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        The zone id of the nas file system.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the nas file system.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of the nas file system.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _FileSystemState:
    def __init__(__self__, *,
                 capacity: Optional[pulumi.Input[int]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 file_system_type: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 protocol_type: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 snapshot_count: Optional[pulumi.Input[int]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemTagArgs']]]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 zone_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FileSystem resources.
        :param pulumi.Input[int] capacity: The capacity of the nas file system. Unit: GiB.
        :param pulumi.Input[str] charge_type: The charge type of the nas file system.
        :param pulumi.Input[str] create_time: The create time of the nas file system.
        :param pulumi.Input[str] description: The description of the nas file system.
        :param pulumi.Input[str] file_system_name: The name of the nas file system.
        :param pulumi.Input[str] file_system_type: The type of the nas file system.
        :param pulumi.Input[str] project_name: The project name of the nas file system.
        :param pulumi.Input[str] protocol_type: The protocol type of the nas file system.
        :param pulumi.Input[str] region_id: The region id of the nas file system.
        :param pulumi.Input[int] snapshot_count: The snapshot count of the nas file system.
        :param pulumi.Input[str] snapshot_id: The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[str] status: The status of the nas file system.
        :param pulumi.Input[str] storage_type: The storage type of the nas file system.
        :param pulumi.Input[Sequence[pulumi.Input['FileSystemTagArgs']]] tags: Tags.
        :param pulumi.Input[str] update_time: The update time of the nas file system.
        :param pulumi.Input[str] version: The version of the nas file system.
        :param pulumi.Input[str] zone_id: The zone id of the nas file system.
        :param pulumi.Input[str] zone_name: The zone name of the nas file system.
        """
        if capacity is not None:
            pulumi.set(__self__, "capacity", capacity)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if file_system_name is not None:
            pulumi.set(__self__, "file_system_name", file_system_name)
        if file_system_type is not None:
            pulumi.set(__self__, "file_system_type", file_system_type)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if protocol_type is not None:
            pulumi.set(__self__, "protocol_type", protocol_type)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)
        if snapshot_count is not None:
            pulumi.set(__self__, "snapshot_count", snapshot_count)
        if snapshot_id is not None:
            pulumi.set(__self__, "snapshot_id", snapshot_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if storage_type is not None:
            pulumi.set(__self__, "storage_type", storage_type)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)
        if zone_name is not None:
            pulumi.set(__self__, "zone_name", zone_name)

    @property
    @pulumi.getter
    def capacity(self) -> Optional[pulumi.Input[int]]:
        """
        The capacity of the nas file system. Unit: GiB.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        The charge type of the nas file system.
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of the nas file system.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the nas file system.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the nas file system.
        """
        return pulumi.get(self, "file_system_name")

    @file_system_name.setter
    def file_system_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_name", value)

    @property
    @pulumi.getter(name="fileSystemType")
    def file_system_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the nas file system.
        """
        return pulumi.get(self, "file_system_type")

    @file_system_type.setter
    def file_system_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_system_type", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of the nas file system.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol type of the nas file system.
        """
        return pulumi.get(self, "protocol_type")

    @protocol_type.setter
    def protocol_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol_type", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The region id of the nas file system.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter(name="snapshotCount")
    def snapshot_count(self) -> Optional[pulumi.Input[int]]:
        """
        The snapshot count of the nas file system.
        """
        return pulumi.get(self, "snapshot_count")

    @snapshot_count.setter
    def snapshot_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "snapshot_count", value)

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> Optional[pulumi.Input[str]]:
        """
        The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "snapshot_id")

    @snapshot_id.setter
    def snapshot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "snapshot_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the nas file system.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> Optional[pulumi.Input[str]]:
        """
        The storage type of the nas file system.
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_type", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FileSystemTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of the nas file system.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version of the nas file system.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The zone id of the nas file system.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> Optional[pulumi.Input[str]]:
        """
        The zone name of the nas file system.
        """
        return pulumi.get(self, "zone_name")

    @zone_name.setter
    def zone_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_name", value)


class FileSystem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileSystemTagArgs']]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage nas file system
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.nas.get_zones()
        # create nas file system
        foo_file_system = volcengine.nas.FileSystem("fooFileSystem",
            file_system_name="acc-test-fs",
            description="acc-test",
            zone_id=foo_zones.zones[0].id,
            capacity=103,
            project_name="default",
            tags=[volcengine.nas.FileSystemTagArgs(
                key="k1",
                value="v1",
            )])
        # create vpc
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16")
        # create subnet
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        # create nas permission group
        foo_permission_group = volcengine.nas.PermissionGroup("fooPermissionGroup",
            permission_group_name="acc-test",
            description="acctest",
            permission_rules=[
                volcengine.nas.PermissionGroupPermissionRuleArgs(
                    cidr_ip="*",
                    rw_mode="RW",
                    use_mode="All_squash",
                ),
                volcengine.nas.PermissionGroupPermissionRuleArgs(
                    cidr_ip="192.168.0.0",
                    rw_mode="RO",
                    use_mode="All_squash",
                ),
            ])
        # create nas mount point
        foo_mount_point = volcengine.nas.MountPoint("fooMountPoint",
            file_system_id=foo_file_system.id,
            mount_point_name="acc-test",
            permission_group_id=foo_permission_group.id,
            subnet_id=foo_subnet.id)
        ```

        ## Import

        NasFileSystem can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:nas/fileSystem:FileSystem default enas-cnbjd3879745****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] capacity: The capacity of the nas file system. Unit: GiB.
        :param pulumi.Input[str] description: The description of the nas file system.
        :param pulumi.Input[str] file_system_name: The name of the nas file system.
        :param pulumi.Input[str] project_name: The project name of the nas file system.
        :param pulumi.Input[str] snapshot_id: The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileSystemTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] zone_id: The zone id of the nas file system.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FileSystemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage nas file system
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.nas.get_zones()
        # create nas file system
        foo_file_system = volcengine.nas.FileSystem("fooFileSystem",
            file_system_name="acc-test-fs",
            description="acc-test",
            zone_id=foo_zones.zones[0].id,
            capacity=103,
            project_name="default",
            tags=[volcengine.nas.FileSystemTagArgs(
                key="k1",
                value="v1",
            )])
        # create vpc
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16")
        # create subnet
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        # create nas permission group
        foo_permission_group = volcengine.nas.PermissionGroup("fooPermissionGroup",
            permission_group_name="acc-test",
            description="acctest",
            permission_rules=[
                volcengine.nas.PermissionGroupPermissionRuleArgs(
                    cidr_ip="*",
                    rw_mode="RW",
                    use_mode="All_squash",
                ),
                volcengine.nas.PermissionGroupPermissionRuleArgs(
                    cidr_ip="192.168.0.0",
                    rw_mode="RO",
                    use_mode="All_squash",
                ),
            ])
        # create nas mount point
        foo_mount_point = volcengine.nas.MountPoint("fooMountPoint",
            file_system_id=foo_file_system.id,
            mount_point_name="acc-test",
            permission_group_id=foo_permission_group.id,
            subnet_id=foo_subnet.id)
        ```

        ## Import

        NasFileSystem can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:nas/fileSystem:FileSystem default enas-cnbjd3879745****
        ```

        :param str resource_name: The name of the resource.
        :param FileSystemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FileSystemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 file_system_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileSystemTagArgs']]]]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FileSystemArgs.__new__(FileSystemArgs)

            if capacity is None and not opts.urn:
                raise TypeError("Missing required property 'capacity'")
            __props__.__dict__["capacity"] = capacity
            __props__.__dict__["description"] = description
            if file_system_name is None and not opts.urn:
                raise TypeError("Missing required property 'file_system_name'")
            __props__.__dict__["file_system_name"] = file_system_name
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["snapshot_id"] = snapshot_id
            __props__.__dict__["tags"] = tags
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["charge_type"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["file_system_type"] = None
            __props__.__dict__["protocol_type"] = None
            __props__.__dict__["region_id"] = None
            __props__.__dict__["snapshot_count"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["storage_type"] = None
            __props__.__dict__["update_time"] = None
            __props__.__dict__["version"] = None
            __props__.__dict__["zone_name"] = None
        super(FileSystem, __self__).__init__(
            'volcengine:nas/fileSystem:FileSystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            capacity: Optional[pulumi.Input[int]] = None,
            charge_type: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            file_system_name: Optional[pulumi.Input[str]] = None,
            file_system_type: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            protocol_type: Optional[pulumi.Input[str]] = None,
            region_id: Optional[pulumi.Input[str]] = None,
            snapshot_count: Optional[pulumi.Input[int]] = None,
            snapshot_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            storage_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileSystemTagArgs']]]]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None,
            zone_name: Optional[pulumi.Input[str]] = None) -> 'FileSystem':
        """
        Get an existing FileSystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] capacity: The capacity of the nas file system. Unit: GiB.
        :param pulumi.Input[str] charge_type: The charge type of the nas file system.
        :param pulumi.Input[str] create_time: The create time of the nas file system.
        :param pulumi.Input[str] description: The description of the nas file system.
        :param pulumi.Input[str] file_system_name: The name of the nas file system.
        :param pulumi.Input[str] file_system_type: The type of the nas file system.
        :param pulumi.Input[str] project_name: The project name of the nas file system.
        :param pulumi.Input[str] protocol_type: The protocol type of the nas file system.
        :param pulumi.Input[str] region_id: The region id of the nas file system.
        :param pulumi.Input[int] snapshot_count: The snapshot count of the nas file system.
        :param pulumi.Input[str] snapshot_id: The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[str] status: The status of the nas file system.
        :param pulumi.Input[str] storage_type: The storage type of the nas file system.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FileSystemTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] update_time: The update time of the nas file system.
        :param pulumi.Input[str] version: The version of the nas file system.
        :param pulumi.Input[str] zone_id: The zone id of the nas file system.
        :param pulumi.Input[str] zone_name: The zone name of the nas file system.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FileSystemState.__new__(_FileSystemState)

        __props__.__dict__["capacity"] = capacity
        __props__.__dict__["charge_type"] = charge_type
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["file_system_name"] = file_system_name
        __props__.__dict__["file_system_type"] = file_system_type
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["protocol_type"] = protocol_type
        __props__.__dict__["region_id"] = region_id
        __props__.__dict__["snapshot_count"] = snapshot_count
        __props__.__dict__["snapshot_id"] = snapshot_id
        __props__.__dict__["status"] = status
        __props__.__dict__["storage_type"] = storage_type
        __props__.__dict__["tags"] = tags
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["version"] = version
        __props__.__dict__["zone_id"] = zone_id
        __props__.__dict__["zone_name"] = zone_name
        return FileSystem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[int]:
        """
        The capacity of the nas file system. Unit: GiB.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> pulumi.Output[str]:
        """
        The charge type of the nas file system.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The create time of the nas file system.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the nas file system.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> pulumi.Output[str]:
        """
        The name of the nas file system.
        """
        return pulumi.get(self, "file_system_name")

    @property
    @pulumi.getter(name="fileSystemType")
    def file_system_type(self) -> pulumi.Output[str]:
        """
        The type of the nas file system.
        """
        return pulumi.get(self, "file_system_type")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The project name of the nas file system.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> pulumi.Output[str]:
        """
        The protocol type of the nas file system.
        """
        return pulumi.get(self, "protocol_type")

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> pulumi.Output[str]:
        """
        The region id of the nas file system.
        """
        return pulumi.get(self, "region_id")

    @property
    @pulumi.getter(name="snapshotCount")
    def snapshot_count(self) -> pulumi.Output[int]:
        """
        The snapshot count of the nas file system.
        """
        return pulumi.get(self, "snapshot_count")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[Optional[str]]:
        """
        The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the nas file system.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Output[str]:
        """
        The storage type of the nas file system.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.FileSystemTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time of the nas file system.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version of the nas file system.
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The zone id of the nas file system.
        """
        return pulumi.get(self, "zone_id")

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> pulumi.Output[str]:
        """
        The zone name of the nas file system.
        """
        return pulumi.get(self, "zone_name")

