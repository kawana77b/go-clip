FROM golang:1.23-bullseye

ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

# Create the user
RUN groupadd --gid $USER_GID $USERNAME \
  && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME 

RUN apt-get update && apt-get install -y \
  zip \
  unzip 

USER $USERNAME

WORKDIR /workspaces

RUN go install github.com/goreleaser/goreleaser/v2@latest && \
  go install github.com/spf13/cobra-cli@latest && \
  go install github.com/Songmu/gocredits/cmd/gocredits@latest

ENV SHELL=/bin/bash

CMD ["/bin/bash"]