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
    'GroupMember',
    'GroupsGroupResult',
    'PermissionSetAssignmentsAssignmentResult',
    'PermissionSetPermissionPolicy',
    'PermissionSetProvisioningsPermissionProvisioningResult',
    'PermissionSetsPermissionSetResult',
    'PermissionSetsPermissionSetPermissionPolicyResult',
    'UserProvisioningsUserProvisioningResult',
    'UsersUserResult',
]

@pulumi.output_type
class GroupMember(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "displayName":
            suggest = "display_name"
        elif key == "identityType":
            suggest = "identity_type"
        elif key == "joinTime":
            suggest = "join_time"
        elif key == "userId":
            suggest = "user_id"
        elif key == "userName":
            suggest = "user_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in GroupMember. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        GroupMember.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        GroupMember.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[str] = None,
                 display_name: Optional[str] = None,
                 email: Optional[str] = None,
                 identity_type: Optional[str] = None,
                 join_time: Optional[str] = None,
                 phone: Optional[str] = None,
                 source: Optional[str] = None,
                 user_id: Optional[str] = None,
                 user_name: Optional[str] = None):
        """
        :param str description: The description of the cloud identity group.
        :param str display_name: The display name of the cloud identity group.
        :param str email: The email of the cloud identity user.
        :param str identity_type: The identity type of the cloud identity user.
        :param str join_time: The join time of the cloud identity user.
        :param str phone: The phone of the cloud identity user.
        :param str source: The source of the cloud identity group.
        :param str user_id: The id of the cloud identity user.
        :param str user_name: The name of the cloud identity user.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if identity_type is not None:
            pulumi.set(__self__, "identity_type", identity_type)
        if join_time is not None:
            pulumi.set(__self__, "join_time", join_time)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)
        if user_name is not None:
            pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the cloud identity group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The display name of the cloud identity group.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def email(self) -> Optional[str]:
        """
        The email of the cloud identity user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="identityType")
    def identity_type(self) -> Optional[str]:
        """
        The identity type of the cloud identity user.
        """
        return pulumi.get(self, "identity_type")

    @property
    @pulumi.getter(name="joinTime")
    def join_time(self) -> Optional[str]:
        """
        The join time of the cloud identity user.
        """
        return pulumi.get(self, "join_time")

    @property
    @pulumi.getter
    def phone(self) -> Optional[str]:
        """
        The phone of the cloud identity user.
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        """
        The source of the cloud identity group.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[str]:
        """
        The id of the cloud identity user.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> Optional[str]:
        """
        The name of the cloud identity user.
        """
        return pulumi.get(self, "user_name")


@pulumi.output_type
class GroupsGroupResult(dict):
    def __init__(__self__, *,
                 created_time: str,
                 description: str,
                 display_name: str,
                 group_id: str,
                 group_name: str,
                 id: str,
                 join_type: str,
                 source: str,
                 updated_time: str):
        """
        :param str created_time: The created time of the cloud identity group.
        :param str description: The description of the cloud identity group.
        :param str display_name: The display name of cloud identity group.
        :param str group_id: The id of the cloud identity group.
        :param str group_name: The name of cloud identity group.
        :param str id: The id of the cloud identity group.
        :param str join_type: The join type of cloud identity group. Valid values: `Auto`, `Manual`.
        :param str source: The source of the cloud identity group.
        :param str updated_time: The updated time of the cloud identity group.
        """
        pulumi.set(__self__, "created_time", created_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "join_type", join_type)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "updated_time", updated_time)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The created time of the cloud identity group.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the cloud identity group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name of cloud identity group.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> str:
        """
        The id of the cloud identity group.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> str:
        """
        The name of cloud identity group.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the cloud identity group.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="joinType")
    def join_type(self) -> str:
        """
        The join type of cloud identity group. Valid values: `Auto`, `Manual`.
        """
        return pulumi.get(self, "join_type")

    @property
    @pulumi.getter
    def source(self) -> str:
        """
        The source of the cloud identity group.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The updated time of the cloud identity group.
        """
        return pulumi.get(self, "updated_time")


@pulumi.output_type
class PermissionSetAssignmentsAssignmentResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 id: str,
                 permission_set_id: str,
                 permission_set_name: str,
                 principal_id: str,
                 principal_type: str,
                 target_id: str):
        """
        :param str create_time: The create time of the cloud identity permission set assignment.
        :param str id: The id of the cloud identity permission set.
        :param str permission_set_id: The id of cloud identity permission set.
        :param str permission_set_name: The name of the cloud identity permission set.
        :param str principal_id: The principal id of cloud identity permission set. When the `principal_type` is `User`, this field is specified to `UserId`. When the `principal_type` is `Group`, this field is specified to `GroupId`.
        :param str principal_type: The principal type of cloud identity permission set. Valid values: `User`, `Group`.
        :param str target_id: The target account id of cloud identity permission set assignment.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "permission_set_id", permission_set_id)
        pulumi.set(__self__, "permission_set_name", permission_set_name)
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "principal_type", principal_type)
        pulumi.set(__self__, "target_id", target_id)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The create time of the cloud identity permission set assignment.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the cloud identity permission set.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="permissionSetId")
    def permission_set_id(self) -> str:
        """
        The id of cloud identity permission set.
        """
        return pulumi.get(self, "permission_set_id")

    @property
    @pulumi.getter(name="permissionSetName")
    def permission_set_name(self) -> str:
        """
        The name of the cloud identity permission set.
        """
        return pulumi.get(self, "permission_set_name")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal id of cloud identity permission set. When the `principal_type` is `User`, this field is specified to `UserId`. When the `principal_type` is `Group`, this field is specified to `GroupId`.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> str:
        """
        The principal type of cloud identity permission set. Valid values: `User`, `Group`.
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        The target account id of cloud identity permission set assignment.
        """
        return pulumi.get(self, "target_id")


