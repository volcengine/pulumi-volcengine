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

__all__ = [
    'FileSystemsResult',
    'AwaitableFileSystemsResult',
    'file_systems',
    'file_systems_output',
]

@pulumi.output_type
class FileSystemsResult:
    """
    A collection of values returned by FileSystems.
    """
    def __init__(__self__, charge_type=None, file_system_name=None, file_systems=None, id=None, ids=None, mount_point_id=None, name_regex=None, output_file=None, permission_group_id=None, project_name=None, protocol_type=None, statuses=None, storage_type=None, tags=None, total_count=None, zone_id=None):
        if charge_type and not isinstance(charge_type, str):
            raise TypeError("Expected argument 'charge_type' to be a str")
        pulumi.set(__self__, "charge_type", charge_type)
        if file_system_name and not isinstance(file_system_name, str):
            raise TypeError("Expected argument 'file_system_name' to be a str")
        pulumi.set(__self__, "file_system_name", file_system_name)
        if file_systems and not isinstance(file_systems, list):
            raise TypeError("Expected argument 'file_systems' to be a list")
        pulumi.set(__self__, "file_systems", file_systems)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if mount_point_id and not isinstance(mount_point_id, str):
            raise TypeError("Expected argument 'mount_point_id' to be a str")
        pulumi.set(__self__, "mount_point_id", mount_point_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if permission_group_id and not isinstance(permission_group_id, str):
            raise TypeError("Expected argument 'permission_group_id' to be a str")
        pulumi.set(__self__, "permission_group_id", permission_group_id)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if protocol_type and not isinstance(protocol_type, str):
            raise TypeError("Expected argument 'protocol_type' to be a str")
        pulumi.set(__self__, "protocol_type", protocol_type)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if storage_type and not isinstance(storage_type, str):
            raise TypeError("Expected argument 'storage_type' to be a str")
        pulumi.set(__self__, "storage_type", storage_type)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[str]:
        """
        The charge type of the nas file system.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> Optional[str]:
        """
        The name of the nas file system.
        """
        return pulumi.get(self, "file_system_name")

    @property
    @pulumi.getter(name="fileSystems")
    def file_systems(self) -> Sequence['outputs.FileSystemsFileSystemResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "file_systems")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="mountPointId")
    def mount_point_id(self) -> Optional[str]:
        return pulumi.get(self, "mount_point_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="permissionGroupId")
    def permission_group_id(self) -> Optional[str]:
        return pulumi.get(self, "permission_group_id")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The project name of the nas file system.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="protocolType")
    def protocol_type(self) -> Optional[str]:
        """
        The protocol type of the nas file system.
        """
        return pulumi.get(self, "protocol_type")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        """
        The status of the nas file system.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> Optional[str]:
        """
        The storage type of the nas file system.
        """
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.FileSystemsTagResult']]:
        """
        Tags of the nas file system.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[str]:
        """
        The zone id of the nas file system.
        """
        return pulumi.get(self, "zone_id")


class AwaitableFileSystemsResult(FileSystemsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return FileSystemsResult(
            charge_type=self.charge_type,
            file_system_name=self.file_system_name,
            file_systems=self.file_systems,
            id=self.id,
            ids=self.ids,
            mount_point_id=self.mount_point_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            permission_group_id=self.permission_group_id,
            project_name=self.project_name,
            protocol_type=self.protocol_type,
            statuses=self.statuses,
            storage_type=self.storage_type,
            tags=self.tags,
            total_count=self.total_count,
            zone_id=self.zone_id)


def file_systems(charge_type: Optional[str] = None,
                 file_system_name: Optional[str] = None,
                 ids: Optional[Sequence[str]] = None,
                 mount_point_id: Optional[str] = None,
                 name_regex: Optional[str] = None,
                 output_file: Optional[str] = None,
                 permission_group_id: Optional[str] = None,
                 project_name: Optional[str] = None,
                 protocol_type: Optional[str] = None,
                 statuses: Optional[Sequence[str]] = None,
                 storage_type: Optional[str] = None,
                 tags: Optional[Sequence[pulumi.InputType['FileSystemsTagArgs']]] = None,
                 zone_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableFileSystemsResult:
    """
    Use this data source to query detailed information of nas file systems
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.nas.zones()
    foo_file_system = []
    for range in [{"value": i} for i in range(0, 3)]:
        foo_file_system.append(volcengine.nas.FileSystem(f"fooFileSystem-{range['value']}",
            file_system_name=f"acc-test-fs-{range['value']}",
            description="acc-test",
            zone_id=foo_zones.zones[0].id,
            capacity=103,
            project_name="default",
            tags=[volcengine.nas.FileSystemTagArgs(
                key="k1",
                value="v1",
            )]))
    foo_file_systems = volcengine.nas.file_systems_output(ids=[__item.id for __item in foo_file_system])
    ```


    :param str charge_type: The charge type of nas file system.
    :param str file_system_name: The name of nas file system. This field supports fuzzy queries.
    :param Sequence[str] ids: A list of nas file system ids.
    :param str mount_point_id: The mount point id of nas file system.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str permission_group_id: The permission group id of nas file system.
    :param str project_name: The project name of nas file system.
    :param str protocol_type: The protocol type of nas file system.
    :param Sequence[str] statuses: The status of nas file system.
    :param str storage_type: The storage type of nas file system.
    :param Sequence[pulumi.InputType['FileSystemsTagArgs']] tags: Tags.
    :param str zone_id: The zone id of nas file system.
    """
    __args__ = dict()
    __args__['chargeType'] = charge_type
    __args__['fileSystemName'] = file_system_name
    __args__['ids'] = ids
    __args__['mountPointId'] = mount_point_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['permissionGroupId'] = permission_group_id
    __args__['projectName'] = project_name
    __args__['protocolType'] = protocol_type
    __args__['statuses'] = statuses
    __args__['storageType'] = storage_type
    __args__['tags'] = tags
    __args__['zoneId'] = zone_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:nas/fileSystems:FileSystems', __args__, opts=opts, typ=FileSystemsResult).value

    return AwaitableFileSystemsResult(
        charge_type=pulumi.get(__ret__, 'charge_type'),
        file_system_name=pulumi.get(__ret__, 'file_system_name'),
        file_systems=pulumi.get(__ret__, 'file_systems'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        mount_point_id=pulumi.get(__ret__, 'mount_point_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        permission_group_id=pulumi.get(__ret__, 'permission_group_id'),
        project_name=pulumi.get(__ret__, 'project_name'),
        protocol_type=pulumi.get(__ret__, 'protocol_type'),
        statuses=pulumi.get(__ret__, 'statuses'),
        storage_type=pulumi.get(__ret__, 'storage_type'),
        tags=pulumi.get(__ret__, 'tags'),
        total_count=pulumi.get(__ret__, 'total_count'),
        zone_id=pulumi.get(__ret__, 'zone_id'))


@_utilities.lift_output_func(file_systems)
def file_systems_output(charge_type: Optional[pulumi.Input[Optional[str]]] = None,
                        file_system_name: Optional[pulumi.Input[Optional[str]]] = None,
                        ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        mount_point_id: Optional[pulumi.Input[Optional[str]]] = None,
                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        permission_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                        project_name: Optional[pulumi.Input[Optional[str]]] = None,
                        protocol_type: Optional[pulumi.Input[Optional[str]]] = None,
                        statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        storage_type: Optional[pulumi.Input[Optional[str]]] = None,
                        tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['FileSystemsTagArgs']]]]] = None,
                        zone_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[FileSystemsResult]:
    """
    Use this data source to query detailed information of nas file systems
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.nas.zones()
    foo_file_system = []
    for range in [{"value": i} for i in range(0, 3)]:
        foo_file_system.append(volcengine.nas.FileSystem(f"fooFileSystem-{range['value']}",
            file_system_name=f"acc-test-fs-{range['value']}",
            description="acc-test",
            zone_id=foo_zones.zones[0].id,
            capacity=103,
            project_name="default",
            tags=[volcengine.nas.FileSystemTagArgs(
                key="k1",
                value="v1",
            )]))
    foo_file_systems = volcengine.nas.file_systems_output(ids=[__item.id for __item in foo_file_system])
    ```


    :param str charge_type: The charge type of nas file system.
    :param str file_system_name: The name of nas file system. This field supports fuzzy queries.
    :param Sequence[str] ids: A list of nas file system ids.
    :param str mount_point_id: The mount point id of nas file system.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str permission_group_id: The permission group id of nas file system.
    :param str project_name: The project name of nas file system.
    :param str protocol_type: The protocol type of nas file system.
    :param Sequence[str] statuses: The status of nas file system.
    :param str storage_type: The storage type of nas file system.
    :param Sequence[pulumi.InputType['FileSystemsTagArgs']] tags: Tags.
    :param str zone_id: The zone id of nas file system.
    """
    ...