Small project to start using golang

You can build and run the image to test it out

```sh
docker build -t <image_name> .
docker run -it --rm <image_name> -[sd] arg1 [args...]
```

Prints the memestring from the args, try it out with the default args of the image

At the request of someone else, if you supply the s flag to memestring then it will print in a retarded manner.

For ease of pasting to discord, the d flag will enclose the output in ``` for a markdown output

```sh
docker build -t <image_name> .
docker run -it --rm <image_name> -[sd] arg1 [args...]
```


For easier use you can avoid using the docker image and use xclip to directly copy to your clipboard (ubuntu)

```sh
docker run -it --rm <image_name> -[sd] [args...] | xclip -selection clipboard
```

Probably easier to make a function and add it to your rc file though (this will need you to enclose the args in "" or '' though)
```sh
func () {
	docker run -it --rm <image_name> s[sd] $1 | xclip -selection clipboard
}
```
