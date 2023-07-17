# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'AuthorizationTokensTokenResult',
    'EndpointsEndpointResult',
    'NamespacesNamespaceResult',
    'RegistriesRegistryResult',
    'RegistriesRegistryDomainResult',
    'RegistriesRegistryStatusResult',
    'RegistriesStatusResult',
    'RegistryDomain',
    'RegistryStatus',
    'RepositoriesRepositoryResult',
    'StateStatus',
    'TagChartAttribute',
    'TagImageAttribute',
    'TagsTagResult',
    'TagsTagChartAttributeResult',
    'TagsTagImageAttributeResult',
    'VpcEndpointVpc',
    'VpcEndpointsEndpointResult',
    'VpcEndpointsEndpointVpcResult',
]

@pulumi.output_type
class AuthorizationTokensTokenResult(dict):
    def __init__(__self__, *,
                 expire_time: str,
                 token: str,
                 username: str):
        """
        :param str expire_time: The expiration time of the temporary access token.
        :param str token: The Temporary access token.
        :param str username: The username for login repository instance.
        """
        pulumi.set(__self__, "expire_time", expire_time)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The expiration time of the temporary access token.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def token(self) -> str:
        """
        The Temporary access token.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The username for login repository instance.
        """
        return pulumi.get(self, "username")


@pulumi.output_type
class EndpointsEndpointResult(dict):
    def __init__(__self__, *,
                 enabled: bool,
                 registry: str,
                 status: str):
        """
        :param bool enabled: Whether public endpoint is enabled.
        :param str registry: The CR instance name.
        :param str status: The status of public endpoint.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "registry", registry)
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether public endpoint is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def registry(self) -> str:
        """
        The CR instance name.
        """
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of public endpoint.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class NamespacesNamespaceResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 name: str):
        """
        :param str create_time: The time when namespace created.
        :param str name: The name of OCI repository.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when namespace created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of OCI repository.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class RegistriesRegistryResult(dict):
    def __init__(__self__, *,
                 charge_type: str,
                 create_time: str,
                 domains: Sequence['outputs.RegistriesRegistryDomainResult'],
                 name: str,
                 status: 'outputs.RegistriesRegistryStatusResult',
                 type: str,
                 user_status: str,
                 username: str):
        """
        :param str charge_type: The charge type of registry.
        :param str create_time: The creation time of registry.
        :param Sequence['RegistriesRegistryDomainArgs'] domains: The domain of registry.
        :param str name: The name of registry.
        :param 'RegistriesRegistryStatusArgs' status: The status of registry.
        :param str type: The type of registry.
        :param str user_status: The status of user.
        :param str username: The username of cr instance.
        """
        pulumi.set(__self__, "charge_type", charge_type)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "domains", domains)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "user_status", user_status)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="chargeType")
    def charge_type(self) -> str:
        """
        The charge type of registry.
        """
        return pulumi.get(self, "charge_type")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of registry.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def domains(self) -> Sequence['outputs.RegistriesRegistryDomainResult']:
        """
        The domain of registry.
        """
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of registry.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.RegistriesRegistryStatusResult':
        """
        The status of registry.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of registry.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userStatus")
    def user_status(self) -> str:
        """
        The status of user.
        """
        return pulumi.get(self, "user_status")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The username of cr instance.
        """
        return pulumi.get(self, "username")


@pulumi.output_type
class RegistriesRegistryDomainResult(dict):
    def __init__(__self__, *,
                 domain: str,
                 type: str):
        """
        :param str domain: The domain of registry.
        :param str type: The type of registry.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        The domain of registry.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of registry.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class RegistriesRegistryStatusResult(dict):
    def __init__(__self__, *,
                 conditions: Sequence[str],
                 phase: str):
        """
        :param Sequence[str] conditions: The condition of registry.
        :param str phase: The phase of status.
        """
        pulumi.set(__self__, "conditions", conditions)
        pulumi.set(__self__, "phase", phase)

    @property
    @pulumi.getter
    def conditions(self) -> Sequence[str]:
        """
        The condition of registry.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def phase(self) -> str:
        """
        The phase of status.
        """
        return pulumi.get(self, "phase")


