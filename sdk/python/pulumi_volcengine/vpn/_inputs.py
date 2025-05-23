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
    'GatewayTagArgs',
    'GatewaysTagArgs',
    'GetGatewaysTagArgs',
]

@pulumi.input_type
class GatewayTagArgs:
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
class GatewaysTagArgs:
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
class GetGatewaysTagArgs:
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


