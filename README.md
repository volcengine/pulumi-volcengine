# Pulumi Volcengine Resource Provider

The Pulumi Volcengine Resource Provider lets you manage [Volcengine](https://www.volcengine.com/) resources.

## Installing

### Install volcengine provider

```bash
pulumi plugin install resource volcengine --server github://api.github.com/volcengine
```

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @volcengine/pulumi
```

or `yarn`:

```bash
yarn add @volcengine/pulumi
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_volcengine
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/volcengine/pulumi-volcengine/sdk/go/...
```

### .NET


To use from .NET, install using `dotnet add package`:


```bash

dotnet add package Volcengine.Pulumi.Volcengine

```

## Configuration

The following configuration points are available for the `volcengine` provider:

- `volcengine:accessKey` (environment: `VOLCENGINE_ACCESS_KEY`) - the API key for `volcengine`
- `volcengine:secretKey` (environment: `VOLCENGINE_SECRET_KEY`) - the API Secret Key for `volcengine`
- `volcengine:region` (environment: `VOLCENGINE_REGION`) - the region in which to deploy resources

[//]: # (## Reference)

[//]: # ()
[//]: # (For detailed reference documentation, please visit [the Pulumi registry]&#40;https://www.pulumi.com/registry/packages/volcengine/api-docs/&#41;.)
