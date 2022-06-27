FROM debian
COPY main main

RUN apt-get update
RUN apt-get install ca-certificates -y

ENTRYPOINT ["./main"]