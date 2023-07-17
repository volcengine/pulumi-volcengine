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
    'DnatEntriesDnatEntryResult',
    'GatewayTag',
    'GatewaysNatGatewayResult',
    'GatewaysNatGatewayEipAddressResult',
    'GatewaysNatGatewayTagResult',
    'GatewaysTagResult',
    'SnatEntriesSnatEntryResult',
]

@pulumi.output_type
class DnatEntriesDnatEntryResult(dict):
    def __init__(__self__, *,
                 dnat_entry_id: str,
                 dnat_entry_name: str,
                 external_ip: str,
                 external_port: str,
                 internal_ip: str,
                 internal_port: str,
                 nat_gateway_id: str,
                 protocol: str,
                 status: str):
        """
        :param str dnat_entry_id: The ID of the DNAT entry.
        :param str dnat_entry_name: The name of the DNAT entry.
        :param str external_ip: Provides the public IP address for public network access.
        :param str external_port: The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
        :param str internal_ip: Provides the internal IP address.
        :param str internal_port: The port or port segment on which the cloud server instance provides services to the public network.
        :param str nat_gateway_id: The id of the NAT gateway.
        :param str protocol: The network protocol.
        :param str status: The network status.
        """
        pulumi.set(__self__, "dnat_entry_id", dnat_entry_id)
        pulumi.set(__self__, "dnat_entry_name", dnat_entry_name)
        pulumi.set(__self__, "external_ip", external_ip)
        pulumi.set(__self__, "external_port", external_port)
        pulumi.set(__self__, "internal_ip", internal_ip)
        pulumi.set(__self__, "internal_port", internal_port)
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="dnatEntryId")
    def dnat_entry_id(self) -> str:
        """
        The ID of the DNAT entry.
        """
        return pulumi.get(self, "dnat_entry_id")

    @property
    @pulumi.getter(name="dnatEntryName")
    def dnat_entry_name(self) -> str:
        """
        The name of the DNAT entry.
        """
        return pulumi.get(self, "dnat_entry_name")

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> str:
        """
        Provides the public IP address for public network access.
        """
        return pulumi.get(self, "external_ip")

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> str:
        """
        The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
        """
        return pulumi.get(self, "external_port")

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> str:
        """
        Provides the internal IP address.
        """
        return pulumi.get(self, "internal_ip")

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> str:
        """
        The port or port segment on which the cloud server instance provides services to the public network.
        """
        return pulumi.get(self, "internal_port")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> str:
        """
        The id of the NAT gateway.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The network protocol.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The network status.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class GatewayTag(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: The Key of Tags.
        :param str value: The Value of Tags.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The Key of Tags.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GatewaysNatGatewayResult(dict):
    def __init__(__self__, *,
                 billing_type: str,
                 business_status: str,
                 creation_time: str,
                 deleted_time: str,
                 description: str,
                 eip_addresses: Sequence['outputs.GatewaysNatGatewayEipAddressResult'],
                 id: str,
                 lock_reason: str,
                 nat_gateway_id: str,
                 nat_gateway_name: str,
                 network_interface_id: str,
                 overdue_time: str,
                 spec: str,
                 status: str,
                 subnet_id: str,
                 tags: Sequence['outputs.GatewaysNatGatewayTagResult'],
                 updated_at: str,
                 vpc_id: str):
        """
        :param str billing_type: The billing type of the NatGateway.
        :param str business_status: Whether the NatGateway is locked.
        :param str creation_time: The creation time of the NatGateway.
        :param str deleted_time: The deleted time of the NatGateway.
        :param str description: The description of the NatGateway.
        :param Sequence['GatewaysNatGatewayEipAddressArgs'] eip_addresses: The eip addresses of the NatGateway.
        :param str id: The ID of the NatGateway.
        :param str lock_reason: The reason why locking NatGateway.
        :param str nat_gateway_id: The ID of the NatGateway.
        :param str nat_gateway_name: The name of the NatGateway.
        :param str network_interface_id: The ID of the network interface.
        :param str overdue_time: The overdue time of the NatGateway.
        :param str spec: The specification of the NatGateway.
        :param str status: The status of the NatGateway.
        :param str subnet_id: The id of the Subnet.
        :param Sequence['GatewaysNatGatewayTagArgs'] tags: Tags.
        :param str updated_at: The update time of the NatGateway.
        :param str vpc_id: The id of the VPC.
        """
        pulumi.set(__self__, "billing_type", billing_type)
        pulumi.set(__self__, "business_status", business_status)
        pulumi.set(__self__, "creation_time", creation_time)
        pulumi.set(__self__, "deleted_time", deleted_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "eip_addresses", eip_addresses)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "lock_reason", lock_reason)
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        pulumi.set(__self__, "nat_gateway_name", nat_gateway_name)
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        pulumi.set(__self__, "overdue_time", overdue_time)
        pulumi.set(__self__, "spec", spec)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "updated_at", updated_at)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> str:
        """
        The billing type of the NatGateway.
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> str:
        """
        Whether the NatGateway is locked.
        """
        return pulumi.get(self, "business_status")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The creation time of the NatGateway.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> str:
        """
        The deleted time of the NatGateway.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the NatGateway.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eipAddresses")
    def eip_addresses(self) -> Sequence['outputs.GatewaysNatGatewayEipAddressResult']:
        """
        The eip addresses of the NatGateway.
        """
        return pulumi.get(self, "eip_addresses")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the NatGateway.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lockReason")
    def lock_reason(self) -> str:
        """
        The reason why locking NatGateway.
        """
        return pulumi.get(self, "lock_reason")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> str:
        """
        The ID of the NatGateway.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="natGatewayName")
    def nat_gateway_name(self) -> str:
        """
        The name of the NatGateway.
        """
        return pulumi.get(self, "nat_gateway_name")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> str:
        """
        The ID of the network interface.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="overdueTime")
    def overdue_time(self) -> str:
        """
        The overdue time of the NatGateway.
        """
        return pulumi.get(self, "overdue_time")

    @property
    @pulumi.getter
    def spec(self) -> str:
        """
        The specification of the NatGateway.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the NatGateway.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The id of the Subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GatewaysNatGatewayTagResult']:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The update time of the NatGateway.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The id of the VPC.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class GatewaysNatGatewayEipAddressResult(dict):
    def __init__(__self__, *,
                 allocation_id: str,
                 eip_address: str,
                 using_status: str):
        """
        :param str allocation_id: The ID of Eip.
        :param str eip_address: The address of Eip.
        :param str using_status: The using status of Eip.
        """
        pulumi.set(__self__, "allocation_id", allocation_id)
        pulumi.set(__self__, "eip_address", eip_address)
        pulumi.set(__self__, "using_status", using_status)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> str:
        """
        The ID of Eip.
        """
        return pulumi.get(self, "allocation_id")

    @property
    @pulumi.getter(name="eipAddress")
    def eip_address(self) -> str:
        """
        The address of Eip.
        """
        return pulumi.get(self, "eip_address")

    @property
    @pulumi.getter(name="usingStatus")
    def using_status(self) -> str:
        """
        The using status of Eip.
        """
        return pulumi.get(self, "using_status")


