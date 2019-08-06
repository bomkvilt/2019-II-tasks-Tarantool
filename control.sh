if [ "$1" == "-b" ]; then
    docker build -t "tests-tests:latest" -f "./tests/docker/tests.dockerfile" .
else
    echo "unexpected flags"
    exit 1
fi
