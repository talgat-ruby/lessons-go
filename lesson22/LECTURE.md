# Lesson 22: Intro into Authentication & Authorization

**Authentication** is the process of verifying the identity of a user, system, or entity. It answers the question, "_Are
you who you claim to be?_" Authentication typically involves proving one's identity through one or more of these
methods:

- Something you know (password, PIN)
- Something you have (security token, mobile device)
- Something you are (biometric data like fingerprint, facial recognition)

**Authorization** is the process of determining what specific resources, systems, or actions a verified user is allowed
to access or perform. It answers the question, "_What are you permitted to do?_" Authorization happens after successful
authentication and defines the user's access rights and permissions.

## Authentication between Services

### Cookie-Based Authentication

- **User authentication**: When a user logs in, the server generates a unique session identifier and stores it in a
  cookie.
- **Cookie transmission**: The cookie is sent to the user’s browser, which stores it locally.
- **Subsequent requests**: On subsequent requests, the browser automatically includes the cookie in the request headers.
- **Server verification**: The server verifies the cookie and checks if it matches the expected session identifier.
- **Authentication success**: If the cookie is valid, the server grants access to the requested resource.

#### Pros

- **Simple implementation**: Cookie-based authentication is easy to implement, requiring minimal effort on both the
  server and client sides.
- **Session management**: Cookies can store additional information related to the user’s session, such as preferences or
  shopping cart contents.
- **Browser support**: Works well with browsers.

#### Cons

- **Security risks**: Cookies are vulnerable to attacks like cross-site scripting (XSS) and cross-site request forgery (
  CSRF), allowing attackers to gain unauthorized access.
- **Less secure for mobile and single-page applications**: Cookies are stored locally on the user’s device, making them
  susceptible to theft or exploitation
  if an attacker gains access to the device.
- **Limited scope**: Cookies are limited to the domain that set them, making it difficult to share authentication
  information between sub-domains or domains.

### Token-Based Authentication (JWT - JSON Web Tokens)

- **User Authentication**: The user provides their credentials (e.g., username and password).
- **Token Generation**: The server generates a JWT token containing the user’s claims (e.g., username, role, expiration
  time).
- **Token Signing**: The server signs the token using a secret key or a public/private key pair.
- **Token Return**: The signed JWT token is returned to the client.
- **Client Storage**: The client stores the JWT token securely (e.g., local storage, cookies).
- **Subsequent Requests**: The client includes the JWT token in the Authorization header of subsequent requests to the
  server.
- **Server Verification**: The server verifies the JWT token’s signature and payload to authenticate the user and
  authorize access to protected resources.

#### Pros

- **Stateless Authentication**: JWTs are self-contained and do not require server-side session management, making them
  stateless and scalable.
- **Compact and Efficient**: JWTs are compact and can be easily transmitted over networks, reducing bandwidth usage and
  improving performance.
- **Security**: JWTs are digitally signed, ensuring data integrity and authenticity. They can also be encrypted to
  protect sensitive information.
- **Single Sign-On (SSO)**: JWTs can facilitate SSO across multiple applications, eliminating the need for
  re-authentication.
- **Flexible and Customizable**: JWTs can be used with various algorithms, such as HMAC SHA256 or RSA, and can be
  customized to fit specific use cases.
- **Easy to Implement**: JWTs are widely supported and have libraries available for most programming languages, making
  implementation relatively straightforward.

#### Cons

- **Token Size Limitations**: JWTs have a maximum size limit (typically 8KB), which can be a constraint for applications
  requiring large amounts of data.
- **Token Revocation**: Revoking a JWT can be challenging, as it requires updating the server-side database and
  invalidating the token.
- **Token Tampering**: While JWTs are digitally signed, tampering with the payload can still occur if the signature is
  not verified on the server-side.
- **Key Management**: Secure key management is crucial for JWTs, as compromised keys can lead to unauthorized access.
- **Limited Support for Complex Scenarios**: JWTs may not be suitable for complex scenarios, such as multi-factor
  authentication or fine-grained access control.
- **Debugging Challenges**: Debugging issues with JWTs can be difficult due to their compact nature and lack of explicit
  session information.

### Session-Based Authentication

- User logs in
- Server creates a session and stores it in memory or database
- Server sends a session ID to the client (cookie)
- Client sends session ID with each request
- Server looks up session information for each request

#### Pros

- Easy to implement
- Immediate session termination
- More control over active sessions
- Simple to revoke access

#### Cons

- Scalability challenges
- Increased server-side storage
- Difficult to manage in distributed systems
- Performance overhead for session lookups

### Honorable Mentions

- OAuth 2.0 / OpenID Connect (Login with Google, Apple Sign-in)
- Two-Factor Authentication (2FA)
- Biometric Authentication
- API Keys

### Recommendations

1. Always use HTTPS
2. Implement token expiration
3. Use secure, httpOnly cookies
4. Rotate encryption keys
5. Implement proper error handling
6. Use multi-factor authentication

## Authentication in DB

There are multiple approaches for different use cases. For now, we will cover basic with "username/password" strategy.
