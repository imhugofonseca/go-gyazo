# Gyazo Server Go
Simple implementation of the [gyazo server](https://github.com/gyazo/Gyazo/blob/master/Server/upload.cgi) in golang that uploads screenshots to a bucket (AWS, DO Spaces, etc..) using [minio](https://github.com/minio/minio-go).

## Why?
Because why not.

##  Configuration
The configuration of the server is done via env variables these are:

`BKT_HOST`: Your bucket host provider ex: `fra1.digitaloceanspaces.com`

`BKT_ACCESS_KEY`: Your bucket access key

`BKT_SECRET_ACCESS_KEY`: Your bucket access secret

`BKT_SPACE_DOMAIN`: The domain where your files will be served from.

`BKT_NAME`: The name of the bucket

## Deployment
A docker image and kubernetes deployment sample are provided in `Dockerfile` and `deployment.yaml`.

Or you can use it directly from github docker registry 
`docker.pkg.github.com/imhugofonseca/go-gyazo/gyazo:latest`

## Gyazo configuration
After having the server up and running you must configure your gyazo app to use the server. You can do so by recompiling your own, check [gyazo repository](https://github.com/gyazo/Gyazo).

## Sample screenshot
Just for clot wheres an example screenshot

https://img.hugo.gg/1bGlaCrE1hzzgW94CjR3Qh0RMDh.png
![Example screenshot](https://img.hugo.gg/1bGlaCrE1hzzgW94CjR3Qh0RMDh.png)

