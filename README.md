# cogynt-datagenerator-go

A tool used for generating data needed for ingesting into the Cogynt system when working with event patterns.

## Running the application

To start the application first build the application from the root directory using:
`go build main.go`

Then run it using:
`./main`

You will then be asked a series of questions including what data sets you want to create, some questions pertaining to each data set, and whether you want to send it to kafka or create a json file.

## Configuration

All configurations can be set in the `config.yaml` file. There are current defaults set for pointing to our personal cogynt environments. If you need to add new configurations remove the file from the .gitignore, commit the defaults, and then add it back to the .gitignore so that we aren't stomping the default configs with our personal ones.

## Event pattern directory

There are also some authoring event patterns, event types, and lexicons built around these data sets. These can be found under there respective folders and can be added via authoring. I will hopefully add instructions on how to upload these at some point but for now ask a team member. The `authoring_data` directory has the files for uploading. Feel free to modify them in authoring and export them to create new ones.
