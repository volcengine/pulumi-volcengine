# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServiceControlPolicyEnablerArgs', 'ServiceControlPolicyEnabler']

@pulumi.input_type
class ServiceControlPolicyEnablerArgs:
    def __init__(__self__):
        """
        The set of arguments for constructing a ServiceControlPolicyEnabler resource.
        """
        pass


class ServiceControlPolicyEnabler(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 __props__=None):
        """
        Provides a resource to manage organization service control policy enabler
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.organization.ServiceControlPolicyEnabler("foo")
        ```

        ## Import

        ServiceControlPolicy enabler can be imported using the default_id (organization:service_control_policy_enable) , e.g.

        ```sh
        $ pulumi import volcengine:organization/serviceControlPolicyEnabler:ServiceControlPolicyEnabler default organization:service_control_policy_enable
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ServiceControlPolicyEnablerArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage organization service control policy enabler
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.organization.ServiceControlPolicyEnabler("foo")
        ```

        ## Import

        ServiceControlPolicy enabler can be imported using the default_id (organization:service_control_policy_enable) , e.g.

        ```sh
        $ pulumi import volcengine:organization/serviceControlPolicyEnabler:ServiceControlPolicyEnabler default organization:service_control_policy_enable
        ```

        :param str resource_name: The name of the resource.
        :param ServiceControlPolicyEnablerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceControlPolicyEnablerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceControlPolicyEnablerArgs.__new__(ServiceControlPolicyEnablerArgs)

        super(ServiceControlPolicyEnabler, __self__).__init__(
            'volcengine:organization/serviceControlPolicyEnabler:ServiceControlPolicyEnabler',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceControlPolicyEnabler':
        """
        Get an existing ServiceControlPolicyEnabler resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceControlPolicyEnablerArgs.__new__(ServiceControlPolicyEnablerArgs)

        return ServiceControlPolicyEnabler(resource_name, opts=opts, __props__=__props__)

