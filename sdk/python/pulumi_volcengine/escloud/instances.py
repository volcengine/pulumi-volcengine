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
    'InstancesResult',
    'AwaitableInstancesResult',
    'instances',
    'instances_output',
]

warnings.warn("""volcengine.escloud.Instances has been deprecated in favor of volcengine.escloud.getInstances""", DeprecationWarning)

@pulumi.output_type
class InstancesResult:
    """
    A collection of values returned by Instances.
    """
    def __init__(__self__, charge_types=None, id=None, ids=None, instances=None, names=None, output_file=None, statuses=None, total_count=None, versions=None, zone_ids=None):
        if charge_types and not isinstance(charge_types, list):
            raise TypeError("Expected argument 'charge_types' to be a list")
        pulumi.set(__self__, "charge_types", charge_types)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)
        if zone_ids and not isinstance(zone_ids, list):
            raise TypeError("Expected argument 'zone_ids' to be a list")
        pulumi.set(__self__, "zone_ids", zone_ids)

    @property
    @pulumi.getter(name="chargeTypes")
    def charge_types(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "charge_types")

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
    @pulumi.getter
    def instances(self) -> Sequence['outputs.InstancesInstanceResult']:
        """
        The collection of instance query.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def names(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "statuses")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of instance query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def versions(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "versions")

    @property
    @pulumi.getter(name="zoneIds")
    def zone_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "zone_ids")


class AwaitableInstancesResult(InstancesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return InstancesResult(
            charge_types=self.charge_types,
            id=self.id,
            ids=self.ids,
            instances=self.instances,
            names=self.names,
            output_file=self.output_file,
            statuses=self.statuses,
            total_count=self.total_count,
            versions=self.versions,
            zone_ids=self.zone_ids)


def instances(charge_types: Optional[Sequence[str]] = None,
              ids: Optional[Sequence[str]] = None,
              names: Optional[Sequence[str]] = None,
              output_file: Optional[str] = None,
              statuses: Optional[Sequence[str]] = None,
              versions: Optional[Sequence[str]] = None,
              zone_ids: Optional[Sequence[str]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableInstancesResult:
    """
    (Deprecated! Recommend use escloud_v2.EscloudInstanceV2 replace) Use this data source to query detailed information of escloud instances
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet_new",
        description="tfdesc",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_instance = volcengine.escloud.Instance("fooInstance", instance_configuration=volcengine.escloud.InstanceInstanceConfigurationArgs(
        version="V6_7",
        zone_number=1,
        enable_https=True,
        admin_user_name="admin",
        admin_password="Password@@",
        charge_type="PostPaid",
        configuration_code="es.standard",
        enable_pure_master=True,
        instance_name="acc-test-0",
        node_specs_assigns=[
            volcengine.escloud.InstanceInstanceConfigurationNodeSpecsAssignArgs(
                type="Master",
                number=3,
                resource_spec_name="es.x4.medium",
                storage_spec_name="es.volume.essd.pl0",
                storage_size=100,
            ),
            volcengine.escloud.InstanceInstanceConfigurationNodeSpecsAssignArgs(
                type="Hot",
                number=2,
                resource_spec_name="es.x4.large",
                storage_spec_name="es.volume.essd.pl0",
                storage_size=100,
            ),
            volcengine.escloud.InstanceInstanceConfigurationNodeSpecsAssignArgs(
                type="Kibana",
                number=1,
                resource_spec_name="kibana.x2.small",
            ),
        ],
        subnet_id=foo_subnet.id,
        project_name="default",
        force_restart_after_scale=False,
    ))
    foo_instances = volcengine.escloud.get_instances_output(ids=[foo_instance.id])
    ```


    :param Sequence[str] charge_types: The charge types of instance.
    :param Sequence[str] ids: A list of instance IDs.
    :param Sequence[str] names: The names of instance.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] statuses: The list status of instance.
    :param Sequence[str] versions: The versions of instance.
    :param Sequence[str] zone_ids: The available zone IDs of instance.
    """
    pulumi.log.warn("""instances is deprecated: volcengine.escloud.Instances has been deprecated in favor of volcengine.escloud.getInstances""")
    __args__ = dict()
    __args__['chargeTypes'] = charge_types
    __args__['ids'] = ids
    __args__['names'] = names
    __args__['outputFile'] = output_file
    __args__['statuses'] = statuses
    __args__['versions'] = versions
    __args__['zoneIds'] = zone_ids
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:escloud/instances:Instances', __args__, opts=opts, typ=InstancesResult).value

    return AwaitableInstancesResult(
        charge_types=pulumi.get(__ret__, 'charge_types'),
        id=pulumi.get(__ret__, 'id'),
        ids=pulumi.get(__ret__, 'ids'),
        instances=pulumi.get(__ret__, 'instances'),
        names=pulumi.get(__ret__, 'names'),
        output_file=pulumi.get(__ret__, 'output_file'),
        statuses=pulumi.get(__ret__, 'statuses'),
        total_count=pulumi.get(__ret__, 'total_count'),
        versions=pulumi.get(__ret__, 'versions'),
        zone_ids=pulumi.get(__ret__, 'zone_ids'))


@_utilities.lift_output_func(instances)
def instances_output(charge_types: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     names: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     output_file: Optional[pulumi.Input[Optional[str]]] = None,
                     statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     versions: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     zone_ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[InstancesResult]:
    """
    (Deprecated! Recommend use escloud_v2.EscloudInstanceV2 replace) Use this data source to query detailed information of escloud instances
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.get_zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-vpc",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-test-subnet_new",
        description="tfdesc",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_instance = volcengine.escloud.Instance("fooInstance", instance_configuration=volcengine.escloud.InstanceInstanceConfigurationArgs(
        version="V6_7",
        zone_number=1,
        enable_https=True,
        admin_user_name="admin",
        admin_password="Password@@",
        charge_type="PostPaid",
        configuration_code="es.standard",
        enable_pure_master=True,
        instance_name="acc-test-0",
        node_specs_assigns=[
            volcengine.escloud.InstanceInstanceConfigurationNodeSpecsAssignArgs(
                type="Master",
                number=3,
                resource_spec_name="es.x4.medium",
                storage_spec_name="es.volume.essd.pl0",
                storage_size=100,
            ),
            volcengine.escloud.InstanceInstanceConfigurationNodeSpecsAssignArgs(
                type="Hot",
                number=2,
                resource_spec_name="es.x4.large",
                storage_spec_name="es.volume.essd.pl0",
                storage_size=100,
            ),
            volcengine.escloud.InstanceInstanceConfigurationNodeSpecsAssignArgs(
                type="Kibana",
                number=1,
                resource_spec_name="kibana.x2.small",
            ),
        ],
        subnet_id=foo_subnet.id,
        project_name="default",
        force_restart_after_scale=False,
    ))
    foo_instances = volcengine.escloud.get_instances_output(ids=[foo_instance.id])
    ```


    :param Sequence[str] charge_types: The charge types of instance.
    :param Sequence[str] ids: A list of instance IDs.
    :param Sequence[str] names: The names of instance.
    :param str output_file: File name where to save data source results.
    :param Sequence[str] statuses: The list status of instance.
    :param Sequence[str] versions: The versions of instance.
    :param Sequence[str] zone_ids: The available zone IDs of instance.
    """
    pulumi.log.warn("""instances is deprecated: volcengine.escloud.Instances has been deprecated in favor of volcengine.escloud.getInstances""")
    ...
