// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The Access Key for Volcengine Provider
func GetAccessKey(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "volcengine:accessKey")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "VOLCENGINE_ACCESS_KEY").(string)
}

// CUSTOMER ENDPOINTS for Volcengine Provider
func GetCustomerEndpoints(ctx *pulumi.Context) string {
	return config.Get(ctx, "volcengine:customerEndpoints")
}

// CUSTOMER HEADERS for Volcengine Provider
func GetCustomerHeaders(ctx *pulumi.Context) string {
	return config.Get(ctx, "volcengine:customerHeaders")
}

// Disable SSL for Volcengine Provider
func GetDisableSsl(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "volcengine:disableSsl")
}

// The Customer Endpoint for Volcengine Provider
func GetEndpoint(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "volcengine:endpoint")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "VOLCENGINE_ENDPOINT").(string)
}

// PROXY URL for Volcengine Provider
func GetProxyUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "volcengine:proxyUrl")
}

// The Region for Volcengine Provider
func GetRegion(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "volcengine:region")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "VOLCENGINE_REGION").(string)
}

// The Secret Key for Volcengine Provider
func GetSecretKey(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "volcengine:secretKey")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "VOLCENGINE_SECRET_KEY").(string)
}

// The Session Token for Volcengine Provider
func GetSessionToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "volcengine:sessionToken")
}