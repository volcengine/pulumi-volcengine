# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GrantRuleArgs', 'GrantRule']

@pulumi.input_type
class GrantRuleArgs:
    def __init__(__self__, *,
                 grant_account_id: pulumi.Input[str],
                 transit_router_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GrantRule resource.
        :param pulumi.Input[str] grant_account_id: Account ID awaiting authorization for intermediate router instance.
        :param pulumi.Input[str] transit_router_id: The id of the transit router.
        :param pulumi.Input[str] description: The description of the rule.
        """
        pulumi.set(__self__, "grant_account_id", grant_account_id)
        pulumi.set(__self__, "transit_router_id", transit_router_id)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="grantAccountId")
    def grant_account_id(self) -> pulumi.Input[str]:
        """
        Account ID awaiting authorization for intermediate router instance.
        """
        return pulumi.get(self, "grant_account_id")

    @grant_account_id.setter
    def grant_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "grant_account_id", value)

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> pulumi.Input[str]:
        """
        The id of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

    @transit_router_id.setter
    def transit_router_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_router_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _GrantRuleState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 grant_account_id: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GrantRule resources.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] grant_account_id: Account ID awaiting authorization for intermediate router instance.
        :param pulumi.Input[str] transit_router_id: The id of the transit router.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if grant_account_id is not None:
            pulumi.set(__self__, "grant_account_id", grant_account_id)
        if transit_router_id is not None:
            pulumi.set(__self__, "transit_router_id", transit_router_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="grantAccountId")
    def grant_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Account ID awaiting authorization for intermediate router instance.
        """
        return pulumi.get(self, "grant_account_id")

    @grant_account_id.setter
    def grant_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "grant_account_id", value)

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

    @transit_router_id.setter
    def transit_router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_id", value)


class GrantRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 grant_account_id: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage transit router grant rule
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.transit_router.GrantRule("foo",
            description="tf-test",
            grant_account_id="200000xxxx",
            transit_router_id="tr-2bzy39uy6u3282dx0efxiqyq0")
        ```

        ## Import

        TransitRouterGrantRule can be imported using the transit router id and accountId, e.g.

        ```sh
         $ pulumi import volcengine:transit_router/grantRule:GrantRule default trId:accountId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] grant_account_id: Account ID awaiting authorization for intermediate router instance.
        :param pulumi.Input[str] transit_router_id: The id of the transit router.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GrantRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage transit router grant rule
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.transit_router.GrantRule("foo",
            description="tf-test",
            grant_account_id="200000xxxx",
            transit_router_id="tr-2bzy39uy6u3282dx0efxiqyq0")
        ```

        ## Import

        TransitRouterGrantRule can be imported using the transit router id and accountId, e.g.

        ```sh
         $ pulumi import volcengine:transit_router/grantRule:GrantRule default trId:accountId
        ```

        :param str resource_name: The name of the resource.
        :param GrantRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GrantRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 grant_account_id: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GrantRuleArgs.__new__(GrantRuleArgs)

            __props__.__dict__["description"] = description
            if grant_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'grant_account_id'")
            __props__.__dict__["grant_account_id"] = grant_account_id
            if transit_router_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_id'")
            __props__.__dict__["transit_router_id"] = transit_router_id
        super(GrantRule, __self__).__init__(
            'volcengine:transit_router/grantRule:GrantRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            grant_account_id: Optional[pulumi.Input[str]] = None,
            transit_router_id: Optional[pulumi.Input[str]] = None) -> 'GrantRule':
        """
        Get an existing GrantRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] grant_account_id: Account ID awaiting authorization for intermediate router instance.
        :param pulumi.Input[str] transit_router_id: The id of the transit router.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GrantRuleState.__new__(_GrantRuleState)

        __props__.__dict__["description"] = description
        __props__.__dict__["grant_account_id"] = grant_account_id
        __props__.__dict__["transit_router_id"] = transit_router_id
        return GrantRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="grantAccountId")
    def grant_account_id(self) -> pulumi.Output[str]:
        """
        Account ID awaiting authorization for intermediate router instance.
        """
        return pulumi.get(self, "grant_account_id")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> pulumi.Output[str]:
        """
        The id of the transit router.
        """
        return pulumi.get(self, "transit_router_id")
