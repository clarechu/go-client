# go-client

简单实现远程调用

```go
package main

import (
	"fmt"
	"github.com/clarechu/go-client/rest"
)

// Get 请求 
func Get(host, username, password string) error {
	var result interface{}
	clientSet, err := rest.RESTClientFor(rest.NewDefaultConfig(host, username, password))
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	clientSet.Get().Resource("projects").
		Name("demo").
		Do().
		Into(&result)

	return err
}


```