FROM golang:1.22
WORKDIR /app
COPY /app/go.mod .
COPY /app/main.go .
#RUN go mod download
#RUN echo "Mostrando archivos en /app:" && ls -la /app
RUN go build -o ms-adapter-tranfer-user .
ENTRYPOINT ["/app/ms-adapter-tranfer-user"]