FROM gradle:7.5.1-jdk17
ADD . /opt/app
WORKDIR /opt/app

ENTRYPOINT ["gradle"]
CMD ["--no-daemon", "functional:clean", "functional:test"]
