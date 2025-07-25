# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['KeyRotationArgs', 'KeyRotation']

@pulumi.input_type
class KeyRotationArgs:
    def __init__(__self__, *,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 keyring_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a KeyRotation resource.
        :param pulumi.Input[str] key_id: The id of the CMK.
        :param pulumi.Input[str] key_name: The name of the CMK.
        :param pulumi.Input[str] keyring_name: The name of the keyring.
        """
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if keyring_name is not None:
            pulumi.set(__self__, "keyring_name", keyring_name)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the CMK.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the CMK.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyringName")
    def keyring_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the keyring.
        """
        return pulumi.get(self, "keyring_name")

    @keyring_name.setter
    def keyring_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "keyring_name", value)


@pulumi.input_type
class _KeyRotationState:
    def __init__(__self__, *,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 keyring_name: Optional[pulumi.Input[str]] = None,
                 rotation_state: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering KeyRotation resources.
        :param pulumi.Input[str] key_id: The id of the CMK.
        :param pulumi.Input[str] key_name: The name of the CMK.
        :param pulumi.Input[str] keyring_name: The name of the keyring.
        :param pulumi.Input[str] rotation_state: The state of the key rotation.
        """
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if keyring_name is not None:
            pulumi.set(__self__, "keyring_name", keyring_name)
        if rotation_state is not None:
            pulumi.set(__self__, "rotation_state", rotation_state)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the CMK.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the CMK.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyringName")
    def keyring_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the keyring.
        """
        return pulumi.get(self, "keyring_name")

    @keyring_name.setter
    def keyring_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "keyring_name", value)

    @property
    @pulumi.getter(name="rotationState")
    def rotation_state(self) -> Optional[pulumi.Input[str]]:
        """
        The state of the key rotation.
        """
        return pulumi.get(self, "rotation_state")

    @rotation_state.setter
    def rotation_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rotation_state", value)


class KeyRotation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 keyring_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage kms key rotation
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.kms.KeyRotation("foo", key_id="m_cn-guilin-boe_63c08fe9-42e8-4c10-a09e-8e8e6xxxxxx")
        ```

        ## Import

        KmsKeyRotation can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:kms/keyRotation:KeyRotation default resource_id
        ```

        or

        ```sh
        $ pulumi import volcengine:kms/keyRotation:KeyRotation default key_name:keyring_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_id: The id of the CMK.
        :param pulumi.Input[str] key_name: The name of the CMK.
        :param pulumi.Input[str] keyring_name: The name of the keyring.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[KeyRotationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage kms key rotation
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.kms.KeyRotation("foo", key_id="m_cn-guilin-boe_63c08fe9-42e8-4c10-a09e-8e8e6xxxxxx")
        ```

        ## Import

        KmsKeyRotation can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:kms/keyRotation:KeyRotation default resource_id
        ```

        or

        ```sh
        $ pulumi import volcengine:kms/keyRotation:KeyRotation default key_name:keyring_name
        ```

        :param str resource_name: The name of the resource.
        :param KeyRotationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyRotationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 keyring_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeyRotationArgs.__new__(KeyRotationArgs)

            __props__.__dict__["key_id"] = key_id
            __props__.__dict__["key_name"] = key_name
            __props__.__dict__["keyring_name"] = keyring_name
            __props__.__dict__["rotation_state"] = None
        super(KeyRotation, __self__).__init__(
            'volcengine:kms/keyRotation:KeyRotation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            key_name: Optional[pulumi.Input[str]] = None,
            keyring_name: Optional[pulumi.Input[str]] = None,
            rotation_state: Optional[pulumi.Input[str]] = None) -> 'KeyRotation':
        """
        Get an existing KeyRotation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_id: The id of the CMK.
        :param pulumi.Input[str] key_name: The name of the CMK.
        :param pulumi.Input[str] keyring_name: The name of the keyring.
        :param pulumi.Input[str] rotation_state: The state of the key rotation.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _KeyRotationState.__new__(_KeyRotationState)

        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["key_name"] = key_name
        __props__.__dict__["keyring_name"] = keyring_name
        __props__.__dict__["rotation_state"] = rotation_state
        return KeyRotation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        The id of the CMK.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[str]:
        """
        The name of the CMK.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="keyringName")
    def keyring_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the keyring.
        """
        return pulumi.get(self, "keyring_name")

    @property
    @pulumi.getter(name="rotationState")
    def rotation_state(self) -> pulumi.Output[str]:
        """
        The state of the key rotation.
        """
        return pulumi.get(self, "rotation_state")

