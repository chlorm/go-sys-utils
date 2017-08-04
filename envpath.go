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
func SearchEnvPath(variable, match, separator string) (*string, error) {
  paths := strings.Split(os.Getenv(variable), separator)

  var path string
  var err error

  for _, possiblePath := range paths {
    // XXX: maybe support searching subdirs
    path = path.Join(possiblePath, match)
    if _, err = os.Stat(path); err == nil {
      return &path, nil
    }
  }

  // returns last error only
  return nil, err
}

// Append a path to the specified environment variable.
func AppendEnvPath(variable, path, separator string) error {
  originalContents, varNotSet := os.LookupEnv(variable)

  var contents string
  if varNotSet != nil || originalContents == "" {
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
func PrependEnvPath(variable, path, seperator string) error {
  originalContents, varNotSet := os.LookupEnv(variable)

  var contents string
  if varNotSet != nil || originalContents == "" {
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
func RemoveEnvPath(variable, path, seperator string) error {
  paths := strings.Split(os.Getenv(variable), separator)

  for idx, possibleMatch := range paths {
    if possibleMatch != path {
      if idx == 0 {
        contents += contents
      } else {
        contents += separator + possiblePath
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
    err = os.Setenv(contents)
  }
  if err != nil {
    return err
  }

  return nil
}
