# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['StateArgs', 'State']

@pulumi.input_type
class StateArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 stopped_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a State resource.
        :param pulumi.Input[str] action: Start or Stop of Instance Action, the value can be `Start`, `Stop` or `ForceStop`. 
               If the target status of the action is consistent with the current status of the instance, the action will not actually be executed.
        :param pulumi.Input[str] instance_id: Id of Instance.
        :param pulumi.Input[str] stopped_mode: Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "instance_id", instance_id)
        if stopped_mode is not None:
            pulumi.set(__self__, "stopped_mode", stopped_mode)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        """
        Start or Stop of Instance Action, the value can be `Start`, `Stop` or `ForceStop`. 
        If the target status of the action is consistent with the current status of the instance, the action will not actually be executed.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Id of Instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="stoppedMode")
    def stopped_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`.
        """
        return pulumi.get(self, "stopped_mode")

    @stopped_mode.setter
    def stopped_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stopped_mode", value)


@pulumi.input_type
class _StateState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 stopped_mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering State resources.
        :param pulumi.Input[str] action: Start or Stop of Instance Action, the value can be `Start`, `Stop` or `ForceStop`. 
               If the target status of the action is consistent with the current status of the instance, the action will not actually be executed.
        :param pulumi.Input[str] instance_id: Id of Instance.
        :param pulumi.Input[str] status: Status of Instance.
        :param pulumi.Input[str] stopped_mode: Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`.
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if stopped_mode is not None:
            pulumi.set(__self__, "stopped_mode", stopped_mode)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        Start or Stop of Instance Action, the value can be `Start`, `Stop` or `ForceStop`. 
        If the target status of the action is consistent with the current status of the instance, the action will not actually be executed.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of Instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of Instance.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="stoppedMode")
    def stopped_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`.
        """
        return pulumi.get(self, "stopped_mode")

    @stopped_mode.setter
    def stopped_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stopped_mode", value)


class State(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 stopped_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage ecs instance state
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.get_zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
            security_group_name="acc-test-security-group",
            vpc_id=foo_vpc.id)
        foo_images = volcengine.ecs.get_images(os_type="Linux",
            visibility="public",
            instance_type_id="ecs.g1.large")
        foo_instance = volcengine.ecs.Instance("fooInstance",
            instance_name="acc-test-ecs",
            image_id=foo_images.images[0].image_id,
            instance_type="ecs.g1.large",
            password="93f0cb0614Aab12",
            instance_charge_type="PostPaid",
            system_volume_type="ESSD_PL0",
            system_volume_size=40,
            subnet_id=foo_subnet.id,
            security_group_ids=[foo_security_group.id])
        foo_state = volcengine.ecs.State("fooState",
            instance_id=foo_instance.id,
            action="Stop",
            stopped_mode="KeepCharging")
        ```

        ## Import

        State Instance can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:ecs/state:State default state:i-mizl7m1kqccg5smt1bdpijuj
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Start or Stop of Instance Action, the value can be `Start`, `Stop` or `ForceStop`. 
               If the target status of the action is consistent with the current status of the instance, the action will not actually be executed.
        :param pulumi.Input[str] instance_id: Id of Instance.
        :param pulumi.Input[str] stopped_mode: Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage ecs instance state
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.get_zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
            security_group_name="acc-test-security-group",
            vpc_id=foo_vpc.id)
        foo_images = volcengine.ecs.get_images(os_type="Linux",
            visibility="public",
            instance_type_id="ecs.g1.large")
        foo_instance = volcengine.ecs.Instance("fooInstance",
            instance_name="acc-test-ecs",
            image_id=foo_images.images[0].image_id,
            instance_type="ecs.g1.large",
            password="93f0cb0614Aab12",
            instance_charge_type="PostPaid",
            system_volume_type="ESSD_PL0",
            system_volume_size=40,
            subnet_id=foo_subnet.id,
            security_group_ids=[foo_security_group.id])
        foo_state = volcengine.ecs.State("fooState",
            instance_id=foo_instance.id,
            action="Stop",
            stopped_mode="KeepCharging")
        ```

        ## Import

        State Instance can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:ecs/state:State default state:i-mizl7m1kqccg5smt1bdpijuj
        ```

        :param str resource_name: The name of the resource.
        :param StateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 stopped_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StateArgs.__new__(StateArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["stopped_mode"] = stopped_mode
            __props__.__dict__["status"] = None
        super(State, __self__).__init__(
            'volcengine:ecs/state:State',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            stopped_mode: Optional[pulumi.Input[str]] = None) -> 'State':
        """
        Get an existing State resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: Start or Stop of Instance Action, the value can be `Start`, `Stop` or `ForceStop`. 
               If the target status of the action is consistent with the current status of the instance, the action will not actually be executed.
        :param pulumi.Input[str] instance_id: Id of Instance.
        :param pulumi.Input[str] status: Status of Instance.
        :param pulumi.Input[str] stopped_mode: Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StateState.__new__(_StateState)

        __props__.__dict__["action"] = action
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["status"] = status
        __props__.__dict__["stopped_mode"] = stopped_mode
        return State(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        Start or Stop of Instance Action, the value can be `Start`, `Stop` or `ForceStop`. 
        If the target status of the action is consistent with the current status of the instance, the action will not actually be executed.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Id of Instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of Instance.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="stoppedMode")
    def stopped_mode(self) -> pulumi.Output[str]:
        """
        Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`.
        """
        return pulumi.get(self, "stopped_mode")

