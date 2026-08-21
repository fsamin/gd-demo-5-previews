# gd-demo-5-previews — step 5: a branch, a URL

Last of the five graduated demo repositories for
[git-deploy-operator](https://github.com/fsamin/git-deploy-operator).

**What this repository proves:** push a branch, get a URL of its own; delete the
branch, it goes away. A preview is a real application — same pipeline, same
gates, its own hostname.

A preview does not know which branch it came from: the platform injects neither
the branch nor the commit. So what makes two preview URLs look *different* is
the code on them. Here that is one block of four constants at the top of
`main.go` — a branch label, a headline, an accent colour and a sentence — which
is exactly the shape of a feature branch.

## Demo

```sh
git-deploy -n demo init --name previews
git-deploy preview enable --pattern 'feat/*' --max 2
git push origin feat/dark-mode
```

Within a poll interval the *Previews* card on the detail page fills in with the
branch, its state and its URL. Open it next to the application's own URL: the
two pages are visibly different.

Three `feat/*` branches ship here on purpose, against a limit of two. The third
one shows up in the card as *waiting for a free slot* rather than not showing up
at all — an absence the page has to explain, or it reads as a bug.

| Branch | Headline | Accent |
|---|---|---|
| `main` | Production | blue |
| `feat/dark-mode` | Dark mode | near-black |
| `feat/pricing` | New pricing | green |
| `feat/search` | Search | orange |

On a preview's detail page, *Edit* and *Delete* are deliberately absent — the
application would recreate the preview at the next poll, and buttons that appear
to work and are silently reverted are worse than no buttons. Its own
*Configuration* page stays fully usable: a preview's variables were copied from
the application when it was created and are its own from then on.

## Locally

```sh
go run .          # http://localhost:8080
```
