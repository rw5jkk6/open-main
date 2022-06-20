import sys, datetime
from termios import PARODD
import pygame
from pygame.locals import QUIT

def tick():
    for event in pygame.event.get():
        if event.type == QUIT:
            pygame.quit()
            sys.exit()

pygame.init()
pygame.key.set_repeat(5,5)
SURFACE=pygame.display.set_mode((600,800))
FPSCLOCK=pygame.time.Clock()

def main():
    fps = 30
    while True:
        tick()

        SURFACE.fill((0,0,0))

        pygame.display.update()
        FPSCLOCK.tick(fps)

if __name__ == "__main__":
    main()
