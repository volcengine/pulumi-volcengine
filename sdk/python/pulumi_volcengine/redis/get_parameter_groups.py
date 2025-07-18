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
    'GetParameterGroupsResult',
    'AwaitableGetParameterGroupsResult',
    'get_parameter_groups',
    'get_parameter_groups_output',
]

@pulumi.output_type
class GetParameterGroupsResult:
    """
    A collection of values returned by getParameterGroups.
    """
    def __init__(__self__, engine_version=None, id=None, name_regex=None, output_file=None, parameter_groups=None, source=None, total_count=None):
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        pulumi.set(__self__, "engine_version", engine_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if parameter_groups and not isinstance(parameter_groups, list):
            raise TypeError("Expected argument 'parameter_groups' to be a list")
        pulumi.set(__self__, "parameter_groups", parameter_groups)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[str]:
        """
        The database version applicable to the parameter template.
        """
        return pulumi.get(self, "engine_version")

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
    @pulumi.getter(name="parameterGroups")
    def parameter_groups(self) -> Sequence['outputs.GetParameterGroupsParameterGroupResult']:
        """
        The details of the parameter template.
        """
        return pulumi.get(self, "parameter_groups")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        """
        The source of creating the parameter template.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetParameterGroupsResult(GetParameterGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetParameterGroupsResult(
            engine_version=self.engine_version,
            id=self.id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            parameter_groups=self.parameter_groups,
            source=self.source,
            total_count=self.total_count)


def get_parameter_groups(engine_version: Optional[str] = None,
                         name_regex: Optional[str] = None,
                         output_file: Optional[str] = None,
                         source: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetParameterGroupsResult:
    """
    Use this data source to query detailed information of redis parameter groups


    :param str engine_version: The Redis database version applicable to the parameter template.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str source: The source of creating the parameter template.
    """
    __args__ = dict()
    __args__['engineVersion'] = engine_version
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['source'] = source
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:redis/getParameterGroups:getParameterGroups', __args__, opts=opts, typ=GetParameterGroupsResult).value

    return AwaitableGetParameterGroupsResult(
        engine_version=pulumi.get(__ret__, 'engine_version'),
        id=pulumi.get(__ret__, 'id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        parameter_groups=pulumi.get(__ret__, 'parameter_groups'),
        source=pulumi.get(__ret__, 'source'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_parameter_groups)
def get_parameter_groups_output(engine_version: Optional[pulumi.Input[Optional[str]]] = None,
                                name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                source: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetParameterGroupsResult]:
    """
    Use this data source to query detailed information of redis parameter groups


    :param str engine_version: The Redis database version applicable to the parameter template.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str source: The source of creating the parameter template.
    """
    ...
