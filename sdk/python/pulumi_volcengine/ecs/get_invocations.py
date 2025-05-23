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
    'GetInvocationsResult',
    'AwaitableGetInvocationsResult',
    'get_invocations',
    'get_invocations_output',
]

@pulumi.output_type
class GetInvocationsResult:
    """
    A collection of values returned by getInvocations.
    """
    def __init__(__self__, command_id=None, command_name=None, command_type=None, id=None, invocation_id=None, invocation_name=None, invocation_statuses=None, invocations=None, name_regex=None, output_file=None, repeat_mode=None, total_count=None):
        if command_id and not isinstance(command_id, str):
            raise TypeError("Expected argument 'command_id' to be a str")
        pulumi.set(__self__, "command_id", command_id)
        if command_name and not isinstance(command_name, str):
            raise TypeError("Expected argument 'command_name' to be a str")
        pulumi.set(__self__, "command_name", command_name)
        if command_type and not isinstance(command_type, str):
            raise TypeError("Expected argument 'command_type' to be a str")
        pulumi.set(__self__, "command_type", command_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if invocation_id and not isinstance(invocation_id, str):
            raise TypeError("Expected argument 'invocation_id' to be a str")
        pulumi.set(__self__, "invocation_id", invocation_id)
        if invocation_name and not isinstance(invocation_name, str):
            raise TypeError("Expected argument 'invocation_name' to be a str")
        pulumi.set(__self__, "invocation_name", invocation_name)
        if invocation_statuses and not isinstance(invocation_statuses, list):
            raise TypeError("Expected argument 'invocation_statuses' to be a list")
        pulumi.set(__self__, "invocation_statuses", invocation_statuses)
        if invocations and not isinstance(invocations, list):
            raise TypeError("Expected argument 'invocations' to be a list")
        pulumi.set(__self__, "invocations", invocations)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if repeat_mode and not isinstance(repeat_mode, str):
            raise TypeError("Expected argument 'repeat_mode' to be a str")
        pulumi.set(__self__, "repeat_mode", repeat_mode)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="commandId")
    def command_id(self) -> Optional[str]:
        """
        The id of the ecs command.
        """
        return pulumi.get(self, "command_id")

    @property
    @pulumi.getter(name="commandName")
    def command_name(self) -> Optional[str]:
        """
        The name of the ecs command.
        """
        return pulumi.get(self, "command_name")

    @property
    @pulumi.getter(name="commandType")
    def command_type(self) -> Optional[str]:
        """
        The type of the ecs command.
        """
        return pulumi.get(self, "command_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="invocationId")
    def invocation_id(self) -> Optional[str]:
        """
        The id of the ecs invocation.
        """
        return pulumi.get(self, "invocation_id")

    @property
    @pulumi.getter(name="invocationName")
    def invocation_name(self) -> Optional[str]:
        """
        The name of the ecs invocation.
        """
        return pulumi.get(self, "invocation_name")

    @property
    @pulumi.getter(name="invocationStatuses")
    def invocation_statuses(self) -> Optional[Sequence[str]]:
        """
        The status of the ecs invocation.
        """
        return pulumi.get(self, "invocation_statuses")

    @property
    @pulumi.getter
    def invocations(self) -> Sequence['outputs.GetInvocationsInvocationResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "invocations")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="repeatMode")
    def repeat_mode(self) -> Optional[str]:
        """
        The repeat mode of the ecs invocation.
        """
        return pulumi.get(self, "repeat_mode")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetInvocationsResult(GetInvocationsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInvocationsResult(
            command_id=self.command_id,
            command_name=self.command_name,
            command_type=self.command_type,
            id=self.id,
            invocation_id=self.invocation_id,
            invocation_name=self.invocation_name,
            invocation_statuses=self.invocation_statuses,
            invocations=self.invocations,
            name_regex=self.name_regex,
            output_file=self.output_file,
            repeat_mode=self.repeat_mode,
            total_count=self.total_count)


def get_invocations(command_id: Optional[str] = None,
                    command_name: Optional[str] = None,
                    command_type: Optional[str] = None,
                    invocation_id: Optional[str] = None,
                    invocation_name: Optional[str] = None,
                    invocation_statuses: Optional[Sequence[str]] = None,
                    name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    repeat_mode: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInvocationsResult:
    """
    Use this data source to query detailed information of ecs invocations
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.ecs.get_invocations(invocation_id="ivk-ych9y4vujvl8j01c****",
        invocation_statuses=["Success"])
    ```


    :param str command_id: The id of ecs command.
    :param str command_name: The name of ecs command. This field support fuzzy query.
    :param str command_type: The type of ecs command. Valid values: `Shell`.
    :param str invocation_id: The id of ecs invocation.
    :param str invocation_name: The name of ecs invocation. This field support fuzzy query.
    :param Sequence[str] invocation_statuses: The list of status of ecs invocation. Valid values: `Pending`, `Scheduled`, `Running`, `Success`, `Failed`, `Stopped`, `PartialFailed`, `Finished`.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str repeat_mode: The repeat mode of ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.
    """
    __args__ = dict()
    __args__['commandId'] = command_id
    __args__['commandName'] = command_name
    __args__['commandType'] = command_type
    __args__['invocationId'] = invocation_id
    __args__['invocationName'] = invocation_name
    __args__['invocationStatuses'] = invocation_statuses
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['repeatMode'] = repeat_mode
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:ecs/getInvocations:getInvocations', __args__, opts=opts, typ=GetInvocationsResult).value

    return AwaitableGetInvocationsResult(
        command_id=pulumi.get(__ret__, 'command_id'),
        command_name=pulumi.get(__ret__, 'command_name'),
        command_type=pulumi.get(__ret__, 'command_type'),
        id=pulumi.get(__ret__, 'id'),
        invocation_id=pulumi.get(__ret__, 'invocation_id'),
        invocation_name=pulumi.get(__ret__, 'invocation_name'),
        invocation_statuses=pulumi.get(__ret__, 'invocation_statuses'),
        invocations=pulumi.get(__ret__, 'invocations'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        repeat_mode=pulumi.get(__ret__, 'repeat_mode'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_invocations)
def get_invocations_output(command_id: Optional[pulumi.Input[Optional[str]]] = None,
                           command_name: Optional[pulumi.Input[Optional[str]]] = None,
                           command_type: Optional[pulumi.Input[Optional[str]]] = None,
                           invocation_id: Optional[pulumi.Input[Optional[str]]] = None,
                           invocation_name: Optional[pulumi.Input[Optional[str]]] = None,
                           invocation_statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           repeat_mode: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInvocationsResult]:
    """
    Use this data source to query detailed information of ecs invocations
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.ecs.get_invocations(invocation_id="ivk-ych9y4vujvl8j01c****",
        invocation_statuses=["Success"])
    ```


    :param str command_id: The id of ecs command.
    :param str command_name: The name of ecs command. This field support fuzzy query.
    :param str command_type: The type of ecs command. Valid values: `Shell`.
    :param str invocation_id: The id of ecs invocation.
    :param str invocation_name: The name of ecs invocation. This field support fuzzy query.
    :param Sequence[str] invocation_statuses: The list of status of ecs invocation. Valid values: `Pending`, `Scheduled`, `Running`, `Success`, `Failed`, `Stopped`, `PartialFailed`, `Finished`.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str repeat_mode: The repeat mode of ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.
    """
    ...
