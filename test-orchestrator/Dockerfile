FROM golang

RUN apt update && \
apt install openjdk-17-jdk maven wget -y

# ENV JUNIT_CONSOLE_LAUNCHER_VERSION="1.11.3"
# ENV JUNIT_CONSOLE_LAUNCHER_URL="https://repo1.maven.org/maven2/org/junit/platform/junit-platform-console-standalone/${JUNIT_CONSOLE_LAUNCHER_VERSION}/junit-platform-console-standalone-${JUNIT_CONSOLE_LAUNCHER_VERSION}.jar"

# RUN wget ${JUNIT_CONSOLE_LAUNCHER_URL}

WORKDIR /workdir

COPY . .

CMD [ "/bin/bash" ]