if [ "$1" == "-b" ]; then
    helm/control.sh -r; sleep 6
    docker build -t "tests-tests:latest" -f "./tests/docker/tests.dockerfile" .
    docker build -t "tests-storage:latest" -f "./storage/docker/storage.dockerfile" .
    helm/control.sh -i
elif [ "$1" == "-r" ]; then
    ./helm/control.sh -r
fi
