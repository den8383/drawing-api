FROM node:12-buster




RUN npm i rimraf -g
RUN curl -L -O https://github.com/protocolbuffers/protobuf/releases/download/v3.11.2/protoc-3.11.2-linux-x86_64.zip
RUN curl -L -O https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-linux-x86_64
RUN unzip protoc-3.11.2-linux-x86_64.zip && cp ./bin/protoc /usr/local/bin/. && chmod +x /usr/local/bin/protoc
RUN cp protoc-gen-grpc-web-1.0.7-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web && chmod +x /usr/local/bin/protoc-gen-grpc-web

RUN apt update && apt upgrade -y
RUN apt install golang -y
RUN go mod init protobuf
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

USER node
