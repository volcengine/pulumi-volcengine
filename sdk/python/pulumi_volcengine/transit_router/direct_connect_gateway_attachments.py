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
    'DirectConnectGatewayAttachmentsResult',
    'AwaitableDirectConnectGatewayAttachmentsResult',
    'direct_connect_gateway_attachments',
    'direct_connect_gateway_attachments_output',
]

@pulumi.output_type
class DirectConnectGatewayAttachmentsResult:
    """
    A collection of values returned by DirectConnectGatewayAttachments.
    """
    def __init__(__self__, attachments=None, direct_connect_gateway_id=None, id=None, output_file=None, total_count=None, transit_router_attachment_ids=None, transit_router_id=None):
        if attachments and not isinstance(attachments, list):
            raise TypeError("Expected argument 'attachments' to be a list")
        pulumi.set(__self__, "attachments", attachments)
        if direct_connect_gateway_id and not isinstance(direct_connect_gateway_id, str):
            raise TypeError("Expected argument 'direct_connect_gateway_id' to be a str")
        pulumi.set(__self__, "direct_connect_gateway_id", direct_connect_gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if transit_router_attachment_ids and not isinstance(transit_router_attachment_ids, list):
            raise TypeError("Expected argument 'transit_router_attachment_ids' to be a list")
        pulumi.set(__self__, "transit_router_attachment_ids", transit_router_attachment_ids)
        if transit_router_id and not isinstance(transit_router_id, str):
            raise TypeError("Expected argument 'transit_router_id' to be a str")
        pulumi.set(__self__, "transit_router_id", transit_router_id)

    @property
    @pulumi.getter
    def attachments(self) -> Sequence['outputs.DirectConnectGatewayAttachmentsAttachmentResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "attachments")

    @property
    @pulumi.getter(name="directConnectGatewayId")
    def direct_connect_gateway_id(self) -> Optional[str]:
        """
        The direct connect gateway id.
        """
        return pulumi.get(self, "direct_connect_gateway_id")

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
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="transitRouterAttachmentIds")
    def transit_router_attachment_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "transit_router_attachment_ids")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> str:
        """
        The id of the transit router.
        """
        return pulumi.get(self, "transit_router_id")


class AwaitableDirectConnectGatewayAttachmentsResult(DirectConnectGatewayAttachmentsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return DirectConnectGatewayAttachmentsResult(
            attachments=self.attachments,
            direct_connect_gateway_id=self.direct_connect_gateway_id,
            id=self.id,
            output_file=self.output_file,
            total_count=self.total_count,
            transit_router_attachment_ids=self.transit_router_attachment_ids,
            transit_router_id=self.transit_router_id)


def direct_connect_gateway_attachments(direct_connect_gateway_id: Optional[str] = None,
                                       output_file: Optional[str] = None,
                                       transit_router_attachment_ids: Optional[Sequence[str]] = None,
                                       transit_router_id: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableDirectConnectGatewayAttachmentsResult:
    """
    Use this data source to query detailed information of transit router direct connect gateway attachments
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.transit_router.direct_connect_gateway_attachments(transit_router_id="tr-2bzy39x27qtxc2dx0eg5qaj05")
    ```


    :param str direct_connect_gateway_id: ID of the direct connection gateway.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] transit_router_attachment_ids: ID of the network instance connection.
    :param str transit_router_id: The id of the transit router.
    """
    __args__ = dict()
    __args__['directConnectGatewayId'] = direct_connect_gateway_id
    __args__['outputFile'] = output_file
    __args__['transitRouterAttachmentIds'] = transit_router_attachment_ids
    __args__['transitRouterId'] = transit_router_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:transit_router/directConnectGatewayAttachments:DirectConnectGatewayAttachments', __args__, opts=opts, typ=DirectConnectGatewayAttachmentsResult).value

    return AwaitableDirectConnectGatewayAttachmentsResult(
        attachments=pulumi.get(__ret__, 'attachments'),
        direct_connect_gateway_id=pulumi.get(__ret__, 'direct_connect_gateway_id'),
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'),
        transit_router_attachment_ids=pulumi.get(__ret__, 'transit_router_attachment_ids'),
        transit_router_id=pulumi.get(__ret__, 'transit_router_id'))


@_utilities.lift_output_func(direct_connect_gateway_attachments)
def direct_connect_gateway_attachments_output(direct_connect_gateway_id: Optional[pulumi.Input[Optional[str]]] = None,
                                              output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                              transit_router_attachment_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                              transit_router_id: Optional[pulumi.Input[str]] = None,
                                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[DirectConnectGatewayAttachmentsResult]:
    """
    Use this data source to query detailed information of transit router direct connect gateway attachments
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.transit_router.direct_connect_gateway_attachments(transit_router_id="tr-2bzy39x27qtxc2dx0eg5qaj05")
    ```


    :param str direct_connect_gateway_id: ID of the direct connection gateway.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] transit_router_attachment_ids: ID of the network instance connection.
    :param str transit_router_id: The id of the transit router.
    """
    ...