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
    'GetResolverEndpointsResult',
    'AwaitableGetResolverEndpointsResult',
    'get_resolver_endpoints',
    'get_resolver_endpoints_output',
]

@pulumi.output_type
class GetResolverEndpointsResult:
    """
    A collection of values returned by getResolverEndpoints.
    """
    def __init__(__self__, direction=None, endpoints=None, id=None, name=None, name_regex=None, output_file=None, project_name=None, status=None, tag_filters=None, total_count=None, vpc_id=None):
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if endpoints and not isinstance(endpoints, list):
            raise TypeError("Expected argument 'endpoints' to be a list")
        pulumi.set(__self__, "endpoints", endpoints)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tag_filters and not isinstance(tag_filters, list):
            raise TypeError("Expected argument 'tag_filters' to be a list")
        pulumi.set(__self__, "tag_filters", tag_filters)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def direction(self) -> Optional[str]:
        """
        The direction of the endpoint.
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def endpoints(self) -> Sequence['outputs.GetResolverEndpointsEndpointResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the endpoint.
        """
        return pulumi.get(self, "name")

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
        The project name of the endpoint.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of the endpoint.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tagFilters")
    def tag_filters(self) -> Optional[Sequence['outputs.GetResolverEndpointsTagFilterResult']]:
        return pulumi.get(self, "tag_filters")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The vpc id of the endpoint.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableGetResolverEndpointsResult(GetResolverEndpointsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverEndpointsResult(
            direction=self.direction,
            endpoints=self.endpoints,
            id=self.id,
            name=self.name,
            name_regex=self.name_regex,
            output_file=self.output_file,
            project_name=self.project_name,
            status=self.status,
            tag_filters=self.tag_filters,
            total_count=self.total_count,
            vpc_id=self.vpc_id)


def get_resolver_endpoints(direction: Optional[str] = None,
                           name: Optional[str] = None,
                           name_regex: Optional[str] = None,
                           output_file: Optional[str] = None,
                           project_name: Optional[str] = None,
                           status: Optional[str] = None,
                           tag_filters: Optional[Sequence[pulumi.InputType['GetResolverEndpointsTagFilterArgs']]] = None,
                           vpc_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResolverEndpointsResult:
    """
    Use this data source to query detailed information of private zone resolver endpoints
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.private_zone.get_resolver_endpoints()
    ```


    :param str direction: The direction of the private zone resolver endpoint.
    :param str name: The name of the private zone resolver endpoint.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str project_name: The project name of the private zone resolver endpoint.
    :param str status: The status of the private zone resolver endpoint.
    :param Sequence[pulumi.InputType['GetResolverEndpointsTagFilterArgs']] tag_filters: List of tag filters.
    :param str vpc_id: The vpc ID of the private zone resolver endpoint.
    """
    __args__ = dict()
    __args__['direction'] = direction
    __args__['name'] = name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['projectName'] = project_name
    __args__['status'] = status
    __args__['tagFilters'] = tag_filters
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:private_zone/getResolverEndpoints:getResolverEndpoints', __args__, opts=opts, typ=GetResolverEndpointsResult).value

    return AwaitableGetResolverEndpointsResult(
        direction=pulumi.get(__ret__, 'direction'),
        endpoints=pulumi.get(__ret__, 'endpoints'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        project_name=pulumi.get(__ret__, 'project_name'),
        status=pulumi.get(__ret__, 'status'),
        tag_filters=pulumi.get(__ret__, 'tag_filters'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(get_resolver_endpoints)
def get_resolver_endpoints_output(direction: Optional[pulumi.Input[Optional[str]]] = None,
                                  name: Optional[pulumi.Input[Optional[str]]] = None,
                                  name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  project_name: Optional[pulumi.Input[Optional[str]]] = None,
                                  status: Optional[pulumi.Input[Optional[str]]] = None,
                                  tag_filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetResolverEndpointsTagFilterArgs']]]]] = None,
                                  vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResolverEndpointsResult]:
    """
    Use this data source to query detailed information of private zone resolver endpoints
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.private_zone.get_resolver_endpoints()
    ```


    :param str direction: The direction of the private zone resolver endpoint.
    :param str name: The name of the private zone resolver endpoint.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str project_name: The project name of the private zone resolver endpoint.
    :param str status: The status of the private zone resolver endpoint.
    :param Sequence[pulumi.InputType['GetResolverEndpointsTagFilterArgs']] tag_filters: List of tag filters.
    :param str vpc_id: The vpc ID of the private zone resolver endpoint.
    """
    ...
