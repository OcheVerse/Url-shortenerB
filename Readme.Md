# Project Name URL lik shortener 

## Description

This project is a URL shortener built using Go. However, it cannot be shipped due to recent changes in Redis.

## Redis Changing to Closed Source

Redis, the popular open-source in-memory data structure store, has recently made the decision to change its licensing model and transition to a closed-source model. This means that the previously available open-source version of Redis is no longer freely available for commercial use.

As a result, this project, which relies on Redis as a key-value store for efficient URL redirection, can no longer be shipped as intended. The reliance on Redis as an integral part of the project's functionality makes it impossible to continue development and distribution without violating the new licensing terms.

## Alternatives

While this project cannot be shipped due to the Redis licensing changes, there are alternative options available for implementing a URL shortener. Some possible alternatives include:

- Using other open-source key-value stores such as Memcached or Couchbase.
- Implementing a custom key-value store within the project using Go's built-in data structures.
- Exploring other third-party services or libraries that provide URL shortening functionality without relying on Redis.

## Conclusion

In light of Redis changing to a closed-source model, this Go project cannot be shipped as originally planned. However, by exploring alternative solutions and adapting the project to use different key-value stores or approaches, it may still be possible to achieve the desired functionality of a URL shortener.

Please note that this README file serves as an explanation for the current state of the project and should be updated accordingly as new developments or alternatives arise.
