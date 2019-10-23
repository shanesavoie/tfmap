package main

import (
    "encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
    "gopkg.in/alecthomas/kingpin.v2"
)

var (
    env            = []string{}
    envMap         = make(map[string]string)
    whitelist      = kingpin.Flag("whitelist", "Prefix for whitelisting environment variables.").String()
    export         = kingpin.Flag("export", "Export to environment variable.").Default("TF_VARS_my_map").String()
)

func main() {
	log.SetFlags(0)
	kingpin.Parse()
    regMatch := regexp.MustCompile("^" + *whitelist + "(.*)")
    regReplace := regexp.MustCompile("^" + *whitelist)

	for _, e := range os.Environ() {
        env = append(env, e)
    }

    sort.Strings(env)

    // Convert env to a map so it can then be converted to JSON.
    for _, e := range env {
        envTmp := strings.SplitN(e, "=", 2)
        if(regMatch.MatchString(envTmp[0])) {
       	    envTmp[0] = regReplace.ReplaceAllString(envTmp[0], "")
            envMap[strings.ToLower(envTmp[0])] = envTmp[1]
        }
    }

    // convert map to JSON
    json, jsonErr := json.Marshal(envMap)
    if jsonErr != nil {
   	    log.Fatalf("error: failed to marshal to JSON: `%s`", jsonErr)
    }
    fmt.Printf("export %s=%s\n", *export, json)
}
