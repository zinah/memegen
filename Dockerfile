# The base go-image
FROM golang:1.14-alpine

# Create a directory for the app
RUN mkdir /memegen

# Copy all files from the current directory to the app directory
COPY . /memegen

# Set working directory
WORKDIR /memegen

# Run command as described:
# go build will build an executable file named server in the current directory
RUN go build -o memeserver .

# Run the server executable
ENTRYPOINT [ "/memegen/memeserver" ]
