# Memegen

A webservice for generating memes by combining cat pictures from Cat As A Service (cataas.com) and programming jokes from JokeAPI (https://jokeapi.dev) as captions. Both are open APIs that don't require any keys.

This is a prototype as the implementation is very naive at times, unfinished and only partially tested.

## Overview

General flow of the application is as follows:
0. Determine the type of environment (prod or dev) and load configuration accordingly (e.g. API urls).
1. Fetch a joke from the JokeAPI, in case of errors, use a default caption.
2. Translate the joke into lolspeak using a simple dictionary preloaded from a static JSON file
3. Fetch a cat picture from `cataas.com/cat `, in case of errors, use a default image.
4. Superimpose the translated caption on the cat picture in a sacred meme format with a traditional meme font.
5. Render a simple HTML template showing the resulting image.

The application can be started with a flag in "production" mode (by no means production ready, only for demo purposes) or development mode. In production mode it will fetch images and jokes from the live APIs (with a fallback in case one of the APIs is down). In development it relies on a setup of Robohydra web server (http://robohydra.org). Robohydra is a powerful HTTP client test tool but here it's basically used to mock both APIs in order not to use live services when writing code. It can be used for manual testing too (included is the setup to test a case when cataas.com is timing out).

## Running the application without Docker

### For "production"/demo purposes
`go run main.go --env=prod`

### For development
1. Install Nodejs
2. Install and run Robohydra to mock APIs
```
cd robohydra
npm install robohydra
npx robohydra test.conf -I plugins
```

Admin for Robohydra will now be at http://localhost:3000/robohydra-admin.

Cat pics and jokes mock APIs are at http://localhost:3000/cat and http://localhost:3000/joke.

3. Run the application

`go run main.go --env=dev`

In the browser go to http://localhost:8080 and enjoy the memes.

## Running the application with Docker

Both main application and Robohydra are dockerized but the setup is extremely
simplified so currently it's not possible to use them both together.

### For "production"/demo purposes
1. Install and set up Docker
2. Build and run meme-server
```
sudo docker build -t meme-server .
sudo docker run -it --rm -p 8080:8080 meme-server --env=prod
```
In the browser go to http://localhost:8080 and enjoy the memes.

### For development
1. Install and set up Docker
2. Build and run Robohydra
```
cd robohydra
sudo docker build -t robohydra .
sudo docker run --name=robohydra -it -d -p 3000:3000 robohydra
```

Admin for Robohydra will now be at http://localhost:3000/robohydra-admin.
Cat pics and jokes mock APIs are at http://localhost:3000/cat and http://localhost:3000/joke.

3. In the root of the project `go run main.go`
If Nodejs is installed, one option is to install nodemon and auto-reload on changes:
```
sudo npm install -g nodemon
nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go
```
There are probably better ways to do this.

4. Kill the Robohydra container when you're done
`sudo docker container kill robohydra`

## Running tests

Run all tests:

`go test ./...`

or with Docker:

1. Build Docker image
`sudo docker build -t meme-server .`
2. Run it in daemon mode
`sudo docker run --name meme-server -it -d -p 8080:8080 meme-server`
3. Run all tests
`sudo docker exec meme-server go test ./...`
4. Kill the container when you're done
`sudo docker container kill meme-server`

## Improvements list

Since this is just a prototype, there is a lot that can be improved:
1. Better dictionary
Currently the dictionary is loaded from a static JSON file that comes from https://github.com/normansimonr/Dumb-Cogs/blob/master/lolz/data/tranzlashun.json - with a few additions. It would be better to have it stored in some sort of database instead. It would also be possible with just a few code modifications to use and siwtch between multiple dictionaries.
2. Smarter placement of the caption in the image
Right now sometimes the caption covers important parts of the image or is weirdly placed depending on the size of the image. Instead of harcoding, calculate the position of the text above and below at least based on the image size. Possibly modify the font size based on image size. Perhaps detect where the heads of cats are in the picture (snout and eyes detection :D).
3. Instead of hardcoded timeouts on both APIS these could be options in the configuration files or flags. There are TODOs in the code to add tests for them.
4. There is a TODO in the code to utilize the joke type that is included in the JokeAPI response to determine if it's a two-part joke (setup and delivery) that should be displayed at the top and the bottom of the image or if it's a single joke that can just be displayed in the bottom.
5. Add more tests to each package.

## Ideas for new features
1. Add a counter for memes viewed.
2. Add a history of viewed memes (can show 5 last memes at /history). As a very simple, naive solution it could dump the images in base64 to e.g. Redis. Add Redis image in Docker file, add connection config option, in the handler save the image every time with timestamped `viewed`, make a new template and display 5 last sorted by the timestamp.
3. Custom meme page - a simple form where you can input top and bottom text, upload image and submit, then load the image with the caption applied.
4. Add `docker-compose` to build and run both the app and Robohydra at the same time, set up network so that the app run in a container sees Robohydra from the other container.