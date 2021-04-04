#version 430

layout(location = 0) uniform mat4 projection;
layout(location = 1) uniform mat4 camera;

layout(location = 0) in vec3 vertexPosition;
layout(location = 1) in vec3 instanceColor;
layout(location = 2) in vec4 transformX;
layout(location = 3) in vec4 transformY;
layout(location = 4) in vec4 transformZ;
layout(location = 5) in vec4 transformS;

out vec3 color;

void main() {
    color = instanceColor;
    gl_Position =
        projection *
        camera *
        mat4(transformX, transformY, transformZ, transformS) *
        vec4(vertexPosition, 1);
}
