# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserGroupAttachmentArgs', 'UserGroupAttachment']

@pulumi.input_type
class UserGroupAttachmentArgs:
    def __init__(__self__, *,
                 user_group_name: pulumi.Input[str],
                 user_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a UserGroupAttachment resource.
        :param pulumi.Input[str] user_group_name: The name of the user group.
        :param pulumi.Input[str] user_name: The name of the user.
        """
        pulumi.set(__self__, "user_group_name", user_group_name)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="userGroupName")
    def user_group_name(self) -> pulumi.Input[str]:
        """
        The name of the user group.
        """
        return pulumi.get(self, "user_group_name")

    @user_group_name.setter
    def user_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_group_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Input[str]:
        """
        The name of the user.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_name", value)


@pulumi.input_type
class _UserGroupAttachmentState:
    def __init__(__self__, *,
                 user_group_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserGroupAttachment resources.
        :param pulumi.Input[str] user_group_name: The name of the user group.
        :param pulumi.Input[str] user_name: The name of the user.
        """
        if user_group_name is not None:
            pulumi.set(__self__, "user_group_name", user_group_name)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="userGroupName")
    def user_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user group.
        """
        return pulumi.get(self, "user_group_name")

    @user_group_name.setter
    def user_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_group_name", value)

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the user.
        """
        return pulumi.get(self, "user_name")

    @user_name.setter
    def user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_name", value)


class UserGroupAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 user_group_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage iam user group attachment
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_user = volcengine.iam.User("fooUser",
            user_name="acc-test-user",
            description="acc test",
            display_name="name")
        foo_user_group = volcengine.iam.UserGroup("fooUserGroup",
            user_group_name="acc-test-group",
            description="acc-test",
            display_name="acctest")
        foo_user_group_attachment = volcengine.iam.UserGroupAttachment("fooUserGroupAttachment",
            user_group_name=foo_user_group.user_group_name,
            user_name=foo_user.user_name)
        ```

        ## Import

        IamUserGroupAttachment can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:iam/userGroupAttachment:UserGroupAttachment default user_group_id:user_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] user_group_name: The name of the user group.
        :param pulumi.Input[str] user_name: The name of the user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserGroupAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage iam user group attachment
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_user = volcengine.iam.User("fooUser",
            user_name="acc-test-user",
            description="acc test",
            display_name="name")
        foo_user_group = volcengine.iam.UserGroup("fooUserGroup",
            user_group_name="acc-test-group",
            description="acc-test",
            display_name="acctest")
        foo_user_group_attachment = volcengine.iam.UserGroupAttachment("fooUserGroupAttachment",
            user_group_name=foo_user_group.user_group_name,
            user_name=foo_user.user_name)
        ```

        ## Import

        IamUserGroupAttachment can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:iam/userGroupAttachment:UserGroupAttachment default user_group_id:user_id
        ```

        :param str resource_name: The name of the resource.
        :param UserGroupAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserGroupAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 user_group_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserGroupAttachmentArgs.__new__(UserGroupAttachmentArgs)

            if user_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_group_name'")
            __props__.__dict__["user_group_name"] = user_group_name
            if user_name is None and not opts.urn:
                raise TypeError("Missing required property 'user_name'")
            __props__.__dict__["user_name"] = user_name
        super(UserGroupAttachment, __self__).__init__(
            'volcengine:iam/userGroupAttachment:UserGroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            user_group_name: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None) -> 'UserGroupAttachment':
        """
        Get an existing UserGroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] user_group_name: The name of the user group.
        :param pulumi.Input[str] user_name: The name of the user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserGroupAttachmentState.__new__(_UserGroupAttachmentState)

        __props__.__dict__["user_group_name"] = user_group_name
        __props__.__dict__["user_name"] = user_name
        return UserGroupAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="userGroupName")
    def user_group_name(self) -> pulumi.Output[str]:
        """
        The name of the user group.
        """
        return pulumi.get(self, "user_group_name")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[str]:
        """
        The name of the user.
        """
        return pulumi.get(self, "user_name")