@pulumi.output_type
class PermissionSetPermissionPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "permissionPolicyType":
            suggest = "permission_policy_type"
        elif key == "inlinePolicyDocument":
            suggest = "inline_policy_document"
        elif key == "permissionPolicyName":
            suggest = "permission_policy_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PermissionSetPermissionPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PermissionSetPermissionPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PermissionSetPermissionPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 permission_policy_type: str,
                 inline_policy_document: Optional[str] = None,
                 permission_policy_name: Optional[str] = None):
        """
        :param str permission_policy_type: The type of the cloud identity permission set policy. Valid values: `System`, `Inline`.
        :param str inline_policy_document: The document of the cloud identity permission set inline policy. When the `permission_policy_type` is `Inline`, this field must be specified.
        :param str permission_policy_name: The name of the cloud identity permission set system policy. When the `permission_policy_type` is `System`, this field must be specified.
        """
        pulumi.set(__self__, "permission_policy_type", permission_policy_type)
        if inline_policy_document is not None:
            pulumi.set(__self__, "inline_policy_document", inline_policy_document)
        if permission_policy_name is not None:
            pulumi.set(__self__, "permission_policy_name", permission_policy_name)

    @property
    @pulumi.getter(name="permissionPolicyType")
    def permission_policy_type(self) -> str:
        """
        The type of the cloud identity permission set policy. Valid values: `System`, `Inline`.
        """
        return pulumi.get(self, "permission_policy_type")

    @property
    @pulumi.getter(name="inlinePolicyDocument")
    def inline_policy_document(self) -> Optional[str]:
        """
        The document of the cloud identity permission set inline policy. When the `permission_policy_type` is `Inline`, this field must be specified.
        """
        return pulumi.get(self, "inline_policy_document")

    @property
    @pulumi.getter(name="permissionPolicyName")
    def permission_policy_name(self) -> Optional[str]:
        """
        The name of the cloud identity permission set system policy. When the `permission_policy_type` is `System`, this field must be specified.
        """
        return pulumi.get(self, "permission_policy_name")


