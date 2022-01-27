# make clean TARGET_DIR=myTargetDir
make create-project TARGET_DIR=myTargetDir SPRING_BOOT_CLI_VERSION=2.6.3 PROJECT_ID=bix COMPONENT_ID=bix PACKAGE_NAME=bix
make add-app-properties TARGET_DIR=myTargetDir
make customise-build-gradle SOURCE_DIR=. TARGET_DIR=myTargetDir
