FROM sapmachine:21-jdk-ubuntu

RUN apt-get update && \
    apt-get install -y curl && \
    curl -fLo coursier https://github.com/coursier/launchers/raw/master/coursier && \
    chmod +x coursier && \
    ./coursier setup -y
