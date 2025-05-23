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
    'GetAlarmNotifyGroupsResult',
    'AwaitableGetAlarmNotifyGroupsResult',
    'get_alarm_notify_groups',
    'get_alarm_notify_groups_output',
]

@pulumi.output_type
class GetAlarmNotifyGroupsResult:
    """
    A collection of values returned by getAlarmNotifyGroups.
    """
    def __init__(__self__, alarm_notify_group_id=None, alarm_notify_group_name=None, groups=None, iam_project_name=None, id=None, output_file=None, receiver_name=None, total_count=None):
        if alarm_notify_group_id and not isinstance(alarm_notify_group_id, str):
            raise TypeError("Expected argument 'alarm_notify_group_id' to be a str")
        pulumi.set(__self__, "alarm_notify_group_id", alarm_notify_group_id)
        if alarm_notify_group_name and not isinstance(alarm_notify_group_name, str):
            raise TypeError("Expected argument 'alarm_notify_group_name' to be a str")
        pulumi.set(__self__, "alarm_notify_group_name", alarm_notify_group_name)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if iam_project_name and not isinstance(iam_project_name, str):
            raise TypeError("Expected argument 'iam_project_name' to be a str")
        pulumi.set(__self__, "iam_project_name", iam_project_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if receiver_name and not isinstance(receiver_name, str):
            raise TypeError("Expected argument 'receiver_name' to be a str")
        pulumi.set(__self__, "receiver_name", receiver_name)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="alarmNotifyGroupId")
    def alarm_notify_group_id(self) -> Optional[str]:
        """
        The id of the notify group.
        """
        return pulumi.get(self, "alarm_notify_group_id")

    @property
    @pulumi.getter(name="alarmNotifyGroupName")
    def alarm_notify_group_name(self) -> Optional[str]:
        """
        Name of the notification group.
        """
        return pulumi.get(self, "alarm_notify_group_name")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetAlarmNotifyGroupsGroupResult']:
        """
        The list of the notify groups.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="iamProjectName")
    def iam_project_name(self) -> Optional[str]:
        """
        The iam project name.
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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="receiverName")
    def receiver_name(self) -> Optional[str]:
        return pulumi.get(self, "receiver_name")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetAlarmNotifyGroupsResult(GetAlarmNotifyGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAlarmNotifyGroupsResult(
            alarm_notify_group_id=self.alarm_notify_group_id,
            alarm_notify_group_name=self.alarm_notify_group_name,
            groups=self.groups,
            iam_project_name=self.iam_project_name,
            id=self.id,
            output_file=self.output_file,
            receiver_name=self.receiver_name,
            total_count=self.total_count)


def get_alarm_notify_groups(alarm_notify_group_id: Optional[str] = None,
                            alarm_notify_group_name: Optional[str] = None,
                            iam_project_name: Optional[str] = None,
                            output_file: Optional[str] = None,
                            receiver_name: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAlarmNotifyGroupsResult:
    """
    Use this data source to query detailed information of tls alarm notify groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.get_alarm_notify_groups()
    ```


    :param str alarm_notify_group_id: The id of the alarm notify group.
    :param str alarm_notify_group_name: The name of the alarm notify group.
    :param str iam_project_name: The name of the iam project.
    :param str output_file: File name where to save data source results.
    :param str receiver_name: The name of the receiver.
    """
    __args__ = dict()
    __args__['alarmNotifyGroupId'] = alarm_notify_group_id
    __args__['alarmNotifyGroupName'] = alarm_notify_group_name
    __args__['iamProjectName'] = iam_project_name
    __args__['outputFile'] = output_file
    __args__['receiverName'] = receiver_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:tls/getAlarmNotifyGroups:getAlarmNotifyGroups', __args__, opts=opts, typ=GetAlarmNotifyGroupsResult).value

    return AwaitableGetAlarmNotifyGroupsResult(
        alarm_notify_group_id=pulumi.get(__ret__, 'alarm_notify_group_id'),
        alarm_notify_group_name=pulumi.get(__ret__, 'alarm_notify_group_name'),
        groups=pulumi.get(__ret__, 'groups'),
        iam_project_name=pulumi.get(__ret__, 'iam_project_name'),
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        receiver_name=pulumi.get(__ret__, 'receiver_name'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_alarm_notify_groups)
def get_alarm_notify_groups_output(alarm_notify_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                   alarm_notify_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                                   iam_project_name: Optional[pulumi.Input[Optional[str]]] = None,
                                   output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   receiver_name: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAlarmNotifyGroupsResult]:
    """
    Use this data source to query detailed information of tls alarm notify groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.get_alarm_notify_groups()
    ```


    :param str alarm_notify_group_id: The id of the alarm notify group.
    :param str alarm_notify_group_name: The name of the alarm notify group.
    :param str iam_project_name: The name of the iam project.
    :param str output_file: File name where to save data source results.
    :param str receiver_name: The name of the receiver.
    """
    ...
