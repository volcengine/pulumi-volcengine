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
    'TopicsResult',
    'AwaitableTopicsResult',
    'topics',
    'topics_output',
]

warnings.warn("""volcengine.kafka.Topics has been deprecated in favor of volcengine.kafka.getTopics""", DeprecationWarning)

@pulumi.output_type
class TopicsResult:
    """
    A collection of values returned by Topics.
    """
    def __init__(__self__, id=None, instance_id=None, name_regex=None, output_file=None, partition_number=None, replica_number=None, topic_name=None, topics=None, total_count=None, user_name=None):
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
        if partition_number and not isinstance(partition_number, int):
            raise TypeError("Expected argument 'partition_number' to be a int")
        pulumi.set(__self__, "partition_number", partition_number)
        if replica_number and not isinstance(replica_number, int):
            raise TypeError("Expected argument 'replica_number' to be a int")
        pulumi.set(__self__, "replica_number", replica_number)
        if topic_name and not isinstance(topic_name, str):
            raise TypeError("Expected argument 'topic_name' to be a str")
        pulumi.set(__self__, "topic_name", topic_name)
        if topics and not isinstance(topics, list):
            raise TypeError("Expected argument 'topics' to be a list")
        pulumi.set(__self__, "topics", topics)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if user_name and not isinstance(user_name, str):
            raise TypeError("Expected argument 'user_name' to be a str")
        pulumi.set(__self__, "user_name", user_name)

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
    @pulumi.getter(name="partitionNumber")
    def partition_number(self) -> Optional[int]:
        """
        The number of partition in the kafka topic.
        """
        return pulumi.get(self, "partition_number")

    @property
    @pulumi.getter(name="replicaNumber")
    def replica_number(self) -> Optional[int]:
        """
        The number of replica in the kafka topic.
        """
        return pulumi.get(self, "replica_number")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[str]:
        """
        The name of the kafka topic.
        """
        return pulumi.get(self, "topic_name")

    @property
    @pulumi.getter
    def topics(self) -> Sequence['outputs.TopicsTopicResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "topics")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[str]:
        """
        The name of SASL user.
        """
        return pulumi.get(self, "user_name")


class AwaitableTopicsResult(TopicsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return TopicsResult(
            id=self.id,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            partition_number=self.partition_number,
            replica_number=self.replica_number,
            topic_name=self.topic_name,
            topics=self.topics,
            total_count=self.total_count,
            user_name=self.user_name)


def topics(instance_id: Optional[str] = None,
           name_regex: Optional[str] = None,
           output_file: Optional[str] = None,
           partition_number: Optional[int] = None,
           replica_number: Optional[int] = None,
           topic_name: Optional[str] = None,
           user_name: Optional[str] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableTopicsResult:
    """
    Use this data source to query detailed information of kafka topics
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_instance = volcengine.kafka.Instance("fooInstance",
        instance_name="acc-test-kafka",
        instance_description="tf-test",
        version="2.2.2",
        compute_spec="kafka.20xrate.hw",
        subnet_id=foo_subnet.id,
        user_name="tf-user",
        user_password="tf-pass!@q1",
        charge_type="PostPaid",
        storage_space=300,
        partition_number=350,
        project_name="default",
        tags=[volcengine.kafka.InstanceTagArgs(
            key="k1",
            value="v1",
        )],
        parameters=[
            volcengine.kafka.InstanceParameterArgs(
                parameter_name="MessageMaxByte",
                parameter_value="12",
            ),
            volcengine.kafka.InstanceParameterArgs(
                parameter_name="LogRetentionHours",
                parameter_value="70",
            ),
        ])
    foo_sasl_user = volcengine.kafka.SaslUser("fooSaslUser",
        user_name="acc-test-user",
        instance_id=foo_instance.id,
        user_password="suqsnis123!",
        description="tf-test",
        all_authority=True,
        password_type="Scram")
    foo_topic = volcengine.kafka.Topic("fooTopic",
        topic_name="acc-test-topic",
        instance_id=foo_instance.id,
        description="tf-test",
        partition_number=15,
        replica_number=3,
        parameters=volcengine.kafka.TopicParametersArgs(
            min_insync_replica_number=2,
            message_max_byte=10,
            log_retention_hours=96,
        ),
        all_authority=False,
        access_policies=[volcengine.kafka.TopicAccessPolicyArgs(
            user_name=foo_sasl_user.user_name,
            access_policy="Pub",
        )])
    default = volcengine.kafka.get_topics_output(instance_id=foo_topic.instance_id)
    ```


    :param str instance_id: The id of kafka instance.
    :param str name_regex: A Name Regex of kafka topic.
    :param str output_file: File name where to save data source results.
    :param int partition_number: The number of partition in kafka topic.
    :param int replica_number: The number of replica in kafka topic.
    :param str topic_name: The name of kafka topic. This field supports fuzzy query.
    :param str user_name: When a user name is specified, only the access policy of the specified user for this Topic will be returned.
    """
    pulumi.log.warn("""topics is deprecated: volcengine.kafka.Topics has been deprecated in favor of volcengine.kafka.getTopics""")
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['partitionNumber'] = partition_number
    __args__['replicaNumber'] = replica_number
    __args__['topicName'] = topic_name
    __args__['userName'] = user_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:kafka/topics:Topics', __args__, opts=opts, typ=TopicsResult).value

    return AwaitableTopicsResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        partition_number=pulumi.get(__ret__, 'partition_number'),
        replica_number=pulumi.get(__ret__, 'replica_number'),
        topic_name=pulumi.get(__ret__, 'topic_name'),
        topics=pulumi.get(__ret__, 'topics'),
        total_count=pulumi.get(__ret__, 'total_count'),
        user_name=pulumi.get(__ret__, 'user_name'))


@_utilities.lift_output_func(topics)
def topics_output(instance_id: Optional[pulumi.Input[str]] = None,
                  name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                  partition_number: Optional[pulumi.Input[Optional[int]]] = None,
                  replica_number: Optional[pulumi.Input[Optional[int]]] = None,
                  topic_name: Optional[pulumi.Input[Optional[str]]] = None,
                  user_name: Optional[pulumi.Input[Optional[str]]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[TopicsResult]:
    """
    Use this data source to query detailed information of kafka topics
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_instance = volcengine.kafka.Instance("fooInstance",
        instance_name="acc-test-kafka",
        instance_description="tf-test",
        version="2.2.2",
        compute_spec="kafka.20xrate.hw",
        subnet_id=foo_subnet.id,
        user_name="tf-user",
        user_password="tf-pass!@q1",
        charge_type="PostPaid",
        storage_space=300,
        partition_number=350,
        project_name="default",
        tags=[volcengine.kafka.InstanceTagArgs(
            key="k1",
            value="v1",
        )],
        parameters=[
            volcengine.kafka.InstanceParameterArgs(
                parameter_name="MessageMaxByte",
                parameter_value="12",
            ),
            volcengine.kafka.InstanceParameterArgs(
                parameter_name="LogRetentionHours",
                parameter_value="70",
            ),
        ])
    foo_sasl_user = volcengine.kafka.SaslUser("fooSaslUser",
        user_name="acc-test-user",
        instance_id=foo_instance.id,
        user_password="suqsnis123!",
        description="tf-test",
        all_authority=True,
        password_type="Scram")
    foo_topic = volcengine.kafka.Topic("fooTopic",
        topic_name="acc-test-topic",
        instance_id=foo_instance.id,
        description="tf-test",
        partition_number=15,
        replica_number=3,
        parameters=volcengine.kafka.TopicParametersArgs(
            min_insync_replica_number=2,
            message_max_byte=10,
            log_retention_hours=96,
        ),
        all_authority=False,
        access_policies=[volcengine.kafka.TopicAccessPolicyArgs(
            user_name=foo_sasl_user.user_name,
            access_policy="Pub",
        )])
    default = volcengine.kafka.get_topics_output(instance_id=foo_topic.instance_id)
    ```


    :param str instance_id: The id of kafka instance.
    :param str name_regex: A Name Regex of kafka topic.
    :param str output_file: File name where to save data source results.
    :param int partition_number: The number of partition in kafka topic.
    :param int replica_number: The number of replica in kafka topic.
    :param str topic_name: The name of kafka topic. This field supports fuzzy query.
    :param str user_name: When a user name is specified, only the access policy of the specified user for this Topic will be returned.
    """
    pulumi.log.warn("""topics is deprecated: volcengine.kafka.Topics has been deprecated in favor of volcengine.kafka.getTopics""")
    ...
