# Lookmarks

[![builds.sr.ht status](https://builds.sr.ht/~bascht/lookmarks.svg)](https://builds.sr.ht/~bascht/lookmarks?)

Quick helper utility initially prototyped with
[Lucas](https://github.com/moonglum/) to generate a lookup list of commonly used
URLs from a standard Yaml file.

Think of it as a kind-of "bookmark builder" which enables you to generate bookmark list from a hierarchically defined Yaml structure:

```yaml
bookmarks:
  - url: https://example.com
    key: example
    sub:
      - url: foo
      - url: bar
      - sub:
        - url: baz
          key: alias
```

Will generate the following lookup table:

| key               | value                       |
|-------------------|-----------------------------|
| example           | https://example.com         |
| example-foo       | https://example.com/foo     |
| example-bar       | https://example.com/bar     |
| example-bar-alias | https://example.com/bar/baz |


While this might not look that impressive on first sight, think about the possibilities for e.g. modelling GitLab's URL structure for a few different repositories by simply defining a few Yaml anchors and letting `lookmarks` take care of the rest:


```yaml
---
gitlab_layout: &gitlab_urls
  - url: "issues"
    key: "issues"
    sub:
    - url: "%{s}"
      key: jump
    - url: "?scope=all&state=opened&assignee_username=bascht"
      key: mine
    - url: "new"
      key: create
  - url: pipelines
    key: pipelines
  - url: merge_requests

bookmarks:
  - key: gitlab
    url: https://gitlab.com/gitlab-org
    sub:
      - url: gitlab-org
        sub: *gitlab_urls
      - url: www-gitlab-com
        sub: *gitlab_urls
```

will generate URLs for:

- gitlab
- gitlab-gitlab-org
- gitlab-gitlab-org-issues
- gitlab-gitlab-org-issues-create
- gitlab-gitlab-org-issues-jump
- gitlab-gitlab-org-issues-mine
- gitlab-gitlab-org-merge_requests
- gitlab-gitlab-org-pipelines
- gitlab-gitlab-org-tags
- gitlab-gitlab-org-wikis
- gitlab-gl-infra
- gitlab-gl-infra-issues
- gitlab-gl-infra-issues-create
- gitlab-gl-infra-issues-jump
- gitlab-gl-infra-issues-mine
- gitlab-gl-infra-merge_requests
- [???] _you probably see where this is going???_

which you can then use to feed them into the tool of your choice, say `fzf` for example:

![screencap](./lookmarks-with-fzf.gif)
