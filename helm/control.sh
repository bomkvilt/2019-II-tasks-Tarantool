pushd $(dirname "${BASH_SOURCE[0]}")
PROJECT=tests
if   [ "$1" = "-i" ]; then  helm upgrade --install ${PROJECT} ${@:2} .
elif [ "$1" = "-d" ]; then  helm install --dry-run --debug --name ${PROJECT} ${@:2} .
elif [ "$1" = "-r" ]; then  helm delete ${PROJECT} --purge
fi 
