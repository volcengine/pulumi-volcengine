# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RouteTableArgs', 'RouteTable']

@pulumi.input_type
class RouteTableArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouteTable resource.
        :param pulumi.Input[str] vpc_id: The id of the VPC.
        :param pulumi.Input[str] description: The description of the route table.
        :param pulumi.Input[str] project_name: The ProjectName of the route table.
        :param pulumi.Input[str] route_table_name: The name of the route table.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if route_table_name is not None:
            pulumi.set(__self__, "route_table_name", route_table_name)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The id of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the route table.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of the route table.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the route table.
        """
        return pulumi.get(self, "route_table_name")

    @route_table_name.setter
    def route_table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_name", value)


@pulumi.input_type
class _RouteTableState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouteTable resources.
        :param pulumi.Input[str] description: The description of the route table.
        :param pulumi.Input[str] project_name: The ProjectName of the route table.
        :param pulumi.Input[str] route_table_name: The name of the route table.
        :param pulumi.Input[str] vpc_id: The id of the VPC.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if route_table_name is not None:
            pulumi.set(__self__, "route_table_name", route_table_name)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the route table.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of the route table.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the route table.
        """
        return pulumi.get(self, "route_table_name")

    @route_table_name.setter
    def route_table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_name", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class RouteTable(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage route table
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpc.RouteTable("foo",
            description="tf-test1",
            project_name="yuwao",
            route_table_name="tf-project-1",
            vpc_id="vpc-2feppmy1ugt1c59gp688n1fld")
        ```

        ## Import

        Route table can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/routeTable:RouteTable default vtb-274e0syt9av407fap8tle16kb
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the route table.
        :param pulumi.Input[str] project_name: The ProjectName of the route table.
        :param pulumi.Input[str] route_table_name: The name of the route table.
        :param pulumi.Input[str] vpc_id: The id of the VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteTableArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage route table
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpc.RouteTable("foo",
            description="tf-test1",
            project_name="yuwao",
            route_table_name="tf-project-1",
            vpc_id="vpc-2feppmy1ugt1c59gp688n1fld")
        ```

        ## Import

        Route table can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/routeTable:RouteTable default vtb-274e0syt9av407fap8tle16kb
        ```

        :param str resource_name: The name of the resource.
        :param RouteTableArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteTableArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 route_table_name: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RouteTableArgs.__new__(RouteTableArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["route_table_name"] = route_table_name
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(RouteTable, __self__).__init__(
            'volcengine:vpc/routeTable:RouteTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            route_table_name: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'RouteTable':
        """
        Get an existing RouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the route table.
        :param pulumi.Input[str] project_name: The ProjectName of the route table.
        :param pulumi.Input[str] route_table_name: The name of the route table.
        :param pulumi.Input[str] vpc_id: The id of the VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouteTableState.__new__(_RouteTableState)

        __props__.__dict__["description"] = description
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["route_table_name"] = route_table_name
        __props__.__dict__["vpc_id"] = vpc_id
        return RouteTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the route table.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[Optional[str]]:
        """
        The ProjectName of the route table.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="routeTableName")
    def route_table_name(self) -> pulumi.Output[str]:
        """
        The name of the route table.
        """
        return pulumi.get(self, "route_table_name")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The id of the VPC.
        """
        return pulumi.get(self, "vpc_id")
