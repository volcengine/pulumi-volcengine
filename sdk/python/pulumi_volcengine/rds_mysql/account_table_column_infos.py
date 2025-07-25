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
    'AccountTableColumnInfosResult',
    'AwaitableAccountTableColumnInfosResult',
    'account_table_column_infos',
    'account_table_column_infos_output',
]

warnings.warn("""volcengine.rds_mysql.AccountTableColumnInfos has been deprecated in favor of volcengine.rds_mysql.getAccountTableColumnInfos""", DeprecationWarning)

@pulumi.output_type
class AccountTableColumnInfosResult:
    """
    A collection of values returned by AccountTableColumnInfos.
    """
    def __init__(__self__, account_name=None, column_name=None, db_name=None, host=None, id=None, instance_id=None, output_file=None, table_infos=None, table_limit=None, table_name=None, total_count=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        pulumi.set(__self__, "account_name", account_name)
        if column_name and not isinstance(column_name, str):
            raise TypeError("Expected argument 'column_name' to be a str")
        pulumi.set(__self__, "column_name", column_name)
        if db_name and not isinstance(db_name, str):
            raise TypeError("Expected argument 'db_name' to be a str")
        pulumi.set(__self__, "db_name", db_name)
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if table_infos and not isinstance(table_infos, list):
            raise TypeError("Expected argument 'table_infos' to be a list")
        pulumi.set(__self__, "table_infos", table_infos)
        if table_limit and not isinstance(table_limit, int):
            raise TypeError("Expected argument 'table_limit' to be a int")
        pulumi.set(__self__, "table_limit", table_limit)
        if table_name and not isinstance(table_name, str):
            raise TypeError("Expected argument 'table_name' to be a str")
        pulumi.set(__self__, "table_name", table_name)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[str]:
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="columnName")
    def column_name(self) -> Optional[str]:
        """
        The name of the column.
        """
        return pulumi.get(self, "column_name")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> str:
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        return pulumi.get(self, "host")

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
    @pulumi.getter(name="tableInfos")
    def table_infos(self) -> Sequence['outputs.AccountTableColumnInfosTableInfoResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "table_infos")

    @property
    @pulumi.getter(name="tableLimit")
    def table_limit(self) -> Optional[int]:
        return pulumi.get(self, "table_limit")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[str]:
        """
        The name of the table.
        """
        return pulumi.get(self, "table_name")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableAccountTableColumnInfosResult(AccountTableColumnInfosResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AccountTableColumnInfosResult(
            account_name=self.account_name,
            column_name=self.column_name,
            db_name=self.db_name,
            host=self.host,
            id=self.id,
            instance_id=self.instance_id,
            output_file=self.output_file,
            table_infos=self.table_infos,
            table_limit=self.table_limit,
            table_name=self.table_name,
            total_count=self.total_count)


def account_table_column_infos(account_name: Optional[str] = None,
                               column_name: Optional[str] = None,
                               db_name: Optional[str] = None,
                               host: Optional[str] = None,
                               instance_id: Optional[str] = None,
                               output_file: Optional[str] = None,
                               table_limit: Optional[int] = None,
                               table_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAccountTableColumnInfosResult:
    """
    Use this data source to query detailed information of rds mysql account table column infos
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.rds_mysql.get_account_table_column_infos(db_name="ddd",
        instance_id="mysql-b51d37110dd1")
    ```


    :param str account_name: The name of the account.
    :param str column_name: The name of the column.
    :param str db_name: The name of the database.
    :param str host: Specify the IP address for the account to access the database. The default value is %.
    :param str instance_id: The id of the mysql instance.
    :param str output_file: File name where to save data source results.
    :param int table_limit: Specify the number of tables in the table column permission information to be returned. If it exceeds the setting, it will be truncated.
    :param str table_name: The name of the table.
    """
    pulumi.log.warn("""account_table_column_infos is deprecated: volcengine.rds_mysql.AccountTableColumnInfos has been deprecated in favor of volcengine.rds_mysql.getAccountTableColumnInfos""")
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['columnName'] = column_name
    __args__['dbName'] = db_name
    __args__['host'] = host
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    __args__['tableLimit'] = table_limit
    __args__['tableName'] = table_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:rds_mysql/accountTableColumnInfos:AccountTableColumnInfos', __args__, opts=opts, typ=AccountTableColumnInfosResult).value

    return AwaitableAccountTableColumnInfosResult(
        account_name=pulumi.get(__ret__, 'account_name'),
        column_name=pulumi.get(__ret__, 'column_name'),
        db_name=pulumi.get(__ret__, 'db_name'),
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        table_infos=pulumi.get(__ret__, 'table_infos'),
        table_limit=pulumi.get(__ret__, 'table_limit'),
        table_name=pulumi.get(__ret__, 'table_name'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(account_table_column_infos)
def account_table_column_infos_output(account_name: Optional[pulumi.Input[Optional[str]]] = None,
                                      column_name: Optional[pulumi.Input[Optional[str]]] = None,
                                      db_name: Optional[pulumi.Input[str]] = None,
                                      host: Optional[pulumi.Input[Optional[str]]] = None,
                                      instance_id: Optional[pulumi.Input[str]] = None,
                                      output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                      table_limit: Optional[pulumi.Input[Optional[int]]] = None,
                                      table_name: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AccountTableColumnInfosResult]:
    """
    Use this data source to query detailed information of rds mysql account table column infos
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.rds_mysql.get_account_table_column_infos(db_name="ddd",
        instance_id="mysql-b51d37110dd1")
    ```


    :param str account_name: The name of the account.
    :param str column_name: The name of the column.
    :param str db_name: The name of the database.
    :param str host: Specify the IP address for the account to access the database. The default value is %.
    :param str instance_id: The id of the mysql instance.
    :param str output_file: File name where to save data source results.
    :param int table_limit: Specify the number of tables in the table column permission information to be returned. If it exceeds the setting, it will be truncated.
    :param str table_name: The name of the table.
    """
    pulumi.log.warn("""account_table_column_infos is deprecated: volcengine.rds_mysql.AccountTableColumnInfos has been deprecated in favor of volcengine.rds_mysql.getAccountTableColumnInfos""")
    ...
