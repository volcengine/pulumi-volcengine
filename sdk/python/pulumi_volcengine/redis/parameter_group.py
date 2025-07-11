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

__all__ = ['ParameterGroupArgs', 'ParameterGroup']

@pulumi.input_type
class ParameterGroupArgs:
    def __init__(__self__, *,
                 engine_version: pulumi.Input[str],
                 param_values: pulumi.Input[Sequence[pulumi.Input['ParameterGroupParamValueArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ParameterGroup resource.
        :param pulumi.Input[str] engine_version: The Redis database version adapted to the parameter template. The value range is as follows; 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.
        :param pulumi.Input[Sequence[pulumi.Input['ParameterGroupParamValueArgs']]] param_values: The list of parameter information that needs to be included in the new parameter template. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[str] description: The remarks information of the parameter template should not exceed 200 characters in length.
        :param pulumi.Input[str] name: Parameter template name. The name needs to meet the following requirements simultaneously: It cannot start with a number or a hyphen (-). Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included. The length should be 2 to 64 characters.
        """
        pulumi.set(__self__, "engine_version", engine_version)
        pulumi.set(__self__, "param_values", param_values)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Input[str]:
        """
        The Redis database version adapted to the parameter template. The value range is as follows; 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter(name="paramValues")
    def param_values(self) -> pulumi.Input[Sequence[pulumi.Input['ParameterGroupParamValueArgs']]]:
        """
        The list of parameter information that needs to be included in the new parameter template. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "param_values")

    @param_values.setter
    def param_values(self, value: pulumi.Input[Sequence[pulumi.Input['ParameterGroupParamValueArgs']]]):
        pulumi.set(self, "param_values", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The remarks information of the parameter template should not exceed 200 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Parameter template name. The name needs to meet the following requirements simultaneously: It cannot start with a number or a hyphen (-). Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included. The length should be 2 to 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ParameterGroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 param_values: Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParamValueArgs']]]] = None):
        """
        Input properties used for looking up and filtering ParameterGroup resources.
        :param pulumi.Input[str] description: The remarks information of the parameter template should not exceed 200 characters in length.
        :param pulumi.Input[str] engine_version: The Redis database version adapted to the parameter template. The value range is as follows; 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.
        :param pulumi.Input[str] name: Parameter template name. The name needs to meet the following requirements simultaneously: It cannot start with a number or a hyphen (-). Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included. The length should be 2 to 64 characters.
        :param pulumi.Input[Sequence[pulumi.Input['ParameterGroupParamValueArgs']]] param_values: The list of parameter information that needs to be included in the new parameter template. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if engine_version is not None:
            pulumi.set(__self__, "engine_version", engine_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if param_values is not None:
            pulumi.set(__self__, "param_values", param_values)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The remarks information of the parameter template should not exceed 200 characters in length.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Redis database version adapted to the parameter template. The value range is as follows; 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.
        """
        return pulumi.get(self, "engine_version")

    @engine_version.setter
    def engine_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine_version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Parameter template name. The name needs to meet the following requirements simultaneously: It cannot start with a number or a hyphen (-). Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included. The length should be 2 to 64 characters.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="paramValues")
    def param_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParamValueArgs']]]]:
        """
        The list of parameter information that needs to be included in the new parameter template. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "param_values")

    @param_values.setter
    def param_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ParameterGroupParamValueArgs']]]]):
        pulumi.set(self, "param_values", value)


class ParameterGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 param_values: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParamValueArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage redis parameter group
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.redis.ParameterGroup("foo",
            description="tf-test-description",
            engine_version="5.0",
            param_values=[
                volcengine.redis.ParameterGroupParamValueArgs(
                    name="active-defrag-cycle-max",
                    value="30",
                ),
                volcengine.redis.ParameterGroupParamValueArgs(
                    name="active-defrag-cycle-min",
                    value="15",
                ),
            ])
        ```

        ## Import

        ParameterGroup can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:redis/parameterGroup:ParameterGroup default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The remarks information of the parameter template should not exceed 200 characters in length.
        :param pulumi.Input[str] engine_version: The Redis database version adapted to the parameter template. The value range is as follows; 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.
        :param pulumi.Input[str] name: Parameter template name. The name needs to meet the following requirements simultaneously: It cannot start with a number or a hyphen (-). Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included. The length should be 2 to 64 characters.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParamValueArgs']]]] param_values: The list of parameter information that needs to be included in the new parameter template. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ParameterGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage redis parameter group
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.redis.ParameterGroup("foo",
            description="tf-test-description",
            engine_version="5.0",
            param_values=[
                volcengine.redis.ParameterGroupParamValueArgs(
                    name="active-defrag-cycle-max",
                    value="30",
                ),
                volcengine.redis.ParameterGroupParamValueArgs(
                    name="active-defrag-cycle-min",
                    value="15",
                ),
            ])
        ```

        ## Import

        ParameterGroup can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:redis/parameterGroup:ParameterGroup default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param ParameterGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ParameterGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 param_values: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParamValueArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ParameterGroupArgs.__new__(ParameterGroupArgs)

            __props__.__dict__["description"] = description
            if engine_version is None and not opts.urn:
                raise TypeError("Missing required property 'engine_version'")
            __props__.__dict__["engine_version"] = engine_version
            __props__.__dict__["name"] = name
            if param_values is None and not opts.urn:
                raise TypeError("Missing required property 'param_values'")
            __props__.__dict__["param_values"] = param_values
        super(ParameterGroup, __self__).__init__(
            'volcengine:redis/parameterGroup:ParameterGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            param_values: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParamValueArgs']]]]] = None) -> 'ParameterGroup':
        """
        Get an existing ParameterGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The remarks information of the parameter template should not exceed 200 characters in length.
        :param pulumi.Input[str] engine_version: The Redis database version adapted to the parameter template. The value range is as follows; 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.
        :param pulumi.Input[str] name: Parameter template name. The name needs to meet the following requirements simultaneously: It cannot start with a number or a hyphen (-). Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included. The length should be 2 to 64 characters.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ParameterGroupParamValueArgs']]]] param_values: The list of parameter information that needs to be included in the new parameter template. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ParameterGroupState.__new__(_ParameterGroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["engine_version"] = engine_version
        __props__.__dict__["name"] = name
        __props__.__dict__["param_values"] = param_values
        return ParameterGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The remarks information of the parameter template should not exceed 200 characters in length.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        """
        The Redis database version adapted to the parameter template. The value range is as follows; 7.0: Redis 7.0. 6.0: Redis 6.0. 5.0: Redis 5.0.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Parameter template name. The name needs to meet the following requirements simultaneously: It cannot start with a number or a hyphen (-). Only Chinese characters, letters, numbers, underscores (_) and hyphens (-) can be included. The length should be 2 to 64 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="paramValues")
    def param_values(self) -> pulumi.Output[Sequence['outputs.ParameterGroupParamValue']]:
        """
        The list of parameter information that needs to be included in the new parameter template. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "param_values")

