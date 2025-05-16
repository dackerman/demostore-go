# Changelog

## 2.0.0 (2025-05-16)

Full Changelog: [v1.0.0...v2.0.0](https://github.com/dackerman/demostore-go/compare/v1.0.0...v2.0.0)

### Features

* **api:** update via SDK Studio ([05a5e57](https://github.com/dackerman/demostore-go/commit/05a5e57ae60f40c2b882771a8cb0b7afcbfb176e))

## 1.0.0 (2025-05-16)

Full Changelog: [v0.4.0...v1.0.0](https://github.com/dackerman/demostore-go/compare/v0.4.0...v1.0.0)

### Features

* **api:** update via SDK Studio ([615e28c](https://github.com/dackerman/demostore-go/commit/615e28ccaaf22a26cdfa6a2ed7d08cc69ed838ae))
* **api:** update via SDK Studio ([eaca349](https://github.com/dackerman/demostore-go/commit/eaca349e2e678b3d07271f407ebd801838484f1f))
* **api:** update via SDK Studio ([2c729db](https://github.com/dackerman/demostore-go/commit/2c729dbc5fb83ceb60b4b89c3f75d71d5d6cdb16))
* **client:** add support for endpoint-specific base URLs in python ([d70e34d](https://github.com/dackerman/demostore-go/commit/d70e34db90bdaa4d684dc8b1d87ddd27b22ab79d))

## 0.4.0 (2025-05-07)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/dackerman/demostore-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** update via SDK Studio ([#64](https://github.com/dackerman/demostore-go/issues/64)) ([a522b77](https://github.com/dackerman/demostore-go/commit/a522b77c4b0f09b5c44f6da229886a6df25cecc0))
* **client:** add support for reading base URL from environment variable ([3f04088](https://github.com/dackerman/demostore-go/commit/3f04088a9ba88945782152f96426b5813fdfc89a))
* **client:** support custom http clients ([#65](https://github.com/dackerman/demostore-go/issues/65)) ([c4e224c](https://github.com/dackerman/demostore-go/commit/c4e224c5385f034d2e3ae8156f3b4e4994e1f79a))


### Bug Fixes

* **client:** clean up reader resources ([2cbd258](https://github.com/dackerman/demostore-go/commit/2cbd258f4f8ea4a99f04ddb8e861e18f868dad00))
* **client:** correctly update body in WithJSONSet ([0523be4](https://github.com/dackerman/demostore-go/commit/0523be479ae2306b2bf55ca0e5601ed561e7d39a))
* **client:** dedup slashes in URL ([#62](https://github.com/dackerman/demostore-go/issues/62)) ([17d8d5d](https://github.com/dackerman/demostore-go/commit/17d8d5d1bdf2b75d7924dc7115d581fde50e03e8))
* **client:** return error on bad custom url instead of panic ([#63](https://github.com/dackerman/demostore-go/issues/63)) ([46306ee](https://github.com/dackerman/demostore-go/commit/46306ee9fce77c2a0f4ee1d1265fd76dc95bd210))
* handle empty bodies in WithJSONSet ([28adf94](https://github.com/dackerman/demostore-go/commit/28adf945e4764877a50643fbf1c0afd17b5039d9))
* **pagination:** handle errors when applying options ([6cef623](https://github.com/dackerman/demostore-go/commit/6cef6238e4e615413ade0de0813bdcf2460ea0ed))
* **test:** return early after test failure ([#60](https://github.com/dackerman/demostore-go/issues/60)) ([28903a1](https://github.com/dackerman/demostore-go/commit/28903a181acd8d22ea565fd8634c19a5501356e9))


### Chores

* add request options to client tests ([#59](https://github.com/dackerman/demostore-go/issues/59)) ([c0532f4](https://github.com/dackerman/demostore-go/commit/c0532f4a07986ea07f1a237f9828848511e8d93b))
* **ci:** add timeout thresholds for CI jobs ([ddabfde](https://github.com/dackerman/demostore-go/commit/ddabfdeeff3ce4145c3ca91bbad1d80fd1d03e99))
* **ci:** only use depot for staging repos ([44999bc](https://github.com/dackerman/demostore-go/commit/44999bc93221fe4cd734916a95376717976b4800))
* **docs:** document pre-request options ([cc03554](https://github.com/dackerman/demostore-go/commit/cc035540085082ef2dd89c518447697656b06136))
* **docs:** improve security documentation ([#57](https://github.com/dackerman/demostore-go/issues/57)) ([4843910](https://github.com/dackerman/demostore-go/commit/4843910bf32df111ae52d09195b91bebb1cb019d))
* fix typos ([#61](https://github.com/dackerman/demostore-go/issues/61)) ([8a38ad7](https://github.com/dackerman/demostore-go/commit/8a38ad7a1105ab87dc45ca8266960fafbb0661cb))
* **internal:** codegen related update ([ac4476a](https://github.com/dackerman/demostore-go/commit/ac4476af7eb95c506e19547f44714f73f0529ef5))
* **internal:** expand CI branch coverage ([f89709b](https://github.com/dackerman/demostore-go/commit/f89709b59bf565632c76e46b0d8fcc04fcc12eab))
* **internal:** reduce CI branch coverage ([f0a7037](https://github.com/dackerman/demostore-go/commit/f0a703737b9a43f2dc70631d50a6e7cc2f584ed7))


### Documentation

* update documentation links to be more uniform ([3476b8e](https://github.com/dackerman/demostore-go/commit/3476b8e12fba9e2e75616d0474b591c3f3559709))

## 0.3.0 (2025-03-20)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/dackerman/demostore-go/compare/v0.2.0...v0.3.0)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#45](https://github.com/dackerman/demostore-go/issues/45)) ([a364e4d](https://github.com/dackerman/demostore-go/commit/a364e4dd9011caafa2ef4cc1f04214e3787e5e1d))
* **api:** update via SDK Studio ([#37](https://github.com/dackerman/demostore-go/issues/37)) ([cc1e8f4](https://github.com/dackerman/demostore-go/commit/cc1e8f4b3aba18165b5d41856e2097245a2de57e))
* **api:** update via SDK Studio ([#38](https://github.com/dackerman/demostore-go/issues/38)) ([7650dc4](https://github.com/dackerman/demostore-go/commit/7650dc4dfa26c4c523a97d13f60c7be653c4a523))
* **api:** update via SDK Studio ([#39](https://github.com/dackerman/demostore-go/issues/39)) ([7f17fd0](https://github.com/dackerman/demostore-go/commit/7f17fd0269e7a85ea4633de7ae7ef6657f8f5595))
* **api:** update via SDK Studio ([#41](https://github.com/dackerman/demostore-go/issues/41)) ([bb1b939](https://github.com/dackerman/demostore-go/commit/bb1b93920c915a69101ec58ebe4842f32bac35e6))
* **api:** update via SDK Studio ([#50](https://github.com/dackerman/demostore-go/issues/50)) ([cde8e7d](https://github.com/dackerman/demostore-go/commit/cde8e7dfe936406c08119e804d5dcc94639cd977))
* **api:** update via SDK Studio ([#51](https://github.com/dackerman/demostore-go/issues/51)) ([a495ac5](https://github.com/dackerman/demostore-go/commit/a495ac5cd0f24c1ba2e394aabbd8dc4bfa7b195f))
* **api:** update via SDK Studio ([#52](https://github.com/dackerman/demostore-go/issues/52)) ([1cb8915](https://github.com/dackerman/demostore-go/commit/1cb891512d9d4a4a12a987ab26a00e84688066a5))
* **api:** update via SDK Studio ([#53](https://github.com/dackerman/demostore-go/issues/53)) ([57a545a](https://github.com/dackerman/demostore-go/commit/57a545a0ba6df1c962c610b3cf9d74b822113607))
* **api:** update via SDK Studio ([#54](https://github.com/dackerman/demostore-go/issues/54)) ([15d3e03](https://github.com/dackerman/demostore-go/commit/15d3e03621d42851fba9ebe578a76d6f47ddb4c5))
* **api:** update via SDK Studio ([#55](https://github.com/dackerman/demostore-go/issues/55)) ([9a71262](https://github.com/dackerman/demostore-go/commit/9a71262cf863aba897ec1b1eef55088ba1208e6d))
* **client:** accept RFC6838 JSON content types ([#46](https://github.com/dackerman/demostore-go/issues/46)) ([24eb0b5](https://github.com/dackerman/demostore-go/commit/24eb0b5fb8889ec199223bdfada31df76eca7abb))
* **client:** allow custom baseurls without trailing slash ([#44](https://github.com/dackerman/demostore-go/issues/44)) ([82d7a06](https://github.com/dackerman/demostore-go/commit/82d7a062c3004f4a73df020e838ddcfd67dbe114))
* **client:** improve default client options support ([#48](https://github.com/dackerman/demostore-go/issues/48)) ([a440add](https://github.com/dackerman/demostore-go/commit/a440adddf7598bc3191673e1696b301c754fda5c))
* **client:** send `X-Stainless-Timeout` header ([#33](https://github.com/dackerman/demostore-go/issues/33)) ([4d8b59a](https://github.com/dackerman/demostore-go/commit/4d8b59a14edd687f4325e302cc0b975c933defda))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#40](https://github.com/dackerman/demostore-go/issues/40)) ([a3722e4](https://github.com/dackerman/demostore-go/commit/a3722e4ce5329de17e5862591da45b3b381692fa))
* do not call path.Base on ContentType ([#36](https://github.com/dackerman/demostore-go/issues/36)) ([6f3f7a9](https://github.com/dackerman/demostore-go/commit/6f3f7a97b80df48ffe623254e88d18eaf8db5065))
* fix early cancel when RequestTimeout is provided for streaming requests ([#35](https://github.com/dackerman/demostore-go/issues/35)) ([40aaa9a](https://github.com/dackerman/demostore-go/commit/40aaa9ac34d6656ff55103fbd95418f5c7bab46e))
* fix unicode encoding for json ([#30](https://github.com/dackerman/demostore-go/issues/30)) ([0840c6c](https://github.com/dackerman/demostore-go/commit/0840c6c9eae3d3571e91daded8b6bb88d9899005))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#34](https://github.com/dackerman/demostore-go/issues/34)) ([5d2b7df](https://github.com/dackerman/demostore-go/commit/5d2b7dfa877c337757d08283f528502f59bd271a))
* **internal:** codegen related update ([#42](https://github.com/dackerman/demostore-go/issues/42)) ([2b3e388](https://github.com/dackerman/demostore-go/commit/2b3e3884a55afb848207fe93dad7a3927fa97962))
* **internal:** remove extra empty newlines ([#49](https://github.com/dackerman/demostore-go/issues/49)) ([9ad7a7f](https://github.com/dackerman/demostore-go/commit/9ad7a7f8bb1a5602f0702d271112d7ebaf0d310b))


### Documentation

* document raw responses ([#32](https://github.com/dackerman/demostore-go/issues/32)) ([f8073c7](https://github.com/dackerman/demostore-go/commit/f8073c78b7cf69699f56c6cc288625879ae0932b))
* update URLs from stainlessapi.com to stainless.com ([#43](https://github.com/dackerman/demostore-go/issues/43)) ([3ca4880](https://github.com/dackerman/demostore-go/commit/3ca48805b8c1d6c3ab9af71a2a8d672c70ac230a))


### Refactors

* tidy up dependencies ([#47](https://github.com/dackerman/demostore-go/issues/47)) ([c1b4db2](https://github.com/dackerman/demostore-go/commit/c1b4db23deb018b7e58c9816c6c92ddfeab46e58))

## 0.2.0 (2025-01-31)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/dackerman/demostore-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** update via SDK Studio ([a87393c](https://github.com/dackerman/demostore-go/commit/a87393cac33b6c826c7f5d8c14854593841e2211))
* **api:** update via SDK Studio ([#26](https://github.com/dackerman/demostore-go/issues/26)) ([5a524de](https://github.com/dackerman/demostore-go/commit/5a524de6d910f00ffb85598919057e3b3012765d))


### Chores

* go live ([#28](https://github.com/dackerman/demostore-go/issues/28)) ([0fb6e1f](https://github.com/dackerman/demostore-go/commit/0fb6e1ffa411c61a8adc92155c1ac9dd9623e8f5))

## 0.1.0 (2025-01-29)

Full Changelog: [v0.1.0-alpha.2...v0.1.0](https://github.com/dackerman/demostore-go/compare/v0.1.0-alpha.2...v0.1.0)

### Features

* **api:** update via SDK Studio ([#23](https://github.com/dackerman/demostore-go/issues/23)) ([b659759](https://github.com/dackerman/demostore-go/commit/b659759df6bc3eee537c0e9efc257d1a4a6b5537))

## 0.1.0-alpha.2 (2025-01-29)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/dackerman/demostore-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** update via SDK Studio ([#16](https://github.com/dackerman/demostore-go/issues/16)) ([0f083fd](https://github.com/dackerman/demostore-go/commit/0f083fd634ae15ae39b873c440517d42f3496250))
* **api:** update via SDK Studio ([#17](https://github.com/dackerman/demostore-go/issues/17)) ([b19a538](https://github.com/dackerman/demostore-go/commit/b19a53883d7ee2da24cf08c2f863ce814b747130))
* **api:** update via SDK Studio ([#18](https://github.com/dackerman/demostore-go/issues/18)) ([cfe80be](https://github.com/dackerman/demostore-go/commit/cfe80be505ecfaacea144cb66bbb3e9460d1e27f))
* **api:** update via SDK Studio ([#20](https://github.com/dackerman/demostore-go/issues/20)) ([527c31c](https://github.com/dackerman/demostore-go/commit/527c31c46d4bed6e17faac5abc75233ff36527cb))
* **api:** update via SDK Studio ([#21](https://github.com/dackerman/demostore-go/issues/21)) ([0a3a430](https://github.com/dackerman/demostore-go/commit/0a3a430c3eefe5eb300de10667c8501c3aaf9a04))


### Chores

* **internal:** codegen related update ([#12](https://github.com/dackerman/demostore-go/issues/12)) ([7c5af9c](https://github.com/dackerman/demostore-go/commit/7c5af9c64c84a6e8afd9a99e2121f09cecf49cee))
* **internal:** codegen related update ([#14](https://github.com/dackerman/demostore-go/issues/14)) ([2f3e1fa](https://github.com/dackerman/demostore-go/commit/2f3e1fa5556e2aed8a8378945074e56836a5c7ad))
* **internal:** codegen related update ([#15](https://github.com/dackerman/demostore-go/issues/15)) ([705e4b7](https://github.com/dackerman/demostore-go/commit/705e4b7264ef6da3addc4f129e12c979115bef51))
* refactor client tests ([#19](https://github.com/dackerman/demostore-go/issues/19)) ([05d5814](https://github.com/dackerman/demostore-go/commit/05d5814b02a81fa57da8c885aae96b4ea033aded))

## 0.1.0-alpha.1 (2024-12-06)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/dackerman/demostore-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([#10](https://github.com/dackerman/demostore-go/issues/10)) ([a2906f7](https://github.com/dackerman/demostore-go/commit/a2906f749efd9d4b6747e13979b269e2deb834aa))
* **api:** update via SDK Studio ([#2](https://github.com/dackerman/demostore-go/issues/2)) ([7d43467](https://github.com/dackerman/demostore-go/commit/7d434674d1101f0610b3b214efd4e66d18422658))
* **api:** update via SDK Studio ([#3](https://github.com/dackerman/demostore-go/issues/3)) ([b447b8b](https://github.com/dackerman/demostore-go/commit/b447b8b16ed22a07cf3f19c0083a591653c0daad))
* **api:** update via SDK Studio ([#4](https://github.com/dackerman/demostore-go/issues/4)) ([0f2e9af](https://github.com/dackerman/demostore-go/commit/0f2e9afc9c9ba5f310da08a72838c3394502f18f))
* **api:** update via SDK Studio ([#5](https://github.com/dackerman/demostore-go/issues/5)) ([844f1a2](https://github.com/dackerman/demostore-go/commit/844f1a2c7e0404b74ab2672407dfaaf1c0b93463))
* **api:** update via SDK Studio ([#6](https://github.com/dackerman/demostore-go/issues/6)) ([58c37d6](https://github.com/dackerman/demostore-go/commit/58c37d6b65d2e72b572ed646e938ec6be859c89a))
* **api:** update via SDK Studio ([#7](https://github.com/dackerman/demostore-go/issues/7)) ([5223292](https://github.com/dackerman/demostore-go/commit/5223292aee799b1e88d646a4e8199f8eca659068))
* **api:** update via SDK Studio ([#8](https://github.com/dackerman/demostore-go/issues/8)) ([626fe38](https://github.com/dackerman/demostore-go/commit/626fe38df2fc8415f0ec07fc5a084299266390f7))
* **api:** update via SDK Studio ([#9](https://github.com/dackerman/demostore-go/issues/9)) ([7ab4239](https://github.com/dackerman/demostore-go/commit/7ab42397fd4021695d5d5f756a1e6975f66b8738))


### Chores

* go live ([43dc8a5](https://github.com/dackerman/demostore-go/commit/43dc8a50f8cb056daf43adb34e58af29ff35a096))
