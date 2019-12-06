# LINK Network Version2

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)

[![codecov](https://codecov.io/gh/line/link/branch/develop/graph/badge.svg?token=JFFuUevpzJ)](https://codecov.io/gh/line/link)

This repository hosts `LINK`, alternative implementation of the LINK Network.

**Node**: Requires [Go 1.13+](https://golang.org/dl/)

**Warnings**: Initial development is in progress, but there has not yet been a stable.

# Quick Start
**Build Docker Image**
```
make build-docker                # build docker image
```

**Configure**
```
./.initialize.sh docker          # prepare keys, validators, initial state, etc.
```

**Run**
```
docker-compose up                # Run a Node and REST
```

**visit with your browser**
* Node: http://localhost:26657/
* REST: http://localhost:1317/swagger-ui/
