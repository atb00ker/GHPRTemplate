# GHPRTemplate

Go-Hasura-Postgres-React Boilerplate Template

## Pre-requisites

- nodejs
- npm | yarn
- docker
- docker-compose
- go

## Usage

# TODO: Application Docker Image
- Run `npm install` in the root of the repository.
- Rename `.env.example` to .env and update configurations as per requirement.
- Build application image `docker-compose build`
- Start application `docker-compose up`

## Creation

1. Initialize NPM: `npm init`
2. Install Hasura: `npm install hasura`
3. Initialize Hasura: `npm `
4. Install React:
```bash
npm install --save react react-dom react-router-dom react-use-websocket react-test-renderer
npm install --save typescript @types/node @types/react @types/react-dom @types/jest
npm install --save-dev webpack webpack-cli webpack-dev-server html-webpack-plugin
npm install --save-dev @babel/core babel-loader @babel/preset-env @babel/preset-react @babel/preset-typescript
npm install --save-dev css-loader style-loader file-loader
npm install --save axios
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
