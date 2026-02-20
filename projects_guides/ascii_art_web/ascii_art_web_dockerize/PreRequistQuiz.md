# 🎯 ASCII-Art-Web-Dockerize Prerequisites Quiz
## Containers · Docker Images · Dockerfile · Multi-Stage Builds · Port Mapping · Garbage Collection

**Time Limit:** 55 minutes  
**Total Questions:** 30  
**Passing Score:** 24/30 (80%)

> ✅ Pass → You're ready to start ASCII-Art-Web-Dockerize  
> ⚠️ This project introduces completely new concepts. A higher score (27+) means you'll work faster.  
> ❌ Score below 20 → Read the Docker getting-started docs fully before attempting this quiz again.

---

## 📋 SECTION 1: CONTAINER CONCEPTS (6 Questions)

### Q1: What is the key difference between a Docker container and a virtual machine?

**A)** Containers are faster to download  
**B)** A VM includes a full OS kernel; a container shares the host OS kernel and only packages the application and its dependencies — making containers much lighter  
**C)** Containers only run on Linux  
**D)** VMs use more RAM by coincidence, not by design  

<details><summary>💡 Answer</summary>

**B) A VM includes a full OS kernel; a container shares the host OS kernel**

| | VM | Container |
|---|---|---|
| Size | GBs (full OS) | MBs (app + deps only) |
| Startup | Minutes | Seconds or less |
| Isolation | Full hardware virtualization | Process-level isolation |
| Kernel | Each VM has its own | Shared with host |

Containers are the right tool for packaging and shipping a web server like this project's.

</details>

---

### Q2: What is a Docker image?

**A)** A running instance of your application  
**B)** A screenshot of your application's UI  
**C)** A read-only, layered snapshot of a filesystem — a blueprint for running containers  
**D)** A Git repository  

<details><summary>💡 Answer</summary>

**C) A read-only, layered snapshot — a blueprint for running containers**

An image is built once from a Dockerfile and stored. You can run many containers from the same image. Think of it as: image = class, container = instance.

</details>

---

### Q3: What is a Docker container?

**A)** The same as an image  
**B)** A running instance of an image — a process isolated from the host with its own filesystem  
**C)** A Dockerfile  
**D)** A directory of source code  

<details><summary>💡 Answer</summary>

**B) A running instance of an image**

`docker run` takes an image and starts a container from it. The container is isolated: its own filesystem (based on the image), its own network namespace, its own process tree. Stopping the container does NOT destroy the image.

</details>

---

### Q4: What is the difference between `docker build` and `docker run`?

**A)** They are the same command  
**B)** `docker build` creates an image from a Dockerfile; `docker run` creates and starts a container from an image  
**C)** `docker run` creates an image; `docker build` runs it  
**D)** `docker build` compiles Go code; `docker run` runs it  

<details><summary>💡 Answer</summary>

**B) `docker build` creates an image; `docker run` starts a container from an image**

```bash
docker build -t ascii-art-web .   # creates image named "ascii-art-web"
docker run -p 8080:8080 ascii-art-web  # starts a container from that image
```

Build once, run many times. The image is portable — it runs the same on any machine with Docker.

</details>

---

### Q5: You built your application image on your Mac. Your teammate has a Linux server. Can they run your image?

**A)** No — Docker images are OS-specific  
**B)** Yes — if both run the same CPU architecture (e.g. both x86_64), Docker images are portable across operating systems  
**C)** Only if they install the same version of Go  
**D)** Only if they have the same username  

<details><summary>💡 Answer</summary>

**B) Yes — Docker images are portable across OS (same CPU architecture)**

This is one of Docker's main selling points: "build once, run anywhere." The container brings everything it needs. The host OS doesn't need Go, the correct library versions, or anything else your app requires. The caveat is CPU architecture: an x86_64 image won't run on ARM without emulation.

</details>

---

### Q6: Your Go server reads `standard.txt` at runtime. If that file is not inside the Docker image, what happens when the container runs?

**A)** Docker automatically finds it on the host  
**B)** The server fails to load the banner file and crashes or returns errors  
**C)** Docker mounts the host filesystem automatically  
**D)** The container downloads it from the internet  

<details><summary>💡 Answer</summary>

**B) The server fails — the container's filesystem only contains what you put in the image**

A container is isolated. Its filesystem comes entirely from the image layers. If you forgot to `COPY` the banner files, templates, or static assets into the image, the server can't find them at runtime. This is one of the most common first-time Docker bugs.

</details>

---

