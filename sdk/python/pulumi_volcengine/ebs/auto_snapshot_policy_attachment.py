# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AutoSnapshotPolicyAttachmentArgs', 'AutoSnapshotPolicyAttachment']

@pulumi.input_type
class AutoSnapshotPolicyAttachmentArgs:
    def __init__(__self__, *,
                 auto_snapshot_policy_id: pulumi.Input[str],
                 volume_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a AutoSnapshotPolicyAttachment resource.
        :param pulumi.Input[str] auto_snapshot_policy_id: The id of the auto snapshot policy.
        :param pulumi.Input[str] volume_id: The id of the volume.
        """
        pulumi.set(__self__, "auto_snapshot_policy_id", auto_snapshot_policy_id)
        pulumi.set(__self__, "volume_id", volume_id)

    @property
    @pulumi.getter(name="autoSnapshotPolicyId")
    def auto_snapshot_policy_id(self) -> pulumi.Input[str]:
        """
        The id of the auto snapshot policy.
        """
        return pulumi.get(self, "auto_snapshot_policy_id")

    @auto_snapshot_policy_id.setter
    def auto_snapshot_policy_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_snapshot_policy_id", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Input[str]:
        """
        The id of the volume.
        """
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "volume_id", value)


@pulumi.input_type
class _AutoSnapshotPolicyAttachmentState:
    def __init__(__self__, *,
                 auto_snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AutoSnapshotPolicyAttachment resources.
        :param pulumi.Input[str] auto_snapshot_policy_id: The id of the auto snapshot policy.
        :param pulumi.Input[str] volume_id: The id of the volume.
        """
        if auto_snapshot_policy_id is not None:
            pulumi.set(__self__, "auto_snapshot_policy_id", auto_snapshot_policy_id)
        if volume_id is not None:
            pulumi.set(__self__, "volume_id", volume_id)

    @property
    @pulumi.getter(name="autoSnapshotPolicyId")
    def auto_snapshot_policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the auto snapshot policy.
        """
        return pulumi.get(self, "auto_snapshot_policy_id")

    @auto_snapshot_policy_id.setter
    def auto_snapshot_policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_snapshot_policy_id", value)

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the volume.
        """
        return pulumi.get(self, "volume_id")

    @volume_id.setter
    def volume_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_id", value)


class AutoSnapshotPolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage ebs auto snapshot policy attachment
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.zones()
        foo_volume = volcengine.ebs.Volume("fooVolume",
            volume_name="acc-test-volume",
            volume_type="ESSD_PL0",
            description="acc-test",
            kind="data",
            size=500,
            zone_id=foo_zones.zones[0].id,
            volume_charge_type="PostPaid",
            project_name="default")
        foo_auto_snapshot_policy = volcengine.ebs.AutoSnapshotPolicy("fooAutoSnapshotPolicy",
            auto_snapshot_policy_name="acc-test-auto-snapshot-policy",
            time_points=[
                "1",
                "5",
                "9",
            ],
            retention_days=-1,
            repeat_weekdays=[
                "2",
                "6",
            ],
            project_name="default",
            tags=[volcengine.ebs.AutoSnapshotPolicyTagArgs(
                key="k1",
                value="v1",
            )])
        foo_auto_snapshot_policy_attachment = volcengine.ebs.AutoSnapshotPolicyAttachment("fooAutoSnapshotPolicyAttachment",
            auto_snapshot_policy_id=foo_auto_snapshot_policy.id,
            volume_id=foo_volume.id)
        ```

        ## Import

        EbsAutoSnapshotPolicyAttachment can be imported using the auto_snapshot_policy_id:volume_id, e.g.

        ```sh
        $ pulumi import volcengine:ebs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_snapshot_policy_id: The id of the auto snapshot policy.
        :param pulumi.Input[str] volume_id: The id of the volume.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AutoSnapshotPolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage ebs auto snapshot policy attachment
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.zones()
        foo_volume = volcengine.ebs.Volume("fooVolume",
            volume_name="acc-test-volume",
            volume_type="ESSD_PL0",
            description="acc-test",
            kind="data",
            size=500,
            zone_id=foo_zones.zones[0].id,
            volume_charge_type="PostPaid",
            project_name="default")
        foo_auto_snapshot_policy = volcengine.ebs.AutoSnapshotPolicy("fooAutoSnapshotPolicy",
            auto_snapshot_policy_name="acc-test-auto-snapshot-policy",
            time_points=[
                "1",
                "5",
                "9",
            ],
            retention_days=-1,
            repeat_weekdays=[
                "2",
                "6",
            ],
            project_name="default",
            tags=[volcengine.ebs.AutoSnapshotPolicyTagArgs(
                key="k1",
                value="v1",
            )])
        foo_auto_snapshot_policy_attachment = volcengine.ebs.AutoSnapshotPolicyAttachment("fooAutoSnapshotPolicyAttachment",
            auto_snapshot_policy_id=foo_auto_snapshot_policy.id,
            volume_id=foo_volume.id)
        ```

        ## Import

        EbsAutoSnapshotPolicyAttachment can be imported using the auto_snapshot_policy_id:volume_id, e.g.

        ```sh
        $ pulumi import volcengine:ebs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param AutoSnapshotPolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoSnapshotPolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_snapshot_policy_id: Optional[pulumi.Input[str]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutoSnapshotPolicyAttachmentArgs.__new__(AutoSnapshotPolicyAttachmentArgs)

            if auto_snapshot_policy_id is None and not opts.urn:
                raise TypeError("Missing required property 'auto_snapshot_policy_id'")
            __props__.__dict__["auto_snapshot_policy_id"] = auto_snapshot_policy_id
            if volume_id is None and not opts.urn:
                raise TypeError("Missing required property 'volume_id'")
            __props__.__dict__["volume_id"] = volume_id
        super(AutoSnapshotPolicyAttachment, __self__).__init__(
            'volcengine:ebs/autoSnapshotPolicyAttachment:AutoSnapshotPolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_snapshot_policy_id: Optional[pulumi.Input[str]] = None,
            volume_id: Optional[pulumi.Input[str]] = None) -> 'AutoSnapshotPolicyAttachment':
        """
        Get an existing AutoSnapshotPolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_snapshot_policy_id: The id of the auto snapshot policy.
        :param pulumi.Input[str] volume_id: The id of the volume.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AutoSnapshotPolicyAttachmentState.__new__(_AutoSnapshotPolicyAttachmentState)

        __props__.__dict__["auto_snapshot_policy_id"] = auto_snapshot_policy_id
        __props__.__dict__["volume_id"] = volume_id
        return AutoSnapshotPolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoSnapshotPolicyId")
    def auto_snapshot_policy_id(self) -> pulumi.Output[str]:
        """
        The id of the auto snapshot policy.
        """
        return pulumi.get(self, "auto_snapshot_policy_id")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[str]:
        """
        The id of the volume.
        """
        return pulumi.get(self, "volume_id")
