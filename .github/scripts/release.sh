#!/usr/bin/env bash
# Copyright 2026 Bhaskar
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


set -euo pipefail


VERSION=${1:-}


if [ -z "$VERSION" ]; then

    echo "Usage:"
    echo "./release.sh v1.0.0"

    exit 1

fi


if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then

    echo "Invalid version format: $VERSION"
    echo "Expected: v1.0.0"

    exit 1

fi


echo "Preparing release $VERSION"



echo "Checking git status..."

if [ -n "$(git status --porcelain)" ]; then

    echo "Working directory is not clean."
    echo "Commit or stash changes before releasing."

    exit 1

fi



echo "Formatting..."

gofmt -w .



echo "Running vet..."

go vet ./...



echo "Running tests..."

go test ./...



echo "Building..."

go build ./...



if git rev-parse "$VERSION" >/dev/null 2>&1; then

    echo "Tag $VERSION already exists."

    exit 1

fi



echo "Creating tag..."

git tag -a "$VERSION" -m "Release $VERSION"



echo "Pushing tag..."

git push origin "$VERSION"



echo
echo "Released $VERSION"