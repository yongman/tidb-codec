// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "fmt"
    "os"
    "encoding/hex"
    "strings"

    "github.com/pingcap/tidb/util/codec"
)

func main() {
    key := os.Args[1]
    fmt.Println("Origin key:\t", key)
    keyBytes, _ := hex.DecodeString(key)
    // encode
    encKey := codec.EncodeBytes([]byte{}, keyBytes)
    fmt.Println("Encoded key:\t", strings.ToUpper(hex.EncodeToString(encKey)))
    // decode
    _, decKey, err := codec.DecodeBytes(keyBytes, nil)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println("Decoded key:\t", strings.ToUpper(hex.EncodeToString(decKey)))
}
