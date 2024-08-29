# Memory Cache Package

## Overview

The `cache` package is a simple and efficient Go library that provides basic operations for managing an in-memory cache. It allows you to store, retrieve, and delete key-value pairs with ease. This package can be useful for scenarios where quick data retrieval is required.

## Features

- **Custom Types:**
  - `Key`: A type representing the key in the cache (string).
  - `Value`: An interface representing the value, allowing any data type.

- **Cache Structure:**
  - `Cache`: The main structure that holds the key-value pairs in a map.

- **Constructor:**
  - `New() *Cache`: Creates and returns a new instance of the cache.

- **Methods:**
  - `Get(key Key) Value`: Retrieves the value associated with a given key. If the key does not exist, it prints a message and returns `nil`.
  - `Set(key Key, value Value)`: Adds or updates a key-value pair in the cache.
  - `Delete(key Key)`: Deletes a key-value pair from the cache.

## Installation

To use the `cache` package in your project, you can install it via `go get`:

```bash
go get github.com/YuriyDubinin/memory-cache