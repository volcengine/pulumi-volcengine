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
    'FileSystemTagArgs',
    'FileSystemsTagArgs',
    'PermissionGroupPermissionRuleArgs',
    'PermissionGroupsFilterArgs',
    'GetFileSystemsTagArgs',
    'GetPermissionGroupsFilterArgs',
]

@pulumi.input_type
class FileSystemTagArgs:
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
class FileSystemsTagArgs:
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
class PermissionGroupPermissionRuleArgs:
    def __init__(__self__, *,
                 cidr_ip: pulumi.Input[str],
                 rw_mode: pulumi.Input[str],
                 use_mode: pulumi.Input[str]):
        """
        :param pulumi.Input[str] cidr_ip: Client IP addresses that are allowed access.
        :param pulumi.Input[str] rw_mode: Permission group read and write rules. The value description is as follows:
               `RW`: Allows reading and writing.
               `RO`: read-only mode.
        :param pulumi.Input[str] use_mode: Permission group user permissions. The value description is as follows:
               `All_squash`: All access users are mapped to anonymous users or user groups.
               `No_all_squash`: The access user is first matched with the local user, and then mapped to an anonymous user or user group after the match fails.
               `Root_squash`: Map the Root user as an anonymous user or user group.
               `No_root_squash`: The Root user maintains the Root account authority.
        """
        pulumi.set(__self__, "cidr_ip", cidr_ip)
        pulumi.set(__self__, "rw_mode", rw_mode)
        pulumi.set(__self__, "use_mode", use_mode)

    @property
    @pulumi.getter(name="cidrIp")
    def cidr_ip(self) -> pulumi.Input[str]:
        """
        Client IP addresses that are allowed access.
        """
        return pulumi.get(self, "cidr_ip")

    @cidr_ip.setter
    def cidr_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr_ip", value)

    @property
    @pulumi.getter(name="rwMode")
    def rw_mode(self) -> pulumi.Input[str]:
        """
        Permission group read and write rules. The value description is as follows:
        `RW`: Allows reading and writing.
        `RO`: read-only mode.
        """
        return pulumi.get(self, "rw_mode")

    @rw_mode.setter
    def rw_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "rw_mode", value)

    @property
    @pulumi.getter(name="useMode")
    def use_mode(self) -> pulumi.Input[str]:
        """
        Permission group user permissions. The value description is as follows:
        `All_squash`: All access users are mapped to anonymous users or user groups.
        `No_all_squash`: The access user is first matched with the local user, and then mapped to an anonymous user or user group after the match fails.
        `Root_squash`: Map the Root user as an anonymous user or user group.
        `No_root_squash`: The Root user maintains the Root account authority.
        """
        return pulumi.get(self, "use_mode")

    @use_mode.setter
    def use_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "use_mode", value)


@pulumi.input_type
class PermissionGroupsFilterArgs:
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: Filters permission groups for specified characteristics based on attributes. The parameters that support filtering are as follows: `PermissionGroupName`, `PermissionGroupId`.
        :param str value: The value of the filter item.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Filters permission groups for specified characteristics based on attributes. The parameters that support filtering are as follows: `PermissionGroupName`, `PermissionGroupId`.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the filter item.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetFileSystemsTagArgs:
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
class GetPermissionGroupsFilterArgs:
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: Filters permission groups for specified characteristics based on attributes. The parameters that support filtering are as follows: `PermissionGroupName`, `PermissionGroupId`.
        :param str value: The value of the filter item.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Filters permission groups for specified characteristics based on attributes. The parameters that support filtering are as follows: `PermissionGroupName`, `PermissionGroupId`.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the filter item.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


