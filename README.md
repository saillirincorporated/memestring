Small project to start using golang

You can build and run the image to test it out

```sh
docker build -t <image_name> .
docker run -it --rm <image_name> [args...]
```

Prints the memestring from the args, try it out with the default args of the image


At the request of someone else, if you supply the s flag to memestring then it will print in a retarded manner

```sh
docker build -t <image_name> .
docker run -it --rm <image_name> [-s] [args...]
```

