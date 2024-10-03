# Lesson 9: HTTP

## HTTP

HTTP is a protocol for fetching resources such as HTML documents. It is the foundation of any data exchange on the Web
and it is a client-server protocol, which means requests are initiated by the recipient, usually the Web browser. A
complete document is reconstructed from the different sub-documents fetched, for instance, text, layout description,
images, videos, scripts, and more.

![diagram](fetching_a_page.png)

_Reference_: [Overview](https://developer.mozilla.org/en-US/docs/Web/HTTP)

### HTTP connection management

![diagram](image.jpg)

_Reference_: [HTTP2 vs HTTP3](https://gcore.com/learning/what-is-http-3/), [HTTP2 compare HTTP3](https://kiwee.eu/blog/http-3-how-it-performs-compared-to-http-2/)

### HTTP Messages

#### Request

![diagram](http_request.png)

- An HTTP [method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods), usually a verb
  like [GET](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/GET), [POST](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/POST),
  or a noun like [OPTIONS](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/OPTIONS)
  or [HEAD](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD) that defines the operation the client wants
  to perform. Typically, a client wants to fetch a resource (using `GET`) or post the value of an HTML form (
  using `POST`), though more operations may be needed in other cases.
- The path of the resource to fetch; the URL of the resource stripped from elements that are obvious from the context,
  for example without the [protocol](https://developer.mozilla.org/en-US/docs/Glossary/Protocol) (`http://`), the
  [domain](https://developer.mozilla.org/en-US/docs/Glossary/Domain) (here, `developer.mozilla.org`), or the
  TCP [port](https://developer.mozilla.org/en-US/docs/Glossary/Port) (here, 80).
- The version of the HTTP protocol.
- Optional [headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers) that convey additional information for
  the servers.
- A body, for some methods like `POST`, similar to those in responses, which contain the resource sent.

#### Response

![diagram](http_response.png)

- The version of the HTTP protocol they follow.
- A [status code](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status), indicating if the request was successful or
  not, and why.
- A status message, a non-authoritative short description of the status code.
- HTTP [headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers), like those for requests.
- Optionally, a body containing the fetched resource.

_Reference_: [HTTP Messages](https://developer.mozilla.org/en-US/docs/Web/HTTP/Messages)
