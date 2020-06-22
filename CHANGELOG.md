# Changelog

## Unreleased

## [2.0] - 2019-12-13

### Added
- Quickstarter-specific memory quotas ([#12](https://github.com/opendevstack/ods-quickstarters/issues/12))
- Quickstarter-specific CPU quotas ([#74](https://github.com/opendevstack/ods-quickstarters/issues/74))
- Add 'release-manager.yml' to each quickstarter ([#53](https://github.com/opendevstack/ods-quickstarters/issues/53))
- Enable WSGI for ds-ml-service quickstarter ([#37](https://github.com/opendevstack/ods-quickstarters/issues/37))
- Add central Tailorfile to easily compare resources ([#44](https://github.com/opendevstack/ods-quickstarters/issues/44))

### Changed
- Quickstarters have been renamed for more consistency when they were moved from `ods-project-quickstarters`
- Switch to OAuth proxy in jupyter-notebook and r-shiny quickstarters ([#46](https://github.com/opendevstack/ods-quickstarters/issues/46))
- Airflow Quickstarter fully provisioned in user's ODS project ([#60](https://github.com/opendevstack/ods-quickstarters/issues/60))
- Rename Airflow QuickStarter to `airflow-cluster` ([#76](https://github.com/opendevstack/ods-quickstarters/issues/76))
- Golang agent misses readme ([#64](https://github.com/opendevstack/ods-quickstarters/issues/64))
- Pass image tag and Git ref as params when provisioning quickstarters ([#41](https://github.com/opendevstack/ods-quickstarters/issues/41))
- fe-vue: Use nodejs10 agent for provisioning and building ([#32](https://github.com/opendevstack/ods-quickstarters/issues/32))
- fe-react: Generated Jenkinsfile should use nodejs10 agent image ([#31](https://github.com/opendevstack/ods-quickstarters/issues/31))
- fe-ionic: Use nodejs10 image for provisioning ([#30](https://github.com/opendevstack/ods-quickstarters/issues/30))
- e2e-cypress: Update to nodejs10 image ([#31](https://github.com/opendevstack/ods-quickstarters/issues/31))
- be-spring-boot: added springCliVersion, updated springframework to 2.2.1 ([#40](https://github.com/opendevstack/ods-quickstarters/pull/40))

### Fixed
- Wrong file permission stops Snyk cli from running in Python agent ([#67](https://github.com/opendevstack/ods-quickstarters/issues/67))
- Spring Boot quickstarter ignores property `no_nexus` ([#61](https://github.com/opendevstack/ods-quickstarters/issues/61))
- be-typescript-express: node version in deployment image doesn't match build image ([#8](https://github.com/opendevstack/ods-quickstarters/issues/8))

### Removed
- `NEXUS_HOST` param for component creation ([#70](https://github.com/opendevstack/ods-quickstarters/issues/70))
- Remove nodejs8 agent image ([#54](https://github.com/opendevstack/ods-quickstarters/issues/54))

## [1.2.0 ods-project-quickstarters] - 2019-10-10

### Added

- Add Go Quickstarter ([#255](https://github.com/opendevstack/ods-project-quickstarters/issues/255)).
- Enable xml unit - unit test results on all quickstarters and refactor build stage ([#299](https://github.com/opendevstack/ods-project-quickstarters/issues/299)).
- Airflow Cluster Quickstarter - pipeline ([#307](https://github.com/opendevstack/ods-project-quickstarters/issues/307)).
- ds-ml-service MRO ready ([#373](https://github.com/opendevstack/ods-project-quickstarters/issues/373)).
- Support Java 11 ([#309](https://github.com/opendevstack/ods-project-quickstarters/issues/309)).
- Clone-environment script should allow to pass branch and skip tagging flags ([#292](https://github.com/opendevstack/ods-project-quickstarters/issues/292)).
- ds-ml-service unified docker file ([#272](https://github.com/opendevstack/ods-project-quickstarters/issues/272)).
- Improve SonarQube support on FE quickstarters ([#212](https://github.com/opendevstack/ods-project-quickstarters/issues/212)).
- Create release manager quickstarter ([#391](https://github.com/opendevstack/ods-project-quickstarters/issues/391)).

### Changed

- Quickstarter `be-docker-plain` now builds a running container based on alpine instead of RHEL ([#260](https://github.com/opendevstack/ods-project-quickstarters/issues/260)).
- Quickstarter `ds-ml-service` with unified docker file ([#272](https://github.com/opendevstack/ods-project-quickstarters/issues/272)).
- Update jenkins agent for 1.2.x release ([#356](https://github.com/opendevstack/ods-project-quickstarters/issues/356)).
- Airflow quickstarter needs rundeck and resource limits update ([#358](https://github.com/opendevstack/ods-project-quickstarters/issues/358)).
- ds-ml-service - Train model pipeline step can fail after port forwarding of training service ([#269](https://github.com/opendevstack/ods-project-quickstarters/issues/269)).
- R quickstarter: standard keyserver port 11371 often blocked ([#298](https://github.com/opendevstack/ods-project-quickstarters/issues/298)).
- Replace upload-templates.sh with tailor ([#38](https://github.com/opendevstack/ods-project-quickstarters/issues/38)).
- Polish be-docker-plain ([#264](https://github.com/opendevstack/ods-project-quickstarters/issues/264)).

### Fixed

- Import of images into other cluster fails - because of missing role for default user ([#345](https://github.com/opendevstack/ods-project-quickstarters/issues/345)).
- fe-react: Make sure npm i is run within node docker image ([#363](https://github.com/opendevstack/ods-project-quickstarters/issues/363)).
- fe-react fails to install jest-junit dependency ([#361](https://github.com/opendevstack/ods-project-quickstarters/issues/361)).
- Enable Junit XML output for unit tests fo Vue QS ([#369](https://github.com/opendevstack/ods-project-quickstarters/issues/369)).
- Fix fe-angular quick starter karma config file manipulation ([#378](https://github.com/opendevstack/ods-project-quickstarters/issues/378)).
- Go quickstarter fails with no tests when collecting test results via junit ([#388](https://github.com/opendevstack/ods-project-quickstarters/issues/388)).
- Jupyter quickstarter app: Kernel won't start ([#268](https://github.com/opendevstack/ods-project-quickstarters/issues/268)).
- Error building Python based Quickstarters on OKD ([#295](https://github.com/opendevstack/ods-project-quickstarters/issues/295)).
- Reduce memory ratio ([#277](https://github.com/opendevstack/ods-project-quickstarters/issues/277)).
- be-docker-plain quickstarter fails in initial deployment ([#260](https://github.com/opendevstack/ods-project-quickstarters/issues/260)).
- be-python-flask: build fails when "python-ldap" is added to requirements.txt ([#250](https://github.com/opendevstack/ods-project-quickstarters/issues/250)).

## [1.1.0 ods-project-quickstarters] - 2019-05-28

### Added
- Rundeck `prepare-continous integration` job can now be used to upgrade an existing git repository ([#110](https://github.com/opendevstack/ods-project-quickstarters/pull/110))
- New quickstarter `be-docker-plain`: useful for starting with a plain `Dockerfile` and no BE/FE framework on top ([#97](https://github.com/opendevstack/ods-project-quickstarters/issues/97))
- Maven/Gradle Jenkins agent `jenkins-agent-maven` now gets Nexus credentials injected as server into `settings.xml` ([#127](https://github.com/opendevstack/ods-project-quickstarters/issues/127))
- New quickstarter `ds_ml_service` for machine learning from model training & testing to production ([#111](https://github.com/opendevstack/ods-project-quickstarters/issues/111))
- Quickstarter `be-python-flask` now provides coverage analysis data to SonarQube
- Quickstarter `fe-angular` now provides coverage analysis data to SonarQube and added SonarQube's linter rules for tslint
- Documentation of all quickstarters and agents added

### Changed
- Python quickstarter should use nexus as artifact repo ([#27](https://github.com/opendevstack/ods-project-quickstarters/issues/27))
- Jupyter & R-Shiny quickstarters are now based on new Openresty-based WAF image ([#103](https://github.com/opendevstack/ods-project-quickstarters/pull/103))
- NodeJS 10 Angular Jenkins agent `nodejs10-angular` replaces `nodejs8-angular` and supports nodeJS 10, Angular CLI 8.0.1 and cypress 3.3.1

### Fixed
- Rshiny quickstarter broken - due to refactoring and webhook proxy introduction ([#200](https://github.com/opendevstack/ods-project-quickstarters/issues/200)) & ([#184](https://github.com/opendevstack/ods-project-quickstarters/issues/184))
- Create-projects.sh seeds wrong jenkins SA rights & misses default SA for webhook proxy bug ([#189](https://github.com/opendevstack/ods-project-quickstarters/issus/189))
- import metadata: docker pull secrets are not created in an existing project - breaks oc import-image ([#202](https://github.com/opendevstack/ods-project-quickstarters/issues/202))
- Import Script is not replacing urls for sonarqube in DCs ([#145](https://github.com/opendevstack/ods-project-quickstarters/issues/145))

## [1.0.2 ods-project-quickstarters] - 2019-04-02

### Fixed
- Angular quickstarter `fe-angular-frontend` compilation failed due to changed dependency ([#129](https://github.com/opendevstack/ods-project-quickstarters/issues/129))
- Spring boot quickstarter `be-springboot` gradle build failed due to dependency update to gradle 4.10 ([#131](https://github.com/opendevstack/ods-project-quickstarters/issues/131))
- Upgrade of repo, thru rundeck job `prepare-continous integration` fails with invalid device ([#124](https://github.com/opendevstack/ods-project-quickstarters/issues/124))
- Jenkins `python agent` requires pip to have proper ssl validation configuration ([#176](https://github.com/opendevstack/ods-project-quickstarters/issues/176))

## [1.0.1 ods-project-quickstarters] - 2019-01-25

### Fixed
- Exclude images in `openshift` and `rhscl` namespace on import ([#102](https://github.com/opendevstack/ods-project-quickstarters/pull/102))
- Maven agent fails when proxy is configured due to invalid XML ([#108](https://github.com/opendevstack/ods-project-quickstarters/pull/108))


## [1.0.0 ods-project-quickstarters] - 2018-12-03

### Added
- Spring Boot Jenkins pipeline surfaces test results (#34)
- Jenkins webhook proxy templates (#81, #82)

### Changed
- Quickstarter build containers (located in the subdirs of https://github.com/opendevstack/ods-project-quickstarters/tree/master/boilerplates) inherit from corresponding Jenkins build agents now rather than replicating the setup
- Rundeck's OC container inherits from `jenkins-agent-base` now. The pull and tag is triggered thru *verify-rundeck-settings* rundeck job (#32)
- The build of a quickstarter component does not upload the tarball to Nexus anymore - instead it uses binary build configs (#9)
- The containers used to connect to openshift now pull the root ca during build, to ensure SSL trust (#12, #54)
- agents now support HTTP/S proxy - inject as ENV - with HTTP_PROXY, HTTPS_PROXY & NO_PROXY (#50)
- Python agent upgraded to 3.6 latest (#24)
- Maven agent now downloads Gradle 4.8.1 during build to increase build performance of components (#23)
- Scala agent now downloads sbt 1.1.6 / scala 2.12 - given an SBT bug - when proxy set, no NEXUS usage
- Update to newest cypress and TypeScript versions (#91)
- Build Jupyter/Rshiny via Jenkins (#92)

### Fixed
- Nodejs 8 quickstarter failed on npm run coverage (#22)
- Rundeck containers not cleaned up (#16, #17)
- Disable inclusion of Nginx server version in HTTP headers (#79)
- Jupyter: install from Nexus (#65)

### Removed
- Remove broken be-database quickstarter (#87)


## [0.1.0 ods-project-quickstarters] - 2018-07-27

Initial release.
