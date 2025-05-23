// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccountArgs, AccountState } from "./account";
export type Account = import("./account").Account;
export const Account: typeof import("./account").Account = null as any;
utilities.lazyLoad(exports, ["Account"], () => require("./account"));

export { AccountsArgs, AccountsResult, AccountsOutputArgs } from "./accounts";
export const accounts: typeof import("./accounts").accounts = null as any;
export const accountsOutput: typeof import("./accounts").accountsOutput = null as any;
utilities.lazyLoad(exports, ["accounts","accountsOutput"], () => require("./accounts"));

export { AllowlistArgs, AllowlistState } from "./allowlist";
export type Allowlist = import("./allowlist").Allowlist;
export const Allowlist: typeof import("./allowlist").Allowlist = null as any;
utilities.lazyLoad(exports, ["Allowlist"], () => require("./allowlist"));

export { AllowlistAssociateArgs, AllowlistAssociateState } from "./allowlistAssociate";
export type AllowlistAssociate = import("./allowlistAssociate").AllowlistAssociate;
export const AllowlistAssociate: typeof import("./allowlistAssociate").AllowlistAssociate = null as any;
utilities.lazyLoad(exports, ["AllowlistAssociate"], () => require("./allowlistAssociate"));

export { AllowlistsArgs, AllowlistsResult, AllowlistsOutputArgs } from "./allowlists";
export const allowlists: typeof import("./allowlists").allowlists = null as any;
export const allowlistsOutput: typeof import("./allowlists").allowlistsOutput = null as any;
utilities.lazyLoad(exports, ["allowlists","allowlistsOutput"], () => require("./allowlists"));

export { DatabaseArgs, DatabaseState } from "./database";
export type Database = import("./database").Database;
export const Database: typeof import("./database").Database = null as any;
utilities.lazyLoad(exports, ["Database"], () => require("./database"));

export { DatabasesArgs, DatabasesResult, DatabasesOutputArgs } from "./databases";
export const databases: typeof import("./databases").databases = null as any;
export const databasesOutput: typeof import("./databases").databasesOutput = null as any;
utilities.lazyLoad(exports, ["databases","databasesOutput"], () => require("./databases"));

export { GetAccountsArgs, GetAccountsResult, GetAccountsOutputArgs } from "./getAccounts";
export const getAccounts: typeof import("./getAccounts").getAccounts = null as any;
export const getAccountsOutput: typeof import("./getAccounts").getAccountsOutput = null as any;
utilities.lazyLoad(exports, ["getAccounts","getAccountsOutput"], () => require("./getAccounts"));

export { GetAllowlistsArgs, GetAllowlistsResult, GetAllowlistsOutputArgs } from "./getAllowlists";
export const getAllowlists: typeof import("./getAllowlists").getAllowlists = null as any;
export const getAllowlistsOutput: typeof import("./getAllowlists").getAllowlistsOutput = null as any;
utilities.lazyLoad(exports, ["getAllowlists","getAllowlistsOutput"], () => require("./getAllowlists"));

export { GetDatabasesArgs, GetDatabasesResult, GetDatabasesOutputArgs } from "./getDatabases";
export const getDatabases: typeof import("./getDatabases").getDatabases = null as any;
export const getDatabasesOutput: typeof import("./getDatabases").getDatabasesOutput = null as any;
utilities.lazyLoad(exports, ["getDatabases","getDatabasesOutput"], () => require("./getDatabases"));

export { GetInstancesArgs, GetInstancesResult, GetInstancesOutputArgs } from "./getInstances";
export const getInstances: typeof import("./getInstances").getInstances = null as any;
export const getInstancesOutput: typeof import("./getInstances").getInstancesOutput = null as any;
utilities.lazyLoad(exports, ["getInstances","getInstancesOutput"], () => require("./getInstances"));

export { GetSchemasArgs, GetSchemasResult, GetSchemasOutputArgs } from "./getSchemas";
export const getSchemas: typeof import("./getSchemas").getSchemas = null as any;
export const getSchemasOutput: typeof import("./getSchemas").getSchemasOutput = null as any;
utilities.lazyLoad(exports, ["getSchemas","getSchemasOutput"], () => require("./getSchemas"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { InstanceReadonlyNodeArgs, InstanceReadonlyNodeState } from "./instanceReadonlyNode";
export type InstanceReadonlyNode = import("./instanceReadonlyNode").InstanceReadonlyNode;
export const InstanceReadonlyNode: typeof import("./instanceReadonlyNode").InstanceReadonlyNode = null as any;
utilities.lazyLoad(exports, ["InstanceReadonlyNode"], () => require("./instanceReadonlyNode"));

export { InstancesArgs, InstancesResult, InstancesOutputArgs } from "./instances";
export const instances: typeof import("./instances").instances = null as any;
export const instancesOutput: typeof import("./instances").instancesOutput = null as any;
utilities.lazyLoad(exports, ["instances","instancesOutput"], () => require("./instances"));

export { SchemaArgs, SchemaState } from "./schema";
export type Schema = import("./schema").Schema;
export const Schema: typeof import("./schema").Schema = null as any;
utilities.lazyLoad(exports, ["Schema"], () => require("./schema"));

export { SchemasArgs, SchemasResult, SchemasOutputArgs } from "./schemas";
export const schemas: typeof import("./schemas").schemas = null as any;
export const schemasOutput: typeof import("./schemas").schemasOutput = null as any;
utilities.lazyLoad(exports, ["schemas","schemasOutput"], () => require("./schemas"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:rds_postgresql/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "volcengine:rds_postgresql/allowlist:Allowlist":
                return new Allowlist(name, <any>undefined, { urn })
            case "volcengine:rds_postgresql/allowlistAssociate:AllowlistAssociate":
                return new AllowlistAssociate(name, <any>undefined, { urn })
            case "volcengine:rds_postgresql/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "volcengine:rds_postgresql/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "volcengine:rds_postgresql/instanceReadonlyNode:InstanceReadonlyNode":
                return new InstanceReadonlyNode(name, <any>undefined, { urn })
            case "volcengine:rds_postgresql/schema:Schema":
                return new Schema(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "rds_postgresql/account", _module)
pulumi.runtime.registerResourceModule("volcengine", "rds_postgresql/allowlist", _module)
pulumi.runtime.registerResourceModule("volcengine", "rds_postgresql/allowlistAssociate", _module)
pulumi.runtime.registerResourceModule("volcengine", "rds_postgresql/database", _module)
pulumi.runtime.registerResourceModule("volcengine", "rds_postgresql/instance", _module)
pulumi.runtime.registerResourceModule("volcengine", "rds_postgresql/instanceReadonlyNode", _module)
pulumi.runtime.registerResourceModule("volcengine", "rds_postgresql/schema", _module)
