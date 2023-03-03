package main

import (
	"net/http"
	"strings"
)

func init() {
	Register("create", Create{})
}

type Create struct{}

func (Create) Body() string {
	return `{
  "type": "apps.deployment",
  "metadata": {
    "namespace": "{{.Namespace}}",
    "labels": {
      "workload.user.cattle.io/workloadselector": "apps.deployment-{{.Namespace}}-{{.ProjectName}}"
    },
    "name": "{{.ProjectName}}"
  },
  "spec": {
    "replicas": 1,
    "template": {
      "spec": {
        "restartPolicy": "Always",
        "containers": [
          {
            "imagePullPolicy": "Always",
            "name": "container-0",
            "volumeMounts": [],
            "image": "{{.Images}}",
            "__active": true
          }
        ],
        "initContainers": [],
        "imagePullSecrets": [
          {
            "name": "{{.Secrets}}"
          }
        ],
        "volumes": [],
        "affinity": {}
      },
      "metadata": {
        "labels": {
          "workload.user.cattle.io/workloadselector": "apps.deployment-{{.Namespace}}-{{.ProjectName}}"
        }
      }
    },
    "selector": {
      "matchLabels": {
        "workload.user.cattle.io/workloadselector": "apps.deployment-{{.Namespace}}-{{.ProjectName}}"
      }
    }
  }
}`
}

func (Create) Url() string {
	return "https://{{.Domain}}/k8s/clusters/{{.Nm}}/v1/apps.deployments"
}

func (c Create) Req(url, data, account, secret string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(account, secret)
	return req, nil
}

/*

 todo publicEndpoints.load balancer

 "ports": [
  {
    "containerPort": "{{.Port}}",
    "hostPort": 0,
    "kind": "ClusterIP",
    "name": "{{.ProjectName}}-{{.Port}}",
    "protocol": "TCP",
    "sourcePort": 0,
    "type": "/v3/project/schemas/containerPort"
  }
]

*/
