// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vedb_mysql

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "volcengine:vedb_mysql/account:Account":
		r = &Account{}
	case "volcengine:vedb_mysql/allowlist:Allowlist":
		r = &Allowlist{}
	case "volcengine:vedb_mysql/allowlistAssociate:AllowlistAssociate":
		r = &AllowlistAssociate{}
	case "volcengine:vedb_mysql/backup:Backup":
		r = &Backup{}
	case "volcengine:vedb_mysql/database:Database":
		r = &Database{}
	case "volcengine:vedb_mysql/endpoint:Endpoint":
		r = &Endpoint{}
	case "volcengine:vedb_mysql/endpointPublicAddress:EndpointPublicAddress":
		r = &EndpointPublicAddress{}
	case "volcengine:vedb_mysql/instance:Instance":
		r = &Instance{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"volcengine",
		"vedb_mysql/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vedb_mysql/allowlist",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vedb_mysql/allowlistAssociate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vedb_mysql/backup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vedb_mysql/database",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vedb_mysql/endpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vedb_mysql/endpointPublicAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vedb_mysql/instance",
		&module{version},
	)
}