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
    'ScalingPoliciesResult',
    'AwaitableScalingPoliciesResult',
    'scaling_policies',
    'scaling_policies_output',
]

warnings.warn("""volcengine.autoscaling.ScalingPolicies has been deprecated in favor of volcengine.autoscaling.getScalingPolicies""", DeprecationWarning)

@pulumi.output_type
class ScalingPoliciesResult:
    """
    A collection of values returned by ScalingPolicies.
    """
    def __init__(__self__, id=None, ids=None, name_regex=None, output_file=None, scaling_group_id=None, scaling_policies=None, scaling_policy_names=None, scaling_policy_type=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if scaling_group_id and not isinstance(scaling_group_id, str):
            raise TypeError("Expected argument 'scaling_group_id' to be a str")
        pulumi.set(__self__, "scaling_group_id", scaling_group_id)
        if scaling_policies and not isinstance(scaling_policies, list):
            raise TypeError("Expected argument 'scaling_policies' to be a list")
        pulumi.set(__self__, "scaling_policies", scaling_policies)
        if scaling_policy_names and not isinstance(scaling_policy_names, list):
            raise TypeError("Expected argument 'scaling_policy_names' to be a list")
        pulumi.set(__self__, "scaling_policy_names", scaling_policy_names)
        if scaling_policy_type and not isinstance(scaling_policy_type, str):
            raise TypeError("Expected argument 'scaling_policy_type' to be a str")
        pulumi.set(__self__, "scaling_policy_type", scaling_policy_type)
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
    @pulumi.getter
    def ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="scalingGroupId")
    def scaling_group_id(self) -> str:
        """
        The id of the scaling group to which the scaling policy belongs.
        """
        return pulumi.get(self, "scaling_group_id")

    @property
    @pulumi.getter(name="scalingPolicies")
    def scaling_policies(self) -> Sequence['outputs.ScalingPoliciesScalingPolicyResult']:
        """
        The collection of scaling policy query.
        """
        return pulumi.get(self, "scaling_policies")

    @property
    @pulumi.getter(name="scalingPolicyNames")
    def scaling_policy_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "scaling_policy_names")

    @property
    @pulumi.getter(name="scalingPolicyType")
    def scaling_policy_type(self) -> Optional[str]:
        """
        The type of the scaling policy.
        """
        return pulumi.get(self, "scaling_policy_type")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of scaling policy query.
        """
        return pulumi.get(self, "total_count")


class AwaitableScalingPoliciesResult(ScalingPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ScalingPoliciesResult(
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            output_file=self.output_file,
            scaling_group_id=self.scaling_group_id,
            scaling_policies=self.scaling_policies,
            scaling_policy_names=self.scaling_policy_names,
            scaling_policy_type=self.scaling_policy_type,
            total_count=self.total_count)


def scaling_policies(ids: Optional[Sequence[str]] = None,
                     name_regex: Optional[str] = None,
                     output_file: Optional[str] = None,
                     scaling_group_id: Optional[str] = None,
                     scaling_policy_names: Optional[Sequence[str]] = None,
                     scaling_policy_type: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableScalingPoliciesResult:
    """
    Use this data source to query detailed information of scaling policies


    :param Sequence[str] ids: A list of scaling policy ids.
    :param str name_regex: A Name Regex of scaling policy.
    :param str output_file: File name where to save data source results.
    :param str scaling_group_id: An id of the scaling group to which the scaling policy belongs.
    :param Sequence[str] scaling_policy_names: A list of scaling policy names.
    :param str scaling_policy_type: A type of scaling policy. Valid values: Scheduled, Recurrence, Manual, Alarm.
    """
    pulumi.log.warn("""scaling_policies is deprecated: volcengine.autoscaling.ScalingPolicies has been deprecated in favor of volcengine.autoscaling.getScalingPolicies""")
    __args__ = dict()
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['scalingGroupId'] = scaling_group_id
    __args__['scalingPolicyNames'] = scaling_policy_names
    __args__['scalingPolicyType'] = scaling_policy_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:autoscaling/scalingPolicies:ScalingPolicies', __args__, opts=opts, typ=ScalingPoliciesResult).value

    return AwaitableScalingPoliciesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        scaling_group_id=pulumi.get(__ret__, 'scaling_group_id'),
        scaling_policies=pulumi.get(__ret__, 'scaling_policies'),
        scaling_policy_names=pulumi.get(__ret__, 'scaling_policy_names'),
        scaling_policy_type=pulumi.get(__ret__, 'scaling_policy_type'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(scaling_policies)
def scaling_policies_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            scaling_group_id: Optional[pulumi.Input[str]] = None,
                            scaling_policy_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            scaling_policy_type: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ScalingPoliciesResult]:
    """
    Use this data source to query detailed information of scaling policies


    :param Sequence[str] ids: A list of scaling policy ids.
    :param str name_regex: A Name Regex of scaling policy.
    :param str output_file: File name where to save data source results.
    :param str scaling_group_id: An id of the scaling group to which the scaling policy belongs.
    :param Sequence[str] scaling_policy_names: A list of scaling policy names.
    :param str scaling_policy_type: A type of scaling policy. Valid values: Scheduled, Recurrence, Manual, Alarm.
    """
    pulumi.log.warn("""scaling_policies is deprecated: volcengine.autoscaling.ScalingPolicies has been deprecated in favor of volcengine.autoscaling.getScalingPolicies""")
    ...
