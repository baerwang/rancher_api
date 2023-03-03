package main

import (
	"net/http"
	"strings"
)

func init() {
	Register("redeploy", Redeploy{})
}

type Redeploy struct{}

func (Redeploy) Body() string {
	return `{
  "containers": [
    {
      "image": "{{.Images}}",
      "imagePullPolicy": "Always",
      "initContainer": false,
      "name": "container-0",
      "ports": [],
      "resources": {
        "type": "/v3/project/schemas/resourceRequirements"
      },
      "restartCount": 0,
      "stdin": false,
      "stdinOnce": false,
      "terminationMessagePath": "/dev/termination-log",
      "terminationMessagePolicy": "File",
      "tty": false,
      "type": "/v3/project/schemas/container"
    }
  ],
  "imagePullSecrets": [
    {
      "name": "{{.Secrets}}",
      "type": "/v3/project/schemas/localObjectReference"
    }
  ],
  "labels": {
    "workload.user.cattle.io/workloadselector": "apps.deployment-{{.Namespace}}-{{.ProjectName}}"
  },
  "name": "{{.ProjectName}}",
  "namespaceId": "{{.Namespace}}",
  "projectId": "{{.ProjectId}}",
  "restartPolicy": "Always",
  "scale": 1,
  "scheduling": {
    "scheduler": "default-scheduler"
  },
  "selector": {
    "matchLabels": {
      "workload.user.cattle.io/workloadselector": "apps.deployment-{{.Namespace}}-{{.ProjectName}}"
    },
    "type": "/v3/project/schemas/labelSelector"
  },
  "state": "active",
  "volumes": [],
  "workloadAnnotations": {
    "deployment.kubernetes.io/revision": "3"
  },
  "workloadLabels": {
    "workload.user.cattle.io/workloadselector": "apps.deployment-{{.Namespace}}-{{.ProjectName}}"
  },
  "workloadMetrics": []
}`
}

func (Redeploy) Url() string {
	return "https://{{.Domain}}/v3/project/{{.Nm}}/workloads/deployment:{{.Namespace}}:{{.ProjectName}}"
}

func (r Redeploy) Req(url, data, account, secret string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(account, secret)
	return req, nil
}