## 📋 SECTION 2: DOCKERFILE SYNTAX (8 Questions)

### Q7: What does the `FROM` instruction do?

**A)** Specifies where to copy files from  
**B)** Sets the base image for your Docker image — every subsequent instruction runs on top of this  
**C)** Specifies the source code location  
**D)** Sets the FROM header in HTTP responses  

<details><summary>💡 Answer</summary>

**B) Sets the base image — every Dockerfile must start with `FROM`**

```dockerfile
FROM golang:1.22-alpine  # start with the official Go Alpine image
```

The base image provides the OS environment. Everything you `RUN`, `COPY`, and configure is added on top of it as layers.

</details>

---

### Q8: What is the difference between `COPY` and `ADD` in a Dockerfile?

**A)** `ADD` is the newer replacement for `COPY`  
**B)** `COPY` copies local files into the image; `ADD` does the same but also handles URLs and auto-extracts tar archives — use `COPY` unless you specifically need `ADD`'s extra features  
**C)** `COPY` is for directories; `ADD` is for files  
**D)** They are identical  

<details><summary>💡 Answer</summary>

**B) `COPY` for simple file copying; `ADD` has extra features but is less predictable**

Docker best practices say: use `COPY` unless you explicitly need URL downloading or tar extraction. `COPY` is explicit and its behavior is always predictable.

</details>

---

### Q9: What does `WORKDIR /app` do?

**A)** Creates an environment variable named `WORKDIR`  
**B)** Sets the working directory for all subsequent instructions (`RUN`, `CMD`, `COPY` destination, etc.)  
**C)** Copies files to `/app`  
**D)** Starts the server from the `/app` directory  

<details><summary>💡 Answer</summary>

**B) Sets the working directory for subsequent instructions**

```dockerfile
WORKDIR /app
COPY . .        # copies to /app/
RUN go build .  # runs in /app/
```

If the directory doesn't exist, `WORKDIR` creates it. Using it instead of `RUN mkdir /app && cd /app` is the idiomatic approach.

</details>

---

### Q10: What is the difference between `RUN`, `CMD`, and `ENTRYPOINT`?

**A)** They are all identical  
**B)** `RUN` executes during image build; `CMD` sets the default command when a container starts; `ENTRYPOINT` sets a fixed executable that `CMD` arguments are passed to  
**C)** `RUN` starts the server; `CMD` builds the code; `ENTRYPOINT` sets environment variables  
**D)** `RUN` is for Linux; `CMD` is for Windows; `ENTRYPOINT` is cross-platform  

<details><summary>💡 Answer</summary>

**B) `RUN` = build time; `CMD` = container start default; `ENTRYPOINT` = fixed executable**

```dockerfile
RUN go build -o server .       # runs during build, creates the binary
CMD ["/app/server"]            # default command when container starts
# or
ENTRYPOINT ["/app/server"]     # always runs this; CMD would add args to it
```

For a simple Go server, `CMD ["/app/server"]` is the standard choice.

</details>

---

### Q11: What does `EXPOSE 8080` do?

**A)** Makes the container accessible on port 8080 automatically  
**B)** Documents which port the container uses — it does NOT automatically publish the port to the host  
**C)** Blocks all ports except 8080  
**D)** Opens port 8080 in the host firewall  

<details><summary>💡 Answer</summary>

**B) `EXPOSE` documents the port — it does NOT publish it**

`EXPOSE` is documentation and metadata. The port is only accessible from the host when you use `-p` in `docker run`:
```bash
docker run -p 8080:8080 ascii-art-web   # host:container port mapping
```

Without `-p 8080:8080`, even if you have `EXPOSE 8080`, the port is not accessible from your browser.

</details>

---

### Q12: Why should you copy `go.mod` and `go.sum` BEFORE copying the rest of the source code?

```dockerfile
# Option A:
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build .

# Option B:
COPY . .
RUN go mod download
RUN go build .
```

**A)** Option B is better — fewer COPY instructions  
**B)** Option A is better — Docker caches each layer. If only source code changed (not dependencies), the `go mod download` layer is reused from cache, making rebuilds much faster  
**C)** They are identical  
**D)** Option A is slower  

<details><summary>💡 Answer</summary>

**B) Option A uses Docker layer caching to skip `go mod download` when only source changed**

Docker rebuilds layers starting from the first changed instruction. If `go.mod` and `go.sum` don't change, the `go mod download` layer is cached. Downloading dependencies is slow — caching it saves significant time on every build where you only changed your Go code.

</details>

---

### Q13: What is a multi-stage Docker build?

