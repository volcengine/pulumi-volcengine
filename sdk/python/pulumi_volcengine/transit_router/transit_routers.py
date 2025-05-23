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
    'TransitRoutersResult',
    'AwaitableTransitRoutersResult',
    'transit_routers',
    'transit_routers_output',
]

warnings.warn("""volcengine.transit_router.TransitRouters has been deprecated in favor of volcengine.transit_router.getTransitRouters""", DeprecationWarning)

@pulumi.output_type
class TransitRoutersResult:
    """
    A collection of values returned by TransitRouters.
    """
    def __init__(__self__, id=None, ids=None, output_file=None, project_name=None, tags=None, total_count=None, transit_router_name=None, transit_routers=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
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
        if transit_router_name and not isinstance(transit_router_name, str):
            raise TypeError("Expected argument 'transit_router_name' to be a str")
        pulumi.set(__self__, "transit_router_name", transit_router_name)
        if transit_routers and not isinstance(transit_routers, list):
            raise TypeError("Expected argument 'transit_routers' to be a list")
        pulumi.set(__self__, "transit_routers", transit_routers)

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
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The ProjectName of the transit router.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.TransitRoutersTagResult']]:
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
    @pulumi.getter(name="transitRouterName")
    def transit_router_name(self) -> Optional[str]:
        """
        The name of the transit router.
        """
        return pulumi.get(self, "transit_router_name")

    @property
    @pulumi.getter(name="transitRouters")
    def transit_routers(self) -> Sequence['outputs.TransitRoutersTransitRouterResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "transit_routers")


class AwaitableTransitRoutersResult(TransitRoutersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return TransitRoutersResult(
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            project_name=self.project_name,
            tags=self.tags,
            total_count=self.total_count,
            transit_router_name=self.transit_router_name,
            transit_routers=self.transit_routers)


def transit_routers(ids: Optional[Sequence[str]] = None,
                    output_file: Optional[str] = None,
                    project_name: Optional[str] = None,
                    tags: Optional[Sequence[pulumi.InputType['TransitRoutersTagArgs']]] = None,
                    transit_router_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableTransitRoutersResult:
    """
    Use this data source to query detailed information of transit routers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.transit_router.TransitRouter("foo",
        transit_router_name="test-tf-acc",
        description="test-tf-acc")
    default = volcengine.transit_router.get_transit_routers_output(ids=[foo.id],
        transit_router_name="test")
    ```


    :param Sequence[str] ids: A list of Transit Router ids.
    :param str output_file: File name where to save data source results.
    :param str project_name: The ProjectName of the transit router.
    :param Sequence[pulumi.InputType['TransitRoutersTagArgs']] tags: Tags.
    :param str transit_router_name: The name info.
    """
    pulumi.log.warn("""transit_routers is deprecated: volcengine.transit_router.TransitRouters has been deprecated in favor of volcengine.transit_router.getTransitRouters""")
    __args__ = dict()
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    __args__['projectName'] = project_name
    __args__['tags'] = tags
    __args__['transitRouterName'] = transit_router_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:transit_router/transitRouters:TransitRouters', __args__, opts=opts, typ=TransitRoutersResult).value

    return AwaitableTransitRoutersResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        output_file=pulumi.get(__ret__, 'output_file'),
        project_name=pulumi.get(__ret__, 'project_name'),
        tags=pulumi.get(__ret__, 'tags'),
        total_count=pulumi.get(__ret__, 'total_count'),
        transit_router_name=pulumi.get(__ret__, 'transit_router_name'),
        transit_routers=pulumi.get(__ret__, 'transit_routers'))


@_utilities.lift_output_func(transit_routers)
def transit_routers_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           project_name: Optional[pulumi.Input[Optional[str]]] = None,
                           tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['TransitRoutersTagArgs']]]]] = None,
                           transit_router_name: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[TransitRoutersResult]:
    """
    Use this data source to query detailed information of transit routers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.transit_router.TransitRouter("foo",
        transit_router_name="test-tf-acc",
        description="test-tf-acc")
    default = volcengine.transit_router.get_transit_routers_output(ids=[foo.id],
        transit_router_name="test")
    ```


    :param Sequence[str] ids: A list of Transit Router ids.
    :param str output_file: File name where to save data source results.
    :param str project_name: The ProjectName of the transit router.
    :param Sequence[pulumi.InputType['TransitRoutersTagArgs']] tags: Tags.
    :param str transit_router_name: The name info.
    """
    pulumi.log.warn("""transit_routers is deprecated: volcengine.transit_router.TransitRouters has been deprecated in favor of volcengine.transit_router.getTransitRouters""")
    ...
