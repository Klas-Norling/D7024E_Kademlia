#FROM alpine:latest

# Use an official Ubuntu base image
FROM ubuntu:latest

FROM golang:1.23s

# Add the commands needed to put your compiled go binary in the container and
# run it when the container starts.
#
# See https://docs.docker.com/engine/reference/builder/ for a reference of all
# the commands you can use in this file.
#
# In order to use this file together with the docker-compose.yml file in the
# same directory, you need to ensure the image you build gets the name
# "kadlab", which you do by using the following command:
#
# $ docker build . -t kadlab

# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
RUN go mod download
 
# Builds your app with optional configuration
RUN go build -o /godocker
 
# Tells Docker which network port your container listens on
#sEXPOSE 8080
 
# Specifies the executable command that runs when the container starts
#CMD [ “/godocker” ]

 CMD /godocker



# Command to keep the container running
#CMD tail -f /dev/null

#CMD go run main.go
