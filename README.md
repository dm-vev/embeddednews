# embeddednews

This project provides a simple Go server that fetches the latest news from the Habr RSS feed and exposes them via a JSON API.

## Usage

Run the server:

```bash
go run ./cmd/server
```

Then access `http://localhost:8080/news` to receive a list of news items. Titles are cleaned to contain only English and Russian letters (with `ё` replaced by `е`), digits, spaces and a set of punctuation characters.

To get the full text of a particular article, call `/article?url=<ARTICLE_URL>`. The response is plain text cleaned with the same rules.
