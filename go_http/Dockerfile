FROM fedora
WORKDIR /app
COPY *.go ./
EXPOSE 8080
RUN dnf -y update
RUN dnf -y install go
RUN go build go_http.go
CMD ./go_http