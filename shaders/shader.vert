#version 410 core

layout (location = 0) in vec3 position;

uniform mat4 ortho;

void main() {
	gl_Position = ortho * vec4(position.x, -position.y, position.z, 1.0);
}
