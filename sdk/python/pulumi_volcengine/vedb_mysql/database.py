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
                 character_set_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Database resource.
        :param pulumi.Input[str] db_name: The name of the database. Naming rules:
               Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
               Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
               The name cannot contain certain reserved words.
        :param pulumi.Input[str] instance_id: The id of the instance.
        :param pulumi.Input[str] character_set_name: Database character set: utf8mb4 (default), utf8, latin1, ascii.
        """
        pulumi.set(__self__, "db_name", db_name)
        pulumi.set(__self__, "instance_id", instance_id)
        if character_set_name is not None:
            pulumi.set(__self__, "character_set_name", character_set_name)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Input[str]:
        """
        The name of the database. Naming rules:
        Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
        Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
        The name cannot contain certain reserved words.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_name", value)

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
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database character set: utf8mb4 (default), utf8, latin1, ascii.
        """
        return pulumi.get(self, "character_set_name")

    @character_set_name.setter
    def character_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "character_set_name", value)


@pulumi.input_type
class _DatabaseState:
    def __init__(__self__, *,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Database resources.
        :param pulumi.Input[str] character_set_name: Database character set: utf8mb4 (default), utf8, latin1, ascii.
        :param pulumi.Input[str] db_name: The name of the database. Naming rules:
               Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
               Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
               The name cannot contain certain reserved words.
        :param pulumi.Input[str] instance_id: The id of the instance.
        """
        if character_set_name is not None:
            pulumi.set(__self__, "character_set_name", character_set_name)
        if db_name is not None:
            pulumi.set(__self__, "db_name", db_name)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database character set: utf8mb4 (default), utf8, latin1, ascii.
        """
        return pulumi.get(self, "character_set_name")

    @character_set_name.setter
    def character_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "character_set_name", value)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database. Naming rules:
        Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
        Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
        The name cannot contain certain reserved words.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_name", value)

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


class Database(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage vedb mysql database
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
        foo_database = volcengine.vedb_mysql.Database("fooDatabase",
            db_name="tf-table",
            instance_id=foo_instance.id)
        ```

        ## Import

        VedbMysqlDatabase can be imported using the instance id and database name, e.g.

        ```sh
        $ pulumi import volcengine:vedb_mysql/database:Database default vedbm-r3xq0zdl****:testdb
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] character_set_name: Database character set: utf8mb4 (default), utf8, latin1, ascii.
        :param pulumi.Input[str] db_name: The name of the database. Naming rules:
               Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
               Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
               The name cannot contain certain reserved words.
        :param pulumi.Input[str] instance_id: The id of the instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage vedb mysql database
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
        foo_database = volcengine.vedb_mysql.Database("fooDatabase",
            db_name="tf-table",
            instance_id=foo_instance.id)
        ```

        ## Import

        VedbMysqlDatabase can be imported using the instance id and database name, e.g.

        ```sh
        $ pulumi import volcengine:vedb_mysql/database:Database default vedbm-r3xq0zdl****:testdb
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
            if db_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_name'")
            __props__.__dict__["db_name"] = db_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(Database, __self__).__init__(
            'volcengine:vedb_mysql/database:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            character_set_name: Optional[pulumi.Input[str]] = None,
            db_name: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'Database':
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] character_set_name: Database character set: utf8mb4 (default), utf8, latin1, ascii.
        :param pulumi.Input[str] db_name: The name of the database. Naming rules:
               Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
               Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
               The name cannot contain certain reserved words.
        :param pulumi.Input[str] instance_id: The id of the instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseState.__new__(_DatabaseState)

        __props__.__dict__["character_set_name"] = character_set_name
        __props__.__dict__["db_name"] = db_name
        __props__.__dict__["instance_id"] = instance_id
        return Database(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> pulumi.Output[str]:
        """
        Database character set: utf8mb4 (default), utf8, latin1, ascii.
        """
        return pulumi.get(self, "character_set_name")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Output[str]:
        """
        The name of the database. Naming rules:
        Unique name. Start with a lowercase letter and end with a letter or number. The length is within 2 to 64 characters.
        Consist of lowercase letters, numbers, underscores (_), or hyphens (-).
        The name cannot contain certain reserved words.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The id of the instance.
        """
        return pulumi.get(self, "instance_id")
