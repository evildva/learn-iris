#include <iostream>
#include "SDL2/SDL.h"

using namespace std;

int main(int argc, char* argv[]){
	SDL_Window *window;

	if (SDL_Init(SDL_INIT_VIDEO|SDL_INIT_AUDIO) != 0) {
        SDL_Log("Unable to initialize SDL: %s", SDL_GetError());
        return 1;
    }

    window = SDL_CreateWindow(
        "An SDL2 window",                  // window title
        SDL_WINDOWPOS_UNDEFINED,           // initial x position
        SDL_WINDOWPOS_UNDEFINED,           // initial y position
        640,                               // width, in pixels
        480,                               // height, in pixels
        SDL_WINDOW_OPENGL                  // flags - see below
    );

    // Check that the window was successfully created
    if (window == NULL) {
        // In the case that the window could not be made...
        printf("Could not create window: %s\n", SDL_GetError());
        return 1;
    }

    int quit = 1;
    do {
        SDL_Event event;
        SDL_WaitEvent(&event);
        switch (event.type) {
        case SDL_QUIT:
            SDL_Log("Event type is %d", event.type);
            quit = 0;
        default:
            SDL_Log("Event type is %d", event.type);
            break;
        }
    } while (quit);

    // Close and destroy the window
    SDL_DestroyWindow(window);
    SDL_Quit();

	return 0;
}