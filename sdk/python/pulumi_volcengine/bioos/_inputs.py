# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterSharedConfigArgs',
    'ClusterVkeConfigArgs',
]

@pulumi.input_type
class ClusterSharedConfigArgs:
    def __init__(__self__, *,
                 enable: pulumi.Input[bool]):
        """
        :param pulumi.Input[bool] enable: Whether to enable a shared cluster.
        """
        pulumi.set(__self__, "enable", enable)

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Input[bool]:
        """
        Whether to enable a shared cluster.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enable", value)


@pulumi.input_type
class ClusterVkeConfigArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 storage_class: pulumi.Input[str]):
        """
        :param pulumi.Input[str] cluster_id: The id of the vke cluster.
        :param pulumi.Input[str] storage_class: The name of the StorageClass that the vke cluster has installed.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "storage_class", storage_class)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The id of the vke cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> pulumi.Input[str]:
        """
        The name of the StorageClass that the vke cluster has installed.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "storage_class", value)

