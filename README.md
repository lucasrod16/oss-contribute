# OSS Contribute

[OSS Contribute](https://osscontribute.com/) is a website I built to help myself and other developers find the best open-source projects to contribute to. Whether you're a beginner or an experienced developer, finding the right project to contribute to can be challenging and time-consuming.

## How It Works

OSS Contribute searches for projects based on the following criteria to ensure you find active and meaningful projects to work on:

- Has an open-source license
- Must not be archived
- Has at least 10 issues labeled with `good first issue` or `help wanted`
- Has at least 500 stars on GitHub
- Has been updated within the last month to ensure it's actively worked on

## Join the Open Source Community

Find your next project to grow your skills and connect with talented people all over the world.

<https://osscontribute.com/>

## Contributing

Contributions are welcome!

If you have an idea for a new feature or find a bug that needs to be fixed, feel free to submit issues and pull requests.

## Deploying Updates to the Site

To deploy updates, create an annotated git tag using [Semantic Versioning](https://semver.org/) and push it to GitHub.

Pushing the tag will trigger the [deployment pipeline](https://github.com/lucasrod16/oss-contribute/blob/main/.github/workflows/deploy.yml), which will deploy the changes to the live environment.

```shell
# Set the version
TAG=<version> # example: v0.1.0

# Create the tag
git tag -a "$TAG" -m "$TAG"

# Push the tag to GitHub
git push origin "$TAG"
```
