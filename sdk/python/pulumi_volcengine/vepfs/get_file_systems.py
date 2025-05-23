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

__all__ = [
    'GetFileSystemsResult',
    'AwaitableGetFileSystemsResult',
    'get_file_systems',
    'get_file_systems_output',
]

@pulumi.output_type
class GetFileSystemsResult:
    """
    A collection of values returned by getFileSystems.
    """
    def __init__(__self__, file_system_name=None, file_systems=None, id=None, ids=None, name_regex=None, output_file=None, project=None, statuses=None, store_type=None, total_count=None, zone_id=None):
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
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if store_type and not isinstance(store_type, str):
            raise TypeError("Expected argument 'store_type' to be a str")
        pulumi.set(__self__, "store_type", store_type)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if zone_id and not isinstance(zone_id, str):
            raise TypeError("Expected argument 'zone_id' to be a str")
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="fileSystemName")
    def file_system_name(self) -> Optional[str]:
        """
        The name of the vepfs file system.
        """
        return pulumi.get(self, "file_system_name")

    @property
    @pulumi.getter(name="fileSystems")
    def file_systems(self) -> Sequence['outputs.GetFileSystemsFileSystemResult']:
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
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def project(self) -> Optional[str]:
        """
        The project name of the vepfs file system.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        """
        The status of the vepfs file system.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="storeType")
    def store_type(self) -> Optional[str]:
        """
        The store type of the vepfs file system.
        """
        return pulumi.get(self, "store_type")

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
        The id of the zone.
        """
        return pulumi.get(self, "zone_id")


class AwaitableGetFileSystemsResult(GetFileSystemsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileSystemsResult(
            file_system_name=self.file_system_name,
            file_systems=self.file_systems,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            project=self.project,
            statuses=self.statuses,
            store_type=self.store_type,
            total_count=self.total_count,
            zone_id=self.zone_id)


def get_file_systems(file_system_name: Optional[str] = None,
                     ids: Optional[Sequence[str]] = None,
                     name_regex: Optional[str] = None,
                     output_file: Optional[str] = None,
                     project: Optional[str] = None,
                     statuses: Optional[Sequence[str]] = None,
                     store_type: Optional[str] = None,
                     zone_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileSystemsResult:
    """
    Use this data source to query detailed information of vepfs file systems
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id="cn-beijing-a",
        vpc_id=foo_vpc.id)
    foo_file_system = volcengine.vepfs.FileSystem("fooFileSystem",
        file_system_name="acc-test-file-system",
        subnet_id=foo_subnet.id,
        store_type="Advance_100",
        description="tf-test",
        capacity=12,
        project="default",
        enable_restripe=False,
        tags=[volcengine.vepfs.FileSystemTagArgs(
            key="k1",
            value="v1",
        )])
    foo_file_systems = volcengine.vepfs.get_file_systems_output(ids=[foo_file_system.id])
    ```


    :param str file_system_name: The Name of Vepfs File System. This field support fuzzy query.
    :param Sequence[str] ids: A list of Vepfs File System IDs.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str project: The project of Vepfs File System.
    :param Sequence[str] statuses: The query status list of Vepfs File System.
    :param str store_type: The Store Type of Vepfs File System.
    :param str zone_id: The zone id of File System.
    """
    __args__ = dict()
    __args__['fileSystemName'] = file_system_name
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['project'] = project
    __args__['statuses'] = statuses
    __args__['storeType'] = store_type
    __args__['zoneId'] = zone_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vepfs/getFileSystems:getFileSystems', __args__, opts=opts, typ=GetFileSystemsResult).value

    return AwaitableGetFileSystemsResult(
        file_system_name=pulumi.get(__ret__, 'file_system_name'),
        file_systems=pulumi.get(__ret__, 'file_systems'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        project=pulumi.get(__ret__, 'project'),
        statuses=pulumi.get(__ret__, 'statuses'),
        store_type=pulumi.get(__ret__, 'store_type'),
        total_count=pulumi.get(__ret__, 'total_count'),
        zone_id=pulumi.get(__ret__, 'zone_id'))


@_utilities.lift_output_func(get_file_systems)
def get_file_systems_output(file_system_name: Optional[pulumi.Input[Optional[str]]] = None,
                            ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            store_type: Optional[pulumi.Input[Optional[str]]] = None,
                            zone_id: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFileSystemsResult]:
    """
    Use this data source to query detailed information of vepfs file systems
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id="cn-beijing-a",
        vpc_id=foo_vpc.id)
    foo_file_system = volcengine.vepfs.FileSystem("fooFileSystem",
        file_system_name="acc-test-file-system",
        subnet_id=foo_subnet.id,
        store_type="Advance_100",
        description="tf-test",
        capacity=12,
        project="default",
        enable_restripe=False,
        tags=[volcengine.vepfs.FileSystemTagArgs(
            key="k1",
            value="v1",
        )])
    foo_file_systems = volcengine.vepfs.get_file_systems_output(ids=[foo_file_system.id])
    ```


    :param str file_system_name: The Name of Vepfs File System. This field support fuzzy query.
    :param Sequence[str] ids: A list of Vepfs File System IDs.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str project: The project of Vepfs File System.
    :param Sequence[str] statuses: The query status list of Vepfs File System.
    :param str store_type: The Store Type of Vepfs File System.
    :param str zone_id: The zone id of File System.
    """
    ...
