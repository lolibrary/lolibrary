#!/usr/bin/env zsh
# Usage: ./build service.foo [--push|-p] [--deploy|-d] [--force|-f]
#   Each stage is broken up into its own function so another script can source it.

args=$*

# set global variables we'll need throughout.
NAMESPACE=lolibrary
VERSION=$(git rev-parse HEAD)
PROJECT="github.com/lolibrary/lolibrary"
BASE="${GOPATH}/src/${PROJECT}"
GCP_PROJECT="lolibrary-180111"

# options; setting defaults.
OPT_DOCKER_PUSH=no
OPT_K8S_DEPLOY=no
OPT_FORCE=no
OPT_PRINT_OPTIONS=no

# args used for specific things.
SERVICE_DIR=""
BINARY_NAME=""
DOCKER_TAG=""
CACHE_DIR=""
SERVICE_NAME=""

function setup/parse_options() {
    while [ "$1" != "" ]; do
        local param=$(echo $1 | awk -F= '{print $1}')
        local value=$(echo $1 | awk -F= '{print $2}')

        case $param in
            -h | --help)
                setup/usage
                ;;
            -d | --deploy)
                OPT_K8S_DEPLOY=yes
                ;;
            -p | --push)
                OPT_DOCKER_PUSH=yes
                ;;
            -v | --print-options)
                OPT_PRINT_OPTIONS=yes
                ;;
            *)
                SERVICE_NAME=$param
                ;;
        esac
        shift
    done
}

# Configure options used to generate our builds.
function setup/configure() {
    if [[ "$SERVICE_NAME" = "" ]]; then
        setup/usage
    fi

    SERVICE_DIR="$BASE/$SERVICE_NAME"
    BINARY_NAME="org.lolibrary.$SERVICE_NAME"
    DOCKER_TAG="gcr.io/$GCP_PROJECT/$BINARY_NAME:$VERSION"
    CACHE_DIR="$BASE/.cache/$SERVICE_NAME"

    if [[ ! -d $SERVICE_DIR ]]; then
        print -P "⚠️  Directory %F{green}$SERVICE_DIR%f does not exist."
        exit 1
    fi
}

# Print usage and exit.
function setup/usage() {
    print -P "🌸 Usage: %F{red}build%f [-p|--push] [-d|--deploy] %F{green}service.foo%f"
    exit
}

function setup/print_options() {
    print -P "%F{red}OPT_DOCKER_PUSH%f: %F{blue}$OPT_DOCKER_PUSH%f"
    print -P "%F{red}OPT_K8S_DEPLOY %f: %F{blue}$OPT_K8S_DEPLOY %f"
    print -P "%F{red}SERVICE_NAME   %f: %F{blue}$SERVICE_NAME   %f"
    print -P "%F{red}SERVICE_DIR    %f: %F{blue}$SERVICE_DIR    %f"
    print -P "%F{red}BINARY_NAME    %f: %F{blue}$BINARY_NAME    %f"
    print -P "%F{red}DOCKER_TAG     %f: %F{blue}$DOCKER_TAG     %f"
    print -P "%F{red}CACHE_DIR      %f: %F{blue}$CACHE_DIR      %f"
}

function build/create_cache() {
    mkdir -p $CACHE_DIR
}

function build/go_build() {
    print -P "👩‍🔧 Building %F{green}$BINARY_NAME%f with %F{red}go build%f"

    CGO_ENABLED=0 GOOS=linux ARCH=amd64 \
        go build -i -installsuffix docker \
            -ldflags "-X \"$PROJECT/foundation.Version=$VERSION\" -X \"$PROJECT/libraries/filters.Version=$VERSION\"" \
            -o "$CACHE_DIR/$BINARY_NAME" "$SERVICE_DIR/main.go"

    print -P "✅ Binary %F{green}$BINARY_NAME%f built"
}

function docker/dockerfile() {
    local dockerfile="
FROM scratch
LABEL maintainer=engineering@lolibrary.org
EXPOSE 8080
ADD $BINARY_NAME /
ENTRYPOINT [\"/$BINARY_NAME\"]
"

    local dockerignore="
**/*
!$BINARY_NAME
"

    echo "$dockerignore" > "$CACHE_DIR/.dockerignore"
    echo "$dockerfile" > "$CACHE_DIR/Dockerfile"
}

function docker/build() {
    print -P "🐳 Building docker image %F{green}$DOCKER_TAG%f"

    docker/dockerfile
    docker build -t $DOCKER_TAG $CACHE_DIR > /dev/null 2>&1
}

function docker/push() {
    print -P "🚢 Pushing docker image to %F{green}$DOCKER_TAG%f"

    docker push $DOCKER_TAG
}

function k8s/deploy() {
    local service=$(echo $SERVICE_NAME | sed "s/\./-/g" | sed "s/^service/s/")

    print -P "⚙️  Updating %F{red}$service%f to %F{cyan}$DOCKER_TAG%f in k8s"

    cat "$SERVICE_DIR/manifest/prod.yml" | sed "s#PLACEHOLDER#\"$DOCKER_TAG\"#" | kubectl apply --record -f -

    print -P "♻️  Waiting for %F{red}$service%f to rollout..."

    kubectl rollout status "deployment/$service" -n $NAMESPACE

    print -P "✅ Deployed %F{red}$SERVICE_NAME%f successfully"
}

function debug_print() {
    if [[ "$VERBOSE" = "true" ]]; then
        print $1
    fi
}

debug_print "entering options parsing"
setup/parse_options $*

debug_print "setting up config"
setup/configure
if [[ $OPT_PRINT_OPTIONS = "yes" ]]; then
    setup/print_options
fi

debug_print "build starting"
# now we start to build the program.
build/create_cache
build/go_build

# now we have a build; deploy to docker.
debug_print "docker building"
docker/build

if [[ $OPT_DOCKER_PUSH = "yes" ]]; then
    docker/push
fi

debug_print "ready to deploy, checking."
# deploy to k8s with a confirmation.
if [[ $OPT_K8S_DEPLOY = "yes" ]]; then
    k8s/deploy
fi
