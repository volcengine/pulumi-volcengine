# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AllowlistAssociatedInstanceArgs',
    'InstanceChargeDetailArgs',
    'InstanceChargeInfoArgs',
    'InstanceEndpointArgs',
    'InstanceEndpointAddressArgs',
    'InstanceEndpointReadOnlyNodeWeightArgs',
    'InstanceNodeArgs',
    'InstanceParameterArgs',
    'InstanceTagArgs',
    'InstancesTagArgs',
    'GetInstancesTagArgs',
]

@pulumi.input_type
class AllowlistAssociatedInstanceArgs:
    def __init__(__self__, *,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_name: Optional[pulumi.Input[str]] = None,
                 vpc: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] instance_id: The id of the postgresql instance.
        :param pulumi.Input[str] instance_name: The name of the postgresql instance.
        :param pulumi.Input[str] vpc: The id of the vpc.
        """
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_name is not None:
            pulumi.set(__self__, "instance_name", instance_name)
        if vpc is not None:
            pulumi.set(__self__, "vpc", vpc)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the postgresql instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the postgresql instance.
        """
        return pulumi.get(self, "instance_name")

    @instance_name.setter
    def instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_name", value)

    @property
    @pulumi.getter
    def vpc(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the vpc.
        """
        return pulumi.get(self, "vpc")

    @vpc.setter
    def vpc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc", value)


@pulumi.input_type
class InstanceChargeDetailArgs:
    def __init__(__self__, *,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 charge_end_time: Optional[pulumi.Input[str]] = None,
                 charge_start_time: Optional[pulumi.Input[str]] = None,
                 charge_status: Optional[pulumi.Input[str]] = None,
                 charge_type: Optional[pulumi.Input[str]] = None,
                 overdue_reclaim_time: Optional[pulumi.Input[str]] = None,
                 overdue_time: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 period_unit: Optional[pulumi.Input[str]] = None,
                 temp_modify_end_time: Optional[pulumi.Input[str]] = None,
                 temp_modify_start_time: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] auto_renew: Whether to automatically renew in prepaid scenarios.
               Autorenew_Enable
               Autorenew_Disable (default).
        :param pulumi.Input[str] charge_end_time: Billing expiry time (yearly and monthly only).
        :param pulumi.Input[str] charge_start_time: Billing start time (pay-as-you-go & monthly subscription).
        :param pulumi.Input[str] charge_status: Pay status. Value:
               normal - normal
               overdue - overdue
               .
        :param pulumi.Input[str] charge_type: Payment type. Value:
               PostPaid - Pay-As-You-Go
               PrePaid - Yearly and monthly (default).
        :param pulumi.Input[str] overdue_reclaim_time: Estimated release time when arrears are closed (pay-as-you-go & monthly subscription).
        :param pulumi.Input[str] overdue_time: Shutdown time in arrears (pay-as-you-go & monthly subscription).
        :param pulumi.Input[int] period: Purchase duration in prepaid scenarios. Default: 1.
        :param pulumi.Input[str] period_unit: The purchase cycle in the prepaid scenario.
               Month - monthly subscription (default)
               Year - Package year.
        :param pulumi.Input[str] temp_modify_end_time: Temporary upgrade of restoration time.
        :param pulumi.Input[str] temp_modify_start_time: Start time of temporary upgrade.
        """
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if charge_end_time is not None:
            pulumi.set(__self__, "charge_end_time", charge_end_time)
        if charge_start_time is not None:
            pulumi.set(__self__, "charge_start_time", charge_start_time)
        if charge_status is not None:
            pulumi.set(__self__, "charge_status", charge_status)
        if charge_type is not None:
            pulumi.set(__self__, "charge_type", charge_type)
        if overdue_reclaim_time is not None:
            pulumi.set(__self__, "overdue_reclaim_time", overdue_reclaim_time)
        if overdue_time is not None:
            pulumi.set(__self__, "overdue_time", overdue_time)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)
        if temp_modify_end_time is not None:
            pulumi.set(__self__, "temp_modify_end_time", temp_modify_end_time)
        if temp_modify_start_time is not None:
            pulumi.set(__self__, "temp_modify_start_time", temp_modify_start_time)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to automatically renew in prepaid scenarios.
        Autorenew_Enable
        Autorenew_Disable (default).
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="chargeEndTime")
    def charge_end_time(self) -> Optional[pulumi.Input[str]]:
        """
        Billing expiry time (yearly and monthly only).
        """
        return pulumi.get(self, "charge_end_time")

    @charge_end_time.setter
    def charge_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_end_time", value)

    @property
    @pulumi.getter(name="chargeStartTime")
    def charge_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        Billing start time (pay-as-you-go & monthly subscription).
        """
        return pulumi.get(self, "charge_start_time")

    @charge_start_time.setter
    def charge_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_start_time", value)

    @property
    @pulumi.getter(name="chargeStatus")
    def charge_status(self) -> Optional[pulumi.Input[str]]:
        """
        Pay status. Value:
        normal - normal
        overdue - overdue
        .
        """
        return pulumi.get(self, "charge_status")

    @charge_status.setter
    def charge_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_status", value)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> Optional[pulumi.Input[str]]:
        """
        Payment type. Value:
        PostPaid - Pay-As-You-Go
        PrePaid - Yearly and monthly (default).
        """
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="overdueReclaimTime")
    def overdue_reclaim_time(self) -> Optional[pulumi.Input[str]]:
        """
        Estimated release time when arrears are closed (pay-as-you-go & monthly subscription).
        """
        return pulumi.get(self, "overdue_reclaim_time")

    @overdue_reclaim_time.setter
    def overdue_reclaim_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "overdue_reclaim_time", value)

    @property
    @pulumi.getter(name="overdueTime")
    def overdue_time(self) -> Optional[pulumi.Input[str]]:
        """
        Shutdown time in arrears (pay-as-you-go & monthly subscription).
        """
        return pulumi.get(self, "overdue_time")

    @overdue_time.setter
    def overdue_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "overdue_time", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Purchase duration in prepaid scenarios. Default: 1.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[str]]:
        """
        The purchase cycle in the prepaid scenario.
        Month - monthly subscription (default)
        Year - Package year.
        """
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period_unit", value)

    @property
    @pulumi.getter(name="tempModifyEndTime")
    def temp_modify_end_time(self) -> Optional[pulumi.Input[str]]:
        """
        Temporary upgrade of restoration time.
        """
        return pulumi.get(self, "temp_modify_end_time")

    @temp_modify_end_time.setter
    def temp_modify_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "temp_modify_end_time", value)

    @property
    @pulumi.getter(name="tempModifyStartTime")
    def temp_modify_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        Start time of temporary upgrade.
        """
        return pulumi.get(self, "temp_modify_start_time")

    @temp_modify_start_time.setter
    def temp_modify_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "temp_modify_start_time", value)


