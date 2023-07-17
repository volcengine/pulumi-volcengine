# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'ZonesResult',
    'AwaitableZonesResult',
    'zones',
    'zones_output',
]

@pulumi.output_type
class ZonesResult:
    """
    A collection of values returned by Zones.
    """
    def __init__(__self__, id=None, master_zones=None, output_file=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if master_zones and not isinstance(master_zones, list):
            raise TypeError("Expected argument 'master_zones' to be a list")
        pulumi.set(__self__, "master_zones", master_zones)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="masterZones")
    def master_zones(self) -> Sequence['outputs.ZonesMasterZoneResult']:
        """
        The master zones list.
        """
        return pulumi.get(self, "master_zones")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of zone query.
        """
        return pulumi.get(self, "total_count")


class AwaitableZonesResult(ZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ZonesResult(
            id=self.id,
            master_zones=self.master_zones,
            output_file=self.output_file,
            total_count=self.total_count)


def zones(output_file: Optional[str] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableZonesResult:
    """
    Use this data source to query detailed information of clb zones
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.clb.zones()
    ```


    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:clb/zones:Zones', __args__, opts=opts, typ=ZonesResult).value

    return AwaitableZonesResult(
        id=__ret__.id,
        master_zones=__ret__.master_zones,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(zones)
def zones_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ZonesResult]:
    """
    Use this data source to query detailed information of clb zones
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.clb.zones()
    ```


    :param str output_file: File name where to save data source results.
    """
    ...