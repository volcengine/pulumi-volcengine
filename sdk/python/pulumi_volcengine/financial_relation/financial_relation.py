# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FinancialRelationArgs', 'FinancialRelation']

@pulumi.input_type
class FinancialRelationArgs:
    def __init__(__self__, *,
                 sub_account_id: pulumi.Input[int],
                 account_alias: Optional[pulumi.Input[str]] = None,
                 auth_lists: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 relation: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a FinancialRelation resource.
        :param pulumi.Input[int] sub_account_id: The sub account id.
        :param pulumi.Input[str] account_alias: The display name of the sub account.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] auth_lists: The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
        :param pulumi.Input[int] relation: The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
        """
        pulumi.set(__self__, "sub_account_id", sub_account_id)
        if account_alias is not None:
            pulumi.set(__self__, "account_alias", account_alias)
        if auth_lists is not None:
            pulumi.set(__self__, "auth_lists", auth_lists)
        if relation is not None:
            pulumi.set(__self__, "relation", relation)

    @property
    @pulumi.getter(name="subAccountId")
    def sub_account_id(self) -> pulumi.Input[int]:
        """
        The sub account id.
        """
        return pulumi.get(self, "sub_account_id")

    @sub_account_id.setter
    def sub_account_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "sub_account_id", value)

    @property
    @pulumi.getter(name="accountAlias")
    def account_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the sub account.
        """
        return pulumi.get(self, "account_alias")

    @account_alias.setter
    def account_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_alias", value)

    @property
    @pulumi.getter(name="authLists")
    def auth_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
        """
        return pulumi.get(self, "auth_lists")

    @auth_lists.setter
    def auth_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "auth_lists", value)

    @property
    @pulumi.getter
    def relation(self) -> Optional[pulumi.Input[int]]:
        """
        The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
        """
        return pulumi.get(self, "relation")

    @relation.setter
    def relation(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "relation", value)


@pulumi.input_type
class _FinancialRelationState:
    def __init__(__self__, *,
                 account_alias: Optional[pulumi.Input[str]] = None,
                 auth_lists: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 relation: Optional[pulumi.Input[int]] = None,
                 relation_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[int]] = None,
                 sub_account_id: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering FinancialRelation resources.
        :param pulumi.Input[str] account_alias: The display name of the sub account.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] auth_lists: The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
        :param pulumi.Input[int] relation: The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
        :param pulumi.Input[str] relation_id: The id of the financial relation.
        :param pulumi.Input[int] status: The status of the financial relation.
        :param pulumi.Input[int] sub_account_id: The sub account id.
        """
        if account_alias is not None:
            pulumi.set(__self__, "account_alias", account_alias)
        if auth_lists is not None:
            pulumi.set(__self__, "auth_lists", auth_lists)
        if relation is not None:
            pulumi.set(__self__, "relation", relation)
        if relation_id is not None:
            pulumi.set(__self__, "relation_id", relation_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if sub_account_id is not None:
            pulumi.set(__self__, "sub_account_id", sub_account_id)

    @property
    @pulumi.getter(name="accountAlias")
    def account_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the sub account.
        """
        return pulumi.get(self, "account_alias")

    @account_alias.setter
    def account_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_alias", value)

    @property
    @pulumi.getter(name="authLists")
    def auth_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]:
        """
        The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
        """
        return pulumi.get(self, "auth_lists")

    @auth_lists.setter
    def auth_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]]):
        pulumi.set(self, "auth_lists", value)

    @property
    @pulumi.getter
    def relation(self) -> Optional[pulumi.Input[int]]:
        """
        The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
        """
        return pulumi.get(self, "relation")

    @relation.setter
    def relation(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "relation", value)

    @property
    @pulumi.getter(name="relationId")
    def relation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the financial relation.
        """
        return pulumi.get(self, "relation_id")

    @relation_id.setter
    def relation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "relation_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[int]]:
        """
        The status of the financial relation.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subAccountId")
    def sub_account_id(self) -> Optional[pulumi.Input[int]]:
        """
        The sub account id.
        """
        return pulumi.get(self, "sub_account_id")

    @sub_account_id.setter
    def sub_account_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "sub_account_id", value)


class FinancialRelation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_alias: Optional[pulumi.Input[str]] = None,
                 auth_lists: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 relation: Optional[pulumi.Input[int]] = None,
                 sub_account_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to manage financial relation
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.financial_relation.FinancialRelation("foo",
            account_alias="acc-test-financial",
            auth_lists=[
                1,
                2,
                3,
            ],
            relation=4,
            sub_account_id=2100260000)
        ```

        ## Import

        FinancialRelation can be imported using the sub_account_id:relation:relation_id, e.g.

        ```sh
        $ pulumi import volcengine:financial_relation/financialRelation:FinancialRelation default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_alias: The display name of the sub account.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] auth_lists: The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
        :param pulumi.Input[int] relation: The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
        :param pulumi.Input[int] sub_account_id: The sub account id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FinancialRelationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage financial relation
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.financial_relation.FinancialRelation("foo",
            account_alias="acc-test-financial",
            auth_lists=[
                1,
                2,
                3,
            ],
            relation=4,
            sub_account_id=2100260000)
        ```

        ## Import

        FinancialRelation can be imported using the sub_account_id:relation:relation_id, e.g.

        ```sh
        $ pulumi import volcengine:financial_relation/financialRelation:FinancialRelation default resource_id
        ```

        :param str resource_name: The name of the resource.
        :param FinancialRelationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FinancialRelationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_alias: Optional[pulumi.Input[str]] = None,
                 auth_lists: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
                 relation: Optional[pulumi.Input[int]] = None,
                 sub_account_id: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FinancialRelationArgs.__new__(FinancialRelationArgs)

            __props__.__dict__["account_alias"] = account_alias
            __props__.__dict__["auth_lists"] = auth_lists
            __props__.__dict__["relation"] = relation
            if sub_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'sub_account_id'")
            __props__.__dict__["sub_account_id"] = sub_account_id
            __props__.__dict__["relation_id"] = None
            __props__.__dict__["status"] = None
        super(FinancialRelation, __self__).__init__(
            'volcengine:financial_relation/financialRelation:FinancialRelation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_alias: Optional[pulumi.Input[str]] = None,
            auth_lists: Optional[pulumi.Input[Sequence[pulumi.Input[int]]]] = None,
            relation: Optional[pulumi.Input[int]] = None,
            relation_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[int]] = None,
            sub_account_id: Optional[pulumi.Input[int]] = None) -> 'FinancialRelation':
        """
        Get an existing FinancialRelation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_alias: The display name of the sub account.
        :param pulumi.Input[Sequence[pulumi.Input[int]]] auth_lists: The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
        :param pulumi.Input[int] relation: The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
        :param pulumi.Input[str] relation_id: The id of the financial relation.
        :param pulumi.Input[int] status: The status of the financial relation.
        :param pulumi.Input[int] sub_account_id: The sub account id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FinancialRelationState.__new__(_FinancialRelationState)

        __props__.__dict__["account_alias"] = account_alias
        __props__.__dict__["auth_lists"] = auth_lists
        __props__.__dict__["relation"] = relation
        __props__.__dict__["relation_id"] = relation_id
        __props__.__dict__["status"] = status
        __props__.__dict__["sub_account_id"] = sub_account_id
        return FinancialRelation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountAlias")
    def account_alias(self) -> pulumi.Output[str]:
        """
        The display name of the sub account.
        """
        return pulumi.get(self, "account_alias")

    @property
    @pulumi.getter(name="authLists")
    def auth_lists(self) -> pulumi.Output[Sequence[int]]:
        """
        The authorization list of financial management. This field is valid and required when the relation is 4. Valid value range is `1-5`.
        """
        return pulumi.get(self, "auth_lists")

    @property
    @pulumi.getter
    def relation(self) -> pulumi.Output[int]:
        """
        The relation of the financial. Valid values: `1`, `4`. `1` means financial custody, `4` means financial management.
        """
        return pulumi.get(self, "relation")

    @property
    @pulumi.getter(name="relationId")
    def relation_id(self) -> pulumi.Output[str]:
        """
        The id of the financial relation.
        """
        return pulumi.get(self, "relation_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[int]:
        """
        The status of the financial relation.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subAccountId")
    def sub_account_id(self) -> pulumi.Output[int]:
        """
        The sub account id.
        """
        return pulumi.get(self, "sub_account_id")

