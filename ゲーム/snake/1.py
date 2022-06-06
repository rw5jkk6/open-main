import sys
import pygame
from pygame.locals import QUIT, Rect

pygame.init()
SURFACE = pygame.display.set_mode((600,600))
FPSCLOCK = pygame.time.Clock()

(W,H)=(20,20)

def paint():
    SURFACE.fill((0,0,0))
    for index in range(20):
        # 縦線が描かれる
        pygame.draw.line(SURFACE, (64, 64, 64), (index*30, 0),(index*30, 600))
        # 横線が描かれる
        pygame.draw.line(SURFACE, (64, 64, 64), (0, index*30),(600, index*30))

    pygame.display.update()

def main():
    while True:
        for event in pygame.event.get():
            if event.type == QUIT:
                pygame.quit()
                sys.exit()
        
        paint()
        FPSCLOCK.tick(5)

if __name__ == "__main__":
    main()