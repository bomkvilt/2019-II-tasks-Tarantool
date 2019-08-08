build_storage() {
    docker build -t "tests-storage" -f "./storage/docker/storage.dockerfile" .
}

build_tests() {
    docker build -t "tests-tests"   -f "./tests/docker/tests.dockerfile" .
}

if   [ "$1" == "-b" ]; then build_storage
elif [ "$1" == "-t" ]; then
    build_storage
    build_tests
    helm/tests/control.sh -i
    sleep 5 # crutch
    kubectl logs pod/tests --container=tests -f
elif [ "$1" == "-r" ]; then
    ./helm/tests/control.sh -r
elif [ "$1" == "-d" ]; then
    gcloud builds submit --config cloudbuild.yaml .
    ./helm/prod/control.sh -i
fi
