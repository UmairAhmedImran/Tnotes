tnotes/
├── .github/
│ ├── workflows/
│ │ ├── ci.yml # defaults.run.working-directory: src
│ │ └── release.yml
│ ├── ISSUE_TEMPLATE/
│ └── PULL_REQUEST_TEMPLATE.md
├── docs/
│ ├── adr/
│ ├── architecture/ # this document lives here
│ ├── flows/
│ └── commands.md
├── src/
│ ├── bin/ # Built binaries (gitignored)
│ │
│ ├── migrations/ # sqlx-cli default location
│ │ ├── <ts>\_init.up.sql / .down.sql
│ │ ├── <ts>\_tags.up.sql / .down.sql
│ │ └── <ts>\_fts.up.sql / .down.sql
│ │
│ ├── internal/
│ │ ├── app/
│ │ │ └── app.go # Composition root: config → db → services → workflow
│ │ │
│ │ ├── cli/ # Cobra layer (thin: parse, call, print)
│ │ │ ├── root.go
│ │ │ ├── note.go
│ │ │ ├── task.go
│ │ │ ├── tag.go
│ │ │ ├── search.go
│ │ │ ├── sync.go
│ │ │ ├── tui.go
│ │ │ └── output.go
│ │ │
│ │ ├── config/
│ │ │ ├── config.go
│ │ │ └── paths.go
│ │ │
│ │ ├── db/
│ │ │ ├── db.go # Open() \*sqlx.DB, pragmas (WAL, foreign_keys)
│ │ │ ├── tx.go # DBTX interface, WithTx helper
│ │ │ ├── testdb.go # In-memory DB for tests
│ │ │ └── models/ # Shared entities + pure rules. Depends on nothing.
│ │ │ ├── note.go # Note struct, Validate(), Preview()
│ │ │ ├── task.go # Task struct, status enum, Validate(), IsOverdue(now)
│ │ │ └── tag.go # Tag struct, Validate(), name normalization
│ │ │
│ │ ├── errs/
│ │ │ └── errs.go # ErrNotFound, ErrInvalid, ErrConflict
│ │ │
│ │ ├── tag/ # Leaf module
│ │ │ ├── repository.go
│ │ │ ├── service.go
│ │ │ └── tag_test.go
│ │ │
│ │ ├── task/ # Imports tag. Reaches notes via SQL join, never imports note.
│ │ │ ├── repository.go # includes task_tags queries
│ │ │ ├── service.go
│ │ │ └── task_test.go
│ │ │
│ │ ├── note/ # Imports tag, task
│ │ │ ├── repository.go # includes note_tags queries, ListByTag
│ │ │ ├── service.go
│ │ │ └── note_test.go
│ │ │
│ │ ├── search/ # FTS5 queries across notes + tasks
│ │ │ ├── service.go # Result type + queries, uses db/models directly
│ │ │ └── search_test.go
│ │ │
│ │ ├── workflow/ # Cross-module operations (see rules below)
│ │ │ ├── workflow.go # Service struct: db + note, task, tag services
│ │ │ ├── complete_note.go # e.g. complete a note and all its tasks in one tx
│ │ │ └── workflow_test.go
│ │ │
│ │ ├── sync/ # The one place with an interface
│ │ │ ├── provider.go # Provider interface: Push, Pull, Status
│ │ │ ├── service.go # Orchestrates providers, conflict handling
│ │ │ ├── service_test.go
│ │ │ ├── git/
│ │ │ │ └── git.go
│ │ │ └── gdrive/
│ │ │ └── gdrive.go
│ │ │
│ │ ├── tui/
│ │ │ ├── app.go
│ │ │ ├── theme.go
│ │ │ ├── keys.go
│ │ │ ├── messages.go
│ │ │ └── views/
│ │ │ ├── dashboard/ (model.go, view.go)
│ │ │ ├── noteeditor/ (model.go, view.go)
│ │ │ └── tasklist/ (model.go, view.go)
│ │ │
│ │ └── platform/
│ │ ├── logger/
│ │ │ └── logger.go
│ │ └── timeutil/
│ │ └── timeutil.go
│ │
│ ├── .env
│ ├── .env.example
│ ├── Makefile # unchanged from the src/ version
│ ├── go.mod
│ ├── main.go
│ └── go.sum
├── .gitignore # src/bin/, src/.env
├── LICENSE
├── README.md
└── CHANGELOG.md

Dependency rules

cli, tui ─┐
├──→ workflow ──→ note ──→ task ──→ tag ──→ db, db/models, errs
├──→ sync ──────→ note, task, tag
└──→ search ────→ db, db/models
app ──→ wires everything (the only package that knows all of them)

- Dependencies point down only. Higher layers import lower ones, and never the reverse. Go enforces this with import cycles, so treat a cycle as a design signal, not an obstacle.
- db/models imports nothing internal (only errs and the standard library). Everything else may import it.
- task never imports note. It reaches notes through SQL joins on note_id. note may import task.
- Repositories never import other modules. They run SQL against shared models. If a query spans tables, the owning module writes the join.
- Services take concrete types, not interfaces. The only interface is sync.Provider, defined in sync where it’s consumed, and app.go injects the git and gdrive implementations.
- Repositories are built from db.DBTX, so the same code runs on *sqlx.DB or inside a *sqlx.Tx. A service needing atomicity calls db.WithTx and constructs tx-bound repositories in the callback.
- search imports no services. It queries FTS tables directly and returns its own result type.
- cli and tui call services or workflows, never repositories. Multi-step rules never live in Cobra handlers or Bubble Tea models, because the other interface would have to duplicate them.

Workflow rules
One file per operation, named for what it does (complete_note.go), not a generic grab-bag.
It depends on the modules; the modules never know it exists.
It owns the transaction: it opens db.WithTx and builds tx-bound repositories or services inside.
It should hold a handful of operations. If you can’t name an operation without the word “and” being the whole point, split it.
When to merge two modules

First try the other fixes for a cycle, in this order: extract the narrow shared piece down (a model method or a small function), move the coordination up into workflow/, or invert one direction with a small interface wired in app/.

Merge when most of these are true:

They change together: a feature touching one almost always touches the other.
They share the same tables, or one can’t exist without the other (a task that requires a note).
Their operations usually need one shared transaction.
Neither module makes sense alone.
Most of the real logic ends up in workflow/, leaving both modules as thin query wrappers.

Don’t merge when the cycle is just one or two calls. That’s a signal to extract or lift, not to combine. If you do merge, collapse the two into one package (for example note containing both note and task repositories and services), keep the models in db/models, and delete the now-unneeded workflow file.

note and task are the likeliest pair to merge, since tasks hang off notes. The structure above starts them separate because note → task is a clean one-way dependency. If you find yourself adding the second or third workflow between them, merge them.
