# blender

Compiled while working with Blender 2.92.0 on Mac OS X.


## 3D viewport

### Navigation

- `Scroll` - Orbit
- `Ctrl + scroll` - Zoom in/out
- `Shift + scroll` - Pan
- `Shift` + `~` - walk camera in View Camera mode (`~` -> `View Camera`)
    - to walk: `W`, `A`, `S`, `D`

### Move objects

- `G`, `R`, `S` - Move/rotate/scale
  - ...along axis: `X`, `Y`, `Z`
  - ...constrain to plane: `Shift` + excluded axis
    - clear: `C`
  - with precision movement: hold `Shift`
  - with incremental movement: hold `Ctrl`
  - cancel: `Esc` or `R click`

### Select/modify object(s)

- select
  - `L click` - select object
    - to select multiple: hold `Shift`
  - `A` - select all
  - `A A` - deselect all

Once selected:
- modify
  - `Ctrl + A` - Apply
    - `All Transforms`: normalize location/rotation/scale
- hide <!-- https://docs.blender.org/manual/en/latest/scene_layout/object/editing/show_hide.html -->
  - `H` - hide all selected objects
  - `Shift` + `H` - hide all **un**selected objects
  - `Alt` + `H` - show hidden objects


### Create/delete objects

- `Shift` + `A` - Add
- `X` - Delete
- `Shift` + `D` - Duplicate

### 3D cursor

- where new objects are created; [docs](https://docs.blender.org/manual/en/latest/editors/3dview/3d_cursor.html)
- `Shift` + `R click` - change cursor position
- `Shift` + `C` - reset cursor

### View

- `~` - View pie
  - Focus on different parts of the scene (e.g.: selected object, camera)
  - <img src="https://user-images.githubusercontent.com/20177171/140260267-8590db04-87be-4677-879f-be633735f34a.png" width="300px" />
- `Z` - Viewpoint Shading pie
  - Change how objects are shaded/displayed in the 3D viewer.
  - Can also change with these buttons (top right): <img src="https://user-images.githubusercontent.com/20177171/158446450-a7b56232-1f25-46ae-9db0-813be1b59d62.png" width="80px" />
  - `Shift` + `Z` - wireframe mode
  - `Alt` + `Z` - X-ray mode

### Switch between modes

- `Ctrl` + `Tab` - switch modes pie
  - <img src="https://user-images.githubusercontent.com/20177171/158446899-fc4f875e-e343-48b4-afea-8b61002519a5.png" width="300px" />
- `Tab` - toggle Edit mode

### Misc

- `T` - Toggle Toolbar
- `N` - Toggle Context
<!-- -->
- `F12` - Render


### Edit mode

This subsection contains shortcuts specific to "Edit mode".



## (default) Render engines

* Eevee
  - Real-time, but less realistic
  - New in Blender 2.8
* Cycles
  - Raytracing engine; slower but more realistic


## Materials



## Notes

`Alt` === `Option`


