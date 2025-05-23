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
    'GetRouteTablesResult',
    'AwaitableGetRouteTablesResult',
    'get_route_tables',
    'get_route_tables_output',
]

@pulumi.output_type
class GetRouteTablesResult:
    """
    A collection of values returned by getRouteTables.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, route_tables=None, tags=None, total_count=None, transit_router_id=None, transit_router_route_table_type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if route_tables and not isinstance(route_tables, list):
            raise TypeError("Expected argument 'route_tables' to be a list")
        pulumi.set(__self__, "route_tables", route_tables)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if transit_router_id and not isinstance(transit_router_id, str):
            raise TypeError("Expected argument 'transit_router_id' to be a str")
        pulumi.set(__self__, "transit_router_id", transit_router_id)
        if transit_router_route_table_type and not isinstance(transit_router_route_table_type, str):
            raise TypeError("Expected argument 'transit_router_route_table_type' to be a str")
        pulumi.set(__self__, "transit_router_route_table_type", transit_router_route_table_type)

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="routeTables")
    def route_tables(self) -> Sequence['outputs.GetRouteTablesRouteTableResult']:
        """
        The list of route tables query.
        """
        return pulumi.get(self, "route_tables")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.GetRouteTablesTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of data query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> str:
        return pulumi.get(self, "transit_router_id")

    @property
    @pulumi.getter(name="transitRouterRouteTableType")
    def transit_router_route_table_type(self) -> Optional[str]:
        """
        The type of route table.
        """
        return pulumi.get(self, "transit_router_route_table_type")


class AwaitableGetRouteTablesResult(GetRouteTablesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteTablesResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            route_tables=self.route_tables,
            tags=self.tags,
            total_count=self.total_count,
            transit_router_id=self.transit_router_id,
            transit_router_route_table_type=self.transit_router_route_table_type)


def get_route_tables(ids: Optional[Sequence[str]] = None,
                     output_file: Optional[str] = None,
                     tags: Optional[Sequence[pulumi.InputType['GetRouteTablesTagArgs']]] = None,
                     transit_router_id: Optional[str] = None,
                     transit_router_route_table_type: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteTablesResult:
    """
    Use this data source to query detailed information of transit router route tables
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_transit_router = volcengine.transit_router.TransitRouter("fooTransitRouter",
        transit_router_name="test-tf-acc",
        description="test-tf-acc")
    foo_route_table = volcengine.transit_router.RouteTable("fooRouteTable",
        description="tf-test-acc-description",
        transit_router_route_table_name="tf-table-test-acc",
        transit_router_id=foo_transit_router.id)
    default = volcengine.transit_router.get_route_tables_output(transit_router_id=foo_transit_router.id,
        ids=[foo_route_table.transit_router_route_table_id])
    ```


    :param Sequence[str] ids: The ids of the transit router route table.
    :param str output_file: File name where to save data source results.
    :param Sequence[pulumi.InputType['GetRouteTablesTagArgs']] tags: Tags.
    :param str transit_router_id: The id of the transit router.
    :param str transit_router_route_table_type: The type of the route table. The value can be System or Custom.
    """
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['tags'] = tags
    __args__['transitRouterId'] = transit_router_id
    __args__['transitRouterRouteTableType'] = transit_router_route_table_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:transit_router/getRouteTables:getRouteTables', __args__, opts=opts, typ=GetRouteTablesResult).value

    return AwaitableGetRouteTablesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        output_file=pulumi.get(__ret__, 'output_file'),
        route_tables=pulumi.get(__ret__, 'route_tables'),
        tags=pulumi.get(__ret__, 'tags'),
        total_count=pulumi.get(__ret__, 'total_count'),
        transit_router_id=pulumi.get(__ret__, 'transit_router_id'),
        transit_router_route_table_type=pulumi.get(__ret__, 'transit_router_route_table_type'))


@_utilities.lift_output_func(get_route_tables)
def get_route_tables_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRouteTablesTagArgs']]]]] = None,
                            transit_router_id: Optional[pulumi.Input[str]] = None,
                            transit_router_route_table_type: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouteTablesResult]:
    """
    Use this data source to query detailed information of transit router route tables
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_transit_router = volcengine.transit_router.TransitRouter("fooTransitRouter",
        transit_router_name="test-tf-acc",
        description="test-tf-acc")
    foo_route_table = volcengine.transit_router.RouteTable("fooRouteTable",
        description="tf-test-acc-description",
        transit_router_route_table_name="tf-table-test-acc",
        transit_router_id=foo_transit_router.id)
    default = volcengine.transit_router.get_route_tables_output(transit_router_id=foo_transit_router.id,
        ids=[foo_route_table.transit_router_route_table_id])
    ```


    :param Sequence[str] ids: The ids of the transit router route table.
    :param str output_file: File name where to save data source results.
    :param Sequence[pulumi.InputType['GetRouteTablesTagArgs']] tags: Tags.
    :param str transit_router_id: The id of the transit router.
    :param str transit_router_route_table_type: The type of the route table. The value can be System or Custom.
    """
    ...
