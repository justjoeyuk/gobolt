dep-windows:
	pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-SDL2{,_mixer,_image,_ttf}

dep-linux-ubuntu:
	apt install libsdl2-dev libsdl2-mixer-dev libsdl2-image-dev libsdl2-ttf-dev libsdl2-gfx-dev

dep-linux-fedora:
	yum install SDL2-devel SDL2_mixer-devel SDL2_image-devel SDL2_ttf-devel SDL2_gfx-devel

dep-linux-arch:
	pacman -S sdl2{,_mixer,_image,_ttf,_gfx}

dep-linux-gentoo:
	emerge -av libsdl2 sdl2-{gfx,image,mixer,ttf}

dep-osx:
	xcode-select --install
	brew install sdl2{,_image,_ttf,_mixer} pkg-config