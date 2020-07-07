package utils

import (
	b64 "encoding/base64"
	"log"
	"path"
	"strings"
	"runtime"
	"fmt"
	coreUtils "github.com/opendevstack/ods-core/tests/utils"
)

func CleanupAndCreateBitbucketProjectAndRepo(quickstarter string, repoName string) {

	values, err := ReadConfiguration()
	if err != nil {
		log.Fatalf("Error reading ods-core.env: %s", err)
	}

	password, _ := b64.StdEncoding.DecodeString(values["CD_USER_PWD_B64"])

	stdout, stderr, err := RunScriptFromBaseDir("tests/scripts/setup_bitbucket_test_project.sh", []string{
		fmt.Sprintf("--bitbucket=%s", values["BITBUCKET_URL"]),
		fmt.Sprintf("--user=%s", values["CD_USER_ID"]),
		fmt.Sprintf("--password=%s", password),
		fmt.Sprintf("--project=%s", coreUtils.PROJECT_NAME),
		fmt.Sprintf("--repository=%s", repoName)},
		[]string{})
	if err != nil {
		fmt.Printf(
			"Execution of `setup_bitbucket_test_project.sh` failed: \nStdOut: %s\nStdErr: %s\nErr: %s\n",
			stdout,
			stderr,
			err)
	}

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)

	// provision build config
	buildConfigName := fmt.Sprintf("prov-%s-%s-%s", 
		quickstarter, coreUtils.PROJECT_NAME, strings.ReplaceAll(values["ODS_GIT_REF"], "/", "-")) 

	stdout, stderr, err = RunCommandWithWorkDir("oc", []string{
		"delete",
		"bc", buildConfigName,
		"-n", coreUtils.PROJECT_NAME_CD,
		}, dir, []string{})
	
	if err != nil {
		fmt.Printf("Error when deleting provisioning bc %s: %s, %s\n", 
			buildConfigName, err, stdout)
	} else {
		fmt.Printf("removed prov build config: %s\n", buildConfigName)
	}

	// quickstarter master branch build
	buildConfigName = fmt.Sprintf("run-%s-%s-master", repoName, coreUtils.PROJECT_NAME) 

	stdout, stderr, err = RunCommandWithWorkDir("oc", []string{
		"delete",
		"bc", buildConfigName,
		"-n", coreUtils.PROJECT_NAME_CD,
		}, dir, []string{})
	if err != nil {
		fmt.Printf("Error when deleting build bc %s: %s, %s\n", buildConfigName, err, stdout)
	} else {
		fmt.Printf("removed (build) build config: %s\n", buildConfigName)
	}

	fmt.Printf("Done\n - cleaned up and created repo: %s\n", repoName)
}
