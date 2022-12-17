package tamplon.snake.gamelogic.snake;

import tamplon.snake.gamelogic.gamefield.Direction;
import org.jetbrains.annotations.NotNull;

@FunctionalInterface
public interface Steerable
{

    void changeDirection(final @NotNull Direction direction);
}
