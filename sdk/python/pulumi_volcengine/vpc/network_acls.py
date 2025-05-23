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

__all__ = [
    'NetworkAclsResult',
    'AwaitableNetworkAclsResult',
    'network_acls',
    'network_acls_output',
]

warnings.warn("""volcengine.vpc.NetworkAcls has been deprecated in favor of volcengine.vpc.getNetworkAcls""", DeprecationWarning)

@pulumi.output_type
class NetworkAclsResult:
    """
    A collection of values returned by NetworkAcls.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, network_acl_name=None, network_acls=None, output_file=None, project_name=None, subnet_id=None, tags=None, total_count=None, vpc_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if network_acl_name and not isinstance(network_acl_name, str):
            raise TypeError("Expected argument 'network_acl_name' to be a str")
        pulumi.set(__self__, "network_acl_name", network_acl_name)
        if network_acls and not isinstance(network_acls, list):
            raise TypeError("Expected argument 'network_acls' to be a list")
        pulumi.set(__self__, "network_acls", network_acls)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if subnet_id and not isinstance(subnet_id, str):
            raise TypeError("Expected argument 'subnet_id' to be a str")
        pulumi.set(__self__, "subnet_id", subnet_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="networkAclName")
    def network_acl_name(self) -> Optional[str]:
        """
        The Name of Network Acl.
        """
        return pulumi.get(self, "network_acl_name")

    @property
    @pulumi.getter(name="networkAcls")
    def network_acls(self) -> Sequence['outputs.NetworkAclsNetworkAclResult']:
        """
        The collection of Network Acl query.
        """
        return pulumi.get(self, "network_acls")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The project name of the network acl.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.NetworkAclsTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of Network Acl query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        """
        The vpc id of Network Acl.
        """
        return pulumi.get(self, "vpc_id")


class AwaitableNetworkAclsResult(NetworkAclsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return NetworkAclsResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            network_acl_name=self.network_acl_name,
            network_acls=self.network_acls,
            output_file=self.output_file,
            project_name=self.project_name,
            subnet_id=self.subnet_id,
            tags=self.tags,
            total_count=self.total_count,
            vpc_id=self.vpc_id)


def network_acls(ids: Optional[Sequence[str]] = None,
                 name_regex: Optional[str] = None,
                 network_acl_name: Optional[str] = None,
                 output_file: Optional[str] = None,
                 project_name: Optional[str] = None,
                 subnet_id: Optional[str] = None,
                 tags: Optional[Sequence[pulumi.InputType['NetworkAclsTagArgs']]] = None,
                 vpc_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableNetworkAclsResult:
    """
    Use this data source to query detailed information of network acls
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpc.get_network_acls(network_acl_name="ms-tf-acl")
    ```


    :param Sequence[str] ids: A list of Network Acl IDs.
    :param str name_regex: A Name Regex of Network Acl.
    :param str network_acl_name: The name of Network Acl.
    :param str output_file: File name where to save data source results.
    :param str project_name: The project name of the network acl.
    :param str subnet_id: The subnet id of Network Acl.
    :param Sequence[pulumi.InputType['NetworkAclsTagArgs']] tags: Tags.
    :param str vpc_id: The vpc id of Network Acl.
    """
    pulumi.log.warn("""network_acls is deprecated: volcengine.vpc.NetworkAcls has been deprecated in favor of volcengine.vpc.getNetworkAcls""")
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['networkAclName'] = network_acl_name
    __args__['outputFile'] = output_file
    __args__['projectName'] = project_name
    __args__['subnetId'] = subnet_id
    __args__['tags'] = tags
    __args__['vpcId'] = vpc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vpc/networkAcls:NetworkAcls', __args__, opts=opts, typ=NetworkAclsResult).value

    return AwaitableNetworkAclsResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        network_acl_name=pulumi.get(__ret__, 'network_acl_name'),
        network_acls=pulumi.get(__ret__, 'network_acls'),
        output_file=pulumi.get(__ret__, 'output_file'),
        project_name=pulumi.get(__ret__, 'project_name'),
        subnet_id=pulumi.get(__ret__, 'subnet_id'),
        tags=pulumi.get(__ret__, 'tags'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'))


@_utilities.lift_output_func(network_acls)
def network_acls_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                        name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                        network_acl_name: Optional[pulumi.Input[Optional[str]]] = None,
                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                        project_name: Optional[pulumi.Input[Optional[str]]] = None,
                        subnet_id: Optional[pulumi.Input[Optional[str]]] = None,
                        tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['NetworkAclsTagArgs']]]]] = None,
                        vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[NetworkAclsResult]:
    """
    Use this data source to query detailed information of network acls
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vpc.get_network_acls(network_acl_name="ms-tf-acl")
    ```


    :param Sequence[str] ids: A list of Network Acl IDs.
    :param str name_regex: A Name Regex of Network Acl.
    :param str network_acl_name: The name of Network Acl.
    :param str output_file: File name where to save data source results.
    :param str project_name: The project name of the network acl.
    :param str subnet_id: The subnet id of Network Acl.
    :param Sequence[pulumi.InputType['NetworkAclsTagArgs']] tags: Tags.
    :param str vpc_id: The vpc id of Network Acl.
    """
    pulumi.log.warn("""network_acls is deprecated: volcengine.vpc.NetworkAcls has been deprecated in favor of volcengine.vpc.getNetworkAcls""")
    ...
