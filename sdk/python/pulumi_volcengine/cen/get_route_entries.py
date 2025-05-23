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
    'GetRouteEntriesResult',
    'AwaitableGetRouteEntriesResult',
    'get_route_entries',
    'get_route_entries_output',
]

@pulumi.output_type
class GetRouteEntriesResult:
    """
    A collection of values returned by getRouteEntries.
    """
    def __init__(__self__, cen_id=None, cen_route_entries=None, destination_cidr_block=None, id=None, instance_id=None, instance_region_id=None, instance_type=None, output_file=None, total_count=None):
        if cen_id and not isinstance(cen_id, str):
            raise TypeError("Expected argument 'cen_id' to be a str")
        pulumi.set(__self__, "cen_id", cen_id)
        if cen_route_entries and not isinstance(cen_route_entries, list):
            raise TypeError("Expected argument 'cen_route_entries' to be a list")
        pulumi.set(__self__, "cen_route_entries", cen_route_entries)
        if destination_cidr_block and not isinstance(destination_cidr_block, str):
            raise TypeError("Expected argument 'destination_cidr_block' to be a str")
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if instance_region_id and not isinstance(instance_region_id, str):
            raise TypeError("Expected argument 'instance_region_id' to be a str")
        pulumi.set(__self__, "instance_region_id", instance_region_id)
        if instance_type and not isinstance(instance_type, str):
            raise TypeError("Expected argument 'instance_type' to be a str")
        pulumi.set(__self__, "instance_type", instance_type)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> str:
        """
        The cen ID of the cen route entry.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter(name="cenRouteEntries")
    def cen_route_entries(self) -> Sequence['outputs.GetRouteEntriesCenRouteEntryResult']:
        """
        The collection of cen route entry query.
        """
        return pulumi.get(self, "cen_route_entries")

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[str]:
        """
        The destination cidr block of the cen route entry.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[str]:
        """
        The instance id of the next hop of the cen route entry.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceRegionId")
    def instance_region_id(self) -> Optional[str]:
        """
        The instance region id of the next hop of the cen route entry.
        """
        return pulumi.get(self, "instance_region_id")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        """
        The instance type of the next hop of the cen route entry.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of cen route entry.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetRouteEntriesResult(GetRouteEntriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteEntriesResult(
            cen_id=self.cen_id,
            cen_route_entries=self.cen_route_entries,
            destination_cidr_block=self.destination_cidr_block,
            id=self.id,
            instance_id=self.instance_id,
            instance_region_id=self.instance_region_id,
            instance_type=self.instance_type,
            output_file=self.output_file,
            total_count=self.total_count)


def get_route_entries(cen_id: Optional[str] = None,
                      destination_cidr_block: Optional[str] = None,
                      instance_id: Optional[str] = None,
                      instance_region_id: Optional[str] = None,
                      instance_type: Optional[str] = None,
                      output_file: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteEntriesResult:
    """
    Use this data source to query detailed information of cen route entries


    :param str cen_id: A cen ID.
    :param str destination_cidr_block: A destination cidr block.
    :param str instance_id: An instance ID.
    :param str instance_region_id: An instance region ID.
    :param str instance_type: An instance type.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['cenId'] = cen_id
    __args__['destinationCidrBlock'] = destination_cidr_block
    __args__['instanceId'] = instance_id
    __args__['instanceRegionId'] = instance_region_id
    __args__['instanceType'] = instance_type
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:cen/getRouteEntries:getRouteEntries', __args__, opts=opts, typ=GetRouteEntriesResult).value

    return AwaitableGetRouteEntriesResult(
        cen_id=pulumi.get(__ret__, 'cen_id'),
        cen_route_entries=pulumi.get(__ret__, 'cen_route_entries'),
        destination_cidr_block=pulumi.get(__ret__, 'destination_cidr_block'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        instance_region_id=pulumi.get(__ret__, 'instance_region_id'),
        instance_type=pulumi.get(__ret__, 'instance_type'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_route_entries)
def get_route_entries_output(cen_id: Optional[pulumi.Input[str]] = None,
                             destination_cidr_block: Optional[pulumi.Input[Optional[str]]] = None,
                             instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                             instance_region_id: Optional[pulumi.Input[Optional[str]]] = None,
                             instance_type: Optional[pulumi.Input[Optional[str]]] = None,
                             output_file: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouteEntriesResult]:
    """
    Use this data source to query detailed information of cen route entries


    :param str cen_id: A cen ID.
    :param str destination_cidr_block: A destination cidr block.
    :param str instance_id: An instance ID.
    :param str instance_region_id: An instance region ID.
    :param str instance_type: An instance type.
    :param str output_file: File name where to save data source results.
    """
    ...
