# Pocket telegram bot
Pocket allows the user to save an article or web page to remote servers for later reading. The article is then sent to the user's Pocket list (synced to all of their devices) for offline reading. Pocket removes clutter from articles, and allows the user to add tags to their articles and to adjust text settings for easier reading.

## Structure
```bash
├── bot.db
├── cmd
│   └── bot
│       └── main.go
├── go.mod
├── go.sum
├── Makefile
└── pkg
    ├── repository
    │   ├── boltdb
    │   │   └── token.go
    │   └── token.go
    ├── server
    │   └── server.go
    └── telegram
        ├── auth.go
        ├── bot.go
        └── handlers.go
```
