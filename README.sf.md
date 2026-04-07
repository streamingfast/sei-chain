## StreamingFast Sei Fork Notes

### Remotes & Branches

```bash
git remote set-url origin https://github.com/sei-protocol/sei-chain.git
git remote add sf git@github.com:streamingfast/sei-chain.git
```

We maintain 3 branches:

- `feature/firehose-tracer`
- `feature/firehose-tracer-at-latest-release-tag`
- `release/firehose`

The `release/firehose` contains Dockerfile and instructions how to build images, this is the branch that should be used to make releases.

The `feature/firehose-tracer` is the PR branch that tracks `origin/main` branch, `feature/firehose-tracer-at-latest-release-tag` is the same as `feature/firehose-tracer` but tracks the latest release tag, which as of time of writing is `v5.5.5`.

### Bumping to new version

```bash
git fetch origin

# Found correct tag to bump to, we will use `v6.0.4`
export VERSION=v6.0.4

git checkout feature/firehose-tracer-at-latest-release-tag
git pull

git merge "${VERSION:?}"
# Fix any conflicts
go test ./...
git commit

git checkout feature/firehose-tracer
git pull
git merge  feature/firehose-tracer-at-latest-release-tag
# Fix any conflicts and merge, but there is usually no conflicts in this step

git checkout release/firehose
git pull

git merge feature/firehose-tracer-at-latest-release-tag
# Fix any conflicts and merge, but there is usually no conflicts in this step

make install
# Run Battlefield tests, check https://github.com/streamingfast/battlefield-ethereum/edit/master/README.md#chain-tests for what to run

git tag "${VERSION:?}-fh3.0"

git push feature/firehose-tracer-at-latest-release-tag release/firehose "${VERSION:?}-fh3.0"
```

#### Building Binary & Docker

Binary is built automatically on GitHub actions when pushing a tag, images will be found `ghcr.io/streamingfast/sei-chain:<version>`
