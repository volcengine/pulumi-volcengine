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
    'GroupsResult',
    'AwaitableGroupsResult',
    'groups',
    'groups_output',
]

@pulumi.output_type
class GroupsResult:
    """
    A collection of values returned by Groups.
    """
    def __init__(__self__, display_name=None, group_name=None, groups=None, id=None, join_type=None, name_regex=None, output_file=None, total_count=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if join_type and not isinstance(join_type, str):
            raise TypeError("Expected argument 'join_type' to be a str")
        pulumi.set(__self__, "join_type", join_type)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The display name of the cloud identity group.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[str]:
        """
        The name of the cloud identity group.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GroupsGroupResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="joinType")
    def join_type(self) -> Optional[str]:
        """
        The email of the cloud identity group.
        """
        return pulumi.get(self, "join_type")

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
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGroupsResult(GroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GroupsResult(
            display_name=self.display_name,
            group_name=self.group_name,
            groups=self.groups,
            id=self.id,
            join_type=self.join_type,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def groups(display_name: Optional[str] = None,
           group_name: Optional[str] = None,
           join_type: Optional[str] = None,
           name_regex: Optional[str] = None,
           output_file: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGroupsResult:
    """
    Use this data source to query detailed information of cloud identity groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_group = []
    for range in [{"value": i} for i in range(0, 2)]:
        foo_group.append(volcengine.cloud_identity.Group(f"fooGroup-{range['value']}",
            description="tf",
            display_name=f"tf-test-group-{range['value']}",
            group_name=f"acc-test-group-{range['value']}",
            join_type="Manual"))
    foo_groups = volcengine.cloud_identity.groups(group_name="acc-test-group",
        join_type="Manual")
    ```


    :param str display_name: The display name of cloud identity group.
    :param str group_name: The name of cloud identity group.
    :param str join_type: The join type of cloud identity group. Valid values: `Auto`, `Manual`.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['displayName'] = display_name
    __args__['groupName'] = group_name
    __args__['joinType'] = join_type
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:cloud_identity/groups:Groups', __args__, opts=opts, typ=GroupsResult).value

    return AwaitableGroupsResult(
        display_name=pulumi.get(__ret__, 'display_name'),
        group_name=pulumi.get(__ret__, 'group_name'),
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        join_type=pulumi.get(__ret__, 'join_type'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(groups)
def groups_output(display_name: Optional[pulumi.Input[Optional[str]]] = None,
                  group_name: Optional[pulumi.Input[Optional[str]]] = None,
                  join_type: Optional[pulumi.Input[Optional[str]]] = None,
                  name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GroupsResult]:
    """
    Use this data source to query detailed information of cloud identity groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_group = []
    for range in [{"value": i} for i in range(0, 2)]:
        foo_group.append(volcengine.cloud_identity.Group(f"fooGroup-{range['value']}",
            description="tf",
            display_name=f"tf-test-group-{range['value']}",
            group_name=f"acc-test-group-{range['value']}",
            join_type="Manual"))
    foo_groups = volcengine.cloud_identity.groups(group_name="acc-test-group",
        join_type="Manual")
    ```


    :param str display_name: The display name of cloud identity group.
    :param str group_name: The name of cloud identity group.
    :param str join_type: The join type of cloud identity group. Valid values: `Auto`, `Manual`.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    """
    ...