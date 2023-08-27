# sql-to-typescript
Convert your SQL schemas into TypeScript files

# Usage
Install Golang and build from source for your os.

Write your sql schema in a file and pass it to the binary
```sh
sql-to-ts schema.sql
```

This will create a new file in the current location which will contain TypeScript code. You can also specify the name of the output file

```sh
sql-to-ts schema.sql types.ts
```

# Supported types
- text, char, varchar will be parsed into string
- integer, numeric will be parsed into number
- json will be parsed into Record<string, unknown>
- arrays are supported: text[] or integer[] will be parsed into either string[] or number[]

# Constrains
- Every field that does not contain 'not null' will be parsed as `T | undefined`

This is an unfinished project, contributions are welcome :))
