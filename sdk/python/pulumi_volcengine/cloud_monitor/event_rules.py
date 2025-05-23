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
    'EventRulesResult',
    'AwaitableEventRulesResult',
    'event_rules',
    'event_rules_output',
]

warnings.warn("""volcengine.cloud_monitor.EventRules has been deprecated in favor of volcengine.cloud_monitor.getEventRules""", DeprecationWarning)

@pulumi.output_type
class EventRulesResult:
    """
    A collection of values returned by EventRules.
    """
    def __init__(__self__, id=None, output_file=None, rule_name=None, rules=None, source=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rule_name and not isinstance(rule_name, str):
            raise TypeError("Expected argument 'rule_name' to be a str")
        pulumi.set(__self__, "rule_name", rule_name)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        pulumi.set(__self__, "source", source)
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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[str]:
        """
        The name of the rule.
        """
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.EventRulesRuleResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        """
        Event source corresponding to pattern matching.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableEventRulesResult(EventRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return EventRulesResult(
            id=self.id,
            output_file=self.output_file,
            rule_name=self.rule_name,
            rules=self.rules,
            source=self.source,
            total_count=self.total_count)


def event_rules(output_file: Optional[str] = None,
                rule_name: Optional[str] = None,
                source: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableEventRulesResult:
    """
    Use this data source to query detailed information of cloud monitor event rules
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cloud_monitor.get_event_rules(rule_name="tftest")
    ```


    :param str output_file: File name where to save data source results.
    :param str rule_name: Rule name, search rules by name using fuzzy search.
    :param str source: Event source.
    """
    pulumi.log.warn("""event_rules is deprecated: volcengine.cloud_monitor.EventRules has been deprecated in favor of volcengine.cloud_monitor.getEventRules""")
    __args__ = dict()
    __args__['outputFile'] = output_file
    __args__['ruleName'] = rule_name
    __args__['source'] = source
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:cloud_monitor/eventRules:EventRules', __args__, opts=opts, typ=EventRulesResult).value

    return AwaitableEventRulesResult(
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        rule_name=pulumi.get(__ret__, 'rule_name'),
        rules=pulumi.get(__ret__, 'rules'),
        source=pulumi.get(__ret__, 'source'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(event_rules)
def event_rules_output(output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       rule_name: Optional[pulumi.Input[Optional[str]]] = None,
                       source: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[EventRulesResult]:
    """
    Use this data source to query detailed information of cloud monitor event rules
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.cloud_monitor.get_event_rules(rule_name="tftest")
    ```


    :param str output_file: File name where to save data source results.
    :param str rule_name: Rule name, search rules by name using fuzzy search.
    :param str source: Event source.
    """
    pulumi.log.warn("""event_rules is deprecated: volcengine.cloud_monitor.EventRules has been deprecated in favor of volcengine.cloud_monitor.getEventRules""")
    ...
