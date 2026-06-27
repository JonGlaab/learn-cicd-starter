# learn-cicd-starter (Notely)



![passing badge lol](https://github.com/JonGlaab/learn-cicd-starter/actions/workflows/ci.yml/badge.svg)

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


name: cd

on:
    push:
        branches: [main]



          # Authenticate and run everything natively using the pre-installed gcloud CLI
          - name: Deploy to Artifact Registry
            env:
              GCP_ACCESS_TOKEN: ${{ secrets.GCP_ACCESS_TOKEN }}
            run: |
              # 1. Force gcloud to look at your lab project
              gcloud config set project project-3df150e7-5b58-40ea-a45
              
              # 2. Authenticate using your temporary token secret
              gcloud config set auth/credential_id "$GCP_ACCESS_TOKEN"
              
              # 3. Kick off the build
              gcloud builds submit --tag us-central1-docker.pkg.dev/project-3df150e7-5b58-40ea-a45/notely-ar-repo/notely-app:latest .