# Go API Template

A template for creating an HTTP API service written in Go

## Creating a New Service

1. Create a new repository (without a _.gitignore_ file) and clone it to your local
2. Copy all of the files (except the _.git_ directory, _go.mod_, _go.sum_ and _README.md_) to the new repository
3. Search and replace the following values using case sensitivity:
   | Value | Replace With | Example |
   | ----- | ------------ | ------- |
   | Template Service | The service's readable name | Awesome Service |
   | template | The service's name in lowercase without "Service" | awesome |
   | template-service | The service's name in kebab case | awesome-service |
   | TemplateService | The service's name in Pascal case | AwesomeService |
   | TS | The service's name abbreviated in uppercase | AS |
   | TemplateDB | The database name for the service | Awesome |
4. Rename the _vscode_ directory to _.vscode_
5. Run `go mod init awesome-service`
6. Run `go build`

## Adding a Database

1. Create the database
2. Search for and uncomment the two _TODO_'s
3. Set the configuration in _.vscode/launch.json_

## Notes

1. This repository will run as a service named template-service. This was done to make testing changes to the project easier.
2. It is not meant for production deployment as is.
