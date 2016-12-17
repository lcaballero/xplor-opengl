package main

import "github.com/go-gl/mathgl/mgl32"

var vertices  = []float32{
	-0.5, -0.5, -0.5,  0.0, 0.0,
	0.5, -0.5, -0.5,  1.0, 0.0,
	0.5,  0.5, -0.5,  1.0, 1.0,
	0.5,  0.5, -0.5,  1.0, 1.0,
	-0.5,  0.5, -0.5,  0.0, 1.0,
	-0.5, -0.5, -0.5,  0.0, 0.0,

	-0.5, -0.5,  0.5,  0.0, 0.0,
	0.5, -0.5,  0.5,  1.0, 0.0,
	0.5,  0.5,  0.5,  1.0, 1.0,
	0.5,  0.5,  0.5,  1.0, 1.0,
	-0.5,  0.5,  0.5,  0.0, 1.0,
	-0.5, -0.5,  0.5,  0.0, 0.0,

	-0.5,  0.5,  0.5,  1.0, 0.0,
	-0.5,  0.5, -0.5,  1.0, 1.0,
	-0.5, -0.5, -0.5,  0.0, 1.0,
	-0.5, -0.5, -0.5,  0.0, 1.0,
	-0.5, -0.5,  0.5,  0.0, 0.0,
	-0.5,  0.5,  0.5,  1.0, 0.0,

	0.5,  0.5,  0.5,  1.0, 0.0,
	0.5,  0.5, -0.5,  1.0, 1.0,
	0.5, -0.5, -0.5,  0.0, 1.0,
	0.5, -0.5, -0.5,  0.0, 1.0,
	0.5, -0.5,  0.5,  0.0, 0.0,
	0.5,  0.5,  0.5,  1.0, 0.0,

	-0.5, -0.5, -0.5,  0.0, 1.0,
	0.5, -0.5, -0.5,  1.0, 1.0,
	0.5, -0.5,  0.5,  1.0, 0.0,
	0.5, -0.5,  0.5,  1.0, 0.0,
	-0.5, -0.5,  0.5,  0.0, 0.0,
	-0.5, -0.5, -0.5,  0.0, 1.0,

	-0.5,  0.5, -0.5,  0.0, 1.0,
	0.5,  0.5, -0.5,  1.0, 1.0,
	0.5,  0.5,  0.5,  1.0, 0.0,
	0.5,  0.5,  0.5,  1.0, 0.0,
	-0.5,  0.5,  0.5,  0.0, 0.0,
	-0.5,  0.5, -0.5,  0.0, 1.0,
};

var cubePositions = []mgl32.Vec3{
	mgl32.Vec3{ 0.0,  0.0,  0.0},
	mgl32.Vec3{ 2.0,  5.0, -15.0},
	mgl32.Vec3{-1.5, -2.2, -2.5},
	mgl32.Vec3{-3.8, -2.0, -12.3},
	mgl32.Vec3{ 2.4, -0.4, -3.5},
	mgl32.Vec3{-1.7,  3.0, -7.5},
	mgl32.Vec3{ 1.3, -2.0, -2.5},
	mgl32.Vec3{ 1.5,  2.0, -2.5},
	mgl32.Vec3{ 1.5,  0.2, -1.5},
	mgl32.Vec3{-1.3,  1.0, -1.5},
}