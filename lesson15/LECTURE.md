# Lesson 15: Introduction to DB

1. **SQL-based relational databases**: These databases use SQL as their primary query language and
   follow the relational model, which organizes data into tables with well-defined relationships between them.
   Examples of SQL-based relational databases include _MySQL_, _PostgreSQL_, _Microsoft SQL Server_, and _Oracle_.
2. **Non-SQL relational databases**: These databases also follow the relational model but do not use SQL
   as their primary query language. Instead, they may use alternative query languages or proprietary query languages.
   Examples of non-SQL relational databases include _IBM DB2_, _Informix_, and _Teradata_.
3. **NoSQL databases**: databases are designed to handle large amounts of unstructured or semi-structured data and
   often use alternative data models, such as key-value, document-oriented, or graph databases.
   Examples of NoSQL databases include _MongoDB_, _Cassandra_, and _Neo4j_.

## SQL

SQL is a programming language designed to interact with relational databases,
allowing users to **store**, **update**, **remove**, **search**, and **retrieve** information from these databases.
It is a standardized language, meaning that it is widely accepted and used across various platforms and systems.

**Key Features**:

- **Declarative language**: SQL is a declarative language, meaning that you specify what you want to do with your data,
  rather than how to do it.
- **Query language**: SQL is a query language, allowing you to ask questions about the data in your database.
- **Relational database management**: SQL is designed to work with relational databases, which store data in tables
  with defined relationships between them.
- **Static, strong typing**: SQL has a static, strong typing system, which means that the data types of
  variables are known at compile time.

```sql
CREATE TABLE books (
    isbn   char(14)      NOT NULL,
    title  varchar(255)  NOT NULL,
    author varchar(255)  NOT NULL,
    price  decimal(5, 2) NOT NULL
);

INSERT INTO books (isbn, title, author, price)
VALUES ('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
       ('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
       ('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);

ALTER TABLE books
    ADD PRIMARY KEY (isbn);
```

## SQL in Go

To access databases in Go, you use a `sql.DB`. You use this type to create statements and transactions,
execute queries, and fetch results.

The first thing you should know is that a `sql.DB` isn’t a database connection.
It also doesn’t map to any particular database software’s notion of a “database” or “schema.”
It’s an abstraction of the interface and existence of a database, which might be as varied as a local file,
accessed through a network connection, or in-memory and in-process.

`sql.DB` performs some important tasks for you behind the scenes:

- It opens and closes connections to the actual underlying database, via the driver.
- It manages a pool of connections as needed, which may be a variety of things as mentioned.

The `sql.DB` abstraction is designed to keep you from worrying about how to manage concurrent
access to the underlying datastore. A connection is marked in-use when you use it to perform a task,
and then returned to the available pool when it’s not in use anymore.
One consequence of this is that if you fail to release connections back to the pool,
you can cause sql.DB to open a lot of connections, potentially running out of
resources (too many connections, too many open file handles, lack of available network ports, etc).
We’ll discuss more about this later.

After creating a `sql.DB`, you can use it to query the database that it represents,
as well as creating statements and transactions.

## ORM

Object-relational mapping (ORM) is a technique (a.k.a. design pattern) of accessing a relational database from an
object-oriented language (Java, for example). There are multiple implementations of ORM in almost every language; for
example: Hibernate for Java, ActiveRecord for Ruby on Rails, Doctrine for PHP, and SQLAlchemy for Python.

Mainly you will meet scepticism using ORM in go. Because they are overhead and not flexible enough. But code generation
is a great tool to use in go. Some of them: [sqlc](https://sqlc.dev/), [pggen](https://github.com/jschaf/pggen) or my
favourite [jet](https://github.com/go-jet/jet).

## Migrations

A database migration is a systematic approach to managing incremental changes to the database schema in a controlled and
versioned manner. It allows you to define and apply a set of changes to the database structure, such as adding tables,
modifying columns, or altering constraints, through a series of migration files.

- Maintain a consistent and versioned database schema across different environments.
- Collaborate effectively with team members by storing migration files in version control.
- Automate the process of applying schema changes during deployments.
- Easily roll back changes if needed.
- Ensure the integrity and reproducibility of your database schema.

## Data Seeding

Data seeding is the process of populating the database with an initial set of data. The benefits are that you
immediately start testing your application with data without having to insert data manually. Since it’s automatic, you
can initialize the database with thousands of data if you want to test your indexing algorithm.
