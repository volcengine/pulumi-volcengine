# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'ClbsResult',
    'AwaitableClbsResult',
    'clbs',
    'clbs_output',
]

@pulumi.output_type
class ClbsResult:
    """
    A collection of values returned by Clbs.
    """
    def __init__(__self__, clbs=None, eni_address=None, id=None, ids=None, load_balancer_name=None, name_regex=None, output_file=None, project_name=None, tags=None, total_count=None, vpc_id=None):
        if clbs and not isinstance(clbs, list):
            raise TypeError("Expected argument 'clbs' to be a list")
        pulumi.set(__self__, "clbs", clbs)
        if eni_address and not isinstance(eni_address, str):
            raise TypeError("Expected argument 'eni_address' to be a str")
        pulumi.set(__self__, "eni_address", eni_address)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if load_balancer_name and not isinstance(load_balancer_name, str):
            raise TypeError("Expected argument 'load_balancer_name' to be a str")
        pulumi.set(__self__, "load_balancer_name", load_balancer_name)
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
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def clbs(self) -> Sequence['outputs.ClbsClbResult']:
        """
        The collection of Clb query.
        """
        return pulumi.get(self, "clbs")

    @property
    @pulumi.getter(name="eniAddress")
    def eni_address(self) -> Optional[str]:
        """
        The Eni address of the Clb.
        """
        return pulumi.get(self, "eni_address")

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
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> Optional[str]:
        """
        The name of the Clb.
        """
        return pulumi.get(self, "load_balancer_name")

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
        The ProjectName of the Clb.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.ClbsTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of Clb query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The vpc ID of the Clb.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableClbsResult(ClbsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ClbsResult(
            clbs=self.clbs,
            eni_address=self.eni_address,
            id=self.id,
            ids=self.ids,
            load_balancer_name=self.load_balancer_name,
            name_regex=self.name_regex,
            output_file=self.output_file,
            project_name=self.project_name,
            tags=self.tags,
            total_count=self.total_count,
            vpc_id=self.vpc_id)


def clbs(eni_address: Optional[str] = None,
         ids: Optional[Sequence[str]] = None,
         load_balancer_name: Optional[str] = None,
         name_regex: Optional[str] = None,
         output_file: Optional[str] = None,
         project_name: Optional[str] = None,
         tags: Optional[Sequence[pulumi.InputType['ClbsTagArgs']]] = None,
         vpc_id: Optional[str] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableClbsResult:
    """
    Use this data source to query detailed information of clbs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.clb.clbs(ids=["clb-273y2ok6ets007fap8txvf6us"])
    ```


    :param str eni_address: The private ip address of the Clb.
    :param Sequence[str] ids: A list of Clb IDs.
    :param str load_balancer_name: The name of the Clb.
    :param str name_regex: A Name Regex of Clb.
    :param str output_file: File name where to save data source results.
    :param str project_name: The ProjectName of Clb.
    :param Sequence[pulumi.InputType['ClbsTagArgs']] tags: Tags.
    :param str vpc_id: The id of the VPC.
    """
    __args__ = dict()
    __args__['eniAddress'] = eni_address
    __args__['ids'] = ids
    __args__['loadBalancerName'] = load_balancer_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['projectName'] = project_name
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:clb/clbs:Clbs', __args__, opts=opts, typ=ClbsResult).value

    return AwaitableClbsResult(
        clbs=__ret__.clbs,
        eni_address=__ret__.eni_address,
        id=__ret__.id,
        ids=__ret__.ids,
        load_balancer_name=__ret__.load_balancer_name,
        name_regex=__ret__.name_regex,
        output_file=__ret__.output_file,
        project_name=__ret__.project_name,
        tags=__ret__.tags,
        total_count=__ret__.total_count,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(clbs)
def clbs_output(eni_address: Optional[pulumi.Input[Optional[str]]] = None,
                ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                load_balancer_name: Optional[pulumi.Input[Optional[str]]] = None,
                name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                project_name: Optional[pulumi.Input[Optional[str]]] = None,
                tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['ClbsTagArgs']]]]] = None,
                vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ClbsResult]:
    """
    Use this data source to query detailed information of clbs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.clb.clbs(ids=["clb-273y2ok6ets007fap8txvf6us"])
    ```


    :param str eni_address: The private ip address of the Clb.
    :param Sequence[str] ids: A list of Clb IDs.
    :param str load_balancer_name: The name of the Clb.
    :param str name_regex: A Name Regex of Clb.
    :param str output_file: File name where to save data source results.
    :param str project_name: The ProjectName of Clb.
    :param Sequence[pulumi.InputType['ClbsTagArgs']] tags: Tags.
    :param str vpc_id: The id of the VPC.
    """
    ...