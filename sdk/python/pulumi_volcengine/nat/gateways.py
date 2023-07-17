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
    'GatewaysResult',
    'AwaitableGatewaysResult',
    'gateways',
    'gateways_output',
]

@pulumi.output_type
class GatewaysResult:
    """
    A collection of values returned by Gateways.
    """
    def __init__(__self__, description=None, id=None, ids=None, name_regex=None, nat_gateway_name=None, nat_gateways=None, output_file=None, spec=None, subnet_id=None, tags=None, total_count=None, vpc_id=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if nat_gateway_name and not isinstance(nat_gateway_name, str):
            raise TypeError("Expected argument 'nat_gateway_name' to be a str")
        pulumi.set(__self__, "nat_gateway_name", nat_gateway_name)
        if nat_gateways and not isinstance(nat_gateways, list):
            raise TypeError("Expected argument 'nat_gateways' to be a list")
        pulumi.set(__self__, "nat_gateways", nat_gateways)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if spec and not isinstance(spec, str):
            raise TypeError("Expected argument 'spec' to be a str")
        pulumi.set(__self__, "spec", spec)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
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
    def description(self) -> Optional[str]:
        """
        The description of the NatGateway.
        """
        return pulumi.get(self, "description")

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
    @pulumi.getter(name="natGatewayName")
    def nat_gateway_name(self) -> Optional[str]:
        """
        The name of the NatGateway.
        """
        return pulumi.get(self, "nat_gateway_name")

    @property
    @pulumi.getter(name="natGateways")
    def nat_gateways(self) -> Sequence['outputs.GatewaysNatGatewayResult']:
        """
        The collection of NatGateway query.
        """
        return pulumi.get(self, "nat_gateways")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def spec(self) -> Optional[str]:
        """
        The specification of the NatGateway.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The ID of the Subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.GatewaysTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of NatGateway query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGatewaysResult(GatewaysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GatewaysResult(
            description=self.description,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            nat_gateway_name=self.nat_gateway_name,
            nat_gateways=self.nat_gateways,
            output_file=self.output_file,
            spec=self.spec,
            subnet_id=self.subnet_id,
            tags=self.tags,
            total_count=self.total_count,
            vpc_id=self.vpc_id)


def gateways(description: Optional[str] = None,
             ids: Optional[Sequence[str]] = None,
             name_regex: Optional[str] = None,
             nat_gateway_name: Optional[str] = None,
             output_file: Optional[str] = None,
             spec: Optional[str] = None,
             subnet_id: Optional[str] = None,
             tags: Optional[Sequence[pulumi.InputType['GatewaysTagArgs']]] = None,
             vpc_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGatewaysResult:
    """
    Use this data source to query detailed information of nat gateways
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.nat.gateways(ids=[
        "ngw-2743w1f6iqby87fap8tvm9kop",
        "ngw-274gwbqe340zk7fap8spkzo7x",
    ])
    ```


    :param str description: The description of the NatGateway.
    :param Sequence[str] ids: The list of NatGateway IDs.
    :param str name_regex: The Name Regex of NatGateway.
    :param str nat_gateway_name: The name of the NatGateway.
    :param str output_file: File name where to save data source results.
    :param str spec: The specification of the NatGateway.
    :param str subnet_id: The id of the Subnet.
    :param Sequence[pulumi.InputType['GatewaysTagArgs']] tags: Tags.
    :param str vpc_id: The id of the VPC.
    """
    __args__ = dict()
    __args__['description'] = description
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['natGatewayName'] = nat_gateway_name
    __args__['outputFile'] = output_file
    __args__['spec'] = spec
    __args__['subnetId'] = subnet_id
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:nat/gateways:Gateways', __args__, opts=opts, typ=GatewaysResult).value

    return AwaitableGatewaysResult(
        description=__ret__.description,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        nat_gateway_name=__ret__.nat_gateway_name,
        nat_gateways=__ret__.nat_gateways,
        output_file=__ret__.output_file,
        spec=__ret__.spec,
        subnet_id=__ret__.subnet_id,
        tags=__ret__.tags,
        total_count=__ret__.total_count,
        vpc_id=__ret__.vpc_id)


@_utilities.lift_output_func(gateways)
def gateways_output(description: Optional[pulumi.Input[Optional[str]]] = None,
                    ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                    nat_gateway_name: Optional[pulumi.Input[Optional[str]]] = None,
                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    spec: Optional[pulumi.Input[Optional[str]]] = None,
                    subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GatewaysTagArgs']]]]] = None,
                    vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GatewaysResult]:
    """
    Use this data source to query detailed information of nat gateways
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.nat.gateways(ids=[
        "ngw-2743w1f6iqby87fap8tvm9kop",
        "ngw-274gwbqe340zk7fap8spkzo7x",
    ])
    ```


    :param str description: The description of the NatGateway.
    :param Sequence[str] ids: The list of NatGateway IDs.
    :param str name_regex: The Name Regex of NatGateway.
    :param str nat_gateway_name: The name of the NatGateway.
    :param str output_file: File name where to save data source results.
    :param str spec: The specification of the NatGateway.
    :param str subnet_id: The id of the Subnet.
    :param Sequence[pulumi.InputType['GatewaysTagArgs']] tags: Tags.
    :param str vpc_id: The id of the VPC.
    """
    ...