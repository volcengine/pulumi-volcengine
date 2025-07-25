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
    'GetRulesResult',
    'AwaitableGetRulesResult',
    'get_rules',
    'get_rules_output',
]

@pulumi.output_type
class GetRulesResult:
    """
    A collection of values returned by getRules.
    """
    def __init__(__self__, id=None, kind=None, name=None, output_file=None, rule_file_names=None, rule_group_names=None, rules=None, status=None, total_count=None, workspace_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rule_file_names and not isinstance(rule_file_names, list):
            raise TypeError("Expected argument 'rule_file_names' to be a list")
        pulumi.set(__self__, "rule_file_names", rule_file_names)
        if rule_group_names and not isinstance(rule_group_names, list):
            raise TypeError("Expected argument 'rule_group_names' to be a list")
        pulumi.set(__self__, "rule_group_names", rule_group_names)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The kind of rule.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="ruleFileNames")
    def rule_file_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "rule_file_names")

    @property
    @pulumi.getter(name="ruleGroupNames")
    def rule_group_names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "rule_group_names")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetRulesRuleResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of rule.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        return pulumi.get(self, "workspace_id")


class AwaitableGetRulesResult(GetRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRulesResult(
            id=self.id,
            kind=self.kind,
            name=self.name,
            output_file=self.output_file,
            rule_file_names=self.rule_file_names,
            rule_group_names=self.rule_group_names,
            rules=self.rules,
            status=self.status,
            total_count=self.total_count,
            workspace_id=self.workspace_id)


def get_rules(kind: Optional[str] = None,
              name: Optional[str] = None,
              output_file: Optional[str] = None,
              rule_file_names: Optional[Sequence[str]] = None,
              rule_group_names: Optional[Sequence[str]] = None,
              status: Optional[str] = None,
              workspace_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRulesResult:
    """
    Use this data source to query detailed information of vmp rules
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vmp.get_rules(kind="Recording",
        workspace_id="baa02ffb-6f22-43c4-841b-ecf90ded****")
    ```


    :param str kind: The kind of rule.
    :param str name: The name of rule.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] rule_file_names: The name of rule file.
    :param Sequence[str] rule_group_names: The name of rule group.
    :param str status: The status of rule.
    :param str workspace_id: The id of workspace.
    """
    __args__ = dict()
    __args__['kind'] = kind
    __args__['name'] = name
    __args__['outputFile'] = output_file
    __args__['ruleFileNames'] = rule_file_names
    __args__['ruleGroupNames'] = rule_group_names
    __args__['status'] = status
    __args__['workspaceId'] = workspace_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vmp/getRules:getRules', __args__, opts=opts, typ=GetRulesResult).value

    return AwaitableGetRulesResult(
        id=pulumi.get(__ret__, 'id'),
        kind=pulumi.get(__ret__, 'kind'),
        name=pulumi.get(__ret__, 'name'),
        output_file=pulumi.get(__ret__, 'output_file'),
        rule_file_names=pulumi.get(__ret__, 'rule_file_names'),
        rule_group_names=pulumi.get(__ret__, 'rule_group_names'),
        rules=pulumi.get(__ret__, 'rules'),
        status=pulumi.get(__ret__, 'status'),
        total_count=pulumi.get(__ret__, 'total_count'),
        workspace_id=pulumi.get(__ret__, 'workspace_id'))


@_utilities.lift_output_func(get_rules)
def get_rules_output(kind: Optional[pulumi.Input[str]] = None,
                     name: Optional[pulumi.Input[Optional[str]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     rule_file_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     rule_group_names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     status: Optional[pulumi.Input[Optional[str]]] = None,
                     workspace_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRulesResult]:
    """
    Use this data source to query detailed information of vmp rules
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.vmp.get_rules(kind="Recording",
        workspace_id="baa02ffb-6f22-43c4-841b-ecf90ded****")
    ```


    :param str kind: The kind of rule.
    :param str name: The name of rule.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] rule_file_names: The name of rule file.
    :param Sequence[str] rule_group_names: The name of rule group.
    :param str status: The status of rule.
    :param str workspace_id: The id of workspace.
    """
    ...
