# Proposal of import path tree for gotk4 libs.

```
import (
	"github.com/gotk4/lib/gtk"
	"github.com/gotk4/lib/glib/v2"

	"github.com/gotk4/grun" // App launcher
)
```

Almost unchanged from [diamondburned's work](github.com/diamondburned/gotk4/pkg) except:
* removed `pkg/` prefix
* moved `gtk/v4` code to `gtk`
* moved `gtk/v3` code to `2old/gtk3`
* moved `gen` code to `2gen`
* added `rsvg` with 3 functions to load and draw svg from cairo


The **gtk4 version** is the base version for the lifespan of the gotk4 repo, it should be THE first class citizen.

The **gtk3 version** can help transitioning but is not for long term use, moved to the side (2old = too old).

The **gen** code would not be grouped with other libs in the list (2gen = to generate).

The main question left is wether libs that don't have a version dir (yet) should be put in a v1 subdir or not, to help a future possible migration (ex cairo to cairo/v1). The gotk4 tree will be active until the release of the gtk6 version, so we could expect it to be 10+ years, and be prepared for that...

```
github.com/gotk4/lib/
                   ├── 2gen
                   │   ├── cmd
                   │   └── girgen
                   ├── 2old
                   │   └── gtk3
                   ├── atk
                   ├── cairo
                   ├── gdk
                   │   ├── v3
                   │   └── v4
                   ├── gdkpixbuf
                   │   └── v2
                   ├── gdkpixdata
                   │   └── v2
                   ├── gdkwayland
                   │   └── v4
                   ├── gdkx11
                   │   ├── v3
                   │   └── v4
                   ├── gio
                   │   └── v2
                   ├── glib
                   │   └── v2
                   ├── graphene
                   ├── gsk
                   │   └── v4
                   ├── gtk        // gtk v4 only
                   ├── pango
                   ├── pangocairo
                   └── rsvg
                       └── v2
```

## Feel free to propose other structures / names.
