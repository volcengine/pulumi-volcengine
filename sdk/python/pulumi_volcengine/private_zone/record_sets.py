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
    'RecordSetsResult',
    'AwaitableRecordSetsResult',
    'record_sets',
    'record_sets_output',
]

@pulumi.output_type
class RecordSetsResult:
    """
    A collection of values returned by RecordSets.
    """
    def __init__(__self__, host=None, id=None, output_file=None, record_set_id=None, record_sets=None, search_mode=None, total_count=None, zid=None):
        if host and not isinstance(host, str):
            raise TypeError("Expected argument 'host' to be a str")
        pulumi.set(__self__, "host", host)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if record_set_id and not isinstance(record_set_id, str):
            raise TypeError("Expected argument 'record_set_id' to be a str")
        pulumi.set(__self__, "record_set_id", record_set_id)
        if record_sets and not isinstance(record_sets, list):
            raise TypeError("Expected argument 'record_sets' to be a list")
        pulumi.set(__self__, "record_sets", record_sets)
        if search_mode and not isinstance(search_mode, str):
            raise TypeError("Expected argument 'search_mode' to be a str")
        pulumi.set(__self__, "search_mode", search_mode)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if zid and not isinstance(zid, int):
            raise TypeError("Expected argument 'zid' to be a int")
        pulumi.set(__self__, "zid", zid)

    @property
    @pulumi.getter
    def host(self) -> Optional[str]:
        """
        The host of the private zone record.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="recordSetId")
    def record_set_id(self) -> Optional[str]:
        """
        The id of the private zone record set.
        """
        return pulumi.get(self, "record_set_id")

    @property
    @pulumi.getter(name="recordSets")
    def record_sets(self) -> Sequence['outputs.RecordSetsRecordSetResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "record_sets")

    @property
    @pulumi.getter(name="searchMode")
    def search_mode(self) -> Optional[str]:
        return pulumi.get(self, "search_mode")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter
    def zid(self) -> int:
        return pulumi.get(self, "zid")


class AwaitableRecordSetsResult(RecordSetsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return RecordSetsResult(
            host=self.host,
            id=self.id,
            output_file=self.output_file,
            record_set_id=self.record_set_id,
            record_sets=self.record_sets,
            search_mode=self.search_mode,
            total_count=self.total_count,
            zid=self.zid)


def record_sets(host: Optional[str] = None,
                output_file: Optional[str] = None,
                record_set_id: Optional[str] = None,
                search_mode: Optional[str] = None,
                zid: Optional[int] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableRecordSetsResult:
    """
    Use this data source to query detailed information of private zone record sets


    :param str host: The host of Private Zone Record Set.
    :param str output_file: File name where to save data source results.
    :param str record_set_id: The id of Private Zone Record Set.
    :param str search_mode: The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
    :param int zid: The zid of Private Zone.
    """
    __args__ = dict()
    __args__['host'] = host
    __args__['outputFile'] = output_file
    __args__['recordSetId'] = record_set_id
    __args__['searchMode'] = search_mode
    __args__['zid'] = zid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:private_zone/recordSets:RecordSets', __args__, opts=opts, typ=RecordSetsResult).value

    return AwaitableRecordSetsResult(
        host=pulumi.get(__ret__, 'host'),
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        record_set_id=pulumi.get(__ret__, 'record_set_id'),
        record_sets=pulumi.get(__ret__, 'record_sets'),
        search_mode=pulumi.get(__ret__, 'search_mode'),
        total_count=pulumi.get(__ret__, 'total_count'),
        zid=pulumi.get(__ret__, 'zid'))


@_utilities.lift_output_func(record_sets)
def record_sets_output(host: Optional[pulumi.Input[Optional[str]]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       record_set_id: Optional[pulumi.Input[Optional[str]]] = None,
                       search_mode: Optional[pulumi.Input[Optional[str]]] = None,
                       zid: Optional[pulumi.Input[int]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[RecordSetsResult]:
    """
    Use this data source to query detailed information of private zone record sets


    :param str host: The host of Private Zone Record Set.
    :param str output_file: File name where to save data source results.
    :param str record_set_id: The id of Private Zone Record Set.
    :param str search_mode: The search mode of query `host`. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
    :param int zid: The zid of Private Zone.
    """
    ...