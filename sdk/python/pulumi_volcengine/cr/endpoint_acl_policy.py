# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EndpointAclPolicyInitArgs', 'EndpointAclPolicy']

@pulumi.input_type
class EndpointAclPolicyInitArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 entry: pulumi.Input[str],
                 registry: pulumi.Input[str],
                 type: pulumi.Input[str]):
        """
        The set of arguments for constructing a EndpointAclPolicy resource.
        :param pulumi.Input[str] description: The description of the acl policy.
        :param pulumi.Input[str] entry: The ip list of the acl policy.
        :param pulumi.Input[str] registry: The registry name.
        :param pulumi.Input[str] type: The type of the acl policy. Valid values: `Public`.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "entry", entry)
        pulumi.set(__self__, "registry", registry)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The description of the acl policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def entry(self) -> pulumi.Input[str]:
        """
        The ip list of the acl policy.
        """
        return pulumi.get(self, "entry")

    @entry.setter
    def entry(self, value: pulumi.Input[str]):
        pulumi.set(self, "entry", value)

    @property
    @pulumi.getter
    def registry(self) -> pulumi.Input[str]:
        """
        The registry name.
        """
        return pulumi.get(self, "registry")

    @registry.setter
    def registry(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the acl policy. Valid values: `Public`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _EndpointAclPolicyState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 entry: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EndpointAclPolicy resources.
        :param pulumi.Input[str] description: The description of the acl policy.
        :param pulumi.Input[str] entry: The ip list of the acl policy.
        :param pulumi.Input[str] registry: The registry name.
        :param pulumi.Input[str] type: The type of the acl policy. Valid values: `Public`.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if entry is not None:
            pulumi.set(__self__, "entry", entry)
        if registry is not None:
            pulumi.set(__self__, "registry", registry)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the acl policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def entry(self) -> Optional[pulumi.Input[str]]:
        """
        The ip list of the acl policy.
        """
        return pulumi.get(self, "entry")

    @entry.setter
    def entry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entry", value)

    @property
    @pulumi.getter
    def registry(self) -> Optional[pulumi.Input[str]]:
        """
        The registry name.
        """
        return pulumi.get(self, "registry")

    @registry.setter
    def registry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the acl policy. Valid values: `Public`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class EndpointAclPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entry: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage cr endpoint acl policy
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_registry = volcengine.cr.Registry("fooRegistry", project="default")
        foo_endpoint = volcengine.cr.Endpoint("fooEndpoint",
            registry=foo_registry.id,
            enabled=True)
        foo_endpoint_acl_policy = []
        for range in [{"value": i} for i in range(0, 3)]:
            foo_endpoint_acl_policy.append(volcengine.cr.EndpointAclPolicy(f"fooEndpointAclPolicy-{range['value']}",
                registry=foo_endpoint.registry,
                type="Public",
                entry=f"192.168.0.{range['value']}",
                description=f"test-{range['value']}"))
        ```

        ## Import

        CrEndpointAclPolicy can be imported using the registry:entry, e.g.

        ```sh
        $ pulumi import volcengine:cr/endpointAclPolicy:EndpointAclPolicy default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the acl policy.
        :param pulumi.Input[str] entry: The ip list of the acl policy.
        :param pulumi.Input[str] registry: The registry name.
        :param pulumi.Input[str] type: The type of the acl policy. Valid values: `Public`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointAclPolicyInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage cr endpoint acl policy
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_registry = volcengine.cr.Registry("fooRegistry", project="default")
        foo_endpoint = volcengine.cr.Endpoint("fooEndpoint",
            registry=foo_registry.id,
            enabled=True)
        foo_endpoint_acl_policy = []
        for range in [{"value": i} for i in range(0, 3)]:
            foo_endpoint_acl_policy.append(volcengine.cr.EndpointAclPolicy(f"fooEndpointAclPolicy-{range['value']}",
                registry=foo_endpoint.registry,
                type="Public",
                entry=f"192.168.0.{range['value']}",
                description=f"test-{range['value']}"))
        ```

        ## Import

        CrEndpointAclPolicy can be imported using the registry:entry, e.g.

        ```sh
        $ pulumi import volcengine:cr/endpointAclPolicy:EndpointAclPolicy default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param EndpointAclPolicyInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointAclPolicyInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entry: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointAclPolicyInitArgs.__new__(EndpointAclPolicyInitArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if entry is None and not opts.urn:
                raise TypeError("Missing required property 'entry'")
            __props__.__dict__["entry"] = entry
            if registry is None and not opts.urn:
                raise TypeError("Missing required property 'registry'")
            __props__.__dict__["registry"] = registry
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(EndpointAclPolicy, __self__).__init__(
            'volcengine:cr/endpointAclPolicy:EndpointAclPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            entry: Optional[pulumi.Input[str]] = None,
            registry: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'EndpointAclPolicy':
        """
        Get an existing EndpointAclPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the acl policy.
        :param pulumi.Input[str] entry: The ip list of the acl policy.
        :param pulumi.Input[str] registry: The registry name.
        :param pulumi.Input[str] type: The type of the acl policy. Valid values: `Public`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointAclPolicyState.__new__(_EndpointAclPolicyState)

        __props__.__dict__["description"] = description
        __props__.__dict__["entry"] = entry
        __props__.__dict__["registry"] = registry
        __props__.__dict__["type"] = type
        return EndpointAclPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the acl policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def entry(self) -> pulumi.Output[str]:
        """
        The ip list of the acl policy.
        """
        return pulumi.get(self, "entry")

    @property
    @pulumi.getter
    def registry(self) -> pulumi.Output[str]:
        """
        The registry name.
        """
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the acl policy. Valid values: `Public`.
        """
        return pulumi.get(self, "type")

