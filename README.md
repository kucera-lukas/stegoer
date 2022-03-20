# client

[![Production deployment](https://img.shields.io/github/deployments/stegoer/client/Production?label=vercel&logo=vercel&logoColor=vercel)](https://stegoer.vercel.app/)
[![Continuous Integration](https://github.com/stegoer/client/actions/workflows/ci.yml/badge.svg)](https://github.com/stegoer/client/actions/workflows/ci.yml)
[![Docs](https://github.com/stegoer/client/actions/workflows/docs.yml/badge.svg)](https://github.com/stegoer/client/actions/workflows/docs.yml)
[![pre-commit.ci status](https://results.pre-commit.ci/badge/github/stegoer/client/main.svg)](https://results.pre-commit.ci/latest/github/stegoer/client/main)
[![code style: prettier](https://img.shields.io/badge/code_style-prettier-ff69b4.svg?style=flat-square)](https://prettier.io/)

Client is using TypeScript, NextJS and GraphQL.

---

Client website: https://stegoer.vercel.app/

Source code: https://github.com/stegoer/client

---

## Installation

### Install dependencies

```sh
npm install
```

### Create the `.env.local` file

Create a `.env.local` file and copy the contents of `.env.local.example` file into the `.env.local` file

```sh
cp .env.local.example .env.local
```

## Development

### Dev server

```sh
npm run dev
```

### GraphQL

[GraphQL Code Generator](https://www.graphql-code-generator.com/)
is used to generate type definitions and hooks for queries and mutations. See `client/src/graphql/codegen.yml` for
configuration options.

#### Codegen

```sh
npm run gen
```

To add a new query or mutation head to `src/graphql/user` or
`src/graphql/image` and add a new file.

To add a new fragment head to the `src/graphql/fragments` folder.

### Docs

```sh
npm run build-docs
```

[TypeDoc](https://github.com/TypeStrong/typedoc) is used to generate documentation
which is then published via the
[Docs GitHub Action](https://github.com/stegoer/client/blob/master/.github/workflows/docs.yml)
on [GitHub Pages](https://pages.github.com/).

## Contributing

```sh
pre-commit install
```

## License

Developed under the [MIT](https://github.com/stegoer/client/blob/master/LICENSE) license.
