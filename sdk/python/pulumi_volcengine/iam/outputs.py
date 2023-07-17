# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'PoliciesPolicyResult',
    'RolesRoleResult',
    'UsersUserResult',
]

@pulumi.output_type
class PoliciesPolicyResult(dict):
    def __init__(__self__, *,
                 create_date: str,
                 description: str,
                 id: str,
                 policy_document: str,
                 policy_name: str,
                 policy_trn: str,
                 policy_type: str,
                 role_attach_date: str,
                 role_name: str,
                 update_date: str,
                 user_attach_date: str,
                 user_name: str):
        """
        :param str create_date: The create time of the Policy.
        :param str description: The description of the Policy.
        :param str id: The ID of the Policy.
        :param str policy_document: The document of the Policy.
        :param str policy_name: The name of the Policy.
        :param str policy_trn: The resource name of the Policy.
        :param str policy_type: The type of the Policy.
        :param str role_attach_date: The role attach time of the Policy.The data show only query with role_name.
        :param str role_name: The name of the IAM role.
        :param str update_date: The update time of the Policy.
        :param str user_attach_date: The user attach time of the Policy.The data show only query with user_name.
        :param str user_name: The name of the IAM user.
        """
        pulumi.set(__self__, "create_date", create_date)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "policy_document", policy_document)
        pulumi.set(__self__, "policy_name", policy_name)
        pulumi.set(__self__, "policy_trn", policy_trn)
        pulumi.set(__self__, "policy_type", policy_type)
        pulumi.set(__self__, "role_attach_date", role_attach_date)
        pulumi.set(__self__, "role_name", role_name)
        pulumi.set(__self__, "update_date", update_date)
        pulumi.set(__self__, "user_attach_date", user_attach_date)
        pulumi.set(__self__, "user_name", user_name)

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
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> str:
        """
        The document of the Policy.
        """
        return pulumi.get(self, "policy_document")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> str:
        """
        The name of the Policy.
        """
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="policyTrn")
    def policy_trn(self) -> str:
        """
        The resource name of the Policy.
        """
        return pulumi.get(self, "policy_trn")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> str:
        """
        The type of the Policy.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="roleAttachDate")
    def role_attach_date(self) -> str:
        """
        The role attach time of the Policy.The data show only query with role_name.
        """
        return pulumi.get(self, "role_attach_date")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> str:
        """
        The name of the IAM role.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter(name="updateDate")
    def update_date(self) -> str:
        """
        The update time of the Policy.
        """
        return pulumi.get(self, "update_date")

    @property
    @pulumi.getter(name="userAttachDate")
    def user_attach_date(self) -> str:
        """
        The user attach time of the Policy.The data show only query with user_name.
        """
        return pulumi.get(self, "user_attach_date")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The name of the IAM user.
        """
        return pulumi.get(self, "user_name")


@pulumi.output_type
class RolesRoleResult(dict):
    def __init__(__self__, *,
                 create_date: str,
                 description: str,
                 id: str,
                 role_name: str,
                 trn: str,
                 trust_policy_document: str):
        """
        :param str create_date: The create time of the Role.
        :param str description: The description of the Role.
        :param str id: The ID of the Role.
        :param str role_name: The name of the Role, comma separated.
        :param str trn: The resource name of the Role.
        :param str trust_policy_document: The trust policy document of the Role.
        """
        pulumi.set(__self__, "create_date", create_date)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "role_name", role_name)
        pulumi.set(__self__, "trn", trn)
        pulumi.set(__self__, "trust_policy_document", trust_policy_document)

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> str:
        """
        The create time of the Role.
        """
        return pulumi.get(self, "create_date")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the Role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the Role.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> str:
        """
        The name of the Role, comma separated.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter
    def trn(self) -> str:
        """
        The resource name of the Role.
        """
        return pulumi.get(self, "trn")

    @property
    @pulumi.getter(name="trustPolicyDocument")
    def trust_policy_document(self) -> str:
        """
        The trust policy document of the Role.
        """
        return pulumi.get(self, "trust_policy_document")


@pulumi.output_type
class UsersUserResult(dict):
    def __init__(__self__, *,
                 account_id: str,
                 create_date: str,
                 trn: str,
                 update_date: str,
                 user_name: str):
        """
        :param str account_id: The account id of the user.
        :param str create_date: The create date of the user.
        :param str trn: The trn of the user.
        :param str update_date: The update date of the user.
        :param str user_name: The name of the user.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "create_date", create_date)
        pulumi.set(__self__, "trn", trn)
        pulumi.set(__self__, "update_date", update_date)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The account id of the user.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> str:
        """
        The create date of the user.
        """
        return pulumi.get(self, "create_date")

    @property
    @pulumi.getter
    def trn(self) -> str:
        """
        The trn of the user.
        """
        return pulumi.get(self, "trn")

    @property
    @pulumi.getter(name="updateDate")
    def update_date(self) -> str:
        """
        The update date of the user.
        """
        return pulumi.get(self, "update_date")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The name of the user.
        """
        return pulumi.get(self, "user_name")

