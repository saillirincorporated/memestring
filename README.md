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


For easier use you can avoid using the docker image and use xclip to directly copy to your clipboard

```sh
docker run -it --rm <image_name> [-s] [args...] | xclip -selection clipboard
```

Probably easier to make a function and add it to your rc file though (this will need you to enclose the args in "" or '' though)
```sh
func () {
	docker run -it --rm <image_name> [-s] $1 | xclip -selection clipboard
}
```
