# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HostGroupArgs', 'HostGroup']

@pulumi.input_type
class HostGroupArgs:
    def __init__(__self__, *,
                 host_group_name: pulumi.Input[str],
                 host_group_type: pulumi.Input[str],
                 auto_update: Optional[pulumi.Input[bool]] = None,
                 host_identifier: Optional[pulumi.Input[str]] = None,
                 host_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iam_project_name: Optional[pulumi.Input[str]] = None,
                 service_logging: Optional[pulumi.Input[bool]] = None,
                 update_end_time: Optional[pulumi.Input[str]] = None,
                 update_start_time: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HostGroup resource.
        :param pulumi.Input[str] host_group_name: The name of host group.
        :param pulumi.Input[str] host_group_type: The type of host group. The value can be IP or Label.
        :param pulumi.Input[bool] auto_update: Whether enable auto update.
        :param pulumi.Input[str] host_identifier: The identifier of host.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] host_ip_lists: The ip list of host group.
        :param pulumi.Input[str] iam_project_name: The project name of iam.
        :param pulumi.Input[bool] service_logging: Whether enable service logging.
        :param pulumi.Input[str] update_end_time: The update end time of log collector.
        :param pulumi.Input[str] update_start_time: The update start time of log collector.
        """
        pulumi.set(__self__, "host_group_name", host_group_name)
        pulumi.set(__self__, "host_group_type", host_group_type)
        if auto_update is not None:
            pulumi.set(__self__, "auto_update", auto_update)
        if host_identifier is not None:
            pulumi.set(__self__, "host_identifier", host_identifier)
        if host_ip_lists is not None:
            pulumi.set(__self__, "host_ip_lists", host_ip_lists)
        if iam_project_name is not None:
            pulumi.set(__self__, "iam_project_name", iam_project_name)
        if service_logging is not None:
            pulumi.set(__self__, "service_logging", service_logging)
        if update_end_time is not None:
            pulumi.set(__self__, "update_end_time", update_end_time)
        if update_start_time is not None:
            pulumi.set(__self__, "update_start_time", update_start_time)

    @property
    @pulumi.getter(name="hostGroupName")
    def host_group_name(self) -> pulumi.Input[str]:
        """
        The name of host group.
        """
        return pulumi.get(self, "host_group_name")

    @host_group_name.setter
    def host_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "host_group_name", value)

    @property
    @pulumi.getter(name="hostGroupType")
    def host_group_type(self) -> pulumi.Input[str]:
        """
        The type of host group. The value can be IP or Label.
        """
        return pulumi.get(self, "host_group_type")

    @host_group_type.setter
    def host_group_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "host_group_type", value)

    @property
    @pulumi.getter(name="autoUpdate")
    def auto_update(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether enable auto update.
        """
        return pulumi.get(self, "auto_update")

    @auto_update.setter
    def auto_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_update", value)

    @property
    @pulumi.getter(name="hostIdentifier")
    def host_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of host.
        """
        return pulumi.get(self, "host_identifier")

    @host_identifier.setter
    def host_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_identifier", value)

    @property
    @pulumi.getter(name="hostIpLists")
    def host_ip_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ip list of host group.
        """
        return pulumi.get(self, "host_ip_lists")

    @host_ip_lists.setter
    def host_ip_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "host_ip_lists", value)

    @property
    @pulumi.getter(name="iamProjectName")
    def iam_project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of iam.
        """
        return pulumi.get(self, "iam_project_name")

    @iam_project_name.setter
    def iam_project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_project_name", value)

    @property
    @pulumi.getter(name="serviceLogging")
    def service_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether enable service logging.
        """
        return pulumi.get(self, "service_logging")

    @service_logging.setter
    def service_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "service_logging", value)

    @property
    @pulumi.getter(name="updateEndTime")
    def update_end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update end time of log collector.
        """
        return pulumi.get(self, "update_end_time")

    @update_end_time.setter
    def update_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_end_time", value)

    @property
    @pulumi.getter(name="updateStartTime")
    def update_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update start time of log collector.
        """
        return pulumi.get(self, "update_start_time")

    @update_start_time.setter
    def update_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_start_time", value)


