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

__all__ = ['BackupArgs', 'Backup']

@pulumi.input_type
class BackupArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 backup_method: Optional[pulumi.Input[str]] = None,
                 backup_policy: Optional[pulumi.Input['BackupBackupPolicyArgs']] = None,
                 backup_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Backup resource.
        :param pulumi.Input[str] instance_id: The id of the instance.
        :param pulumi.Input[str] backup_method: Backup method. Currently, only physical backup is supported. The value is Physical.
        :param pulumi.Input['BackupBackupPolicyArgs'] backup_policy: Data backup strategy for instances.
        :param pulumi.Input[str] backup_type: Backup type. Currently, only full backup is supported. The value is Full.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        if backup_method is not None:
            pulumi.set(__self__, "backup_method", backup_method)
        if backup_policy is not None:
            pulumi.set(__self__, "backup_policy", backup_policy)
        if backup_type is not None:
            pulumi.set(__self__, "backup_type", backup_type)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The id of the instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="backupMethod")
    def backup_method(self) -> Optional[pulumi.Input[str]]:
        """
        Backup method. Currently, only physical backup is supported. The value is Physical.
        """
        return pulumi.get(self, "backup_method")

    @backup_method.setter
    def backup_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_method", value)

    @property
    @pulumi.getter(name="backupPolicy")
    def backup_policy(self) -> Optional[pulumi.Input['BackupBackupPolicyArgs']]:
        """
        Data backup strategy for instances.
        """
        return pulumi.get(self, "backup_policy")

    @backup_policy.setter
    def backup_policy(self, value: Optional[pulumi.Input['BackupBackupPolicyArgs']]):
        pulumi.set(self, "backup_policy", value)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[pulumi.Input[str]]:
        """
        Backup type. Currently, only full backup is supported. The value is Full.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_type", value)


@pulumi.input_type
class _BackupState:
    def __init__(__self__, *,
                 backup_id: Optional[pulumi.Input[str]] = None,
                 backup_method: Optional[pulumi.Input[str]] = None,
                 backup_policy: Optional[pulumi.Input['BackupBackupPolicyArgs']] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Backup resources.
        :param pulumi.Input[str] backup_id: The id of the backup.
        :param pulumi.Input[str] backup_method: Backup method. Currently, only physical backup is supported. The value is Physical.
        :param pulumi.Input['BackupBackupPolicyArgs'] backup_policy: Data backup strategy for instances.
        :param pulumi.Input[str] backup_type: Backup type. Currently, only full backup is supported. The value is Full.
        :param pulumi.Input[str] instance_id: The id of the instance.
        """
        if backup_id is not None:
            pulumi.set(__self__, "backup_id", backup_id)
        if backup_method is not None:
            pulumi.set(__self__, "backup_method", backup_method)
        if backup_policy is not None:
            pulumi.set(__self__, "backup_policy", backup_policy)
        if backup_type is not None:
            pulumi.set(__self__, "backup_type", backup_type)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the backup.
        """
        return pulumi.get(self, "backup_id")

    @backup_id.setter
    def backup_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_id", value)

    @property
    @pulumi.getter(name="backupMethod")
    def backup_method(self) -> Optional[pulumi.Input[str]]:
        """
        Backup method. Currently, only physical backup is supported. The value is Physical.
        """
        return pulumi.get(self, "backup_method")

    @backup_method.setter
    def backup_method(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_method", value)

    @property
    @pulumi.getter(name="backupPolicy")
    def backup_policy(self) -> Optional[pulumi.Input['BackupBackupPolicyArgs']]:
        """
        Data backup strategy for instances.
        """
        return pulumi.get(self, "backup_policy")

    @backup_policy.setter
    def backup_policy(self, value: Optional[pulumi.Input['BackupBackupPolicyArgs']]):
        pulumi.set(self, "backup_policy", value)

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> Optional[pulumi.Input[str]]:
        """
        Backup type. Currently, only full backup is supported. The value is Full.
        """
        return pulumi.get(self, "backup_type")

    @backup_type.setter
    def backup_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_type", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class Backup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_method: Optional[pulumi.Input[str]] = None,
                 backup_policy: Optional[pulumi.Input[pulumi.InputType['BackupBackupPolicyArgs']]] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage vedb mysql backup
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[2].id,
            vpc_id=foo_vpc.id)
        foo_instance = volcengine.vedb_mysql.Instance("fooInstance",
            charge_type="PostPaid",
            storage_charge_type="PostPaid",
            db_engine_version="MySQL_8_0",
            db_minor_version="3.0",
            node_number=2,
            node_spec="vedb.mysql.x4.large",
            subnet_id=foo_subnet.id,
            instance_name="tf-test",
            project_name="testA",
            tags=[
                volcengine.vedb_mysql.InstanceTagArgs(
                    key="tftest",
                    value="tftest",
                ),
                volcengine.vedb_mysql.InstanceTagArgs(
                    key="tftest2",
                    value="tftest2",
                ),
            ])
        foo_backup = volcengine.vedb_mysql.Backup("fooBackup",
            instance_id=foo_instance.id,
            backup_policy=volcengine.vedb_mysql.BackupBackupPolicyArgs(
                backup_time="18:00Z-20:00Z",
                full_backup_period="Monday,Tuesday,Wednesday",
                backup_retention_period=8,
            ))
        ```

        ## Import

        VedbMysqlBackup can be imported using the instance id and backup id, e.g.

        ```sh
        $ pulumi import volcengine:vedb_mysql/backup:Backup default instanceID:backupId
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_method: Backup method. Currently, only physical backup is supported. The value is Physical.
        :param pulumi.Input[pulumi.InputType['BackupBackupPolicyArgs']] backup_policy: Data backup strategy for instances.
        :param pulumi.Input[str] backup_type: Backup type. Currently, only full backup is supported. The value is Full.
        :param pulumi.Input[str] instance_id: The id of the instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage vedb mysql backup
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-vpc",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-test-subnet",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[2].id,
            vpc_id=foo_vpc.id)
        foo_instance = volcengine.vedb_mysql.Instance("fooInstance",
            charge_type="PostPaid",
            storage_charge_type="PostPaid",
            db_engine_version="MySQL_8_0",
            db_minor_version="3.0",
            node_number=2,
            node_spec="vedb.mysql.x4.large",
            subnet_id=foo_subnet.id,
            instance_name="tf-test",
            project_name="testA",
            tags=[
                volcengine.vedb_mysql.InstanceTagArgs(
                    key="tftest",
                    value="tftest",
                ),
                volcengine.vedb_mysql.InstanceTagArgs(
                    key="tftest2",
                    value="tftest2",
                ),
            ])
        foo_backup = volcengine.vedb_mysql.Backup("fooBackup",
            instance_id=foo_instance.id,
            backup_policy=volcengine.vedb_mysql.BackupBackupPolicyArgs(
                backup_time="18:00Z-20:00Z",
                full_backup_period="Monday,Tuesday,Wednesday",
                backup_retention_period=8,
            ))
        ```

        ## Import

        VedbMysqlBackup can be imported using the instance id and backup id, e.g.

        ```sh
        $ pulumi import volcengine:vedb_mysql/backup:Backup default instanceID:backupId
        ```

        :param str resource_name: The name of the resource.
        :param BackupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_method: Optional[pulumi.Input[str]] = None,
                 backup_policy: Optional[pulumi.Input[pulumi.InputType['BackupBackupPolicyArgs']]] = None,
                 backup_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupArgs.__new__(BackupArgs)

            __props__.__dict__["backup_method"] = backup_method
            __props__.__dict__["backup_policy"] = backup_policy
            __props__.__dict__["backup_type"] = backup_type
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["backup_id"] = None
        super(Backup, __self__).__init__(
            'volcengine:vedb_mysql/backup:Backup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_id: Optional[pulumi.Input[str]] = None,
            backup_method: Optional[pulumi.Input[str]] = None,
            backup_policy: Optional[pulumi.Input[pulumi.InputType['BackupBackupPolicyArgs']]] = None,
            backup_type: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'Backup':
        """
        Get an existing Backup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_id: The id of the backup.
        :param pulumi.Input[str] backup_method: Backup method. Currently, only physical backup is supported. The value is Physical.
        :param pulumi.Input[pulumi.InputType['BackupBackupPolicyArgs']] backup_policy: Data backup strategy for instances.
        :param pulumi.Input[str] backup_type: Backup type. Currently, only full backup is supported. The value is Full.
        :param pulumi.Input[str] instance_id: The id of the instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupState.__new__(_BackupState)

        __props__.__dict__["backup_id"] = backup_id
        __props__.__dict__["backup_method"] = backup_method
        __props__.__dict__["backup_policy"] = backup_policy
        __props__.__dict__["backup_type"] = backup_type
        __props__.__dict__["instance_id"] = instance_id
        return Backup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> pulumi.Output[str]:
        """
        The id of the backup.
        """
        return pulumi.get(self, "backup_id")

    @property
    @pulumi.getter(name="backupMethod")
    def backup_method(self) -> pulumi.Output[str]:
        """
        Backup method. Currently, only physical backup is supported. The value is Physical.
        """
        return pulumi.get(self, "backup_method")

    @property
    @pulumi.getter(name="backupPolicy")
    def backup_policy(self) -> pulumi.Output['outputs.BackupBackupPolicy']:
        """
        Data backup strategy for instances.
        """
        return pulumi.get(self, "backup_policy")

    @property
    @pulumi.getter(name="backupType")
    def backup_type(self) -> pulumi.Output[str]:
        """
        Backup type. Currently, only full backup is supported. The value is Full.
        """
        return pulumi.get(self, "backup_type")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The id of the instance.
        """
        return pulumi.get(self, "instance_id")
