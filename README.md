# cogynt-datagenerator-go

A tool used for generating data needed for ingesting into the Cogynt system when working with event patterns.

## Running the application

To start the application first build the application from the root directory using:
`go build main.go`

Then run it using:
`./main`

You will then be asked a series of questions including what data sets you want to create, some questions pertaining to each data set, and whether you want to send it to kafka or create a json file.

## Configuration

All configurations can be set in the `config.yaml` file. There are current defaults set for pointing to our personal cogynt environments.
