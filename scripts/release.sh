#!/bin/bash

set -euo pipefail

verify_on_main_branch() {
    local current_branch
    current_branch=$(git rev-parse --abbrev-ref HEAD)
    if [[ "$current_branch" != "main" ]]; then
        echo "Error: Must be on the 'main' branch to run the release script (currently on $current_branch)."
        exit 1
    fi
}

validate_args() {
    if [[ $# -ne 1 ]]; then
        echo "Usage: $0 <semantic version> (e.g., v1.0.0)"
        exit 1
    fi
}

validate_tag_begins_with_v() {
    local tag=$1
    if [[ ! "$tag" =~ ^v ]]; then
        echo "Error: tag must start with 'v' (e.g., v1.0.0)"
        exit 1
    fi
}

validate_semver() {
    local tag=$1
    # strip the v from the tag since it is technically not a valid semantic version
    # https://semver.org/spec/v2.0.0.html#is-v123-a-semantic-version
    tag=${tag#v}

    # https://semver.org/spec/v2.0.0.html#is-there-a-suggested-regular-expression-regex-to-check-a-semver-string
    # See: https://regex101.com/r/vkijKf/1/
    semver_regex="^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(?:-((?:0|[1-9][0-9]*|[0-9]*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9][0-9]*|[0-9]*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$"

    if ! grep -E "$semver_regex" <<< "$tag"; then
        echo "Error: tag must follow semantic versioning (e.g., v1.0.0)"
        exit 1
    fi
}

validate_version_greater_than_latest() {
    local tag=$1
    tag=${tag#v}

    local latest_tag
    latest_tag=$(git describe --tags --abbrev=0)
    latest_tag=${latest_tag#v}

    if [[ "$(printf "%s\n%s" "$tag" "$latest_tag" | sort -V | tail -n 1)" != "$tag" ]]; then
        echo "Error: Provided version ($tag) must be greater than the latest tag ($latest_tag)"
        exit 1
    fi
}

verify_on_main_branch
validate_args "$@"
validate_tag_begins_with_v "$@"

TAG=$1
validate_semver "$TAG"
validate_version_greater_than_latest "$TAG"

git tag -a "$TAG" -m "$TAG"
git push origin "$TAG"

echo "Tag $TAG created and pushed successfully."
echo "View the release workflow: https://github.com/lucasrod16/oss-contribute/actions/workflows/deploy.yml"
