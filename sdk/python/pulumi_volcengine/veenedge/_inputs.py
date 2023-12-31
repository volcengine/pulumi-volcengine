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
    'CloudServerBillingConfigArgs',
    'CloudServerCustomDataArgs',
    'CloudServerNetworkConfigArgs',
    'CloudServerScheduleStrategyArgs',
    'CloudServerStorageConfigArgs',
    'CloudServerStorageConfigDataDiskListArgs',
    'CloudServerStorageConfigSystemDiskArgs',
]

@pulumi.input_type
class CloudServerBillingConfigArgs:
    def __init__(__self__, *,
                 bandwidth_billing_method: pulumi.Input[str],
                 computing_billing_method: pulumi.Input[str]):
        """
        :param pulumi.Input[str] bandwidth_billing_method: The method of bandwidth billing. The value can be `MonthlyP95` or `DailyPeak`.
        :param pulumi.Input[str] computing_billing_method: The method of computing billing. The value can be `MonthlyPeak` or `DailyPeak`.
        """
        pulumi.set(__self__, "bandwidth_billing_method", bandwidth_billing_method)
        pulumi.set(__self__, "computing_billing_method", computing_billing_method)

    @property
    @pulumi.getter(name="bandwidthBillingMethod")
    def bandwidth_billing_method(self) -> pulumi.Input[str]:
        """
        The method of bandwidth billing. The value can be `MonthlyP95` or `DailyPeak`.
        """
        return pulumi.get(self, "bandwidth_billing_method")

    @bandwidth_billing_method.setter
    def bandwidth_billing_method(self, value: pulumi.Input[str]):
        pulumi.set(self, "bandwidth_billing_method", value)

    @property
    @pulumi.getter(name="computingBillingMethod")
    def computing_billing_method(self) -> pulumi.Input[str]:
        """
        The method of computing billing. The value can be `MonthlyPeak` or `DailyPeak`.
        """
        return pulumi.get(self, "computing_billing_method")

    @computing_billing_method.setter
    def computing_billing_method(self, value: pulumi.Input[str]):
        pulumi.set(self, "computing_billing_method", value)


