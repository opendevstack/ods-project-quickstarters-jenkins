package docker_plain

import (
	"fmt"
	"testing"
	"runtime"
	"path/filepath"
	coreUtils "github.com/opendevstack/ods-core/tests/utils"
	utils "github.com/opendevstack/ods-quickstarters/tests/utils"
	"io/ioutil"
)

func TestJenkinsFile(t *testing.T) {

	values, err := utils.ReadConfiguration()
	if err != nil {
		t.Fatal(err)
	}

	_, filename, _, _ := runtime.Caller(0)
	quickstarterPath := filepath.Dir(filename)
	quickstarterName := filepath.Base(quickstarterPath)
	fmt.Printf("path:%s", quickstarterName)
	
	componentId := fmt.Sprintf("%s-test", quickstarterName)

	// cleanup and create bb resources for this test
	utils.CleanupAndCreateBitbucketProjectAndRepo(
		quickstarterName, componentId)

	// run provision job for quickstarter
	stages, err := utils.RunJenkinsFile(
		"ods-quickstarters",
		values["ODS_BITBUCKET_PROJECT"],
		values["ODS_GIT_REF"],
		coreUtils.PROJECT_NAME,
		fmt.Sprintf("%s/Jenkinsfile", quickstarterName),
		coreUtils.PROJECT_NAME_CD,
		coreUtils.EnvPair{
			Name:  "COMPONENT_ID",
			Value: componentId,
		},
		coreUtils.EnvPair{
			Name:  "GIT_URL_HTTP",
			Value: fmt.Sprintf("%s/%s/%s.git", values["REPO_BASE"], coreUtils.PROJECT_NAME, componentId),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Provision Build for %s returned:\n%s", componentId, stages)
	
	expected, err := ioutil.ReadFile("golden/jenkins-provision-stages.json")
	if err != nil {
		t.Fatal(err)
	}
	
	expectedAsString := string(expected)
	if stdout != expectedAsString {
		t.Fatalf("Actual jenkins stages from prov run: %s don't match -golden:\n'%s'\n-jenkins response:\n'%s'",
			componentId, expectedAsString, stdout)
	}

	// run master build of provisioned quickstarter in project's cd jenkins
	stages, err = utils.RunJenkinsFile(
		componentId,
		coreUtils.PROJECT_NAME,
		"master",
		coreUtils.PROJECT_NAME,
		"Jenkinsfile",
		coreUtils.PROJECT_NAME_CD,
		coreUtils.EnvPair{
			Name:  "COMPONENT_ID",
			Value: componentId,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Master (code) build for %s returned:\n%s", componentId, stages)

	expected, err = ioutil.ReadFile("golden/jenkins-build-stages.json")
	if err != nil {
		t.Fatal(err)
	}
	
	expectedAsString = string(expected)
	if stdout != expectedAsString {
		t.Fatalf("Actual jenkins stages from build run: %s don't match -golden:\n'%s'\n-jenkins response:\n'%s'",
			componentId, expectedAsString, stdout)
	}

	resourcesInTest := coreUtils.Resources{
		Namespace:         coreUtils.PROJECT_NAME_DEV,
		ImageTags:         []coreUtils.ImageTag{{Name: componentId, Tag: "latest"}},
		BuildConfigs:      []string{componentId},
		DeploymentConfigs: []string{componentId},
		Services:          []string{componentId},
		ImageStreams:      []string{componentId},
	}

	coreUtils.CheckResources(resourcesInTest, t)

}
