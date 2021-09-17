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

## Setting up authoring v1.x

To set up your patterns for authoring you should walk through the following steps in the authoring application.

1. Upload lexicons under the "Lexicon" tab

- Click the button with the cloud icon with the arrow in the center.
- Drag or click to upload any .lcn files for the data set you are using to create the lexicons. For example:
  `authoring_data/v1.x/person_phone_call_job/person_names.lcn`

2. Under the "Authoring" tab upload the event pattern and event type files

- Click the button with the cloud icon with the arrow in the center.
- Drag or click to upload any .cet and .cep files. For example:
  `authoring_data/v1.x/person_phone_call_job/person_phone_call_job.cet`
  `authoring_data/v1.x/person_phone_call_job/person_phone_call_job.cep`

3. Create a new deployment or create one by uploading a .dts file under the "Deployment tab.
4. When the deployment has been created under the dropdown click "Deploy"
5. In the pop up select all the patterns you want to deploy under that deployment and click "Deploy"

After this the status should eventually change to "Running" and and you should see event the deployment and event definitions appear in Workstation. Then run the data generator to push to the topics created by the deployment.
