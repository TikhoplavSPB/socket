FROM golang:alpine as build
COPY . /go/src
RUN cd /go/src \
	&& go build -o app \
	&& cp app /bin/

FROM alpine
COPY --from=build /bin/app /bin
CMD ["/bin/app"]
