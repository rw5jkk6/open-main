import pygame

pygame.init()
SURFACE = pygame.display.set_mode((600,600))
FPSCLOCK = pygame.time.Clock()


def paint():
    SURFACE.fill((0,0,0))

    pygame.display.update()

def main():
    while True:
        
        paint()
        FPSCLOCK.tick(5)

if __name__ == "__main__":
    main()
