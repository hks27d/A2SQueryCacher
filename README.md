# 📌 A2S Query Cacher

> [!CAUTION]
> - This software is a conceptual mock-up designed to illustrate functionality.
> - It is not intended for actual deployment.

## 🌐 Description
It caches A2S responses matching FFFFFFFF54 (A2S_INFO), FFFFFFFF55 (A2S_PING) and FFFFFFFF41 (A2S_PLAYER).

## 📗 Requirements
1. go >= 1.23.5
2. iptables or other software that mimic the same behaviour for packets redirection as shown below.

## 🔵 Running
> [!WARNING]
> - Bind IP and game server IP must have different values than the default ones.
> - The value for cacheTTL must be between 1 and 30.

1. Download the software.
2. Add the provided iptables rules.
3. Run the compiled binary (A2SQueryCacher).
4. Use command line arguments or the JSON file created at first run.

## 🔌 Building
1. Download and install Go.
2. Clone the repository.
3. Go into its own directory.
4. Build: `go build -o A2SQueryCacher -ldflags="-s -w" cmd/A2SQueryCacher/main.go`

## ⚡ Redirect packets matching A2S_INFO, A2S_PING and A2S_PLAYER
```
iptables -t nat -A PREROUTING -p udp --dport 27015 --match string --algo kmp --hex-string '|FFFFFFFF54|' -j REDIRECT --to-ports 9110
iptables -t nat -A PREROUTING -p udp --dport 27015 --match string --algo kmp --hex-string '|FFFFFFFF55|' -j REDIRECT --to-ports 9110
iptables -t nat -A PREROUTING -p udp --dport 27015 --match string --algo kmp --hex-string '|FFFFFFFF41|' -j REDIRECT --to-ports 9110
```

## 🔎 CLI arguments 

> [!IMPORTANT]
> - CLI arguments are checked first.
> - If they are not provided the program will continue automatically with the JSON config.

```
bindip - Local IP address to bind (default: "0.0.0.0")
gameserverip - Game server IP address (default: "127.0.0.1")
bindPort - Local port to bind (default: 9110)
gameserverport - Game server port (default: 27015)
cacheTTL - Cache TTL in seconds (default: 10)
threads - Number of worker threads (default: 4)
```

## ⭐ Show your support
Give a ⭐ if this project helped you.
