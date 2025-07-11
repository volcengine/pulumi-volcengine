# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FlowLogActiveArgs', 'FlowLogActive']

@pulumi.input_type
class FlowLogActiveArgs:
    def __init__(__self__, *,
                 flow_log_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a FlowLogActive resource.
        :param pulumi.Input[str] flow_log_id: The ID of flow log.
        """
        pulumi.set(__self__, "flow_log_id", flow_log_id)

    @property
    @pulumi.getter(name="flowLogId")
    def flow_log_id(self) -> pulumi.Input[str]:
        """
        The ID of flow log.
        """
        return pulumi.get(self, "flow_log_id")

    @flow_log_id.setter
    def flow_log_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "flow_log_id", value)


@pulumi.input_type
class _FlowLogActiveState:
    def __init__(__self__, *,
                 flow_log_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FlowLogActive resources.
        :param pulumi.Input[str] flow_log_id: The ID of flow log.
        :param pulumi.Input[str] status: The status of flow log.
        """
        if flow_log_id is not None:
            pulumi.set(__self__, "flow_log_id", flow_log_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="flowLogId")
    def flow_log_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of flow log.
        """
        return pulumi.get(self, "flow_log_id")

    @flow_log_id.setter
    def flow_log_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flow_log_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of flow log.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class FlowLogActive(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flow_log_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage flow log active
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.get_zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16",
            project_name="default")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        foo_flow_log = volcengine.vpc.FlowLog("fooFlowLog",
            flow_log_name="acc-test-flow-log",
            description="acc-test",
            resource_type="subnet",
            resource_id=foo_subnet.id,
            traffic_type="All",
            log_project_name="acc-test-project",
            log_topic_name="acc-test-topic",
            aggregation_interval=10,
            project_name="default",
            tags=[volcengine.vpc.FlowLogTagArgs(
                key="k1",
                value="v1",
            )])
        foo_flow_log_active = volcengine.vpc.FlowLogActive("fooFlowLogActive", flow_log_id=foo_flow_log.id)
        ```

        ## Import

        FlowLogActive can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:vpc/flowLogActive:FlowLogActive default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] flow_log_id: The ID of flow log.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowLogActiveArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage flow log active
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.get_zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16",
            project_name="default")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        foo_flow_log = volcengine.vpc.FlowLog("fooFlowLog",
            flow_log_name="acc-test-flow-log",
            description="acc-test",
            resource_type="subnet",
            resource_id=foo_subnet.id,
            traffic_type="All",
            log_project_name="acc-test-project",
            log_topic_name="acc-test-topic",
            aggregation_interval=10,
            project_name="default",
            tags=[volcengine.vpc.FlowLogTagArgs(
                key="k1",
                value="v1",
            )])
        foo_flow_log_active = volcengine.vpc.FlowLogActive("fooFlowLogActive", flow_log_id=foo_flow_log.id)
        ```

        ## Import

        FlowLogActive can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:vpc/flowLogActive:FlowLogActive default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param FlowLogActiveArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowLogActiveArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 flow_log_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlowLogActiveArgs.__new__(FlowLogActiveArgs)

            if flow_log_id is None and not opts.urn:
                raise TypeError("Missing required property 'flow_log_id'")
            __props__.__dict__["flow_log_id"] = flow_log_id
            __props__.__dict__["status"] = None
        super(FlowLogActive, __self__).__init__(
            'volcengine:vpc/flowLogActive:FlowLogActive',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            flow_log_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'FlowLogActive':
        """
        Get an existing FlowLogActive resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] flow_log_id: The ID of flow log.
        :param pulumi.Input[str] status: The status of flow log.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlowLogActiveState.__new__(_FlowLogActiveState)

        __props__.__dict__["flow_log_id"] = flow_log_id
        __props__.__dict__["status"] = status
        return FlowLogActive(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="flowLogId")
    def flow_log_id(self) -> pulumi.Output[str]:
        """
        The ID of flow log.
        """
        return pulumi.get(self, "flow_log_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of flow log.
        """
        return pulumi.get(self, "status")

