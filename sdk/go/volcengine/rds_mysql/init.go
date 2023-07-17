// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "volcengine:rds_mysql/account:Account":
		r = &Account{}
	case "volcengine:rds_mysql/allowlist:Allowlist":
		r = &Allowlist{}
	case "volcengine:rds_mysql/allowlistAssociate:AllowlistAssociate":
		r = &AllowlistAssociate{}
	case "volcengine:rds_mysql/database:Database":
		r = &Database{}
	case "volcengine:rds_mysql/instance:Instance":
		r = &Instance{}
	case "volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode":
		r = &InstanceReadonlyNode{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := volcengine.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"volcengine",
		"rds_mysql/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"rds_mysql/allowlist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"rds_mysql/allowlistAssociate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"rds_mysql/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"rds_mysql/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"rds_mysql/instanceReadonlyNode",
		&module{version},
	)
}