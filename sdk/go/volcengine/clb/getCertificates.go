// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

// Use this data source to query detailed information of certificates
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/clb"
//
// )
// func main() {
// pulumi.Run(func(ctx *pulumi.Context) error {
// var fooCertificate []*clb.Certificate
//
//	for index := 0; index < 3; index++ {
//	    key0 := index
//	    val0 := index
//
// __res, err := clb.NewCertificate(ctx, fmt.Sprintf("fooCertificate-%v", key0), &clb.CertificateArgs{
// CertificateName: pulumi.String(fmt.Sprintf("acc-test-certificate-%v", val0)),
// Description: pulumi.String("acc-test-demo"),
// PublicKey: pulumi.String(`-----BEGIN CERTIFICATE-----
// MIICWDCCAcGgAwIBAgIJAP7vOtjPtQIjMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
// BAYTAkNOMRMwEQYDVQQIDApjbi1iZWlqaW5nMSEwHwYDVQQKDBhJbnRlcm5ldCBX
// aWRnaXRzIFB0eSBMdGQwHhcNMjAxMDIwMDYxOTUxWhcNMjAxMTE5MDYxOTUxWjBF
// MQswCQYDVQQGEwJDTjETMBEGA1UECAwKY24tYmVpamluZzEhMB8GA1UECgwYSW50
// ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKB
// gQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9BVuFIBoU8nrP
// Y9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2CNIzxr9DjCzN5
// tWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQABo1AwTjAdBgNV
// HQ4EFgQUYDwuuqC2a2UPrfm1v31vE7+GRM4wHwYDVR0jBBgwFoAUYDwuuqC2a2UP
// rfm1v31vE7+GRM4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOBgQAovSB0
// 5JRKrg7lYR/KlTuKHmozfyL9UER0/dpTSoqsCyt8yc1BbtAKUJWh09BujBE1H22f
// lKvCAjhPmnNdfd/l9GrmAWNDWEDPLdUTkGSkKAScMpdS+mLmOBuYWgdnOtq3eQGf
// t07tlBL+dtzrrohHpfLeuNyYb40g8VQdp3RRRQ==
// -----END CERTIFICATE-----`),
// PrivateKey: pulumi.String(`-----BEGIN RSA PRIVATE KEY-----
// MIICXAIBAAKBgQDEdoyaJ0kdtjtbLRx5X9qwI7FblhJPRcScvhQSE8P5y/b/T8J9
// BVuFIBoU8nrPY9ABz4JFklZ6SznxLbFBqtXoJTmzV6ixyjjH+AGEw6hCiA8Pqy2C
// NIzxr9DjCzN5tWruiHqO60O3Bve6cHipH0VyLAhrB85mflvOZSH4xGsJkwIDAQAB
// AoGARe2oaCo5lTDK+c4Zx3392hoqQ94r0DmWHPBvNmwAooYd+YxLPrLMe5sMjY4t
// dmohnLNevCK1Uzw5eIX6BNSo5CORBcIDRmiAgwiYiS3WOv2+qi9g5uIdMiDr+EED
// K8wZJjB5E2WyfxL507vtW4T5L36yfr8SkmqH3GvzpI2jCqECQQDsy0AmBzyfK0tG
// Nw1+iF9SReJWgb1f5iHvz+6Dt5ueVQngrl/5++Gp5bNoaQMkLEDsy0iHIj9j43ji
// 0DON05uDAkEA1GXgGn8MXXKyuzYuoyYXCBH7aF579d7KEGET/jjnXx9DHcfRJZBY
// B9ghMnnonSOGboF04Zsdd3xwYF/3OHYssQJAekd/SeQEzyE5TvoQ8t2Tc9X4yrlW
// xNX/gmp6/fPr3biGUEtb7qi+4NBodCt+XsingmB7hKUP3RJTk7T2WnAC5wJAMqHi
// jY5x3SkFkHl3Hq9q2CKpQxUbCd7FXqg1wum/xj5GmqfSpNjHE3+jUkwbdrJMTrWP
// rmRy3tQMWf0mixAo0QJBAN4IcZChanq8cZyNqqoNbxGm4hkxUmE0W4hxHmLC2CYZ
// V4JpNm8dpi4CiMWLasF6TYlVMgX+aPxYRUWc/qqf1/Q=
// -----END RSA PRIVATE KEY-----`),
// ProjectName: pulumi.String("default"),
// Tags: clb.CertificateTagArray{
// &clb.CertificateTagArgs{
// Key: pulumi.String("k1"),
// Value: pulumi.String("v1"),
// },
// },
// })
// if err != nil {
// return err
// }
// fooCertificate = append(fooCertificate, __res)
// }
// _ = clb.GetCertificatesOutput(ctx, clb.GetCertificatesOutputArgs{
// Ids: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-functions-volcengine:clb-getCertificates:getCertificates.pp:16,9-29),
// }, nil);
// return nil
// })
// }
// ```
func GetCertificates(ctx *pulumi.Context, args *GetCertificatesArgs, opts ...pulumi.InvokeOption) (*GetCertificatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCertificatesResult
	err := ctx.Invoke("volcengine:clb/getCertificates:getCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesArgs struct {
	// The name of the Certificate.
	CertificateName *string `pulumi:"certificateName"`
	// The list of Certificate IDs.
	Ids []string `pulumi:"ids"`
	// The Name Regex of Certificate.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ProjectName of Certificate.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []GetCertificatesTag `pulumi:"tags"`
}

// A collection of values returned by getCertificates.
type GetCertificatesResult struct {
	// The name of the Certificate.
	CertificateName *string `pulumi:"certificateName"`
	// The collection of Certificate query.
	Certificates []GetCertificatesCertificate `pulumi:"certificates"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ProjectName of the Certificate.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []GetCertificatesTag `pulumi:"tags"`
	// The total count of Certificate query.
	TotalCount int `pulumi:"totalCount"`
}

func GetCertificatesOutput(ctx *pulumi.Context, args GetCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetCertificatesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCertificatesResult, error) {
			args := v.(GetCertificatesArgs)
			r, err := GetCertificates(ctx, &args, opts...)
			var s GetCertificatesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCertificatesResultOutput)
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesOutputArgs struct {
	// The name of the Certificate.
	CertificateName pulumi.StringPtrInput `pulumi:"certificateName"`
	// The list of Certificate IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Name Regex of Certificate.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ProjectName of Certificate.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Tags.
	Tags GetCertificatesTagArrayInput `pulumi:"tags"`
}

func (GetCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getCertificates.
type GetCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesResult)(nil)).Elem()
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutput() GetCertificatesResultOutput {
	return o
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutputWithContext(ctx context.Context) GetCertificatesResultOutput {
	return o
}

// The name of the Certificate.
func (o GetCertificatesResultOutput) CertificateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.CertificateName }).(pulumi.StringPtrOutput)
}

// The collection of Certificate query.
func (o GetCertificatesResultOutput) Certificates() GetCertificatesCertificateArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []GetCertificatesCertificate { return v.Certificates }).(GetCertificatesCertificateArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCertificatesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetCertificatesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o GetCertificatesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ProjectName of the Certificate.
func (o GetCertificatesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCertificatesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Tags.
func (o GetCertificatesResultOutput) Tags() GetCertificatesTagArrayOutput {
	return o.ApplyT(func(v GetCertificatesResult) []GetCertificatesTag { return v.Tags }).(GetCertificatesTagArrayOutput)
}

// The total count of Certificate query.
func (o GetCertificatesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetCertificatesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCertificatesResultOutput{})
}
