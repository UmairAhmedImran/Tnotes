# T-Notes — Commands & Libraries (Master Document)

## Purpose
This document lists the full CLI surface for **T-Notes** (commands, subcommands, flags, examples) and the recommended Go libraries/modules for building the application. Use this to scaffold Cobra commands or generate CLI help text.

---

## How to Use This Document
Each command block includes:
- Description
- Flags
- Example usages
- Notes for implementation

A separate section describes Go library choices.

---

## Default Behavior — `tnotes` (no args)
Running `tnotes` with no arguments should show a compact dashboard.

### Suggested Dashboard Layout
- **Header:** `T-Notes — <#notes> notes • <#tasks> tasks • <#due> due today`
- **Recent Notes:** ID, title, tags, modified time, short preview
- **Due Tasks:** ID, task text, due date, priority
- **Footer:** hints like `tnotes add`, `tnotes search`, `tnotes ui`

---

## Global Flags
- `--json`, `--yaml`
- `--quiet`, `-q`
- `--verbose`, `-v`
- `--yes`, `-y`
- `--dry-run`
- `--format <table|long|compact|md>`
- `--workspace <name>`

---

# FULL COMMAND SURFACE
Below are all top‑level commands.

---

## `tnotes add` — Create a note or task
**Flags:**
- `--title`, `-t`
- `--body`, `-b`
- `--file <path>`
- `--tags`, `-g`
- `--format <md|text>`
- `--pin`, `--unpin`
- `--secure`
- `--task` (task mode)
- `--due <date>`
- `--priority <low|med|high>`
- `--repeat <pattern>`
- `--estimate <minutes>`
- `--source <url>`
- `--meta key=value`
- `--interactive`

**Examples:**
- `tnotes add -t "Daily Log" -b "Wrote tests" -g journal`
- `tnotes add --interactive --task --due 2025-08-15 --priority high`

---

## `tnotes edit` — Edit an existing note/task
Usage: `tnotes edit <id>`

**Flags:**
- `--title`
- `--body`, `--file`
- `--tags`, `--add-tag`, `--remove-tag`
- `--pin`, `--unpin`
- `--editor <editor>`
- `--no-open`

**Examples:**
- `tnotes edit 23`
- `tnotes edit 23 --title "Updated plan" --add-tag urgent`

---

## `tnotes view` — Show full note or task
**Flags:**
- `--raw`
- `--json`
- `--history`
- `--follow-links`

**Examples:**
- `tnotes view 23`
- `tnotes view 23 --raw`

---

## `tnotes list` — List notes and tasks
**Flags:**
- `--all`
- `--tasks`
- `--tag <tag>`
- `--workspace <ws>`
- `--from <date>` / `--to <date>`
- `--due-soon <N>`
- `--status <open|done|archived>`
- `--sort <field>`
- `--format table|long|compact|json|yaml`

**Examples:**
- `tnotes list --tag work --format table --sort modified --desc`
- `tnotes list --tasks --due-soon 7 --format json`

---

## `tnotes search` — Full-text & fielded search
**Flags:**
- `--fuzzy`
- `--regex`
- `--case-sensitive`
- `--fields title,body,tags`
- `--limit`, `--offset`
- `--interactive`

**Examples:**
- `tnotes search "meeting notes" --fuzzy --limit 10`
- `tnotes search "tag:ideas title:ai"`

---

## `tnotes delete`
**Flags:**
- `--hard`
- `--confirm`
- `--tag <tag>`
- `--older-than <duration>`
- `--dry-run`

**Examples:**
- `tnotes delete 23`
- `tnotes delete --tag archived --older-than 365d --dry-run`

---

## `tnotes tag`
**Subcommands:**
- `tag list`
- `tag add <tag> --to <id>`
- `tag remove <tag> --from <id>`
- `tag rename <old> <new>`
- `tag merge <target> <source...>`
- `tag delete <tag> --confirm`

---

## `tnotes task`
**Subcommands:**
- `task add`
- `task list`
- `task done <id>` / `task reopen <id>`
- `task assign <id> --to <user>`
- `task link <task-id> --note <note-id>`
- `task snooze <id> --until <date>`
- `task backlog`
- `task today`

---

## `tnotes link`
- `link add <from-id> <to-id> --type <kind>`
- `link remove <from-id> <to-id>`
- `link list <id>`
- `backlink <id>`

---

## `tnotes sync`
**Subcommands:**
- `sync init --provider <git|s3|drive>`
- `sync status`
- `sync push` / `sync pull`
- `sync auto --interval <duration>`
- `sync conflict resolve <id>`

---

## `tnotes export` / `tnotes import`
**Export Flags:**
- `--format md|json|yaml|html|pdf`
- `--output <path>`
- `--split-by tag|workspace`

**Import Flags:**
- `--from <path>`
- `--format markdown|jrnl|notable`
- `--dedupe`
- `--map-tags "a:b"`

---

## Encryption Commands
- `tnotes encrypt <id>`
- `tnotes decrypt <id>`
- `tnotes key generate`
- `tnotes key list`
- `tnotes key delete`

---

## `tnotes ui` — Terminal UI
**Flags:**
- `--theme <dark|light>`
- `--port <num>` (optional preview server)

Includes:
- Full-screen navigation
- Keybindings (j/k/o/e/t/d/q)
- Search + preview + edit

---

## `tnotes plugin`
- `plugin list`
- `plugin install <url>`
- `plugin remove <name>`
- `plugin update <name>`
- `plugin run <name> -- <args>`

---

## Maintenance Commands
- `tnotes reindex`
- `tnotes index status`
- `tnotes index optimize`
- `tnotes backup`
- `tnotes restore`
- `tnotes doctor`

---

## Config Commands
- `config set <key> <value>`
- `config get <key>`
- `config edit`
- `config reset --all`

---

## Misc Commands
- `tnotes completion <shell>`
- `tnotes version`
- `tnotes help`

---

# Recommended Development Order
1. add, edit, view, list, search
2. tag + task system
3. config, completion, version
4. sync, import/export, encryption
5. TUI (`tnotes ui`)

---

# Libraries
- **Cobra** — CLI framework
- **Viper** — config management
- **Bubble Tea** — TUI
- **Lipgloss** — styling
- **Glamour** — Markdown rendering
- **SQLite / Badger** — storage
- **Bleve / Meilisearch** — search index

---

_End of Master Document._

