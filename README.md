![Toggle](toggle_logo.png)

# The Vision

A startup company has a vision for a service that provides a tailored news feed to its users.

News will be acquired from a multitude of online sources. Each news article will be analyzed (using AI) and tagged with one or more keywords, 
before being stored in a database within the company's infrastructure. Users of the service are able to specify tags for news areas that interest them. When
a user opens their dashboard, they should see a feed of the most recent and relevant news articles.

# The Mission

There are few architectural guidelines so it’s completely up to you; however, to support their growth, the company have opted for
the microservice route. Obviously, a number of services are going to be 
required, but don't panic, we're not expecting you to implement _everything_ (although that would be impressive)!

Your task is to implement two services as follows:

- A `user` service that will provide a user's chosen tags
- A `news article` service that will provide a feed of news articles

## 1. User Service

The `User` service is a microservice that stores each user's tag selection.

### Internal Endpoint

_Protocol, method, signature, etc. - (TBD, by candidate!)_

**Role:**

This endpoint is called by the `News Article` service.

**Behaviour:**

For a given user, return all the tags specified.


## 2. News Article Service

The `News Article` service is a microservice that stores news articles.

Each article must contain a `Title`, `Timestamp` and list of `Tags`.

This service also provides an external endpoint that allows users to retrieve articles, filtered and sorted by their timestamp.

### Public Endpoint

_Protocol, method, signature, etc. - (TBD, by candidate!)_

**Role:**

This endpoint is called from user's browser/phone/tablet.

**Predicate**

> An article is included in the response if it has at least 1 tag matching those of the user.

**Behaviour:**

Return news articles, filtered by the tags of the current user.

### Public Endpoint

_Protocol, method, signature, etc. - (TBD, by candidate!)_

Request must include 2 tags

**Role:**

This endpoint is called from user's browser/phone/tablet.

**Predicate**

> An article is included in the response if its tags contain _both_ of the tags specified in the request.

The product team mentioned that they might need to change these predicate values. Modifying both the number 
of tags in the request and that used in the matching criteria. So, bonus points if you make these configurable! ;)


**Behaviour:**

Return news articles, filtered by the tags included in the request.


## Prerequisites
- Handle all failure cases
- Your code should be tested
- Provide a `docker-compose.yaml` file for any third party services that you use 
- Provide a clear explanation of your approach and design choices (while submitting your pull request)
- Provide a proper `README.md`:
    - Explain how to setup and run your code
    - Include all information you consider useful for a seamless coworker on-boarding

## Workflow
- Create a new branch
- Commit and push to this branch
- Submit a pull request once you have finished

We will then write a review for your pull request!

## Bonus

- Add metrics / request tracing / authentication 📈
- Add whatever you think is necessary to make the app awesome ✨
