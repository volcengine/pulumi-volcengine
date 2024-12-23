# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcArgs', 'Vpc']

@pulumi.input_type
class VpcArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 desc: pulumi.Input[str],
                 vpc_name: pulumi.Input[str],
                 cidr: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Vpc resource.
        :param pulumi.Input[str] cluster_name: The name of the cluster.
        :param pulumi.Input[str] desc: The description of the VPC.
        :param pulumi.Input[str] vpc_name: The name of the VPC.
        :param pulumi.Input[str] cidr: The cidr info.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "desc", desc)
        pulumi.set(__self__, "vpc_name", vpc_name)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def desc(self) -> pulumi.Input[str]:
        """
        The description of the VPC.
        """
        return pulumi.get(self, "desc")

    @desc.setter
    def desc(self, value: pulumi.Input[str]):
        pulumi.set(self, "desc", value)

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> pulumi.Input[str]:
        """
        The name of the VPC.
        """
        return pulumi.get(self, "vpc_name")

    @vpc_name.setter
    def vpc_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_name", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The cidr info.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)


@pulumi.input_type
class _VpcState:
    def __init__(__self__, *,
                 cidr: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 vpc_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Vpc resources.
        :param pulumi.Input[str] cidr: The cidr info.
        :param pulumi.Input[str] cluster_name: The name of the cluster.
        :param pulumi.Input[str] desc: The description of the VPC.
        :param pulumi.Input[str] vpc_name: The name of the VPC.
        """
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if desc is not None:
            pulumi.set(__self__, "desc", desc)
        if vpc_name is not None:
            pulumi.set(__self__, "vpc_name", vpc_name)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The cidr info.
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the VPC.
        """
        return pulumi.get(self, "desc")

    @desc.setter
    def desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desc", value)

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPC.
        """
        return pulumi.get(self, "vpc_name")

    @vpc_name.setter
    def vpc_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_name", value)


class Vpc(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 vpc_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage veenedge vpc
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.veenedge.Vpc("foo",
            cluster_name="b****t02",
            vpc_name="tf-test-2")
        ```

        ## Import

        VPC can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:veenedge/vpc:Vpc default vpc-mizl7m1k
        ```
        If you need to create a VPC, you need to apply for permission from the administrator in advance.
        You can only delete the vpc from web consul

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: The cidr info.
        :param pulumi.Input[str] cluster_name: The name of the cluster.
        :param pulumi.Input[str] desc: The description of the VPC.
        :param pulumi.Input[str] vpc_name: The name of the VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage veenedge vpc
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.veenedge.Vpc("foo",
            cluster_name="b****t02",
            vpc_name="tf-test-2")
        ```

        ## Import

        VPC can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:veenedge/vpc:Vpc default vpc-mizl7m1k
        ```
        If you need to create a VPC, you need to apply for permission from the administrator in advance.
        You can only delete the vpc from web consul

        :param str resource_name: The name of the resource.
        :param VpcArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 vpc_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcArgs.__new__(VpcArgs)

            __props__.__dict__["cidr"] = cidr
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            if desc is None and not opts.urn:
                raise TypeError("Missing required property 'desc'")
            __props__.__dict__["desc"] = desc
            if vpc_name is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_name'")
            __props__.__dict__["vpc_name"] = vpc_name
        super(Vpc, __self__).__init__(
            'volcengine:veenedge/vpc:Vpc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            desc: Optional[pulumi.Input[str]] = None,
            vpc_name: Optional[pulumi.Input[str]] = None) -> 'Vpc':
        """
        Get an existing Vpc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: The cidr info.
        :param pulumi.Input[str] cluster_name: The name of the cluster.
        :param pulumi.Input[str] desc: The description of the VPC.
        :param pulumi.Input[str] vpc_name: The name of the VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcState.__new__(_VpcState)

        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["desc"] = desc
        __props__.__dict__["vpc_name"] = vpc_name
        return Vpc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[Optional[str]]:
        """
        The cidr info.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter
    def desc(self) -> pulumi.Output[str]:
        """
        The description of the VPC.
        """
        return pulumi.get(self, "desc")

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> pulumi.Output[str]:
        """
        The name of the VPC.
        """
        return pulumi.get(self, "vpc_name")

