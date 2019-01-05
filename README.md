**This project is still in an early phase. It's usable and working but not yet recommended for production.**

# gotime

gotime is a command line application to track time and to simplify office work.
It's able to track time, to create HTML and PDF reports and to create invoice drafts for a few web-based invoicing application.

`gotime` tracks time locally. Instead of implementing billing, etc. it will rely on 3rd-party tools, i.e. mostly cloud-based services.

## Basic usage
A typical session looks like this:
```bash
gotime start acme
gotime stop
gotime report --month 0
```

## Documentation
Documentation about the command line is available at [docs/markdown](./docs/markdown/gotime.md)

## Data model
The data is stored in a few JSON files on disk. It's easy to backup and still fast.
Right now it's on disk at `$HOME/.gotime/`.
The available commands make it easy to export it into different target formats. Most commands are supporting plain text
and JSON at this time. Other formats may be added in the future.

### Projects
gotime supports nested projects. The separator character is the slash '/'.
The simplest form is a project without any subprojects.

If you need to bill a single client for several distinct project then it's better to use subprojects.
For example:
```bash
gotime create project client1 client1/web client1/backend
gotime start client1/web
gotime stop
gotime report --split project -p client1
```

This will create a report on all projects which belong to client1 with the tracked time per project.

## Tracking time
### Start
### Stop
### Cancel
### View

## Reporting

### Date and time filters
### Splitting options

The reporting allows to split the tracked time into multiple pieces. Multiple levels of splitting is possible.
Possible values:
- year
- month
- project

For example, `-split year` creates a report where the tracked time is grouped into the years which are covered by tracked time.
`-split year,project` will renders groups of years which are further grouped by the projects which were tracked in this particular year. 
`-split project,month` groups the tracked time by project first and then lists the summaries grouped by month for each project. 

### Plain Text Reports

### HTML Reports

### PDF Reports
#### api2pdf service integration

## Create Invoices
### Create invoices at sevdesk.com

## Import from other tools

### Import data from Watson
```bash
gotime import watson
```

### Import data from Fanurio
This needs a custom CSV export (to be documented).
```bash
gotime import fanurio complete-export.csv
```

## License
To be decided