@pulumi.output_type
class RegistriesStatusResult(dict):
    def __init__(__self__, *,
                 condition: Optional[str] = None,
                 phase: Optional[str] = None):
        """
        :param str condition: The condition of registry.
        :param str phase: The phase of status.
        """
        if condition is not None:
            pulumi.set(__self__, "condition", condition)
        if phase is not None:
            pulumi.set(__self__, "phase", phase)

    @property
    @pulumi.getter
    def condition(self) -> Optional[str]:
        """
        The condition of registry.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def phase(self) -> Optional[str]:
        """
        The phase of status.
        """
        return pulumi.get(self, "phase")


@pulumi.output_type
class RegistryDomain(dict):
    def __init__(__self__, *,
                 domain: Optional[str] = None,
                 type: Optional[str] = None):
        """
        :param str domain: The domain of registry.
        :param str type: The type of registry.
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        The domain of registry.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of registry.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class RegistryStatus(dict):
    def __init__(__self__, *,
                 conditions: Optional[Sequence[str]] = None,
                 phase: Optional[str] = None):
        """
        :param Sequence[str] conditions: The condition of registry.
        :param str phase: The phase status of registry.
        """
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if phase is not None:
            pulumi.set(__self__, "phase", phase)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[Sequence[str]]:
        """
        The condition of registry.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def phase(self) -> Optional[str]:
        """
        The phase status of registry.
        """
        return pulumi.get(self, "phase")


@pulumi.output_type
class RepositoriesRepositoryResult(dict):
    def __init__(__self__, *,
                 access_level: str,
                 create_time: str,
                 description: str,
                 name: str,
                 namespace: str,
                 update_time: str):
        """
        :param str access_level: The access level of repository.
        :param str create_time: The creation time of repository.
        :param str description: The description of repository.
        :param str name: The name of repository.
        :param str namespace: The namespace of repository.
        :param str update_time: The last update time of repository.
        """
        pulumi.set(__self__, "access_level", access_level)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> str:
        """
        The access level of repository.
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of repository.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of repository.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of repository.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        The namespace of repository.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update time of repository.
        """
        return pulumi.get(self, "update_time")


@pulumi.output_type
class StateStatus(dict):
    def __init__(__self__, *,
                 conditions: Optional[Sequence[str]] = None,
                 phase: Optional[str] = None):
        """
        :param Sequence[str] conditions: The condition of instance.
        :param str phase: The phase status of instance.
        """
        if conditions is not None:
            pulumi.set(__self__, "conditions", conditions)
        if phase is not None:
            pulumi.set(__self__, "phase", phase)

    @property
    @pulumi.getter
    def conditions(self) -> Optional[Sequence[str]]:
        """
        The condition of instance.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter
    def phase(self) -> Optional[str]:
        """
        The phase status of instance.
        """
        return pulumi.get(self, "phase")


@pulumi.output_type
class TagChartAttribute(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "apiVersion":
            suggest = "api_version"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in TagChartAttribute. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        TagChartAttribute.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        TagChartAttribute.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 api_version: Optional[str] = None,
                 name: Optional[str] = None,
                 version: Optional[str] = None):
        """
        :param str api_version: The Helm version.
        :param str name: The name of OCI product.
        :param str version: The Helm Chart version.
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", api_version)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[str]:
        """
        The Helm version.
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of OCI product.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        The Helm Chart version.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class TagImageAttribute(dict):
    def __init__(__self__, *,
                 architecture: Optional[str] = None,
                 author: Optional[str] = None,
                 digest: Optional[str] = None,
                 os: Optional[str] = None):
        """
        :param str architecture: The image architecture.
        :param str author: The image author.
        :param str digest: The digest of image.
        :param str os: The iamge os.
        """
        if architecture is not None:
            pulumi.set(__self__, "architecture", architecture)
        if author is not None:
            pulumi.set(__self__, "author", author)
        if digest is not None:
            pulumi.set(__self__, "digest", digest)
        if os is not None:
            pulumi.set(__self__, "os", os)

    @property
    @pulumi.getter
    def architecture(self) -> Optional[str]:
        """
        The image architecture.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter
    def author(self) -> Optional[str]:
        """
        The image author.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter
    def digest(self) -> Optional[str]:
        """
        The digest of image.
        """
        return pulumi.get(self, "digest")

    @property
    @pulumi.getter
    def os(self) -> Optional[str]:
        """
        The iamge os.
        """
        return pulumi.get(self, "os")