**A)** Building the image in multiple terminals simultaneously  
**B)** Using multiple `FROM` instructions in one Dockerfile — the first stage builds the app, the second stage creates a minimal runtime image using only the compiled binary  
**C)** Building the image multiple times and keeping the best one  
**D)** A Docker feature for Windows only  

<details><summary>💡 Answer</summary>

**B) Multiple `FROM` instructions — build stage creates binary, final stage is minimal**

```dockerfile
# Stage 1: Build
FROM golang:1.22-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/server .   # copy only the binary
COPY standard.txt shadow.txt thinkertoy.txt ./
COPY templates/ templates/
CMD ["./server"]
```

The final image contains only alpine + binary + data files. No Go toolchain, no source code. Result: ~10MB instead of ~800MB.

</details>

---

### Q14: What base image should you use for the **final stage** of a Go multi-stage build?

**A)** `golang:latest` — same as the build stage  
**B)** `alpine:latest` or `scratch` — minimal base images for small final images  
**C)** `ubuntu:latest` — a full OS for stability  
**D)** `node:latest` — for serving web content  

<details><summary>💡 Answer</summary>

**B) `alpine:latest` or `scratch`**

- `alpine` is a minimal Linux (~5MB) — has shell, basic utilities, and libc
- `scratch` is empty (0 bytes) — requires a fully static binary (use `CGO_ENABLED=0`)

For a Go web server:
```dockerfile
FROM alpine:latest     # easiest choice
# or
FROM scratch          # requires: RUN go build with CGO_ENABLED=0
```

Never use `golang:latest` in the final stage — it's ~800MB and includes the entire Go toolchain you don't need at runtime.

</details>

---

## 📋 SECTION 3: DOCKER CLI & LABELS (6 Questions)

### Q15: What does `docker build -t ascii-art-web .` do?

**A)** Runs the container named `ascii-art-web`  
**B)** Builds an image named `ascii-art-web` using the `Dockerfile` in the current directory (`.`)  
**C)** Pulls the `ascii-art-web` image from Docker Hub  
**D)** Tags an existing image  

<details><summary>💡 Answer</summary>

**B) Builds an image named `ascii-art-web` from the current directory's Dockerfile**

- `-t ascii-art-web` — tag (name) for the image
- `.` — the build context (directory Docker sends to the daemon)

After this, `docker images` will show `ascii-art-web` in the list.

</details>

---

### Q16: What does `-p 8080:8080` mean in `docker run -p 8080:8080 ascii-art-web`?

**A)** Build on port 8080 and run on port 8080  
**B)** Map host port 8080 to container port 8080 — requests to `localhost:8080` are forwarded into the container  
**C)** Use port 8080 twice for redundancy  
**D)** Expose port 8080 externally and internally  

<details><summary>💡 Answer</summary>

**B) Map host port 8080 → container port 8080**

Format: `-p HOST_PORT:CONTAINER_PORT`. Your Go server inside the container listens on `:8080`. The `-p 8080:8080` flag makes the host's port 8080 forward to the container's port 8080. You could use different ports: `-p 9090:8080` would make `localhost:9090` reach the container's server.

</details>

---

### Q17: How do you run a container in detached mode (background)?

**A)** `docker run --background ascii-art-web`  
**B)** `docker run -d ascii-art-web`  
**C)** `docker run -b ascii-art-web`  
**D)** `docker run & ascii-art-web`  

<details><summary>💡 Answer</summary>

**B) `docker run -d ascii-art-web`**

`-d` (detached) runs the container in the background and prints the container ID. Without `-d`, the container's output goes to your terminal and Ctrl+C stops it. With `-d`, use `docker logs <id>` to see output and `docker stop <id>` to stop it.

</details>

---

### Q18: What does `docker ps -a` show vs `docker ps`?

**A)** Identical output  
**B)** `docker ps` shows only running containers; `docker ps -a` shows all containers including stopped ones  
**C)** `docker ps -a` shows images; `docker ps` shows containers  
**D)** `docker ps -a` shows more detailed stats  

<details><summary>💡 Answer</summary>

**B) `docker ps` = running only; `docker ps -a` = all including stopped**

Stopped containers still exist on disk until you `docker rm` them. They consume disk space. `docker ps -a` is how you see (and clean up) containers that have exited.

</details>

---

### Q19: How do you add a label to your Docker image?

**A)** In the Dockerfile: `LABEL maintainer="your@email.com"`  
**B)** As a command-line flag: `docker build --label maintainer="email" .`  
**C)** In `go.mod`  
**D)** Both A and B work  

