# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'BucketObjectsResult',
    'AwaitableBucketObjectsResult',
    'bucket_objects',
    'bucket_objects_output',
]

@pulumi.output_type
class BucketObjectsResult:
    """
    A collection of values returned by BucketObjects.
    """
    def __init__(__self__, bucket_name=None, id=None, name_regex=None, object_name=None, objects=None, output_file=None, total_count=None):
        if bucket_name and not isinstance(bucket_name, str):
            raise TypeError("Expected argument 'bucket_name' to be a str")
        pulumi.set(__self__, "bucket_name", bucket_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if object_name and not isinstance(object_name, str):
            raise TypeError("Expected argument 'object_name' to be a str")
        pulumi.set(__self__, "object_name", object_name)
        if objects and not isinstance(objects, list):
            raise TypeError("Expected argument 'objects' to be a list")
        pulumi.set(__self__, "objects", objects)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> Optional[str]:
        return pulumi.get(self, "object_name")

    @property
    @pulumi.getter
    def objects(self) -> Sequence['outputs.BucketObjectsObjectResult']:
        """
        The collection of TOS Object query.
        """
        return pulumi.get(self, "objects")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of TOS Object query.
        """
        return pulumi.get(self, "total_count")


class AwaitableBucketObjectsResult(BucketObjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return BucketObjectsResult(
            bucket_name=self.bucket_name,
            id=self.id,
            name_regex=self.name_regex,
            object_name=self.object_name,
            objects=self.objects,
            output_file=self.output_file,
            total_count=self.total_count)


def bucket_objects(bucket_name: Optional[str] = None,
                   name_regex: Optional[str] = None,
                   object_name: Optional[str] = None,
                   output_file: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableBucketObjectsResult:
    """
    Use this data source to query detailed information of tos objects
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tos.bucket_objects(bucket_name="test")
    ```


    :param str bucket_name: The name the TOS bucket.
    :param str name_regex: A Name Regex of TOS Object.
    :param str object_name: The name the TOS Object.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['bucketName'] = bucket_name
    __args__['nameRegex'] = name_regex
    __args__['objectName'] = object_name
    __args__['outputFile'] = output_file
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('volcengine:tos/bucketObjects:BucketObjects', __args__, opts=opts, typ=BucketObjectsResult).value

    return AwaitableBucketObjectsResult(
        bucket_name=__ret__.bucket_name,
        id=__ret__.id,
        name_regex=__ret__.name_regex,
        object_name=__ret__.object_name,
        objects=__ret__.objects,
        output_file=__ret__.output_file,
        total_count=__ret__.total_count)


@_utilities.lift_output_func(bucket_objects)
def bucket_objects_output(bucket_name: Optional[pulumi.Input[str]] = None,
                          name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                          object_name: Optional[pulumi.Input[Optional[str]]] = None,
                          output_file: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[BucketObjectsResult]:
    """
    Use this data source to query detailed information of tos objects
    ## Example Usage

    ```python
    import pulumi
    import pulumi_volcengine as volcengine

    default = volcengine.tos.bucket_objects(bucket_name="test")
    ```


    :param str bucket_name: The name the TOS bucket.
    :param str name_regex: A Name Regex of TOS Object.
    :param str object_name: The name the TOS Object.
    :param str output_file: File name where to save data source results.
    """
    ...