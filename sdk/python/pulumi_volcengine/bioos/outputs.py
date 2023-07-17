# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClusterSharedConfig',
    'ClusterVkeConfig',
    'ClustersItemResult',
    'WorkspacesItemResult',
]

@pulumi.output_type
class ClusterSharedConfig(dict):
    def __init__(__self__, *,
                 enable: bool):
        """
        :param bool enable: Whether to enable a shared cluster.
        """
        pulumi.set(__self__, "enable", enable)

    @property
    @pulumi.getter
    def enable(self) -> bool:
        """
        Whether to enable a shared cluster.
        """
        return pulumi.get(self, "enable")


@pulumi.output_type
class ClusterVkeConfig(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "clusterId":
            suggest = "cluster_id"
        elif key == "storageClass":
            suggest = "storage_class"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClusterVkeConfig. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClusterVkeConfig.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClusterVkeConfig.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cluster_id: str,
                 storage_class: str):
        """
        :param str cluster_id: The id of the vke cluster.
        :param str storage_class: The name of the StorageClass that the vke cluster has installed.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "storage_class", storage_class)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The id of the vke cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> str:
        """
        The name of the StorageClass that the vke cluster has installed.
        """
        return pulumi.get(self, "storage_class")


@pulumi.output_type
class ClustersItemResult(dict):
    def __init__(__self__, *,
                 bound: bool,
                 description: str,
                 id: str,
                 name: str,
                 public: bool,
                 start_time: int,
                 stopped_time: int,
                 vke_config_id: str,
                 vke_config_storage_class: str):
        """
        :param bool bound: Whether there is a bound workspace.
        :param str description: The description of the cluster.
        :param str id: The id of the bioos cluster.
        :param str name: The name of the cluster.
        :param bool public: whether it is a public cluster.
        :param int start_time: The start time of the cluster.
        :param int stopped_time: The end time of the cluster.
        :param str vke_config_id: The id of the vke cluster.
        :param str vke_config_storage_class: The name of the StorageClass that the vke cluster has installed.
        """
        pulumi.set(__self__, "bound", bound)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "public", public)
        pulumi.set(__self__, "start_time", start_time)
        pulumi.set(__self__, "stopped_time", stopped_time)
        pulumi.set(__self__, "vke_config_id", vke_config_id)
        pulumi.set(__self__, "vke_config_storage_class", vke_config_storage_class)

    @property
    @pulumi.getter
    def bound(self) -> bool:
        """
        Whether there is a bound workspace.
        """
        return pulumi.get(self, "bound")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the bioos cluster.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def public(self) -> bool:
        """
        whether it is a public cluster.
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> int:
        """
        The start time of the cluster.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="stoppedTime")
    def stopped_time(self) -> int:
        """
        The end time of the cluster.
        """
        return pulumi.get(self, "stopped_time")

    @property
    @pulumi.getter(name="vkeConfigId")
    def vke_config_id(self) -> str:
        """
        The id of the vke cluster.
        """
        return pulumi.get(self, "vke_config_id")

    @property
    @pulumi.getter(name="vkeConfigStorageClass")
    def vke_config_storage_class(self) -> str:
        """
        The name of the StorageClass that the vke cluster has installed.
        """
        return pulumi.get(self, "vke_config_storage_class")


@pulumi.output_type
class WorkspacesItemResult(dict):
    def __init__(__self__, *,
                 cover_download_url: str,
                 create_time: int,
                 description: str,
                 id: str,
                 name: str,
                 owner_name: str,
                 role: str,
                 s3_bucket: str,
                 update_time: int):
        """
        :param str cover_download_url: The URL of the cover.
        :param int create_time: The creation time of the workspace.
        :param str description: The description of the workspace.
        :param str id: The id of the workspace.
        :param str name: The name of the workspace.
        :param str owner_name: The name of the owner of the workspace.
        :param str role: The role of the user.
        :param str s3_bucket: S3 bucket address.
        :param int update_time: The update time of the workspace.
        """
        pulumi.set(__self__, "cover_download_url", cover_download_url)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "owner_name", owner_name)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "s3_bucket", s3_bucket)
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="coverDownloadUrl")
    def cover_download_url(self) -> str:
        """
        The URL of the cover.
        """
        return pulumi.get(self, "cover_download_url")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> int:
        """
        The creation time of the workspace.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the workspace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the workspace.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerName")
    def owner_name(self) -> str:
        """
        The name of the owner of the workspace.
        """
        return pulumi.get(self, "owner_name")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The role of the user.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> str:
        """
        S3 bucket address.
        """
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> int:
        """
        The update time of the workspace.
        """
        return pulumi.get(self, "update_time")

