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
    'MountPointsResult',
    'AwaitableMountPointsResult',
    'mount_points',
    'mount_points_output',
]

@pulumi.output_type
class MountPointsResult:
    """
    A collection of values returned by MountPoints.
    """
    def __init__(__self__, file_system_id=None, id=None, mount_point_id=None, mount_point_name=None, mount_points=None, output_file=None, total_count=None, vpcs_id=None):
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        pulumi.set(__self__, "file_system_id", file_system_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mount_point_id and not isinstance(mount_point_id, str):
            raise TypeError("Expected argument 'mount_point_id' to be a str")
        pulumi.set(__self__, "mount_point_id", mount_point_id)
        if mount_point_name and not isinstance(mount_point_name, str):
            raise TypeError("Expected argument 'mount_point_name' to be a str")
        pulumi.set(__self__, "mount_point_name", mount_point_name)
        if mount_points and not isinstance(mount_points, list):
            raise TypeError("Expected argument 'mount_points' to be a list")
        pulumi.set(__self__, "mount_points", mount_points)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpcs_id and not isinstance(vpcs_id, str):
            raise TypeError("Expected argument 'vpcs_id' to be a str")
        pulumi.set(__self__, "vpcs_id", vpcs_id)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        """
        The id of the file system.
        """
        return pulumi.get(self, "file_system_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mountPointId")
    def mount_point_id(self) -> Optional[str]:
        """
        The id of the mount point.
        """
        return pulumi.get(self, "mount_point_id")

    @property
    @pulumi.getter(name="mountPointName")
    def mount_point_name(self) -> Optional[str]:
        """
        The name of the mount point.
        """
        return pulumi.get(self, "mount_point_name")

    @property
    @pulumi.getter(name="mountPoints")
    def mount_points(self) -> Sequence['outputs.MountPointsMountPointResult']:
        """
        The list of the mount point.
        """
        return pulumi.get(self, "mount_points")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of nas mount points query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcsId")
    def vpcs_id(self) -> Optional[str]:
        return pulumi.get(self, "vpcs_id")


class AwaitableMountPointsResult(MountPointsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return MountPointsResult(
            file_system_id=self.file_system_id,
            id=self.id,
            mount_point_id=self.mount_point_id,
            mount_point_name=self.mount_point_name,
            mount_points=self.mount_points,
            output_file=self.output_file,
            total_count=self.total_count,
            vpcs_id=self.vpcs_id)


def mount_points(file_system_id: Optional[str] = None,
                 mount_point_id: Optional[str] = None,
                 mount_point_name: Optional[str] = None,
                 output_file: Optional[str] = None,
                 vpcs_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableMountPointsResult:
    """
    Use this data source to query detailed information of nas mount points


    :param str file_system_id: The id of the file system.
    :param str mount_point_id: The id of the mount point.
    :param str mount_point_name: The name of the mount point.
    :param str output_file: File name where to save data source results.
    :param str vpcs_id: The id of the vpc.
    """
    __args__ = dict()
    __args__['fileSystemId'] = file_system_id
    __args__['mountPointId'] = mount_point_id
    __args__['mountPointName'] = mount_point_name
    __args__['outputFile'] = output_file
    __args__['vpcsId'] = vpcs_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:nas/mountPoints:MountPoints', __args__, opts=opts, typ=MountPointsResult).value

    return AwaitableMountPointsResult(
        file_system_id=pulumi.get(__ret__, 'file_system_id'),
        id=pulumi.get(__ret__, 'id'),
        mount_point_id=pulumi.get(__ret__, 'mount_point_id'),
        mount_point_name=pulumi.get(__ret__, 'mount_point_name'),
        mount_points=pulumi.get(__ret__, 'mount_points'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpcs_id=pulumi.get(__ret__, 'vpcs_id'))


@_utilities.lift_output_func(mount_points)
def mount_points_output(file_system_id: Optional[pulumi.Input[str]] = None,
                        mount_point_id: Optional[pulumi.Input[Optional[str]]] = None,
                        mount_point_name: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        vpcs_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[MountPointsResult]:
    """
    Use this data source to query detailed information of nas mount points


    :param str file_system_id: The id of the file system.
    :param str mount_point_id: The id of the mount point.
    :param str mount_point_name: The name of the mount point.
    :param str output_file: File name where to save data source results.
    :param str vpcs_id: The id of the vpc.
    """
    ...
