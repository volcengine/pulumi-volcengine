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
from ._inputs import *

__all__ = ['AutoSnapshotPolicyArgs', 'AutoSnapshotPolicy']

@pulumi.input_type
class AutoSnapshotPolicyArgs:
    def __init__(__self__, *,
                 auto_snapshot_policy_name: pulumi.Input[str],
                 retention_days: pulumi.Input[int],
                 time_points: pulumi.Input[Sequence[pulumi.Input[str]]],
                 project_name: Optional[pulumi.Input[str]] = None,
                 repeat_days: Optional[pulumi.Input[int]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AutoSnapshotPolicyTagArgs']]]] = None):
        """
        The set of arguments for constructing a AutoSnapshotPolicy resource.
        :param pulumi.Input[str] auto_snapshot_policy_name: The name of the auto snapshot policy.
        :param pulumi.Input[int] retention_days: The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        :param pulumi.Input[str] project_name: The project name of the auto snapshot policy.
        :param pulumi.Input[int] repeat_days: Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        :param pulumi.Input[Sequence[pulumi.Input['AutoSnapshotPolicyTagArgs']]] tags: Tags.
        """
        pulumi.set(__self__, "auto_snapshot_policy_name", auto_snapshot_policy_name)
        pulumi.set(__self__, "retention_days", retention_days)
        pulumi.set(__self__, "time_points", time_points)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if repeat_days is not None:
            pulumi.set(__self__, "repeat_days", repeat_days)
        if repeat_weekdays is not None:
            pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="autoSnapshotPolicyName")
    def auto_snapshot_policy_name(self) -> pulumi.Input[str]:
        """
        The name of the auto snapshot policy.
        """
        return pulumi.get(self, "auto_snapshot_policy_name")

    @auto_snapshot_policy_name.setter
    def auto_snapshot_policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_snapshot_policy_name", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Input[int]:
        """
        The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: pulumi.Input[int]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        """
        return pulumi.get(self, "time_points")

    @time_points.setter
    def time_points(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "time_points", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of the auto snapshot policy.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="repeatDays")
    def repeat_days(self) -> Optional[pulumi.Input[int]]:
        """
        Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        """
        return pulumi.get(self, "repeat_days")

    @repeat_days.setter
    def repeat_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repeat_days", value)

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        """
        return pulumi.get(self, "repeat_weekdays")

    @repeat_weekdays.setter
    def repeat_weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "repeat_weekdays", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutoSnapshotPolicyTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutoSnapshotPolicyTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _AutoSnapshotPolicyState:
    def __init__(__self__, *,
                 auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 repeat_days: Optional[pulumi.Input[int]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AutoSnapshotPolicyTagArgs']]]] = None,
                 time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None,
                 volume_nums: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering AutoSnapshotPolicy resources.
        :param pulumi.Input[str] auto_snapshot_policy_name: The name of the auto snapshot policy.
        :param pulumi.Input[str] created_at: The creation time of the auto snapshot policy.
        :param pulumi.Input[str] project_name: The project name of the auto snapshot policy.
        :param pulumi.Input[int] repeat_days: Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        :param pulumi.Input[int] retention_days: The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        :param pulumi.Input[str] status: The status of the auto snapshot policy.
        :param pulumi.Input[Sequence[pulumi.Input['AutoSnapshotPolicyTagArgs']]] tags: Tags.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        :param pulumi.Input[str] updated_at: The updated time of the auto snapshot policy.
        :param pulumi.Input[int] volume_nums: The number of volumes associated with the auto snapshot policy.
        """
        if auto_snapshot_policy_name is not None:
            pulumi.set(__self__, "auto_snapshot_policy_name", auto_snapshot_policy_name)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if repeat_days is not None:
            pulumi.set(__self__, "repeat_days", repeat_days)
        if repeat_weekdays is not None:
            pulumi.set(__self__, "repeat_weekdays", repeat_weekdays)
        if retention_days is not None:
            pulumi.set(__self__, "retention_days", retention_days)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if time_points is not None:
            pulumi.set(__self__, "time_points", time_points)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)
        if volume_nums is not None:
            pulumi.set(__self__, "volume_nums", volume_nums)

    @property
    @pulumi.getter(name="autoSnapshotPolicyName")
    def auto_snapshot_policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the auto snapshot policy.
        """
        return pulumi.get(self, "auto_snapshot_policy_name")

    @auto_snapshot_policy_name.setter
    def auto_snapshot_policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_snapshot_policy_name", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the auto snapshot policy.
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of the auto snapshot policy.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="repeatDays")
    def repeat_days(self) -> Optional[pulumi.Input[int]]:
        """
        Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        """
        return pulumi.get(self, "repeat_days")

    @repeat_days.setter
    def repeat_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "repeat_days", value)

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        """
        return pulumi.get(self, "repeat_weekdays")

    @repeat_weekdays.setter
    def repeat_weekdays(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "repeat_weekdays", value)

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> Optional[pulumi.Input[int]]:
        """
        The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        """
        return pulumi.get(self, "retention_days")

    @retention_days.setter
    def retention_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_days", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the auto snapshot policy.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AutoSnapshotPolicyTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AutoSnapshotPolicyTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        """
        return pulumi.get(self, "time_points")

    @time_points.setter
    def time_points(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "time_points", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        The updated time of the auto snapshot policy.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)

    @property
    @pulumi.getter(name="volumeNums")
    def volume_nums(self) -> Optional[pulumi.Input[int]]:
        """
        The number of volumes associated with the auto snapshot policy.
        """
        return pulumi.get(self, "volume_nums")

    @volume_nums.setter
    def volume_nums(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_nums", value)


class AutoSnapshotPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 repeat_days: Optional[pulumi.Input[int]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoSnapshotPolicyTagArgs']]]]] = None,
                 time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage ebs auto snapshot policy
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.ebs.AutoSnapshotPolicy("foo",
            auto_snapshot_policy_name="acc-test-auto-snapshot-policy",
            project_name="default",
            repeat_weekdays=[
                "2",
                "6",
            ],
            retention_days=-1,
            tags=[volcengine.ebs.AutoSnapshotPolicyTagArgs(
                key="k1",
                value="v1",
            )],
            time_points=[
                "1",
                "5",
                "9",
            ])
        ```

        ## Import

        EbsAutoSnapshotPolicy can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_snapshot_policy_name: The name of the auto snapshot policy.
        :param pulumi.Input[str] project_name: The project name of the auto snapshot policy.
        :param pulumi.Input[int] repeat_days: Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        :param pulumi.Input[int] retention_days: The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoSnapshotPolicyTagArgs']]]] tags: Tags.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AutoSnapshotPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage ebs auto snapshot policy
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.ebs.AutoSnapshotPolicy("foo",
            auto_snapshot_policy_name="acc-test-auto-snapshot-policy",
            project_name="default",
            repeat_weekdays=[
                "2",
                "6",
            ],
            retention_days=-1,
            tags=[volcengine.ebs.AutoSnapshotPolicyTagArgs(
                key="k1",
                value="v1",
            )],
            time_points=[
                "1",
                "5",
                "9",
            ])
        ```

        ## Import

        EbsAutoSnapshotPolicy can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param AutoSnapshotPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AutoSnapshotPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 repeat_days: Optional[pulumi.Input[int]] = None,
                 repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoSnapshotPolicyTagArgs']]]]] = None,
                 time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AutoSnapshotPolicyArgs.__new__(AutoSnapshotPolicyArgs)

            if auto_snapshot_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'auto_snapshot_policy_name'")
            __props__.__dict__["auto_snapshot_policy_name"] = auto_snapshot_policy_name
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["repeat_days"] = repeat_days
            __props__.__dict__["repeat_weekdays"] = repeat_weekdays
            if retention_days is None and not opts.urn:
                raise TypeError("Missing required property 'retention_days'")
            __props__.__dict__["retention_days"] = retention_days
            __props__.__dict__["tags"] = tags
            if time_points is None and not opts.urn:
                raise TypeError("Missing required property 'time_points'")
            __props__.__dict__["time_points"] = time_points
            __props__.__dict__["created_at"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["updated_at"] = None
            __props__.__dict__["volume_nums"] = None
        super(AutoSnapshotPolicy, __self__).__init__(
            'volcengine:ebs/autoSnapshotPolicy:AutoSnapshotPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_snapshot_policy_name: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            repeat_days: Optional[pulumi.Input[int]] = None,
            repeat_weekdays: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            retention_days: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoSnapshotPolicyTagArgs']]]]] = None,
            time_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            updated_at: Optional[pulumi.Input[str]] = None,
            volume_nums: Optional[pulumi.Input[int]] = None) -> 'AutoSnapshotPolicy':
        """
        Get an existing AutoSnapshotPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_snapshot_policy_name: The name of the auto snapshot policy.
        :param pulumi.Input[str] created_at: The creation time of the auto snapshot policy.
        :param pulumi.Input[str] project_name: The project name of the auto snapshot policy.
        :param pulumi.Input[int] repeat_days: Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] repeat_weekdays: The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        :param pulumi.Input[int] retention_days: The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        :param pulumi.Input[str] status: The status of the auto snapshot policy.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AutoSnapshotPolicyTagArgs']]]] tags: Tags.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] time_points: The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        :param pulumi.Input[str] updated_at: The updated time of the auto snapshot policy.
        :param pulumi.Input[int] volume_nums: The number of volumes associated with the auto snapshot policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AutoSnapshotPolicyState.__new__(_AutoSnapshotPolicyState)

        __props__.__dict__["auto_snapshot_policy_name"] = auto_snapshot_policy_name
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["repeat_days"] = repeat_days
        __props__.__dict__["repeat_weekdays"] = repeat_weekdays
        __props__.__dict__["retention_days"] = retention_days
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["time_points"] = time_points
        __props__.__dict__["updated_at"] = updated_at
        __props__.__dict__["volume_nums"] = volume_nums
        return AutoSnapshotPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoSnapshotPolicyName")
    def auto_snapshot_policy_name(self) -> pulumi.Output[str]:
        """
        The name of the auto snapshot policy.
        """
        return pulumi.get(self, "auto_snapshot_policy_name")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The creation time of the auto snapshot policy.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The project name of the auto snapshot policy.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="repeatDays")
    def repeat_days(self) -> pulumi.Output[Optional[int]]:
        """
        Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. Only one of `repeat_weekdays, repeat_days` can be specified.
        """
        return pulumi.get(self, "repeat_days")

    @property
    @pulumi.getter(name="repeatWeekdays")
    def repeat_weekdays(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. Only one of `repeat_weekdays, repeat_days` can be specified.
        """
        return pulumi.get(self, "repeat_weekdays")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[int]:
        """
        The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the auto snapshot policy.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.AutoSnapshotPolicyTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timePoints")
    def time_points(self) -> pulumi.Output[Sequence[str]]:
        """
        The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.
        """
        return pulumi.get(self, "time_points")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        The updated time of the auto snapshot policy.
        """
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="volumeNums")
    def volume_nums(self) -> pulumi.Output[int]:
        """
        The number of volumes associated with the auto snapshot policy.
        """
        return pulumi.get(self, "volume_nums")
