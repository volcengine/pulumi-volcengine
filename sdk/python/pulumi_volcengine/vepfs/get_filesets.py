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
    'GetFilesetsResult',
    'AwaitableGetFilesetsResult',
    'get_filesets',
    'get_filesets_output',
]

@pulumi.output_type
class GetFilesetsResult:
    """
    A collection of values returned by getFilesets.
    """
    def __init__(__self__, file_system_id=None, fileset_id=None, fileset_name=None, fileset_path=None, filesets=None, id=None, name_regex=None, output_file=None, statuses=None, total_count=None):
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        pulumi.set(__self__, "file_system_id", file_system_id)
        if fileset_id and not isinstance(fileset_id, str):
            raise TypeError("Expected argument 'fileset_id' to be a str")
        pulumi.set(__self__, "fileset_id", fileset_id)
        if fileset_name and not isinstance(fileset_name, str):
            raise TypeError("Expected argument 'fileset_name' to be a str")
        pulumi.set(__self__, "fileset_name", fileset_name)
        if fileset_path and not isinstance(fileset_path, str):
            raise TypeError("Expected argument 'fileset_path' to be a str")
        pulumi.set(__self__, "fileset_path", fileset_path)
        if filesets and not isinstance(filesets, list):
            raise TypeError("Expected argument 'filesets' to be a list")
        pulumi.set(__self__, "filesets", filesets)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter(name="filesetId")
    def fileset_id(self) -> Optional[str]:
        """
        The id of the vepfs fileset.
        """
        return pulumi.get(self, "fileset_id")

    @property
    @pulumi.getter(name="filesetName")
    def fileset_name(self) -> Optional[str]:
        """
        The name of the vepfs fileset.
        """
        return pulumi.get(self, "fileset_name")

    @property
    @pulumi.getter(name="filesetPath")
    def fileset_path(self) -> Optional[str]:
        """
        The path of the vepfs fileset.
        """
        return pulumi.get(self, "fileset_path")

    @property
    @pulumi.getter
    def filesets(self) -> Sequence['outputs.GetFilesetsFilesetResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "filesets")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

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
    def statuses(self) -> Optional[Sequence[str]]:
        """
        The status of the vepfs fileset.
        """
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetFilesetsResult(GetFilesetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFilesetsResult(
            file_system_id=self.file_system_id,
            fileset_id=self.fileset_id,
            fileset_name=self.fileset_name,
            fileset_path=self.fileset_path,
            filesets=self.filesets,
            id=self.id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            statuses=self.statuses,
            total_count=self.total_count)


def get_filesets(file_system_id: Optional[str] = None,
                 fileset_id: Optional[str] = None,
                 fileset_name: Optional[str] = None,
                 fileset_path: Optional[str] = None,
                 name_regex: Optional[str] = None,
                 output_file: Optional[str] = None,
                 statuses: Optional[Sequence[str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFilesetsResult:
    """
    Use this data source to query detailed information of vepfs filesets
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
    foo_fileset = volcengine.vepfs.Fileset("fooFileset",
        file_system_id=foo_file_system.id,
        fileset_name="acc-test-fileset",
        fileset_path="/tf-test/",
        max_iops=100,
        max_bandwidth=10,
        file_limit=20,
        capacity_limit=30)
    foo_filesets = volcengine.vepfs.get_filesets_output(file_system_id=foo_file_system.id,
        fileset_id=foo_fileset.id)
    ```


    :param str file_system_id: The id of Vepfs File System.
    :param str fileset_id: The id of Vepfs Fileset.
    :param str fileset_name: The name of Vepfs Fileset. This field support fuzzy query.
    :param str fileset_path: The path of Vepfs Fileset. This field support fuzzy query.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] statuses: The query status list of Vepfs Fileset.
    """
    __args__ = dict()
    __args__['fileSystemId'] = file_system_id
    __args__['filesetId'] = fileset_id
    __args__['filesetName'] = fileset_name
    __args__['filesetPath'] = fileset_path
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['statuses'] = statuses
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vepfs/getFilesets:getFilesets', __args__, opts=opts, typ=GetFilesetsResult).value

    return AwaitableGetFilesetsResult(
        file_system_id=pulumi.get(__ret__, 'file_system_id'),
        fileset_id=pulumi.get(__ret__, 'fileset_id'),
        fileset_name=pulumi.get(__ret__, 'fileset_name'),
        fileset_path=pulumi.get(__ret__, 'fileset_path'),
        filesets=pulumi.get(__ret__, 'filesets'),
        id=pulumi.get(__ret__, 'id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        statuses=pulumi.get(__ret__, 'statuses'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_filesets)
def get_filesets_output(file_system_id: Optional[pulumi.Input[str]] = None,
                        fileset_id: Optional[pulumi.Input[Optional[str]]] = None,
                        fileset_name: Optional[pulumi.Input[Optional[str]]] = None,
                        fileset_path: Optional[pulumi.Input[Optional[str]]] = None,
                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFilesetsResult]:
    """
    Use this data source to query detailed information of vepfs filesets
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
    foo_fileset = volcengine.vepfs.Fileset("fooFileset",
        file_system_id=foo_file_system.id,
        fileset_name="acc-test-fileset",
        fileset_path="/tf-test/",
        max_iops=100,
        max_bandwidth=10,
        file_limit=20,
        capacity_limit=30)
    foo_filesets = volcengine.vepfs.get_filesets_output(file_system_id=foo_file_system.id,
        fileset_id=foo_fileset.id)
    ```


    :param str file_system_id: The id of Vepfs File System.
    :param str fileset_id: The id of Vepfs Fileset.
    :param str fileset_name: The name of Vepfs Fileset. This field support fuzzy query.
    :param str fileset_path: The path of Vepfs Fileset. This field support fuzzy query.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] statuses: The query status list of Vepfs Fileset.
    """
    ...
