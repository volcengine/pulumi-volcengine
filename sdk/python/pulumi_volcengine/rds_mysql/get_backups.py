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
    'GetBackupsResult',
    'AwaitableGetBackupsResult',
    'get_backups',
    'get_backups_output',
]

@pulumi.output_type
class GetBackupsResult:
    """
    A collection of values returned by getBackups.
    """
    def __init__(__self__, backup_end_time=None, backup_id=None, backup_method=None, backup_start_time=None, backup_status=None, backup_type=None, backups=None, create_type=None, id=None, instance_id=None, output_file=None, total_count=None):
        if backup_end_time and not isinstance(backup_end_time, str):
            raise TypeError("Expected argument 'backup_end_time' to be a str")
        pulumi.set(__self__, "backup_end_time", backup_end_time)
        if backup_id and not isinstance(backup_id, str):
            raise TypeError("Expected argument 'backup_id' to be a str")
        pulumi.set(__self__, "backup_id", backup_id)
        if backup_method and not isinstance(backup_method, str):
            raise TypeError("Expected argument 'backup_method' to be a str")
        pulumi.set(__self__, "backup_method", backup_method)
        if backup_start_time and not isinstance(backup_start_time, str):
            raise TypeError("Expected argument 'backup_start_time' to be a str")
        pulumi.set(__self__, "backup_start_time", backup_start_time)
        if backup_status and not isinstance(backup_status, str):
            raise TypeError("Expected argument 'backup_status' to be a str")
        pulumi.set(__self__, "backup_status", backup_status)
        if backup_type and not isinstance(backup_type, str):
            raise TypeError("Expected argument 'backup_type' to be a str")
        pulumi.set(__self__, "backup_type", backup_type)
        if backups and not isinstance(backups, list):
            raise TypeError("Expected argument 'backups' to be a list")
        pulumi.set(__self__, "backups", backups)
        if create_type and not isinstance(create_type, str):
            raise TypeError("Expected argument 'create_type' to be a str")
        pulumi.set(__self__, "create_type", create_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="backupEndTime")
    def backup_end_time(self) -> Optional[str]:
        """
        The end time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
        """
        return pulumi.get(self, "backup_end_time")

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> Optional[str]:
        """
        The id of the backup.
        """
        return pulumi.get(self, "backup_id")

    @property
    @pulumi.getter(name="backupMethod")
    def backup_method(self) -> Optional[str]:
        """
        Backup type, value: Physical: Physical backup. Logical: Logical backup.
        """
        return pulumi.get(self, "backup_method")

    @property
    @pulumi.getter(name="backupStartTime")
    def backup_start_time(self) -> Optional[str]:
        """
        The start time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).
        """
        return pulumi.get(self, "backup_start_time")

    @property
    @pulumi.getter(name="backupStatus")
    def backup_status(self) -> Optional[str]:
        """
        Backup status, values: Success. Failed. Running.
        """
        return pulumi.get(self, "backup_status")

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[str]:
        """
        Backup method, values:
        Full: Full backup under physical backup type or library table backup under logical backup type.
        Increment: Incremental backup under physical backup type (created by the system).
        DumpAll: Full database backup under logical backup type.
        Description:
        There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter
    def backups(self) -> Sequence['outputs.GetBackupsBackupResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "backups")

    @property
    @pulumi.getter(name="createType")
    def create_type(self) -> Optional[str]:
        """
        Creator of backup. Values: System. User.
        """
        return pulumi.get(self, "create_type")

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


class AwaitableGetBackupsResult(GetBackupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackupsResult(
            backup_end_time=self.backup_end_time,
            backup_id=self.backup_id,
            backup_method=self.backup_method,
            backup_start_time=self.backup_start_time,
            backup_status=self.backup_status,
            backup_type=self.backup_type,
            backups=self.backups,
            create_type=self.create_type,
            id=self.id,
            instance_id=self.instance_id,
            output_file=self.output_file,
            total_count=self.total_count)


def get_backups(backup_end_time: Optional[str] = None,
                backup_id: Optional[str] = None,
                backup_method: Optional[str] = None,
                backup_start_time: Optional[str] = None,
                backup_status: Optional[str] = None,
                backup_type: Optional[str] = None,
                create_type: Optional[str] = None,
                instance_id: Optional[str] = None,
                output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackupsResult:
    """
    Use this data source to query detailed information of rds mysql backups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.rds_mysql.get_backups(instance_id="mysql-b51d37110dd1")
    ```


    :param str backup_end_time: The end time of the backup.
    :param str backup_id: The id of the backup.
    :param str backup_method: Backup type, value: Physical: Physical backup. Default value. Logical: Logical backup. Description: There is no default value. When this field is not passed, backups of all states under the query conditions limited by other fields are returned.
    :param str backup_start_time: The start time of the backup.
    :param str backup_status: Backup status, values: Success: Success. Failed: Failed. Running: In progress. Description: There is no default value. When this field is not passed, all backups in all states under the query conditions limited by other fields are returned.
    :param str backup_type: Backup method, value: Full: Full backup under physical backup type or library table backup under logical backup type. Increment: Incremental backup under physical backup type. DumpAll: Full database backup under logical backup type. Description: There is no default value. When this field is not passed, all backups of all methods under the query conditions limited by other fields are returned.
    :param str create_type: Creator of backup. Values: System: System. User: User. Description: There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
    :param str instance_id: The id of the instance.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['backupEndTime'] = backup_end_time
    __args__['backupId'] = backup_id
    __args__['backupMethod'] = backup_method
    __args__['backupStartTime'] = backup_start_time
    __args__['backupStatus'] = backup_status
    __args__['backupType'] = backup_type
    __args__['createType'] = create_type
    __args__['instanceId'] = instance_id
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:rds_mysql/getBackups:getBackups', __args__, opts=opts, typ=GetBackupsResult).value

    return AwaitableGetBackupsResult(
        backup_end_time=pulumi.get(__ret__, 'backup_end_time'),
        backup_id=pulumi.get(__ret__, 'backup_id'),
        backup_method=pulumi.get(__ret__, 'backup_method'),
        backup_start_time=pulumi.get(__ret__, 'backup_start_time'),
        backup_status=pulumi.get(__ret__, 'backup_status'),
        backup_type=pulumi.get(__ret__, 'backup_type'),
        backups=pulumi.get(__ret__, 'backups'),
        create_type=pulumi.get(__ret__, 'create_type'),
        id=pulumi.get(__ret__, 'id'),
        instance_id=pulumi.get(__ret__, 'instance_id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_backups)
def get_backups_output(backup_end_time: Optional[pulumi.Input[Optional[str]]] = None,
                       backup_id: Optional[pulumi.Input[Optional[str]]] = None,
                       backup_method: Optional[pulumi.Input[Optional[str]]] = None,
                       backup_start_time: Optional[pulumi.Input[Optional[str]]] = None,
                       backup_status: Optional[pulumi.Input[Optional[str]]] = None,
                       backup_type: Optional[pulumi.Input[Optional[str]]] = None,
                       create_type: Optional[pulumi.Input[Optional[str]]] = None,
                       instance_id: Optional[pulumi.Input[Optional[str]]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackupsResult]:
    """
    Use this data source to query detailed information of rds mysql backups
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.rds_mysql.get_backups(instance_id="mysql-b51d37110dd1")
    ```


    :param str backup_end_time: The end time of the backup.
    :param str backup_id: The id of the backup.
    :param str backup_method: Backup type, value: Physical: Physical backup. Default value. Logical: Logical backup. Description: There is no default value. When this field is not passed, backups of all states under the query conditions limited by other fields are returned.
    :param str backup_start_time: The start time of the backup.
    :param str backup_status: Backup status, values: Success: Success. Failed: Failed. Running: In progress. Description: There is no default value. When this field is not passed, all backups in all states under the query conditions limited by other fields are returned.
    :param str backup_type: Backup method, value: Full: Full backup under physical backup type or library table backup under logical backup type. Increment: Incremental backup under physical backup type. DumpAll: Full database backup under logical backup type. Description: There is no default value. When this field is not passed, all backups of all methods under the query conditions limited by other fields are returned.
    :param str create_type: Creator of backup. Values: System: System. User: User. Description: There is no default value. When this field is not passed, all types of backups under the query conditions limited by other fields are returned.
    :param str instance_id: The id of the instance.
    :param str output_file: File name where to save data source results.
    """
    ...
