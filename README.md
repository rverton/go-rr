# Go webapp blueprint

Go web app blueprint for dynamic content with a type-safe SQL backend.

Technologies used:
* [Echo](https://echo.labstack.com/) web framework
* sqlite
* [sqlc](https://sqlc.dev/) for type-safe SQL
* react-router for frontend routing and data fetching

## Usage

Development:

```
make sqlc/generate
make frontend/watch
make run
```
