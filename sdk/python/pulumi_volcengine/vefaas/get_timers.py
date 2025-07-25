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
    'GetTimersResult',
    'AwaitableGetTimersResult',
    'get_timers',
    'get_timers_output',
]

@pulumi.output_type
class GetTimersResult:
    """
    A collection of values returned by getTimers.
    """
    def __init__(__self__, function_id=None, id=None, items=None, name_regex=None, output_file=None, total_count=None):
        if function_id and not isinstance(function_id, str):
            raise TypeError("Expected argument 'function_id' to be a str")
        pulumi.set(__self__, "function_id", function_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="functionId")
    def function_id(self) -> str:
        """
        The ID of Function.
        """
        return pulumi.get(self, "function_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.GetTimersItemResult']:
        """
        The list of timer trigger.
        """
        return pulumi.get(self, "items")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetTimersResult(GetTimersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTimersResult(
            function_id=self.function_id,
            id=self.id,
            items=self.items,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def get_timers(function_id: Optional[str] = None,
               name_regex: Optional[str] = None,
               output_file: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTimersResult:
    """
    Use this data source to query detailed information of vefaas timers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.vefaas.get_timers(function_id="g79asxxx")
    ```


    :param str function_id: The ID of Function.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['functionId'] = function_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vefaas/getTimers:getTimers', __args__, opts=opts, typ=GetTimersResult).value

    return AwaitableGetTimersResult(
        function_id=pulumi.get(__ret__, 'function_id'),
        id=pulumi.get(__ret__, 'id'),
        items=pulumi.get(__ret__, 'items'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_timers)
def get_timers_output(function_id: Optional[pulumi.Input[str]] = None,
                      name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTimersResult]:
    """
    Use this data source to query detailed information of vefaas timers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.vefaas.get_timers(function_id="g79asxxx")
    ```


    :param str function_id: The ID of Function.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    """
    ...