@pulumi.input_type
class CloudServerCustomDataArgs:
    def __init__(__self__, *,
                 data: pulumi.Input[str]):
        """
        :param pulumi.Input[str] data: The custom data info.
        """
        pulumi.set(__self__, "data", data)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Input[str]:
        """
        The custom data info.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: pulumi.Input[str]):
        pulumi.set(self, "data", value)


@pulumi.input_type
class CloudServerNetworkConfigArgs:
    def __init__(__self__, *,
                 bandwidth_peak: pulumi.Input[str],
                 custom_external_interface_name: Optional[pulumi.Input[str]] = None,
                 custom_internal_interface_name: Optional[pulumi.Input[str]] = None,
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 internal_bandwidth_peak: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] bandwidth_peak: The peak of bandwidth.
        :param pulumi.Input[str] custom_external_interface_name: The name of custom external interface.
        :param pulumi.Input[str] custom_internal_interface_name: The name of custom internal interface.
        :param pulumi.Input[bool] enable_ipv6: Whether enable ipv6.
        :param pulumi.Input[str] internal_bandwidth_peak: The internal peak of bandwidth.
        """
        pulumi.set(__self__, "bandwidth_peak", bandwidth_peak)
        if custom_external_interface_name is not None:
            pulumi.set(__self__, "custom_external_interface_name", custom_external_interface_name)
        if custom_internal_interface_name is not None:
            pulumi.set(__self__, "custom_internal_interface_name", custom_internal_interface_name)
        if enable_ipv6 is not None:
            pulumi.set(__self__, "enable_ipv6", enable_ipv6)
        if internal_bandwidth_peak is not None:
            pulumi.set(__self__, "internal_bandwidth_peak", internal_bandwidth_peak)

    @property
    @pulumi.getter(name="bandwidthPeak")
    def bandwidth_peak(self) -> pulumi.Input[str]:
        """
        The peak of bandwidth.
        """
        return pulumi.get(self, "bandwidth_peak")

    @bandwidth_peak.setter
    def bandwidth_peak(self, value: pulumi.Input[str]):
        pulumi.set(self, "bandwidth_peak", value)

    @property
    @pulumi.getter(name="customExternalInterfaceName")
    def custom_external_interface_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of custom external interface.
        """
        return pulumi.get(self, "custom_external_interface_name")

    @custom_external_interface_name.setter
    def custom_external_interface_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_external_interface_name", value)

    @property
    @pulumi.getter(name="customInternalInterfaceName")
    def custom_internal_interface_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of custom internal interface.
        """
        return pulumi.get(self, "custom_internal_interface_name")

    @custom_internal_interface_name.setter
    def custom_internal_interface_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_internal_interface_name", value)

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether enable ipv6.
        """
        return pulumi.get(self, "enable_ipv6")

    @enable_ipv6.setter
    def enable_ipv6(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ipv6", value)

    @property
    @pulumi.getter(name="internalBandwidthPeak")
    def internal_bandwidth_peak(self) -> Optional[pulumi.Input[str]]:
        """
        The internal peak of bandwidth.
        """
        return pulumi.get(self, "internal_bandwidth_peak")

    @internal_bandwidth_peak.setter
    def internal_bandwidth_peak(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "internal_bandwidth_peak", value)


@pulumi.input_type
class CloudServerScheduleStrategyArgs:
    def __init__(__self__, *,
                 network_strategy: pulumi.Input[str],
                 price_strategy: pulumi.Input[str],
                 schedule_strategy: pulumi.Input[str]):
        """
        :param pulumi.Input[str] network_strategy: The network strategy.
        :param pulumi.Input[str] price_strategy: The price strategy. The value can be `high_priority` or `low_priority`.
        :param pulumi.Input[str] schedule_strategy: The type of schedule strategy. The value can be `dispersion` or `concentration`.
        """
        pulumi.set(__self__, "network_strategy", network_strategy)
        pulumi.set(__self__, "price_strategy", price_strategy)
        pulumi.set(__self__, "schedule_strategy", schedule_strategy)

    @property
    @pulumi.getter(name="networkStrategy")
    def network_strategy(self) -> pulumi.Input[str]:
        """
        The network strategy.
        """
        return pulumi.get(self, "network_strategy")

    @network_strategy.setter
    def network_strategy(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_strategy", value)

    @property
    @pulumi.getter(name="priceStrategy")
    def price_strategy(self) -> pulumi.Input[str]:
        """
        The price strategy. The value can be `high_priority` or `low_priority`.
        """
        return pulumi.get(self, "price_strategy")

    @price_strategy.setter
    def price_strategy(self, value: pulumi.Input[str]):
        pulumi.set(self, "price_strategy", value)

    @property
    @pulumi.getter(name="scheduleStrategy")
    def schedule_strategy(self) -> pulumi.Input[str]:
        """
        The type of schedule strategy. The value can be `dispersion` or `concentration`.
        """
        return pulumi.get(self, "schedule_strategy")

    @schedule_strategy.setter
    def schedule_strategy(self, value: pulumi.Input[str]):
        pulumi.set(self, "schedule_strategy", value)


@pulumi.input_type
class CloudServerStorageConfigArgs:
    def __init__(__self__, *,
                 system_disk: pulumi.Input['CloudServerStorageConfigSystemDiskArgs'],
                 data_disk_lists: Optional[pulumi.Input[Sequence[pulumi.Input['CloudServerStorageConfigDataDiskListArgs']]]] = None):
        """
        :param pulumi.Input['CloudServerStorageConfigSystemDiskArgs'] system_disk: The disk info of system.
        :param pulumi.Input[Sequence[pulumi.Input['CloudServerStorageConfigDataDiskListArgs']]] data_disk_lists: The disk list info of data.
        """
        pulumi.set(__self__, "system_disk", system_disk)
        if data_disk_lists is not None:
            pulumi.set(__self__, "data_disk_lists", data_disk_lists)

    @property
    @pulumi.getter(name="systemDisk")
    def system_disk(self) -> pulumi.Input['CloudServerStorageConfigSystemDiskArgs']:
        """
        The disk info of system.
        """
        return pulumi.get(self, "system_disk")

    @system_disk.setter
    def system_disk(self, value: pulumi.Input['CloudServerStorageConfigSystemDiskArgs']):
        pulumi.set(self, "system_disk", value)

    @property
    @pulumi.getter(name="dataDiskLists")
    def data_disk_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CloudServerStorageConfigDataDiskListArgs']]]]:
        """
        The disk list info of data.
        """
        return pulumi.get(self, "data_disk_lists")

    @data_disk_lists.setter
    def data_disk_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CloudServerStorageConfigDataDiskListArgs']]]]):
        pulumi.set(self, "data_disk_lists", value)


@pulumi.input_type
class CloudServerStorageConfigDataDiskListArgs:
    def __init__(__self__, *,
                 capacity: pulumi.Input[str],
                 storage_type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] capacity: The capacity of storage.
        :param pulumi.Input[str] storage_type: The type of storage. The value can be `CloudBlockHDD` or `CloudBlockSSD`.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "storage_type", storage_type)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Input[str]:
        """
        The capacity of storage.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: pulumi.Input[str]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Input[str]:
        """
        The type of storage. The value can be `CloudBlockHDD` or `CloudBlockSSD`.
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_type", value)


@pulumi.input_type
class CloudServerStorageConfigSystemDiskArgs:
    def __init__(__self__, *,
                 capacity: pulumi.Input[str],
                 storage_type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] capacity: The capacity of storage.
        :param pulumi.Input[str] storage_type: The type of storage. The value can be `CloudBlockHDD` or `CloudBlockSSD`.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "storage_type", storage_type)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Input[str]:
        """
        The capacity of storage.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: pulumi.Input[str]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Input[str]:
        """
        The type of storage. The value can be `CloudBlockHDD` or `CloudBlockSSD`.
        """
        return pulumi.get(self, "storage_type")

    @storage_type.setter
    def storage_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_type", value)


