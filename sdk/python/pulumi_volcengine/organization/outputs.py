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
    'AccountTag',
    'AccountsAccountResult',
    'AccountsAccountTagResult',
    'OrganizationsOrganizationResult',
    'ServiceControlPoliciesPolicyResult',
    'UnitsUnitResult',
]

@pulumi.output_type
class AccountTag(dict):
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

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class AccountsAccountResult(dict):
    def __init__(__self__, *,
                 account_id: str,
                 account_name: str,
                 allow_console: int,
                 allow_exit: int,
                 created_time: str,
                 delete_uk: str,
                 deleted_time: str,
                 description: str,
                 iam_role: str,
                 id: str,
                 is_owner: int,
                 join_type: int,
                 org_id: str,
                 org_type: int,
                 org_unit_id: str,
                 org_unit_name: str,
                 org_verification_id: str,
                 owner: str,
                 show_name: str,
                 tags: Sequence['outputs.AccountsAccountTagResult'],
                 updated_time: str):
        """
        :param str account_id: The id of the account.
        :param str account_name: The name of the account.
        :param int allow_console: Whether to allow the account enable console. `0` means allowed, `1` means not allowed.
        :param int allow_exit: Whether to allow exit the organization. `0` means allowed, `1` means not allowed.
        :param str created_time: The created time of the account.
        :param str delete_uk: The delete uk of the account.
        :param str deleted_time: The deleted time of the account.
        :param str description: The description of the account.
        :param str iam_role: The name of the iam role.
        :param str id: The id of the account.
        :param int is_owner: Whether the account is owner. `0` means not owner, `1` means owner.
        :param int join_type: The join type of the account. `0` means create, `1` means invitation.
        :param str org_id: The id of the organization.
        :param int org_type: The type of the organization. `1` means business organization.
        :param str org_unit_id: The id of the organization unit.
        :param str org_unit_name: The name of the organization unit.
        :param str org_verification_id: The id of the organization verification.
        :param str owner: The owner id of the account.
        :param str show_name: The show name of the account.
        :param Sequence['AccountsAccountTagArgs'] tags: Tags.
        :param str updated_time: The updated time of the account.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "allow_console", allow_console)
        pulumi.set(__self__, "allow_exit", allow_exit)
        pulumi.set(__self__, "created_time", created_time)
        pulumi.set(__self__, "delete_uk", delete_uk)
        pulumi.set(__self__, "deleted_time", deleted_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "iam_role", iam_role)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "is_owner", is_owner)
        pulumi.set(__self__, "join_type", join_type)
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "org_type", org_type)
        pulumi.set(__self__, "org_unit_id", org_unit_id)
        pulumi.set(__self__, "org_unit_name", org_unit_name)
        pulumi.set(__self__, "org_verification_id", org_verification_id)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "show_name", show_name)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "updated_time", updated_time)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The id of the account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        """
        The name of the account.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="allowConsole")
    def allow_console(self) -> int:
        """
        Whether to allow the account enable console. `0` means allowed, `1` means not allowed.
        """
        return pulumi.get(self, "allow_console")

    @property
    @pulumi.getter(name="allowExit")
    def allow_exit(self) -> int:
        """
        Whether to allow exit the organization. `0` means allowed, `1` means not allowed.
        """
        return pulumi.get(self, "allow_exit")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The created time of the account.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="deleteUk")
    def delete_uk(self) -> str:
        """
        The delete uk of the account.
        """
        return pulumi.get(self, "delete_uk")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> str:
        """
        The deleted time of the account.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the account.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="iamRole")
    def iam_role(self) -> str:
        """
        The name of the iam role.
        """
        return pulumi.get(self, "iam_role")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the account.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isOwner")
    def is_owner(self) -> int:
        """
        Whether the account is owner. `0` means not owner, `1` means owner.
        """
        return pulumi.get(self, "is_owner")

    @property
    @pulumi.getter(name="joinType")
    def join_type(self) -> int:
        """
        The join type of the account. `0` means create, `1` means invitation.
        """
        return pulumi.get(self, "join_type")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        """
        The id of the organization.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="orgType")
    def org_type(self) -> int:
        """
        The type of the organization. `1` means business organization.
        """
        return pulumi.get(self, "org_type")

    @property
    @pulumi.getter(name="orgUnitId")
    def org_unit_id(self) -> str:
        """
        The id of the organization unit.
        """
        return pulumi.get(self, "org_unit_id")

    @property
    @pulumi.getter(name="orgUnitName")
    def org_unit_name(self) -> str:
        """
        The name of the organization unit.
        """
        return pulumi.get(self, "org_unit_name")

    @property
    @pulumi.getter(name="orgVerificationId")
    def org_verification_id(self) -> str:
        """
        The id of the organization verification.
        """
        return pulumi.get(self, "org_verification_id")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The owner id of the account.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="showName")
    def show_name(self) -> str:
        """
        The show name of the account.
        """
        return pulumi.get(self, "show_name")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.AccountsAccountTagResult']:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The updated time of the account.
        """
        return pulumi.get(self, "updated_time")


@pulumi.output_type
class AccountsAccountTagResult(dict):
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

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The Value of Tags.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class OrganizationsOrganizationResult(dict):
    def __init__(__self__, *,
                 account_id: int,
                 account_name: str,
                 created_time: str,
                 delete_uk: str,
                 deleted_time: str,
                 description: str,
                 id: str,
                 main_name: str,
                 name: str,
                 owner: str,
                 status: int,
                 type: int,
                 updated_time: str):
        """
        :param int account_id: The account id of the organization owner.
        :param str account_name: The account name of the organization owner.
        :param str created_time: The created time of the organization.
        :param str delete_uk: The delete uk of the organization.
        :param str deleted_time: The deleted time of the organization.
        :param str description: The description of the organization.
        :param str id: The id of the organization.
        :param str main_name: The main name of the organization owner.
        :param str name: The name of the organization.
        :param str owner: The owner id of the organization.
        :param int status: The status of the organization.
        :param int type: The type of the organization.
        :param str updated_time: The updated time of the organization.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "created_time", created_time)
        pulumi.set(__self__, "delete_uk", delete_uk)
        pulumi.set(__self__, "deleted_time", deleted_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "main_name", main_name)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "updated_time", updated_time)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> int:
        """
        The account id of the organization owner.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> str:
        """
        The account name of the organization owner.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The created time of the organization.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="deleteUk")
    def delete_uk(self) -> str:
        """
        The delete uk of the organization.
        """
        return pulumi.get(self, "delete_uk")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> str:
        """
        The deleted time of the organization.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the organization.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the organization.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mainName")
    def main_name(self) -> str:
        """
        The main name of the organization owner.
        """
        return pulumi.get(self, "main_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the organization.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The owner id of the organization.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def status(self) -> int:
        """
        The status of the organization.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> int:
        """
        The type of the organization.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The updated time of the organization.
        """
        return pulumi.get(self, "updated_time")


