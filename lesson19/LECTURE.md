# Lesson 16: REST API & Pagination

A REST API (Representational State Transfer Application Programming Interface) is a way to enable communication between
client and server over the web. It follows the principles of the REST architectural style, which emphasize simplicity,
scalability, and stateless interactions.

## REST API

### Features

#### Statelessness

Each request from the client to the server must contain all the information needed to process the request. The server
does not store client state between requests.

#### Client-Server Separation

The client (e.g., a web or mobile app) and server (e.g., a backend system) operate independently. The client makes
requests, and the server provides resources.

#### Uniform Interface

REST APIs use consistent and standardized resource URIs, making them easy to understand and interact with.

#### Resource-Based

REST treats everything as a resource, which can be represented using URIs (Uniform Resource Identifiers). For example:

`/users` for a list of users
`/users/123` for a specific user

#### HTTP Methods

REST APIs leverage standard HTTP methods to perform operations:

- `GET`: Retrieve a resource
- `POST`: Create a resource
- `PUT`: Update a resource (or create if it doesn’t exist)
- `PATCH`: Partially update a resource
- `DELETE`: Remove a resource

#### Representation of Resources

Resources are represented in formats like JSON, XML, or others, with JSON being the most commonly used.

#### Stateless Communication

All API calls are independent; session state is not stored on the server.

### Advantages

Advantages of REST APIs

- **Simplicity**: Uses standard HTTP methods and status codes.
- **Scalability**: Statelessness makes it easy to scale server components.
- **Flexibility**: Can handle different types of clients (e.g., mobile, web).
- **Interoperability**: Use of standard protocols ensures wide compatibility.

## Pagination

Pagination is a technique used to divide a large dataset into smaller, more manageable chunks (pages). This improves
performance and user experience by limiting the amount of data sent in a single API response.

1. **Improved Performance**: Reduces the amount of data sent over the network, improving API speed.
2. **Better User Experience**: Prevents overwhelming the client with large amounts of data.
3. **Scalability**: Makes it easier to handle large datasets by retrieving data in parts.

### Types

#### Offset-Based Pagination

Data is retrieved based on a starting index (offset) and the number of items to retrieve (limit).

```shell
GET /users?offset=0&limit=10
```

#### Cursor-Based Pagination

Uses a pointer (cursor) to fetch the next set of items.

```shell
GET /users?cursor=abc123&limit=5
```

## Graceful Shutdown

A graceful shutdown is a controlled process of stopping a system, application, or service in a way that prevents data
loss, ensures ongoing tasks are completed, and allows for proper cleanup of resources. Key characteristics include:

- Stopping new incoming tasks while completing existing critical operations
- Releasing system resources like file handles, network connections, and memory
- Saving any unsaved data or application state
- Closing connections and transactions safely
- Logging shutdown information for diagnostics

In programming, this is often implemented using signal handlers or shutdown hooks that catch termination signals and
initiate a systematic shutdown sequence, preventing abrupt termination that could cause data corruption or system
instability.
