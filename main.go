// Copyright © 2020 Ulrich Anhalt <ulrich.anhalt@gmail.com>
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
	"os"

	"github.com/ulranh/hana_sql_exporter/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "02-01-2006 15:04:05",
	})

	if debugLogsEnabled() {
		log.SetLevel(log.DebugLevel)
		log.Info("Debug logs enabled")
	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("Debug logs disabled")
	}

	cmd.Execute()
}

func debugLogsEnabled() bool {
	debugEnabled := os.Getenv("DEBUG_ENABLED")

	if debugEnabled == "true" {
		return true
	}

	return false
}
