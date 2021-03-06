# [4.5.0](https://github.com/ae-lexs/ae_albums_api/compare/v4.4.0...v4.5.0) (2022-06-01)


### Features

* Dockerize the project. ([1691b62](https://github.com/ae-lexs/ae_albums_api/commit/1691b62221cd4d9a7cc480da24b3b632f8f6159c))

# [4.4.0](https://github.com/ae-lexs/ae_albums_api/compare/v4.3.0...v4.4.0) (2022-05-21)


### Features

* Add route to GET an album by ID. ([a68e939](https://github.com/ae-lexs/ae_albums_api/commit/a68e939881c30ba072fcf1dc4e1ab76c097972df))

# [4.3.0](https://github.com/ae-lexs/ae_albums_api/compare/v4.2.0...v4.3.0) (2022-05-21)


### Features

* Return AlbumNotFoundError in GetByID method of Album repository. ([e43b4a8](https://github.com/ae-lexs/ae_albums_api/commit/e43b4a8787f94016d46364176faacd3b166f7246))

# [4.2.0](https://github.com/ae-lexs/ae_albums_api/compare/v4.1.0...v4.2.0) (2022-05-21)


### Features

* Add GetByID method to album repository. ([a41f0e2](https://github.com/ae-lexs/ae_albums_api/commit/a41f0e2102d3d055e12a113adc01ab329aafd8f4))
* Add NotFound response to the Get method of Album handler. ([6b58511](https://github.com/ae-lexs/ae_albums_api/commit/6b5851124059c07e806c68cd8954eab5c75a7d5d))

# [4.1.0](https://github.com/ae-lexs/ae_albums_api/compare/v4.0.0...v4.1.0) (2022-05-07)


### Features

* Decode request body in the albumREST.Create method. ([6d0512e](https://github.com/ae-lexs/ae_albums_api/commit/6d0512e4a8bd72d15ecb67a203e546e9f5713ee6))

# [4.0.0](https://github.com/ae-lexs/ae_albums_api/compare/v3.0.0...v4.0.0) (2022-05-07)


### Features

* Create a http server for the application. ([87a5f20](https://github.com/ae-lexs/ae_albums_api/commit/87a5f206d82dd6d5c62913546499642910ff845c))


### BREAKING CHANGES

* Refactor the main package using gin as router in a http server.

# [3.0.0](https://github.com/ae-lexs/ae_albums_api/compare/v2.0.3...v3.0.0) (2022-05-07)


### Bug Fixes

* Handler error in Create method of album repository. ([2d891e8](https://github.com/ae-lexs/ae_albums_api/commit/2d891e88c1a104e1b4beb74f1dba554c436b99c9))


### Features

* Add GetAll method to the album repository. ([0330f7b](https://github.com/ae-lexs/ae_albums_api/commit/0330f7b291ff9315b6584fda72127aaccfe4fb02))
* Add model package. ([14ac52c](https://github.com/ae-lexs/ae_albums_api/commit/14ac52ceb1f78cd146d643f1791e2f5a71f3ddec))


### Performance Improvements

* Refactor albumREST handler and album route. ([778b97a](https://github.com/ae-lexs/ae_albums_api/commit/778b97af9cab0f5eef09a113034fca9fc4c4a726))
* Refactor the album respository. ([f64014f](https://github.com/ae-lexs/ae_albums_api/commit/f64014f59b3dc3de42d2fd3e4f273b6607a69805))


### BREAKING CHANGES

* Use the new album repository version based on the album model.

## [2.0.3](https://github.com/ae-lexs/ae_albums_api/compare/v2.0.2...v2.0.3) (2022-05-07)


### Performance Improvements

* Add NewAlbumREST to make repository private. ([a8213c4](https://github.com/ae-lexs/ae_albums_api/commit/a8213c428ab7ac4d5b4dd3eaff2bda4e44d3a66e))

## [2.0.2](https://github.com/ae-lexs/ae_albums_api/compare/v2.0.1...v2.0.2) (2022-05-07)


### Performance Improvements

* Refactor TestCreateAlbum using table driven tests. ([57e0c98](https://github.com/ae-lexs/ae_albums_api/commit/57e0c980a425355b4e4c2b7d541fae665ef69d3b))
* Refactor TestGetAlbums using table driven tests. ([6ca6bd0](https://github.com/ae-lexs/ae_albums_api/commit/6ca6bd074ec0864d2600e94bcced8467740fcf04))

## [2.0.1](https://github.com/ae-lexs/ae_albums_api/compare/v2.0.0...v2.0.1) (2022-05-06)


### Performance Improvements

* Rename album handler tests. ([fd0d4eb](https://github.com/ae-lexs/ae_albums_api/commit/fd0d4ebac894d4c66f6d8adc4f73220e95bcfa53))
* Rename config and client methods. ([59465bb](https://github.com/ae-lexs/ae_albums_api/commit/59465bb55c2f7bb0d0cfe183a3bb9c620e8c6c09))
* Rename the items of the album handler. ([a2404dc](https://github.com/ae-lexs/ae_albums_api/commit/a2404dc5c2d7a4e8871ccfd03d07dfce8003d871))
* Rename the items of the album repository. ([b0e1949](https://github.com/ae-lexs/ae_albums_api/commit/b0e1949876be9ba2d8bbbdb1eee3b81a78c4ed00))
* Rename the items of the album route. ([3362c44](https://github.com/ae-lexs/ae_albums_api/commit/3362c44eaf2ec15ee075af7d9fdaecdfdfd629a3))

# [2.0.0](https://github.com/ae-lexs/ae_albums_api/compare/v1.7.0...v2.0.0) (2022-04-30)


### Features

* Update README file adding Installation and Usage sections. ([5a3b427](https://github.com/ae-lexs/ae_albums_api/commit/5a3b4275edfc5f1338f5d601a3ca0fa9c7dca975))


### Performance Improvements

* Tidy the packages. ([2efefe5](https://github.com/ae-lexs/ae_albums_api/commit/2efefe5b5de9e79bdd1966fc9825dce4b2a71a10))


### BREAKING CHANGES

* Remove unused packages.

# [1.7.0](https://github.com/ae-lexs/ae_albums_api/compare/v1.6.0...v1.7.0) (2022-04-26)


### Features

* Add client package. ([463df1a](https://github.com/ae-lexs/ae_albums_api/commit/463df1a3d22a371f8a37752de3932f46e5835dc8))
* Add config package. ([1ce434a](https://github.com/ae-lexs/ae_albums_api/commit/1ce434ac8f12310404f49955506b7a368eb09720))
* Refactor main using config and client packages. ([bde5d5b](https://github.com/ae-lexs/ae_albums_api/commit/bde5d5b70db57420a103e7b1c664dc1809089b0c))

# [1.6.0](https://github.com/ae-lexs/ae_albums_api/compare/v1.5.0...v1.6.0) (2022-04-25)


### Bug Fixes

* Fix album the creation using the gorm.Model properties. ([a46e954](https://github.com/ae-lexs/ae_albums_api/commit/a46e9547498fcbee209c9ca5e1c2967d89d3adaa))


### Features

* Add CreateAlbum to AlbumHandler. ([75beafb](https://github.com/ae-lexs/ae_albums_api/commit/75beafb418daffdefd466fca91d3913e6a3323d3))
* Add GetAlbums to AlbumHandler. ([3db3918](https://github.com/ae-lexs/ae_albums_api/commit/3db39184ca48c34c27d5516d2a8afb2b70f947ff))


### Performance Improvements

* Refactor the routes using a handler. ([c28b10c](https://github.com/ae-lexs/ae_albums_api/commit/c28b10cec9f713628d3f3cf81b2617bc88d9c115))

# [1.5.0](https://github.com/ae-lexs/ae_albums_api/compare/v1.4.0...v1.5.0) (2022-04-19)


### Features

* Add album handler. ([6632b3f](https://github.com/ae-lexs/ae_albums_api/commit/6632b3f85a36f19e4df682d6301c2948c89023ae))
* Add postgres service in docker compose. ([0ea63f4](https://github.com/ae-lexs/ae_albums_api/commit/0ea63f44c4be4ebfd16683544f0dd2f9e777f436))
* Add route package. ([33934b8](https://github.com/ae-lexs/ae_albums_api/commit/33934b81461ce7c28699abec257564038b9afab1))


### Performance Improvements

* Refactor repository package. ([d6bbac7](https://github.com/ae-lexs/ae_albums_api/commit/d6bbac75fb9d3d168b610e1ad259df5296f36ff2))
* Remove handler package. ([0f8dc59](https://github.com/ae-lexs/ae_albums_api/commit/0f8dc594dba11387f29e3b91a009e3f074becc08))
* Update the naming of TestGetAlbums. ([10a09e6](https://github.com/ae-lexs/ae_albums_api/commit/10a09e68026da5f9aaa34bcc76f4b665db0e97cc))

# [1.4.0](https://github.com/ae-lexs/ae_albums_api/compare/v1.3.0...v1.4.0) (2022-03-02)


### Features

* Add lint and test to the development workflow. ([6158c57](https://github.com/ae-lexs/ae_albums_api/commit/6158c57aef7e9efdaaa518bd12057fa0a0282be3))

# [1.3.0](https://github.com/ae-lexs/ae_albums_api/compare/v1.2.0...v1.3.0) (2022-03-02)


### Features

* Add Album entity. ([ef7832f](https://github.com/ae-lexs/ae_albums_api/commit/ef7832fea5f850969355c08bae0e568ae746a64c))
* Add Make commands. ([a5e4ac6](https://github.com/ae-lexs/ae_albums_api/commit/a5e4ac65ad1c0a9c5a572bb439ef771f2bf248e7))
* Add PostgresAlbumRepository. ([8a765e5](https://github.com/ae-lexs/ae_albums_api/commit/8a765e579259d27309cdf79725bcd8bb3177d3f3))

# [1.2.0](https://github.com/ae-lexs/ae_albums_api/compare/v1.1.0...v1.2.0) (2022-03-01)


### Features

* Add CircleCI badge. ([6f65e4d](https://github.com/ae-lexs/ae_albums_api/commit/6f65e4dd4e2e52d4b2a35816c73e1dee4e2164e7))

# [1.1.0](https://github.com/ae-lexs/ae_albums_api/compare/v1.0.0...v1.1.0) (2022-03-01)


### Features

* Add base server. ([e9549dc](https://github.com/ae-lexs/ae_albums_api/commit/e9549dcfd290cd51f90a644b29469f205a752335))
* Add go.mod file. ([5377ce0](https://github.com/ae-lexs/ae_albums_api/commit/5377ce09d3f3a02c207169ad03e262e192bff4fb))

# 1.0.0 (2022-03-01)


### Features

* Set up semantic release. ([053674a](https://github.com/ae-lexs/ae_albums_api/commit/053674a347612b26e421bc99056f36ffcb723c05))
