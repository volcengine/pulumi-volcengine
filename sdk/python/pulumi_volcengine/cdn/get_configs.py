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
    'GetConfigsResult',
    'AwaitableGetConfigsResult',
    'get_configs',
    'get_configs_output',
]

@pulumi.output_type
class GetConfigsResult:
    """
    A collection of values returned by getConfigs.
    """
    def __init__(__self__, domain=None, domain_configs=None, id=None, output_file=None, total_count=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if domain_configs and not isinstance(domain_configs, list):
            raise TypeError("Expected argument 'domain_configs' to be a list")
        pulumi.set(__self__, "domain_configs", domain_configs)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if total_count and not isinstance(total_count, int):
            raise TypeError("Expected argument 'total_count' to be a int")
        pulumi.set(__self__, "total_count", total_count)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        The domain name.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="domainConfigs")
    def domain_configs(self) -> Sequence['outputs.GetConfigsDomainConfigResult']:
        """
        The collection of query.
        """
        return pulumi.get(self, "domain_configs")

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
    @pulumi.getter(name="totalCount")
    def total_count(self) -> int:
        """
        The total count of query.
        """
        return pulumi.get(self, "total_count")


class AwaitableGetConfigsResult(GetConfigsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConfigsResult(
            domain=self.domain,
            domain_configs=self.domain_configs,
            id=self.id,
            output_file=self.output_file,
            total_count=self.total_count)


def get_configs(domain: Optional[str] = None,
                output_file: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConfigsResult:
    """
    Use this data source to query detailed information of cdn configs
    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_volcengine as volcengine

    foo_cdn_certificate = volcengine.cdn.CdnCertificate("fooCdnCertificate",
        certificate="",
        private_key="",
        desc="tftest",
        source="cdn_cert_hosting")
    foo_cdn_domain = volcengine.cdn.CdnDomain("fooCdnDomain",
        domain="tftest.byte-test.com",
        service_type="web",
        tags=[
            volcengine.cdn.CdnDomainTagArgs(
                key="tfkey1",
                value="tfvalue1",
            ),
            volcengine.cdn.CdnDomainTagArgs(
                key="tfkey2",
                value="tfvalue2",
            ),
        ],
        domain_config=pulumi.Output.json_dumps({
            "OriginProtocol": "https",
            "Origin": [{
                "OriginAction": {
                    "OriginLines": [{
                        "Address": "1.1.1.1",
                        "HttpPort": "80",
                        "HttpsPort": "443",
                        "InstanceType": "ip",
                        "OriginType": "primary",
                        "PrivateBucketAccess": False,
                        "Weight": "2",
                    }],
                },
            }],
            "HTTPS": {
                "CertInfo": {
                    "CertId": foo_cdn_certificate.id,
                },
                "DisableHttp": False,
                "HTTP2": True,
                "Switch": True,
                "Ocsp": False,
                "TlsVersion": [
                    "tlsv1.1",
                    "tlsv1.2",
                ],
            },
        }))
    foo_configs = volcengine.cdn.get_configs_output(domain=foo_cdn_domain.id)
    ```


    :param str domain: The domain name.
    :param str output_file: File name where to save data source results.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('volcengine:cdn/getConfigs:getConfigs', __args__, opts=opts, typ=GetConfigsResult).value

    return AwaitableGetConfigsResult(
        domain=pulumi.get(__ret__, 'domain'),
        domain_configs=pulumi.get(__ret__, 'domain_configs'),
        id=pulumi.get(__ret__, 'id'),
        output_file=pulumi.get(__ret__, 'output_file'),
        total_count=pulumi.get(__ret__, 'total_count'))


@_utilities.lift_output_func(get_configs)
def get_configs_output(domain: Optional[pulumi.Input[str]] = None,
                       output_file: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConfigsResult]:
    """
    Use this data source to query detailed information of cdn configs
    ## Example Usage

    ```python
    import pulumi
    import json
    import pulumi_volcengine as volcengine

    foo_cdn_certificate = volcengine.cdn.CdnCertificate("fooCdnCertificate",
        certificate="",
        private_key="",
        desc="tftest",
        source="cdn_cert_hosting")
    foo_cdn_domain = volcengine.cdn.CdnDomain("fooCdnDomain",
        domain="tftest.byte-test.com",
        service_type="web",
        tags=[
            volcengine.cdn.CdnDomainTagArgs(
                key="tfkey1",
                value="tfvalue1",
            ),
            volcengine.cdn.CdnDomainTagArgs(
                key="tfkey2",
                value="tfvalue2",
            ),
        ],
        domain_config=pulumi.Output.json_dumps({
            "OriginProtocol": "https",
            "Origin": [{
                "OriginAction": {
                    "OriginLines": [{
                        "Address": "1.1.1.1",
                        "HttpPort": "80",
                        "HttpsPort": "443",
                        "InstanceType": "ip",
                        "OriginType": "primary",
                        "PrivateBucketAccess": False,
                        "Weight": "2",
                    }],
                },
            }],
            "HTTPS": {
                "CertInfo": {
                    "CertId": foo_cdn_certificate.id,
                },
                "DisableHttp": False,
                "HTTP2": True,
                "Switch": True,
                "Ocsp": False,
                "TlsVersion": [
                    "tlsv1.1",
                    "tlsv1.2",
                ],
            },
        }))
    foo_configs = volcengine.cdn.get_configs_output(domain=foo_cdn_domain.id)
    ```


    :param str domain: The domain name.
    :param str output_file: File name where to save data source results.
    """
    ...
