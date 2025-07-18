// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vmp

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
	case "volcengine:vmp/alertingRule:AlertingRule":
		r = &AlertingRule{}
	case "volcengine:vmp/contact:Contact":
		r = &Contact{}
	case "volcengine:vmp/contactGroup:ContactGroup":
		r = &ContactGroup{}
	case "volcengine:vmp/notifyGroupPolicy:NotifyGroupPolicy":
		r = &NotifyGroupPolicy{}
	case "volcengine:vmp/notifyPolicy:NotifyPolicy":
		r = &NotifyPolicy{}
	case "volcengine:vmp/notifyTemplate:NotifyTemplate":
		r = &NotifyTemplate{}
	case "volcengine:vmp/ruleFile:RuleFile":
		r = &RuleFile{}
	case "volcengine:vmp/workspace:Workspace":
		r = &Workspace{}
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
		"vmp/alertingRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vmp/contact",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vmp/contactGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vmp/notifyGroupPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vmp/notifyPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vmp/notifyTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vmp/ruleFile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"vmp/workspace",
		&module{version},
	)
}
