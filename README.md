## What is this?
This is a simple program to convert a 3d-model .obj file to a lua script-file like this:

model={
vertices={
{x=-0.182532,y=1.294653,z=-0.580947,w=1},
{x=-0.182532,y=1.294653,z=-0.580947,w=1},
{x=-0.182532,y=1.294653,z=-0.580947,w=1},
....
},
faces={
{1,2,4,3},
{3,4,8,7},
{7,8,6,5},
....
}
}

I made this so I could easily import 3d-models to Pico-8.
Currently only the above data (vertices and faces) are exported as I only care about those but I will maybe add support for texture coordinates etc. in the future.

## Instructions

To use simply run the program with the filepath of the obj-file like so "go run ./src /home/me/3dmodel.obj"