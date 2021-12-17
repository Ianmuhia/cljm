FROM golang:1.17.5 as development
# Add a work directory
WORKDIR /app/cmd
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Install Reflex for development
RUN go install github.com/cespare/reflex@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest



# Expose port
EXPOSE 8000
# Start app
CMD reflex -g '*.go' go run *.go --start-service
