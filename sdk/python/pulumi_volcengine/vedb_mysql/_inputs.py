# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AccountAccountPrivilegeArgs',
    'BackupBackupPolicyArgs',
    'InstanceTagArgs',
    'InstancesTagArgs',
    'GetInstancesTagArgs',
]

@pulumi.input_type
class AccountAccountPrivilegeArgs:
    def __init__(__self__, *,
                 account_privilege: pulumi.Input[str],
                 db_name: pulumi.Input[str],
                 account_privilege_detail: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] account_privilege: Authorization database privilege types: 
               ReadWrite: Read and write privilege.
               ReadOnly: Read-only privilege.
               DDLOnly: Only DDL privilege.
               DMLOnly: Only DML privilege.
               Custom: Custom privilege.
        :param pulumi.Input[str] db_name: Database name requiring authorization.
        :param pulumi.Input[str] account_privilege_detail: The specific SQL operation permissions contained in the permission type are separated by English commas (,) between multiple strings.
               When used as a request parameter in the CreateDatabase interface, when the AccountPrivilege value is Custom, this parameter is required. Value range (multiple selections allowed): SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER. When used as a return parameter in the DescribeDatabases interface, regardless of the value of AccountPrivilege, the details of the SQL operation permissions contained in this permission type are returned. For the specific SQL operation permissions contained in each permission type, please refer to the account permission list.
        """
        pulumi.set(__self__, "account_privilege", account_privilege)
        pulumi.set(__self__, "db_name", db_name)
        if account_privilege_detail is not None:
            pulumi.set(__self__, "account_privilege_detail", account_privilege_detail)

    @property
    @pulumi.getter(name="accountPrivilege")
    def account_privilege(self) -> pulumi.Input[str]:
        """
        Authorization database privilege types: 
        ReadWrite: Read and write privilege.
        ReadOnly: Read-only privilege.
        DDLOnly: Only DDL privilege.
        DMLOnly: Only DML privilege.
        Custom: Custom privilege.
        """
        return pulumi.get(self, "account_privilege")

    @account_privilege.setter
    def account_privilege(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_privilege", value)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Input[str]:
        """
        Database name requiring authorization.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter(name="accountPrivilegeDetail")
    def account_privilege_detail(self) -> Optional[pulumi.Input[str]]:
        """
        The specific SQL operation permissions contained in the permission type are separated by English commas (,) between multiple strings.
        When used as a request parameter in the CreateDatabase interface, when the AccountPrivilege value is Custom, this parameter is required. Value range (multiple selections allowed): SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER. When used as a return parameter in the DescribeDatabases interface, regardless of the value of AccountPrivilege, the details of the SQL operation permissions contained in this permission type are returned. For the specific SQL operation permissions contained in each permission type, please refer to the account permission list.
        """
        return pulumi.get(self, "account_privilege_detail")

    @account_privilege_detail.setter
    def account_privilege_detail(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_privilege_detail", value)


@pulumi.input_type
class BackupBackupPolicyArgs:
    def __init__(__self__, *,
                 backup_retention_period: pulumi.Input[int],
                 backup_time: pulumi.Input[str],
                 full_backup_period: pulumi.Input[str]):
        """
        :param pulumi.Input[int] backup_retention_period: Data backup retention period, value: 7 to 30 days.
        :param pulumi.Input[str] backup_time: The time for executing the backup task has an interval window of 2 hours and must be an even-hour time. Format: HH:mmZ-HH:mmZ (UTC time).
        :param pulumi.Input[str] full_backup_period: Full backup period. It is recommended to select at least 2 days per week for full backup. Multiple values are separated by English commas (,). Values: Monday: Monday. Tuesday: Tuesday. Wednesday: Wednesday. Thursday: Thursday. Friday: Friday. Saturday: Saturday. Sunday: Sunday.
        """
        pulumi.set(__self__, "backup_retention_period", backup_retention_period)
        pulumi.set(__self__, "backup_time", backup_time)
        pulumi.set(__self__, "full_backup_period", full_backup_period)

    @property
    @pulumi.getter(name="backupRetentionPeriod")
    def backup_retention_period(self) -> pulumi.Input[int]:
        """
        Data backup retention period, value: 7 to 30 days.
        """
        return pulumi.get(self, "backup_retention_period")

    @backup_retention_period.setter
    def backup_retention_period(self, value: pulumi.Input[int]):
        pulumi.set(self, "backup_retention_period", value)

    @property
    @pulumi.getter(name="backupTime")
    def backup_time(self) -> pulumi.Input[str]:
        """
        The time for executing the backup task has an interval window of 2 hours and must be an even-hour time. Format: HH:mmZ-HH:mmZ (UTC time).
        """
        return pulumi.get(self, "backup_time")

    @backup_time.setter
    def backup_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_time", value)

    @property
    @pulumi.getter(name="fullBackupPeriod")
    def full_backup_period(self) -> pulumi.Input[str]:
        """
        Full backup period. It is recommended to select at least 2 days per week for full backup. Multiple values are separated by English commas (,). Values: Monday: Monday. Tuesday: Tuesday. Wednesday: Wednesday. Thursday: Thursday. Friday: Friday. Saturday: Saturday. Sunday: Sunday.
        """
        return pulumi.get(self, "full_backup_period")

    @full_backup_period.setter
    def full_backup_period(self, value: pulumi.Input[str]):
        pulumi.set(self, "full_backup_period", value)


@pulumi.input_type
class InstanceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: The Key of Tags.
        :param pulumi.Input[str] value: The Value of Tags.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The Key of Tags.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class InstancesTagArgs:
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: The Key of Tags.
        :param str value: The Value of Tags.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The Key of Tags.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetInstancesTagArgs:
    def __init__(__self__, *,
                 key: str,
                 value: str):
        """
        :param str key: The Key of Tags.
        :param str value: The Value of Tags.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        The Key of Tags.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: str):
        pulumi.set(self, "value", value)


