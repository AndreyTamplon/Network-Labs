package tamplon.snake.gamelogic.snake.impl;

import tamplon.snake.gamelogic.gamefield.BoundedPoint;
import tamplon.snake.gamelogic.gamefield.Direction;
import org.jetbrains.annotations.NotNull;
import tamplon.snake.gamelogic.snake.Snake;

@FunctionalInterface
public interface SnakeImplementationSupplier<SnakeType extends Snake>
{

    @NotNull SnakeType get(
            final int id,
            final @NotNull BoundedPoint head,
            final @NotNull Direction direction,
            final int size);
}
