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
    'GetVpcEndpointZonesResult',
    'AwaitableGetVpcEndpointZonesResult',
    'get_vpc_endpoint_zones',
    'get_vpc_endpoint_zones_output',
]

@pulumi.output_type
class GetVpcEndpointZonesResult:
    """
    A collection of values returned by getVpcEndpointZones.
    """
    def __init__(__self__, endpoint_id=None, id=None, output_file=None, total_count=None, vpc_endpoint_zones=None):
        if endpoint_id and not isinstance(endpoint_id, str):
            raise TypeError("Expected argument 'endpoint_id' to be a str")
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_endpoint_zones and not isinstance(vpc_endpoint_zones, list):
            raise TypeError("Expected argument 'vpc_endpoint_zones' to be a list")
        pulumi.set(__self__, "vpc_endpoint_zones", vpc_endpoint_zones)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[str]:
        return pulumi.get(self, "endpoint_id")

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
        Returns the total amount of the data list.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcEndpointZones")
    def vpc_endpoint_zones(self) -> Sequence['outputs.GetVpcEndpointZonesVpcEndpointZoneResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "vpc_endpoint_zones")


class AwaitableGetVpcEndpointZonesResult(GetVpcEndpointZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcEndpointZonesResult(
            endpoint_id=self.endpoint_id,
            id=self.id,
            output_file=self.output_file,
            total_count=self.total_count,
            vpc_endpoint_zones=self.vpc_endpoint_zones)


def get_vpc_endpoint_zones(endpoint_id: Optional[str] = None,
                           output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcEndpointZonesResult:
    """
    Use this data source to query detailed information of privatelink vpc endpoint zones
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
        security_group_name="acc-test-security-group",
        vpc_id=foo_vpc.id)
    foo_clb = volcengine.clb.Clb("fooClb",
        type="public",
        subnet_id=foo_subnet.id,
        load_balancer_spec="small_1",
        description="acc-test-demo",
        load_balancer_name="acc-test-clb",
        load_balancer_billing_type="PostPaid",
        eip_billing_config=volcengine.clb.ClbEipBillingConfigArgs(
            isp="BGP",
            eip_billing_type="PostPaidByBandwidth",
            bandwidth=1,
        ),
        tags=[volcengine.clb.ClbTagArgs(
            key="k1",
            value="v1",
        )])
    foo_vpc_endpoint_service = volcengine.privatelink.VpcEndpointService("fooVpcEndpointService",
        resources=[volcengine.privatelink.VpcEndpointServiceResourceArgs(
            resource_id=foo_clb.id,
            resource_type="CLB",
        )],
        description="acc-test",
        auto_accept_enabled=True)
    foo_vpc_endpoint = volcengine.privatelink.VpcEndpoint("fooVpcEndpoint",
        security_group_ids=[foo_security_group.id],
        service_id=foo_vpc_endpoint_service.id,
        endpoint_name="acc-test-ep",
        description="acc-test")
    foo_vpc_endpoint_zone = volcengine.privatelink.VpcEndpointZone("fooVpcEndpointZone",
        endpoint_id=foo_vpc_endpoint.id,
        subnet_id=foo_subnet.id,
        private_ip_address="172.16.0.251")
    foo_vpc_endpoint_zones = volcengine.privatelink.get_vpc_endpoint_zones_output(endpoint_id=foo_vpc_endpoint_zone.endpoint_id)
    ```


    :param str endpoint_id: The endpoint id of query.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['endpointId'] = endpoint_id
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:privatelink/getVpcEndpointZones:getVpcEndpointZones', __args__, opts=opts, typ=GetVpcEndpointZonesResult).value

    return AwaitableGetVpcEndpointZonesResult(
        endpoint_id=pulumi.get(__ret__, 'endpoint_id'),
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpc_endpoint_zones=pulumi.get(__ret__, 'vpc_endpoint_zones'))


@_utilities.lift_output_func(get_vpc_endpoint_zones)
def get_vpc_endpoint_zones_output(endpoint_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcEndpointZonesResult]:
    """
    Use this data source to query detailed information of privatelink vpc endpoint zones
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
        security_group_name="acc-test-security-group",
        vpc_id=foo_vpc.id)
    foo_clb = volcengine.clb.Clb("fooClb",
        type="public",
        subnet_id=foo_subnet.id,
        load_balancer_spec="small_1",
        description="acc-test-demo",
        load_balancer_name="acc-test-clb",
        load_balancer_billing_type="PostPaid",
        eip_billing_config=volcengine.clb.ClbEipBillingConfigArgs(
            isp="BGP",
            eip_billing_type="PostPaidByBandwidth",
            bandwidth=1,
        ),
        tags=[volcengine.clb.ClbTagArgs(
            key="k1",
            value="v1",
        )])
    foo_vpc_endpoint_service = volcengine.privatelink.VpcEndpointService("fooVpcEndpointService",
        resources=[volcengine.privatelink.VpcEndpointServiceResourceArgs(
            resource_id=foo_clb.id,
            resource_type="CLB",
        )],
        description="acc-test",
        auto_accept_enabled=True)
    foo_vpc_endpoint = volcengine.privatelink.VpcEndpoint("fooVpcEndpoint",
        security_group_ids=[foo_security_group.id],
        service_id=foo_vpc_endpoint_service.id,
        endpoint_name="acc-test-ep",
        description="acc-test")
    foo_vpc_endpoint_zone = volcengine.privatelink.VpcEndpointZone("fooVpcEndpointZone",
        endpoint_id=foo_vpc_endpoint.id,
        subnet_id=foo_subnet.id,
        private_ip_address="172.16.0.251")
    foo_vpc_endpoint_zones = volcengine.privatelink.get_vpc_endpoint_zones_output(endpoint_id=foo_vpc_endpoint_zone.endpoint_id)
    ```


    :param str endpoint_id: The endpoint id of query.
    :param str output_file: File name where to save data source results.
    """
    ...
