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
    'EscloudNodeAvailableSpecsResult',
    'AwaitableEscloudNodeAvailableSpecsResult',
    'escloud_node_available_specs',
    'escloud_node_available_specs_output',
]

warnings.warn("""volcengine.escloud_v2.EscloudNodeAvailableSpecs has been deprecated in favor of volcengine.escloud_v2.getEscloudNodeAvailableSpecs""", DeprecationWarning)

@pulumi.output_type
class EscloudNodeAvailableSpecsResult:
    """
    A collection of values returned by EscloudNodeAvailableSpecs.
    """
    def __init__(__self__, id=None, instance_id=None, node_specs=None, output_file=None, total_count=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if node_specs and not isinstance(node_specs, list):
            raise TypeError("Expected argument 'node_specs' to be a list")
        pulumi.set(__self__, "node_specs", node_specs)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
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
    def instance_id(self) -> Optional[str]:
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="nodeSpecs")
    def node_specs(self) -> Sequence['outputs.EscloudNodeAvailableSpecsNodeSpecResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "node_specs")

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


class AwaitableEscloudNodeAvailableSpecsResult(EscloudNodeAvailableSpecsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return EscloudNodeAvailableSpecsResult(
            id=self.id,
            instance_id=self.instance_id,
            node_specs=self.node_specs,
            output_file=self.output_file,
            total_count=self.total_count)


def escloud_node_available_specs(instance_id: Optional[str] = None,
                                 output_file: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableEscloudNodeAvailableSpecsResult:
    """
    Use this data source to query detailed information of escloud node available specs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.escloud_v2.get_escloud_node_available_specs()
    ```


    :param str instance_id: The id of the instance.
    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""escloud_node_available_specs is deprecated: volcengine.escloud_v2.EscloudNodeAvailableSpecs has been deprecated in favor of volcengine.escloud_v2.getEscloudNodeAvailableSpecs""")
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:escloud_v2/escloudNodeAvailableSpecs:EscloudNodeAvailableSpecs', __args__, opts=opts, typ=EscloudNodeAvailableSpecsResult).value

    return AwaitableEscloudNodeAvailableSpecsResult(
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        node_specs=pulumi.get(__ret__, 'node_specs'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(escloud_node_available_specs)
def escloud_node_available_specs_output(instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                                        output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[EscloudNodeAvailableSpecsResult]:
    """
    Use this data source to query detailed information of escloud node available specs
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.escloud_v2.get_escloud_node_available_specs()
    ```


    :param str instance_id: The id of the instance.
    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""escloud_node_available_specs is deprecated: volcengine.escloud_v2.EscloudNodeAvailableSpecs has been deprecated in favor of volcengine.escloud_v2.getEscloudNodeAvailableSpecs""")
    ...
