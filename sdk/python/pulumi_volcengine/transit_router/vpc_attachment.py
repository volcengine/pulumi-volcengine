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

__all__ = ['VpcAttachmentArgs', 'VpcAttachment']

@pulumi.input_type
class VpcAttachmentArgs:
    def __init__(__self__, *,
                 attach_points: pulumi.Input[Sequence[pulumi.Input['VpcAttachmentAttachPointArgs']]],
                 transit_router_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 transit_router_attachment_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcAttachment resource.
        :param pulumi.Input[Sequence[pulumi.Input['VpcAttachmentAttachPointArgs']]] attach_points: The attach points of transit router vpc attachment.
        :param pulumi.Input[str] transit_router_id: The id of the transit router.
        :param pulumi.Input[str] vpc_id: The ID of vpc.
        :param pulumi.Input[str] description: The description of the transit router vpc attachment.
        :param pulumi.Input[str] transit_router_attachment_name: The name of the transit router vpc attachment.
        """
        pulumi.set(__self__, "attach_points", attach_points)
        pulumi.set(__self__, "transit_router_id", transit_router_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if transit_router_attachment_name is not None:
            pulumi.set(__self__, "transit_router_attachment_name", transit_router_attachment_name)

    @property
    @pulumi.getter(name="attachPoints")
    def attach_points(self) -> pulumi.Input[Sequence[pulumi.Input['VpcAttachmentAttachPointArgs']]]:
        """
        The attach points of transit router vpc attachment.
        """
        return pulumi.get(self, "attach_points")

    @attach_points.setter
    def attach_points(self, value: pulumi.Input[Sequence[pulumi.Input['VpcAttachmentAttachPointArgs']]]):
        pulumi.set(self, "attach_points", value)

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
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of vpc.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the transit router vpc attachment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="transitRouterAttachmentName")
    def transit_router_attachment_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the transit router vpc attachment.
        """
        return pulumi.get(self, "transit_router_attachment_name")

    @transit_router_attachment_name.setter
    def transit_router_attachment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_attachment_name", value)


@pulumi.input_type
class _VpcAttachmentState:
    def __init__(__self__, *,
                 attach_points: Optional[pulumi.Input[Sequence[pulumi.Input['VpcAttachmentAttachPointArgs']]]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
                 transit_router_attachment_name: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcAttachment resources.
        :param pulumi.Input[Sequence[pulumi.Input['VpcAttachmentAttachPointArgs']]] attach_points: The attach points of transit router vpc attachment.
        :param pulumi.Input[str] creation_time: The create time.
        :param pulumi.Input[str] description: The description of the transit router vpc attachment.
        :param pulumi.Input[str] status: The status of the transit router.
        :param pulumi.Input[str] transit_router_attachment_id: The id of the transit router attachment.
        :param pulumi.Input[str] transit_router_attachment_name: The name of the transit router vpc attachment.
        :param pulumi.Input[str] transit_router_id: The id of the transit router.
        :param pulumi.Input[str] update_time: The update time.
        :param pulumi.Input[str] vpc_id: The ID of vpc.
        """
        if attach_points is not None:
            pulumi.set(__self__, "attach_points", attach_points)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if transit_router_attachment_id is not None:
            pulumi.set(__self__, "transit_router_attachment_id", transit_router_attachment_id)
        if transit_router_attachment_name is not None:
            pulumi.set(__self__, "transit_router_attachment_name", transit_router_attachment_name)
        if transit_router_id is not None:
            pulumi.set(__self__, "transit_router_id", transit_router_id)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="attachPoints")
    def attach_points(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcAttachmentAttachPointArgs']]]]:
        """
        The attach points of transit router vpc attachment.
        """
        return pulumi.get(self, "attach_points")

    @attach_points.setter
    def attach_points(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcAttachmentAttachPointArgs']]]]):
        pulumi.set(self, "attach_points", value)

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
        The description of the transit router vpc attachment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the transit router attachment.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @transit_router_attachment_id.setter
    def transit_router_attachment_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_attachment_id", value)

    @property
    @pulumi.getter(name="transitRouterAttachmentName")
    def transit_router_attachment_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the transit router vpc attachment.
        """
        return pulumi.get(self, "transit_router_attachment_name")

    @transit_router_attachment_name.setter
    def transit_router_attachment_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_router_attachment_name", value)

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

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of vpc.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class VpcAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attach_points: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcAttachmentAttachPointArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 transit_router_attachment_name: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage transit router vpc attachment
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.transit_router.VpcAttachment("foo",
            attach_points=[
                volcengine.transit_router.VpcAttachmentAttachPointArgs(
                    subnet_id="subnet-3refsrxdswsn45zsk2hmdg4zx",
                    zone_id="cn-beijing-a",
                ),
                volcengine.transit_router.VpcAttachmentAttachPointArgs(
                    subnet_id="subnet-2d68bh74345q858ozfekrm8fj",
                    zone_id="cn-beijing-a",
                ),
            ],
            description="desc",
            transit_router_attachment_name="tfname1",
            transit_router_id="tr-2d6fr7f39unsw58ozfe1ow21x",
            vpc_id="vpc-2bysvq1xx543k2dx0eeulpeiv")
        ```

        ## Import

        TransitRouterVpcAttachment can be imported using the transitRouterId:attachmentId, e.g.

        ```sh
         $ pulumi import volcengine:transit_router/vpcAttachment:VpcAttachment default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcAttachmentAttachPointArgs']]]] attach_points: The attach points of transit router vpc attachment.
        :param pulumi.Input[str] description: The description of the transit router vpc attachment.
        :param pulumi.Input[str] transit_router_attachment_name: The name of the transit router vpc attachment.
        :param pulumi.Input[str] transit_router_id: The id of the transit router.
        :param pulumi.Input[str] vpc_id: The ID of vpc.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage transit router vpc attachment
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.transit_router.VpcAttachment("foo",
            attach_points=[
                volcengine.transit_router.VpcAttachmentAttachPointArgs(
                    subnet_id="subnet-3refsrxdswsn45zsk2hmdg4zx",
                    zone_id="cn-beijing-a",
                ),
                volcengine.transit_router.VpcAttachmentAttachPointArgs(
                    subnet_id="subnet-2d68bh74345q858ozfekrm8fj",
                    zone_id="cn-beijing-a",
                ),
            ],
            description="desc",
            transit_router_attachment_name="tfname1",
            transit_router_id="tr-2d6fr7f39unsw58ozfe1ow21x",
            vpc_id="vpc-2bysvq1xx543k2dx0eeulpeiv")
        ```

        ## Import

        TransitRouterVpcAttachment can be imported using the transitRouterId:attachmentId, e.g.

        ```sh
         $ pulumi import volcengine:transit_router/vpcAttachment:VpcAttachment default tr-2d6fr7mzya2gw58ozfes5g2oh:tr-attach-7qthudw0ll6jmc****
        ```

        :param str resource_name: The name of the resource.
        :param VpcAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attach_points: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcAttachmentAttachPointArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 transit_router_attachment_name: Optional[pulumi.Input[str]] = None,
                 transit_router_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcAttachmentArgs.__new__(VpcAttachmentArgs)

            if attach_points is None and not opts.urn:
                raise TypeError("Missing required property 'attach_points'")
            __props__.__dict__["attach_points"] = attach_points
            __props__.__dict__["description"] = description
            __props__.__dict__["transit_router_attachment_name"] = transit_router_attachment_name
            if transit_router_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_router_id'")
            __props__.__dict__["transit_router_id"] = transit_router_id
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["transit_router_attachment_id"] = None
            __props__.__dict__["update_time"] = None
        super(VpcAttachment, __self__).__init__(
            'volcengine:transit_router/vpcAttachment:VpcAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attach_points: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcAttachmentAttachPointArgs']]]]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            transit_router_attachment_id: Optional[pulumi.Input[str]] = None,
            transit_router_attachment_name: Optional[pulumi.Input[str]] = None,
            transit_router_id: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'VpcAttachment':
        """
        Get an existing VpcAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcAttachmentAttachPointArgs']]]] attach_points: The attach points of transit router vpc attachment.
        :param pulumi.Input[str] creation_time: The create time.
        :param pulumi.Input[str] description: The description of the transit router vpc attachment.
        :param pulumi.Input[str] status: The status of the transit router.
        :param pulumi.Input[str] transit_router_attachment_id: The id of the transit router attachment.
        :param pulumi.Input[str] transit_router_attachment_name: The name of the transit router vpc attachment.
        :param pulumi.Input[str] transit_router_id: The id of the transit router.
        :param pulumi.Input[str] update_time: The update time.
        :param pulumi.Input[str] vpc_id: The ID of vpc.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcAttachmentState.__new__(_VpcAttachmentState)

        __props__.__dict__["attach_points"] = attach_points
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["status"] = status
        __props__.__dict__["transit_router_attachment_id"] = transit_router_attachment_id
        __props__.__dict__["transit_router_attachment_name"] = transit_router_attachment_name
        __props__.__dict__["transit_router_id"] = transit_router_id
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["vpc_id"] = vpc_id
        return VpcAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attachPoints")
    def attach_points(self) -> pulumi.Output[Sequence['outputs.VpcAttachmentAttachPoint']]:
        """
        The attach points of transit router vpc attachment.
        """
        return pulumi.get(self, "attach_points")

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
        The description of the transit router vpc attachment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the transit router.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="transitRouterAttachmentId")
    def transit_router_attachment_id(self) -> pulumi.Output[str]:
        """
        The id of the transit router attachment.
        """
        return pulumi.get(self, "transit_router_attachment_id")

    @property
    @pulumi.getter(name="transitRouterAttachmentName")
    def transit_router_attachment_name(self) -> pulumi.Output[str]:
        """
        The name of the transit router vpc attachment.
        """
        return pulumi.get(self, "transit_router_attachment_name")

    @property
    @pulumi.getter(name="transitRouterId")
    def transit_router_id(self) -> pulumi.Output[str]:
        """
        The id of the transit router.
        """
        return pulumi.get(self, "transit_router_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of vpc.
        """
        return pulumi.get(self, "vpc_id")

