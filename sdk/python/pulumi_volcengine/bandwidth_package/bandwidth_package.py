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

__all__ = ['BandwidthPackageArgs', 'BandwidthPackage']

@pulumi.input_type
class BandwidthPackageArgs:
    def __init__(__self__, *,
                 bandwidth: pulumi.Input[int],
                 bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 security_protection_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]] = None):
        """
        The set of arguments for constructing a BandwidthPackage resource.
        :param pulumi.Input[int] bandwidth: Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
        :param pulumi.Input[str] bandwidth_package_name: The name of the bandwidth package.
        :param pulumi.Input[str] billing_type: BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
        :param pulumi.Input[str] description: The description of the bandwidth package.
        :param pulumi.Input[str] isp: Route type, default to BGP.
        :param pulumi.Input[int] period: Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
        :param pulumi.Input[str] project_name: The project name of the bandwidth package.
        :param pulumi.Input[str] protocol: The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_protection_types: Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
        :param pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]] tags: Tags.
        """
        pulumi.set(__self__, "bandwidth", bandwidth)
        if bandwidth_package_name is not None:
            pulumi.set(__self__, "bandwidth_package_name", bandwidth_package_name)
        if billing_type is not None:
            pulumi.set(__self__, "billing_type", billing_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if isp is not None:
            pulumi.set(__self__, "isp", isp)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if security_protection_types is not None:
            pulumi.set(__self__, "security_protection_types", security_protection_types)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Input[int]:
        """
        Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: pulumi.Input[int]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="bandwidthPackageName")
    def bandwidth_package_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bandwidth package.
        """
        return pulumi.get(self, "bandwidth_package_name")

    @bandwidth_package_name.setter
    def bandwidth_package_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_name", value)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> Optional[pulumi.Input[str]]:
        """
        BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
        """
        return pulumi.get(self, "billing_type")

    @billing_type.setter
    def billing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the bandwidth package.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def isp(self) -> Optional[pulumi.Input[str]]:
        """
        Route type, default to BGP.
        """
        return pulumi.get(self, "isp")

    @isp.setter
    def isp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isp", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of the bandwidth package.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="securityProtectionTypes")
    def security_protection_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
        """
        return pulumi.get(self, "security_protection_types")

    @security_protection_types.setter
    def security_protection_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_protection_types", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _BandwidthPackageState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 security_protection_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering BandwidthPackage resources.
        :param pulumi.Input[int] bandwidth: Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
        :param pulumi.Input[str] bandwidth_package_name: The name of the bandwidth package.
        :param pulumi.Input[str] billing_type: BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
        :param pulumi.Input[str] description: The description of the bandwidth package.
        :param pulumi.Input[str] isp: Route type, default to BGP.
        :param pulumi.Input[int] period: Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
        :param pulumi.Input[str] project_name: The project name of the bandwidth package.
        :param pulumi.Input[str] protocol: The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_protection_types: Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
        :param pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]] tags: Tags.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if bandwidth_package_name is not None:
            pulumi.set(__self__, "bandwidth_package_name", bandwidth_package_name)
        if billing_type is not None:
            pulumi.set(__self__, "billing_type", billing_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if isp is not None:
            pulumi.set(__self__, "isp", isp)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if security_protection_types is not None:
            pulumi.set(__self__, "security_protection_types", security_protection_types)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="bandwidthPackageName")
    def bandwidth_package_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bandwidth package.
        """
        return pulumi.get(self, "bandwidth_package_name")

    @bandwidth_package_name.setter
    def bandwidth_package_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bandwidth_package_name", value)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> Optional[pulumi.Input[str]]:
        """
        BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
        """
        return pulumi.get(self, "billing_type")

    @billing_type.setter
    def billing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the bandwidth package.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def isp(self) -> Optional[pulumi.Input[str]]:
        """
        Route type, default to BGP.
        """
        return pulumi.get(self, "isp")

    @isp.setter
    def isp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isp", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The project name of the bandwidth package.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="securityProtectionTypes")
    def security_protection_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
        """
        return pulumi.get(self, "security_protection_types")

    @security_protection_types.setter
    def security_protection_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_protection_types", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]]):
        pulumi.set(self, "tags", value)


class BandwidthPackage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 security_protection_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.bandwidth_package.BandwidthPackage("foo",
            bandwidth=10,
            bandwidth_package_name="acc-test-bp",
            billing_type="PostPaidByBandwidth",
            description="acc-test",
            isp="BGP",
            protocol="IPv4",
            security_protection_types=["AntiDDoS_Enhanced"],
            tags=[volcengine.bandwidth_package.BandwidthPackageTagArgs(
                key="k1",
                value="v1",
            )])
        ```

        ## Import

        BandwidthPackage can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:bandwidth_package/bandwidthPackage:BandwidthPackage default bwp-2zeo05qre24nhrqpy****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
        :param pulumi.Input[str] bandwidth_package_name: The name of the bandwidth package.
        :param pulumi.Input[str] billing_type: BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
        :param pulumi.Input[str] description: The description of the bandwidth package.
        :param pulumi.Input[str] isp: Route type, default to BGP.
        :param pulumi.Input[int] period: Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
        :param pulumi.Input[str] project_name: The project name of the bandwidth package.
        :param pulumi.Input[str] protocol: The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_protection_types: Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]] tags: Tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BandwidthPackageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.bandwidth_package.BandwidthPackage("foo",
            bandwidth=10,
            bandwidth_package_name="acc-test-bp",
            billing_type="PostPaidByBandwidth",
            description="acc-test",
            isp="BGP",
            protocol="IPv4",
            security_protection_types=["AntiDDoS_Enhanced"],
            tags=[volcengine.bandwidth_package.BandwidthPackageTagArgs(
                key="k1",
                value="v1",
            )])
        ```

        ## Import

        BandwidthPackage can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:bandwidth_package/bandwidthPackage:BandwidthPackage default bwp-2zeo05qre24nhrqpy****
        ```

        :param str resource_name: The name of the resource.
        :param BandwidthPackageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BandwidthPackageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 security_protection_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BandwidthPackageArgs.__new__(BandwidthPackageArgs)

            if bandwidth is None and not opts.urn:
                raise TypeError("Missing required property 'bandwidth'")
            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["bandwidth_package_name"] = bandwidth_package_name
            __props__.__dict__["billing_type"] = billing_type
            __props__.__dict__["description"] = description
            __props__.__dict__["isp"] = isp
            __props__.__dict__["period"] = period
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["protocol"] = protocol
            __props__.__dict__["security_protection_types"] = security_protection_types
            __props__.__dict__["tags"] = tags
        super(BandwidthPackage, __self__).__init__(
            'volcengine:bandwidth_package/bandwidthPackage:BandwidthPackage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            bandwidth_package_name: Optional[pulumi.Input[str]] = None,
            billing_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            isp: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            security_protection_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]]] = None) -> 'BandwidthPackage':
        """
        Get an existing BandwidthPackage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
        :param pulumi.Input[str] bandwidth_package_name: The name of the bandwidth package.
        :param pulumi.Input[str] billing_type: BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
        :param pulumi.Input[str] description: The description of the bandwidth package.
        :param pulumi.Input[str] isp: Route type, default to BGP.
        :param pulumi.Input[int] period: Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
        :param pulumi.Input[str] project_name: The project name of the bandwidth package.
        :param pulumi.Input[str] protocol: The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_protection_types: Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]] tags: Tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BandwidthPackageState.__new__(_BandwidthPackageState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["bandwidth_package_name"] = bandwidth_package_name
        __props__.__dict__["billing_type"] = billing_type
        __props__.__dict__["description"] = description
        __props__.__dict__["isp"] = isp
        __props__.__dict__["period"] = period
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["security_protection_types"] = security_protection_types
        __props__.__dict__["tags"] = tags
        return BandwidthPackage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[int]:
        """
        Bandwidth upper limit of shared bandwidth package, unit: Mbps. When BillingType is set to PrePaid: the value range is 5 to 5000. When BillingType is set to PostPaidByBandwidth: the value range is 2 to 5000. When BillingType is set to PostPaidByTraffic: the value range is 2 to 2000. When BillingType is set to PayBy95Peak: the value range is 2 to 5000.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="bandwidthPackageName")
    def bandwidth_package_name(self) -> pulumi.Output[str]:
        """
        The name of the bandwidth package.
        """
        return pulumi.get(self, "bandwidth_package_name")

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> pulumi.Output[Optional[str]]:
        """
        BillingType of the bandwidth package. Valid values: `PrePaid`, `PostPaidByBandwidth`(Default), `PostPaidByTraffic`, `PayBy95Peak`. The billing method of IPv6 does not include `PrePaid`, and the billing method is only based on the `PostPaidByBandwidth`.
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the bandwidth package.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def isp(self) -> pulumi.Output[Optional[str]]:
        """
        Route type, default to BGP.
        """
        return pulumi.get(self, "isp")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        """
        Duration of purchasing shared bandwidth package on an annual or monthly basis. The valid value range in 1~9 or 12, 24 or 36. Default value is 1. The period unit defaults to `Month`.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The project name of the bandwidth package.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The IP protocol values for shared bandwidth packages are as follows: `IPv4`: IPv4 protocol. `IPv6`: IPv6 protocol.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="securityProtectionTypes")
    def security_protection_types(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Security protection types for shared bandwidth packages. Parameter - N: Indicates the number of security protection types, currently only supports taking 1. Value: `AntiDDoS_Enhanced` or left blank.If the value is `AntiDDoS_Enhanced`, then will create a shared bandwidth package with enhanced protection, which supports adding basic protection type public IP addresses.If left blank, it indicates a shared bandwidth package with basic protection, which supports the addition of public IP addresses with enhanced protection.
        """
        return pulumi.get(self, "security_protection_types")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.BandwidthPackageTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")
