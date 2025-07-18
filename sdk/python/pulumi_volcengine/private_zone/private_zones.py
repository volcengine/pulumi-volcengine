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

__all__ = [
    'PrivateZonesResult',
    'AwaitablePrivateZonesResult',
    'private_zones',
    'private_zones_output',
]

warnings.warn("""volcengine.private_zone.PrivateZones has been deprecated in favor of volcengine.private_zone.getPrivateZones""", DeprecationWarning)

@pulumi.output_type
class PrivateZonesResult:
    """
    A collection of values returned by PrivateZones.
    """
    def __init__(__self__, id=None, key_word=None, line_mode=None, name_regex=None, output_file=None, private_zones=None, project_name=None, recursion_mode=None, region=None, search_mode=None, tag_filters=None, total_count=None, vpc_id=None, zid=None, zone_name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_word and not isinstance(key_word, str):
            raise TypeError("Expected argument 'key_word' to be a str")
        pulumi.set(__self__, "key_word", key_word)
        if line_mode and not isinstance(line_mode, int):
            raise TypeError("Expected argument 'line_mode' to be a int")
        pulumi.set(__self__, "line_mode", line_mode)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if private_zones and not isinstance(private_zones, list):
            raise TypeError("Expected argument 'private_zones' to be a list")
        pulumi.set(__self__, "private_zones", private_zones)
        if project_name and not isinstance(project_name, str):
            raise TypeError("Expected argument 'project_name' to be a str")
        pulumi.set(__self__, "project_name", project_name)
        if recursion_mode and not isinstance(recursion_mode, bool):
            raise TypeError("Expected argument 'recursion_mode' to be a bool")
        pulumi.set(__self__, "recursion_mode", recursion_mode)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if search_mode and not isinstance(search_mode, str):
            raise TypeError("Expected argument 'search_mode' to be a str")
        pulumi.set(__self__, "search_mode", search_mode)
        if tag_filters and not isinstance(tag_filters, list):
            raise TypeError("Expected argument 'tag_filters' to be a list")
        pulumi.set(__self__, "tag_filters", tag_filters)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError("Expected argument 'vpc_id' to be a str")
        pulumi.set(__self__, "vpc_id", vpc_id)
        if zid and not isinstance(zid, int):
            raise TypeError("Expected argument 'zid' to be a int")
        pulumi.set(__self__, "zid", zid)
        if zone_name and not isinstance(zone_name, str):
            raise TypeError("Expected argument 'zone_name' to be a str")
        pulumi.set(__self__, "zone_name", zone_name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyWord")
    def key_word(self) -> Optional[str]:
        return pulumi.get(self, "key_word")

    @property
    @pulumi.getter(name="lineMode")
    def line_mode(self) -> Optional[int]:
        """
        The line mode of the private zone, specified whether the intelligent mode and the load balance function is enabled.
        """
        return pulumi.get(self, "line_mode")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="privateZones")
    def private_zones(self) -> Sequence['outputs.PrivateZonesPrivateZoneResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "private_zones")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[str]:
        """
        The project name of the private zone.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="recursionMode")
    def recursion_mode(self) -> Optional[bool]:
        """
        Whether the recursion mode of the private zone is enabled.
        """
        return pulumi.get(self, "recursion_mode")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The region of the private zone.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="searchMode")
    def search_mode(self) -> Optional[str]:
        return pulumi.get(self, "search_mode")

    @property
    @pulumi.getter(name="tagFilters")
    def tag_filters(self) -> Optional[Sequence['outputs.PrivateZonesTagFilterResult']]:
        return pulumi.get(self, "tag_filters")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[str]:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter
    def zid(self) -> Optional[int]:
        """
        The id of the private zone.
        """
        return pulumi.get(self, "zid")

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> Optional[str]:
        """
        The id of the private zone.
        """
        return pulumi.get(self, "zone_name")


class AwaitablePrivateZonesResult(PrivateZonesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return PrivateZonesResult(
            id=self.id,
            key_word=self.key_word,
            line_mode=self.line_mode,
            name_regex=self.name_regex,
            output_file=self.output_file,
            private_zones=self.private_zones,
            project_name=self.project_name,
            recursion_mode=self.recursion_mode,
            region=self.region,
            search_mode=self.search_mode,
            tag_filters=self.tag_filters,
            total_count=self.total_count,
            vpc_id=self.vpc_id,
            zid=self.zid,
            zone_name=self.zone_name)


def private_zones(key_word: Optional[str] = None,
                  line_mode: Optional[int] = None,
                  name_regex: Optional[str] = None,
                  output_file: Optional[str] = None,
                  project_name: Optional[str] = None,
                  recursion_mode: Optional[bool] = None,
                  region: Optional[str] = None,
                  search_mode: Optional[str] = None,
                  tag_filters: Optional[Sequence[pulumi.InputType['PrivateZonesTagFilterArgs']]] = None,
                  vpc_id: Optional[str] = None,
                  zid: Optional[int] = None,
                  zone_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitablePrivateZonesResult:
    """
    Use this data source to query detailed information of private zones
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.private_zone.get_private_zones(line_mode=3,
        recursion_mode=True,
        search_mode="EXACT",
        zid=770000,
        zone_name="volces.com")
    ```


    :param str key_word: The keyword of zone name.
    :param int line_mode: The line mode of Private Zone, specified whether the intelligent mode and the load balance function is enabled.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str project_name: The project name of the private zone.
    :param bool recursion_mode: Whether the recursion mode of Private Zone is enabled.
    :param str region: The region of Private Zone.
    :param str search_mode: The search mode of query. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
    :param Sequence[pulumi.InputType['PrivateZonesTagFilterArgs']] tag_filters: List of tag filters.
    :param str vpc_id: The vpc id associated to Private Zone.
    :param int zid: The zid of Private Zone.
    :param str zone_name: The name of Private Zone.
    """
    pulumi.log.warn("""private_zones is deprecated: volcengine.private_zone.PrivateZones has been deprecated in favor of volcengine.private_zone.getPrivateZones""")
    __args__ = dict()
    __args__['keyWord'] = key_word
    __args__['lineMode'] = line_mode
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['projectName'] = project_name
    __args__['recursionMode'] = recursion_mode
    __args__['region'] = region
    __args__['searchMode'] = search_mode
    __args__['tagFilters'] = tag_filters
    __args__['vpcId'] = vpc_id
    __args__['zid'] = zid
    __args__['zoneName'] = zone_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:private_zone/privateZones:PrivateZones', __args__, opts=opts, typ=PrivateZonesResult).value

    return AwaitablePrivateZonesResult(
        id=pulumi.get(__ret__, 'id'),
        key_word=pulumi.get(__ret__, 'key_word'),
        line_mode=pulumi.get(__ret__, 'line_mode'),
        name_regex=pulumi.get(__ret__, 'name_regex'),
        output_file=pulumi.get(__ret__, 'output_file'),
        private_zones=pulumi.get(__ret__, 'private_zones'),
        project_name=pulumi.get(__ret__, 'project_name'),
        recursion_mode=pulumi.get(__ret__, 'recursion_mode'),
        region=pulumi.get(__ret__, 'region'),
        search_mode=pulumi.get(__ret__, 'search_mode'),
        tag_filters=pulumi.get(__ret__, 'tag_filters'),
        total_count=pulumi.get(__ret__, 'total_count'),
        vpc_id=pulumi.get(__ret__, 'vpc_id'),
        zid=pulumi.get(__ret__, 'zid'),
        zone_name=pulumi.get(__ret__, 'zone_name'))


@_utilities.lift_output_func(private_zones)
def private_zones_output(key_word: Optional[pulumi.Input[Optional[str]]] = None,
                         line_mode: Optional[pulumi.Input[Optional[int]]] = None,
                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                         project_name: Optional[pulumi.Input[Optional[str]]] = None,
                         recursion_mode: Optional[pulumi.Input[Optional[bool]]] = None,
                         region: Optional[pulumi.Input[Optional[str]]] = None,
                         search_mode: Optional[pulumi.Input[Optional[str]]] = None,
                         tag_filters: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['PrivateZonesTagFilterArgs']]]]] = None,
                         vpc_id: Optional[pulumi.Input[Optional[str]]] = None,
                         zid: Optional[pulumi.Input[Optional[int]]] = None,
                         zone_name: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[PrivateZonesResult]:
    """
    Use this data source to query detailed information of private zones
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    foo = volcengine.private_zone.get_private_zones(line_mode=3,
        recursion_mode=True,
        search_mode="EXACT",
        zid=770000,
        zone_name="volces.com")
    ```


    :param str key_word: The keyword of zone name.
    :param int line_mode: The line mode of Private Zone, specified whether the intelligent mode and the load balance function is enabled.
    :param str name_regex: A Name Regex of Resource.
    :param str output_file: File name where to save data source results.
    :param str project_name: The project name of the private zone.
    :param bool recursion_mode: Whether the recursion mode of Private Zone is enabled.
    :param str region: The region of Private Zone.
    :param str search_mode: The search mode of query. Valid values: `LIKE`, `EXACT`. Default is `LIKE`.
    :param Sequence[pulumi.InputType['PrivateZonesTagFilterArgs']] tag_filters: List of tag filters.
    :param str vpc_id: The vpc id associated to Private Zone.
    :param int zid: The zid of Private Zone.
    :param str zone_name: The name of Private Zone.
    """
    pulumi.log.warn("""private_zones is deprecated: volcengine.private_zone.PrivateZones has been deprecated in favor of volcengine.private_zone.getPrivateZones""")
    ...
