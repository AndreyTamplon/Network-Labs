package tamplon.snake.gamelogic.snake;

import tamplon.snake.gamelogic.gamefield.BoundedPoint;
import tamplon.snake.gamelogic.gamefield.Direction;
import org.jetbrains.annotations.NotNull;

import java.util.function.Consumer;

public interface SnakeInfo
{

    @NotNull BoundedPoint getHead();

    void forEachSegment(final @NotNull Consumer<@NotNull BoundedPoint> action);

    int getPlayerId();

    @NotNull Direction getDirection();

    boolean isZombie();
}
