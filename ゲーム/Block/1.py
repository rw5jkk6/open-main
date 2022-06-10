import sys
from termios import PARODD
import pygame
from pygame.locals import QUIT, KEYDOWN, K_LEFT, K_RIGHT, Rect

class Block:
    def __init__(self, col, rect):
        self.col = col
        self.rect = rect

    def draw(self):
        pygame.draw.rect(SURFACE, self.col, self.rect)

def tick():
    global BLOCKS
    for event in pygame.event.get():
        if event.type == QUIT:
            pygame.quit()
            sys.exit()
        elif event.type == KEYDOWN:
            if event.key == K_LEFT:
                PADDLE.rect.centerx -= 10
            elif event.key == K_RIGHT:
                PADDLE.rect.centerx += 10

pygame.init()
pygame.key.set_repeat(5,5)
SURFACE=pygame.display.set_mode((600,800))
FPSCLOCK=pygame.time.Clock()
PADDLE= Block((242, 242, 0), Rect(300, 700, 100, 30))

def main():
    fps = 30
    while True:
        tick()

        SURFACE.fill((0,0,0))
        PADDLE.draw()

        pygame.display.update()
        FPSCLOCK.tick(fps)

if __name__ == "__main__":
    main()
