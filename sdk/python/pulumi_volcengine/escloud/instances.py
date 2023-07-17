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
    'InstancesResult',
    'AwaitableInstancesResult',
    'instances',
    'instances_output',
]

@pulumi.output_type
class InstancesResult:
    """
    A collection of values returned by Instances.
    """
    def __init__(__self__, charge_types=None, id=None, ids=None, instances=None, names=None, output_file=None, statuses=None, total_count=None, versions=None, zone_ids=None):
        if charge_types and not isinstance(charge_types, list):
            raise TypeError("Expected argument 'charge_types' to be a list")
        pulumi.set(__self__, "charge_types", charge_types)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)
        if zone_ids and not isinstance(zone_ids, list):
            raise TypeError("Expected argument 'zone_ids' to be a list")
        pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="chargeTypes")
    def charge_types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "charge_types")

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
    @pulumi.getter
    def instances(self) -> Sequence['outputs.InstancesInstanceResult']:
        """
        The collection of instance query.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of instance query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def versions(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "versions")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "zone_ids")


class AwaitableInstancesResult(InstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstancesResult(
            charge_types=self.charge_types,
            id=self.id,
            ids=self.ids,
            instances=self.instances,
            names=self.names,
            output_file=self.output_file,
            statuses=self.statuses,
            total_count=self.total_count,
            versions=self.versions,
            zone_ids=self.zone_ids)


def instances(charge_types: Optional[Sequence[str]] = None,
              ids: Optional[Sequence[str]] = None,
              names: Optional[Sequence[str]] = None,
              output_file: Optional[str] = None,
              statuses: Optional[Sequence[str]] = None,
              versions: Optional[Sequence[str]] = None,
              zone_ids: Optional[Sequence[str]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstancesResult:
    """
    Use this data source to query detailed information of escloud instances
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.escloud.instances(ids=["d3gftqjvnah74eie"],
        statuses=["Running"])
    ```


    :param Sequence[str] charge_types: The charge types of instance.
    :param Sequence[str] ids: A list of instance IDs.
    :param Sequence[str] names: The names of instance.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] statuses: The list status of instance.
    :param Sequence[str] versions: The versions of instance.
    :param Sequence[str] zone_ids: The available zone IDs of instance.
    """
    __args__ = dict()
    __args__['chargeTypes'] = charge_types
    __args__['ids'] = ids
    __args__['names'] = names
    __args__['outputFile'] = output_file
    __args__['statuses'] = statuses
    __args__['versions'] = versions
    __args__['zoneIds'] = zone_ids
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:escloud/instances:Instances', __args__, opts=opts, typ=InstancesResult).value

    return AwaitableInstancesResult(
        charge_types=__ret__.charge_types,
        id=__ret__.id,
        ids=__ret__.ids,
        instances=__ret__.instances,
        names=__ret__.names,
        output_file=__ret__.output_file,
        statuses=__ret__.statuses,
        total_count=__ret__.total_count,
        versions=__ret__.versions,
        zone_ids=__ret__.zone_ids)


@_utilities.lift_output_func(instances)
def instances_output(charge_types: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     versions: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     zone_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstancesResult]:
    """
    Use this data source to query detailed information of escloud instances
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.escloud.instances(ids=["d3gftqjvnah74eie"],
        statuses=["Running"])
    ```


    :param Sequence[str] charge_types: The charge types of instance.
    :param Sequence[str] ids: A list of instance IDs.
    :param Sequence[str] names: The names of instance.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] statuses: The list status of instance.
    :param Sequence[str] versions: The versions of instance.
    :param Sequence[str] zone_ids: The available zone IDs of instance.
    """
    ...