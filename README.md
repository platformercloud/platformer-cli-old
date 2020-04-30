# Platformer CLI
Platformer Tools can be used to manage, and deploy your Platformer project from the command line.

* Deploy code and assets to your Firebase projects
* Covert configuration files to K8s format.

To get started with the Platformer CLI, read the full list of commands below

## Install
``go get platformer/cli``

## Options
``-V, --version`` - output the version number   
``-P, --project <alias_or_project_id>`` - the Platformer  project to use for this command  
`` --token <token>  `` - supply an auth token for this command  
``-h, --help  `` - output usage information     
``--debug`` - print verbose debug output and keep a debug log file      

## Commands 
* ``platformer init`` - Creates a scaffolding project if run from an empty directory. If run from an existing source directory, converts it to a project directory by creating a .platformer directory.
* ``platformer login`` - Log the CLI into Platformer
* ``platformer logout`` - Log the CLI out of the Platformer
* ``platformer convert <type>`` - CLI convert existing configuration to different configuration files
* ``platformer project list`` - List all Platformer projects
* ``platformer deploy`` - Deploy 
* ``platformer help`` - Display help information
* ``platformer use`` -  Set an active Platformer for your working directory.