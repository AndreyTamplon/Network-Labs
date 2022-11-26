package tamplon.snake.gui.util;

import java.awt.*;
import java.util.Collection;
import java.util.HashSet;
import java.util.concurrent.ThreadLocalRandom;

public final class Colours
{
    public static final Color BACKGROUND_COLOUR = new Color(43, 43, 43);
    public static final Color FOREGROUND_COLOUR = new Color(212, 212, 212);

    public static final Color NAME_COLOUR = FOREGROUND_COLOUR;
    public static final Color DEAD_SNAKE_COLOUR = FOREGROUND_COLOUR;
    public static final Color RED = new Color(199, 84, 80);
    public static final Color GREEN = new Color(73, 156, 84);
    public static final Color LIGHT_GRAY = new Color(135, 147, 154);
    public static final Color BLUE = new Color(53, 146, 196);
    public static final Color YELLOW = new Color(244, 175, 61);

    public static final Color INTERFACE_BACKGROUND = new Color(255, 255, 255);
    public static final Color GAME_PANEL_BACKGROUND = new Color(170, 255, 192);
    public static final Color LIGHT_LINING = new Color(206, 255, 178);
    public static final Color LINING = new Color(81, 81, 81);
    public static final Color DARK_LINING = new Color(203, 255, 187);
    public static final Color TEXT = new Color(0, 0, 0);
    public static final Color SCROLL_THUMB = new Color(78, 78, 78);
    public static final Color TEXT_ENTRY_FORM = new Color(255, 255, 255);
    public static final Color TOOLTIP = new Color(75, 77, 75);

    private static final Collection<Color> snakeColours = new HashSet<>();

    static
    {
        snakeColours.add(Color.CYAN);
        snakeColours.add(Color.MAGENTA);
        snakeColours.add(LIGHT_GRAY);
        snakeColours.add(BLUE);
        snakeColours.add(YELLOW);
    }

    private Colours()
    {
        throw new UnsupportedOperationException("This is a utility class and cannot be instantiated");
    }

    public static Color getRandomColour()
    {
        var num = ThreadLocalRandom.current().nextInt(0, snakeColours.size());
        for (final var t : snakeColours)
        {
            --num;
            if (num < 0)
            {
                return t;
            }
        }
        throw new AssertionError();
    }
}
