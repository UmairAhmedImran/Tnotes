tnotes/
│
├── cmd/                            # Cobra command definitions (entry layer)
│   ├── root.go
│   ├── add.go
│   ├── edit.go
│   └── list.go
│
├── internal/                       # Non-exported app logic (core engine)
│   ├── model/                      # Data structures / DB models
│   │   ├── note.go
│   │   ├── task.go
│   │   ├── tag.go
│   │   └── config.go
│   │
│   ├── db/                         # Persistent layer (sqlite / bolt / badger)
│   │   ├── db.go                   # Open/close connection once
│   │
│   ├── service/                    # Core logic per domain
│   │   ├── init_service.go         # Bootstraps DB, config, etc.
│   │   ├── add_service.go         # add/edit/view/list/search
│   │   ├── edit_service.go          # tag CRUD + merge
│   │   └── list_service.go         # task CRUD + done/reopen/snooze
│   ├── view/                       # CLI + future TUI renderers
│   │   ├── add_view.go            # render note
│   │   ├── edit_view.go            # render task list/detail
│   │   ├── list_view.go            # tabular output with filters
│   │   ├── search_view.go          # fuzzy/regex result view
│   │
│   └── utils/                      # Shared helpers
│       ├── time.go
│       ├── io.go
│       ├── errors.go
│       ├── format.go
│       └── prompt.go
│
├── go.mod
├── go.sum
├── main.go                         # Initializes Cobra root command
└── README.md