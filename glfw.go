package OctaForce

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"runtime"
)

var (
	window *glfw.Window

	WindowWidth  = 1280
	WindowHeight = 720

	mouseButtonPrimary   = 0
	mouseButtonSecondary = 1
	mouseButtonTertiary  = 2
	mouseButtonCount     = 3

	lastTime         float64
	mouseJustPressed [3]bool
)

func initGLFW() {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)

	window, err = glfw.CreateWindow(WindowWidth, WindowHeight, "Test", nil, nil)
	if err != nil {
		glfw.Terminate()
		panic(err)
	}

	window.MakeContextCurrent()
	glfw.SwapInterval(1)
}

var glfwButtonIndexByID = map[glfw.MouseButton]int{
	glfw.MouseButton1: mouseButtonPrimary,
	glfw.MouseButton2: mouseButtonSecondary,
	glfw.MouseButton3: mouseButtonTertiary,
}

var glfwButtonIDByIndex = map[int]glfw.MouseButton{
	mouseButtonPrimary:   glfw.MouseButton1,
	mouseButtonSecondary: glfw.MouseButton2,
	mouseButtonTertiary:  glfw.MouseButton3,
}

func processEvents() {
	glfw.PollEvents()

	// Mouse Info
	mouseX, mouseY := window.GetCursorPos()
	mouseMovement = mgl32.Vec2{float32(mouseX) - mousePos.X(), float32(mouseY) - mousePos.Y()}
	mousePos = mgl32.Vec2{float32(mouseX), float32(mouseY)}
}

func DisplaySize() [2]float32 {
	w, h := window.GetSize()
	return [2]float32{float32(w), float32(h)}
}

func FramebufferSize() [2]float32 {
	w, h := window.GetFramebufferSize()
	return [2]float32{float32(w), float32(h)}
}

func postRender() {
	window.SwapBuffers()
}

// KeyPressed returns true when the Key is pressed.
func KeyPressed(key keyTyp) bool {
	return window.GetKey(glfw.Key(key)) == glfw.Press
}

type keyTyp int