@pulumi.output_type
class PermissionSetProvisioningsPermissionProvisioningResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 id: str,
                 permission_set_id: str,
                 permission_set_name: str,
                 target_id: str,
                 update_time: str):
        """
        :param str create_time: The create time of the cloud identity permission set provisioning.
        :param str id: The id of the cloud identity permission set.
        :param str permission_set_id: The id of cloud identity permission set.
        :param str permission_set_name: The name of the cloud identity permission set.
        :param str target_id: The target account id of cloud identity permission set.
        :param str update_time: The update time of the cloud identity permission set provisioning.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "permission_set_id", permission_set_id)
        pulumi.set(__self__, "permission_set_name", permission_set_name)
        pulumi.set(__self__, "target_id", target_id)
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The create time of the cloud identity permission set provisioning.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the cloud identity permission set.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="permissionSetId")
    def permission_set_id(self) -> str:
        """
        The id of cloud identity permission set.
        """
        return pulumi.get(self, "permission_set_id")

    @property
    @pulumi.getter(name="permissionSetName")
    def permission_set_name(self) -> str:
        """
        The name of the cloud identity permission set.
        """
        return pulumi.get(self, "permission_set_name")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        The target account id of cloud identity permission set.
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The update time of the cloud identity permission set provisioning.
        """
        return pulumi.get(self, "update_time")


