# wstitle

`wstitle` is intended to be bound under a key in i3 or sway. It will pop up a
dialogue containing the current name (if available). It will set the name of the
workspace to entered name.

There are two modi: Default (and per command line argument `-window`, the title
of the currently focused window is taken. With the argument `-workspace`, the
current workspace title will be selected.

## Installation

```sh
make build
cp wstitle ~/bin
```

## Configuration
Add a line like the following to your i3 configuration:

```
bindsym F12 exec --no-startup-id wstitle
```

# License
I promise I will not send my lawyers after you if you would use it.
