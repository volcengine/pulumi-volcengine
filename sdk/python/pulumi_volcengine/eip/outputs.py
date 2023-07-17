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
    'AddressTag',
    'AddressesAddressResult',
    'AddressesAddressTagResult',
    'AddressesTagResult',
]

@pulumi.output_type
class AddressTag(dict):
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
class AddressesAddressResult(dict):
    def __init__(__self__, *,
                 allocation_id: str,
                 allocation_time: str,
                 bandwidth: int,
                 billing_type: str,
                 business_status: str,
                 deleted_time: str,
                 description: str,
                 eip_address: str,
                 expired_time: str,
                 id: str,
                 instance_id: str,
                 instance_type: str,
                 isp: str,
                 lock_reason: str,
                 name: str,
                 overdue_time: str,
                 project_name: str,
                 status: str,
                 tags: Sequence['outputs.AddressesAddressTagResult'],
                 updated_at: str):
        """
        :param str allocation_id: The id of the EIP address.
        :param str allocation_time: The allocation time of the EIP.
        :param int bandwidth: The peek bandwidth of the EIP.
        :param str billing_type: The billing type of the EIP.
        :param str business_status: The business status of the EIP.
        :param str deleted_time: The deleted time of the EIP.
        :param str description: The description of the EIP.
        :param str eip_address: The EIP ip address of the EIP.
        :param str expired_time: The expired time of the EIP.
        :param str id: The id of the EIP address.
        :param str instance_id: The instance id which be associated to the EIP.
        :param str instance_type: The type of the associated instance.
        :param str isp: An ISP of EIP Address, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.
        :param str lock_reason: The lock reason of the EIP.
        :param str name: A name of EIP.
        :param str overdue_time: The overdue time of the EIP.
        :param str project_name: The ProjectName of EIP.
        :param str status: A status of EIP, the value can be `Attaching` or `Detaching` or `Attached` or `Available`.
        :param Sequence['AddressesAddressTagArgs'] tags: Tags.
        :param str updated_at: The last update time of the EIP.
        """
        pulumi.set(__self__, "allocation_id", allocation_id)
        pulumi.set(__self__, "allocation_time", allocation_time)
        pulumi.set(__self__, "bandwidth", bandwidth)
        pulumi.set(__self__, "billing_type", billing_type)
        pulumi.set(__self__, "business_status", business_status)
        pulumi.set(__self__, "deleted_time", deleted_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "eip_address", eip_address)
        pulumi.set(__self__, "expired_time", expired_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_type", instance_type)
        pulumi.set(__self__, "isp", isp)
        pulumi.set(__self__, "lock_reason", lock_reason)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "overdue_time", overdue_time)
        pulumi.set(__self__, "project_name", project_name)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> str:
        """
        The id of the EIP address.
        """
        return pulumi.get(self, "allocation_id")

    @property
    @pulumi.getter(name="allocationTime")
    def allocation_time(self) -> str:
        """
        The allocation time of the EIP.
        """
        return pulumi.get(self, "allocation_time")

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        """
        The peek bandwidth of the EIP.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> str:
        """
        The billing type of the EIP.
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> str:
        """
        The business status of the EIP.
        """
        return pulumi.get(self, "business_status")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> str:
        """
        The deleted time of the EIP.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the EIP.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eipAddress")
    def eip_address(self) -> str:
        """
        The EIP ip address of the EIP.
        """
        return pulumi.get(self, "eip_address")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> str:
        """
        The expired time of the EIP.
        """
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the EIP address.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The instance id which be associated to the EIP.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> str:
        """
        The type of the associated instance.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def isp(self) -> str:
        """
        An ISP of EIP Address, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.
        """
        return pulumi.get(self, "isp")

    @property
    @pulumi.getter(name="lockReason")
    def lock_reason(self) -> str:
        """
        The lock reason of the EIP.
        """
        return pulumi.get(self, "lock_reason")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A name of EIP.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overdueTime")
    def overdue_time(self) -> str:
        """
        The overdue time of the EIP.
        """
        return pulumi.get(self, "overdue_time")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> str:
        """
        The ProjectName of EIP.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        A status of EIP, the value can be `Attaching` or `Detaching` or `Attached` or `Available`.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.AddressesAddressTagResult']:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        """
        The last update time of the EIP.
        """
        return pulumi.get(self, "updated_at")


@pulumi.output_type
class AddressesAddressTagResult(dict):
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
class AddressesTagResult(dict):
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