<details><summary>💡 Answer</summary>

**D) Both A (Dockerfile LABEL) and B (`--label` flag) work**

Dockerfile `LABEL` is preferred for permanent metadata about the image (author, version, description) because it's version-controlled with the code:
```dockerfile
LABEL maintainer="your@email.com" \
      version="1.0" \
      description="ASCII Art Web Application"
```

The `--label` flag at build time is useful for dynamic labels like build timestamps.

</details>

---

### Q20: How do you verify that labels were applied to your image?

**A)** `docker labels ascii-art-web`  
**B)** `docker inspect ascii-art-web`  
**C)** `docker ps ascii-art-web`  
**D)** `cat Dockerfile`  

<details><summary>💡 Answer</summary>

**B) `docker inspect ascii-art-web`**

`docker inspect` outputs full JSON metadata about an image or container, including the `Labels` object. Pipe through grep to filter:
```bash
docker inspect ascii-art-web | grep -A 10 '"Labels"'
```

</details>

---

## 📋 SECTION 4: GARBAGE COLLECTION & CLEANUP (5 Questions)

### Q21: What is a "dangling" Docker image?

**A)** An image that failed to build  
**B)** An image with no tag (shown as `<none>`) — typically an intermediate layer from a previous build that got superseded by a new build with the same tag  
**C)** An image that is currently running  
**D)** An image stored on Docker Hub  

<details><summary>💡 Answer</summary>

**B) An untagged image — a superseded intermediate build layer**

Every time you run `docker build -t ascii-art-web .`, the old layers that were previously tagged as `ascii-art-web` lose their tag. They become dangling images: taking up disk space, listed as `<none>:<none>` in `docker images -a`.

</details>

---

### Q22: What does `docker system prune` do?

**A)** Removes all images, containers, and volumes — use with extreme caution  
**B)** Removes all stopped containers, all dangling images, and unused networks  
**C)** Removes only stopped containers  
**D)** Removes only dangling images  

<details><summary>💡 Answer</summary>

**B) Removes stopped containers, dangling images, and unused networks**

`docker system prune` is safe to run routinely — it doesn't remove running containers or tagged images. For more aggressive cleanup:
```bash
docker system prune -a   # also removes unused (not just dangling) images
docker system prune -a --volumes  # also removes unused volumes
```
Always know what you're deleting — run `docker system df` first to see disk usage.

</details>

---

### Q23: After rebuilding your image 10 times, you run `docker images -a` and see many `<none>` images. They take up 2GB. What command cleans them up?

**A)** `docker rm -f $(docker ps -a -q)`  
**B)** `docker image prune`  
**C)** `docker rmi $(docker images -q)`  
**D)** `docker system reset`  

<details><summary>💡 Answer</summary>

**B) `docker image prune`**

`docker image prune` removes all dangling (untagged) images. Option C would remove ALL images including your current tagged ones — dangerous. Option A removes containers, not images.

</details>

---

### Q24: What is a `.dockerignore` file and why do you need one?

**A)** It tells Docker which files to delete after building  
**B)** It tells Docker which files to exclude from the build context — preventing large, unnecessary files from being sent to the Docker daemon  
**C)** It specifies which images to ignore when pulling  
**D)** It's a configuration file for Docker Desktop  

<details><summary>💡 Answer</summary>

**B) It excludes files from the build context**

Without `.dockerignore`, `docker build .` sends your ENTIRE directory (including `.git`, test files, local binaries, etc.) to the Docker daemon. This wastes time and can accidentally include secrets or large files in your image.

```
# .dockerignore
.git
*.test
*.md
.env
```

</details>

---

### Q25: What should you typically include in `.dockerignore`?

**A)** All Go source files  
**B)** The Dockerfile itself  
**C)** `.git`, compiled binaries, test files, local `.env` files, editor directories  
**D)** The `go.mod` file  

<details><summary>💡 Answer</summary>

**C) `.git`, compiled binaries, test files, `.env`, editor directories**

```
# .dockerignore example
.git
.gitignore
*.test
*_test.go
*.md
.env
.DS_Store
.vscode/
.idea/
```

Never exclude `go.mod`, `go.sum`, or your source files — the build needs them. Never exclude banner files, templates, or static assets — the runtime needs them.

</details>

---

## 📋 SECTION 5: BEST PRACTICES & TRICKY CASES (5 Questions)

### Q26: Your Go binary runs as the `root` user inside the container. Why is this a security concern?

