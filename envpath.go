// Copyright (c) 2017, Cody Opel <codyopel@gmail.com>
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

// Package for manipulating separator delimited paths in environment variables.

package sysutils

import (
  "os"
  "path"
  "strings"
)

// Check all paths returned by the specified environment variable for
// the existence of a match.
func SearchEnvPath(variable, separator, match string) (*string, error) {
  paths := strings.Split(os.Getenv(variable), separator)

  var possibleMatch string
  var err error

  for _, possiblePath := range paths {
    // XXX: maybe support searching subdirs
    possibleMatch = path.Join(possiblePath, match)
    if _, err = os.Stat(possibleMatch); err == nil {
      return &possibleMatch, nil
    }
  }

  // returns last error only
  return nil, err
}

// Append a path to the specified environment variable.
func AppendEnvPath(variable, separator, path string) error {
  originalContents, varSet := os.LookupEnv(variable)

  var contents string
  if varSet == false || originalContents == "" {
    contents = path
  } else {
    contents = path + separator + variable
  }

  if err := os.Setenv(variable, contents); err != nil {
    return err
  }

  return nil
}

// Prepends a path to the specified environment variable.
func PrependEnvPath(variable, seperator, path string) error {
  originalContents, varSet := os.LookupEnv(variable)

  var contents string
  if varSet == false || originalContents == "" {
    contents = path
  } else {
    contents = variable + seperator + path
  }

  if err := os.Setenv(variable, contents); err != nil {
    return err
  }

  return nil
}

// Removes a path from the specified environment variable.
func RemoveEnvPath(variable, separator, path string) error {
  paths := strings.Split(os.Getenv(variable), separator)

  var contents string
  for idx, possibleMatch := range paths {
    if possibleMatch != path {
      if idx == 0 {
        contents += contents
      } else {
        contents += separator + possibleMatch
      }
    } else {
      // skip matching strings to be removed
      continue
    }
  }

  var err error
  if contents == "" {
    err = os.Unsetenv(variable)
  } else {
    err = os.Setenv(variable, contents)
  }
  if err != nil {
    return err
  }

  return nil
}
