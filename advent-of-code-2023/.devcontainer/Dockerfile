FROM sapmachine:17-jdk-ubuntu

RUN apt-get update && \
    apt-get install -y git curl && \
    curl -fLo coursier https://github.com/coursier/launchers/raw/master/coursier && \
    chmod +x coursier && \
    ./coursier setup -y
