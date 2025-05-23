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
    'GetHostGroupsResult',
    'AwaitableGetHostGroupsResult',
    'get_host_groups',
    'get_host_groups_output',
]

@pulumi.output_type
class GetHostGroupsResult:
    """
    A collection of values returned by getHostGroups.
    """
    def __init__(__self__, auto_update=None, host_group_id=None, host_group_name=None, host_identifier=None, iam_project_name=None, id=None, infos=None, output_file=None, service_logging=None, total_count=None):
        if auto_update and not isinstance(auto_update, bool):
            raise TypeError("Expected argument 'auto_update' to be a bool")
        pulumi.set(__self__, "auto_update", auto_update)
        if host_group_id and not isinstance(host_group_id, str):
            raise TypeError("Expected argument 'host_group_id' to be a str")
        pulumi.set(__self__, "host_group_id", host_group_id)
        if host_group_name and not isinstance(host_group_name, str):
            raise TypeError("Expected argument 'host_group_name' to be a str")
        pulumi.set(__self__, "host_group_name", host_group_name)
        if host_identifier and not isinstance(host_identifier, str):
            raise TypeError("Expected argument 'host_identifier' to be a str")
        pulumi.set(__self__, "host_identifier", host_identifier)
        if iam_project_name and not isinstance(iam_project_name, str):
            raise TypeError("Expected argument 'iam_project_name' to be a str")
        pulumi.set(__self__, "iam_project_name", iam_project_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if infos and not isinstance(infos, list):
            raise TypeError("Expected argument 'infos' to be a list")
        pulumi.set(__self__, "infos", infos)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if service_logging and not isinstance(service_logging, bool):
            raise TypeError("Expected argument 'service_logging' to be a bool")
        pulumi.set(__self__, "service_logging", service_logging)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="autoUpdate")
    def auto_update(self) -> Optional[bool]:
        """
        Whether enable auto update.
        """
        return pulumi.get(self, "auto_update")

    @property
    @pulumi.getter(name="hostGroupId")
    def host_group_id(self) -> Optional[str]:
        """
        The id of host group.
        """
        return pulumi.get(self, "host_group_id")

    @property
    @pulumi.getter(name="hostGroupName")
    def host_group_name(self) -> Optional[str]:
        """
        The name of host group.
        """
        return pulumi.get(self, "host_group_name")

    @property
    @pulumi.getter(name="hostIdentifier")
    def host_identifier(self) -> Optional[str]:
        """
        The identifier of host.
        """
        return pulumi.get(self, "host_identifier")

    @property
    @pulumi.getter(name="iamProjectName")
    def iam_project_name(self) -> Optional[str]:
        """
        The project name of iam.
        """
        return pulumi.get(self, "iam_project_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def infos(self) -> Sequence['outputs.GetHostGroupsInfoResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "infos")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="serviceLogging")
    def service_logging(self) -> Optional[bool]:
        """
        Whether enable service logging.
        """
        return pulumi.get(self, "service_logging")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetHostGroupsResult(GetHostGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostGroupsResult(
            auto_update=self.auto_update,
            host_group_id=self.host_group_id,
            host_group_name=self.host_group_name,
            host_identifier=self.host_identifier,
            iam_project_name=self.iam_project_name,
            id=self.id,
            infos=self.infos,
            output_file=self.output_file,
            service_logging=self.service_logging,
            total_count=self.total_count)


def get_host_groups(auto_update: Optional[bool] = None,
                    host_group_id: Optional[str] = None,
                    host_group_name: Optional[str] = None,
                    host_identifier: Optional[str] = None,
                    iam_project_name: Optional[str] = None,
                    output_file: Optional[str] = None,
                    service_logging: Optional[bool] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostGroupsResult:
    """
    Use this data source to query detailed information of tls host groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.get_host_groups(host_group_id="fbea6619-7b0c-40f3-ac7e-45c63e3f676e",
        host_group_name="cn")
    ```


    :param bool auto_update: Whether enable auto update.
    :param str host_group_id: The id of host group.
    :param str host_group_name: The name of host group.
    :param str host_identifier: The identifier of host.
    :param str iam_project_name: The project name of iam.
    :param str output_file: File name where to save data source results.
    :param bool service_logging: Whether enable service logging.
    """
    __args__ = dict()
    __args__['autoUpdate'] = auto_update
    __args__['hostGroupId'] = host_group_id
    __args__['hostGroupName'] = host_group_name
    __args__['hostIdentifier'] = host_identifier
    __args__['iamProjectName'] = iam_project_name
    __args__['outputFile'] = output_file
    __args__['serviceLogging'] = service_logging
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:tls/getHostGroups:getHostGroups', __args__, opts=opts, typ=GetHostGroupsResult).value

    return AwaitableGetHostGroupsResult(
        auto_update=pulumi.get(__ret__, 'auto_update'),
        host_group_id=pulumi.get(__ret__, 'host_group_id'),
        host_group_name=pulumi.get(__ret__, 'host_group_name'),
        host_identifier=pulumi.get(__ret__, 'host_identifier'),
        iam_project_name=pulumi.get(__ret__, 'iam_project_name'),
        id=pulumi.get(__ret__, 'id'),
        infos=pulumi.get(__ret__, 'infos'),
        output_file=pulumi.get(__ret__, 'output_file'),
        service_logging=pulumi.get(__ret__, 'service_logging'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_host_groups)
def get_host_groups_output(auto_update: Optional[pulumi.Input[Optional[bool]]] = None,
                           host_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                           host_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                           host_identifier: Optional[pulumi.Input[Optional[str]]] = None,
                           iam_project_name: Optional[pulumi.Input[Optional[str]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           service_logging: Optional[pulumi.Input[Optional[bool]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHostGroupsResult]:
    """
    Use this data source to query detailed information of tls host groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.get_host_groups(host_group_id="fbea6619-7b0c-40f3-ac7e-45c63e3f676e",
        host_group_name="cn")
    ```


    :param bool auto_update: Whether enable auto update.
    :param str host_group_id: The id of host group.
    :param str host_group_name: The name of host group.
    :param str host_identifier: The identifier of host.
    :param str iam_project_name: The project name of iam.
    :param str output_file: File name where to save data source results.
    :param bool service_logging: Whether enable service logging.
    """
    ...