@pulumi.input_type
class InstanceChargeInfoArgs:
    def __init__(__self__, *,
                 charge_type: pulumi.Input[str],
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 period_unit: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] auto_renew: Whether to automatically renew in prepaid scenarios.
        :param pulumi.Input[int] period: Purchase duration in prepaid scenarios. Default: 1.
        :param pulumi.Input[str] period_unit: The purchase cycle in the prepaid scenario.
               Month - monthly subscription (default)
               Year - Package year.
        """
        pulumi.set(__self__, "charge_type", charge_type)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "charge_type")

    @charge_type.setter
    def charge_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "charge_type", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to automatically renew in prepaid scenarios.
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Purchase duration in prepaid scenarios. Default: 1.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[str]]:
        """
        The purchase cycle in the prepaid scenario.
        Month - monthly subscription (default)
        Year - Package year.
        """
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period_unit", value)


@pulumi.input_type
class InstanceEndpointArgs:
    def __init__(__self__, *,
                 addresses: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceEndpointAddressArgs']]]] = None,
                 auto_add_new_nodes: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_read_only: Optional[pulumi.Input[str]] = None,
                 enable_read_write_splitting: Optional[pulumi.Input[str]] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 endpoint_name: Optional[pulumi.Input[str]] = None,
                 endpoint_type: Optional[pulumi.Input[str]] = None,
                 read_only_node_weights: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceEndpointReadOnlyNodeWeightArgs']]]] = None,
                 read_write_mode: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input['InstanceEndpointAddressArgs']]] addresses: Address list.
        :param pulumi.Input[str] auto_add_new_nodes: When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.
        :param pulumi.Input[str] description: Address description.
        :param pulumi.Input[str] enable_read_only: Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.
        :param pulumi.Input[str] enable_read_write_splitting: Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.
        :param pulumi.Input[str] endpoint_id: Instance connection terminal ID.
        :param pulumi.Input[str] endpoint_name: The instance connection terminal name.
        :param pulumi.Input[str] endpoint_type: Terminal type:
               Cluster: The default terminal. (created by default)
               Primary: Primary node terminal.
               Custom: Custom terminal.
               Direct: Direct connection to the terminal. (Only the operation and maintenance side)
               AllNode: All node terminals. (Only the operation and maintenance side).
        :param pulumi.Input[Sequence[pulumi.Input['InstanceEndpointReadOnlyNodeWeightArgs']]] read_only_node_weights: The list of nodes configured by the connection terminal and the corresponding read-only weights.
        :param pulumi.Input[str] read_write_mode: Read and write mode:
               ReadWrite: read and write
               ReadOnly: read only (default).
        """
        if addresses is not None:
            pulumi.set(__self__, "addresses", addresses)
        if auto_add_new_nodes is not None:
            pulumi.set(__self__, "auto_add_new_nodes", auto_add_new_nodes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_read_only is not None:
            pulumi.set(__self__, "enable_read_only", enable_read_only)
        if enable_read_write_splitting is not None:
            pulumi.set(__self__, "enable_read_write_splitting", enable_read_write_splitting)
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if endpoint_name is not None:
            pulumi.set(__self__, "endpoint_name", endpoint_name)
        if endpoint_type is not None:
            pulumi.set(__self__, "endpoint_type", endpoint_type)
        if read_only_node_weights is not None:
            pulumi.set(__self__, "read_only_node_weights", read_only_node_weights)
        if read_write_mode is not None:
            pulumi.set(__self__, "read_write_mode", read_write_mode)

    @property
    @pulumi.getter
    def addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceEndpointAddressArgs']]]]:
        """
        Address list.
        """
        return pulumi.get(self, "addresses")

    @addresses.setter
    def addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceEndpointAddressArgs']]]]):
        pulumi.set(self, "addresses", value)

    @property
    @pulumi.getter(name="autoAddNewNodes")
    def auto_add_new_nodes(self) -> Optional[pulumi.Input[str]]:
        """
        When the terminal type is read-write terminal or read-only terminal, it supports setting whether new nodes are automatically added.
        """
        return pulumi.get(self, "auto_add_new_nodes")

    @auto_add_new_nodes.setter
    def auto_add_new_nodes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_add_new_nodes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Address description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableReadOnly")
    def enable_read_only(self) -> Optional[pulumi.Input[str]]:
        """
        Whether global read-only is enabled, value: Enable: Enable. Disable: Disabled.
        """
        return pulumi.get(self, "enable_read_only")

    @enable_read_only.setter
    def enable_read_only(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enable_read_only", value)

    @property
    @pulumi.getter(name="enableReadWriteSplitting")
    def enable_read_write_splitting(self) -> Optional[pulumi.Input[str]]:
        """
        Whether read-write separation is enabled, value: Enable: Enable. Disable: Disabled.
        """
        return pulumi.get(self, "enable_read_write_splitting")

    @enable_read_write_splitting.setter
    def enable_read_write_splitting(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enable_read_write_splitting", value)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance connection terminal ID.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> Optional[pulumi.Input[str]]:
        """
        The instance connection terminal name.
        """
        return pulumi.get(self, "endpoint_name")

    @endpoint_name.setter
    def endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_name", value)

    @property
    @pulumi.getter(name="endpointType")
    def endpoint_type(self) -> Optional[pulumi.Input[str]]:
        """
        Terminal type:
        Cluster: The default terminal. (created by default)
        Primary: Primary node terminal.
        Custom: Custom terminal.
        Direct: Direct connection to the terminal. (Only the operation and maintenance side)
        AllNode: All node terminals. (Only the operation and maintenance side).
        """
        return pulumi.get(self, "endpoint_type")

    @endpoint_type.setter
    def endpoint_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_type", value)

    @property
    @pulumi.getter(name="readOnlyNodeWeights")
    def read_only_node_weights(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceEndpointReadOnlyNodeWeightArgs']]]]:
        """
        The list of nodes configured by the connection terminal and the corresponding read-only weights.
        """
        return pulumi.get(self, "read_only_node_weights")

    @read_only_node_weights.setter
    def read_only_node_weights(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceEndpointReadOnlyNodeWeightArgs']]]]):
        pulumi.set(self, "read_only_node_weights", value)

    @property
    @pulumi.getter(name="readWriteMode")
    def read_write_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Read and write mode:
        ReadWrite: read and write
        ReadOnly: read only (default).
        """
        return pulumi.get(self, "read_write_mode")

    @read_write_mode.setter
    def read_write_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "read_write_mode", value)


@pulumi.input_type
class InstanceEndpointAddressArgs:
    def __init__(__self__, *,
                 dns_visibility: Optional[pulumi.Input[bool]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 eip_id: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 network_type: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[bool] dns_visibility: DNS Visibility.
        :param pulumi.Input[str] domain: Connect domain name.
        :param pulumi.Input[str] eip_id: The ID of the EIP, only valid for Public addresses.
        :param pulumi.Input[str] ip_address: The IP Address.
        :param pulumi.Input[str] network_type: Network address type, temporarily Private, Public, PublicService.
        :param pulumi.Input[str] port: The Port.
        :param pulumi.Input[str] subnet_id: Subnet ID of the RDS PostgreSQL instance.
        """
        if dns_visibility is not None:
            pulumi.set(__self__, "dns_visibility", dns_visibility)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if eip_id is not None:
            pulumi.set(__self__, "eip_id", eip_id)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if network_type is not None:
            pulumi.set(__self__, "network_type", network_type)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="dnsVisibility")
    def dns_visibility(self) -> Optional[pulumi.Input[bool]]:
        """
        DNS Visibility.
        """
        return pulumi.get(self, "dns_visibility")

    @dns_visibility.setter
    def dns_visibility(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dns_visibility", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        Connect domain name.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the EIP, only valid for Public addresses.
        """
        return pulumi.get(self, "eip_id")

    @eip_id.setter
    def eip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_id", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP Address.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> Optional[pulumi.Input[str]]:
        """
        Network address type, temporarily Private, Public, PublicService.
        """
        return pulumi.get(self, "network_type")

    @network_type.setter
    def network_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_type", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        """
        The Port.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet ID of the RDS PostgreSQL instance.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class InstanceEndpointReadOnlyNodeWeightArgs:
    def __init__(__self__, *,
                 node_id: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] node_id: Node ID.
        :param pulumi.Input[str] node_type: Node type. Value: Primary: Primary node.
               Secondary: Standby node.
               ReadOnly: Read-only node.
        :param pulumi.Input[int] weight: The weight of the node.
        """
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)
        if node_type is not None:
            pulumi.set(__self__, "node_type", node_type)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        Node ID.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[pulumi.Input[str]]:
        """
        Node type. Value: Primary: Primary node.
        Secondary: Standby node.
        ReadOnly: Read-only node.
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight of the node.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class InstanceNodeArgs:
    def __init__(__self__, *,
                 create_time: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[int]] = None,
                 node_id: Optional[pulumi.Input[str]] = None,
                 node_spec: Optional[pulumi.Input[str]] = None,
                 node_status: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 region_id: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 v_cpu: Optional[pulumi.Input[int]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] create_time: Node creation local time.
        :param pulumi.Input[str] instance_id: Instance ID.
        :param pulumi.Input[int] memory: Memory size in GB.
        :param pulumi.Input[str] node_id: Node ID.
        :param pulumi.Input[str] node_spec: The specification of primary node and secondary node.
        :param pulumi.Input[str] node_status: Node state, value: aligned with instance state.
        :param pulumi.Input[str] node_type: Node type. Value: Primary: Primary node.
               Secondary: Standby node.
               ReadOnly: Read-only node.
        :param pulumi.Input[str] region_id: The region of the RDS PostgreSQL instance.
        :param pulumi.Input[str] update_time: The update time of the RDS PostgreSQL instance.
        :param pulumi.Input[int] v_cpu: CPU size.
        :param pulumi.Input[str] zone_id: The available zone of the RDS PostgreSQL instance.
        """
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if memory is not None:
            pulumi.set(__self__, "memory", memory)
        if node_id is not None:
            pulumi.set(__self__, "node_id", node_id)
        if node_spec is not None:
            pulumi.set(__self__, "node_spec", node_spec)
        if node_status is not None:
            pulumi.set(__self__, "node_status", node_status)
        if node_type is not None:
            pulumi.set(__self__, "node_type", node_type)
        if region_id is not None:
            pulumi.set(__self__, "region_id", region_id)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if v_cpu is not None:
            pulumi.set(__self__, "v_cpu", v_cpu)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Node creation local time.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Instance ID.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def memory(self) -> Optional[pulumi.Input[int]]:
        """
        Memory size in GB.
        """
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory", value)

    @property
    @pulumi.getter(name="nodeId")
    def node_id(self) -> Optional[pulumi.Input[str]]:
        """
        Node ID.
        """
        return pulumi.get(self, "node_id")

    @node_id.setter
    def node_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_id", value)

    @property
    @pulumi.getter(name="nodeSpec")
    def node_spec(self) -> Optional[pulumi.Input[str]]:
        """
        The specification of primary node and secondary node.
        """
        return pulumi.get(self, "node_spec")

    @node_spec.setter
    def node_spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_spec", value)

    @property
    @pulumi.getter(name="nodeStatus")
    def node_status(self) -> Optional[pulumi.Input[str]]:
        """
        Node state, value: aligned with instance state.
        """
        return pulumi.get(self, "node_status")

    @node_status.setter
    def node_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_status", value)

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> Optional[pulumi.Input[str]]:
        """
        Node type. Value: Primary: Primary node.
        Secondary: Standby node.
        ReadOnly: Read-only node.
        """
        return pulumi.get(self, "node_type")

    @node_type.setter
    def node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_type", value)

    @property
    @pulumi.getter(name="regionId")
    def region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The region of the RDS PostgreSQL instance.
        """
        return pulumi.get(self, "region_id")

    @region_id.setter
    def region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region_id", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of the RDS PostgreSQL instance.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter(name="vCpu")
    def v_cpu(self) -> Optional[pulumi.Input[int]]:
        """
        CPU size.
        """
        return pulumi.get(self, "v_cpu")

    @v_cpu.setter
    def v_cpu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "v_cpu", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The available zone of the RDS PostgreSQL instance.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class InstanceParameterArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] name: Parameter name.
        :param pulumi.Input[str] value: Parameter value.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Parameter name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Parameter value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class InstanceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: The Key of Tags.
        :param pulumi.Input[str] value: The Value of Tags.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The Key of Tags.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class InstancesTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetInstancesTagArgs:
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

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


