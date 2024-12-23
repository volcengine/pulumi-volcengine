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
                 c_type: Optional[pulumi.Input[str]] = None,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 collate: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Database resource.
        :param pulumi.Input[str] db_name: The name of database.
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        :param pulumi.Input[str] c_type: Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
        :param pulumi.Input[str] character_set_name: Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
        :param pulumi.Input[str] collate: The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
        :param pulumi.Input[str] owner: The owner of database.
        """
        pulumi.set(__self__, "db_name", db_name)
        pulumi.set(__self__, "instance_id", instance_id)
        if c_type is not None:
            pulumi.set(__self__, "c_type", c_type)
        if character_set_name is not None:
            pulumi.set(__self__, "character_set_name", character_set_name)
        if collate is not None:
            pulumi.set(__self__, "collate", collate)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Input[str]:
        """
        The name of database.
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
    @pulumi.getter(name="cType")
    def c_type(self) -> Optional[pulumi.Input[str]]:
        """
        Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
        """
        return pulumi.get(self, "c_type")

    @c_type.setter
    def c_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "c_type", value)

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
        """
        return pulumi.get(self, "character_set_name")

    @character_set_name.setter
    def character_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "character_set_name", value)

    @property
    @pulumi.getter
    def collate(self) -> Optional[pulumi.Input[str]]:
        """
        The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
        """
        return pulumi.get(self, "collate")

    @collate.setter
    def collate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collate", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of database.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)


@pulumi.input_type
class _DatabaseState:
    def __init__(__self__, *,
                 c_type: Optional[pulumi.Input[str]] = None,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 collate: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 db_status: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Database resources.
        :param pulumi.Input[str] c_type: Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
        :param pulumi.Input[str] character_set_name: Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
        :param pulumi.Input[str] collate: The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
        :param pulumi.Input[str] db_name: The name of database.
        :param pulumi.Input[str] db_status: The status of the RDS database.
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        :param pulumi.Input[str] owner: The owner of database.
        """
        if c_type is not None:
            pulumi.set(__self__, "c_type", c_type)
        if character_set_name is not None:
            pulumi.set(__self__, "character_set_name", character_set_name)
        if collate is not None:
            pulumi.set(__self__, "collate", collate)
        if db_name is not None:
            pulumi.set(__self__, "db_name", db_name)
        if db_status is not None:
            pulumi.set(__self__, "db_status", db_status)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)

    @property
    @pulumi.getter(name="cType")
    def c_type(self) -> Optional[pulumi.Input[str]]:
        """
        Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
        """
        return pulumi.get(self, "c_type")

    @c_type.setter
    def c_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "c_type", value)

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
        """
        return pulumi.get(self, "character_set_name")

    @character_set_name.setter
    def character_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "character_set_name", value)

    @property
    @pulumi.getter
    def collate(self) -> Optional[pulumi.Input[str]]:
        """
        The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
        """
        return pulumi.get(self, "collate")

    @collate.setter
    def collate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "collate", value)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of database.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter(name="dbStatus")
    def db_status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the RDS database.
        """
        return pulumi.get(self, "db_status")

    @db_status.setter
    def db_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "db_status", value)

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

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        The owner of database.
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)


class Database(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 c_type: Optional[pulumi.Input[str]] = None,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 collate: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage rds postgresql database
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.rds_postgresql.Database("foo",
            c_type="C",
            collate="zh_CN.utf8",
            db_name="acc-test",
            instance_id="postgres-95*******233")
        ```

        ## Import

        Database can be imported using the instanceId:dbName, e.g.

        ```sh
        $ pulumi import volcengine:rds_postgresql/database:Database default postgres-ca7b7019****:dbname
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] c_type: Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
        :param pulumi.Input[str] character_set_name: Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
        :param pulumi.Input[str] collate: The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
        :param pulumi.Input[str] db_name: The name of database.
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        :param pulumi.Input[str] owner: The owner of database.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage rds postgresql database
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.rds_postgresql.Database("foo",
            c_type="C",
            collate="zh_CN.utf8",
            db_name="acc-test",
            instance_id="postgres-95*******233")
        ```

        ## Import

        Database can be imported using the instanceId:dbName, e.g.

        ```sh
        $ pulumi import volcengine:rds_postgresql/database:Database default postgres-ca7b7019****:dbname
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
                 c_type: Optional[pulumi.Input[str]] = None,
                 character_set_name: Optional[pulumi.Input[str]] = None,
                 collate: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseArgs.__new__(DatabaseArgs)

            __props__.__dict__["c_type"] = c_type
            __props__.__dict__["character_set_name"] = character_set_name
            __props__.__dict__["collate"] = collate
            if db_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_name'")
            __props__.__dict__["db_name"] = db_name
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["owner"] = owner
            __props__.__dict__["db_status"] = None
        super(Database, __self__).__init__(
            'volcengine:rds_postgresql/database:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            c_type: Optional[pulumi.Input[str]] = None,
            character_set_name: Optional[pulumi.Input[str]] = None,
            collate: Optional[pulumi.Input[str]] = None,
            db_name: Optional[pulumi.Input[str]] = None,
            db_status: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None) -> 'Database':
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] c_type: Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
        :param pulumi.Input[str] character_set_name: Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
        :param pulumi.Input[str] collate: The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
        :param pulumi.Input[str] db_name: The name of database.
        :param pulumi.Input[str] db_status: The status of the RDS database.
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        :param pulumi.Input[str] owner: The owner of database.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseState.__new__(_DatabaseState)

        __props__.__dict__["c_type"] = c_type
        __props__.__dict__["character_set_name"] = character_set_name
        __props__.__dict__["collate"] = collate
        __props__.__dict__["db_name"] = db_name
        __props__.__dict__["db_status"] = db_status
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["owner"] = owner
        return Database(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cType")
    def c_type(self) -> pulumi.Output[str]:
        """
        Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.
        """
        return pulumi.get(self, "c_type")

    @property
    @pulumi.getter(name="characterSetName")
    def character_set_name(self) -> pulumi.Output[str]:
        """
        Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.
        """
        return pulumi.get(self, "character_set_name")

    @property
    @pulumi.getter
    def collate(self) -> pulumi.Output[str]:
        """
        The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.
        """
        return pulumi.get(self, "collate")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Output[str]:
        """
        The name of database.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter(name="dbStatus")
    def db_status(self) -> pulumi.Output[str]:
        """
        The status of the RDS database.
        """
        return pulumi.get(self, "db_status")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        The owner of database.
        """
        return pulumi.get(self, "owner")

