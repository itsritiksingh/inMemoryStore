# QuickCache (Persistant inMemory key-value store) (Go Library)

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Usage via cli](#usage-via-cli)
- [How it Works](#how-it-works)
  - [Write Ahead Log (WAL)](#write-ahead-log-wal)
  - [String Sorted Tables (SSTables)](#string-sorted-tables-sstables)
- [Contributing](#contributing)

## Introduction

This is an in-memory key-value store written in Go, designed to provide tolerance against system crashes. It achieves crash tolerance through the use of a Write Ahead Log (WAL) and String Sorted Tables (SSTables). This library is suitable for scenarios where you need a fast, efficient, and crash-resistant in-memory key-value store for your Go applications.

### Inspired From : "The Log-Structured Merge-Tree (LSM-Tree)" by Patrick O'Neil, Edward Cheng, Dieter Gawlick, and Elizabeth O'Neil (1996)

## Features

- In-memory key-value storage.
- Crash tolerance against system crashes.
- Efficient and high-performance storage.
- Write Ahead Log (WAL) for durability.
- String Sorted Tables (SSTables) for efficient data storage.
- Easy-to-use Go library.

## Getting Started

### Installation

You can install this library using Go modules. To add it to your project, run:

```bash
go get github.com/itsritiksingh/inMemoryStore
```

### Usage
```go
func main() {
	store := store.Init()
	store.Put("hello", "world")
	store.Put("key","value")
	store.Upsert("hello", "world1")

	err , val := store.Get("hello")
	if err != nil {
		fmt.Errorf("error fetching key %v",err)
	}
	fmt.Println(val)
	fmt.Println(store.GetAllKeys())
}
```

### Usage via cli
```bash
#list all keys
keys

#get a value
get <keyname>

#store a key value pair
set <key> <value>

```
## How it Works

### Write Ahead Log (WAL)

This library utilizes a Write Ahead Log (WAL) to ensure durability. When you put or delete key-value pairs, the changes are first written to the WAL before they are applied to the in-memory store. In the event of a system crash, the WAL can be replayed to recover the previous state of the store, ensuring data integrity.

### String Sorted Tables (SSTables)

The concept of String Sorted Tables (SSTables) is inspired by the book. SSTables are a way of organizing key-value data on disk in a sorted order, making it efficient for read operations.

#### Key Features of SSTables:
1. Sorted Structure: SSTables maintain key-value pairs in a sorted order based on the keys. This sorting allows for efficient lookups and range queries.

2. Merging and Compaction: When multiple SSTables are used, they can be merged efficiently to remove duplicate keys and retain only the most recent value for each key. This process ensures that the number of segments is manageable, reducing the overhead of checking multiple hash maps during lookups.


## Contributing

Contributions to this project are welcome! If you want to contribute, please fork this repository, make your changes, and submit a pull request. Be sure to follow the code of conduct and the contribution guidelines.
