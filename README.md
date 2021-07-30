# GHPRTemplate

Go-Hasura-Postgres-React Boilerplate Template

## Pre-requisites

- nodejs
- npm | yarn
- docker
- docker-compose
- go

## Using template

- Delete contents of `src/models/` and run `hasura` again.
- Delete contents of `src/controllers/`.
- Search for the keyword `gphr` and remove all instances.
- Sample react code in [PERNTemplate views](https://github.com/atb00ker/PERNTemplate/tree/main/src/views).
- `.env.dev` should during development to run actions outside of containers.

## Running Template

- Rename `.env.example` to `.env` and update configurations as per requirement.
- Build application image `docker-compose build`
- Start application `docker-compose up`

## Setup

1. Run `npm install --include=dev` in the root of the repository.
2. Get Go dependencies: `go mod tidy`
3. Open a terminal and run:
```bash
npm run docker
npm run hasura
```
4. On a second terminal:
```bash
npm run metadata
npm run migrate
npm run metadata-reload
npm run actions
```

## Testing

1. Run `npm install --include=dev` in the root of the repository.
2. Get Go dependencies: `go mod tidy`
3. Open a terminal and run:
```bash
npm run docker-test
npm run hasura-test
```
4. On a second terminal:
```bash
npm run metadata-test
npm run migrate-test
npm run metadata-reload
npm run actions-test
```

## Creation

1. Initialize NPM: `npm init`
2. Install Hasura: `npm install hasura`
3. Initialize Hasura: `npx hasura init`
4. Install React:
```bash
npm install --save react react-dom react-router-dom react-use-websocket react-test-renderer
npm install --save typescript @types/node @types/react @types/react-dom @types/jest
npm install --save-dev webpack webpack-cli webpack-dev-server html-webpack-plugin dotenv-webpack
npm install --save-dev @babel/core babel-loader @babel/preset-env @babel/preset-react @babel/preset-typescript
npm install --save-dev css-loader style-loader file-loader
npm install --save axios
# GraphQL
npm install --save-dev apollo-client react-apollo apollo-cache-inmemory apollo-link-http graphql-tag 
# Testing
npm install --save-dev @testing-library/react @testing-library/jest-dom jest react-test-renderer
npm install --save @types/jest
```
5. Setup Go development tools & dependencies:
```bash
go get
go install github.com/go-delve/delve/cmd/dlv
go install github.com/pilu/fresh
```

### Example Queries

```
mutation {
  insert_gphr_one(object: {gphr: "value1"}) {
    Id,
    gphr
  }
}

subscription {
  ghpr{
    Id,
    ghpr
  }
}

mutation ($ghpr: String!) {
  insert_ghpr_one(object: {ghpr: $ghpr}) {
    Id,
  	ghpr
  }
} {"ghpr": "value1"}

mutation {
  InsertGhpr(ghpr: "actions5") {
    Id
    ghpr
  }
}
```
