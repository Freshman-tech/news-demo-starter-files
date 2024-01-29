# News demo web app

News application built with Go. It was created following the tutorial created by
[ayoisiah](https://github.com/ayoisaiah) at: https://freshman.tech/web-development-with-go/

## Prerequisites

- You need to have [Go](https://golang.org/dl/) installed on your computer. The
version used at creation was **1.21.6**.

- The application also depends on the [godotenv package](https://github.com/joho/godotenv).

- Finally, you need to have an API key from [News API](https://newsapi.org). You can sign up for a free key [here](https://newsapi.org/account).

## Get started

- Clone or download this repository to your filesystem.

```bash
git clone https://github.com/drmattheath/news-demo-web-app
```

- Install the `godotenv` package.

```bash
go get github.com/joho/godotenv
```

- Create a `.env` file, then populate it with your News API key and the port you wish the application to be available at.

```env
PORT=
NEWS_API_KEY=
```

- Compile the code.

```bash
go build
```

- Execute the binary to start the application.

```bash
./news-demo-web-app
```

