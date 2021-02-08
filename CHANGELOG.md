# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.1.0] - 2020-02-08
### Changed
- golang to 1.15
### Added 
- logger using zip
- add labels to secret
- add tests directory to mock repository and logger
- add coverage

## [1.0.0] - 2019-12-10
### Added
- CRUD to manage secrets inside kubernetes using rest api.
- ENCODING_REQUEST variable to accepted only encoded requests.
- Add tests in appcontext, controller and usecase
- Add deploy script to control tag and branch versions