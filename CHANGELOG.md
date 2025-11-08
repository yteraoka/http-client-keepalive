# Changelog

## [v0.5.1](https://github.com/yteraoka/http-client-keepalive/compare/v0.5.0...v0.5.1) - 2025-11-08
- tagpr の導入 by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/49
- fix log.Printf() args by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/50
- tagpr の対象 branch を修正 by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/51
- use `time.Since` instead of `time.Now().Sub` (gosimple) by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/53
- Update dependency golang to v1.23.1 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/47
- Update dependency golang to v1.23.2 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/54
- Update dependency golang to v1.24.6 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/55
- Update actions/create-github-app-token action to v2 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/57
- pinact by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/58
- Update dependency go to v1.24.6 - autoclosed by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/56
- Update actions/checkout action to v4.3.0 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/59
- Update actions/create-github-app-token action to v2.1.4 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/60
- Update actions/checkout action to v5 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/61
- Update goreleaser/goreleaser-action action to v6.4.0 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/64
- Update Songmu/tagpr action to v1.9.0 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/65
- Update actions/setup-go action to v6 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/66
- Update dependency golang to v1.25.4 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/62
- Update dependency go to v1.25.4 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/63

## [v0.5.0](https://github.com/yteraoka/http-client-keepalive/compare/v0.4.0...v0.5.0) - 2024-09-06
- support HTTP_PROXY by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/48

## [v0.4.0](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.9...v0.4.0) - 2024-07-12
- goreleaser v2 に合わせて --rm-dist を --clean に変更 by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/46

## [v0.3.9](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.8...v0.3.9) - 2024-07-12
- .goreleaser.yaml の更新 by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/45

## [v0.3.8](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.7...v0.3.8) - 2024-07-12
- Update actions/setup-go action to v5 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/40
- Update goreleaser/goreleaser-action action to v6 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/43
- Update dependency golang to v1.22.5 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/42
- Update module github.com/google/uuid to v1.6.0 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/41
- Update module github.com/jessevdk/go-flags to v1.6.1 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/44

## [v0.3.7](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.6...v0.3.7) - 2023-12-06
- Update dependency golang to v1.21.5 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/39

## [v0.3.6](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.5...v0.3.6) - 2023-11-08
- Update goreleaser/goreleaser-action action to v4 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/27
- Update actions/setup-go action to v4 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/31
- Update module github.com/google/uuid to v1.3.1 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/33
- Update dependency golang to v1.21.0 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/32
- Update dependency golang to v1.21.2 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/35
- Update actions/checkout action to v4 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/34
- Update goreleaser/goreleaser-action action to v5 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/36
- Update dependency golang to v1.21.4 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/37
- Update module github.com/google/uuid to v1.4.0 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/38

## [v0.3.5](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.4...v0.3.5) - 2023-02-16
- trace が 3 以上の場合に response header を出力する by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/30

## [v0.3.4](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.3...v0.3.4) - 2023-02-06
- goreleaser で arm64 binary も build する by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/29

## [v0.3.3](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.2...v0.3.3) - 2023-02-06
- body の byte 数を出力 by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/28

## [v0.3.2](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.1...v0.3.2) - 2023-02-06
- `--uuid` オプションで uuid parameter を追加する by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/26

## [v0.3.1](https://github.com/yteraoka/http-client-keepalive/compare/v0.3.0...v0.3.1) - 2023-01-31
- Configure Renovate by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/17
- Update module github.com/jessevdk/go-flags to v1.5.0 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/18
- Update goreleaser/goreleaser-action action to v3 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/23
- Update actions/checkout action to v3 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/21
- Update actions/setup-go action to v3 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/22
- Update module go to 1.19 by @renovate[bot] in https://github.com/yteraoka/http-client-keepalive/pull/19
- actions/setup-go で install する version も 1.19 に更新 by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/24
- io.Copy のエラーをチェック by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/25

## [v0.3.0](https://github.com/yteraoka/http-client-keepalive/compare/v0.2.5...v0.3.0) - 2022-05-26
- sleep-at-end オプションを追加 by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/15
- go 1.18 by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/16

## [v0.2.5](https://github.com/yteraoka/http-client-keepalive/compare/v0.2.4...v0.2.5) - 2021-01-25
- Call client.CloseIdleConnections before exit by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/14

## [v0.2.4](https://github.com/yteraoka/http-client-keepalive/compare/v0.2.3...v0.2.4) - 2021-01-23
- Replace `--random-sleep-max-ms` with `--sleep-range-ms` by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/13

## [v0.2.3](https://github.com/yteraoka/http-client-keepalive/compare/v0.2.2...v0.2.3) - 2020-05-18
- Delete unnecessary call RoundTrip() by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/10
- Add httptrace.DNSStart, DNSDone by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/11
- Split the httptrace functions into a separate file by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/12

## [v0.2.2](https://github.com/yteraoka/http-client-keepalive/compare/v0.2.1...v0.2.2) - 2020-05-15
- fix typo by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/7
- Add .gitignore by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/8
- Update go version in workflow by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/9

## [v0.2.1](https://github.com/yteraoka/http-client-keepalive/compare/v0.2.0...v0.2.1) - 2020-05-15
- set some message to lower priority by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/6

## [v0.2.0](https://github.com/yteraoka/http-client-keepalive/compare/v0.1.8...v0.2.0) - 2020-05-15
- Update README by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/4
- Add httptrace by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/5

## [v0.1.8](https://github.com/yteraoka/http-client-keepalive/compare/v0.1.7...v0.1.8) - 2020-05-14
- Override MaxIdleConns with MaxIdleConnsPerHost if MaxIdleConnsPerHost… by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/3

## [v0.1.7](https://github.com/yteraoka/http-client-keepalive/compare/v0.1.6...v0.1.7) - 2020-04-28
- disable interval sleep by default by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/2

## [v0.1.6](https://github.com/yteraoka/http-client-keepalive/compare/v0.1.5...v0.1.6) - 2020-04-28
- Add disable-http-keepalive flag by @yteraoka in https://github.com/yteraoka/http-client-keepalive/pull/1
