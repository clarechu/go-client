/*
Copyright 2020 The go-harbor Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
*/

package example

import (
	"context"
	"fmt"
	"github.com/clarechu/go-client/rest"
	"os"
)

func Get(host, username, password string) error {
	var result interface{}
	clientSet, err := rest.RESTClientFor(rest.NewDefaultConfig(host, username, password))
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	clientSet.Get().Resource(context.Background(), "projects").
		Name("demo").
		Do().
		Into(&result)

	return err
}

func Stream(host, username, password string) error {
	stdout := os.Stdout
	clientSet, err := rest.RESTClientFor(rest.NewDefaultConfig(host, username, password))
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return clientSet.Get().Resource(context.Background(), "projects").
		Name("demo").
		Stream(stdout)
}

func Socket(host, username, password string) error {
	stdout := os.Stdout
	stdin := os.Stdin
	clientSet, err := rest.RESTClientFor(rest.NewDefaultConfig(host, username, password).SetSchema(rest.WsSchema))
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return clientSet.Get().Resource(context.Background(), "echo").
		Websocket(stdin, stdout)
}
