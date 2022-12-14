package tamplon.snake.gamelogic.snake;

import tamplon.snake.gamelogic.gamefield.BoundedPoint;
import org.jetbrains.annotations.NotNull;

import java.util.function.Function;

public interface Snake extends Steerable, SnakeInfo
{

    void zombify();

    void move(final @NotNull Function<@NotNull BoundedPoint, @NotNull Boolean> isFood);
}
