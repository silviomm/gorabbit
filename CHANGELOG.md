# Change Log
All notable changes to this project will be documented in this file.
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [0.2.0] - 2023-11-15

### Fixed
- `queues delete` connected to CurrentContext rabbit but used LocalRabbit config to get queue names to delete;

### Changed
- `gorabbit move-msgs` command is now `gorabbit msg shovel`

### Added
- CMD `queues new`;
- CMD `msg send`;
- `-q` flag on `msg shovel` CMD to choose which queues will be shoveled; 
- `-q` flag on `queues delete` CMD to choose which queues will be deleted (defaults to all);
- Feedback on `queues delete` of error or successful deletion;
- `context` commands help section;

## [0.1.0] - 2023-10-12

### Added
- CLI Configuration file;
- Default context;
- Prompt to add new context;
- CMD `context`: `creation`, `list`, `set`, `current`;
- CMD `queues`: `get`, `delete`;
- CMD `move-msgs`;

#### Dev - Internal
- Added bumpversion configuration;
- Added go build release workflow;