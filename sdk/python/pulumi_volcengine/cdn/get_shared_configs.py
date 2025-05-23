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
    'GetSharedConfigsResult',
    'AwaitableGetSharedConfigsResult',
    'get_shared_configs',
    'get_shared_configs_output',
]

@pulumi.output_type
class GetSharedConfigsResult:
    """
    A collection of values returned by getSharedConfigs.
    """
    def __init__(__self__, config_datas=None, config_name=None, config_type=None, config_type_lists=None, id=None, output_file=None, project_name=None, total_count=None):
        if config_datas and not isinstance(config_datas, list):
            raise TypeError("Expected argument 'config_datas' to be a list")
        pulumi.set(__self__, "config_datas", config_datas)
        if config_name and not isinstance(config_name, str):
            raise TypeError("Expected argument 'config_name' to be a str")
        pulumi.set(__self__, "config_name", config_name)
        if config_type and not isinstance(config_type, str):
            raise TypeError("Expected argument 'config_type' to be a str")
        pulumi.set(__self__, "config_type", config_type)
        if config_type_lists and not isinstance(config_type_lists, list):
            raise TypeError("Expected argument 'config_type_lists' to be a list")
        pulumi.set(__self__, "config_type_lists", config_type_lists)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="configDatas")
    def config_datas(self) -> Sequence['outputs.GetSharedConfigsConfigDataResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "config_datas")

    @property
    @pulumi.getter(name="configName")
    def config_name(self) -> Optional[str]:
        """
        The name of the config.
        """
        return pulumi.get(self, "config_name")

    @property
    @pulumi.getter(name="configType")
    def config_type(self) -> Optional[str]:
        """
        The type of the config.
        """
        return pulumi.get(self, "config_type")

    @property
    @pulumi.getter(name="configTypeLists")
    def config_type_lists(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "config_type_lists")

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
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The name of the project.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetSharedConfigsResult(GetSharedConfigsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSharedConfigsResult(
            config_datas=self.config_datas,
            config_name=self.config_name,
            config_type=self.config_type,
            config_type_lists=self.config_type_lists,
            id=self.id,
            output_file=self.output_file,
            project_name=self.project_name,
            total_count=self.total_count)


def get_shared_configs(config_name: Optional[str] = None,
                       config_type: Optional[str] = None,
                       config_type_lists: Optional[Sequence[str]] = None,
                       output_file: Optional[str] = None,
                       project_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSharedConfigsResult:
    """
    Use this data source to query detailed information of cdn shared configs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cdn.get_shared_configs(config_name="tf-test",
        config_type="allow_ip_access_rule",
        project_name="default")
    ```


    :param str config_name: The name of the shared config.
    :param str config_type: The type of the shared config.
    :param Sequence[str] config_type_lists: The config type list. The parameter value can be a combination of available values for ConfigType. ConfigType and ConfigTypeList cannot be specified at the same time.
    :param str output_file: File name where to save data source results.
    :param str project_name: The name of the project.
    """
    __args__ = dict()
    __args__['configName'] = config_name
    __args__['configType'] = config_type
    __args__['configTypeLists'] = config_type_lists
    __args__['outputFile'] = output_file
    __args__['projectName'] = project_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:cdn/getSharedConfigs:getSharedConfigs', __args__, opts=opts, typ=GetSharedConfigsResult).value

    return AwaitableGetSharedConfigsResult(
        config_datas=pulumi.get(__ret__, 'config_datas'),
        config_name=pulumi.get(__ret__, 'config_name'),
        config_type=pulumi.get(__ret__, 'config_type'),
        config_type_lists=pulumi.get(__ret__, 'config_type_lists'),
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        project_name=pulumi.get(__ret__, 'project_name'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_shared_configs)
def get_shared_configs_output(config_name: Optional[pulumi.Input[Optional[str]]] = None,
                              config_type: Optional[pulumi.Input[Optional[str]]] = None,
                              config_type_lists: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              project_name: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSharedConfigsResult]:
    """
    Use this data source to query detailed information of cdn shared configs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cdn.get_shared_configs(config_name="tf-test",
        config_type="allow_ip_access_rule",
        project_name="default")
    ```


    :param str config_name: The name of the shared config.
    :param str config_type: The type of the shared config.
    :param Sequence[str] config_type_lists: The config type list. The parameter value can be a combination of available values for ConfigType. ConfigType and ConfigTypeList cannot be specified at the same time.
    :param str output_file: File name where to save data source results.
    :param str project_name: The name of the project.
    """
    ...
