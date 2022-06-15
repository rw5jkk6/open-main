import sys
import time
import pygame
from pygame.locals import QUIT, Rect

pygame.init()
SURFACE = pygame.display.set_mode((600,600))
FPSCLOCK = pygame.time.Clock()
score = time.time()

(W,H)=(20,20)

def paint(mf):
    SURFACE.fill((0,0,0))
    mess_score = mf.render('{:.1f}'.format(time.time()-score), 
                            True, (255, 255, 0))
    SURFACE.blit(mess_score, (500, 10))
    pygame.display.update()

def main():
    myfont=pygame.font.SysFont(None,60)
    while True:
        for event in pygame.event.get():
            if event.type == QUIT:
                pygame.quit()
                sys.exit()
        
        paint(myfont)
        
        FPSCLOCK.tick(5)

if __name__ == "__main__":
    main()