@pulumi.output_type
class PermissionSetsPermissionSetResult(dict):
    def __init__(__self__, *,
                 created_time: str,
                 description: str,
                 id: str,
                 name: str,
                 permission_policies: Sequence['outputs.PermissionSetsPermissionSetPermissionPolicyResult'],
                 permission_set_id: str,
                 relay_state: str,
                 session_duration: int,
                 updated_time: str):
        """
        :param str created_time: The create time of the cloud identity permission set.
        :param str description: The description of the cloud identity permission set.
        :param str id: The id of the cloud identity permission set.
        :param str name: The name of the cloud identity permission set.
        :param Sequence['PermissionSetsPermissionSetPermissionPolicyArgs'] permission_policies: The policies of the cloud identity permission set.
        :param str permission_set_id: The id of the cloud identity permission set.
        :param str relay_state: The relay state of the cloud identity permission set.
        :param int session_duration: The session duration of the cloud identity permission set.
        :param str updated_time: The updated time of the cloud identity permission set.
        """
        pulumi.set(__self__, "created_time", created_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "permission_policies", permission_policies)
        pulumi.set(__self__, "permission_set_id", permission_set_id)
        pulumi.set(__self__, "relay_state", relay_state)
        pulumi.set(__self__, "session_duration", session_duration)
        pulumi.set(__self__, "updated_time", updated_time)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The create time of the cloud identity permission set.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the cloud identity permission set.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the cloud identity permission set.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the cloud identity permission set.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="permissionPolicies")
    def permission_policies(self) -> Sequence['outputs.PermissionSetsPermissionSetPermissionPolicyResult']:
        """
        The policies of the cloud identity permission set.
        """
        return pulumi.get(self, "permission_policies")

    @property
    @pulumi.getter(name="permissionSetId")
    def permission_set_id(self) -> str:
        """
        The id of the cloud identity permission set.
        """
        return pulumi.get(self, "permission_set_id")

    @property
    @pulumi.getter(name="relayState")
    def relay_state(self) -> str:
        """
        The relay state of the cloud identity permission set.
        """
        return pulumi.get(self, "relay_state")

    @property
    @pulumi.getter(name="sessionDuration")
    def session_duration(self) -> int:
        """
        The session duration of the cloud identity permission set.
        """
        return pulumi.get(self, "session_duration")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The updated time of the cloud identity permission set.
        """
        return pulumi.get(self, "updated_time")


@pulumi.output_type
class PermissionSetsPermissionSetPermissionPolicyResult(dict):
    def __init__(__self__, *,
                 create_time: str,
                 permission_policy_document: str,
                 permission_policy_name: str,
                 permission_policy_type: str):
        """
        :param str create_time: The create time of the cloud identity permission set policy.
        :param str permission_policy_document: The document of the cloud identity permission set policy.
        :param str permission_policy_name: The name of the cloud identity permission set policy.
        :param str permission_policy_type: The type of the cloud identity permission set policy.
        """
        pulumi.set(__self__, "create_time", create_time)
        pulumi.set(__self__, "permission_policy_document", permission_policy_document)
        pulumi.set(__self__, "permission_policy_name", permission_policy_name)
        pulumi.set(__self__, "permission_policy_type", permission_policy_type)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The create time of the cloud identity permission set policy.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="permissionPolicyDocument")
    def permission_policy_document(self) -> str:
        """
        The document of the cloud identity permission set policy.
        """
        return pulumi.get(self, "permission_policy_document")

    @property
    @pulumi.getter(name="permissionPolicyName")
    def permission_policy_name(self) -> str:
        """
        The name of the cloud identity permission set policy.
        """
        return pulumi.get(self, "permission_policy_name")

    @property
    @pulumi.getter(name="permissionPolicyType")
    def permission_policy_type(self) -> str:
        """
        The type of the cloud identity permission set policy.
        """
        return pulumi.get(self, "permission_policy_type")


@pulumi.output_type
class UserProvisioningsUserProvisioningResult(dict):
    def __init__(__self__, *,
                 created_time: str,
                 deletion_strategy: str,
                 department_names: Sequence[str],
                 description: str,
                 duplication_strategy: str,
                 duplication_suffix: str,
                 id: str,
                 identity_source_strategy: str,
                 principal_id: str,
                 principal_name: str,
                 principal_type: str,
                 provision_status: str,
                 target_id: str,
                 updated_time: str,
                 user_provisioning_id: str):
        """
        :param str created_time: The created time of the cloud identity user provisioning.
        :param str deletion_strategy: The deletion strategy of the cloud identity user provisioning.
        :param Sequence[str] department_names: The department names of the cloud identity user provisioning.
        :param str description: The description of the cloud identity user provisioning.
        :param str duplication_strategy: The duplication strategy of the cloud identity user provisioning.
        :param str duplication_suffix: The duplication suffix of the cloud identity user provisioning.
        :param str id: The id of the cloud identity user provisioning.
        :param str identity_source_strategy: The identity source strategy of the cloud identity user provisioning.
        :param str principal_id: The principal id of the cloud identity user provisioning.
        :param str principal_name: The principal name of the cloud identity user provisioning.
        :param str principal_type: The principal type of the cloud identity user provisioning.
        :param str provision_status: The status of the cloud identity user provisioning.
        :param str target_id: The target account id of the cloud identity user provisioning.
        :param str updated_time: The updated time of the cloud identity user provisioning.
        :param str user_provisioning_id: The id of the cloud identity user provisioning.
        """
        pulumi.set(__self__, "created_time", created_time)
        pulumi.set(__self__, "deletion_strategy", deletion_strategy)
        pulumi.set(__self__, "department_names", department_names)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "duplication_strategy", duplication_strategy)
        pulumi.set(__self__, "duplication_suffix", duplication_suffix)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "identity_source_strategy", identity_source_strategy)
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "principal_name", principal_name)
        pulumi.set(__self__, "principal_type", principal_type)
        pulumi.set(__self__, "provision_status", provision_status)
        pulumi.set(__self__, "target_id", target_id)
        pulumi.set(__self__, "updated_time", updated_time)
        pulumi.set(__self__, "user_provisioning_id", user_provisioning_id)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The created time of the cloud identity user provisioning.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="deletionStrategy")
    def deletion_strategy(self) -> str:
        """
        The deletion strategy of the cloud identity user provisioning.
        """
        return pulumi.get(self, "deletion_strategy")

    @property
    @pulumi.getter(name="departmentNames")
    def department_names(self) -> Sequence[str]:
        """
        The department names of the cloud identity user provisioning.
        """
        return pulumi.get(self, "department_names")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the cloud identity user provisioning.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="duplicationStrategy")
    def duplication_strategy(self) -> str:
        """
        The duplication strategy of the cloud identity user provisioning.
        """
        return pulumi.get(self, "duplication_strategy")

    @property
    @pulumi.getter(name="duplicationSuffix")
    def duplication_suffix(self) -> str:
        """
        The duplication suffix of the cloud identity user provisioning.
        """
        return pulumi.get(self, "duplication_suffix")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the cloud identity user provisioning.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identitySourceStrategy")
    def identity_source_strategy(self) -> str:
        """
        The identity source strategy of the cloud identity user provisioning.
        """
        return pulumi.get(self, "identity_source_strategy")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal id of the cloud identity user provisioning.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalName")
    def principal_name(self) -> str:
        """
        The principal name of the cloud identity user provisioning.
        """
        return pulumi.get(self, "principal_name")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> str:
        """
        The principal type of the cloud identity user provisioning.
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="provisionStatus")
    def provision_status(self) -> str:
        """
        The status of the cloud identity user provisioning.
        """
        return pulumi.get(self, "provision_status")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        The target account id of the cloud identity user provisioning.
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The updated time of the cloud identity user provisioning.
        """
        return pulumi.get(self, "updated_time")

    @property
    @pulumi.getter(name="userProvisioningId")
    def user_provisioning_id(self) -> str:
        """
        The id of the cloud identity user provisioning.
        """
        return pulumi.get(self, "user_provisioning_id")


@pulumi.output_type
class UsersUserResult(dict):
    def __init__(__self__, *,
                 created_time: str,
                 description: str,
                 display_name: str,
                 email: str,
                 id: str,
                 identity_type: str,
                 phone: str,
                 source: str,
                 updated_time: str,
                 user_id: str,
                 user_name: str):
        """
        :param str created_time: The created time of the cloud identity user.
        :param str description: The description of the cloud identity user.
        :param str display_name: The display name of cloud identity user.
        :param str email: The email of the cloud identity user.
        :param str id: The id of the cloud identity user.
        :param str identity_type: The identity type of the cloud identity user.
        :param str phone: The phone of the cloud identity user.
        :param str source: The source of cloud identity user. Valid values: `Sync`, `Manual`.
        :param str updated_time: The updated time of the cloud identity user.
        :param str user_id: The id of the cloud identity user.
        :param str user_name: The name of cloud identity user.
        """
        pulumi.set(__self__, "created_time", created_time)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "email", email)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "identity_type", identity_type)
        pulumi.set(__self__, "phone", phone)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "updated_time", updated_time)
        pulumi.set(__self__, "user_id", user_id)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The created time of the cloud identity user.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the cloud identity user.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name of cloud identity user.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The email of the cloud identity user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the cloud identity user.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identityType")
    def identity_type(self) -> str:
        """
        The identity type of the cloud identity user.
        """
        return pulumi.get(self, "identity_type")

    @property
    @pulumi.getter
    def phone(self) -> str:
        """
        The phone of the cloud identity user.
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter
    def source(self) -> str:
        """
        The source of cloud identity user. Valid values: `Sync`, `Manual`.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="updatedTime")
    def updated_time(self) -> str:
        """
        The updated time of the cloud identity user.
        """
        return pulumi.get(self, "updated_time")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        """
        The id of the cloud identity user.
        """
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        """
        The name of cloud identity user.
        """
        return pulumi.get(self, "user_name")

