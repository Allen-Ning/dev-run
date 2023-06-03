Available commands:

- `clone`: Clone the specified GitHub repositories.
- `docker-up`: Run Docker Compose in the downloaded repositories.
- `list-services`: List all available services in the downloaded repositories.
- `run-service`: Run a particular service within all downloaded repositories.

For more information about each command, run `myproject [command] --help`.

## Configuration

The tool requires a GitHub personal access token to access the repositories. Set the `GITHUB_TOKEN` environment variable to your personal access token.

Additionally, set the `REPOSITORIES` environment variable to a comma-separated list of the GitHub repositories you want to download.
