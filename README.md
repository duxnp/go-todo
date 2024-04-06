# Todos

[Go](https://go.dev/) + [Echo](https://echo.labstack.com/) + [Templ](https://templ.guide/) + [HTMX](https://htmx.org/) + [Tailwind CSS](https://tailwindcss.com/)

I guess there will be some pages for adding and deleting Todos, but this is mostly a bunch of tests to learn how to make a web app with Go.

I used [Go Blueprint](https://github.com/Melkeydev/go-blueprint) to generate a project from its project templates.

## Getting Started

[Install Go](https://go.dev/doc/install).

[Install Air](https://github.com/cosmtrek/air?tab=readme-ov-file#installation). Air is used to help live reload the app during development.

[Install Node](https://nodejs.org/en). Node is used to download javascript libraries like HTMX and the Tailwind CLI. You could skip Node if your CSS didn't need a build step and downloaded the necessary javascript libraries manually.

Run `npm install` to download the dependencies listed in package.json.

TODO: What about DB migration?

Run `make watch` to start the dev server.

Take a look at [this](vscode.md) for an idea on a development workflow when using VS Code.

## MakeFile

run all make commands with clean tests

```bash
make all build
```

build the application

```bash
make build
```

run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB container

```bash
make docker-down
```

live reload the application

```bash
make watch
```

run the test suite

```bash
make test
```

clean up binary from the last build

```bash
make clean
```
