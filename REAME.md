We are using this structure for our backend application

## Go Directories (Base)

| Folder    | Description                                                                                                                          |
| --------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| /cmd      | Main applications for this project.                                                                                                  |
| /internal | Private application and library code.                                                                                                |
| /pkg      | Library code that's ok to use by external applications.                                                                              |
| /vendor   | Application dependencies (managed manually or by your favorite dependency management tool like the new built-in Go Modules feature). |

## Service Application Directories

| Folder | Description                                                          |
| ------ | -------------------------------------------------------------------- |
| /api   | OpenAPI/Swagger specs, JSON schema files, protocol definition files. |

## Common Application Directories

| Folder       | Description                                                                                       |
| ------------ | ------------------------------------------------------------------------------------------------- |
| /configs     | Configuration file templates or default configs.                                                  |
| /init        | System init (systemd, upstart, sysv) and process manager/supervisor (runit, supervisord) configs. |
| /scripts     | Scripts to perform various build, install, analysis, etc operations.                              |
| /build       | Packaging and Continuous Integration.                                                             |
| /deployments | IaaS, PaaS, system and container orchestration deployment configurations and templates.           |
| /test        | Additional external test apps and test data.                                                      |

## Other Directories

| Folder       | Description                                                                             |
| ------------ | --------------------------------------------------------------------------------------- |
| /docs        | Design and user documents (in addition to your godoc generated documentation).          |
| /tools       | Supporting tools for this project.                                                      |
| /examples    | Examples for your applications and/or public libraries.                                 |
| /third_party | External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).    |
| /githooks    | Git hooks.                                                                              |
| /assets      | Other assets to go along with your repository (images, logos, etc).                     |
| /website     | This is the place to put your project's website data if you are not using GitHub pages. |
