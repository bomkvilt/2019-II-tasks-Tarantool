build_storage() {
    docker build -t "tests-storage:latest" -f "./storage/docker/storage.dockerfile" .
}

build_tests() {
    docker build -t "tests-tests:latest"   -f "./tests/docker/tests.dockerfile" .
}

if   [ "$1" == "-b" ]; then build_storage
elif [ "$1" == "-t" ]; then
    build_storage
    build_tests
    helm/control.sh -i
    sleep 5 # crutch
    kubectl logs pod/tests --container=tests -f
elif [ "$1" == "-r" ]; then
    ./helm/control.sh -r
fi
