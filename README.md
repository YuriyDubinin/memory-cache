# Cache Package

## Overview

The `cache` package is a simple and efficient Go library for managing an in-memory cache. It allows you to store, retrieve, and delete key-value pairs, as well as automatically clean up expired entries. This package is useful in scenarios where quick data retrieval and controlled data expiration are required.

## Features

- **Custom Types:**
  - `Key`: A type representing the key in the cache (`string`).
  - `Value`: An interface representing the value, allowing any data type.

- **Cache Structure:**
  - `Cache`: The main structure that holds the key-value pairs in a map.

- **Constructor:**
  - `New() *Cache`: Creates and returns a new instance of the cache.

- **Methods:**
  - `Get(key Key) (Value, error)`: Retrieves the value associated with a given key. If the key does not exist, it returns an error indicating that the key was not found.
  - `Set(key Key, value Value, lifetime time.Duration)`: Adds or updates a key-value pair in the cache.
  - `Delete(key Key) error`: Deletes a key-value pair from the cache. If the key does not exist, it returns an error indicating that the key was not found.
  - `Automatic Data Expiration`: Entries with expired lifetimes are automatically cleared from the cache.

- **Automatic Clearing of Outdated Data:**
  - The cache runs a background goroutine that periodically checks and removes expired entries. The interval for checking is set to 1 second by default.
  - `Stop Mechanism`: The cache has a mechanism to stop the automatic clearing process when necessary, by sending a signal through a channel.


## Installation

To use the `cache` package in your project, you can install it via `go get`:

```bash
go get github.com/YuriyDubinin/memory-cache
```

## Testing

To test the `cache` package, you can run the tests file with:

```bash
go test -v
```