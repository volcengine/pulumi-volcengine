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

__all__ = ['TransitRouterArgs', 'TransitRouter']

@pulumi.input_type
class TransitRouterArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 transit_router_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TransitRouter resource.
        :param pulumi.Input[str] description: The description of the transit router.
        :param pulumi.Input[str] transit_router_name: The name of the transit router.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if transit_router_name is not None:
            pulumi.set(__self__, "transit_router_name", transit_router_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the transit router.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="transitRouterName")
    def transit_router_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the transit router.
        """
        return pulumi.get(self, "transit_router_name")

    @transit_router_name.setter
    def transit_router_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_name", value)


@pulumi.input_type
class _TransitRouterState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 business_status: Optional[pulumi.Input[str]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 overdue_time: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transit_router_attachments: Optional[pulumi.Input[Sequence[pulumi.Input['TransitRouterTransitRouterAttachmentArgs']]]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 transit_router_name: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TransitRouter resources.
        :param pulumi.Input[str] account_id: The ID of account.
        :param pulumi.Input[str] business_status: The business status of the transit router.
        :param pulumi.Input[str] creation_time: The create time.
        :param pulumi.Input[str] description: The description of the transit router.
        :param pulumi.Input[str] overdue_time: The overdue time.
        :param pulumi.Input[str] status: The status of the transit router.
        :param pulumi.Input[Sequence[pulumi.Input['TransitRouterTransitRouterAttachmentArgs']]] transit_router_attachments: The attachments of transit router.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        :param pulumi.Input[str] transit_router_name: The name of the transit router.
        :param pulumi.Input[str] update_time: The update time.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if business_status is not None:
            pulumi.set(__self__, "business_status", business_status)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if overdue_time is not None:
            pulumi.set(__self__, "overdue_time", overdue_time)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transit_router_attachments is not None:
            pulumi.set(__self__, "transit_router_attachments", transit_router_attachments)
        if transit_router_id is not None:
            pulumi.set(__self__, "transit_router_id", transit_router_id)
        if transit_router_name is not None:
            pulumi.set(__self__, "transit_router_name", transit_router_name)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of account.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> Optional[pulumi.Input[str]]:
        """
        The business status of the transit router.
        """
        return pulumi.get(self, "business_status")

    @business_status.setter
    def business_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "business_status", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the transit router.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="overdueTime")
    def overdue_time(self) -> Optional[pulumi.Input[str]]:
        """
        The overdue time.
        """
        return pulumi.get(self, "overdue_time")

    @overdue_time.setter
    def overdue_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "overdue_time", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the transit router.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="transitRouterAttachments")
    def transit_router_attachments(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TransitRouterTransitRouterAttachmentArgs']]]]:
        """
        The attachments of transit router.
        """
        return pulumi.get(self, "transit_router_attachments")

    @transit_router_attachments.setter
    def transit_router_attachments(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TransitRouterTransitRouterAttachmentArgs']]]]):
        pulumi.set(self, "transit_router_attachments", value)

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

    @transit_router_id.setter
    def transit_router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_id", value)

    @property
    @pulumi.getter(name="transitRouterName")
    def transit_router_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the transit router.
        """
        return pulumi.get(self, "transit_router_name")

    @transit_router_name.setter
    def transit_router_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_name", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update time.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class TransitRouter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 transit_router_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage transit router
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.transit_router.TransitRouter("foo",
            description="acc-test",
            transit_router_name="acc-test-tr")
        ```

        ## Import

        TransitRouter can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:transit_router/transitRouter:TransitRouter default tr-2d6fr7mzya2gw58ozfes5g2oh
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the transit router.
        :param pulumi.Input[str] transit_router_name: The name of the transit router.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TransitRouterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage transit router
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.transit_router.TransitRouter("foo",
            description="acc-test",
            transit_router_name="acc-test-tr")
        ```

        ## Import

        TransitRouter can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:transit_router/transitRouter:TransitRouter default tr-2d6fr7mzya2gw58ozfes5g2oh
        ```

        :param str resource_name: The name of the resource.
        :param TransitRouterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransitRouterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 transit_router_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransitRouterArgs.__new__(TransitRouterArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["transit_router_name"] = transit_router_name
            __props__.__dict__["account_id"] = None
            __props__.__dict__["business_status"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["overdue_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["transit_router_attachments"] = None
            __props__.__dict__["transit_router_id"] = None
            __props__.__dict__["update_time"] = None
        super(TransitRouter, __self__).__init__(
            'volcengine:transit_router/transitRouter:TransitRouter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            business_status: Optional[pulumi.Input[str]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            overdue_time: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            transit_router_attachments: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransitRouterTransitRouterAttachmentArgs']]]]] = None,
            transit_router_id: Optional[pulumi.Input[str]] = None,
            transit_router_name: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'TransitRouter':
        """
        Get an existing TransitRouter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The ID of account.
        :param pulumi.Input[str] business_status: The business status of the transit router.
        :param pulumi.Input[str] creation_time: The create time.
        :param pulumi.Input[str] description: The description of the transit router.
        :param pulumi.Input[str] overdue_time: The overdue time.
        :param pulumi.Input[str] status: The status of the transit router.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransitRouterTransitRouterAttachmentArgs']]]] transit_router_attachments: The attachments of transit router.
        :param pulumi.Input[str] transit_router_id: The ID of the transit router.
        :param pulumi.Input[str] transit_router_name: The name of the transit router.
        :param pulumi.Input[str] update_time: The update time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransitRouterState.__new__(_TransitRouterState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["business_status"] = business_status
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["overdue_time"] = overdue_time
        __props__.__dict__["status"] = status
        __props__.__dict__["transit_router_attachments"] = transit_router_attachments
        __props__.__dict__["transit_router_id"] = transit_router_id
        __props__.__dict__["transit_router_name"] = transit_router_name
        __props__.__dict__["update_time"] = update_time
        return TransitRouter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The ID of account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> pulumi.Output[str]:
        """
        The business status of the transit router.
        """
        return pulumi.get(self, "business_status")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The create time.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the transit router.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="overdueTime")
    def overdue_time(self) -> pulumi.Output[str]:
        """
        The overdue time.
        """
        return pulumi.get(self, "overdue_time")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the transit router.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterAttachments")
    def transit_router_attachments(self) -> pulumi.Output[Sequence['outputs.TransitRouterTransitRouterAttachment']]:
        """
        The attachments of transit router.
        """
        return pulumi.get(self, "transit_router_attachments")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> pulumi.Output[str]:
        """
        The ID of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

    @property
    @pulumi.getter(name="transitRouterName")
    def transit_router_name(self) -> pulumi.Output[str]:
        """
        The name of the transit router.
        """
        return pulumi.get(self, "transit_router_name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time.
        """
        return pulumi.get(self, "update_time")
