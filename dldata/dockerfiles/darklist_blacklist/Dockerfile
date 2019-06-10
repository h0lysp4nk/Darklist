# Specify docker image
FROM golang:alpine

# Specify maintainer
MAINTAINER Holy, https://github.com/h0lysp4nk

# Specify user
USER root

# Add basics first
RUN apk update && apk upgrade && apk add \
	bash 
# Remove APK cache
RUN rm -rf /var/cache/apk/*

# Create Golang folder
RUN mkdir -p /bootstrap && mkdir -p /go/bin && mkdir -p /go/pkg

# Copy script
COPY ./start.sh /bootstrap/start.sh

# Set permissions
RUN chmod -R 755 /bootstrap && chmod -R 755 /go

# Run the application
ENTRYPOINT ["/bootstrap/start.sh"]

