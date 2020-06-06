#!/usr/bin/env zsh
# Usage: ./build service.foo [--push|-p] [--deploy|-d] [--force|-f]
#   Each stage is broken up into its own function so another script can source it.

args=$*

# set global variables we'll need throughout.
NAMESPACE=default
GCP_PROJECT="lolibrary-180111"

# options; setting defaults.
OPT_DOCKER_PUSH=no
OPT_K8S_DEPLOY=no
OPT_FORCE=no
OPT_PRINT_OPTIONS=no

# args used for specific things.
BRANCH=""
GIT_SHA=""
DOCKER_SHA_TAG=""
DOCKER_BRANCH_TAG=""
VERBOSE=""

function setup/parse_options() {
    debug_print "parsing options"

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
            -v | --verbose)
                VERBOSE="true"
                ;;
            *)
                BRANCH=$param
                ;;
        esac
        shift
    done
}

function setup/configure() {
    if [[ "$BRANCH" = "" ]]; then
        setup/usage
    fi
}

# Print usage and exit.
function setup/usage() {
    debug_print "printing usage"

    print -P "🌸 Usage: %F{red}build.zsh%f [-d|--deploy] %F{green}branch%f"
    exit
}

cache=""
function setup/clean() {
    debug_print "tearing down previous cache"

    previous=$(pwd)
    cache="$(pwd)/cache"

    if [[ "$CLEAN_CACHE" = "true" ]]; then
      rm -rf $cache
    fi

    if [[ ! -d $cache ]]; then
      git clone git@github.com:lolibrary/sakura.git -q $cache
    fi

    cd $cache

    if [[ "$VERBOSE" = "true" ]]; then
      git fetch
      git checkout --force origin/$BRANCH
    else
      git fetch -q
      git checkout --force -q origin/$BRANCH
    fi

    print -P "🌳 %F{blue}$BRANCH%f"

    GIT_SHA="$(git rev-parse HEAD)"
    COMMIT_MESSAGE=$(git log -n 1 --format=%s)
    print -P "🐙 %F{green}$GIT_SHA%f"
    print -P "💬 $COMMIT_MESSAGE"

    DOCKER_SHA_TAG="gcr.io/$GCP_PROJECT/sakura:$GIT_SHA"
    DOCKER_BRANCH_TAG="gcr.io/$GCP_PROJECT/sakura:$BRANCH"

    cd $previous
}

function install/php() {
  print -P -n "    %F{red}composer%f    "

  if [[ "$VERBOSE" = "true" ]]; then
    composer install --no-interaction --working-dir $cache
  else
    composer install --no-interaction --working-dir $cache --quiet
  fi

  print "✅"
}

function install/node() {
  print -P -n "    %F{red}node.js%f     "

  if [[ "$VERBOSE" = "true" ]]; then
    npm install --prefix $cache
  else
    npm install --prefix $cache --silent &>/dev/null
  fi

  print "✅"
}

function install/mix() {
  print -P -n "    %F{red}laravel-mix%f "

  if [[ "$VERBOSE" = "true" ]]; then
    npm run production --prefix $cache
  else
    npm run production --prefix $cache --silent &>/dev/null
  fi

  print "✅"
}

function docker/build() {
    print ""
    print -P "🐳 %F{blue}gcr.io/$GCP_PROJECT%f/%F{red}sakura%f:%F{green}$GIT_SHA%f"
    print -P "   %F{blue}gcr.io/$GCP_PROJECT%f/%F{red}sakura%f:%F{green}$BRANCH%f"
    print ""
    print "👷‍♀️ Building docker images"
    print ""

    if [[ "$VERBOSE" = "true" ]]; then
      docker build -t $DOCKER_SHA_TAG .
      docker build -t $DOCKER_BRANCH_TAG .
    else
      docker build -t $DOCKER_SHA_TAG . > /dev/null 2>&1
      docker build -t $DOCKER_BRANCH_TAG . > /dev/null 2>&1
    fi
}

function docker/push() {
    print -P "🚢 Pushing docker images"

    if [[ "$VERBOSE" = "true" ]]; then
      docker push $DOCKER_SHA_TAG
      docker push $DOCKER_BRANCH_TAG
    else
      print -P "    %F{blue}gcr.io/$GCP_PROJECT%f/%F{red}sakura%f:%F{green}$GIT_SHA%f"
      docker push $DOCKER_SHA_TAG > /dev/null 2>&1

      print -P "    %F{blue}gcr.io/$GCP_PROJECT%f/%F{red}sakura%f:%F{green}$BRANCH%f"
      docker push $DOCKER_BRANCH_TAG > /dev/null 2>&1
    fi
}

function k8s/deploy() {
    local service="s-sakura"

    print -P "☸️  Updating %F{red}deployment/s-sakura%f to %F{cyan}$GIT_SHA%f"
    cat "$(pwd)/manifests/prod.yml" | sed "s#PLACEHOLDER#\"$DOCKER_SHA_TAG\"#" | kubectl apply --record -f -

    print -P "♻️  Waiting for %F{red}$service%f to rollout..."

    kubectl rollout status "deployment/$service" -n $NAMESPACE

    print -P "✅ Deployed %F{red}service.sakura%f successfully"
}

function debug_print() {
    if [[ "$VERBOSE" = "true" ]]; then
        print $1
    fi
}

setup/parse_options $*
setup/configure
setup/clean

print ""
print "📲 Installing dependencies"
install/php
install/node
install/mix

docker/build
docker/push

k8s/deploy