@pulumi.input_type
class _HostGroupState:
    def __init__(__self__, *,
                 abnormal_heartbeat_status_count: Optional[pulumi.Input[int]] = None,
                 agent_latest_version: Optional[pulumi.Input[str]] = None,
                 auto_update: Optional[pulumi.Input[bool]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 host_count: Optional[pulumi.Input[int]] = None,
                 host_group_name: Optional[pulumi.Input[str]] = None,
                 host_group_type: Optional[pulumi.Input[str]] = None,
                 host_identifier: Optional[pulumi.Input[str]] = None,
                 host_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iam_project_name: Optional[pulumi.Input[str]] = None,
                 modify_time: Optional[pulumi.Input[str]] = None,
                 normal_heartbeat_status_count: Optional[pulumi.Input[int]] = None,
                 rule_count: Optional[pulumi.Input[int]] = None,
                 service_logging: Optional[pulumi.Input[bool]] = None,
                 update_end_time: Optional[pulumi.Input[str]] = None,
                 update_start_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HostGroup resources.
        :param pulumi.Input[int] abnormal_heartbeat_status_count: The abnormal heartbeat status count of host.
        :param pulumi.Input[str] agent_latest_version: The latest version of log collector.
        :param pulumi.Input[bool] auto_update: Whether enable auto update.
        :param pulumi.Input[str] create_time: The create time of host group.
        :param pulumi.Input[int] host_count: The count of host.
        :param pulumi.Input[str] host_group_name: The name of host group.
        :param pulumi.Input[str] host_group_type: The type of host group. The value can be IP or Label.
        :param pulumi.Input[str] host_identifier: The identifier of host.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] host_ip_lists: The ip list of host group.
        :param pulumi.Input[str] iam_project_name: The project name of iam.
        :param pulumi.Input[str] modify_time: The modify time of host group.
        :param pulumi.Input[int] normal_heartbeat_status_count: The normal heartbeat status count of host.
        :param pulumi.Input[int] rule_count: The rule count of host.
        :param pulumi.Input[bool] service_logging: Whether enable service logging.
        :param pulumi.Input[str] update_end_time: The update end time of log collector.
        :param pulumi.Input[str] update_start_time: The update start time of log collector.
        """
        if abnormal_heartbeat_status_count is not None:
            pulumi.set(__self__, "abnormal_heartbeat_status_count", abnormal_heartbeat_status_count)
        if agent_latest_version is not None:
            pulumi.set(__self__, "agent_latest_version", agent_latest_version)
        if auto_update is not None:
            pulumi.set(__self__, "auto_update", auto_update)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if host_count is not None:
            pulumi.set(__self__, "host_count", host_count)
        if host_group_name is not None:
            pulumi.set(__self__, "host_group_name", host_group_name)
        if host_group_type is not None:
            pulumi.set(__self__, "host_group_type", host_group_type)
        if host_identifier is not None:
            pulumi.set(__self__, "host_identifier", host_identifier)
        if host_ip_lists is not None:
            pulumi.set(__self__, "host_ip_lists", host_ip_lists)
        if iam_project_name is not None:
            pulumi.set(__self__, "iam_project_name", iam_project_name)
        if modify_time is not None:
            pulumi.set(__self__, "modify_time", modify_time)
        if normal_heartbeat_status_count is not None:
            pulumi.set(__self__, "normal_heartbeat_status_count", normal_heartbeat_status_count)
        if rule_count is not None:
            pulumi.set(__self__, "rule_count", rule_count)
        if service_logging is not None:
            pulumi.set(__self__, "service_logging", service_logging)
        if update_end_time is not None:
            pulumi.set(__self__, "update_end_time", update_end_time)
        if update_start_time is not None:
            pulumi.set(__self__, "update_start_time", update_start_time)

    @property
    @pulumi.getter(name="abnormalHeartbeatStatusCount")
    def abnormal_heartbeat_status_count(self) -> Optional[pulumi.Input[int]]:
        """
        The abnormal heartbeat status count of host.
        """
        return pulumi.get(self, "abnormal_heartbeat_status_count")

    @abnormal_heartbeat_status_count.setter
    def abnormal_heartbeat_status_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "abnormal_heartbeat_status_count", value)

    @property
    @pulumi.getter(name="agentLatestVersion")
    def agent_latest_version(self) -> Optional[pulumi.Input[str]]:
        """
        The latest version of log collector.
        """
        return pulumi.get(self, "agent_latest_version")

    @agent_latest_version.setter
    def agent_latest_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "agent_latest_version", value)

    @property
    @pulumi.getter(name="autoUpdate")
    def auto_update(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether enable auto update.
        """
        return pulumi.get(self, "auto_update")

    @auto_update.setter
    def auto_update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_update", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of host group.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="hostCount")
    def host_count(self) -> Optional[pulumi.Input[int]]:
        """
        The count of host.
        """
        return pulumi.get(self, "host_count")

    @host_count.setter
    def host_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "host_count", value)

    @property
    @pulumi.getter(name="hostGroupName")
    def host_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of host group.
        """
        return pulumi.get(self, "host_group_name")

    @host_group_name.setter
    def host_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_group_name", value)

    @property
    @pulumi.getter(name="hostGroupType")
    def host_group_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of host group. The value can be IP or Label.
        """
        return pulumi.get(self, "host_group_type")

    @host_group_type.setter
    def host_group_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_group_type", value)

    @property
    @pulumi.getter(name="hostIdentifier")
    def host_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of host.
        """
        return pulumi.get(self, "host_identifier")

    @host_identifier.setter
    def host_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_identifier", value)

    @property
    @pulumi.getter(name="hostIpLists")
    def host_ip_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The ip list of host group.
        """
        return pulumi.get(self, "host_ip_lists")

    @host_ip_lists.setter
    def host_ip_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "host_ip_lists", value)

    @property
    @pulumi.getter(name="iamProjectName")
    def iam_project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of iam.
        """
        return pulumi.get(self, "iam_project_name")

    @iam_project_name.setter
    def iam_project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_project_name", value)

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> Optional[pulumi.Input[str]]:
        """
        The modify time of host group.
        """
        return pulumi.get(self, "modify_time")

    @modify_time.setter
    def modify_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_time", value)

    @property
    @pulumi.getter(name="normalHeartbeatStatusCount")
    def normal_heartbeat_status_count(self) -> Optional[pulumi.Input[int]]:
        """
        The normal heartbeat status count of host.
        """
        return pulumi.get(self, "normal_heartbeat_status_count")

    @normal_heartbeat_status_count.setter
    def normal_heartbeat_status_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "normal_heartbeat_status_count", value)

    @property
    @pulumi.getter(name="ruleCount")
    def rule_count(self) -> Optional[pulumi.Input[int]]:
        """
        The rule count of host.
        """
        return pulumi.get(self, "rule_count")

    @rule_count.setter
    def rule_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rule_count", value)

    @property
    @pulumi.getter(name="serviceLogging")
    def service_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether enable service logging.
        """
        return pulumi.get(self, "service_logging")

    @service_logging.setter
    def service_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "service_logging", value)

    @property
    @pulumi.getter(name="updateEndTime")
    def update_end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update end time of log collector.
        """
        return pulumi.get(self, "update_end_time")

    @update_end_time.setter
    def update_end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_end_time", value)

    @property
    @pulumi.getter(name="updateStartTime")
    def update_start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update start time of log collector.
        """
        return pulumi.get(self, "update_start_time")

    @update_start_time.setter
    def update_start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_start_time", value)


class HostGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_update: Optional[pulumi.Input[bool]] = None,
                 host_group_name: Optional[pulumi.Input[str]] = None,
                 host_group_type: Optional[pulumi.Input[str]] = None,
                 host_identifier: Optional[pulumi.Input[str]] = None,
                 host_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iam_project_name: Optional[pulumi.Input[str]] = None,
                 service_logging: Optional[pulumi.Input[bool]] = None,
                 update_end_time: Optional[pulumi.Input[str]] = None,
                 update_start_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage tls host group
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.tls.HostGroup("foo",
            auto_update=False,
            host_group_name="tfgroup",
            host_group_type="Label",
            host_identifier="tf-controller",
            service_logging=False)
        ```

        ## Import

        Tls Host Group can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:tls/hostGroup:HostGroup default edf052s21s*******dc15
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_update: Whether enable auto update.
        :param pulumi.Input[str] host_group_name: The name of host group.
        :param pulumi.Input[str] host_group_type: The type of host group. The value can be IP or Label.
        :param pulumi.Input[str] host_identifier: The identifier of host.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] host_ip_lists: The ip list of host group.
        :param pulumi.Input[str] iam_project_name: The project name of iam.
        :param pulumi.Input[bool] service_logging: Whether enable service logging.
        :param pulumi.Input[str] update_end_time: The update end time of log collector.
        :param pulumi.Input[str] update_start_time: The update start time of log collector.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HostGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage tls host group
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.tls.HostGroup("foo",
            auto_update=False,
            host_group_name="tfgroup",
            host_group_type="Label",
            host_identifier="tf-controller",
            service_logging=False)
        ```

        ## Import

        Tls Host Group can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:tls/hostGroup:HostGroup default edf052s21s*******dc15
        ```

        :param str resource_name: The name of the resource.
        :param HostGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HostGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_update: Optional[pulumi.Input[bool]] = None,
                 host_group_name: Optional[pulumi.Input[str]] = None,
                 host_group_type: Optional[pulumi.Input[str]] = None,
                 host_identifier: Optional[pulumi.Input[str]] = None,
                 host_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 iam_project_name: Optional[pulumi.Input[str]] = None,
                 service_logging: Optional[pulumi.Input[bool]] = None,
                 update_end_time: Optional[pulumi.Input[str]] = None,
                 update_start_time: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HostGroupArgs.__new__(HostGroupArgs)

            __props__.__dict__["auto_update"] = auto_update
            if host_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'host_group_name'")
            __props__.__dict__["host_group_name"] = host_group_name
            if host_group_type is None and not opts.urn:
                raise TypeError("Missing required property 'host_group_type'")
            __props__.__dict__["host_group_type"] = host_group_type
            __props__.__dict__["host_identifier"] = host_identifier
            __props__.__dict__["host_ip_lists"] = host_ip_lists
            __props__.__dict__["iam_project_name"] = iam_project_name
            __props__.__dict__["service_logging"] = service_logging
            __props__.__dict__["update_end_time"] = update_end_time
            __props__.__dict__["update_start_time"] = update_start_time
            __props__.__dict__["abnormal_heartbeat_status_count"] = None
            __props__.__dict__["agent_latest_version"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["host_count"] = None
            __props__.__dict__["modify_time"] = None
            __props__.__dict__["normal_heartbeat_status_count"] = None
            __props__.__dict__["rule_count"] = None
        super(HostGroup, __self__).__init__(
            'volcengine:tls/hostGroup:HostGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            abnormal_heartbeat_status_count: Optional[pulumi.Input[int]] = None,
            agent_latest_version: Optional[pulumi.Input[str]] = None,
            auto_update: Optional[pulumi.Input[bool]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            host_count: Optional[pulumi.Input[int]] = None,
            host_group_name: Optional[pulumi.Input[str]] = None,
            host_group_type: Optional[pulumi.Input[str]] = None,
            host_identifier: Optional[pulumi.Input[str]] = None,
            host_ip_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            iam_project_name: Optional[pulumi.Input[str]] = None,
            modify_time: Optional[pulumi.Input[str]] = None,
            normal_heartbeat_status_count: Optional[pulumi.Input[int]] = None,
            rule_count: Optional[pulumi.Input[int]] = None,
            service_logging: Optional[pulumi.Input[bool]] = None,
            update_end_time: Optional[pulumi.Input[str]] = None,
            update_start_time: Optional[pulumi.Input[str]] = None) -> 'HostGroup':
        """
        Get an existing HostGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] abnormal_heartbeat_status_count: The abnormal heartbeat status count of host.
        :param pulumi.Input[str] agent_latest_version: The latest version of log collector.
        :param pulumi.Input[bool] auto_update: Whether enable auto update.
        :param pulumi.Input[str] create_time: The create time of host group.
        :param pulumi.Input[int] host_count: The count of host.
        :param pulumi.Input[str] host_group_name: The name of host group.
        :param pulumi.Input[str] host_group_type: The type of host group. The value can be IP or Label.
        :param pulumi.Input[str] host_identifier: The identifier of host.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] host_ip_lists: The ip list of host group.
        :param pulumi.Input[str] iam_project_name: The project name of iam.
        :param pulumi.Input[str] modify_time: The modify time of host group.
        :param pulumi.Input[int] normal_heartbeat_status_count: The normal heartbeat status count of host.
        :param pulumi.Input[int] rule_count: The rule count of host.
        :param pulumi.Input[bool] service_logging: Whether enable service logging.
        :param pulumi.Input[str] update_end_time: The update end time of log collector.
        :param pulumi.Input[str] update_start_time: The update start time of log collector.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HostGroupState.__new__(_HostGroupState)

        __props__.__dict__["abnormal_heartbeat_status_count"] = abnormal_heartbeat_status_count
        __props__.__dict__["agent_latest_version"] = agent_latest_version
        __props__.__dict__["auto_update"] = auto_update
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["host_count"] = host_count
        __props__.__dict__["host_group_name"] = host_group_name
        __props__.__dict__["host_group_type"] = host_group_type
        __props__.__dict__["host_identifier"] = host_identifier
        __props__.__dict__["host_ip_lists"] = host_ip_lists
        __props__.__dict__["iam_project_name"] = iam_project_name
        __props__.__dict__["modify_time"] = modify_time
        __props__.__dict__["normal_heartbeat_status_count"] = normal_heartbeat_status_count
        __props__.__dict__["rule_count"] = rule_count
        __props__.__dict__["service_logging"] = service_logging
        __props__.__dict__["update_end_time"] = update_end_time
        __props__.__dict__["update_start_time"] = update_start_time
        return HostGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="abnormalHeartbeatStatusCount")
    def abnormal_heartbeat_status_count(self) -> pulumi.Output[int]:
        """
        The abnormal heartbeat status count of host.
        """
        return pulumi.get(self, "abnormal_heartbeat_status_count")

    @property
    @pulumi.getter(name="agentLatestVersion")
    def agent_latest_version(self) -> pulumi.Output[str]:
        """
        The latest version of log collector.
        """
        return pulumi.get(self, "agent_latest_version")

    @property
    @pulumi.getter(name="autoUpdate")
    def auto_update(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether enable auto update.
        """
        return pulumi.get(self, "auto_update")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The create time of host group.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="hostCount")
    def host_count(self) -> pulumi.Output[int]:
        """
        The count of host.
        """
        return pulumi.get(self, "host_count")

    @property
    @pulumi.getter(name="hostGroupName")
    def host_group_name(self) -> pulumi.Output[str]:
        """
        The name of host group.
        """
        return pulumi.get(self, "host_group_name")

    @property
    @pulumi.getter(name="hostGroupType")
    def host_group_type(self) -> pulumi.Output[str]:
        """
        The type of host group. The value can be IP or Label.
        """
        return pulumi.get(self, "host_group_type")

    @property
    @pulumi.getter(name="hostIdentifier")
    def host_identifier(self) -> pulumi.Output[Optional[str]]:
        """
        The identifier of host.
        """
        return pulumi.get(self, "host_identifier")

    @property
    @pulumi.getter(name="hostIpLists")
    def host_ip_lists(self) -> pulumi.Output[Sequence[str]]:
        """
        The ip list of host group.
        """
        return pulumi.get(self, "host_ip_lists")

    @property
    @pulumi.getter(name="iamProjectName")
    def iam_project_name(self) -> pulumi.Output[Optional[str]]:
        """
        The project name of iam.
        """
        return pulumi.get(self, "iam_project_name")

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> pulumi.Output[str]:
        """
        The modify time of host group.
        """
        return pulumi.get(self, "modify_time")

    @property
    @pulumi.getter(name="normalHeartbeatStatusCount")
    def normal_heartbeat_status_count(self) -> pulumi.Output[int]:
        """
        The normal heartbeat status count of host.
        """
        return pulumi.get(self, "normal_heartbeat_status_count")

    @property
    @pulumi.getter(name="ruleCount")
    def rule_count(self) -> pulumi.Output[int]:
        """
        The rule count of host.
        """
        return pulumi.get(self, "rule_count")

    @property
    @pulumi.getter(name="serviceLogging")
    def service_logging(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether enable service logging.
        """
        return pulumi.get(self, "service_logging")

    @property
    @pulumi.getter(name="updateEndTime")
    def update_end_time(self) -> pulumi.Output[str]:
        """
        The update end time of log collector.
        """
        return pulumi.get(self, "update_end_time")

    @property
    @pulumi.getter(name="updateStartTime")
    def update_start_time(self) -> pulumi.Output[str]:
        """
        The update start time of log collector.
        """
        return pulumi.get(self, "update_start_time")
