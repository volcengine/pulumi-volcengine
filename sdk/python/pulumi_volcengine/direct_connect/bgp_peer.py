# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BgpPeerArgs', 'BgpPeer']

@pulumi.input_type
class BgpPeerArgs:
    def __init__(__self__, *,
                 remote_asn: pulumi.Input[int],
                 virtual_interface_id: pulumi.Input[str],
                 auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_peer_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BgpPeer resource.
        :param pulumi.Input[int] remote_asn: The remote asn of bgp peer.
        :param pulumi.Input[str] virtual_interface_id: The id of virtual interface.
        :param pulumi.Input[str] auth_key: The auth key of bgp peer.
        :param pulumi.Input[str] bgp_peer_name: The name of bgp peer.
        :param pulumi.Input[str] description: The description of bgp peer.
        """
        pulumi.set(__self__, "remote_asn", remote_asn)
        pulumi.set(__self__, "virtual_interface_id", virtual_interface_id)
        if auth_key is not None:
            pulumi.set(__self__, "auth_key", auth_key)
        if bgp_peer_name is not None:
            pulumi.set(__self__, "bgp_peer_name", bgp_peer_name)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="remoteAsn")
    def remote_asn(self) -> pulumi.Input[int]:
        """
        The remote asn of bgp peer.
        """
        return pulumi.get(self, "remote_asn")

    @remote_asn.setter
    def remote_asn(self, value: pulumi.Input[int]):
        pulumi.set(self, "remote_asn", value)

    @property
    @pulumi.getter(name="virtualInterfaceId")
    def virtual_interface_id(self) -> pulumi.Input[str]:
        """
        The id of virtual interface.
        """
        return pulumi.get(self, "virtual_interface_id")

    @virtual_interface_id.setter
    def virtual_interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "virtual_interface_id", value)

    @property
    @pulumi.getter(name="authKey")
    def auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        The auth key of bgp peer.
        """
        return pulumi.get(self, "auth_key")

    @auth_key.setter
    def auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_key", value)

    @property
    @pulumi.getter(name="bgpPeerName")
    def bgp_peer_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of bgp peer.
        """
        return pulumi.get(self, "bgp_peer_name")

    @bgp_peer_name.setter
    def bgp_peer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_peer_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of bgp peer.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _BgpPeerState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_peer_id: Optional[pulumi.Input[str]] = None,
                 bgp_peer_name: Optional[pulumi.Input[str]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 local_asn: Optional[pulumi.Input[int]] = None,
                 remote_asn: Optional[pulumi.Input[int]] = None,
                 session_status: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 virtual_interface_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpPeer resources.
        :param pulumi.Input[str] account_id: The id of account.
        :param pulumi.Input[str] auth_key: The auth key of bgp peer.
        :param pulumi.Input[str] bgp_peer_id: The id of bgp peer.
        :param pulumi.Input[str] bgp_peer_name: The name of bgp peer.
        :param pulumi.Input[str] creation_time: The create time of bgp peer.
        :param pulumi.Input[str] description: The description of bgp peer.
        :param pulumi.Input[int] local_asn: The local asn of bgp peer.
        :param pulumi.Input[int] remote_asn: The remote asn of bgp peer.
        :param pulumi.Input[str] session_status: The session status of bgp peer.
        :param pulumi.Input[str] status: The status of bgp peer.
        :param pulumi.Input[str] update_time: The update time of bgp peer.
        :param pulumi.Input[str] virtual_interface_id: The id of virtual interface.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if auth_key is not None:
            pulumi.set(__self__, "auth_key", auth_key)
        if bgp_peer_id is not None:
            pulumi.set(__self__, "bgp_peer_id", bgp_peer_id)
        if bgp_peer_name is not None:
            pulumi.set(__self__, "bgp_peer_name", bgp_peer_name)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if local_asn is not None:
            pulumi.set(__self__, "local_asn", local_asn)
        if remote_asn is not None:
            pulumi.set(__self__, "remote_asn", remote_asn)
        if session_status is not None:
            pulumi.set(__self__, "session_status", session_status)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if virtual_interface_id is not None:
            pulumi.set(__self__, "virtual_interface_id", virtual_interface_id)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of account.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="authKey")
    def auth_key(self) -> Optional[pulumi.Input[str]]:
        """
        The auth key of bgp peer.
        """
        return pulumi.get(self, "auth_key")

    @auth_key.setter
    def auth_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_key", value)

    @property
    @pulumi.getter(name="bgpPeerId")
    def bgp_peer_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of bgp peer.
        """
        return pulumi.get(self, "bgp_peer_id")

    @bgp_peer_id.setter
    def bgp_peer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_peer_id", value)

    @property
    @pulumi.getter(name="bgpPeerName")
    def bgp_peer_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of bgp peer.
        """
        return pulumi.get(self, "bgp_peer_name")

    @bgp_peer_name.setter
    def bgp_peer_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_peer_name", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of bgp peer.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of bgp peer.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="localAsn")
    def local_asn(self) -> Optional[pulumi.Input[int]]:
        """
        The local asn of bgp peer.
        """
        return pulumi.get(self, "local_asn")

    @local_asn.setter
    def local_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "local_asn", value)

    @property
    @pulumi.getter(name="remoteAsn")
    def remote_asn(self) -> Optional[pulumi.Input[int]]:
        """
        The remote asn of bgp peer.
        """
        return pulumi.get(self, "remote_asn")

    @remote_asn.setter
    def remote_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "remote_asn", value)

    @property
    @pulumi.getter(name="sessionStatus")
    def session_status(self) -> Optional[pulumi.Input[str]]:
        """
        The session status of bgp peer.
        """
        return pulumi.get(self, "session_status")

    @session_status.setter
    def session_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_status", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of bgp peer.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of bgp peer.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter(name="virtualInterfaceId")
    def virtual_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of virtual interface.
        """
        return pulumi.get(self, "virtual_interface_id")

    @virtual_interface_id.setter
    def virtual_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "virtual_interface_id", value)


