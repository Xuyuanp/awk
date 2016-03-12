/*
 * Copyright 2016 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package awk

import (
	"bufio"
	"io"
	"reflect"
	"strconv"
	"strings"
)

// Awk tool
func Awk(r io.Reader, delimiter string, f interface{}) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, delimiter)
		apply(f, words...)
	}
	return scanner.Err()
}

func apply(f interface{}, args ...string) {
	argvs := make([]reflect.Value, len(args))
	fv := reflect.ValueOf(f)
	ft := reflect.TypeOf(f)
	for ni := 0; ni < ft.NumIn(); ni++ {
		tin := ft.In(ni)
		var in interface{}
		switch tin.Kind() {
		case reflect.String:
			in = args[ni]
		case reflect.Int:
			in, _ = strconv.Atoi(args[ni])
		case reflect.Int64:
			in, _ = strconv.ParseInt(args[ni], 10, 64)
		case reflect.Float32:
			in64, _ := strconv.ParseFloat(args[ni], 32)
			in = float32(in64)
		case reflect.Float64:
			in, _ = strconv.ParseFloat(args[ni], 64)
		case reflect.Bool:
			in, _ = strconv.ParseBool(args[ni])
		}

		argvs[ni] = reflect.ValueOf(in)
	}
	fv.Call(argvs)
}
