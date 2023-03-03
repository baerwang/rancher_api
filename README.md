# use rancher api small tool

## build app

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o app
```

## create deployments command

```shell
./app -policy create -domain <rancher.example.com> -nm <xxxx> -ns dev -access <api-access> -secret <api-secret> -project <project-name> -images <example.com/master/project:version> -secrets <secrets>
```

## update deployments command

> nm and projectId field command line use the same value

```shell
./app -policy redeploy -domain <rancher.example.com> -nm <xx:xx> -ns <namespace> -access <api-access> -secret <api-secret> -project <project-name> -images <example.com/master/project:version> -secrets <secrets> -projectId <xx:xx>
```