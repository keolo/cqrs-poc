FROM scratch
EXPOSE 8080
ENTRYPOINT ["/cqrs-poc"]
COPY ./bin/ /