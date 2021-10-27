# Contributing to DarsHub

A big welcome and thank you for considering contributing to DarsHub projects!

Reading and following these guidelines will help us make the contribution process easy and effective for everyone involved. It also communicates that you agree to respect the time of the developers managing and developing these projects. In return, we will reciprocate that respect by addressing your issue, assessing changes, and helping you finalize your pull requests.

---

## Getting Started

Contributions are made to this repo via [Issues](https://dev.azure.com/DarsHub-Orga/DarsHub/_workitems/assignedtome/) and Pull Requests ([PRs](https://dev.azure.com/DarsHub-Orga/_git/DarsHub/pullrequests?_a=mine)). A few general guidelines that cover both:

- Search for existing Issues and PRs before creating your own.
- We work hard to makes sure issues are handled in a timely manner but, depending on the impact, it could take a while to investigate the root cause. A friendly ping in the comment thread to the submitter or a contributor can help draw attention if your issue is blocking.

---

## Issues

All issues are collected in [**Open( Backlog )**](https://dev.azure.com/DarsHub-Orga/DarsHub/_backlogs/backlog/Developers/Backlog%20items).
If an issue was selected, then you move it to **wip(Work-In-Progess)**. This way the team members know that this issue is being worked on.

Each issue includes a meaningful **title** and a good **description** of what exactly is being done.

In the description you are welcome to use Markdown to make the text better, e.g. with tables, images, links or other Markdown features.

- the first heading should start with '#'.
- if needed, a checkbock to check off the tasks also helps
  ```
  [x] done
  [ ] open
  ```

To better manage the issues, **labels** should be used to mark them.

The following pre-selection can be used directly:

- backend/frontend/test
- ci/cd
- docs
- ops

multiple selection is possible :D

Issues should be used to report problems with the library, request a new feature, or to discuss potential changes before a PR is created. When you create a new Issue, a template will be loaded that will guide you through collecting and providing the information we need to investigate.

Adding a [reaction](https://github.blog/2016-03-10-add-reactions-to-pull-requests-issues-and-comments/) can also help be indicating to our maintainers that a particular problem is affecting more than just the reporter.

If the issue could be processed completely, it is moved to **check**.

Now a merge request can be created and the changes can be checked by another person.

If there was no problem and the merge was approved, the issue can be finally closed by moving it to **Closed**

---

## Pull Requests

PRs to our libraries are always welcome and can be a quick way to get your fix or improvement slated for the next release. In general, PRs should:

- Only fix/add the functionality in question **OR** address wide-spread whitespace/style issues, not both.
- Add unit or integration tests for fixed or changed functionality (if a test suite already exists).
- Address a single concern in the least number of changed lines as possible.
- Be accompanied by a complete Pull Request template (loaded automatically when a PR is created).

For changes that address core functionality or would require breaking changes (e.g. a major release), it's best to open an Issue to discuss your proposal first. This is not required but can save time creating and reviewing changes.

---

## Commits

Each commit should start with a type, include a short description, and end with the issue(s) number.

If an issue has been fixed, the issue number should appear in the commit message with a preceding `fix`:

```
fix(#7): TestIssue

//or

fix: TestIssue (#7)

```

### Type

| type                  | description                                                      |
| --------------------- | ---------------------------------------------------------------- |
| ops                   | changes for containers, CI/CD pipeline, etc.                     |
| doc                   | changes to docs (README, GUIDES, WIKI)                           |
| fix                   | a bugfix, compilation of the project/ YAML files of the pipeline |
| backend/frontend/test | changes on code for frontend/backend and tests                   |

---

## Branch

Create a branch locally with a succinct but descriptive name

### Switch to the desired branch

```shell
$ git switch main

$ git checkout dev
```

### Get all changes

```
$ git fetch
$ git pull
```

### Create and switch to a new branch

If an issue is edited, a new branch should be created.
The only requirement would be that either the issue number is in the branch name or that the comments always contain an issue number so that they can be traced later.

```
$ git switch --create 1234-commit-analysis

$ git checkout -b 1234-commit-analysis
```

```
$ git switch --create commit-analysis
$ git checkout -b commit-analysis
$ git commit -m "fix: stuf (#1234)"
//or
$ git commit -m "fix(#1234): stuf (#1234)"
```

In general, we follow the [Git-flow-Workflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow#:~:text=Der%20Git-flow-Workflow%20ist,Release%20des%20Projekts%20konzipiert%20wurde.&text=Git-flow%20ist%20hervorragend%20f%C3%BCr,Release-Zyklus%20nach%20Zeitplan%20geeignet)
