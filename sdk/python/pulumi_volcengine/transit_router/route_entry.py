# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RouteEntryArgs', 'RouteEntry']

@pulumi.input_type
class RouteEntryArgs:
    def __init__(__self__, *,
                 destination_cidr_block: pulumi.Input[str],
                 transit_router_route_entry_next_hop_type: pulumi.Input[str],
                 transit_router_route_table_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_name: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_next_hop_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouteEntry resource.
        :param pulumi.Input[str] destination_cidr_block: The target network segment of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_next_hop_type: The next hop type of the routing entry. The value can be Attachment or BlackHole.
        :param pulumi.Input[str] transit_router_route_table_id: The id of the route table.
        :param pulumi.Input[str] description: Description of the transit router route entry.
        :param pulumi.Input[str] transit_router_route_entry_name: The name of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_next_hop_id: The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
        """
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "transit_router_route_entry_next_hop_type", transit_router_route_entry_next_hop_type)
        pulumi.set(__self__, "transit_router_route_table_id", transit_router_route_table_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if transit_router_route_entry_name is not None:
            pulumi.set(__self__, "transit_router_route_entry_name", transit_router_route_entry_name)
        if transit_router_route_entry_next_hop_id is not None:
            pulumi.set(__self__, "transit_router_route_entry_next_hop_id", transit_router_route_entry_next_hop_id)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Input[str]:
        """
        The target network segment of the route entry.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="transitRouterRouteEntryNextHopType")
    def transit_router_route_entry_next_hop_type(self) -> pulumi.Input[str]:
        """
        The next hop type of the routing entry. The value can be Attachment or BlackHole.
        """
        return pulumi.get(self, "transit_router_route_entry_next_hop_type")

    @transit_router_route_entry_next_hop_type.setter
    def transit_router_route_entry_next_hop_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_route_entry_next_hop_type", value)

    @property
    @pulumi.getter(name="transitRouterRouteTableId")
    def transit_router_route_table_id(self) -> pulumi.Input[str]:
        """
        The id of the route table.
        """
        return pulumi.get(self, "transit_router_route_table_id")

    @transit_router_route_table_id.setter
    def transit_router_route_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_route_table_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the transit router route entry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="transitRouterRouteEntryName")
    def transit_router_route_entry_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the route entry.
        """
        return pulumi.get(self, "transit_router_route_entry_name")

    @transit_router_route_entry_name.setter
    def transit_router_route_entry_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_entry_name", value)

    @property
    @pulumi.getter(name="transitRouterRouteEntryNextHopId")
    def transit_router_route_entry_next_hop_id(self) -> Optional[pulumi.Input[str]]:
        """
        The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
        """
        return pulumi.get(self, "transit_router_route_entry_next_hop_id")

    @transit_router_route_entry_next_hop_id.setter
    def transit_router_route_entry_next_hop_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_entry_next_hop_id", value)


@pulumi.input_type
class _RouteEntryState:
    def __init__(__self__, *,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_id: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_name: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_next_hop_id: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_next_hop_type: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_type: Optional[pulumi.Input[str]] = None,
                 transit_router_route_table_id: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouteEntry resources.
        :param pulumi.Input[str] creation_time: The creation time of the route entry.
        :param pulumi.Input[str] description: Description of the transit router route entry.
        :param pulumi.Input[str] destination_cidr_block: The target network segment of the route entry.
        :param pulumi.Input[str] status: The status of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_id: The id of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_name: The name of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_next_hop_id: The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
        :param pulumi.Input[str] transit_router_route_entry_next_hop_type: The next hop type of the routing entry. The value can be Attachment or BlackHole.
        :param pulumi.Input[str] transit_router_route_entry_type: The type of the route entry.
        :param pulumi.Input[str] transit_router_route_table_id: The id of the route table.
        :param pulumi.Input[str] update_time: The update time of the route entry.
        """
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_cidr_block is not None:
            pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transit_router_route_entry_id is not None:
            pulumi.set(__self__, "transit_router_route_entry_id", transit_router_route_entry_id)
        if transit_router_route_entry_name is not None:
            pulumi.set(__self__, "transit_router_route_entry_name", transit_router_route_entry_name)
        if transit_router_route_entry_next_hop_id is not None:
            pulumi.set(__self__, "transit_router_route_entry_next_hop_id", transit_router_route_entry_next_hop_id)
        if transit_router_route_entry_next_hop_type is not None:
            pulumi.set(__self__, "transit_router_route_entry_next_hop_type", transit_router_route_entry_next_hop_type)
        if transit_router_route_entry_type is not None:
            pulumi.set(__self__, "transit_router_route_entry_type", transit_router_route_entry_type)
        if transit_router_route_table_id is not None:
            pulumi.set(__self__, "transit_router_route_table_id", transit_router_route_table_id)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of the route entry.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the transit router route entry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The target network segment of the route entry.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the route entry.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="transitRouterRouteEntryId")
    def transit_router_route_entry_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the route entry.
        """
        return pulumi.get(self, "transit_router_route_entry_id")

    @transit_router_route_entry_id.setter
    def transit_router_route_entry_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_entry_id", value)

    @property
    @pulumi.getter(name="transitRouterRouteEntryName")
    def transit_router_route_entry_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the route entry.
        """
        return pulumi.get(self, "transit_router_route_entry_name")

    @transit_router_route_entry_name.setter
    def transit_router_route_entry_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_entry_name", value)

    @property
    @pulumi.getter(name="transitRouterRouteEntryNextHopId")
    def transit_router_route_entry_next_hop_id(self) -> Optional[pulumi.Input[str]]:
        """
        The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
        """
        return pulumi.get(self, "transit_router_route_entry_next_hop_id")

    @transit_router_route_entry_next_hop_id.setter
    def transit_router_route_entry_next_hop_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_entry_next_hop_id", value)

    @property
    @pulumi.getter(name="transitRouterRouteEntryNextHopType")
    def transit_router_route_entry_next_hop_type(self) -> Optional[pulumi.Input[str]]:
        """
        The next hop type of the routing entry. The value can be Attachment or BlackHole.
        """
        return pulumi.get(self, "transit_router_route_entry_next_hop_type")

    @transit_router_route_entry_next_hop_type.setter
    def transit_router_route_entry_next_hop_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_entry_next_hop_type", value)

    @property
    @pulumi.getter(name="transitRouterRouteEntryType")
    def transit_router_route_entry_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the route entry.
        """
        return pulumi.get(self, "transit_router_route_entry_type")

    @transit_router_route_entry_type.setter
    def transit_router_route_entry_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_entry_type", value)

    @property
    @pulumi.getter(name="transitRouterRouteTableId")
    def transit_router_route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the route table.
        """
        return pulumi.get(self, "transit_router_route_table_id")

    @transit_router_route_table_id.setter
    def transit_router_route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_route_table_id", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of the route entry.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class RouteEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_name: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_next_hop_id: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_next_hop_type: Optional[pulumi.Input[str]] = None,
                 transit_router_route_table_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage transit router route entry
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.transit_router.RouteEntry("foo",
            description="tf test 23",
            destination_cidr_block="192.168.0.0/24",
            transit_router_route_entry_name="tf-entry-23",
            transit_router_route_entry_next_hop_type="BlackHole",
            transit_router_route_table_id="tr-rtb-12b7qd3fmzf2817q7y2jkbd55")
        ```

        ## Import

        transit router route entry can be imported using the table and entry id, e.g.

        ```sh
         $ pulumi import volcengine:transit_router/routeEntry:RouteEntry default tr-rtb-12b7qd3fmzf2817q7y2jkbd55:tr-rte-1i5i8khf9m58gae5kcx6***
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the transit router route entry.
        :param pulumi.Input[str] destination_cidr_block: The target network segment of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_name: The name of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_next_hop_id: The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
        :param pulumi.Input[str] transit_router_route_entry_next_hop_type: The next hop type of the routing entry. The value can be Attachment or BlackHole.
        :param pulumi.Input[str] transit_router_route_table_id: The id of the route table.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage transit router route entry
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.transit_router.RouteEntry("foo",
            description="tf test 23",
            destination_cidr_block="192.168.0.0/24",
            transit_router_route_entry_name="tf-entry-23",
            transit_router_route_entry_next_hop_type="BlackHole",
            transit_router_route_table_id="tr-rtb-12b7qd3fmzf2817q7y2jkbd55")
        ```

        ## Import

        transit router route entry can be imported using the table and entry id, e.g.

        ```sh
         $ pulumi import volcengine:transit_router/routeEntry:RouteEntry default tr-rtb-12b7qd3fmzf2817q7y2jkbd55:tr-rte-1i5i8khf9m58gae5kcx6***
        ```

        :param str resource_name: The name of the resource.
        :param RouteEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_name: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_next_hop_id: Optional[pulumi.Input[str]] = None,
                 transit_router_route_entry_next_hop_type: Optional[pulumi.Input[str]] = None,
                 transit_router_route_table_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouteEntryArgs.__new__(RouteEntryArgs)

            __props__.__dict__["description"] = description
            if destination_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            __props__.__dict__["transit_router_route_entry_name"] = transit_router_route_entry_name
            __props__.__dict__["transit_router_route_entry_next_hop_id"] = transit_router_route_entry_next_hop_id
            if transit_router_route_entry_next_hop_type is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_route_entry_next_hop_type'")
            __props__.__dict__["transit_router_route_entry_next_hop_type"] = transit_router_route_entry_next_hop_type
            if transit_router_route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_route_table_id'")
            __props__.__dict__["transit_router_route_table_id"] = transit_router_route_table_id
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["transit_router_route_entry_id"] = None
            __props__.__dict__["transit_router_route_entry_type"] = None
            __props__.__dict__["update_time"] = None
        super(RouteEntry, __self__).__init__(
            'volcengine:transit_router/routeEntry:RouteEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            transit_router_route_entry_id: Optional[pulumi.Input[str]] = None,
            transit_router_route_entry_name: Optional[pulumi.Input[str]] = None,
            transit_router_route_entry_next_hop_id: Optional[pulumi.Input[str]] = None,
            transit_router_route_entry_next_hop_type: Optional[pulumi.Input[str]] = None,
            transit_router_route_entry_type: Optional[pulumi.Input[str]] = None,
            transit_router_route_table_id: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'RouteEntry':
        """
        Get an existing RouteEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_time: The creation time of the route entry.
        :param pulumi.Input[str] description: Description of the transit router route entry.
        :param pulumi.Input[str] destination_cidr_block: The target network segment of the route entry.
        :param pulumi.Input[str] status: The status of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_id: The id of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_name: The name of the route entry.
        :param pulumi.Input[str] transit_router_route_entry_next_hop_id: The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
        :param pulumi.Input[str] transit_router_route_entry_next_hop_type: The next hop type of the routing entry. The value can be Attachment or BlackHole.
        :param pulumi.Input[str] transit_router_route_entry_type: The type of the route entry.
        :param pulumi.Input[str] transit_router_route_table_id: The id of the route table.
        :param pulumi.Input[str] update_time: The update time of the route entry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouteEntryState.__new__(_RouteEntryState)

        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["destination_cidr_block"] = destination_cidr_block
        __props__.__dict__["status"] = status
        __props__.__dict__["transit_router_route_entry_id"] = transit_router_route_entry_id
        __props__.__dict__["transit_router_route_entry_name"] = transit_router_route_entry_name
        __props__.__dict__["transit_router_route_entry_next_hop_id"] = transit_router_route_entry_next_hop_id
        __props__.__dict__["transit_router_route_entry_next_hop_type"] = transit_router_route_entry_next_hop_type
        __props__.__dict__["transit_router_route_entry_type"] = transit_router_route_entry_type
        __props__.__dict__["transit_router_route_table_id"] = transit_router_route_table_id
        __props__.__dict__["update_time"] = update_time
        return RouteEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The creation time of the route entry.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the transit router route entry.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[str]:
        """
        The target network segment of the route entry.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the route entry.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterRouteEntryId")
    def transit_router_route_entry_id(self) -> pulumi.Output[str]:
        """
        The id of the route entry.
        """
        return pulumi.get(self, "transit_router_route_entry_id")

    @property
    @pulumi.getter(name="transitRouterRouteEntryName")
    def transit_router_route_entry_name(self) -> pulumi.Output[str]:
        """
        The name of the route entry.
        """
        return pulumi.get(self, "transit_router_route_entry_name")

    @property
    @pulumi.getter(name="transitRouterRouteEntryNextHopId")
    def transit_router_route_entry_next_hop_id(self) -> pulumi.Output[Optional[str]]:
        """
        The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.
        """
        return pulumi.get(self, "transit_router_route_entry_next_hop_id")

    @property
    @pulumi.getter(name="transitRouterRouteEntryNextHopType")
    def transit_router_route_entry_next_hop_type(self) -> pulumi.Output[str]:
        """
        The next hop type of the routing entry. The value can be Attachment or BlackHole.
        """
        return pulumi.get(self, "transit_router_route_entry_next_hop_type")

    @property
    @pulumi.getter(name="transitRouterRouteEntryType")
    def transit_router_route_entry_type(self) -> pulumi.Output[str]:
        """
        The type of the route entry.
        """
        return pulumi.get(self, "transit_router_route_entry_type")

    @property
    @pulumi.getter(name="transitRouterRouteTableId")
    def transit_router_route_table_id(self) -> pulumi.Output[str]:
        """
        The id of the route table.
        """
        return pulumi.get(self, "transit_router_route_table_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time of the route entry.
        """
        return pulumi.get(self, "update_time")

