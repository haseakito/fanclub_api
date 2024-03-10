# Desginful - Fanclud SNS

## Overview

- [About](#About)
- [Features](#Features)
- [Technology](#Technology)
- [Development](#Development)

## About

Designful is a social networking service exclusively designed for creators and fan. We have ambition to replace major giants in this field such as [Pixiv](https://www.pixiv.net/) and [OnlyFans](https://onlyfans.com/). With this goal in mind, we try to offer more simple, easy and intuitive user interface on the client side, while super fast, less resource intensive, flexible, scalable and most importantly secure and reliable backend is implemented on the server side.

## Features

### Current

Designful aims to implement all of the features covered by the competitors and also add new features to help creators and fans engage more **simple, easy and fun way**. 

Here are the lists of features worth mentioning.

- Customizable billboard and categories

- Authentication and Authorization

- Image and video upload

- User follow

- Post like

- Place orders and checkout

- Notifications

- Webhooks handling events such as user signup

### Future scope

As we're a small group of developers, we have limited resource and capabilities. Therefore, we have decided to prioritize the most important features first. As a record for roadmap, We have gathered the features that we'are going to implement in the next months. 

- Schedule post publish date

This include scheduler such as cron or background processing such as [AWS Batch](https://aws.amazon.com/batch/).

- Chat and messaging

This include the real time messaging with [socket.io](https://socket.io/).

- Assets transformation

This include adaptive birate streaming, making clips from uploaded video, and offering more insightful analytics for creators. 

- Advanced searching and filtering

This include the full-text search powered with [Meilisearch](https://www.meilisearch.com/) or [Elasticsearch](https://www.elastic.co/elasticsearch).

- Multi channel notifications

This include email and push notification support on top of in app notifications.

- Flexible pricing

This include facilitating easy launch of discount and campaign.

## Technology

<p style="display: inline">
    <!-- Backend -->
    <img src="https://img.shields.io/badge/-go-00ADD8.svg?logo=go&style=for-the-badge&logoColor=white">
    <!-- Middleware -->
    <img src="https://img.shields.io/badge/-nginx-009639.svg?logo=nginx&style=for-the-badge">
    <img src="https://img.shields.io/badge/-mysql-4479A1.svg?logo=mysql&style=for-the-badge&logoColor=white">
    <img src="https://img.shields.io/badge/-redis-DC382D.svg?logo=redis&style=for-the-badge&logoColor=white">
    <img src="https://img.shields.io/badge/-elasticsearch-005571.svg?logo=elasticsearch&style=for-the-badge">
    <!-- Saas -->
    <img src="https://img.shields.io/badge/-stripe-008CDD.svg?logo=stripe&style=for-the-badge&logoColor=white">
    <img src="https://img.shields.io/badge/-clerk-6C47FF.svg?logo=clerk&style=for-the-badge">
    <img src="https://img.shields.io/badge/-sentry-362D59.svg?logo=sentry&style=for-the-badge">
    <img src="https://img.shields.io/badge/-novu-000000.svg?logo=novu&style=for-the-badge">    
    <!-- Infrastracture -->
    <img src="https://img.shields.io/badge/-Docker-1488C6.svg?logo=docker&style=for-the-badge">
    <img src="https://img.shields.io/badge/-terraform-20232A?logo=terraform&style=for-the-badge">
    <img src="https://img.shields.io/badge/-githubactions-2088FF.svg?logo=github-actions&style=for-the-badge&logoColor=white">
    <img src="https://img.shields.io/badge/-awsfargate-232F3E.svg?logo=awsfargate&style=for-the-badge">
    <img src="https://img.shields.io/badge/-awslambda-232F3E.svg?logo=awslambda&style=for-the-badge">
    <!-- Analytics -->
    <img src="https://img.shields.io/badge/-googleanalytics-FFFFFF.svg?logo=googleanalytics&style=for-the-badge">
</p>

### Backend

- [Echo](https://echo.labstack.com/): Web framework
- [Ent](https://entgo.io/): ORM

### Saas

- [Stripe](https://stripe.com/): Checkout
- [Clerk](https://clerk.com/): Authentication and Authorization
- [Cloudinary](https://cloudinary.com/): Assets upload, transformation, and optimization
- [Sentry](https://sentry.io/welcome/): Error monitoring
- [Novu](https://novu.co/): Notifications

### Infrastructure

- [Docker](https://www.docker.com/): Container
- [Terraform](https://www.terraform.io/): IaC
- [Github Actions](https://github.com/features/actions): CI/CD
- [AWS](https://aws.amazon.com/): Cloud Infrastruture

## Development

For local development, here are some basic commands with `Makefile` to get started.
Make sure that you follow [this instruction](https://ngrok.com/docs/integrations/clerk/webhooks/) to install `ngrok` to run `make listen` command.

```
# Expose local endpoint to receive webhook from Clerk
make listen

# Start container
make start

# Stop container
make stop

# Run unit tests
make test
```
