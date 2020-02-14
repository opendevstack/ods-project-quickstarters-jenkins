package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	coreUtils "github.com/opendevstack/ods-core/tests/utils"
	v1 "github.com/openshift/api/build/v1"
	buildClientV1 "github.com/openshift/client-go/build/clientset/versioned/typed/build/v1"
	"io/ioutil"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"os"
	"strings"
	"time"
)

func RunJenkinsFile(repository string, repositoryProject string, branch string, projectName string, jenkinsFile string, jenkinsNamespace string, envVars ...coreUtils.EnvPair) error {
	values, err := ReadConfiguration()
	if err != nil {
		return err
	}

	request := coreUtils.RequestBuild{
		Repository: repository,
		Branch:     branch,
		Project:    repositoryProject,
		Env: append([]coreUtils.EnvPair{
			{
				Name:  "PROJECT_ID",
				Value: projectName,
			},
			{
				Name:  "CD_USER_TYPE",
				Value: "general",
			},
			{
				Name:  "CD_USER_ID_B64",
				Value: values["CD_USER_ID_B64"],
			},
			{
				Name:  "PIPELINE_TRIGGER_SECRET",
				Value: values["PIPELINE_TRIGGER_SECRET_B64"],
			},
			{
				Name:  "ODS_GIT_REF",
				Value: "cicdtests",
			},
			{
				Name:  "ODS_IMAGE_TAG",
				Value: values["ODS_IMAGE_TAG"],
			},
		}, envVars...),
	}

	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("Could not marchal json: %s", err)
	}

	pipelineNamePrefix := strings.ToLower(strings.Split(jenkinsFile, "/")[0])
	pipelineName := fmt.Sprintf("ods-corejob-%s-%s", pipelineNamePrefix, projectName)
	buildName := fmt.Sprintf("%s-%s-1", pipelineName, branch)

	println(buildName)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	response, err := http.Post(
		fmt.Sprintf("https://webhook-proxy-%s.172.17.0.1.nip.io/build?trigger_secret=%s&jenkinsfile_path=%s&component=%s",
			jenkinsNamespace,
			values["PIPELINE_TRIGGER_SECRET"],
			jenkinsFile,
			pipelineName),
		"application/json",
		bytes.NewBuffer(body))

	if err != nil {
		return err
	}

	if response.StatusCode >= http.StatusAccepted {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Could not post request: %s", string(bodyBytes))
	}

	if response.StatusCode >= http.StatusAccepted {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Could not post request: %s", string(bodyBytes))
	}

	config, err := coreUtils.GetOCClient()
	if err != nil {
		return fmt.Errorf("Error creating OC config: %s", err)
	}

	buildClient, err := buildClientV1.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("Error creating Build client: %s", err)
	}

	time.Sleep(10 * time.Second)
	build, err := buildClient.Builds(jenkinsNamespace).Get(buildName, metav1.GetOptions{})
	count := 0
	max := 240
	for (err != nil || build.Status.Phase == v1.BuildPhaseNew || build.Status.Phase == v1.BuildPhasePending || build.Status.Phase == v1.BuildPhaseRunning) && count < max {
		build, err = buildClient.Builds(jenkinsNamespace).Get(buildName, metav1.GetOptions{})
		time.Sleep(2 * time.Second)
		if err != nil {
			fmt.Printf("Build is still not available\n")
		} else {
			fmt.Printf("Waiting for build. Current status: %s\n", build.Status.Phase)
		}
		count++
	}

	stdout, stderr, err := coreUtils.RunCommand(
		"oc",
		[]string{
			"project", jenkinsNamespace,
		}, []string{})

	workspace, ok := os.LookupEnv("GITHUB_WORKSPACE")
	var script string
	if ok {
		script = fmt.Sprintf("%s/ods-core/tests/scripts/utils/print-jenkins-log.sh", workspace)
	} else {
		script = "../../ods-core/tests/scripts/utils/print-jenkins-log.sh"
	}

	stdout, stderr, err = coreUtils.RunCommand(
		script,
		[]string{
			buildName,
		}, []string{})

	if err != nil {
		panic(err)
	}

	if count >= max || build.Status.Phase != v1.BuildPhaseComplete {

		if count >= max {
			return fmt.Errorf(
				"Timeout during build: \nStdOut: %s\nStdErr: %s",
				stdout,
				stderr)
		} else {
			return fmt.Errorf(
				"Error during build: \nStdOut: %s\nStdErr: %s",
				stdout,
				stderr)
		}

	}

	return nil
}
