FROM golang:1.22.5 as build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN go build -o /myapp ./cmd

FROM golang:1.22.5 as run

# Copy the application executable from the build image
COPY --from=build /myapp /myapp

WORKDIR /app
EXPOSE 5001
CMD ["/myapp"]
