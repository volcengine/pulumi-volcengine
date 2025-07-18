# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DatabaseArgs', 'Database']

@pulumi.input_type
class DatabaseArgs:
    def __init__(__self__, *,
                 db_name: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 db_desc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Database resource.
        :param pulumi.Input[str] db_name: Name database.
               illustrate:
               Unique name.
               The length is 2~64 characters.
               Start with a letter and end with a letter or number.
               Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
               Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        :param pulumi.Input[str] character_set_name: Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        :param pulumi.Input[str] db_desc: The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
        """
        pulumi.set(__self__, "db_name", db_name)
        pulumi.set(__self__, "instance_id", instance_id)
        if character_set_name is not None:
            pulumi.set(__self__, "character_set_name", character_set_name)
        if db_desc is not None:
            pulumi.set(__self__, "db_desc", db_desc)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Input[str]:
        """
        Name database.
        illustrate:
        Unique name.
        The length is 2~64 characters.
        Start with a letter and end with a letter or number.
        Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
        Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        """
        return pulumi.get(self, "character_set_name")

    @character_set_name.setter
    def character_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "character_set_name", value)

    @property
    @pulumi.getter(name="dbDesc")
    def db_desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
        """
        return pulumi.get(self, "db_desc")

    @db_desc.setter
    def db_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_desc", value)


@pulumi.input_type
class _DatabaseState:
    def __init__(__self__, *,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 db_desc: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Database resources.
        :param pulumi.Input[str] character_set_name: Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        :param pulumi.Input[str] db_desc: The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
        :param pulumi.Input[str] db_name: Name database.
               illustrate:
               Unique name.
               The length is 2~64 characters.
               Start with a letter and end with a letter or number.
               Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
               Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        """
        if character_set_name is not None:
            pulumi.set(__self__, "character_set_name", character_set_name)
        if db_desc is not None:
            pulumi.set(__self__, "db_desc", db_desc)
        if db_name is not None:
            pulumi.set(__self__, "db_name", db_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        """
        return pulumi.get(self, "character_set_name")

    @character_set_name.setter
    def character_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "character_set_name", value)

    @property
    @pulumi.getter(name="dbDesc")
    def db_desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
        """
        return pulumi.get(self, "db_desc")

    @db_desc.setter
    def db_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_desc", value)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name database.
        illustrate:
        Unique name.
        The length is 2~64 characters.
        Start with a letter and end with a letter or number.
        Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
        Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class Database(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 db_desc: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage rds mysql database
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.get_zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-project1",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-subnet-test-2",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        foo_instance = volcengine.rds_mysql.Instance("fooInstance",
            db_engine_version="MySQL_5_7",
            node_spec="rds.mysql.1c2g",
            primary_zone_id=foo_zones.zones[0].id,
            secondary_zone_id=foo_zones.zones[0].id,
            storage_space=80,
            subnet_id=foo_subnet.id,
            instance_name="acc-test",
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
            db_name="acc-test",
            instance_id=foo_instance.id,
            db_desc="test-update")
        ```

        ## Import

        Database can be imported using the instanceId:dbName, e.g.

        ```sh
        $ pulumi import volcengine:rds_mysql/database:Database default mysql-42b38c769c4b:dbname
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] character_set_name: Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        :param pulumi.Input[str] db_desc: The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
        :param pulumi.Input[str] db_name: Name database.
               illustrate:
               Unique name.
               The length is 2~64 characters.
               Start with a letter and end with a letter or number.
               Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
               Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage rds mysql database
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_zones = volcengine.ecs.get_zones()
        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-project1",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-subnet-test-2",
            cidr_block="172.16.0.0/24",
            zone_id=foo_zones.zones[0].id,
            vpc_id=foo_vpc.id)
        foo_instance = volcengine.rds_mysql.Instance("fooInstance",
            db_engine_version="MySQL_5_7",
            node_spec="rds.mysql.1c2g",
            primary_zone_id=foo_zones.zones[0].id,
            secondary_zone_id=foo_zones.zones[0].id,
            storage_space=80,
            subnet_id=foo_subnet.id,
            instance_name="acc-test",
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
            db_name="acc-test",
            instance_id=foo_instance.id,
            db_desc="test-update")
        ```

        ## Import

        Database can be imported using the instanceId:dbName, e.g.

        ```sh
        $ pulumi import volcengine:rds_mysql/database:Database default mysql-42b38c769c4b:dbname
        ```

        :param str resource_name: The name of the resource.
        :param DatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 db_desc: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseArgs.__new__(DatabaseArgs)

            __props__.__dict__["character_set_name"] = character_set_name
            __props__.__dict__["db_desc"] = db_desc
            if db_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_name'")
            __props__.__dict__["db_name"] = db_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(Database, __self__).__init__(
            'volcengine:rds_mysql/database:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            character_set_name: Optional[pulumi.Input[str]] = None,
            db_desc: Optional[pulumi.Input[str]] = None,
            db_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'Database':
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] character_set_name: Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        :param pulumi.Input[str] db_desc: The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
        :param pulumi.Input[str] db_name: Name database.
               illustrate:
               Unique name.
               The length is 2~64 characters.
               Start with a letter and end with a letter or number.
               Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
               Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseState.__new__(_DatabaseState)

        __props__.__dict__["character_set_name"] = character_set_name
        __props__.__dict__["db_desc"] = db_desc
        __props__.__dict__["db_name"] = db_name
        __props__.__dict__["instance_id"] = instance_id
        return Database(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> pulumi.Output[Optional[str]]:
        """
        Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        """
        return pulumi.get(self, "character_set_name")

    @property
    @pulumi.getter(name="dbDesc")
    def db_desc(self) -> pulumi.Output[Optional[str]]:
        """
        The description information of the database, with a length not exceeding 256 characters. This field is optional. If this field is not set, or if this field is set but the length of the description information is 0, then the description information is empty.
        """
        return pulumi.get(self, "db_desc")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Output[str]:
        """
        Name database.
        illustrate:
        Unique name.
        The length is 2~64 characters.
        Start with a letter and end with a letter or number.
        Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
        Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

