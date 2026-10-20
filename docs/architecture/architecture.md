tnotes/
├── .github/
│ ├── workflows/
│ │ ├── ci.yml
│ │ └── release.yml
│ ├── ISSUE_TEMPLATE/
│ └── PULL_REQUEST_TEMPLATE.md
├── docs/
│ ├── adr/
│ ├── architecture/
│ ├── flows/
│ └── commands.md
├── src/
│ ├── bin/ # Built binaries (gitignored)
│ ├── cmd/
│ │ └── main.go
│ ├── internal/
│ │ ├── app/
│ │ ├── cli/
│ │ ├── config/
│ │ ├── db/
│ │ ├── errs/
│ │ ├── tag/
│ │ ├── note/
│ │ ├── task/
│ │ ├── search/
│ │ ├── sync/
│ │ │ ├── git/
│ │ │ └── gdrive/
│ │ ├── tui/
│ │ └── platform/
│ ├── migrations/
│ ├── .env
│ ├── .env.example
│ ├── Makefile
│ ├── go.mod
│ └── go.sum
├── .gitignore
├── LICENSE
├── README.md
└── CHANGELOG.md

In detail src would contain

src/
│
├── bin/ # Built binaries (gitignored)
│
├── cmd/
│ └── main.go # Entry point: calls cli.Execute() only
│
├── migrations/ # sqlx-cli default location
│ ├── <ts>\_init.up.sql
│ ├── <ts>\_init.down.sql
│ ├── <ts>\_tags.up.sql
│ ├── <ts>\_tags.down.sql
│ ├── <ts>\_fts.up.sql
│ └── <ts>\_fts.down.sql
│
├── internal/
│ ├── app/
│ │ └── app.go # Composition root: config → db → services
│ │
│ ├── cli/ # Cobra layer (thin)
│ │ ├── root.go # Root command, global flags, Execute()
│ │ ├── note.go
│ │ ├── task.go
│ │ ├── tag.go
│ │ ├── search.go
│ │ ├── sync.go
│ │ ├── tui.go
│ │ └── output.go # table/JSON printing helpers
│ │
│ ├── config/
│ │ ├── config.go
│ │ └── paths.go # XDG config/data dirs, DB path
│ │
│ ├── db/
│ │ ├── db.go # Open() \*sqlx.DB, pragmas (WAL, foreign_keys)
│ │ ├── tx.go # WithTx helper
│ │ └── testdb.go # In-memory DB for tests
│ │
│ ├── errs/
│ │ └── errs.go # ErrNotFound, ErrInvalid, ErrConflict
│ │
│ ├── tag/ # No dependencies on other features
│ │ ├── tag.go # Tag struct, Validate()
│ │ ├── repository.go # Concrete struct, sqlx queries
│ │ ├── service.go
│ │ └── tag_test.go
│ │
│ ├── note/ # Imports tag
│ │ ├── note.go # Note struct, Validate()
│ │ ├── repository.go
│ │ ├── service.go
│ │ └── note_test.go
│ │
│ ├── task/ # Imports tag
│ │ ├── task.go # Task struct, status enum, Validate()
│ │ ├── repository.go
│ │ ├── service.go
│ │ └── task_test.go
│ │
│ ├── search/ # FTS5 queries across notes + tasks
│ │ ├── search.go # Result type
│ │ ├── service.go
│ │ └── search_test.go
│ │
│ ├── sync/ # The one place with an interface
│ │ ├── provider.go # Provider interface: Push, Pull, Status
│ │ ├── service.go # Orchestrates providers, conflict handling
│ │ ├── service_test.go
│ │ ├── git/
│ │ │ └── git.go # implements sync.Provider
│ │ └── gdrive/
│ │ └── gdrive.go # implements sync.Provider
│ │
│ ├── tui/
│ │ ├── app.go # Root Bubble Tea model, view routing
│ │ ├── theme.go
│ │ ├── keys.go
│ │ ├── messages.go
│ │ └── views/
│ │ ├── dashboard/
│ │ │ ├── model.go
│ │ │ └── view.go
│ │ ├── noteeditor/
│ │ │ ├── model.go
│ │ │ └── view.go
│ │ └── tasklist/
│ │ ├── model.go
│ │ └── view.go
│ │
│ └── platform/
│ ├── logger/
│ │ └── logger.go
│ └── timeutil/
│ └── timeutil.go
│
├── .env
├── .env.example
├── .gitignore
├── Makefile
├── go.mod
└── go.sum

Dependency rules

- tag depends on nothing internal except db and errs.
- note and task import tag (one direction), so there are no import cycles and no need for a shared domain package.
- search imports db and runs FTS queries directly. It doesn’t need to import note or task; it returns its own result type.
- sync defines Provider in its own package (where it’s consumed), and git/ and gdrive/ implement it. sync doesn’t import them; app.go constructs them and passes them in.
- Only app/ knows about everything.
