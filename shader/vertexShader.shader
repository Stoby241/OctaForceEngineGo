#version 430

layout(location = 0) uniform mat4 projection;
layout(location = 1) uniform mat4 camera;

layout(location = 2) uniform mat4 transform;
layout(location = 3) uniform vec3 inColor;

layout(location = 0) in vec3 vertexPosition;

out vec3 color;

void main() {
    color = inColor;
    gl_Position = projection * camera * transform * vec4(vertexPosition, 1);
}
