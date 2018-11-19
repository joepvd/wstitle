# wstitle

`wstitle` is intended to be bound under a key in i3 or sway. It will pop up a
dialogue containing the current name (if available). It will set the name of the
workspace to entered name.

## Installation

Ensure your GOPATH is set up. Get the dependencies:

```sh
go get -u github.com/gen2brain/dlgs
go get -u go.i3wm.org/i3
go build .
cp wstitle ~/bin # Or wherever your treasure trove is located
```

## Configuration
Add a line like the following to your i3 configuration:

```
bindsym F12 exec --no-startup-id wstitle
```

# License
I promise I will not send my lawyers after you if you would use it.
