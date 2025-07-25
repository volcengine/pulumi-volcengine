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
    'GetConsumerGroupsResult',
    'AwaitableGetConsumerGroupsResult',
    'get_consumer_groups',
    'get_consumer_groups_output',
]

@pulumi.output_type
class GetConsumerGroupsResult:
    """
    A collection of values returned by getConsumerGroups.
    """
    def __init__(__self__, consumer_group_name=None, consumer_groups=None, iam_project_name=None, id=None, name_regex=None, output_file=None, project_id=None, project_name=None, topic_id=None, topic_name=None, total_count=None):
        if consumer_group_name and not isinstance(consumer_group_name, str):
            raise TypeError("Expected argument 'consumer_group_name' to be a str")
        pulumi.set(__self__, "consumer_group_name", consumer_group_name)
        if consumer_groups and not isinstance(consumer_groups, list):
            raise TypeError("Expected argument 'consumer_groups' to be a list")
        pulumi.set(__self__, "consumer_groups", consumer_groups)
        if iam_project_name and not isinstance(iam_project_name, str):
            raise TypeError("Expected argument 'iam_project_name' to be a str")
        pulumi.set(__self__, "iam_project_name", iam_project_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if topic_id and not isinstance(topic_id, str):
            raise TypeError("Expected argument 'topic_id' to be a str")
        pulumi.set(__self__, "topic_id", topic_id)
        if topic_name and not isinstance(topic_name, str):
            raise TypeError("Expected argument 'topic_name' to be a str")
        pulumi.set(__self__, "topic_name", topic_name)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="consumerGroupName")
    def consumer_group_name(self) -> Optional[str]:
        """
        The name of the consumer group.
        """
        return pulumi.get(self, "consumer_group_name")

    @property
    @pulumi.getter(name="consumerGroups")
    def consumer_groups(self) -> Sequence['outputs.GetConsumerGroupsConsumerGroupResult']:
        """
        List of log service consumption groups.
        """
        return pulumi.get(self, "consumer_groups")

    @property
    @pulumi.getter(name="iamProjectName")
    def iam_project_name(self) -> Optional[str]:
        return pulumi.get(self, "iam_project_name")

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
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[str]:
        """
        The log project ID to which the consumption group belongs.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The name of the log item to which the consumption group belongs.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> Optional[str]:
        """
        The list of log topic ids to be consumed by the consumer group.
        """
        return pulumi.get(self, "topic_id")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[str]:
        return pulumi.get(self, "topic_name")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetConsumerGroupsResult(GetConsumerGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConsumerGroupsResult(
            consumer_group_name=self.consumer_group_name,
            consumer_groups=self.consumer_groups,
            iam_project_name=self.iam_project_name,
            id=self.id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            project_id=self.project_id,
            project_name=self.project_name,
            topic_id=self.topic_id,
            topic_name=self.topic_name,
            total_count=self.total_count)


def get_consumer_groups(consumer_group_name: Optional[str] = None,
                        iam_project_name: Optional[str] = None,
                        name_regex: Optional[str] = None,
                        output_file: Optional[str] = None,
                        project_id: Optional[str] = None,
                        project_name: Optional[str] = None,
                        topic_id: Optional[str] = None,
                        topic_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConsumerGroupsResult:
    """
    Use this data source to query detailed information of tls consumer groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.get_consumer_groups()
    ```


    :param str consumer_group_name: The name of the consumer group.
    :param str iam_project_name: IAM log project name.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str project_id: The log project ID to which the consumption group belongs.
    :param str project_name: The name of the log item to which the consumption group belongs.
    :param str topic_id: The log topic ID to which the consumer belongs.
    :param str topic_name: The name of the log topic to which the consumption group belongs.
    """
    __args__ = dict()
    __args__['consumerGroupName'] = consumer_group_name
    __args__['iamProjectName'] = iam_project_name
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['projectId'] = project_id
    __args__['projectName'] = project_name
    __args__['topicId'] = topic_id
    __args__['topicName'] = topic_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:tls/getConsumerGroups:getConsumerGroups', __args__, opts=opts, typ=GetConsumerGroupsResult).value

    return AwaitableGetConsumerGroupsResult(
        consumer_group_name=pulumi.get(__ret__, 'consumer_group_name'),
        consumer_groups=pulumi.get(__ret__, 'consumer_groups'),
        iam_project_name=pulumi.get(__ret__, 'iam_project_name'),
        id=pulumi.get(__ret__, 'id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        project_id=pulumi.get(__ret__, 'project_id'),
        project_name=pulumi.get(__ret__, 'project_name'),
        topic_id=pulumi.get(__ret__, 'topic_id'),
        topic_name=pulumi.get(__ret__, 'topic_name'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_consumer_groups)
def get_consumer_groups_output(consumer_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                               iam_project_name: Optional[pulumi.Input[Optional[str]]] = None,
                               name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                               output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               project_id: Optional[pulumi.Input[Optional[str]]] = None,
                               project_name: Optional[pulumi.Input[Optional[str]]] = None,
                               topic_id: Optional[pulumi.Input[Optional[str]]] = None,
                               topic_name: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConsumerGroupsResult]:
    """
    Use this data source to query detailed information of tls consumer groups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.get_consumer_groups()
    ```


    :param str consumer_group_name: The name of the consumer group.
    :param str iam_project_name: IAM log project name.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str project_id: The log project ID to which the consumption group belongs.
    :param str project_name: The name of the log item to which the consumption group belongs.
    :param str topic_id: The log topic ID to which the consumer belongs.
    :param str topic_name: The name of the log topic to which the consumption group belongs.
    """
    ...
