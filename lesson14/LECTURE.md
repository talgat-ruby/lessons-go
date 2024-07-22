# Lesson 14: Docker

First please all install [docker desktop](https://www.docker.com/products/docker-desktop/). It will install:

1. **Docker Engine** 
2. **Docker CLI client** 
3. Docker Scout (additional subscription may apply)
4. Docker Build 
5. Docker Extensions 
6. **Docker Compose** 
7. Docker Content Trust 
8. Kubernetes
9. Credential Helper

In order to check if you have successfully installed docker desktop run the following command:

```shell
$ docker version
Client:
 Version:           27.0.3
 API version:       1.46
 Go version:        go1.21.11
 Git commit:        7d4bcd8
 Built:             Fri Jun 28 23:59:41 2024
 OS/Arch:           darwin/arm64
 Context:           desktop-linux

Server: Docker Desktop 4.32.0 (157355)
 Engine:
  Version:          27.0.3
  API version:      1.46 (minimum version 1.24)
  Go version:       go1.21.11
  Git commit:       662f78c
  Built:            Sat Jun 29 00:02:44 2024
  OS/Arch:          linux/arm64
  Experimental:     false
 containerd:
  Version:          1.7.18
  GitCommit:        ae71819c4f5e67bb4d5ae76a6b735f29cc25774e
 runc:
  Version:          1.7.18
  GitCommit:        v1.1.13-0-g58aa920
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
```

## Image

Images are objects that contain everything an app needs to run. 
This includes an OS filesystem, the application, and all dependencies. 
If you’re a OOP developer, they’re similar to **classes**.

Check your images with:

```shell
$ docker images
```

Let's run our first container:

```shell
$ docker run --name test -it ubuntu:latest bash
```