class BgpPeer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_peer_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[int]] = None,
                 virtual_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage direct connect bgp peer
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.direct_connect.BgpPeer("foo",
            description="tf-test",
            remote_asn=2000,
            virtual_interface_id="dcv-62vi13v131tsn3gd6il****")
        ```

        ## Import

        DirectConnectBgpPeer can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:direct_connect/bgpPeer:BgpPeer default bgp-2752hz4teko3k7fap8u4c****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_key: The auth key of bgp peer.
        :param pulumi.Input[str] bgp_peer_name: The name of bgp peer.
        :param pulumi.Input[str] description: The description of bgp peer.
        :param pulumi.Input[int] remote_asn: The remote asn of bgp peer.
        :param pulumi.Input[str] virtual_interface_id: The id of virtual interface.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpPeerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage direct connect bgp peer
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.direct_connect.BgpPeer("foo",
            description="tf-test",
            remote_asn=2000,
            virtual_interface_id="dcv-62vi13v131tsn3gd6il****")
        ```

        ## Import

        DirectConnectBgpPeer can be imported using the id, e.g.

        ```sh
        $ pulumi import volcengine:direct_connect/bgpPeer:BgpPeer default bgp-2752hz4teko3k7fap8u4c****
        ```

        :param str resource_name: The name of the resource.
        :param BgpPeerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpPeerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_key: Optional[pulumi.Input[str]] = None,
                 bgp_peer_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[int]] = None,
                 virtual_interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpPeerArgs.__new__(BgpPeerArgs)

            __props__.__dict__["auth_key"] = auth_key
            __props__.__dict__["bgp_peer_name"] = bgp_peer_name
            __props__.__dict__["description"] = description
            if remote_asn is None and not opts.urn:
                raise TypeError("Missing required property 'remote_asn'")
            __props__.__dict__["remote_asn"] = remote_asn
            if virtual_interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_interface_id'")
            __props__.__dict__["virtual_interface_id"] = virtual_interface_id
            __props__.__dict__["account_id"] = None
            __props__.__dict__["bgp_peer_id"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["local_asn"] = None
            __props__.__dict__["session_status"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["update_time"] = None
        super(BgpPeer, __self__).__init__(
            'volcengine:direct_connect/bgpPeer:BgpPeer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            auth_key: Optional[pulumi.Input[str]] = None,
            bgp_peer_id: Optional[pulumi.Input[str]] = None,
            bgp_peer_name: Optional[pulumi.Input[str]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            local_asn: Optional[pulumi.Input[int]] = None,
            remote_asn: Optional[pulumi.Input[int]] = None,
            session_status: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            virtual_interface_id: Optional[pulumi.Input[str]] = None) -> 'BgpPeer':
        """
        Get an existing BgpPeer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The id of account.
        :param pulumi.Input[str] auth_key: The auth key of bgp peer.
        :param pulumi.Input[str] bgp_peer_id: The id of bgp peer.
        :param pulumi.Input[str] bgp_peer_name: The name of bgp peer.
        :param pulumi.Input[str] creation_time: The create time of bgp peer.
        :param pulumi.Input[str] description: The description of bgp peer.
        :param pulumi.Input[int] local_asn: The local asn of bgp peer.
        :param pulumi.Input[int] remote_asn: The remote asn of bgp peer.
        :param pulumi.Input[str] session_status: The session status of bgp peer.
        :param pulumi.Input[str] status: The status of bgp peer.
        :param pulumi.Input[str] update_time: The update time of bgp peer.
        :param pulumi.Input[str] virtual_interface_id: The id of virtual interface.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpPeerState.__new__(_BgpPeerState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["auth_key"] = auth_key
        __props__.__dict__["bgp_peer_id"] = bgp_peer_id
        __props__.__dict__["bgp_peer_name"] = bgp_peer_name
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["local_asn"] = local_asn
        __props__.__dict__["remote_asn"] = remote_asn
        __props__.__dict__["session_status"] = session_status
        __props__.__dict__["status"] = status
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["virtual_interface_id"] = virtual_interface_id
        return BgpPeer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The id of account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="authKey")
    def auth_key(self) -> pulumi.Output[Optional[str]]:
        """
        The auth key of bgp peer.
        """
        return pulumi.get(self, "auth_key")

    @property
    @pulumi.getter(name="bgpPeerId")
    def bgp_peer_id(self) -> pulumi.Output[str]:
        """
        The id of bgp peer.
        """
        return pulumi.get(self, "bgp_peer_id")

    @property
    @pulumi.getter(name="bgpPeerName")
    def bgp_peer_name(self) -> pulumi.Output[str]:
        """
        The name of bgp peer.
        """
        return pulumi.get(self, "bgp_peer_name")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The create time of bgp peer.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of bgp peer.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="localAsn")
    def local_asn(self) -> pulumi.Output[int]:
        """
        The local asn of bgp peer.
        """
        return pulumi.get(self, "local_asn")

    @property
    @pulumi.getter(name="remoteAsn")
    def remote_asn(self) -> pulumi.Output[int]:
        """
        The remote asn of bgp peer.
        """
        return pulumi.get(self, "remote_asn")

    @property
    @pulumi.getter(name="sessionStatus")
    def session_status(self) -> pulumi.Output[str]:
        """
        The session status of bgp peer.
        """
        return pulumi.get(self, "session_status")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of bgp peer.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time of bgp peer.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="virtualInterfaceId")
    def virtual_interface_id(self) -> pulumi.Output[str]:
        """
        The id of virtual interface.
        """
        return pulumi.get(self, "virtual_interface_id")

