# Provider Sentry

`provider-sentry` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Sentry API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/justtrack/provider-sentry):
```
up ctp provider install justtrack/provider-sentry:v0.0.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-sentry
spec:
  package: justtrack/provider-sentry:v0.0.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/justtrack/provider-sentry).

## Developing

Run code-generation pipeline:
```console
make generate
```

Run against a Kubernetes cluster:

```console
make run
```

Build binary:

```console
make build
```

Publishing to upbound marketplace:

First you need to make sure that you are authenticated
```shell
up organization list
```

Now build the binary with
```shell
make build
```
and copy the path after `xpkg saved to ...`

Eventually publish the package with
```shell
up xpkg push justtrack/provider-sentry:vX.Y.Z -f $FILENAME
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/justtrack/provider-sentry/issues).
