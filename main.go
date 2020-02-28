/*
 * Copyright 2018. Akamai Technologies, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"os"
	"os/exec"
	"flag"
	"fmt"
	"log"
	"encoding/json"
	"strings"

) 

type Input struct  {
	command string
}

type Output struct {
	Result string
}

func main() {

	var input Input

	err := json.NewDecoder(os.Stdin).Decode(&input)
	if err != nil {
		log.Fatal(err)
	}

        flag.Parse()    
        arg1 := strings.Join(flag.Args()[:1], " ")

	var output Output
	output.Result, err = RunCMD(arg1, flag.Args()[1:])
	if err != nil {
		log.Fatal(err)
	}
	jsonBody, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", jsonBody)

	
}

// RunCMD is a simple wrapper around terminal commands
func RunCMD(path string, args []string) (out string, err error) {

    cmd := exec.Command(path, args...)

    var b []byte
    b, err = cmd.CombinedOutput()
    if err != nil {
        log.Fatal(b, err)
    }
    out = strings.TrimSpace(string(b))

    return
}
