# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AclEntryArgs', 'AclEntry']

@pulumi.input_type
class AclEntryArgs:
    def __init__(__self__, *,
                 acl_id: pulumi.Input[str],
                 entry: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AclEntry resource.
        :param pulumi.Input[str] acl_id: The ID of Acl.
        :param pulumi.Input[str] entry: The content of the AclEntry.
        :param pulumi.Input[str] description: The description of the AclEntry.
        """
        pulumi.set(__self__, "acl_id", acl_id)
        pulumi.set(__self__, "entry", entry)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="aclId")
    def acl_id(self) -> pulumi.Input[str]:
        """
        The ID of Acl.
        """
        return pulumi.get(self, "acl_id")

    @acl_id.setter
    def acl_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "acl_id", value)

    @property
    @pulumi.getter
    def entry(self) -> pulumi.Input[str]:
        """
        The content of the AclEntry.
        """
        return pulumi.get(self, "entry")

    @entry.setter
    def entry(self, value: pulumi.Input[str]):
        pulumi.set(self, "entry", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the AclEntry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _AclEntryState:
    def __init__(__self__, *,
                 acl_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entry: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AclEntry resources.
        :param pulumi.Input[str] acl_id: The ID of Acl.
        :param pulumi.Input[str] description: The description of the AclEntry.
        :param pulumi.Input[str] entry: The content of the AclEntry.
        """
        if acl_id is not None:
            pulumi.set(__self__, "acl_id", acl_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if entry is not None:
            pulumi.set(__self__, "entry", entry)

    @property
    @pulumi.getter(name="aclId")
    def acl_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of Acl.
        """
        return pulumi.get(self, "acl_id")

    @acl_id.setter
    def acl_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "acl_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the AclEntry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def entry(self) -> Optional[pulumi.Input[str]]:
        """
        The content of the AclEntry.
        """
        return pulumi.get(self, "entry")

    @entry.setter
    def entry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entry", value)


class AclEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entry: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage acl entry
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_acl = volcengine.clb.Acl("fooAcl",
            acl_name="tf-test-3",
            description="tf-test")
        foo_acl_entry = volcengine.clb.AclEntry("fooAclEntry",
            acl_id=foo_acl.id,
            description="tf acl entry desc demo",
            entry="192.2.2.1/32")
        ```

        ## Import

        AclEntry can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:clb/aclEntry:AclEntry default ID is a string concatenated with colons(AclId:Entry)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl_id: The ID of Acl.
        :param pulumi.Input[str] description: The description of the AclEntry.
        :param pulumi.Input[str] entry: The content of the AclEntry.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AclEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage acl entry
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_acl = volcengine.clb.Acl("fooAcl",
            acl_name="tf-test-3",
            description="tf-test")
        foo_acl_entry = volcengine.clb.AclEntry("fooAclEntry",
            acl_id=foo_acl.id,
            description="tf acl entry desc demo",
            entry="192.2.2.1/32")
        ```

        ## Import

        AclEntry can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:clb/aclEntry:AclEntry default ID is a string concatenated with colons(AclId:Entry)
        ```

        :param str resource_name: The name of the resource.
        :param AclEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AclEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acl_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entry: Optional[pulumi.Input[str]] = None,
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
            __props__ = AclEntryArgs.__new__(AclEntryArgs)

            if acl_id is None and not opts.urn:
                raise TypeError("Missing required property 'acl_id'")
            __props__.__dict__["acl_id"] = acl_id
            __props__.__dict__["description"] = description
            if entry is None and not opts.urn:
                raise TypeError("Missing required property 'entry'")
            __props__.__dict__["entry"] = entry
        super(AclEntry, __self__).__init__(
            'volcengine:clb/aclEntry:AclEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acl_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            entry: Optional[pulumi.Input[str]] = None) -> 'AclEntry':
        """
        Get an existing AclEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl_id: The ID of Acl.
        :param pulumi.Input[str] description: The description of the AclEntry.
        :param pulumi.Input[str] entry: The content of the AclEntry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AclEntryState.__new__(_AclEntryState)

        __props__.__dict__["acl_id"] = acl_id
        __props__.__dict__["description"] = description
        __props__.__dict__["entry"] = entry
        return AclEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="aclId")
    def acl_id(self) -> pulumi.Output[str]:
        """
        The ID of Acl.
        """
        return pulumi.get(self, "acl_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the AclEntry.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def entry(self) -> pulumi.Output[str]:
        """
        The content of the AclEntry.
        """
        return pulumi.get(self, "entry")
