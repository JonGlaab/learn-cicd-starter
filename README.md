# learn-cicd-starter (Notely)

https://github.com/JonGlaab/learn-cicd-starter/actions/workflows/ci.yml/badge.svg

![hot bunny action](https://img.magnific.com/premium-photo/white-rabbit-garden-fluffy-bunny-green-grass-summer-time_288539-2320.jpg)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!
Jon's version of Boot.dev's Notely app