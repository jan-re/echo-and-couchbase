# Intro

This is a simple Go application integrated with a Couchbase database. It exists because I wanted to learn how the `github.com/labstack/echo` framework worked and I also wanted to learn about Couchbase.

# Functionality

This is the quintessential CRUD application. We're dealing with two API entities:

1. books, housed on `/api/books`
2. themes, housed on `/api/themes`

There are many books. There is a smaller number of themes. Each book can be connected to a variable number of themes.

Supported actions are:

1. Listing all books and themes
2. Creating new books and themes
3. Getting a concrete book or theme by its ID
4. Updating a concrete book or theme
5. Deleting a concrete book or theme
6. Listing all books which feature a given theme
