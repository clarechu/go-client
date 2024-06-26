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
	"testing"
	"time"
)

/*
func TestArtifact(t *testing.T) {
	err := Artifact("xxx.xx.xx", "username", "password")
	if err != nil {
		t.Errorf("%v", err)
	}
}
*/

func TestSocket(t *testing.T) {
	type args struct {
		host     string
		username string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "aa",
			args: args{
				host: "127.0.0.1:8080",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Socket(tt.args.host, tt.args.username, tt.args.password); (err != nil) != tt.wantErr {

				t.Errorf("Socket() error = %v, wantErr %v", err, tt.wantErr)
			}
			time.Sleep(100 * time.Second)
		})
	}
}