@pulumi.output_type
class ServiceControlPoliciesPolicyResult(dict):
    def __init__(__self__, *,
                 create_date: str,
                 description: str,
                 id: str,
                 policy_name: str,
                 policy_type: str,
                 statement: str,
                 update_date: str):
        """
        :param str create_date: The create time of the Policy.
        :param str description: The description of the Policy.
        :param str id: The ID of the Policy.
        :param str policy_name: The name of the Policy.
        :param str policy_type: The type of policy. The value can be System or Custom.
        :param str statement: The statement of the Policy.
        :param str update_date: The update time of the Policy.
        """
        pulumi.set(__self__, "create_date", create_date)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "policy_name", policy_name)
        pulumi.set(__self__, "policy_type", policy_type)
        pulumi.set(__self__, "statement", statement)
        pulumi.set(__self__, "update_date", update_date)

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> str:
        """
        The create time of the Policy.
        """
        return pulumi.get(self, "create_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the Policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Policy.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> str:
        """
        The name of the Policy.
        """
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> str:
        """
        The type of policy. The value can be System or Custom.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter
    def statement(self) -> str:
        """
        The statement of the Policy.
        """
        return pulumi.get(self, "statement")

    @property
    @pulumi.getter(name="updateDate")
    def update_date(self) -> str:
        """
        The update time of the Policy.
        """
        return pulumi.get(self, "update_date")


@pulumi.output_type
class UnitsUnitResult(dict):
    def __init__(__self__, *,
                 created_time: str,
                 delete_uk: str,
                 deleted_time: str,
                 depth: int,
                 description: str,
                 id: str,
                 name: str,
                 org_id: str,
                 org_type: int,
                 owner: str,
                 parent_id: str,
                 updated_time: str):
        """
        :param str created_time: The created time of the organization unit.
        :param str delete_uk: Delete marker.
        :param str deleted_time: The deleted time of the organization unit.
        :param int depth: The depth of the organization unit.
        :param str description: The description of the organization unit.
        :param str id: The id of the organization unit.
        :param str name: The name of the organization unit.
        :param str org_id: The id of the organization.
        :param int org_type: The organization type.
        :param str owner: The owner of the organization unit.
        :param str parent_id: Parent Unit ID.
        :param str updated_time: The updated time of the organization unit.
        """
        pulumi.set(__self__, "created_time", created_time)
        pulumi.set(__self__, "delete_uk", delete_uk)
        pulumi.set(__self__, "deleted_time", deleted_time)
        pulumi.set(__self__, "depth", depth)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "org_type", org_type)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "parent_id", parent_id)
        pulumi.set(__self__, "updated_time", updated_time)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The created time of the organization unit.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="deleteUk")
    def delete_uk(self) -> str:
        """
        Delete marker.
        """
        return pulumi.get(self, "delete_uk")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> str:
        """
        The deleted time of the organization unit.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter
    def depth(self) -> int:
        """
        The depth of the organization unit.
        """
        return pulumi.get(self, "depth")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the organization unit.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the organization unit.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the organization unit.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        """
        The id of the organization.
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="orgType")
    def org_type(self) -> int:
        """
        The organization type.
        """
        return pulumi.get(self, "org_type")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The owner of the organization unit.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> str:
        """
        Parent Unit ID.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The updated time of the organization unit.
        """
        return pulumi.get(self, "updated_time")

