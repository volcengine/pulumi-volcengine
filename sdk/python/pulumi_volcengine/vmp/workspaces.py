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
    'WorkspacesResult',
    'AwaitableWorkspacesResult',
    'workspaces',
    'workspaces_output',
]

warnings.warn("""volcengine.vmp.Workspaces has been deprecated in favor of volcengine.vmp.getWorkspaces""", DeprecationWarning)

@pulumi.output_type
class WorkspacesResult:
    """
    A collection of values returned by Workspaces.
    """
    def __init__(__self__, id=None, ids=None, instance_type_ids=None, name=None, output_file=None, project_name=None, statuses=None, tags=None, total_count=None, workspaces=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instance_type_ids and not isinstance(instance_type_ids, list):
            raise TypeError("Expected argument 'instance_type_ids' to be a list")
        pulumi.set(__self__, "instance_type_ids", instance_type_ids)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if workspaces and not isinstance(workspaces, list):
            raise TypeError("Expected argument 'workspaces' to be a list")
        pulumi.set(__self__, "workspaces", workspaces)

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
    @pulumi.getter(name="instanceTypeIds")
    def instance_type_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "instance_type_ids")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of workspace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The project name of vmp workspace.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Sequence['outputs.WorkspacesTagResult']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def workspaces(self) -> Sequence['outputs.WorkspacesWorkspaceResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "workspaces")


class AwaitableWorkspacesResult(WorkspacesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return WorkspacesResult(
            id=self.id,
            ids=self.ids,
            instance_type_ids=self.instance_type_ids,
            name=self.name,
            output_file=self.output_file,
            project_name=self.project_name,
            statuses=self.statuses,
            tags=self.tags,
            total_count=self.total_count,
            workspaces=self.workspaces)


def workspaces(ids: Optional[Sequence[str]] = None,
               instance_type_ids: Optional[Sequence[str]] = None,
               name: Optional[str] = None,
               output_file: Optional[str] = None,
               project_name: Optional[str] = None,
               statuses: Optional[Sequence[str]] = None,
               tags: Optional[Sequence[pulumi.InputType['WorkspacesTagArgs']]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableWorkspacesResult:
    """
    Use this data source to query detailed information of vmp workspaces
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_workspace = volcengine.vmp.Workspace("fooWorkspace",
        instance_type_id="vmp.standard.15d",
        delete_protection_enabled=False,
        description="acc-test-1",
        username="admin123",
        password="*******")
    foo_workspaces = volcengine.vmp.get_workspaces_output(ids=[foo_workspace.id])
    ```


    :param Sequence[str] ids: A list of Workspace IDs.
    :param Sequence[str] instance_type_ids: A list of Instance Type IDs.
    :param str name: The name of workspace.
    :param str output_file: File name where to save data source results.
    :param str project_name: The project name of vmp workspace.
    :param Sequence[str] statuses: A list of Workspace status.
    :param Sequence[pulumi.InputType['WorkspacesTagArgs']] tags: The tags of vmp workspace.
    """
    pulumi.log.warn("""workspaces is deprecated: volcengine.vmp.Workspaces has been deprecated in favor of volcengine.vmp.getWorkspaces""")
    __args__ = dict()
    __args__['ids'] = ids
    __args__['instanceTypeIds'] = instance_type_ids
    __args__['name'] = name
    __args__['outputFile'] = output_file
    __args__['projectName'] = project_name
    __args__['statuses'] = statuses
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:vmp/workspaces:Workspaces', __args__, opts=opts, typ=WorkspacesResult).value

    return AwaitableWorkspacesResult(
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instance_type_ids=pulumi.get(__ret__, 'instance_type_ids'),
        name=pulumi.get(__ret__, 'name'),
        output_file=pulumi.get(__ret__, 'output_file'),
        project_name=pulumi.get(__ret__, 'project_name'),
        statuses=pulumi.get(__ret__, 'statuses'),
        tags=pulumi.get(__ret__, 'tags'),
        total_count=pulumi.get(__ret__, 'total_count'),
        workspaces=pulumi.get(__ret__, 'workspaces'))


@_utilities.lift_output_func(workspaces)
def workspaces_output(ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      instance_type_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      name: Optional[pulumi.Input[Optional[str]]] = None,
                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                      project_name: Optional[pulumi.Input[Optional[str]]] = None,
                      statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                      tags: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['WorkspacesTagArgs']]]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[WorkspacesResult]:
    """
    Use this data source to query detailed information of vmp workspaces
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_workspace = volcengine.vmp.Workspace("fooWorkspace",
        instance_type_id="vmp.standard.15d",
        delete_protection_enabled=False,
        description="acc-test-1",
        username="admin123",
        password="*******")
    foo_workspaces = volcengine.vmp.get_workspaces_output(ids=[foo_workspace.id])
    ```


    :param Sequence[str] ids: A list of Workspace IDs.
    :param Sequence[str] instance_type_ids: A list of Instance Type IDs.
    :param str name: The name of workspace.
    :param str output_file: File name where to save data source results.
    :param str project_name: The project name of vmp workspace.
    :param Sequence[str] statuses: A list of Workspace status.
    :param Sequence[pulumi.InputType['WorkspacesTagArgs']] tags: The tags of vmp workspace.
    """
    pulumi.log.warn("""workspaces is deprecated: volcengine.vmp.Workspaces has been deprecated in favor of volcengine.vmp.getWorkspaces""")
    ...
