/*
 * Copyright 2020 Aletheia Ware LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package network

import (
	"fmt"
	"github.com/AletheiaWareLLC/perusergo"
	"io"
	"log"
	"net/http"
	"strings"
)

type tcpNetwork struct {
	//
}

func NewTCPNetwork() perusergo.Network {
	return &tcpNetwork{}
}

func (n *tcpNetwork) Get(address string) (string, io.ReadCloser, error) {
	response, err := http.Get(address)
	if err != nil {
		return "", nil, err
	}
	log.Println("Response:", response)
	switch response.StatusCode {
	case http.StatusOK:
		mime := response.Header.Get("Content-Type")
		mime = strings.Split(mime, ";")[0]
		return mime, response.Body, nil
	}
	return "", nil, fmt.Errorf("HTTP Get: %s", response.Status)
}