**A)** It's not — containers are always isolated  
**B)** If an attacker exploits your app, they get root privileges inside the container, which increases the blast radius of the compromise  
**C)** Root users can't run web servers  
**D)** Docker prevents root from binding to ports  

<details><summary>💡 Answer</summary>

**B) Root inside the container means full container control if exploited**

Best practice: create a non-root user in your Dockerfile:
```dockerfile
RUN adduser -D -u 1000 appuser
USER appuser
```

Not always required for this project, but good to know. The spec may or may not require it.

</details>

---

### Q27: Your container starts but the website shows "no banner file" errors. You confirm the binary is in the image. What do you check next?

**A)** The port mapping  
**B)** Whether the banner files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`) and `templates/` directory were copied into the image  
**C)** Whether Docker is up to date  
**D)** Whether the Go version is correct  

<details><summary>💡 Answer</summary>

**B) Whether banner files and templates were copied into the image**

A Go binary alone is not enough. Your server reads files at runtime. Every file your server reads must be inside the container's filesystem. Check your Dockerfile — do you have `COPY` instructions for the banner files and the `templates/` directory?

```dockerfile
COPY standard.txt shadow.txt thinkertoy.txt ./
COPY templates/ templates/
COPY static/ static/
```

</details>

---

### Q28: What is the correct command to build your image AND run it, all in one step?

**A)** `docker buildrun .`  
**B)** There is no single command — you always `docker build` first, then `docker run`  
**C)** `docker compose up`  
**D)** `docker start .`  

<details><summary>💡 Answer</summary>

**B) There is no single command — build first, then run**

```bash
docker build -t ascii-art-web .
docker run -p 8080:8080 ascii-art-web
```

`docker compose up` can build and run, but Compose requires a `docker-compose.yml` file. For this project, the two-step approach is expected.

</details>

---

### Q29: You rebuild your image but `docker run` still shows old behavior. What is the most likely cause?

**A)** Docker cached the build  
**B)** You're running a container from the old image — you need to stop the old container, remove it, and run a new one  
**C)** The binary didn't compile  
**D)** The port is blocked  

<details><summary>💡 Answer</summary>

**B) You're running an old container — stop it, remove it, run fresh**

`docker run` creates a NEW container from the image. If you have an old container still running (or stopped), it's based on the old image. The new image doesn't affect running/stopped containers:
```bash
docker stop <old_container_id>
docker rm <old_container_id>
docker run -p 8080:8080 ascii-art-web   # new container from new image
```

</details>

---

### Q30: After stopping a container with `docker stop`, what commands are needed to fully clean it up?

**A)** Nothing — stopped containers are automatically removed  
**B)** `docker rm <container_id>` to remove the stopped container, and `docker rmi <image>` if you also want to remove the image  
**C)** `docker delete <container_id>`  
**D)** `docker stop` also removes it  

<details><summary>💡 Answer</summary>

**B) `docker rm` to remove stopped container; `docker rmi` for the image**

`docker stop` only stops the running process — the container still exists in a stopped state, consuming disk space. `docker rm` removes it. `docker rmi` is separate and removes the image (not the container). The lifecycle: `build` → `run` → `stop` → `rm` → `rmi`.

Use `docker run --rm` to auto-remove the container when it stops:
```bash
docker run --rm -p 8080:8080 ascii-art-web
```

</details>

---

## 📊 Score Interpretation

| Score | Result |
|---|---|
| 28–30 ✅ | **Exceptional.** Deep Docker understanding — start immediately. |
| 24–27 ✅ | **Ready.** Review missed questions (especially multi-stage builds) before starting. |
| 18–23 ⚠️ | **Study first.** Read the Docker getting-started guide and Dockerfile best practices fully. |
| Below 18 ❌ | **Not ready.** Run `docker run hello-world`, work through the official Docker tutorial, then retry. |

---

## 🔍 Review Map

| Questions Missed | Topic to Study |
|---|---|
| Q1–Q6 | Container vs VM, images vs containers, `docker build` vs `docker run`, filesystem isolation |
| Q7–Q14 | Dockerfile instructions (`FROM`, `COPY`, `RUN`, `CMD`, `EXPOSE`, `WORKDIR`), layer caching, multi-stage builds |
| Q15–Q20 | `docker build -t`, `-p` port mapping, `-d` detached, `docker ps`, `docker inspect`, labels |
| Q21–Q25 | Dangling images, `docker system prune`, `docker image prune`, `.dockerignore` |
| Q26–Q30 | Non-root users, runtime file requirements, stale containers, `docker rm` vs `docker rmi` |