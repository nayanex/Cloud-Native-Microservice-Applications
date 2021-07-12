# set the base image. Since we're running 
# a Python application a Python base image is used
FROM golang:alpine
# set a key-value label for the Docker image
LABEL maintainer="Nayanexx.py"
# copy files from the host to the container filesystem. 
# For example, all the files in the current directory
# to the  `/app` directory in the container
COPY . /go/src/app
#  defines the working directory within the container
WORKDIR /go/src/app
# run commands within the container. 
# For example, invoke a pip command 
# to install dependencies defined in the requirements.txt file. 
RUN go mod init
RUN go build -o helloworld 

EXPOSE 6111
# provide a command to run on container start. 
# For example, start the `app.py` application.
CMD ["./helloworld" ]
