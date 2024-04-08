# Brand Monitor

## Goal

The goal is to `simulate a search on Google Search`, where we have a simple front end making a call to an API. It's necessary to use Golang.

## Project Structure

The `backend folder` contains the HTTP API with one endpoint `/`. This endpoint expects a query parameter called `search_string`.

```html
http://localhost:8080/?search_string=Corinthians
```

The `front end folder` contains the `CSS` and `JS` for the `index.html` present in the project root folder.

## CI/CD

When new commits are pushed to the `main` branch, this will trigger GitHub Actions to `build the Docker image and upload it to my personal Docker Hub`. Every new image uploaded to the Docker Hub will `trigger a deployment in the Azure App Service` hosting this application.

## Result

To check the result and consume it, use the [deployed site link](https://romerogabriel.github.io/brandmonitor-test/).

## Warnings

This project use [Azure App Service free tier](https://azure.microsoft.com/en-us/pricing/details/app-service/windows/#pricing) to host the application. This can result in a [cold start problem](https://medium.com/@ilakk2023/overcoming-the-cold-start-problem-in-microservices-strategies-and-aws-solutions-2f93fc1e59a6).