@pulumi.output_type
class TagsTagResult(dict):
    def __init__(__self__, *,
                 chart_attribute: 'outputs.TagsTagChartAttributeResult',
                 digest: str,
                 image_attributes: Sequence['outputs.TagsTagImageAttributeResult'],
                 name: str,
                 push_time: str,
                 size: int,
                 type: str):
        """
        :param 'TagsTagChartAttributeArgs' chart_attribute: The chart attribute,valid when tag type is Chart.
        :param str digest: The digest of image.
        :param Sequence['TagsTagImageAttributeArgs'] image_attributes: The list of image attributes,valid when tag type is Image.
        :param str name: The name of OCI product tag.
        :param str push_time: The last push time of OCI product.
        :param int size: The size of OCI product.
        :param str type: The type of OCI product tag.
        """
        pulumi.set(__self__, "chart_attribute", chart_attribute)
        pulumi.set(__self__, "digest", digest)
        pulumi.set(__self__, "image_attributes", image_attributes)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "push_time", push_time)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="chartAttribute")
    def chart_attribute(self) -> 'outputs.TagsTagChartAttributeResult':
        """
        The chart attribute,valid when tag type is Chart.
        """
        return pulumi.get(self, "chart_attribute")

    @property
    @pulumi.getter
    def digest(self) -> str:
        """
        The digest of image.
        """
        return pulumi.get(self, "digest")

    @property
    @pulumi.getter(name="imageAttributes")
    def image_attributes(self) -> Sequence['outputs.TagsTagImageAttributeResult']:
        """
        The list of image attributes,valid when tag type is Image.
        """
        return pulumi.get(self, "image_attributes")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of OCI product tag.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pushTime")
    def push_time(self) -> str:
        """
        The last push time of OCI product.
        """
        return pulumi.get(self, "push_time")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size of OCI product.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of OCI product tag.
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class TagsTagChartAttributeResult(dict):
    def __init__(__self__, *,
                 api_version: str,
                 name: str,
                 version: str):
        """
        :param str api_version: The Helm version.
        :param str name: The name of OCI product tag.
        :param str version: The Helm Chart version.
        """
        pulumi.set(__self__, "api_version", api_version)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> str:
        """
        The Helm version.
        """
        return pulumi.get(self, "api_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of OCI product tag.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The Helm Chart version.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class TagsTagImageAttributeResult(dict):
    def __init__(__self__, *,
                 architecture: str,
                 author: str,
                 digest: str,
                 os: str):
        """
        :param str architecture: The image architecture.
        :param str author: The image author.
        :param str digest: The digest of image.
        :param str os: The iamge os.
        """
        pulumi.set(__self__, "architecture", architecture)
        pulumi.set(__self__, "author", author)
        pulumi.set(__self__, "digest", digest)
        pulumi.set(__self__, "os", os)

    @property
    @pulumi.getter
    def architecture(self) -> str:
        """
        The image architecture.
        """
        return pulumi.get(self, "architecture")

    @property
    @pulumi.getter
    def author(self) -> str:
        """
        The image author.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter
    def digest(self) -> str:
        """
        The digest of image.
        """
        return pulumi.get(self, "digest")

    @property
    @pulumi.getter
    def os(self) -> str:
        """
        The iamge os.
        """
        return pulumi.get(self, "os")


@pulumi.output_type
class VpcEndpointVpc(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountId":
            suggest = "account_id"
        elif key == "subnetId":
            suggest = "subnet_id"
        elif key == "vpcId":
            suggest = "vpc_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VpcEndpointVpc. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VpcEndpointVpc.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VpcEndpointVpc.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_id: Optional[int] = None,
                 subnet_id: Optional[str] = None,
                 vpc_id: Optional[str] = None):
        """
        :param int account_id: The id of the account. When you need to expose the Enterprise Edition instance to a VPC under another primary account, you need to specify the ID of the primary account to which the VPC belongs.
        :param str subnet_id: The id of the subnet. If not specified, the subnet with the most remaining IPs under the VPC will be automatically selected.
        :param str vpc_id: The id of the vpc.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[int]:
        """
        The id of the account. When you need to expose the Enterprise Edition instance to a VPC under another primary account, you need to specify the ID of the primary account to which the VPC belongs.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The id of the subnet. If not specified, the subnet with the most remaining IPs under the VPC will be automatically selected.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The id of the vpc.
        """
        return pulumi.get(self, "vpc_id")


@pulumi.output_type
class VpcEndpointsEndpointResult(dict):
    def __init__(__self__, *,
                 registry: str,
                 vpcs: Sequence['outputs.VpcEndpointsEndpointVpcResult']):
        """
        :param str registry: The CR registry name.
        :param Sequence['VpcEndpointsEndpointVpcArgs'] vpcs: List of vpc information.
        """
        pulumi.set(__self__, "registry", registry)
        pulumi.set(__self__, "vpcs", vpcs)

    @property
    @pulumi.getter
    def registry(self) -> str:
        """
        The CR registry name.
        """
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter
    def vpcs(self) -> Sequence['outputs.VpcEndpointsEndpointVpcResult']:
        """
        List of vpc information.
        """
        return pulumi.get(self, "vpcs")


@pulumi.output_type
class VpcEndpointsEndpointVpcResult(dict):
    def __init__(__self__, *,
                 account_id: int,
                 create_time: str,
                 ip: str,
                 region: str,
                 status: str,
                 subnet_id: str,
                 vpc_id: str):
        """
        :param int account_id: The id of the account.
        :param str create_time: The creation time.
        :param str ip: The IP address of the mirror repository in the VPC.
        :param str region: The region id.
        :param str status: The status of the vpc endpoint.
        :param str subnet_id: The ID of the subnet.
        :param str vpc_id: The ID of the vpc.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "ip", ip)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> int:
        """
        The id of the account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def ip(self) -> str:
        """
        The IP address of the mirror repository in the VPC.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The region id.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the vpc endpoint.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> str:
        """
        The ID of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        The ID of the vpc.
        """
        return pulumi.get(self, "vpc_id")

