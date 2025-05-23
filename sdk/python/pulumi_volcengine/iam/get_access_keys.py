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
    'GetAccessKeysResult',
    'AwaitableGetAccessKeysResult',
    'get_access_keys',
    'get_access_keys_output',
]

@pulumi.output_type
class GetAccessKeysResult:
    """
    A collection of values returned by getAccessKeys.
    """
    def __init__(__self__, access_key_metadatas=None, id=None, name_regex=None, output_file=None, total_count=None, user_name=None):
        if access_key_metadatas and not isinstance(access_key_metadatas, list):
            raise TypeError("Expected argument 'access_key_metadatas' to be a list")
        pulumi.set(__self__, "access_key_metadatas", access_key_metadatas)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accessKeyMetadatas")
    def access_key_metadatas(self) -> Sequence['outputs.GetAccessKeysAccessKeyMetadataResult']:
        """
        The collection of access keys.
        """
        return pulumi.get(self, "access_key_metadatas")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of user query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[str]:
        """
        The user name.
        """
        return pulumi.get(self, "user_name")


class AwaitableGetAccessKeysResult(GetAccessKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessKeysResult(
            access_key_metadatas=self.access_key_metadatas,
            id=self.id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count,
            user_name=self.user_name)


def get_access_keys(name_regex: Optional[str] = None,
                    output_file: Optional[str] = None,
                    user_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessKeysResult:
    """
    Use this data source to query detailed information of iam access keys
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.iam.get_access_keys()
    ```


    :param str name_regex: A Name Regex of IAM.
    :param str output_file: File name where to save data source results.
    :param str user_name: The user names.
    """
    __args__ = dict()
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['userName'] = user_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:iam/getAccessKeys:getAccessKeys', __args__, opts=opts, typ=GetAccessKeysResult).value

    return AwaitableGetAccessKeysResult(
        access_key_metadatas=pulumi.get(__ret__, 'access_key_metadatas'),
        id=pulumi.get(__ret__, 'id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'),
        user_name=pulumi.get(__ret__, 'user_name'))


@_utilities.lift_output_func(get_access_keys)
def get_access_keys_output(name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                           output_file: Optional[pulumi.Input[Optional[str]]] = None,
                           user_name: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccessKeysResult]:
    """
    Use this data source to query detailed information of iam access keys
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.iam.get_access_keys()
    ```


    :param str name_regex: A Name Regex of IAM.
    :param str output_file: File name where to save data source results.
    :param str user_name: The user names.
    """
    ...
