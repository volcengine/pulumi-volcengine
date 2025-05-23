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
    'RuleAppliersResult',
    'AwaitableRuleAppliersResult',
    'rule_appliers',
    'rule_appliers_output',
]

warnings.warn("""volcengine.tls.RuleAppliers has been deprecated in favor of volcengine.tls.getRuleAppliers""", DeprecationWarning)

@pulumi.output_type
class RuleAppliersResult:
    """
    A collection of values returned by RuleAppliers.
    """
    def __init__(__self__, host_group_id=None, id=None, output_file=None, rules=None, total_count=None):
        if host_group_id and not isinstance(host_group_id, str):
            raise TypeError("Expected argument 'host_group_id' to be a str")
        pulumi.set(__self__, "host_group_id", host_group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="hostGroupId")
    def host_group_id(self) -> str:
        return pulumi.get(self, "host_group_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.RuleAppliersRuleResult']:
        """
        The rules list.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableRuleAppliersResult(RuleAppliersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return RuleAppliersResult(
            host_group_id=self.host_group_id,
            id=self.id,
            output_file=self.output_file,
            rules=self.rules,
            total_count=self.total_count)


def rule_appliers(host_group_id: Optional[str] = None,
                  output_file: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableRuleAppliersResult:
    """
    Use this data source to query detailed information of tls rule appliers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.get_rule_appliers(host_group_id="fbea6619-7b0c-40f3-ac7e-45c63e3f676e")
    ```


    :param str host_group_id: The host group id.
    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""rule_appliers is deprecated: volcengine.tls.RuleAppliers has been deprecated in favor of volcengine.tls.getRuleAppliers""")
    __args__ = dict()
    __args__['hostGroupId'] = host_group_id
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:tls/ruleAppliers:RuleAppliers', __args__, opts=opts, typ=RuleAppliersResult).value

    return AwaitableRuleAppliersResult(
        host_group_id=pulumi.get(__ret__, 'host_group_id'),
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        rules=pulumi.get(__ret__, 'rules'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(rule_appliers)
def rule_appliers_output(host_group_id: Optional[pulumi.Input[str]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[RuleAppliersResult]:
    """
    Use this data source to query detailed information of tls rule appliers
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tls.get_rule_appliers(host_group_id="fbea6619-7b0c-40f3-ac7e-45c63e3f676e")
    ```


    :param str host_group_id: The host group id.
    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""rule_appliers is deprecated: volcengine.tls.RuleAppliers has been deprecated in favor of volcengine.tls.getRuleAppliers""")
    ...
