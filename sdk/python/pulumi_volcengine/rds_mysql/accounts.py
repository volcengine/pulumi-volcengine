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
    'AccountsResult',
    'AwaitableAccountsResult',
    'accounts',
    'accounts_output',
]

warnings.warn("""volcengine.rds_mysql.Accounts has been deprecated in favor of volcengine.rds_mysql.getAccounts""", DeprecationWarning)

@pulumi.output_type
class AccountsResult:
    """
    A collection of values returned by Accounts.
    """
    def __init__(__self__, account_name=None, accounts=None, id=None, instance_id=None, name_regex=None, output_file=None, total_count=None):
        if account_name and not isinstance(account_name, str):
            raise TypeError("Expected argument 'account_name' to be a str")
        pulumi.set(__self__, "account_name", account_name)
        if accounts and not isinstance(accounts, list):
            raise TypeError("Expected argument 'accounts' to be a list")
        pulumi.set(__self__, "accounts", accounts)
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
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[str]:
        """
        The name of the database account.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter
    def accounts(self) -> Sequence['outputs.AccountsAccountResult']:
        """
        The collection of RDS instance account query.
        """
        return pulumi.get(self, "accounts")

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
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of database account query.
        """
        return pulumi.get(self, "total_count")


class AwaitableAccountsResult(AccountsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return AccountsResult(
            account_name=self.account_name,
            accounts=self.accounts,
            id=self.id,
            instance_id=self.instance_id,
            name_regex=self.name_regex,
            output_file=self.output_file,
            total_count=self.total_count)


def accounts(account_name: Optional[str] = None,
             instance_id: Optional[str] = None,
             name_regex: Optional[str] = None,
             output_file: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableAccountsResult:
    """
    Use this data source to query detailed information of rds mysql accounts
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
    foo_instance = volcengine.rds_mysql.Instance("fooInstance",
        instance_name="acc-test-rds-mysql",
        db_engine_version="MySQL_5_7",
        node_spec="rds.mysql.1c2g",
        primary_zone_id=foo_zones.zones[0].id,
        secondary_zone_id=foo_zones.zones[0].id,
        storage_space=80,
        subnet_id=foo_subnet.id,
        lower_case_table_names="1",
        charge_info=volcengine.rds_mysql.InstanceChargeInfoArgs(
            charge_type="PostPaid",
        ),
        parameters=[
            volcengine.rds_mysql.InstanceParameterArgs(
                parameter_name="auto_increment_increment",
                parameter_value="2",
            ),
            volcengine.rds_mysql.InstanceParameterArgs(
                parameter_name="auto_increment_offset",
                parameter_value="4",
            ),
        ])
    foo_database = volcengine.rds_mysql.Database("fooDatabase",
        db_name="acc-test-db",
        instance_id=foo_instance.id)
    foo_account = volcengine.rds_mysql.Account("fooAccount",
        account_name="acc-test-account",
        account_password="93f0cb0614Aab12",
        account_type="Normal",
        instance_id=foo_instance.id,
        account_privileges=[volcengine.rds_mysql.AccountAccountPrivilegeArgs(
            db_name=foo_database.db_name,
            account_privilege="Custom",
            account_privilege_detail="SELECT,INSERT",
        )])
    foo_accounts = volcengine.rds_mysql.get_accounts_output(instance_id=foo_instance.id,
        account_name=foo_account.account_name)
    ```


    :param str account_name: The name of the database account. This field supports fuzzy query.
    :param str instance_id: The id of the RDS instance.
    :param str name_regex: A Name Regex of database account.
    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""accounts is deprecated: volcengine.rds_mysql.Accounts has been deprecated in favor of volcengine.rds_mysql.getAccounts""")
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['instanceId'] = instance_id
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:rds_mysql/accounts:Accounts', __args__, opts=opts, typ=AccountsResult).value

    return AwaitableAccountsResult(
        account_name=pulumi.get(__ret__, 'account_name'),
        accounts=pulumi.get(__ret__, 'accounts'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(accounts)
def accounts_output(account_name: Optional[pulumi.Input[Optional[str]]] = None,
                    instance_id: Optional[pulumi.Input[str]] = None,
                    name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                    output_file: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[AccountsResult]:
    """
    Use this data source to query detailed information of rds mysql accounts
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
    foo_instance = volcengine.rds_mysql.Instance("fooInstance",
        instance_name="acc-test-rds-mysql",
        db_engine_version="MySQL_5_7",
        node_spec="rds.mysql.1c2g",
        primary_zone_id=foo_zones.zones[0].id,
        secondary_zone_id=foo_zones.zones[0].id,
        storage_space=80,
        subnet_id=foo_subnet.id,
        lower_case_table_names="1",
        charge_info=volcengine.rds_mysql.InstanceChargeInfoArgs(
            charge_type="PostPaid",
        ),
        parameters=[
            volcengine.rds_mysql.InstanceParameterArgs(
                parameter_name="auto_increment_increment",
                parameter_value="2",
            ),
            volcengine.rds_mysql.InstanceParameterArgs(
                parameter_name="auto_increment_offset",
                parameter_value="4",
            ),
        ])
    foo_database = volcengine.rds_mysql.Database("fooDatabase",
        db_name="acc-test-db",
        instance_id=foo_instance.id)
    foo_account = volcengine.rds_mysql.Account("fooAccount",
        account_name="acc-test-account",
        account_password="93f0cb0614Aab12",
        account_type="Normal",
        instance_id=foo_instance.id,
        account_privileges=[volcengine.rds_mysql.AccountAccountPrivilegeArgs(
            db_name=foo_database.db_name,
            account_privilege="Custom",
            account_privilege_detail="SELECT,INSERT",
        )])
    foo_accounts = volcengine.rds_mysql.get_accounts_output(instance_id=foo_instance.id,
        account_name=foo_account.account_name)
    ```


    :param str account_name: The name of the database account. This field supports fuzzy query.
    :param str instance_id: The id of the RDS instance.
    :param str name_regex: A Name Regex of database account.
    :param str output_file: File name where to save data source results.
    """
    pulumi.log.warn("""accounts is deprecated: volcengine.rds_mysql.Accounts has been deprecated in favor of volcengine.rds_mysql.getAccounts""")
    ...
