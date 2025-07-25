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
    'GetVpcsResult',
    'AwaitableGetVpcsResult',
    'get_vpcs',
    'get_vpcs_output',
]

@pulumi.output_type
class GetVpcsResult:
    """
    A collection of values returned by getVpcs.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, output_file=None, project_name=None, tags=None, total_count=None, vpc_name=None, vpc_owner_id=None, vpcs=None):
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
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_name and not isinstance(vpc_name, str):
            raise TypeError("Expected argument 'vpc_name' to be a str")
        pulumi.set(__self__, "vpc_name", vpc_name)
        if vpc_owner_id and not isinstance(vpc_owner_id, int):
            raise TypeError("Expected argument 'vpc_owner_id' to be a int")
        pulumi.set(__self__, "vpc_owner_id", vpc_owner_id)
        if vpcs and not isinstance(vpcs, list):
            raise TypeError("Expected argument 'vpcs' to be a list")
        pulumi.set(__self__, "vpcs", vpcs)

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
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The ProjectName of the VPC.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.GetVpcsTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of Vpc query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> Optional[str]:
        """
        The name of VPC.
        """
        return pulumi.get(self, "vpc_name")

    @property
    @pulumi.getter(name="vpcOwnerId")
    def vpc_owner_id(self) -> Optional[int]:
        return pulumi.get(self, "vpc_owner_id")

    @property
    @pulumi.getter
    def vpcs(self) -> Sequence['outputs.GetVpcsVpcResult']:
        """
        The collection of Vpc query.
        """
        return pulumi.get(self, "vpcs")


class AwaitableGetVpcsResult(GetVpcsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcsResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            project_name=self.project_name,
            tags=self.tags,
            total_count=self.total_count,
            vpc_name=self.vpc_name,
            vpc_owner_id=self.vpc_owner_id,
            vpcs=self.vpcs)


def get_vpcs(ids: Optional[Sequence[str]] = None,
             name_regex: Optional[str] = None,
             output_file: Optional[str] = None,
             project_name: Optional[str] = None,
             tags: Optional[Sequence[pulumi.InputType['GetVpcsTagArgs']]] = None,
             vpc_name: Optional[str] = None,
             vpc_owner_id: Optional[int] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcsResult:
    """
    Use this data source to query detailed information of vpcs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpc.get_vpcs(ids=["vpc-mizl7m1kqccg5smt1bdpijuj"])
    ```


    :param Sequence[str] ids: A list of VPC IDs.
    :param str name_regex: A Name Regex of Vpc.
    :param str output_file: File name where to save data source results.
    :param str project_name: The ProjectName of the VPC.
    :param Sequence[pulumi.InputType['GetVpcsTagArgs']] tags: Tags.
    :param str vpc_name: The vpc name to query.
    :param int vpc_owner_id: The owner ID of the vpc.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['projectName'] = project_name
    __args__['tags'] = tags
    __args__['vpcName'] = vpc_name
    __args__['vpcOwnerId'] = vpc_owner_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vpc/getVpcs:getVpcs', __args__, opts=opts, typ=GetVpcsResult).value

    return AwaitableGetVpcsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        project_name=pulumi.get(__ret__, 'project_name'),
        tags=pulumi.get(__ret__, 'tags'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpc_name=pulumi.get(__ret__, 'vpc_name'),
        vpc_owner_id=pulumi.get(__ret__, 'vpc_owner_id'),
        vpcs=pulumi.get(__ret__, 'vpcs'))


@_utilities.lift_output_func(get_vpcs)
def get_vpcs_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    project_name: Optional[pulumi.Input[Optional[str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetVpcsTagArgs']]]]] = None,
                    vpc_name: Optional[pulumi.Input[Optional[str]]] = None,
                    vpc_owner_id: Optional[pulumi.Input[Optional[int]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcsResult]:
    """
    Use this data source to query detailed information of vpcs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpc.get_vpcs(ids=["vpc-mizl7m1kqccg5smt1bdpijuj"])
    ```


    :param Sequence[str] ids: A list of VPC IDs.
    :param str name_regex: A Name Regex of Vpc.
    :param str output_file: File name where to save data source results.
    :param str project_name: The ProjectName of the VPC.
    :param Sequence[pulumi.InputType['GetVpcsTagArgs']] tags: Tags.
    :param str vpc_name: The vpc name to query.
    :param int vpc_owner_id: The owner ID of the vpc.
    """
    ...
