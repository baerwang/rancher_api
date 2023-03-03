package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
)

func main() {
	var policy, domain, nm, namespace, access, secret, project, projectId, images, secrets, port string

	var isTls bool

	// todo use cobra optimization command
	flag.StringVar(&policy, "policy", "", "create or redeploy")
	flag.StringVar(&domain, "domain", "", "")
	flag.StringVar(&nm, "nm", "", "")
	flag.StringVar(&access, "access", "", "access_key")
	flag.StringVar(&secret, "secret", "", "secret_key")
	flag.StringVar(&namespace, "ns", "default", "namespace")
	flag.StringVar(&project, "project", "", "project.name")
	flag.StringVar(&projectId, "projectId", "", "project.id")
	flag.StringVar(&images, "images", "", "images")
	flag.StringVar(&secrets, "secrets", "", "secrets")
	flag.StringVar(&port, "port", "", "port")
	flag.BoolVar(&isTls, "tls", true, "tls")
	flag.Parse()

	a, ok := Actions[policy]
	if !ok {
		panic("policy not is exist [create or redeploy]")
	}

	rancher := Rancher{Domain: domain, Nm: nm, Access: access, Secret: secret,
		Namespace: namespace, ProjectName: project, Port: port, ProjectId: projectId, Images: images, Secrets: secrets}

	url := parse(rancher, a.Url())
	data := parse(rancher, a.Body())

	req, err := a.Req(url, data, access, secret)
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	if isTls {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	do, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer do.Body.Close()

	fmt.Println("status code:", do.StatusCode)

	resp := bytes.Buffer{}
	if _, err = io.Copy(&resp, do.Body); err != nil {
		panic(err)
	}
	fmt.Println("resp:", resp.String())
}

func parse(r Rancher, data string) string {
	var (
		sb = strings.Builder{}
		t2 = template.Must(template.New("rancher").Parse(data))
	)

	if err := t2.Execute(&sb, r); err != nil {
		panic(err)
	}

	return sb.String()
}
