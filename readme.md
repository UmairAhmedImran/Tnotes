# Tnotes

Tnotes is a terminal-based note-taking tool designed for managing project-related and personal notes during development. It functions similar to a todo list but with enhanced note-taking capabilities.

## Installation Guide

### Prerequisites

Ensure Go is installed on your system. Follow the official installation guide: [Go Documentation](https://go.dev/doc/install)

### Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/UmairAhmedImran/Tnotes
   cd Tnotes
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Initialize the Project**

   ```bash
   go run main.go init
   ```

## Usage

### Adding Notes

Add a note with both title and content:

```bash
go run main.go add -t "title" -c "content"
```

Alternatively, add a note with just a title (you'll be prompted to enter content):

```bash
go run main.go add -t "title"
```

When prompted, type your content and press Enter.

### Viewing Notes

List all note titles:

```bash
go run main.go list
```

List notes with content summary:

```bash
go run main.go list -r y
```

