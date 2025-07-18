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
    'RegionsResult',
    'AwaitableRegionsResult',
    'regions',
    'regions_output',
]

warnings.warn("""volcengine.rabbitmq.Regions has been deprecated in favor of volcengine.rabbitmq.getRegions""", DeprecationWarning)

@pulumi.output_type
class RegionsResult:
    """
    A collection of values returned by Regions.
    """
    def __init__(__self__, id=None, output_file=None, regions=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)
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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def regions(self) -> Sequence['outputs.RegionsRegionResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableRegionsResult(RegionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return RegionsResult(
            id=self.id,
            output_file=self.output_file,
            regions=self.regions,
            total_count=self.total_count)


def regions(output_file: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableRegionsResult:
    """
    Use this data source to query detailed information of rabbitmq regions
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.rabbitmq.get_regions()
    ```


    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""regions is deprecated: volcengine.rabbitmq.Regions has been deprecated in favor of volcengine.rabbitmq.getRegions""")
    __args__ = dict()
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:rabbitmq/regions:Regions', __args__, opts=opts, typ=RegionsResult).value

    return AwaitableRegionsResult(
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        regions=pulumi.get(__ret__, 'regions'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(regions)
def regions_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[RegionsResult]:
    """
    Use this data source to query detailed information of rabbitmq regions
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.rabbitmq.get_regions()
    ```


    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""regions is deprecated: volcengine.rabbitmq.Regions has been deprecated in favor of volcengine.rabbitmq.getRegions""")
    ...
