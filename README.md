# GoRedis

A minimal Redis compatible server built in Go that supports basic RESP (Redis Serialization Protocol) parsing and a few core commands.

<img width="1920" alt="Screenshot 2025-06-12 at 7 09 47 AM" src="https://github.com/user-attachments/assets/b6c90b0e-840d-4f86-b4a0-ff4bd8bfe5b9" />

---

## Features

- RESP (Redis Serialization Protocol) parser implemented from scratch.
- Handles core Redis commands like `PING`, `SET`, `GET`, `HSET`, `HGET`, and `HGETALL`.
- Uses Go standard library only. No external dependencies.
- Supports writing persistent data using Append-Only File (AOF) format.

---

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/Prtik12/GoRedis.git
cd GoRedis
````

### 2. Run the Server

Use this command:

```bash
go run .
```

> **Note:** Running `go run .` is preferred instead of `go run main.go` to ensure all files in the package are included.

---

## Connecting to the Server

You can connect using the official `redis-cli`:

### Option 1: Using `redis-cli` (Recommended)

#### 🐧 Linux / 🖥 macOS

Most systems already have `redis-cli` or you can install it using:

```bash
# macOS
brew install redis

# Ubuntu/Debian
sudo apt install redis-tools
```

#### 🪟 Windows

Use [Redis for Windows](https://github.com/tporadowski/redis/releases) or install via WSL and run:

```bash
redis-cli -p 6379
```

### Option 2: Using `netcat` (for debugging)

```bash
nc localhost 6379
```

Then manually enter RESP commands like:

```
*1\r\n$4\r\nPING\r\n
```

---

## Example Usage with redis-cli

Start your server:

```bash
go run .
```

In another terminal:

```bash
redis-cli -p 6379
```

Then:

```bash
127.0.0.1:6379> PING
PONG

127.0.0.1:6379> ECHO "hello"
"hello"

127.0.0.1:6379> SET key1 value1
OK

127.0.0.1:6379> GET key1
"value1"

127.0.0.1:6379> HSET user name Alice
(integer) 1

127.0.0.1:6379> HGET user name
"Alice"

127.0.0.1:6379> HGETALL user
1) "name"
2) "Alice"
```

---

## Supported Commands

| Command   | Description                          |
| --------- | ------------------------------------ |
| `PING`    | Heartbeat check, returns `PONG`.     |
| `SET`     | Stores a string value for a key.     |
| `GET`     | Retrieves the value of a key.        |
| `HSET`    | Sets a field in a hash.              |
| `HGET`    | Gets the value of a field in a hash. |
| `HGETALL` | Returns all fields & values in hash. |

---

## Notes

* This is **not a full Redis implementation** 
* It’s a learning project focused on core ideas like TCP server handling, RESP protocol parsing, and in-memory data storage.
* Persistence is handled using a basic AOF file that logs commands.
