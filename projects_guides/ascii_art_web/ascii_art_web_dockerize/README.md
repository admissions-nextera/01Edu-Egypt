# ASCII-Art-Web-Dockerize Project Guide

> **Before you start:** Install Docker and run `docker run hello-world` to confirm it works. Read what Docker is and why it exists before writing a single line. You cannot containerize something you do not understand.

---

## Objectives

By completing this project you will learn:

1. **Containerization** — What a container is and how it differs from a virtual machine
2. **Docker Images** — Building a reproducible, portable snapshot of your application
3. **Dockerfile** — Writing instructions that turn your code into an image
4. **Docker Labels** — Attaching metadata to Docker objects
5. **Port Mapping** — Exposing a container's internal port to the host
6. **Garbage Collection** — Understanding and cleaning up unused Docker objects

---

## Prerequisites — Topics You Must Know Before Starting

### 1. ASCII-Art-Web or ASCII-Art-Stylize (Completed)
- Working Go web server with GET `/` and POST `/ascii-art`

### 2. Docker Concepts
- What is a container?
- What is an image?
- What is the difference between `docker build` and `docker run`?
- Search: **"Docker explained for beginners"**
- https://docs.docker.com/get-started/

### 3. Dockerfile Syntax
- `FROM`, `WORKDIR`, `COPY`, `RUN`, `EXPOSE`, `CMD`
- What each instruction does and in what order they run
- Search: **"Dockerfile instructions explained"**
- https://docs.docker.com/engine/reference/builder/

### 4. Docker CLI
- `docker build`, `docker run`, `docker ps`, `docker stop`, `docker rm`, `docker images`
- Search: **"docker CLI cheat sheet"**

**Read before starting:**
- https://docs.docker.com/develop/develop-images/dockerfile_best-practices/
- https://docs.docker.com/config/labels-custom-metadata/

---

## Project Structure

```
ascii-art-web-dockerize/
├── main.go
├── handlers.go
├── banner.go
├── templates/
│   └── index.html
├── static/
│   └── style.css
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
├── Dockerfile
└── go.mod
```

---

## Milestone 1 — Understand What You Are Containerizing

**This milestone has no code.** Answer these questions before writing a Dockerfile.

**Questions to answer:**
- What does your Go server need to run? (Binary? Files? Which ones?)
- Which files must be inside the container for the server to work correctly?
- What port does your server listen on inside the container?
- What is the difference between `go build` and `go run` — and which is better inside a container?
- What is a multi-stage Docker build and why does it produce a smaller image?

Search: **"Docker multi-stage build Go"** — read and understand it before moving on.

---

## Milestone 2 — Write the Dockerfile

**Goal:** `docker build` produces an image that runs your server.

**Questions to answer:**
- Which base image should you use for the build stage? Which for the final stage?
- Why is `golang:alpine` better than `golang:latest` for production images?
- What is the correct order of `COPY` and `RUN go build` — and why does order matter for layer caching?
- Which files need to be copied into the final image? (Not just the binary — think about what your server reads at runtime.)
- What does `EXPOSE` do — and what does it NOT do?

**Code Placeholder:**
```dockerfile
# Dockerfile

# --- Stage 1: Build ---
# Choose a Go base image

# Set the working directory inside the container

# Copy go.mod and go.sum first (why before copying the rest?)
# Download dependencies

# Copy the rest of the source code
# Build the binary with the correct output name

# --- Stage 2: Run ---
# Choose a minimal base image (e.g. alpine or scratch)

# Set working directory

# Copy the binary from the build stage
# Copy the banner files and templates into the image

# Add metadata labels (author, version, description)

# Expose the port your server listens on

# Define the command to run when the container starts
```

**Resources:**
- https://docs.docker.com/develop/develop-images/dockerfile_best-practices/
- Search: **"golang docker multi-stage build example"**

**Verify:**
```bash
docker build -t ascii-art-web .
# Should complete without errors
docker images
# Your image should appear in the list
```

---

## Milestone 3 — Add Labels

**Goal:** Your image and container carry metadata as Docker labels.

**Questions to answer:**
- What are Docker labels and what format do they use?
- Which labels should your image have? (At minimum: author, version, description)
- Where in the Dockerfile do you add labels?

