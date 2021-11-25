exactVersion=$(ls -lah /usr/lib/jvm | grep "java-11-openjdk-11.*\.x86_64" | awk '{print $NF}' | head -1) && \
alternatives --set java /usr/lib/jvm/${exactVersion}/bin/java && \
alternatives --set javac /usr/lib/jvm/${exactVersion}/bin/javac && \
#update-alternatives -s java-11-openjdk
#export JAVA_HOME=/usr/lib/jvm/java-11-openjdk/
#export PATH=$PATH:$JAVA_HOME
java -version
javac -version
