# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AlarmArgs', 'Alarm']

@pulumi.input_type
class AlarmArgs:
    def __init__(__self__, *,
                 alarm_name: pulumi.Input[str],
                 alarm_notify_groups: pulumi.Input[Sequence[pulumi.Input[str]]],
                 condition: pulumi.Input[str],
                 project_id: pulumi.Input[str],
                 query_requests: pulumi.Input[Sequence[pulumi.Input['AlarmQueryRequestArgs']]],
                 request_cycle: pulumi.Input['AlarmRequestCycleArgs'],
                 alarm_period: Optional[pulumi.Input[int]] = None,
                 alarm_period_detail: Optional[pulumi.Input['AlarmAlarmPeriodDetailArgs']] = None,
                 status: Optional[pulumi.Input[bool]] = None,
                 trigger_period: Optional[pulumi.Input[int]] = None,
                 user_define_msg: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Alarm resource.
        :param pulumi.Input[str] alarm_name: The name of the alarm.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_notify_groups: List of notification groups corresponding to the alarm.
        :param pulumi.Input[str] condition: Alarm trigger condition.
        :param pulumi.Input[str] project_id: The project id.
        :param pulumi.Input[Sequence[pulumi.Input['AlarmQueryRequestArgs']]] query_requests: Search and analyze sentences, 1~3 can be configured.
        :param pulumi.Input['AlarmRequestCycleArgs'] request_cycle: The execution period of the alarm task.
        :param pulumi.Input[int] alarm_period: Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        :param pulumi.Input['AlarmAlarmPeriodDetailArgs'] alarm_period_detail: Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        :param pulumi.Input[bool] status: Whether to enable the alert policy. The default value is true, that is, on.
        :param pulumi.Input[int] trigger_period: Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
        :param pulumi.Input[str] user_define_msg: Customize the alarm notification content.
        """
        pulumi.set(__self__, "alarm_name", alarm_name)
        pulumi.set(__self__, "alarm_notify_groups", alarm_notify_groups)
        pulumi.set(__self__, "condition", condition)
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "query_requests", query_requests)
        pulumi.set(__self__, "request_cycle", request_cycle)
        if alarm_period is not None:
            pulumi.set(__self__, "alarm_period", alarm_period)
        if alarm_period_detail is not None:
            pulumi.set(__self__, "alarm_period_detail", alarm_period_detail)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if trigger_period is not None:
            pulumi.set(__self__, "trigger_period", trigger_period)
        if user_define_msg is not None:
            pulumi.set(__self__, "user_define_msg", user_define_msg)

    @property
    @pulumi.getter(name="alarmName")
    def alarm_name(self) -> pulumi.Input[str]:
        """
        The name of the alarm.
        """
        return pulumi.get(self, "alarm_name")

    @alarm_name.setter
    def alarm_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "alarm_name", value)

    @property
    @pulumi.getter(name="alarmNotifyGroups")
    def alarm_notify_groups(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of notification groups corresponding to the alarm.
        """
        return pulumi.get(self, "alarm_notify_groups")

    @alarm_notify_groups.setter
    def alarm_notify_groups(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "alarm_notify_groups", value)

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Input[str]:
        """
        Alarm trigger condition.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: pulumi.Input[str]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The project id.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="queryRequests")
    def query_requests(self) -> pulumi.Input[Sequence[pulumi.Input['AlarmQueryRequestArgs']]]:
        """
        Search and analyze sentences, 1~3 can be configured.
        """
        return pulumi.get(self, "query_requests")

    @query_requests.setter
    def query_requests(self, value: pulumi.Input[Sequence[pulumi.Input['AlarmQueryRequestArgs']]]):
        pulumi.set(self, "query_requests", value)

    @property
    @pulumi.getter(name="requestCycle")
    def request_cycle(self) -> pulumi.Input['AlarmRequestCycleArgs']:
        """
        The execution period of the alarm task.
        """
        return pulumi.get(self, "request_cycle")

    @request_cycle.setter
    def request_cycle(self, value: pulumi.Input['AlarmRequestCycleArgs']):
        pulumi.set(self, "request_cycle", value)

    @property
    @pulumi.getter(name="alarmPeriod")
    def alarm_period(self) -> Optional[pulumi.Input[int]]:
        """
        Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        """
        return pulumi.get(self, "alarm_period")

    @alarm_period.setter
    def alarm_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alarm_period", value)

    @property
    @pulumi.getter(name="alarmPeriodDetail")
    def alarm_period_detail(self) -> Optional[pulumi.Input['AlarmAlarmPeriodDetailArgs']]:
        """
        Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        """
        return pulumi.get(self, "alarm_period_detail")

    @alarm_period_detail.setter
    def alarm_period_detail(self, value: Optional[pulumi.Input['AlarmAlarmPeriodDetailArgs']]):
        pulumi.set(self, "alarm_period_detail", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the alert policy. The default value is true, that is, on.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="triggerPeriod")
    def trigger_period(self) -> Optional[pulumi.Input[int]]:
        """
        Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
        """
        return pulumi.get(self, "trigger_period")

    @trigger_period.setter
    def trigger_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trigger_period", value)

    @property
    @pulumi.getter(name="userDefineMsg")
    def user_define_msg(self) -> Optional[pulumi.Input[str]]:
        """
        Customize the alarm notification content.
        """
        return pulumi.get(self, "user_define_msg")

    @user_define_msg.setter
    def user_define_msg(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_define_msg", value)


@pulumi.input_type
class _AlarmState:
    def __init__(__self__, *,
                 alarm_id: Optional[pulumi.Input[str]] = None,
                 alarm_name: Optional[pulumi.Input[str]] = None,
                 alarm_notify_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_period: Optional[pulumi.Input[int]] = None,
                 alarm_period_detail: Optional[pulumi.Input['AlarmAlarmPeriodDetailArgs']] = None,
                 condition: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 query_requests: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmQueryRequestArgs']]]] = None,
                 request_cycle: Optional[pulumi.Input['AlarmRequestCycleArgs']] = None,
                 status: Optional[pulumi.Input[bool]] = None,
                 trigger_period: Optional[pulumi.Input[int]] = None,
                 user_define_msg: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Alarm resources.
        :param pulumi.Input[str] alarm_id: The alarm id.
        :param pulumi.Input[str] alarm_name: The name of the alarm.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_notify_groups: List of notification groups corresponding to the alarm.
        :param pulumi.Input[int] alarm_period: Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        :param pulumi.Input['AlarmAlarmPeriodDetailArgs'] alarm_period_detail: Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        :param pulumi.Input[str] condition: Alarm trigger condition.
        :param pulumi.Input[str] project_id: The project id.
        :param pulumi.Input[Sequence[pulumi.Input['AlarmQueryRequestArgs']]] query_requests: Search and analyze sentences, 1~3 can be configured.
        :param pulumi.Input['AlarmRequestCycleArgs'] request_cycle: The execution period of the alarm task.
        :param pulumi.Input[bool] status: Whether to enable the alert policy. The default value is true, that is, on.
        :param pulumi.Input[int] trigger_period: Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
        :param pulumi.Input[str] user_define_msg: Customize the alarm notification content.
        """
        if alarm_id is not None:
            pulumi.set(__self__, "alarm_id", alarm_id)
        if alarm_name is not None:
            pulumi.set(__self__, "alarm_name", alarm_name)
        if alarm_notify_groups is not None:
            pulumi.set(__self__, "alarm_notify_groups", alarm_notify_groups)
        if alarm_period is not None:
            pulumi.set(__self__, "alarm_period", alarm_period)
        if alarm_period_detail is not None:
            pulumi.set(__self__, "alarm_period_detail", alarm_period_detail)
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if query_requests is not None:
            pulumi.set(__self__, "query_requests", query_requests)
        if request_cycle is not None:
            pulumi.set(__self__, "request_cycle", request_cycle)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if trigger_period is not None:
            pulumi.set(__self__, "trigger_period", trigger_period)
        if user_define_msg is not None:
            pulumi.set(__self__, "user_define_msg", user_define_msg)

    @property
    @pulumi.getter(name="alarmId")
    def alarm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The alarm id.
        """
        return pulumi.get(self, "alarm_id")

    @alarm_id.setter
    def alarm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_id", value)

    @property
    @pulumi.getter(name="alarmName")
    def alarm_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the alarm.
        """
        return pulumi.get(self, "alarm_name")

    @alarm_name.setter
    def alarm_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alarm_name", value)

    @property
    @pulumi.getter(name="alarmNotifyGroups")
    def alarm_notify_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of notification groups corresponding to the alarm.
        """
        return pulumi.get(self, "alarm_notify_groups")

    @alarm_notify_groups.setter
    def alarm_notify_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "alarm_notify_groups", value)

    @property
    @pulumi.getter(name="alarmPeriod")
    def alarm_period(self) -> Optional[pulumi.Input[int]]:
        """
        Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        """
        return pulumi.get(self, "alarm_period")

    @alarm_period.setter
    def alarm_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "alarm_period", value)

    @property
    @pulumi.getter(name="alarmPeriodDetail")
    def alarm_period_detail(self) -> Optional[pulumi.Input['AlarmAlarmPeriodDetailArgs']]:
        """
        Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        """
        return pulumi.get(self, "alarm_period_detail")

    @alarm_period_detail.setter
    def alarm_period_detail(self, value: Optional[pulumi.Input['AlarmAlarmPeriodDetailArgs']]):
        pulumi.set(self, "alarm_period_detail", value)

    @property
    @pulumi.getter
    def condition(self) -> Optional[pulumi.Input[str]]:
        """
        Alarm trigger condition.
        """
        return pulumi.get(self, "condition")

    @condition.setter
    def condition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project id.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="queryRequests")
    def query_requests(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AlarmQueryRequestArgs']]]]:
        """
        Search and analyze sentences, 1~3 can be configured.
        """
        return pulumi.get(self, "query_requests")

    @query_requests.setter
    def query_requests(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AlarmQueryRequestArgs']]]]):
        pulumi.set(self, "query_requests", value)

    @property
    @pulumi.getter(name="requestCycle")
    def request_cycle(self) -> Optional[pulumi.Input['AlarmRequestCycleArgs']]:
        """
        The execution period of the alarm task.
        """
        return pulumi.get(self, "request_cycle")

    @request_cycle.setter
    def request_cycle(self, value: Optional[pulumi.Input['AlarmRequestCycleArgs']]):
        pulumi.set(self, "request_cycle", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable the alert policy. The default value is true, that is, on.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="triggerPeriod")
    def trigger_period(self) -> Optional[pulumi.Input[int]]:
        """
        Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
        """
        return pulumi.get(self, "trigger_period")

    @trigger_period.setter
    def trigger_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "trigger_period", value)

    @property
    @pulumi.getter(name="userDefineMsg")
    def user_define_msg(self) -> Optional[pulumi.Input[str]]:
        """
        Customize the alarm notification content.
        """
        return pulumi.get(self, "user_define_msg")

    @user_define_msg.setter
    def user_define_msg(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_define_msg", value)


class Alarm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alarm_name: Optional[pulumi.Input[str]] = None,
                 alarm_notify_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_period: Optional[pulumi.Input[int]] = None,
                 alarm_period_detail: Optional[pulumi.Input[pulumi.InputType['AlarmAlarmPeriodDetailArgs']]] = None,
                 condition: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 query_requests: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmQueryRequestArgs']]]]] = None,
                 request_cycle: Optional[pulumi.Input[pulumi.InputType['AlarmRequestCycleArgs']]] = None,
                 status: Optional[pulumi.Input[bool]] = None,
                 trigger_period: Optional[pulumi.Input[int]] = None,
                 user_define_msg: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage tls alarm
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.tls.Alarm("foo",
            alarm_name="test",
            alarm_notify_groups=["3019107f-28a2-4208-a2b6-c33fcb97ac3a"],
            alarm_period_detail=volcengine.tls.AlarmAlarmPeriodDetailArgs(
                email=2,
                general_webhook=3,
                phone=10,
                sms=10,
            ),
            condition="$1.errNum>0",
            project_id="cc44f8b6-0328-4622-b043-023fca735cd4",
            query_requests=[volcengine.tls.AlarmQueryRequestArgs(
                end_time_offset=0,
                number=1,
                query="Failed | select count(*) as errNum",
                start_time_offset=-15,
                topic_id="af1a2240-ba62-4f18-b421-bde2f9684e57",
            )],
            request_cycle=volcengine.tls.AlarmRequestCycleArgs(
                time=11,
                type="Period",
            ),
            user_define_msg="test for terraform")
        ```

        ## Import

        tls alarm can be imported using the id and project id, e.g.

        ```sh
         $ pulumi import volcengine:tls/alarm:Alarm default projectId:fc************
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alarm_name: The name of the alarm.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_notify_groups: List of notification groups corresponding to the alarm.
        :param pulumi.Input[int] alarm_period: Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        :param pulumi.Input[pulumi.InputType['AlarmAlarmPeriodDetailArgs']] alarm_period_detail: Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        :param pulumi.Input[str] condition: Alarm trigger condition.
        :param pulumi.Input[str] project_id: The project id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmQueryRequestArgs']]]] query_requests: Search and analyze sentences, 1~3 can be configured.
        :param pulumi.Input[pulumi.InputType['AlarmRequestCycleArgs']] request_cycle: The execution period of the alarm task.
        :param pulumi.Input[bool] status: Whether to enable the alert policy. The default value is true, that is, on.
        :param pulumi.Input[int] trigger_period: Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
        :param pulumi.Input[str] user_define_msg: Customize the alarm notification content.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlarmArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage tls alarm
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.tls.Alarm("foo",
            alarm_name="test",
            alarm_notify_groups=["3019107f-28a2-4208-a2b6-c33fcb97ac3a"],
            alarm_period_detail=volcengine.tls.AlarmAlarmPeriodDetailArgs(
                email=2,
                general_webhook=3,
                phone=10,
                sms=10,
            ),
            condition="$1.errNum>0",
            project_id="cc44f8b6-0328-4622-b043-023fca735cd4",
            query_requests=[volcengine.tls.AlarmQueryRequestArgs(
                end_time_offset=0,
                number=1,
                query="Failed | select count(*) as errNum",
                start_time_offset=-15,
                topic_id="af1a2240-ba62-4f18-b421-bde2f9684e57",
            )],
            request_cycle=volcengine.tls.AlarmRequestCycleArgs(
                time=11,
                type="Period",
            ),
            user_define_msg="test for terraform")
        ```

        ## Import

        tls alarm can be imported using the id and project id, e.g.

        ```sh
         $ pulumi import volcengine:tls/alarm:Alarm default projectId:fc************
        ```

        :param str resource_name: The name of the resource.
        :param AlarmArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlarmArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alarm_name: Optional[pulumi.Input[str]] = None,
                 alarm_notify_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 alarm_period: Optional[pulumi.Input[int]] = None,
                 alarm_period_detail: Optional[pulumi.Input[pulumi.InputType['AlarmAlarmPeriodDetailArgs']]] = None,
                 condition: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 query_requests: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmQueryRequestArgs']]]]] = None,
                 request_cycle: Optional[pulumi.Input[pulumi.InputType['AlarmRequestCycleArgs']]] = None,
                 status: Optional[pulumi.Input[bool]] = None,
                 trigger_period: Optional[pulumi.Input[int]] = None,
                 user_define_msg: Optional[pulumi.Input[str]] = None,
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
            __props__ = AlarmArgs.__new__(AlarmArgs)

            if alarm_name is None and not opts.urn:
                raise TypeError("Missing required property 'alarm_name'")
            __props__.__dict__["alarm_name"] = alarm_name
            if alarm_notify_groups is None and not opts.urn:
                raise TypeError("Missing required property 'alarm_notify_groups'")
            __props__.__dict__["alarm_notify_groups"] = alarm_notify_groups
            __props__.__dict__["alarm_period"] = alarm_period
            __props__.__dict__["alarm_period_detail"] = alarm_period_detail
            if condition is None and not opts.urn:
                raise TypeError("Missing required property 'condition'")
            __props__.__dict__["condition"] = condition
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if query_requests is None and not opts.urn:
                raise TypeError("Missing required property 'query_requests'")
            __props__.__dict__["query_requests"] = query_requests
            if request_cycle is None and not opts.urn:
                raise TypeError("Missing required property 'request_cycle'")
            __props__.__dict__["request_cycle"] = request_cycle
            __props__.__dict__["status"] = status
            __props__.__dict__["trigger_period"] = trigger_period
            __props__.__dict__["user_define_msg"] = user_define_msg
            __props__.__dict__["alarm_id"] = None
        super(Alarm, __self__).__init__(
            'volcengine:tls/alarm:Alarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alarm_id: Optional[pulumi.Input[str]] = None,
            alarm_name: Optional[pulumi.Input[str]] = None,
            alarm_notify_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            alarm_period: Optional[pulumi.Input[int]] = None,
            alarm_period_detail: Optional[pulumi.Input[pulumi.InputType['AlarmAlarmPeriodDetailArgs']]] = None,
            condition: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            query_requests: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmQueryRequestArgs']]]]] = None,
            request_cycle: Optional[pulumi.Input[pulumi.InputType['AlarmRequestCycleArgs']]] = None,
            status: Optional[pulumi.Input[bool]] = None,
            trigger_period: Optional[pulumi.Input[int]] = None,
            user_define_msg: Optional[pulumi.Input[str]] = None) -> 'Alarm':
        """
        Get an existing Alarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alarm_id: The alarm id.
        :param pulumi.Input[str] alarm_name: The name of the alarm.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alarm_notify_groups: List of notification groups corresponding to the alarm.
        :param pulumi.Input[int] alarm_period: Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        :param pulumi.Input[pulumi.InputType['AlarmAlarmPeriodDetailArgs']] alarm_period_detail: Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        :param pulumi.Input[str] condition: Alarm trigger condition.
        :param pulumi.Input[str] project_id: The project id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AlarmQueryRequestArgs']]]] query_requests: Search and analyze sentences, 1~3 can be configured.
        :param pulumi.Input[pulumi.InputType['AlarmRequestCycleArgs']] request_cycle: The execution period of the alarm task.
        :param pulumi.Input[bool] status: Whether to enable the alert policy. The default value is true, that is, on.
        :param pulumi.Input[int] trigger_period: Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
        :param pulumi.Input[str] user_define_msg: Customize the alarm notification content.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlarmState.__new__(_AlarmState)

        __props__.__dict__["alarm_id"] = alarm_id
        __props__.__dict__["alarm_name"] = alarm_name
        __props__.__dict__["alarm_notify_groups"] = alarm_notify_groups
        __props__.__dict__["alarm_period"] = alarm_period
        __props__.__dict__["alarm_period_detail"] = alarm_period_detail
        __props__.__dict__["condition"] = condition
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["query_requests"] = query_requests
        __props__.__dict__["request_cycle"] = request_cycle
        __props__.__dict__["status"] = status
        __props__.__dict__["trigger_period"] = trigger_period
        __props__.__dict__["user_define_msg"] = user_define_msg
        return Alarm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alarmId")
    def alarm_id(self) -> pulumi.Output[str]:
        """
        The alarm id.
        """
        return pulumi.get(self, "alarm_id")

    @property
    @pulumi.getter(name="alarmName")
    def alarm_name(self) -> pulumi.Output[str]:
        """
        The name of the alarm.
        """
        return pulumi.get(self, "alarm_name")

    @property
    @pulumi.getter(name="alarmNotifyGroups")
    def alarm_notify_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        List of notification groups corresponding to the alarm.
        """
        return pulumi.get(self, "alarm_notify_groups")

    @property
    @pulumi.getter(name="alarmPeriod")
    def alarm_period(self) -> pulumi.Output[Optional[int]]:
        """
        Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        """
        return pulumi.get(self, "alarm_period")

    @property
    @pulumi.getter(name="alarmPeriodDetail")
    def alarm_period_detail(self) -> pulumi.Output[Optional['outputs.AlarmAlarmPeriodDetail']]:
        """
        Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.
        """
        return pulumi.get(self, "alarm_period_detail")

    @property
    @pulumi.getter
    def condition(self) -> pulumi.Output[str]:
        """
        Alarm trigger condition.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project id.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="queryRequests")
    def query_requests(self) -> pulumi.Output[Sequence['outputs.AlarmQueryRequest']]:
        """
        Search and analyze sentences, 1~3 can be configured.
        """
        return pulumi.get(self, "query_requests")

    @property
    @pulumi.getter(name="requestCycle")
    def request_cycle(self) -> pulumi.Output['outputs.AlarmRequestCycle']:
        """
        The execution period of the alarm task.
        """
        return pulumi.get(self, "request_cycle")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable the alert policy. The default value is true, that is, on.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="triggerPeriod")
    def trigger_period(self) -> pulumi.Output[Optional[int]]:
        """
        Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.
        """
        return pulumi.get(self, "trigger_period")

    @property
    @pulumi.getter(name="userDefineMsg")
    def user_define_msg(self) -> pulumi.Output[Optional[str]]:
        """
        Customize the alarm notification content.
        """
        return pulumi.get(self, "user_define_msg")
