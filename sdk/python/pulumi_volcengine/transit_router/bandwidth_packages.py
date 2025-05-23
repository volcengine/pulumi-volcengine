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
    'BandwidthPackagesResult',
    'AwaitableBandwidthPackagesResult',
    'bandwidth_packages',
    'bandwidth_packages_output',
]

warnings.warn("""volcengine.transit_router.BandwidthPackages has been deprecated in favor of volcengine.transit_router.getBandwidthPackages""", DeprecationWarning)

@pulumi.output_type
class BandwidthPackagesResult:
    """
    A collection of values returned by BandwidthPackages.
    """
    def __init__(__self__, bandwidth_packages=None, id=None, ids=None, local_geographic_region_set_id=None, output_file=None, peer_geographic_region_set_id=None, project_name=None, tags=None, total_count=None, transit_router_bandwidth_package_name=None, transit_router_peer_attachment_id=None):
        if bandwidth_packages and not isinstance(bandwidth_packages, list):
            raise TypeError("Expected argument 'bandwidth_packages' to be a list")
        pulumi.set(__self__, "bandwidth_packages", bandwidth_packages)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if local_geographic_region_set_id and not isinstance(local_geographic_region_set_id, str):
            raise TypeError("Expected argument 'local_geographic_region_set_id' to be a str")
        pulumi.set(__self__, "local_geographic_region_set_id", local_geographic_region_set_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if peer_geographic_region_set_id and not isinstance(peer_geographic_region_set_id, str):
            raise TypeError("Expected argument 'peer_geographic_region_set_id' to be a str")
        pulumi.set(__self__, "peer_geographic_region_set_id", peer_geographic_region_set_id)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if transit_router_bandwidth_package_name and not isinstance(transit_router_bandwidth_package_name, str):
            raise TypeError("Expected argument 'transit_router_bandwidth_package_name' to be a str")
        pulumi.set(__self__, "transit_router_bandwidth_package_name", transit_router_bandwidth_package_name)
        if transit_router_peer_attachment_id and not isinstance(transit_router_peer_attachment_id, str):
            raise TypeError("Expected argument 'transit_router_peer_attachment_id' to be a str")
        pulumi.set(__self__, "transit_router_peer_attachment_id", transit_router_peer_attachment_id)

    @property
    @pulumi.getter(name="bandwidthPackages")
    def bandwidth_packages(self) -> Sequence['outputs.BandwidthPackagesBandwidthPackageResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "bandwidth_packages")

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
    @pulumi.getter(name="localGeographicRegionSetId")
    def local_geographic_region_set_id(self) -> Optional[str]:
        """
        The local geographic region set ID.
        """
        return pulumi.get(self, "local_geographic_region_set_id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="peerGeographicRegionSetId")
    def peer_geographic_region_set_id(self) -> Optional[str]:
        """
        The peer geographic region set ID.
        """
        return pulumi.get(self, "peer_geographic_region_set_id")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The ProjectName of the transit router bandwidth package.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.BandwidthPackagesTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="transitRouterBandwidthPackageName")
    def transit_router_bandwidth_package_name(self) -> Optional[str]:
        """
        The name of the transit router bandwidth package.
        """
        return pulumi.get(self, "transit_router_bandwidth_package_name")

    @property
    @pulumi.getter(name="transitRouterPeerAttachmentId")
    def transit_router_peer_attachment_id(self) -> Optional[str]:
        """
        The ID of the peer attachment.
        """
        return pulumi.get(self, "transit_router_peer_attachment_id")


class AwaitableBandwidthPackagesResult(BandwidthPackagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return BandwidthPackagesResult(
            bandwidth_packages=self.bandwidth_packages,
            id=self.id,
            ids=self.ids,
            local_geographic_region_set_id=self.local_geographic_region_set_id,
            output_file=self.output_file,
            peer_geographic_region_set_id=self.peer_geographic_region_set_id,
            project_name=self.project_name,
            tags=self.tags,
            total_count=self.total_count,
            transit_router_bandwidth_package_name=self.transit_router_bandwidth_package_name,
            transit_router_peer_attachment_id=self.transit_router_peer_attachment_id)


def bandwidth_packages(ids: Optional[Sequence[str]] = None,
                       local_geographic_region_set_id: Optional[str] = None,
                       output_file: Optional[str] = None,
                       peer_geographic_region_set_id: Optional[str] = None,
                       project_name: Optional[str] = None,
                       tags: Optional[Sequence[pulumi.InputType['BandwidthPackagesTagArgs']]] = None,
                       transit_router_bandwidth_package_name: Optional[str] = None,
                       transit_router_peer_attachment_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableBandwidthPackagesResult:
    """
    Use this data source to query detailed information of transit router bandwidth packages
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_bandwidth_package = volcengine.transit_router.BandwidthPackage("fooBandwidthPackage",
        transit_router_bandwidth_package_name="acc-tf-test",
        description="acc-test",
        bandwidth=2,
        period=1,
        renew_type="Manual")
    foo_bandwidth_packages = volcengine.transit_router.get_bandwidth_packages_output(ids=[foo_bandwidth_package.id])
    ```


    :param Sequence[str] ids: The ID list of the TransitRouter bandwidth package.
    :param str local_geographic_region_set_id: The local geographic region set ID.
    :param str output_file: File name where to save data source results.
    :param str peer_geographic_region_set_id: The peer geographic region set ID.
    :param str project_name: The ProjectName of the TransitRouter bandwidth package.
    :param Sequence[pulumi.InputType['BandwidthPackagesTagArgs']] tags: Tags.
    :param str transit_router_bandwidth_package_name: The name of the TransitRouter bandwidth package.
    :param str transit_router_peer_attachment_id: The ID of the peer attachment.
    """
    pulumi.log.warn("""bandwidth_packages is deprecated: volcengine.transit_router.BandwidthPackages has been deprecated in favor of volcengine.transit_router.getBandwidthPackages""")
    __args__ = dict()
    __args__['ids'] = ids
    __args__['localGeographicRegionSetId'] = local_geographic_region_set_id
    __args__['outputFile'] = output_file
    __args__['peerGeographicRegionSetId'] = peer_geographic_region_set_id
    __args__['projectName'] = project_name
    __args__['tags'] = tags
    __args__['transitRouterBandwidthPackageName'] = transit_router_bandwidth_package_name
    __args__['transitRouterPeerAttachmentId'] = transit_router_peer_attachment_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:transit_router/bandwidthPackages:BandwidthPackages', __args__, opts=opts, typ=BandwidthPackagesResult).value

    return AwaitableBandwidthPackagesResult(
        bandwidth_packages=pulumi.get(__ret__, 'bandwidth_packages'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        local_geographic_region_set_id=pulumi.get(__ret__, 'local_geographic_region_set_id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        peer_geographic_region_set_id=pulumi.get(__ret__, 'peer_geographic_region_set_id'),
        project_name=pulumi.get(__ret__, 'project_name'),
        tags=pulumi.get(__ret__, 'tags'),
        total_count=pulumi.get(__ret__, 'total_count'),
        transit_router_bandwidth_package_name=pulumi.get(__ret__, 'transit_router_bandwidth_package_name'),
        transit_router_peer_attachment_id=pulumi.get(__ret__, 'transit_router_peer_attachment_id'))


@_utilities.lift_output_func(bandwidth_packages)
def bandwidth_packages_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                              local_geographic_region_set_id: Optional[pulumi.Input[Optional[str]]] = None,
                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                              peer_geographic_region_set_id: Optional[pulumi.Input[Optional[str]]] = None,
                              project_name: Optional[pulumi.Input[Optional[str]]] = None,
                              tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['BandwidthPackagesTagArgs']]]]] = None,
                              transit_router_bandwidth_package_name: Optional[pulumi.Input[Optional[str]]] = None,
                              transit_router_peer_attachment_id: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[BandwidthPackagesResult]:
    """
    Use this data source to query detailed information of transit router bandwidth packages
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_bandwidth_package = volcengine.transit_router.BandwidthPackage("fooBandwidthPackage",
        transit_router_bandwidth_package_name="acc-tf-test",
        description="acc-test",
        bandwidth=2,
        period=1,
        renew_type="Manual")
    foo_bandwidth_packages = volcengine.transit_router.get_bandwidth_packages_output(ids=[foo_bandwidth_package.id])
    ```


    :param Sequence[str] ids: The ID list of the TransitRouter bandwidth package.
    :param str local_geographic_region_set_id: The local geographic region set ID.
    :param str output_file: File name where to save data source results.
    :param str peer_geographic_region_set_id: The peer geographic region set ID.
    :param str project_name: The ProjectName of the TransitRouter bandwidth package.
    :param Sequence[pulumi.InputType['BandwidthPackagesTagArgs']] tags: Tags.
    :param str transit_router_bandwidth_package_name: The name of the TransitRouter bandwidth package.
    :param str transit_router_peer_attachment_id: The ID of the peer attachment.
    """
    pulumi.log.warn("""bandwidth_packages is deprecated: volcengine.transit_router.BandwidthPackages has been deprecated in favor of volcengine.transit_router.getBandwidthPackages""")
    ...
