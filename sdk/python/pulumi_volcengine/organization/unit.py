# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UnitArgs', 'Unit']

@pulumi.input_type
class UnitArgs:
    def __init__(__self__, *,
                 parent_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Unit resource.
        :param pulumi.Input[str] parent_id: Parent Organization Unit ID.
        :param pulumi.Input[str] description: Description of the organization unit.
        :param pulumi.Input[str] name: Name of the organization unit.
        """
        pulumi.set(__self__, "parent_id", parent_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Input[str]:
        """
        Parent Organization Unit ID.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the organization unit.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the organization unit.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _UnitState:
    def __init__(__self__, *,
                 depth: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 org_type: Optional[pulumi.Input[int]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Unit resources.
        :param pulumi.Input[int] depth: The depth of the organization unit.
        :param pulumi.Input[str] description: Description of the organization unit.
        :param pulumi.Input[str] name: Name of the organization unit.
        :param pulumi.Input[str] org_id: The id of the organization.
        :param pulumi.Input[int] org_type: The organization type.
        :param pulumi.Input[str] owner: The owner of the organization unit.
        :param pulumi.Input[str] parent_id: Parent Organization Unit ID.
        """
        if depth is not None:
            pulumi.set(__self__, "depth", depth)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if org_type is not None:
            pulumi.set(__self__, "org_type", org_type)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if parent_id is not None:
            pulumi.set(__self__, "parent_id", parent_id)

    @property
    @pulumi.getter
    def depth(self) -> Optional[pulumi.Input[int]]:
        """
        The depth of the organization unit.
        """
        return pulumi.get(self, "depth")

    @depth.setter
    def depth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "depth", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the organization unit.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the organization unit.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the organization.
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="orgType")
    def org_type(self) -> Optional[pulumi.Input[int]]:
        """
        The organization type.
        """
        return pulumi.get(self, "org_type")

    @org_type.setter
    def org_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "org_type", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of the organization unit.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[str]]:
        """
        Parent Organization Unit ID.
        """
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_id", value)


class Unit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage organization unit
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_organization = volcengine.organization.Organization("fooOrganization")
        foo_units = volcengine.organization.units()
        foo_unit = volcengine.organization.Unit("fooUnit",
            parent_id=[unit.id for unit in foo_units.units if unit.parent_id == "0"][0],
            description="tf-test")
        ```

        ## Import

        OrganizationUnit can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:organization/unit:Unit default ID
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the organization unit.
        :param pulumi.Input[str] name: Name of the organization unit.
        :param pulumi.Input[str] parent_id: Parent Organization Unit ID.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UnitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage organization unit
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_organization = volcengine.organization.Organization("fooOrganization")
        foo_units = volcengine.organization.units()
        foo_unit = volcengine.organization.Unit("fooUnit",
            parent_id=[unit.id for unit in foo_units.units if unit.parent_id == "0"][0],
            description="tf-test")
        ```

        ## Import

        OrganizationUnit can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:organization/unit:Unit default ID
        ```

        :param str resource_name: The name of the resource.
        :param UnitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UnitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UnitArgs.__new__(UnitArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if parent_id is None and not opts.urn:
                raise TypeError("Missing required property 'parent_id'")
            __props__.__dict__["parent_id"] = parent_id
            __props__.__dict__["depth"] = None
            __props__.__dict__["org_id"] = None
            __props__.__dict__["org_type"] = None
            __props__.__dict__["owner"] = None
        super(Unit, __self__).__init__(
            'volcengine:organization/unit:Unit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            depth: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            org_type: Optional[pulumi.Input[int]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            parent_id: Optional[pulumi.Input[str]] = None) -> 'Unit':
        """
        Get an existing Unit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] depth: The depth of the organization unit.
        :param pulumi.Input[str] description: Description of the organization unit.
        :param pulumi.Input[str] name: Name of the organization unit.
        :param pulumi.Input[str] org_id: The id of the organization.
        :param pulumi.Input[int] org_type: The organization type.
        :param pulumi.Input[str] owner: The owner of the organization unit.
        :param pulumi.Input[str] parent_id: Parent Organization Unit ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UnitState.__new__(_UnitState)

        __props__.__dict__["depth"] = depth
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["org_type"] = org_type
        __props__.__dict__["owner"] = owner
        __props__.__dict__["parent_id"] = parent_id
        return Unit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def depth(self) -> pulumi.Output[int]:
        """
        The depth of the organization unit.
        """
        return pulumi.get(self, "depth")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the organization unit.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the organization unit.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        The id of the organization.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="orgType")
    def org_type(self) -> pulumi.Output[int]:
        """
        The organization type.
        """
        return pulumi.get(self, "org_type")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        The owner of the organization unit.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[str]:
        """
        Parent Organization Unit ID.
        """
        return pulumi.get(self, "parent_id")
