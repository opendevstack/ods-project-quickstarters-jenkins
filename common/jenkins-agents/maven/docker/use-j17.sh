#!/bin/bash

echo "Switching to java 17:"
exactVersion=$(ls -lah /usr/lib/jvm | grep "java-17-openjdk-17.*\.x86_64" | awk '{print $NF}' | head -1)
alternatives --set java /usr/lib/jvm/${exactVersion}/bin/java || exit 1
alternatives --set javac /usr/lib/jvm/${exactVersion}/bin/javac || exit 1

# Check it again
java -version | grep -q 17 || exit 1
javac -version | grep -q 17 || exit 1

java -version
javac -version
echo "JAVA_HOME: $JAVA_HOME"