const (
	KeyApostrophe   = keyTyp(glfw.KeyApostrophe)
	KeyComma        = keyTyp(glfw.KeyComma)
	KeyMinus        = keyTyp(glfw.KeyMinus)
	KeyPeriod       = keyTyp(glfw.KeyPeriod)
	KeySlash        = keyTyp(glfw.KeySlash)
	Key0            = keyTyp(glfw.Key0)
	Key1            = keyTyp(glfw.Key1)
	Key2            = keyTyp(glfw.Key2)
	Key3            = keyTyp(glfw.Key3)
	Key4            = keyTyp(glfw.Key4)
	Key5            = keyTyp(glfw.Key5)
	Key6            = keyTyp(glfw.Key6)
	Key7            = keyTyp(glfw.Key7)
	Key8            = keyTyp(glfw.Key8)
	Key9            = keyTyp(glfw.Key9)
	KeySemicolon    = keyTyp(glfw.KeySemicolon)
	KeyEqual        = keyTyp(glfw.KeyEqual)
	KeyA            = keyTyp(glfw.KeyA)
	KeyB            = keyTyp(glfw.KeyB)
	KeyC            = keyTyp(glfw.KeyC)
	KeyD            = keyTyp(glfw.KeyD)
	KeyE            = keyTyp(glfw.KeyE)
	KeyF            = keyTyp(glfw.KeyF)
	KeyG            = keyTyp(glfw.KeyG)
	KeyH            = keyTyp(glfw.KeyH)
	KeyI            = keyTyp(glfw.KeyI)
	KeyJ            = keyTyp(glfw.KeyJ)
	KeyK            = keyTyp(glfw.KeyK)
	KeyL            = keyTyp(glfw.KeyL)
	KeyM            = keyTyp(glfw.KeyM)
	KeyN            = keyTyp(glfw.KeyN)
	KeyO            = keyTyp(glfw.KeyO)
	KeyP            = keyTyp(glfw.KeyP)
	KeyQ            = keyTyp(glfw.KeyQ)
	KeyR            = keyTyp(glfw.KeyR)
	KeyS            = keyTyp(glfw.KeyS)
	KeyT            = keyTyp(glfw.KeyT)
	KeyU            = keyTyp(glfw.KeyU)
	KeyV            = keyTyp(glfw.KeyV)
	KeyW            = keyTyp(glfw.KeyW)
	KeyX            = keyTyp(glfw.KeyX)
	KeyY            = keyTyp(glfw.KeyY)
	KeyZ            = keyTyp(glfw.KeyZ)
	KeyLeftBracket  = keyTyp(glfw.KeyLeftBracket)
	KeyBackslash    = keyTyp(glfw.KeyBackslash)
	KeyRightBracket = keyTyp(glfw.KeyRightBracket)
	KeyGraveAccent  = keyTyp(glfw.KeyGraveAccent)
	KeyWorld1       = keyTyp(glfw.KeyWorld1)
	KeyWorld2       = keyTyp(glfw.KeyWorld2)
	KeyEscape       = keyTyp(glfw.KeyEscape)
	KeyEnter        = keyTyp(glfw.KeyEnter)
	KeyTab          = keyTyp(glfw.KeyTab)
	KeyBackspace    = keyTyp(glfw.KeyBackspace)
	KeyInsert       = keyTyp(glfw.KeyInsert)
	KeyDelete       = keyTyp(glfw.KeyDelete)
	KeyRight        = keyTyp(glfw.KeyRight)
	KeyLeft         = keyTyp(glfw.KeyLeft)
	KeyDown         = keyTyp(glfw.KeyDown)
	KeyUp           = keyTyp(glfw.KeyUp)
	KeyPageUp       = keyTyp(glfw.KeyPageUp)
	KeyPageDown     = keyTyp(glfw.KeyPageDown)
	KeyHome         = keyTyp(glfw.KeyHome)
	KeyEnd          = keyTyp(glfw.KeyEnd)
	KeyCapsLock     = keyTyp(glfw.KeyCapsLock)
	KeyScrollLock   = keyTyp(glfw.KeyScrollLock)
	KeyNumLock      = keyTyp(glfw.KeyNumLock)
	KeyPrintScreen  = keyTyp(glfw.KeyPrintScreen)
	KeyPause        = keyTyp(glfw.KeyPause)
	KeyF1           = keyTyp(glfw.KeyF1)
	KeyF2           = keyTyp(glfw.KeyF2)
	KeyF3           = keyTyp(glfw.KeyF3)
	KeyF4           = keyTyp(glfw.KeyF4)
	KeyF5           = keyTyp(glfw.KeyF5)
	KeyF6           = keyTyp(glfw.KeyF6)
	KeyF7           = keyTyp(glfw.KeyF7)
	KeyF8           = keyTyp(glfw.KeyF8)
	KeyF9           = keyTyp(glfw.KeyF9)
	KeyF10          = keyTyp(glfw.KeyF10)
	KeyF11          = keyTyp(glfw.KeyF11)
	KeyF12          = keyTyp(glfw.KeyF12)
	KeyF13          = keyTyp(glfw.KeyF13)
	KeyF14          = keyTyp(glfw.KeyF14)
	KeyF15          = keyTyp(glfw.KeyF15)
	KeyF16          = keyTyp(glfw.KeyF16)
	KeyF17          = keyTyp(glfw.KeyF17)
	KeyF18          = keyTyp(glfw.KeyF17)
	KeyF19          = keyTyp(glfw.KeyF19)
	KeyF20          = keyTyp(glfw.KeyF20)
	KeyF21          = keyTyp(glfw.KeyF21)
	KeyF22          = keyTyp(glfw.KeyF22)
	KeyF23          = keyTyp(glfw.KeyF23)
	KeyF24          = keyTyp(glfw.KeyF24)
	KeyF25          = keyTyp(glfw.KeyF25)
	KeyKP0          = keyTyp(glfw.KeyF25)
	KeyKP1          = keyTyp(glfw.KeyKP1)
	KeyKP2          = keyTyp(glfw.KeyKP2)
	KeyKP3          = keyTyp(glfw.KeyKP3)
	KeyKP4          = keyTyp(glfw.KeyKP4)
	KeyKP5          = keyTyp(glfw.KeyKP5)
	KeyKP6          = keyTyp(glfw.KeyKP6)
	KeyKP7          = keyTyp(glfw.KeyKP7)
	KeyKP8          = keyTyp(glfw.KeyKP8)
	KeyKP9          = keyTyp(glfw.KeyKP9)
	KeyKPDecimal    = keyTyp(glfw.KeyKPDecimal)
	KeyKPDivide     = keyTyp(glfw.KeyKPDivide)
	KeyKPMultiply   = keyTyp(glfw.KeyKPMultiply)
	KeyKPSubtract   = keyTyp(glfw.KeyKPSubtract)
	KeyKPAdd        = keyTyp(glfw.KeyKPAdd)
	KeyKPEnter      = keyTyp(glfw.KeyKPEnter)
	KeyKPEqual      = keyTyp(glfw.KeyKPEqual)
	KeyLeftShift    = keyTyp(glfw.KeyLeftShift)
	KeyLeftControl  = keyTyp(glfw.KeyLeftControl)
	KeyLeftAlt      = keyTyp(glfw.KeyLeftAlt)
	KeyLeftSuper    = keyTyp(glfw.KeyLeftSuper)
	KeyRightShift   = keyTyp(glfw.KeyRightShift)
	KeyRightControl = keyTyp(glfw.KeyRightControl)
	KeyRightAlt     = keyTyp(glfw.KeyRightAlt)
	KeyRightSuper   = keyTyp(glfw.KeyRightSuper)
	KeyMenu         = keyTyp(glfw.KeyMenu)
	KeyLast         = keyTyp(glfw.KeyLast)
)

// MouseButtonPressed returns true when the mouse key is pressed.
func MouseButtonPressed(button mouseButtonTyp) bool {
	return window.GetMouseButton(glfw.MouseButton(button)) == glfw.Press
}

type mouseButtonTyp int

const (
	MouseButton2      = mouseButtonTyp(glfw.MouseButton2)
	MouseButton3      = mouseButtonTyp(glfw.MouseButton3)
	MouseButton4      = mouseButtonTyp(glfw.MouseButton4)
	MouseButton5      = mouseButtonTyp(glfw.MouseButton5)
	MouseButton6      = mouseButtonTyp(glfw.MouseButton6)
	MouseButton7      = mouseButtonTyp(glfw.MouseButton7)
	MouseButton8      = mouseButtonTyp(glfw.MouseButton8)
	MouseButtonLast   = mouseButtonTyp(glfw.MouseButtonLast)
	MouseButtonLeft   = mouseButtonTyp(glfw.MouseButtonLeft)
	MouseButtonRight  = mouseButtonTyp(glfw.MouseButtonRight)
	MouseButtonMiddle = mouseButtonTyp(glfw.MouseButtonMiddle)
)

var mousePos mgl32.Vec2

// GetMousePos returns the position of the Mouse relative to the screen.
func GetMousePos() mgl32.Vec2 {
	return mousePos
}

var mouseMovement mgl32.Vec2

// GetMouseMovement returns the relative movement.
func GetMouseMovement() mgl32.Vec2 {
	return mouseMovement
}