@pulumi.output_type
class GatewaysNatGatewayTagResult(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: The Key of Tags.
        :param str value: The Value of Tags.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The Key of Tags.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GatewaysTagResult(dict):
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: The Key of Tags.
        :param str value: The Value of Tags.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The Key of Tags.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class SnatEntriesSnatEntryResult(dict):
    def __init__(__self__, *,
                 eip_address: str,
                 eip_id: str,
                 id: str,
                 nat_gateway_id: str,
                 snat_entry_id: str,
                 snat_entry_name: str,
                 source_cidr: str,
                 status: str,
                 subnet_id: str):
        """
        :param str eip_address: The public ip address used by the SNAT entry.
        :param str eip_id: An id of the public ip address used by the SNAT entry.
        :param str id: The id of the SNAT entry.
        :param str nat_gateway_id: An id of the nat gateway to which the entry belongs.
        :param str snat_entry_id: The id of the SNAT entry.
        :param str snat_entry_name: A name of SNAT entry.
        :param str source_cidr: The SourceCidr of SNAT entry.
        :param str status: The status of the SNAT entry.
        :param str subnet_id: An id of the subnet that is required to access the Internet.
        """
        pulumi.set(__self__, "eip_address", eip_address)
        pulumi.set(__self__, "eip_id", eip_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        pulumi.set(__self__, "snat_entry_id", snat_entry_id)
        pulumi.set(__self__, "snat_entry_name", snat_entry_name)
        pulumi.set(__self__, "source_cidr", source_cidr)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="eipAddress")
    def eip_address(self) -> str:
        """
        The public ip address used by the SNAT entry.
        """
        return pulumi.get(self, "eip_address")

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> str:
        """
        An id of the public ip address used by the SNAT entry.
        """
        return pulumi.get(self, "eip_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the SNAT entry.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> str:
        """
        An id of the nat gateway to which the entry belongs.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="snatEntryId")
    def snat_entry_id(self) -> str:
        """
        The id of the SNAT entry.
        """
        return pulumi.get(self, "snat_entry_id")

    @property
    @pulumi.getter(name="snatEntryName")
    def snat_entry_name(self) -> str:
        """
        A name of SNAT entry.
        """
        return pulumi.get(self, "snat_entry_name")

    @property
    @pulumi.getter(name="sourceCidr")
    def source_cidr(self) -> str:
        """
        The SourceCidr of SNAT entry.
        """
        return pulumi.get(self, "source_cidr")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the SNAT entry.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        An id of the subnet that is required to access the Internet.
        """
        return pulumi.get(self, "subnet_id")

