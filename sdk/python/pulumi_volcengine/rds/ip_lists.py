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

__all__ = [
    'IpListsResult',
    'AwaitableIpListsResult',
    'ip_lists',
    'ip_lists_output',
]

warnings.warn("""volcengine.rds.IpLists has been deprecated in favor of volcengine.rds.getIpLists""", DeprecationWarning)

@pulumi.output_type
class IpListsResult:
    """
    A collection of values returned by IpLists.
    """
    def __init__(__self__, id=None, instance_id=None, name_regex=None, output_file=None, rds_ip_lists=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rds_ip_lists and not isinstance(rds_ip_lists, list):
            raise TypeError("Expected argument 'rds_ip_lists' to be a list")
        pulumi.set(__self__, "rds_ip_lists", rds_ip_lists)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="rdsIpLists")
    def rds_ip_lists(self) -> Sequence['outputs.IpListsRdsIpListResult']:
        """
        The collection of RDS ip list account query.
        """
        return pulumi.get(self, "rds_ip_lists")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of RDS ip list query.
        """
        return pulumi.get(self, "total_count")


class AwaitableIpListsResult(IpListsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IpListsResult(
            id=self.id,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            rds_ip_lists=self.rds_ip_lists,
            total_count=self.total_count)


def ip_lists(instance_id: Optional[str] = None,
             name_regex: Optional[str] = None,
             output_file: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIpListsResult:
    """
    (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds ip lists
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.rds.get_ip_lists(instance_id="mysql-0fdd3bab2e7c")
    ```


    :param str instance_id: The id of the RDS instance.
    :param str name_regex: A Name Regex of RDS ip list.
    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""ip_lists is deprecated: volcengine.rds.IpLists has been deprecated in favor of volcengine.rds.getIpLists""")
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:rds/ipLists:IpLists', __args__, opts=opts, typ=IpListsResult).value

    return AwaitableIpListsResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        rds_ip_lists=pulumi.get(__ret__, 'rds_ip_lists'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(ip_lists)
def ip_lists_output(instance_id: Optional[pulumi.Input[str]] = None,
                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IpListsResult]:
    """
    (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds ip lists
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.rds.get_ip_lists(instance_id="mysql-0fdd3bab2e7c")
    ```


    :param str instance_id: The id of the RDS instance.
    :param str name_regex: A Name Regex of RDS ip list.
    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""ip_lists is deprecated: volcengine.rds.IpLists has been deprecated in favor of volcengine.rds.getIpLists""")
    ...