**Code Placeholder:**
```dockerfile
# Add to your Dockerfile after the FROM in the final stage

LABEL maintainer="..."
LABEL version="..."
LABEL description="..."
```

**Verify:**
```bash
docker inspect ascii-art-web | grep -A 10 "Labels"
# Should show your labels
```

---

## Milestone 4 — Run the Container

**Goal:** `docker run` starts the container and the web server is accessible from your browser.

**Questions to answer:**
- What does `-p 8080:8080` mean in a `docker run` command?
- What is the difference between running in the foreground and with `-d` (detached)?
- How do you view logs from a running container?
- How do you stop a running container?

**Verify:**
```bash
docker run -p 8080:8080 ascii-art-web
```
Open `http://localhost:8080` in your browser. The site must work exactly as it did before containerization.

---

## Milestone 5 — Garbage Collection

**Goal:** You understand and can clean up unused Docker objects.

**Questions to answer:**
- What are dangling images? How do they appear?
- What command removes all stopped containers?
- What command removes all unused images?
- What does `docker system prune` do — and what does the `-a` flag add?
- Why is it important to clean up unused objects in production?

**Verify:**
```bash
docker ps -a        # list all containers including stopped ones
docker images -a    # list all images including intermediate ones
docker system df    # show disk usage by Docker objects
```

Run a build several times and watch dangling images accumulate. Then clean them up. Understand what you are deleting before you delete it.

**Resource:** Search: **"docker garbage collection prune"**

---

## Milestone 6 — Dockerfile Best Practices Review

Before submission, go through your Dockerfile and answer these:

- Are your `COPY` instructions ordered so that rarely-changing files (go.mod) come before frequently-changing ones (source code)? This maximizes layer cache reuse.
- Is your final image as small as possible? Did you use a multi-stage build?
- Does your container run as root? (It should not in production — search: **"Docker non-root user"**)
- Is there a `.dockerignore` file that prevents unnecessary files from being sent to the build context?

**Code Placeholder:**
```
# .dockerignore

# Exclude files that should not be in the build context
# (what belongs here?)
```

---

## Debugging Checklist

- Does `docker build` fail on `go build`? Check that all source files are copied before the build step and that `go.mod` is present.
- Does the server start but the page shows no CSS or templates? The banner files, templates, and static files must be copied into the final image — not just the binary.
- Does `docker run` fail with "port already in use"? Another process is using port 8080. Stop it or use a different host port: `-p 9090:8080`.
- Do you have dozens of dangling images? Run `docker image prune` to clean them up.
- Does the container exit immediately after starting? Check the logs with `docker logs <container_id>` to see the error.

---

## Key Commands

| Command | What It Does |
|---|---|
| `docker build -t name .` | Build an image from the Dockerfile in the current directory |
| `docker run -p 8080:8080 name` | Run a container, mapping host port to container port |
| `docker run -d -p 8080:8080 name` | Run detached (in background) |
| `docker ps` | List running containers |
| `docker ps -a` | List all containers including stopped |
| `docker logs <id>` | View container output |
| `docker stop <id>` | Stop a running container |
| `docker rm <id>` | Remove a stopped container |
| `docker images` | List all images |
| `docker rmi <id>` | Remove an image |
| `docker system prune` | Remove all stopped containers and dangling images |
| `docker inspect <name>` | Show detailed metadata about an image or container |

---

## Submission Checklist

- [ ] `docker build -t ascii-art-web .` completes without errors
- [ ] `docker run -p 8080:8080 ascii-art-web` starts the server
- [ ] Website is fully functional inside the container
- [ ] Dockerfile uses a multi-stage build
- [ ] Final image is based on a minimal base image
- [ ] Image has at least three labels (author, version, description)
- [ ] Labels are visible with `docker inspect`
- [ ] `.dockerignore` file exists and excludes unnecessary files
- [ ] Port is correctly exposed and mapped
- [ ] Container can be stopped cleanly with `docker stop`
- [ ] Stopped containers and unused images can be cleaned up with prune
- [ ] Dockerfile follows best practices (layer order, minimal image, non-root if possible)