echo "Removing bootstrap..."
rm -f bootstrap

echo "Removing zip........."
rm -f gilson-tracker-backend.zip

echo "Building bootstrap..."
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap *.go
if [ $? -ne 0 ]; then
	echo "Bad"
	exit 1 
fi

echo "Building zip........."
zip gilson-tracker-backend.zip bootstrap
if [ $? -ne 0 ]; then
	echo "Bad"
	exit 2 
fi
