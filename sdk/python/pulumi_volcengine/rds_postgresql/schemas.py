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
    'SchemasResult',
    'AwaitableSchemasResult',
    'schemas',
    'schemas_output',
]

@pulumi.output_type
class SchemasResult:
    """
    A collection of values returned by Schemas.
    """
    def __init__(__self__, db_name=None, id=None, instance_id=None, output_file=None, schemas=None, total_count=None):
        if db_name and not isinstance(db_name, str):
            raise TypeError("Expected argument 'db_name' to be a str")
        pulumi.set(__self__, "db_name", db_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if schemas and not isinstance(schemas, list):
            raise TypeError("Expected argument 'schemas' to be a list")
        pulumi.set(__self__, "schemas", schemas)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> Optional[str]:
        """
        The name of the database.
        """
        return pulumi.get(self, "db_name")

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def schemas(self) -> Sequence['outputs.SchemasSchemaResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "schemas")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableSchemasResult(SchemasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SchemasResult(
            db_name=self.db_name,
            id=self.id,
            instance_id=self.instance_id,
            output_file=self.output_file,
            schemas=self.schemas,
            total_count=self.total_count)


def schemas(db_name: Optional[str] = None,
            instance_id: Optional[str] = None,
            output_file: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSchemasResult:
    """
    Use this data source to query detailed information of rds postgresql schemas
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-project1",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-subnet-test-2",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_instance = volcengine.rds_postgresql.Instance("fooInstance",
        db_engine_version="PostgreSQL_12",
        node_spec="rds.postgres.1c2g",
        primary_zone_id=foo_zones.zones[0].id,
        secondary_zone_id=foo_zones.zones[0].id,
        storage_space=40,
        subnet_id=foo_subnet.id,
        instance_name="acc-test-1",
        charge_info=volcengine.rds_postgresql.InstanceChargeInfoArgs(
            charge_type="PostPaid",
        ),
        project_name="default",
        tags=[volcengine.rds_postgresql.InstanceTagArgs(
            key="tfk1",
            value="tfv1",
        )],
        parameters=[
            volcengine.rds_postgresql.InstanceParameterArgs(
                name="auto_explain.log_analyze",
                value="off",
            ),
            volcengine.rds_postgresql.InstanceParameterArgs(
                name="auto_explain.log_format",
                value="text",
            ),
        ])
    foo_database = volcengine.rds_postgresql.Database("fooDatabase",
        db_name="acc-test",
        instance_id=foo_instance.id,
        c_type="C",
        collate="zh_CN.utf8")
    foo_account = volcengine.rds_postgresql.Account("fooAccount",
        account_name="acc-test-account",
        account_password="9wc@********12",
        account_type="Normal",
        instance_id=foo_instance.id,
        account_privileges="Inherit,Login,CreateRole,CreateDB")
    foo1 = volcengine.rds_postgresql.Account("foo1",
        account_name="acc-test-account1",
        account_password="9wc@*******12",
        account_type="Normal",
        instance_id=foo_instance.id,
        account_privileges="Inherit,Login,CreateRole,CreateDB")
    foo_schema = volcengine.rds_postgresql.Schema("fooSchema",
        db_name=foo_database.db_name,
        instance_id=foo_instance.id,
        owner=foo_account.account_name,
        schema_name="acc-test-schema")
    foo_schemas = volcengine.rds_postgresql.schemas_output(db_name=foo_schema.db_name,
        instance_id=foo_instance.id)
    ```


    :param str db_name: The name of the database.
    :param str instance_id: The id of the instance.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['dbName'] = db_name
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:rds_postgresql/schemas:Schemas', __args__, opts=opts, typ=SchemasResult).value

    return AwaitableSchemasResult(
        db_name=pulumi.get(__ret__, 'db_name'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        schemas=pulumi.get(__ret__, 'schemas'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(schemas)
def schemas_output(db_name: Optional[pulumi.Input[Optional[str]]] = None,
                   instance_id: Optional[pulumi.Input[str]] = None,
                   output_file: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SchemasResult]:
    """
    Use this data source to query detailed information of rds postgresql schemas
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo_zones = volcengine.ecs.zones()
    foo_vpc = volcengine.vpc.Vpc("fooVpc",
        vpc_name="acc-test-project1",
        cidr_block="172.16.0.0/16")
    foo_subnet = volcengine.vpc.Subnet("fooSubnet",
        subnet_name="acc-subnet-test-2",
        cidr_block="172.16.0.0/24",
        zone_id=foo_zones.zones[0].id,
        vpc_id=foo_vpc.id)
    foo_instance = volcengine.rds_postgresql.Instance("fooInstance",
        db_engine_version="PostgreSQL_12",
        node_spec="rds.postgres.1c2g",
        primary_zone_id=foo_zones.zones[0].id,
        secondary_zone_id=foo_zones.zones[0].id,
        storage_space=40,
        subnet_id=foo_subnet.id,
        instance_name="acc-test-1",
        charge_info=volcengine.rds_postgresql.InstanceChargeInfoArgs(
            charge_type="PostPaid",
        ),
        project_name="default",
        tags=[volcengine.rds_postgresql.InstanceTagArgs(
            key="tfk1",
            value="tfv1",
        )],
        parameters=[
            volcengine.rds_postgresql.InstanceParameterArgs(
                name="auto_explain.log_analyze",
                value="off",
            ),
            volcengine.rds_postgresql.InstanceParameterArgs(
                name="auto_explain.log_format",
                value="text",
            ),
        ])
    foo_database = volcengine.rds_postgresql.Database("fooDatabase",
        db_name="acc-test",
        instance_id=foo_instance.id,
        c_type="C",
        collate="zh_CN.utf8")
    foo_account = volcengine.rds_postgresql.Account("fooAccount",
        account_name="acc-test-account",
        account_password="9wc@********12",
        account_type="Normal",
        instance_id=foo_instance.id,
        account_privileges="Inherit,Login,CreateRole,CreateDB")
    foo1 = volcengine.rds_postgresql.Account("foo1",
        account_name="acc-test-account1",
        account_password="9wc@*******12",
        account_type="Normal",
        instance_id=foo_instance.id,
        account_privileges="Inherit,Login,CreateRole,CreateDB")
    foo_schema = volcengine.rds_postgresql.Schema("fooSchema",
        db_name=foo_database.db_name,
        instance_id=foo_instance.id,
        owner=foo_account.account_name,
        schema_name="acc-test-schema")
    foo_schemas = volcengine.rds_postgresql.schemas_output(db_name=foo_schema.db_name,
        instance_id=foo_instance.id)
    ```


    :param str db_name: The name of the database.
    :param str instance_id: The id of the instance.
    :param str output_file: File name where to save data source results.
    """
    ...