make clean
make create-project TARGET_DIR=myTargetDir SPRING_BOOT_CLI_VERSION=2.3.1 PROJECT_ID=myproject COMPONENT_ID=mycomponent PACKAGE_NAME=mypackagename
make add-app-properties TARGET_DIR=myTargetDir
make customise-build-gradle SOURCE_DIR=. TARGET_DIR=myTargetDir
